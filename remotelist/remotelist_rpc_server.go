package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/rpc"
	remotelist "ppgti/remotelist/pkg"
)

func loadData() *remotelist.RemoteList {
	content, err := ioutil.ReadFile(remotelist.Filename)
	if err != nil {
		return remotelist.NewRemoteList()
	}
	rlist := remotelist.NewRemoteList()
	err = json.Unmarshal(content, &rlist)
	if err != nil {
		return remotelist.NewRemoteList()
	}
	return rlist
}

func main() {
	list := loadData()
	rpcs := rpc.NewServer()
	rpcs.Register(list)
	l, e := net.Listen("tcp", ":5002")
	if e != nil {
		fmt.Println("listen error:", e)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err == nil {
			go rpcs.ServeConn(conn)
		} else {
			break
		}
	}
}
