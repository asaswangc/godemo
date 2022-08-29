package main

import (
	"goframework/src/goframework/rabbitmq"
)

func main() {
	rabbitmq.Producer("127.0.0.1", 5672, "rabbit", "rabbit")
}
