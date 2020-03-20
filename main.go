package main

import (
	random "math/rand"
	"time"
)

func init() {
	random.Seed(time.Now().UnixNano())
}

func main() {

}
