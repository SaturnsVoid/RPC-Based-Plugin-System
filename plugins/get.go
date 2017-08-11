// Plugin System project main.go
package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Plugin struct {
	listener net.Listener
}

func (p Plugin) Get(arg string, ret *string) error {
	resp, err := http.Get(arg)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	*ret = string(html)
	return nil
}
func (p Plugin) Exit(arg int, ret *int) error {
	os.Exit(0)
	return nil
}

func main() {
	p := &Plugin{}
	err := rpc.Register(p)
	if err != nil {
		log.Fatal("Cannot register plugin: ", err)
	}
	p.listener, err = net.Listen("tcp", "127.0.0.1:55555")
	if err != nil {
		log.Fatal("Cannot listen: ", err)
	}
	rpc.Accept(p.listener)
}
