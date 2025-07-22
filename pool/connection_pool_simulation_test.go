package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func connectToService() interface{} {
	time.Sleep(time.Second)
	return struct{}{}
}

func warmCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New)
	}
	return p
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		p := warmCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatal(err)
		}
		defer server.Close()
		wg.Done()
		for {
			accept, err := server.Accept()
			if err != nil {
				log.Println("Error accepting connection:", err)
				continue
			}
			conn := p.Get()
			fmt.Println("Connection established:", conn)
			p.Put(conn)
			accept.Close()
		}
	}()
	return &wg
}

func init() {
	daemon := startNetworkDaemon()
	daemon.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
