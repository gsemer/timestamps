FROM golang:1.18.0-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /timestamps ./cmd/main.go

EXPOSE 8000

CMD [ "/timestamps" ]