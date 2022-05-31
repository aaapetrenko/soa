package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/hw5_grpc/mafia/mafia"
	"log"
	"math/rand"
	"os"
	"time"
)

type User struct {
	name   string
	status string
}

func main() {
	var gamer_name = []string{"", "", "", ""}
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
	request := &mafia.Request{Message: u.name}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.Name(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}
	var answ = response.GetMessage()
	if "There are no places" == answ {
		log.Printf("%s\n", answ)
		return
	} else {
		u.status = answ
		log.Println(":) Success, you are", answ)
	}

	for {
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		request = &mafia.Request{Message: u.name}
		response, err := client.Ping(ctx, request)
		if err != nil {
			log.Fatalln(err)
		}
		var tmp = response.GetMessage()
		if tmp == "OK" {
			gamer_name[0] = response.Name1
			gamer_name[1] = response.Name2
			gamer_name[2] = response.Name3
			gamer_name[3] = response.Name4
			log.Printf(":) List of players: %s %s %s %s\n", response.Name1, response.Name2, response.Name3, response.Name4)
			break
		}
	}
	log.Println(":) woohoo, let's go")
	flagDayNight := 0 // день
	flagVote := 0
	for {
		if flagDayNight == 0 {
			log.Println(":) Now is the day")
		} else {
			log.Println(":) Now is the night")
		}
		log.Println(":) Enter the command -->")
		var command = ""
		fmt.Fscan(os.Stdin, &command)
		if command == "exit" {
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			request = &mafia.Request{Message: u.name}
			client.Exit(ctx, request)
			break
		}
		if flagDayNight == 0 {
			if command == "vote" && flagVote == 0 {
				log.Printf("%v", gamer_name)
				randGamer := ""
				for {
					var randInd = rand.Intn(len(gamer_name))
					randGamer = gamer_name[randInd]
					if randGamer != u.name && randGamer != "" {
						break
					}
				}
				ctx, cancel = context.WithTimeout(context.Background(), time.Second)
				request = &mafia.Request{Message: randGamer}
				response, err = client.Vote(ctx, request)
				if response.Message == "mafia" || response.Message == "civilian" {
					log.Printf("This is the end. Status: %s", response.Message)
					break
				}
				if err != nil {
					log.Fatalln(err)
				}
				flagVote = 1
				log.Printf("you voted for %s", randGamer)
			} else if command == "change" {
				flagDayNight = 1
				flagVote = 0
				spirit := ""
				for {
					ctx, cancel = context.WithTimeout(context.Background(), time.Second)
					request = &mafia.Request{Message: u.name}
					response, err := client.PingResultDay(ctx, request)
					if response.Status == "END" {
						log.Printf("This is the end. Status: %s", response.Spirit)
						break
					}
					if err != nil {
						log.Fatalln(err)
					}
					var tmp = response.Status
					if tmp == "OK" {
						spirit = response.Spirit
						break
					}
				}
				if spirit == u.name {
					log.Printf("I died :(( %s", spirit)
					break
				}
				for i, v := range gamer_name {
					if v == spirit {
						gamer_name[i] = ""
					}
				}
				flagDayNight = 1
				log.Printf("he died :(( %s %v", spirit, gamer_name)
			} else if command == "members" {
				ctx, cancel = context.WithTimeout(context.Background(), time.Second)
				request = &mafia.Request{Message: u.name}
				response, err := client.PingMembers(ctx, request)
				if response.Message == "END" {
					log.Printf("This is the end. Status: %s", response.Name1)
					break
				}
				if err != nil {
					log.Fatalln(err)
				}
				var tmp = response.GetMessage()
				if tmp == "OK" {
					gamer_name[0] = response.Name1
					gamer_name[1] = response.Name2
					gamer_name[2] = response.Name3
					gamer_name[3] = response.Name4
					log.Printf(":) List of players: %s %s %s %s\n", response.Name1, response.Name2, response.Name3, response.Name4)
				}
			}
		} else {
			if command == "change" {
				flagDayNight = 0
				ctx, cancel = context.WithTimeout(context.Background(), time.Second)
				request = &mafia.Request{Message: u.name}
				response, err := client.ChangeTime(ctx, request)
				if response.Message == "YES" {
					log.Printf("I died :((((")
					break
				}
				if err != nil {
					log.Fatalln(err)
				}
			}
			if u.status == "civilian" || u.status == "doctor" {
				continue
			}
			if command == "vote" {
				randGamer := ""
				for {
					var randInd = rand.Intn(len(gamer_name))
					randGamer = gamer_name[randInd]
					if randGamer != u.name && randGamer != "" {
						break
					}
				}
				ctx, cancel = context.WithTimeout(context.Background(), time.Second)
				request = &mafia.Request{Message: randGamer}
				response, err = client.NightVote(ctx, request)
				if response.Message == "mafia" || response.Message == "civilian" {
					log.Printf("This is the end. Status: %s", response.Message)
					break
				}
			} else if command == "members" {
				ctx, cancel = context.WithTimeout(context.Background(), time.Second)
				request = &mafia.Request{Message: u.name}
				response, err := client.PingMembers(ctx, request)
				if response.Message == "END" {
					log.Printf("This is the end. Status: %s", response.Name1)
					break
				}
				if err != nil {
					log.Fatalln(err)
				}
				var tmp = response.GetMessage()
				if tmp == "OK" {
					gamer_name[0] = response.Name1
					gamer_name[1] = response.Name2
					gamer_name[2] = response.Name3
					gamer_name[3] = response.Name4
					log.Printf(":) List of players: %s %s %s %s\n", response.Name1, response.Name2, response.Name3, response.Name4)
				}
			}
		}
	}
}
