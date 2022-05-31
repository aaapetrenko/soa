package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/hw5_grpc/mafia/mafia"
	"log"
	"os"
	"time"
)

type User struct {
	name   string
	status string
}

func main() {
	log.Println(":) Client running ...")
	var u User
	log.Println(":) Enter your user name -->")
	fmt.Fscan(os.Stdin, &u.name)
	log.Printf(":) Hello, %s", u.name)
	conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := mafia.NewReverseClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for {
		log.Println(":) Enter the command -->")
		var command = ""
		fmt.Fscan(os.Stdin, &command)
		if command == "exit" {
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			request := &mafia.Request{Url2: u.name}
			client.Exit(ctx, request)
			break
		}
		if command == "do" {
			log.Println(":) Enter 2 url -->")
			url1 := ""
			url2 := ""
			fmt.Fscan(os.Stdin, &url1)
			fmt.Fscan(os.Stdin, &url2)
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			request := &mafia.Request{Url1: url1, Url2: url2}
			response, err := client.Do(ctx, request)
			if err != nil {
				log.Fatalln(err)
			}
			log.Printf(response.Pong)
		} else if command == "get" {
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			request := &mafia.Request{Url1: "url1"}
			response, err := client.Answer(ctx, request)
			if err != nil {
				log.Fatalln(err)
			}
			log.Printf(response.Url1)
		}
	}

}
