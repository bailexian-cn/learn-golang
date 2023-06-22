```go
func exeCommand(ip string, port int, password, command string) (string, error) {
	sshConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password(tc.RootPW),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	hostAddress := fmt.Sprintf("%s:%d", ip, port)

	connection, err := ssh.Dial("tcp", hostAddress, sshConfig)
	if err != nil {
		return "", fmt.Errorf("ssh: failed to dial IP %s, error: %s", hostAddress, err.Error())
	}

	session, err := connection.NewSession()
	if err != nil {
		return "", fmt.Errorf("ssh: failed to create session, error: %s", err.Error())
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		return "", fmt.Errorf("ssh: failed to run command `%s`, error: %s", command, err.Error())
	}

	output := strings.Trim(b.String(), "\n")
	return output, nil
}
```

