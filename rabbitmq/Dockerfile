FROM golang:latest

WORKDIR /

#COPY go.mod go.sum ./
#COPY mafia/server/main.go ./
COPY . .

RUN go mod download && go mod verify

ENTRYPOINT go run mafia/server/main.go
