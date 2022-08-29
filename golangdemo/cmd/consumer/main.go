package main

import (
	"golangdemo/src/framework/rabbitmq"
)

func main() {
	rabbitmq.Consumer("127.0.0.1", 5672, "rabbit", "rabbit")
}
