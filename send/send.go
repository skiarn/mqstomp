package main

import (
	"flag"
	"log"
	"strings"

	"github.com/go-stomp/stomp"
	"github.com/go-stomp/stomp/frame"
)

var serverAddr = flag.String("server", "localhost:61613", "STOMP server")
var message = flag.String("message", "", "Message to to sent")
var queueName = flag.String("queue", "/queue/test", "Destination queue")
var login = flag.String("login", "admin=admin", "login user=pwd credentials")
var host = flag.String("host", "/", "host header")

var headers headersFlags

func main() {
	flag.Var(&headers, "header", "Headers for message. -header=\"CusomHeader=Value\"")
	flag.Parse()
	if *message == "" {
		log.Fatal("Message required.")
	}

	conn, err := stomp.Dial("tcp", *serverAddr, connOptions()...)
	if err != nil {
		log.Fatal("cannot connect to server", err.Error())
	}

	err = conn.Send(*queueName, "text/plain",
		[]byte(*message),
		headers.Get()...,
	)
	if err != nil {
		log.Fatal(err)
	}
	conn.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}

func connOptions() []func(*stomp.Conn) error {
	l := strings.Split(*login, "=")
	if len(l) != 2 {
		log.Fatalf("Expected connection option login: %s have format: user=pwd.", *login)
	}
	return []func(*stomp.Conn) error{
		stomp.ConnOpt.Login(l[0], l[1]),
		stomp.ConnOpt.Host(*host),
	}
}

type headersFlags []string

func (hf *headersFlags) Get() (r []func(*frame.Frame) error) {
	for _, h := range *hf {
		header := strings.Split(h, "=")
		f := stomp.SendOpt.Header(header[0], header[1])
		r = append(r, f)
	}
	return r
}

func (hf headersFlags) String() string {
	return strings.Join(hf[:], ",")
}

func (hf *headersFlags) Set(value string) error {
	h := strings.Split(value, "=")
	if len(h) != 2 {
		log.Fatalf("Expected header: %s have format: header1=value1.", value)
	}
	*hf = append(*hf, value)
	return nil
}
