package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) (value int) {
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "king", "queen", "jack", "ten":
		value = 10
	}
	return
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	panic("Please implement the FirstTurn function")
}
