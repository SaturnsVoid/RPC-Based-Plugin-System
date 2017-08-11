// Plugin System project main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Plugin Controler")
	err := exec.Command("./plugins/get.plugin").Start()
	if err != nil {
		log.Fatal("Cannot start plugin", err)
	}
	time.Sleep(1 * time.Second)
	client, err := rpc.Dial("tcp", "127.0.0.1:55555")
	if err != nil {
		log.Fatal("Cannot create RPC client: ", err)
	}
	for {
		fmt.Print("Command: ")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		if scan.Text() == "get" {
			var Data string
			err = client.Call("Plugin.Get", "http://checkip.amazonaws.com", &Data)
			if err != nil {
				log.Fatal("Error calling Revert: ", err)
			}
			fmt.Println("Plugin:", Data)
		} else if scan.Text() == "exit" {
			var n int
			client.Call("Plugin.Exit", 0, &n)
			fmt.Println("Plugin Closed, Closing Self.")
			os.Exit(0)
		}
	}
}
