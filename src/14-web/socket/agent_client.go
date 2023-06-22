package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

var (
	exitSemaphore chan bool
)

type Result string

const (
	Success = Result("success")
	Failed  = Result("failed")
)

const (
	UsbAgentSock = "/var/lib/vm-command-agent/test/agent.sock"
)

type Message struct {
	Bus     string `json:"bus"`
	Device  string `json:"device"`
	USBPort string `json:"usb_port"`
	Port    string `json:"port"`
	Action  string `json:"action"`
}

type Response struct {
	Result        Result `json:"result,omitempty"`
	FailedMessage string `json:"failedMessage,omitempty"`
	Output        string `json:"output,omitempty"`
}

type Request struct {
	Ip      string `json:"ip,omitempty"`
	SshPort string `json:"sshPort,omitempty"`
	Command string `json:"command"`
}

func sendMessage(message *Request) (*Response, error) {
	unixSock := UsbAgentSock
	uaddr, err := net.ResolveUnixAddr("unix", unixSock)
	if err != nil {
		return nil, err
	}

	// Connect server with unix socket
	uconn, err := net.DialUnix("unix", nil, uaddr)
	if err != nil {
		return nil, err
	}

	// Close unix socket when exit this function
	defer uconn.Close()

	// Wait to receive response
	response := Response{}
	go onMessageReceived(uconn, &response)

	msg, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	// send message by unix sock
	_, err = uconn.Write(msg)
	if err != nil {
		return nil, err
	}

	// Wait server response
	// change this duration bigger than server sleep time to get correct response
	exitSemaphore = make(chan bool)
	select {
	case <-time.After(time.Duration(2) * time.Second):
		log.Println(fmt.Sprintf(" Wait response timeout"))
	case <-exitSemaphore:
	}

	close(exitSemaphore)

	return &response, nil
}

func onMessageReceived(conn *net.UnixConn, response *Response) {
	// Read information from response
	data, err := parseResponse(conn)

	if err != nil {
		log.Println(err)
	} else {
		if err = json.Unmarshal(data, response); err != nil {
			log.Println(err)
		}
	}

	// Exit when receive data from server
	exitSemaphore <- true
}

func parseResponse(conn *net.UnixConn) ([]byte, error) {
	buf := make([]byte, 20960)
	if nr, err := conn.Read(buf); err != nil {
		return nil, err
	} else {
		return buf[0:nr], nil
	}
}
