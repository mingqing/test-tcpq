package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	// 监听或者连接的地址
	listenAddr string
	readData   bool
	writeData  bool
	interval   int
)

func init() {
	flag.StringVar(&listenAddr, "l", "127.0.0.1:4444", "set listen address")
	flag.BoolVar(&readData, "r", false, "open read data")
	flag.BoolVar(&writeData, "w", false, "open write data")
	flag.IntVar(&interval, "i", 1000*1000*1000, "1(second) = 1000 * 1000 * 1000(nanosecond)")
	flag.Parse()
}

func main() {
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for {
			if writeData || readData {
				if writeData {
					conn.Write([]byte("0123456789"))
					fmt.Println("send 10 byte to", conn.RemoteAddr())
				}
				if readData {
					buf := make([]byte, 10)
					conn.Read(buf)
					fmt.Println("read 10 byte:", string(buf), "from", conn.RemoteAddr())
				}
			} else {
				fmt.Println("waiting", interval, "nanosecond")
			}

			time.Sleep(time.Duration(interval) * time.Nanosecond)
		}
	}
}
