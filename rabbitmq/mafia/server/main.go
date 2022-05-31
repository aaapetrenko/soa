package main

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/hw5_grpc/mafia/mafia"
	"log"
	"net"
	"net/http"
	"strconv"
)

type server struct {
	mafia.UnimplementedReverseServer
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
	return
}

var global_result = ""

func (s *server) Do(ctx context.Context, request *mafia.Request) (*mafia.SResponse, error) {
	gl_url := "https://en.wikipedia.org/wiki/Main_Page"
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/%2f")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	body := "\n WAY --> "
	count := strconv.Itoa(0)
	first_url := request.Url1
	second_url := request.Url2
	if first_url != second_url {
		count = strconv.Itoa(2)
		req, _ := http.Get(first_url)
		defer req.Body.Close()
		//b, _ := ioutil.ReadAll(req.Body)
		//fmt.Printf("%s\n", b)
		body += first_url + "\n"
		body += gl_url + "\n"
		req, _ = http.Get(second_url)
		defer req.Body.Close()
		//b, _ = ioutil.ReadAll(req.Body)
		//fmt.Printf("%s\n", b)
		body += second_url + "\n"
		body += "ANSWER --> "
		body += count
	} else {
		req, _ := http.Get(first_url)
		defer req.Body.Close()
		body += first_url + "\n"
		body += "ANSWER --> "
		body += count
	}
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body)})
	failOnError(err, "Failed to publish a message")
	global_result = body
	return &mafia.SResponse{Pong: "OK"}, nil
}

func (s *server) Answer(ctx context.Context, request *mafia.Request) (*mafia.Response, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/%2f")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")
	go func() {
		for d := range msgs {
			global_result = string(d.Body)
			//log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	return &mafia.Response{Url1: global_result}, nil
}

func (s *server) Exit(ctx context.Context, request *mafia.Request) (*mafia.SResponse, error) {
	return &mafia.SResponse{Pong: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	mafia.RegisterReverseServer(srv, &server{})

	log.Fatalln(srv.Serve(lis))

}
