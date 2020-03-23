package main

import (
	"fmt"
	random "math/rand"
	"time"

	"github.com/koungkub/thw/rand"

	"github.com/koungkub/thw/hand"
	"github.com/koungkub/thw/io"
)

func init() {
	random.Seed(time.Now().UnixNano())
}

func main() {

	fmt.Println(io.Welcome())
	predictorState := true

	for {

		fmt.Println(io.Turn(predictorState))

		player := io.Scan("Player: ")
		err := hand.Validate(predictorState, player)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		ai := rand.RandomHands(!predictorState)
		fmt.Printf("AI: %v\n", ai)

		err = hand.Compare(player, ai)
		if err != nil {
			fmt.Println("No winner !!")
			predictorState = !predictorState
			continue
		}

		confirm := io.Scan("Do you want to play again [Y]es or [N]o ?: ")
		if confirm == "Y" || confirm == "YES" {
			predictorState = !predictorState
			continue
		}

		fmt.Println("OK, bye !!")
		break
	}
}
