package main

import (
	"fmt"
	"net/rpc"
	"os"
	"sync"
)

// Args struct that contains input data
type Args struct {
	A, B int
}

// Connection struct that has all the details of a connection
type Connection struct {
	service string
	client  *rpc.Client
	wg      sync.WaitGroup
}

// Init returns a new connecton, receives a `ip:port` string
func Init(serv string) *Connection {
	return &Connection{service: serv}
}

// Bind creates a connection with the previously provided data
func (conn *Connection) Bind() error {
	var err error
	conn.client, err = rpc.Dial("tcp", conn.service)

	return err
}

// RPCall calss the remote function
func (conn *Connection) RPCall(funcName string, args interface{}, result interface{}) {
	conn.client.Call(funcName, args, result)
	conn.wg.Done()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}

	conn := Init(os.Args[1])
	err := conn.Bind()
	if err != nil {
		fmt.Printf("RPC init failed")
	}

	args := Args{17, 8}

	var reply1, reply2, reply3 string

	conn.wg.Add(1)
	go conn.RPCall("Arith.Time", args, &reply1)
	conn.wg.Wait()

	conn.wg.Add(1)
	go conn.RPCall("Arith.CPUUsage", args, &reply2)
	conn.wg.Wait()

	conn.wg.Add(1)
	go conn.RPCall("Arith.MemStats", args, &reply3)
	conn.wg.Wait()

	fmt.Printf("Arith.Time: %s\n", reply1)
	fmt.Printf("Arith.CpuUsage: %s\n", reply2)
	fmt.Printf("Arith.MemStats: %s\n", reply3)

}
