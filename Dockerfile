FROM golang:1.16.3 AS build

ENV GO111MODULE=on

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /usr/bin/liproduce ./cmd/liproduce

RUN go build -o /usr/bin/migrations ./cmd/migrations

RUN chmod +x ./build/commands.sh

EXPOSE 8080

ENV ENVIRONMENT=production

ENTRYPOINT [ "/usr/src/app/build/commands.sh" ]
