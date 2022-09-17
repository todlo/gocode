package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"golang.org/x/term"
)

var (
	target           = flag.String("target", "", "Hostname or address of the pihole you want to update.")
	port             = flag.String("port", "22", "The SSH port number on the target for this ssh session.")
	sshdir           = flag.String("sshdir", "", "Location of ssh keys for pubkey authentication.")
	keyfile          = flag.String("keyfile", "", "File name of private ssh key.")
	user             = flag.String("user", "", "The username that we're trying to log in as.")
	yubikey          = flag.Bool("yubikey", true, "If false, use 6-digit authenticator code as 2fa.")
	ignoreKnownHosts = flag.Bool("ignore-known-hosts", false, "[INSECURE] Ignore parsing of known hosts file.")
)

type sshcon struct {
	sc         *ssh.Client
	scc        *ssh.ClientConfig
	pp, pw, sd []byte
}

func main() {
	flag.Parse()
	var s sshcon
	if err := s.connect(*target, *port, *sshdir); err != nil {
		log.Fatal(err)
	}
	defer s.sc.Close()

	session, err := s.sc.NewSession()
	if err != nil {
		log.Fatalf("unable to create session %v", err)
	}
	defer session.Close()

	// Set IO
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	in, err := session.StdinPipe()
	if err != nil {
		log.Fatal("error with stdinpipe: ", err)
	}

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}
	if err := session.Start("sudo pihole -up"); err != nil {
		fmt.Println("failed to run command: ", err)
	}
	io.Copy(in, bytes.NewBufferString(string(s.pw)+"\n"))
	time.Sleep(1 * time.Second)
	if *yubikey {
		fmt.Printf("\nPlease touch yubikey again...\n")
	}
	if !*yubikey {
		io.Copy(in, bytes.NewBufferString(string(s.sd)+"\n"))
		time.Sleep(1 * time.Second)
	}
	err = session.Wait()
	if err != nil {
		log.Fatalf("wait err: %v", err)
	}
	session.Signal(ssh.SIGINT)
}

func (s *sshcon) connect(host, port, sshdir string) error {
	if sshdir == "" {
		hd, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		sshdir = hd + "/.ssh"
	}
	key, err := os.ReadFile(filepath.Join(sshdir, *keyfile))
	if err != nil {
		return err
	}

	s.pp, err = getcreds("Enter ssh passphrase: ")
	if err != nil {
		return err
	}
	fmt.Println()
	s.pw, err = getcreds("Enter sudo password: ")
	if err != nil {
		return err
	}
	fmt.Println()
	if !*yubikey {
		s.sd, err = getcreds("Please enter 6-digit code: ")
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Please touch yubikey...")
	}
	signer, err := ssh.ParsePrivateKeyWithPassphrase(key, s.pp)
	if err != nil {
		return err
	}

	cb := func(user, instruction string, questions []string, echos []bool) ([]string, error) {
		if len(questions) < 1 {
			return nil, nil
		}
		var answers []string
		fmt.Printf("\n%s", questions[0])
		num, err := term.ReadPassword(0)
		if err != nil {
			return nil, err
		}
		fmt.Println()
		answers = append(answers, string(num))
		return answers, nil
	}
	hostKeyCb, err := hccb(sshdir, *ignoreKnownHosts)
	if err != nil {
		return err
	}
	s.scc = &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
			ssh.RetryableAuthMethod(ssh.KeyboardInteractiveChallenge(cb), 3),
		},
		HostKeyCallback:   hostKeyCb,
		HostKeyAlgorithms: []string{ssh.KeyAlgoED25519},
	}

	s.sc, err = ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), s.scc)
	if err != nil {
		return err
	}

	return err
}

func hccb(sshdir string, ignore bool) (ssh.HostKeyCallback, error) {
	if ignore {
		return ssh.InsecureIgnoreHostKey(), nil
	}
	return knownhosts.New(sshdir + "/known_hosts")
}

func getcreds(instruction string) ([]byte, error) {
	fmt.Print(instruction)
	pw, err := term.ReadPassword(0)
	if err != nil {
		return nil, err
	}
	return []byte(pw), nil
}
