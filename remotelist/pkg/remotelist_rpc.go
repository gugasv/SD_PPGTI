package remotelist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
)

var Filename string = "listfile.json"

type RemoteRequest struct {
	Id    int
	Value int
}

type RemList struct {
	mu    sync.Mutex
	Items []int
}

type RemoteList struct {
	Lists map[int]*RemList
}

func (l *RemList) Append(req *RemoteRequest) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Items = append(l.Items, req.Value)
	fmt.Println(l.Items)
	return nil
}

func (l *RemList) Remove() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.Items) > 0 {
		l.Items = l.Items[:len(l.Items)-1]
		fmt.Println(l.Items)
	} else {
		fmt.Println("Lista vazia")
		return errors.New("Lista vazia")
	}
	return nil
}

func (l *RemoteList) Append(req *RemoteRequest, reply *bool) error {
	if _, contains := l.Lists[req.Id]; !contains {
		l.Lists[req.Id] = new(RemList)
	}
	go l.Lists[req.Id].Append(req)
	go l.Save()

	*reply = true
	return nil
}

func (l *RemoteList) Remove(id int, reply *bool) error {
	if _, contains := l.Lists[id]; contains {
		go l.Lists[id].Remove()
		go l.Save()
	} else {
		return errors.New("Lista não encontrada")
	}
	*reply = true
	return nil
}

func (l *RemoteList) Get(req *RemoteRequest, reply *int) error {
	if ll, contains := l.Lists[req.Id]; contains {
		if req.Value < len(ll.Items) {
			*reply = ll.Items[req.Value]
		} else {
			*reply = -1
		}
	} else {
		return errors.New("Lista não encontrada")
	}
	return nil
}

func (l *RemoteList) Size(id int, reply *int) error {
	if ll, contains := l.Lists[id]; contains {
		*reply = len(ll.Items)
	} else {
		return errors.New("Lista não encontrada")
	}
	return nil
}

func (l *RemoteList) Save() error {
	jsonByte, _ := json.Marshal(l)
	err := ioutil.WriteFile(Filename, jsonByte, 0644)
	if err != nil {
		return err
	}
	return nil
}

func NewRemoteList() *RemoteList {
	l := new(RemoteList)
	l.Lists = make(map[int]*RemList)
	return l
}
