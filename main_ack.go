package main

import (
  "fmt"
  // Denne relative stien må ligge i GOPATH
  // Og GO111MODULE må settes til "auto"
  "github.com/uia-worker/turingmaskin/ack"
)

func main() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Printf("ackermann (%d, %d) is: %d\n", i, j, ack.Ack(i, j))
		}
	}
}
