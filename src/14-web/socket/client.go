package main


func main()  {
	r := Request{
		Ip: "10.253.8.91",
		SshPort: "22",
		Command: "ls /root",
	}
	sendMessage(&r)
}