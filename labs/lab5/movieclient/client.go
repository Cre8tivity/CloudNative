// Package main imlements a client for movieinfo service
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Cre8tivity/CloudNative/labs/lab5/movieapi"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultTitle = "Pulp fiction"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := movieapi.NewMovieInfoClient(conn)

	// Contact the server and print out its response.
	title := defaultTitle
	if len(os.Args) > 1 {
		title = os.Args[1]
	}
	// Timeout if server doesn't respond
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: title})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for %s: %d %s %v", title, r.GetYear(), r.GetDirector(), r.GetCast())

	// Start test code:
	title2 := "Transformers"
	year2 := "2007"
	director2 := "Michael Bay"
	cast2 := []string{"Shia LaBeouf", "Megan Fox", "Josh Duhamel", "Tyrese Gibson"}

	r2, err := c.SetMovieInfo(ctx, &movieapi.MovieRequest{Title: title2, Year: year2, Director: director2, Cast: cast2})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Set Movie Info for %s: %d %s %v\n", title2, r2.GetYear(), r2.GetDirector(), r2.GetCast())

	r3, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: title2})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for %s: %d %s %v", title2, r3.GetYear(), r3.GetDirector(), r3.GetCast())
}
