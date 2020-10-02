package main

import "fmt"

//StorePlayer1Name for later usage
var StorePlayer1Name string

//StorePlayer2Name for later usage
var StorePlayer2Name string

//StorePlayerMove for later usage
var StorePlayerMove int

func main() {
	fmt.Println("Zendesk Take Home Assignment - TTT")
	getPlayersName()
	var gameEnded bool
	gameBoard := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	currentTurn := 1
	// fmt.Println(len(gameBoard))
	// fmt.Println(gameEnded)
	// fmt.Println(currentTurn)
	fmt.Println("Key 99 to quit game!")
	for gameEnded != true {
		getGameBoard(gameBoard)
		player := 0
		if currentTurn%2 == 1 {
			fmt.Printf("Please Make Your Move:%v\n", StorePlayer1Name)
			player = 1
		} else {
			fmt.Printf("Please Make Your Move:%v\n", StorePlayer2Name)
			player = 2
		}
		// fmt.Println(player)
		fmt.Println("Please enter a value from 0-8")

		move := getPlayerMove()

		if StorePlayerMove == 99 {
			fmt.Println("Thank you, See you again!")
			return
		}

		gameBoard = applyPlayerMove(move, player, gameBoard)
		results := checkWinCondition(gameBoard)
		if results > 0 {
			fmt.Printf("Player %d win\n\n", results)
		} else if currentTurn == 9 {
			fmt.Println("Draw Game!")
			gameEnded = true
		} else {
			currentTurn++
		}
	}
}

//Asking for Player 1's  name
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

//Asking for Player 2's name
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

//Prompt the player to key a move
func getPlayerMove() int {
	var PlayerMove int
	fmt.Scan(&PlayerMove)
	StorePlayerMove = PlayerMove
	return StorePlayerMove
}

func applyPlayerMove(StorePlayerMove int, player int, board [9]int) [9]int {

	// Check if the current spot is  occupied
	if board[StorePlayerMove] != 0 {
		fmt.Println("Space no available")
		StorePlayerMove = getPlayerMove()
		board = applyPlayerMove(StorePlayerMove, player, board)
	} else {
		if player == 1 {
			board[StorePlayerMove] = 1
		} else if player == 2 {
			board[StorePlayerMove] = 10
		}
	}

	for StorePlayerMove > 9 && StorePlayerMove < 0 {
		fmt.Println("Please enter a number from 0 - 9")
		StorePlayerMove = getPlayerMove()
	}

	if player == 1 {
		board[StorePlayerMove] = 1
	} else if player == 2 {
		board[StorePlayerMove] = 10
	}
	return board
}

//hard-coded win condition :p
func checkWinCondition(b [9]int) int {

	sums := [8]int{0, 0, 0, 0, 0, 0, 0, 0}

	sums[0] = b[2] + b[4] + b[6]
	sums[1] = b[0] + b[3] + b[6]
	sums[2] = b[1] + b[4] + b[7]
	sums[3] = b[2] + b[5] + b[8]
	sums[4] = b[0] + b[4] + b[8]
	sums[5] = b[6] + b[7] + b[8]
	sums[6] = b[3] + b[4] + b[5]
	sums[7] = b[0] + b[1] + b[2]
	for _, v := range sums {
		if v == 3 {
			return 1
		} else if v == 30 {
			return 2
		}
	}
	return 0
}
