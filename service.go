package main

import (
	"example.com/grouprpc/group"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type GroupService struct {
	group group.Group
}

func setSuccess(err error, b *bool) error {
	*b = err == nil
	return err
}

func (g *GroupService) AddStudent(student_id int, success *bool) error {
	return setSuccess(g.group.AddStudent(student_id), success)
}

func main() {
	g := new(GroupService)
	if err := rpc.Register(g); err != nil {
		log.Fatalln(err)
	}
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server Started")
	http.Serve(l, nil)

}
