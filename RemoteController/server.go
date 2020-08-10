package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	// _ "github.com/go-sql-driver/mysql"
)

// Args struct that contains input data
type Args struct {
	A, B int
}

// Arith type to bind the RPC functions
type Arith int

// Time returns the local time on the VM
func (t *Arith) Time(args *Args, reply *string) error {
	fmt.Println("RPC Time")
	*reply = string(time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

// CPUUsage returns CPU Usage(%)
func (t *Arith) CPUUsage(args *Args, reply *string) error {
	fmt.Println("RPC CPUUsage")
	percent, _ := cpu.Percent(time.Second, false)
	*reply = strings.Trim(strings.Replace(fmt.Sprint(percent), " ", ", ", -1), "[]")
	return nil
}

// MemStats returns memory consumption (it should do that, but not yet)
func (t *Arith) MemStats(args *Args, reply *string) error {
	fmt.Println("RPC MemStats")
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)

	*reply = strconv.FormatUint(mem.TotalAlloc, 10)

	return nil
}

func main() {

	// db, err := sql.Open("mysql",
	// 	"zimmer:test1234@tcp(127.0.0.1:3306)/gotest")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// Binds the RPC function with an int
	arith := new(Arith)
	rpc.Register(arith)

	// Launches the server
	tcpAddr, err := net.ResolveTCPAddr("tcp", "192.168.1.8:8888")
	checkError(err)

	// Listens to the binded port
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// Accepts new connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
