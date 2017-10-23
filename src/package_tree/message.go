package main 

import (
	"regexp"
	"fmt"
	"strings"
)

type command int

type Command interface {
	Command() command
}

const (
	INDEX command = iota
	QUERY
	REMOVE
	ERROR
)

func(c command) Command() command {
	return c
}

type Message struct {
	cmd Command
	pkg string
	dep []string
}

func NewMessage(message string) Message {
	r, err := regexp.Compile(`^(INDEX|REMOVE|QUERY)\|([\w\-+]+)\|([\w,\-+]*)`)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	s := r.FindStringSubmatch(message)
	deps := make([]string, 0)
	if len(s) == 0 {
		return Message{ERROR,"", deps}
	}
	if len(s) == 4 {
		deps = strings.Split(s[3],",")
	}

	switch s[1] {
	case "INDEX":
		return Message{INDEX, s[2], deps}
	case "REMOVE":
		return Message{REMOVE, s[2], deps}
	case "QUERY":
		return Message{QUERY, s[2], deps}
	default:
		return Message{ERROR, "", deps}
	}
}
