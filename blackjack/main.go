package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	b "raylibcom.com/main/blackjack/buttons"
	c "raylibcom.com/main/blackjack/cards"
)

func main() {
	var screenWidth, screenHeight int32 = 800, 450
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	// Load textures, create deck and initialize arrays of cards
	textures := c.LoadTextures()

	var deck c.Deck
	deck.GenerateDeck(textures)

	var playerCards []c.Card
	var dealerCards []c.Card

	// Give 2 cards to the player and 1 to the dealer
	playerCards = append(playerCards, deck.DrawCard(), deck.DrawCard())
	dealerCards = append(dealerCards, deck.DrawCard())

	var stay = false
	// Runs every frame
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		var playerCardsValue, dealerCardsValue = updateCardCount(playerCards, dealerCards)
		paintLoop(playerCards, dealerCards, playerCardsValue, dealerCardsValue)

		// Check for button clicks
		if b.HitButton.CheckClick("Hit") {
			playerCards = append(playerCards, deck.DrawCard())
		}

		if b.StayButton.CheckClick("Stay") {
			stay = true
			for dealerCardsValue <= 17 {
				dealerCards = append(dealerCards, deck.DrawCard())
				playerCardsValue, dealerCardsValue = updateCardCount(playerCards, dealerCards)
			}
		}

		// Restarts Game

		if gameOver(playerCardsValue, dealerCardsValue, stay) {
			b.RestartButton.DrawButton("Restart")
			if b.RestartButton.CheckClick("Restart") {
				stay = false
				playerCards = nil
				dealerCards = nil
				playerCards = append(playerCards, deck.DrawCard(), deck.DrawCard())
				dealerCards = append(dealerCards, deck.DrawCard())
			}
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()

}

// Runs graphics loop
func paintLoop(playerCards, dealerCards []c.Card, playerCardsValue, dealerCardsValue int) {

	// Hit and stay buttons
	b.HitButton.DrawButton("Hit")
	b.StayButton.DrawButton("Stay")

	// X- position for painting cards
	var posX int32 = 50

	// Paint cards for each card in player's hand
	for i := 0; i < len(playerCards); i++ {
		playerCards[i].PaintCard(posX, 250)
		posX += 110
	}
	posX = 50

	// Paint cards for each card in dealer's hand
	for i := 0; i < len(dealerCards); i++ {
		dealerCards[i].PaintCard(posX, 50)
		posX += 110
	}

	// Paint Score
	rl.DrawText(fmt.Sprint(playerCardsValue), 600, 300, 30, rl.Black)
	rl.DrawText(fmt.Sprint(dealerCardsValue), 600, 150, 30, rl.Black)
}

// Check win conditions
func gameOver(playerCardsValue, dealerCardsValue int, stay bool) bool {
	if stay && playerCardsValue > dealerCardsValue && playerCardsValue <= 21 {
		b.GameOverRec.DrawButton("You've won")
		return true
	} else if stay && playerCardsValue < dealerCardsValue && dealerCardsValue <= 21 {
		b.GameOverRec.DrawButton("You've lost")
		return true
	}
	if playerCardsValue >= 22 {
		b.GameOverRec.DrawButton("You've Lost")
		return true
	} else if dealerCardsValue >= 17 && playerCardsValue > dealerCardsValue {
		b.GameOverRec.DrawButton("You've Won!")
		return true
	} else if dealerCardsValue >= 22 {
		b.GameOverRec.DrawButton("You've Won!")
		return true
	}
	if stay && playerCardsValue == dealerCardsValue && playerCardsValue <= 21 && dealerCardsValue <= 21 {
		b.GameOverRec.DrawButton("You've tied")
	}
	return false
}

// Returns sum of card values
func updateCardCount(playerCards, dealerCards []c.Card) (int, int) {
	dealer := 0
	player := 0

	for i := 0; i < len(playerCards); i++ {
		player += playerCards[i].Value
	}
	for i := 0; i < len(dealerCards); i++ {
		dealer += dealerCards[i].Value
	}

	return player, dealer
}
