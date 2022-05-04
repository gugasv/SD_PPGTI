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

func readOption(msg string) int {
	fmt.Print(msg)
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
	var op int = readOption("Escolha a opção: (1-append; 2-remove; 3-size; 4-get)")

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
		op = readOption("Escolha a opção: (1-append; 2-remove; 3-size; 4-get)")
	}
	client.Close()
}

func callAppend(client *rpc.Client) {
	// Append
	var reply bool
	param := new(remotelist.RemoteRequest)
	param.Id = readOption("Identificador da lista: ")
	param.Value = readOption("Valor: ")
	err := client.Call("RemoteList.Append", param, &reply)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Elemento adicionado:", reply)
	}
}

func callRemove(client *rpc.Client) {
	// Remove
	var reply bool
	id := readOption("Identificador da lista: ")
	err := client.Call("RemoteList.Remove", id, &reply)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Elemento removido:", reply)
	}
}

func callGet(client *rpc.Client) {
	// Get
	param := new(remotelist.RemoteRequest)
	param.Id = readOption("Identificador da lista: ")
	param.Value = readOption("Posição: ")
	var reply int
	err := client.Call("RemoteList.Get", param, &reply)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Elemento: ", reply)
	}
}

func callSize(client *rpc.Client) {
	// Size
	var reply int
	id := readOption("Identificador da lista: ")
	err := client.Call("RemoteList.Size", id, &reply)
	if err != nil {
		fmt.Println("Error:", err)
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
