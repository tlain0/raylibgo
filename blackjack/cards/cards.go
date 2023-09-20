package cards

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Loads card suit textures before game is started
func LoadTextures() []rl.Texture2D {
	textures := []rl.Texture2D{rl.LoadTexture("./resources/spade.png"), rl.LoadTexture("./resources/diamond.png"), rl.LoadTexture("./resources/club.png"), rl.LoadTexture("./resources/heart.png")}
	for i := 0; i < len(textures); i++ {
		textures[i].Height = 30
		textures[i].Width = 30
	}
	return textures
}

type Card struct {
	Value            int
	Suit             string
	Icon             rl.Texture2D
	SpecialCardValue string
}
type Deck struct {
	Cards []Card
}

// Generates cards from 1 - K for each of the 4 suits
func (deck *Deck) GenerateDeck(textures []rl.Texture2D) {
	suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}

	// String for displaying the card value so the picture cards work
	n := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	// Repeat four times (for each suit)
	for i := 0; i < len(suits); i++ {

		// 14 times for each card in suit ( 1 -> K )
		for f := 1; f < 14; f++ {

			// J, Q, and K are valued at 10
			// If f > 10 then set Card.Value = 10
			if f > 10 {
				deck.Cards = append(deck.Cards, Card{Value: 10, Suit: suits[i], Icon: textures[i], SpecialCardValue: n[f-1]})
			} else {
				deck.Cards = append(deck.Cards, Card{Value: f, Suit: suits[i], Icon: textures[i], SpecialCardValue: n[f-1]})
			}
		}
	}
}

// Generates randint and picks (removes) card from deck
func (deck *Deck) DrawCard() Card {
	rand.Seed(time.Now().UnixNano())
	randint := rand.Intn(len(deck.Cards))

	pick := deck.Cards[randint]
	deck.Cards = append(deck.Cards[:randint], deck.Cards[randint+1:]...)

	return pick
}

// Paints card on screen
func (card *Card) PaintCard(posX, posY int32) {
	suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}

	var color rl.Color = rl.Red
	if card.Suit == suits[0] || card.Suit == suits[2] {
		color = rl.Black
	}

	rl.DrawRectangle(posX, posY, 100, 150, rl.White)
	rl.DrawText(card.SpecialCardValue, posX+7, posY+5, 30, color)
	rl.DrawTexture(card.Icon, posX+2, posY+40, rl.White)

}
