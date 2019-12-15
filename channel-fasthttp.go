package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/valyala/fasthttp"
)

type CommandType int

const (
	GetCommand = iota
	SetCommand
	IncCommand
)

type Command struct {
	ty        CommandType
	name      string
	val       int
	replyChan chan int
}

func startCounterManager(initvals map[string]int) chan<- Command {
	counters := make(map[string]int)
	for k, v := range initvals {
		counters[k] = v
	}

	cmds := make(chan Command)

	go func() {
		for cmd := range cmds {
			switch cmd.ty {
			case GetCommand:
				if val, ok := counters[cmd.name]; ok {
					cmd.replyChan <- val
				} else {
					cmd.replyChan <- -1
				}
			case SetCommand:
				counters[cmd.name] = cmd.val
				cmd.replyChan <- cmd.val
			case IncCommand:
				if _, ok := counters[cmd.name]; ok {
					counters[cmd.name]++
					cmd.replyChan <- counters[cmd.name]
				} else {
					cmd.replyChan <- -1
				}
			default:
				log.Fatal("unknown command type", cmd.ty)
			}
		}
	}()
	return cmds
}

type Server struct {
	cmds chan<- Command
}

func (h *Server) inc(ctx *fasthttp.RequestCtx) {
	log.Printf("inc %v", ctx.Request)
	args := ctx.Request.URI().QueryArgs().GetBool("name")
	var name string
	if args {
		name = "i"
	} else {
		name = "i"
	}
	replyChan := make(chan int)
	h.cmds <- Command{ty: IncCommand, name: name, replyChan: replyChan}

	reply := <-replyChan
	if reply >= 0 {
		fmt.Fprintf(ctx, "ok\n")
	} else {
		fmt.Fprintf(ctx, "%s not found\n", name)
	}
}

func main() {
	server := Server{startCounterManager(map[string]int{"i": 0, "j": 0})}

	portnum := 8000
	if len(os.Args) > 1 {
		portnum, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Going to listen on port %d\n", portnum)
	log.Fatal(fasthttp.ListenAndServe("localhost:"+strconv.Itoa(portnum), server.inc))
}
