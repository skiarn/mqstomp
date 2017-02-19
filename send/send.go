package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/go-stomp/stomp"
	"github.com/go-stomp/stomp/frame"
)

var serverAddr = flag.String("server", "localhost:61613", "STOMP server.")
var message = flag.String("message", "", "Message to to sent")
var queueName = flag.String("queue", "/queue/test", "Destination queue")

var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
	stomp.ConnOpt.Login("admin", "admin"),
	stomp.ConnOpt.Host("/"),
}
var headers headersFlags

func main() {
	flag.Var(&headers, "header", "Headers for message. Example JMSXGroupID:group1")
	flag.Parse()
	if *message == "" {
		log.Fatal("Message required.")
	}
	headers.Set("persistent=true")
	headers.Set("JMSXGroupID=queue1")

	conn, err := stomp.Dial("tcp", *serverAddr, options...)
	if err != nil {
		log.Fatal("cannot connect to server", err.Error())
	}

	headers := headers.Get()
	err = conn.Send(*queueName, "text/plain",
		[]byte(*message),
		//stomp.SendOpt.Header("CustomHeader1", "H1"),
		//stomp.SendOpt.Header("CustomHeader2", "H2"),
		//stomp.SendOpt.Header("persistent", "true"),
		//stomp.SendOpt.Header("JMSXGroupID", "queue1"),
		headers...,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sent message: %s \n", *message)
	conn.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}

type headersFlags []header

type header struct {
	Header string
	Value  string
}

func (hf *headersFlags) Get() (r []func(*frame.Frame) error) {
	for _, h := range *hf {
		f := func(*frame.Frame) error {
			return stomp.SendOpt.Header(h.Header, h.Value)
		}
		r = append(r, f)
	}
	return

}
func (hf *headersFlags) String() string {
	return "Headers for message."
}

func (hf *headersFlags) Set(value string) error {
	h := strings.Split(value, "=")
	if len(h) != 2 {
		log.Fatalf("Expected header: %s have format: header1=value1.", value)
	}
	*hf = append(*hf, header{h[0], h[1]})
	return nil
}
