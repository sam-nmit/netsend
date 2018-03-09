package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var (
	listen = flag.Int("l", 0, "Listen [-l 6000]")
	web    = flag.Int("w", 0, "Web [-w 127.0.0.1:6000]")
	send   = flag.String("s", "", "send [-s 127.0.0.1:6000]")
)

func listenerToPort(l *net.Listener) int {
	return (*l).Addr().(*net.TCPAddr).Port
}

func main() {

	flag.Parse()
	if *listen != 0 { //TCP Receving
		l, _ := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *listen))
		c, _ := l.Accept()
		b, _ := ioutil.ReadAll(c)
		fmt.Print(string(b))
	} else if *web != 0 { //WEB Receving
		pipe, _ := ioutil.ReadAll(os.Stdin)
		sPipe := string(pipe)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, sPipe)
		})
		l, _ := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *web))
		fmt.Println(listenerToPort(&l))
		log.Fatal(http.Serve(l, nil))
	} else if *send != "" { //Sending
		lAddr := *send
		if !strings.Contains(lAddr, ":") {
			lAddr = fmt.Sprintf("127.0.0.1:%s", lAddr)
		}
		c, err := net.Dial("tcp", lAddr)
		if err != nil {
			log.Fatal("Failed to connect.")
		} else {
			io.Copy(c, os.Stdin)
			c.Close()
			log.Printf("Sent.")
		}
	}
}
