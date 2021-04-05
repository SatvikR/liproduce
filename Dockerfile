FROM golang:1.16.3 AS build

ENV GO111MODULE=on

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /usr/bin/liproduce ./cmd/liproduce

RUN go build -o /usr/bin/migrations ./cmd/migrations

FROM scratch AS bin

COPY --from=build /usr/bin/liproduce /usr/bin/
COPY --from=build /usr/src/app/build/commands.sh /scripts/commands.sh

EXPOSE 8080

ENTRYPOINT [ "/scripts/commands.sh" ]
