package main

import "fmt"

//StorePlayer1Name for later uses
var StorePlayer1Name string

//StorePlayer2Name for later uses
var StorePlayer2Name string

func main() {
	fmt.Println("Zendesk Take Home Assignment - TTT")
	getPlayersName()
	var gameEnded bool
	gameBoard := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	currentTurn := 1
	fmt.Println(len(gameBoard))
	fmt.Println(gameEnded)
	fmt.Println(currentTurn)

	getGameBoard(gameBoard)

}

// asking for player 1's  name
func getPlayersName() string {
	fmt.Println("Input player one's name :")
	var Player1Name string
	fmt.Scanln(&Player1Name)
	if len(Player1Name) != 0 {
		fmt.Println("Welcome", Player1Name, " : Your marker is X.")
		getPlayer2Name(Player1Name)

	} else {
		fmt.Println("Please enter a valid input :")
		getPlayersName()
	}
	StorePlayer1Name = Player1Name
	return Player1Name
}

//asking for Player 2's name
func getPlayer2Name(name string) string {
	fmt.Println("Input player two's name :")
	var player2Name string
	fmt.Scanln(&player2Name)
	if len(player2Name) != 0 && player2Name != name {
		fmt.Println("Welcome", player2Name, " : Your marker is O.")

	} else {
		fmt.Println("Please enter a valid input :")
		getPlayer2Name(name)
	}
	StorePlayer2Name = player2Name
	return player2Name
}

//Displaying the Gameboard on the terminal
func getGameBoard(board [9]int) {
	for i, v := range board {
		if v == 0 {
			fmt.Printf("%d", i)
		} else if v == 1 {
			fmt.Printf("X")
		} else if v == 10 {
			fmt.Printf("O")
		}
		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" | ")
		}
	}
}
