package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	addr = flag.String("addr", ":8080", "listen addr")
)

func main() {
	flag.Parse()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received", r.URL.Path)
		ip, err := GetOutboundIP()
		if err != nil {
			log.Println("error :", err)
		}
		w.Write([]byte(ip + "\n"))
		w.Write([]byte("pong\n"))
	})

	go func() {
		log.Println("listening on", *addr)
		http.ListenAndServe(*addr, nil)
	}()

	log.Println("start")

	<-sigChan

	log.Println("stop")
}

func GetOutboundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	_ = conn.Close()

	return localAddr.IP.String(), nil
}
