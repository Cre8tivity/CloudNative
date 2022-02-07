// Demonstration of channels with a chat application
// Copyright Â© 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Chat is a server that lets clients chat with each other.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// changed
type client struct {
	channel chan<- string // an outgoing message channel
	name    string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.channel <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			// changed
			var online []string

			for cList := range clients {
				online = append(online, cList.name)
			}
			cli.channel <- fmt.Sprintf("%d clients connected: < %s >", len(clients), strings.Join(online, ", "))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.channel)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	var who string
	// changed
	var count bool = false
	ch <- "say your name"
	input := bufio.NewScanner(conn)
	for input.Scan() {
		if !count {

			who = input.Text()
			if who == "" {
				ch <- "You are " + conn.RemoteAddr().String()
				who = conn.RemoteAddr().String()
			} else {
				ch <- "You are " + who
			}

			messages <- who + " has arrived"

			entering <- client{ch, who}

			// testing: fmt.Println("here")
		} else {
			messages <- "[ " + who + "] " + ": " + input.Text()
		}
		count = true
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client{ch, who}
	messages <- who + " has left" // changed
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
