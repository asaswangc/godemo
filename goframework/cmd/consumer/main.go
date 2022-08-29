package main

import (
	"goframework/src/goframework/rabbitmq"
)

func main() {
	rabbitmq.Consumer("127.0.0.1", 5672, "rabbit", "rabbit")
}
