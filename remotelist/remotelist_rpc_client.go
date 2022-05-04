package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	remotelist "ppgti/remotelist/pkg"
	"strconv"
	"strings"
)

func readOption() int {
	var reader = bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	op, err := strconv.Atoi(strings.TrimSpace(input))

	for err != nil {
		fmt.Print("Digite uma opção válida! ")
		input, _ = reader.ReadString('\n')
		op, err = strconv.Atoi(strings.TrimSpace(input))
	}

	return op
}

func menuLoop(client *rpc.Client) {
	var op int = readOption()

	for op != 0 {
		switch op {
		case 1:
			callAppend(client)
		case 2:
			callRemove(client)
		case 3:
			callSize(client)
		case 4:
			callGet(client)
		default:
			fmt.Print("Digite uma opção válida! ")
		}
		op = readOption()
	}
	client.Close()
}

func callAppend(client *rpc.Client) {
	// Append
	var reply bool
	param := new(remotelist.RemoteRequest)
	param.Id = 0
	param.Value = 10
	err := client.Call("RemoteList.Append", param, &reply)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento adicionado:", reply)
	}
}

func callRemove(client *rpc.Client) {
	// Remove
	var reply bool
	err := client.Call("RemoteList.Remove", 0, &reply)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento removido:", reply)
	}
}

func callGet(client *rpc.Client) {
	// Get
	param := new(remotelist.RemoteRequest)
	param.Id = 0
	param.Value = 0
	var reply int
	err := client.Call("RemoteList.Get", param, &reply)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento: ", reply)
	}
}

func callSize(client *rpc.Client) {
	// Size
	var reply int
	err := client.Call("RemoteList.Size", 0, &reply)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Tamanho:", reply)
	}
}

func main() {
	client, err := rpc.Dial("tcp", ":5002")
	if err != nil {
		fmt.Print("dialing:", err)
	} else {
		menuLoop(client)
	}
}
