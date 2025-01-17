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
	var (
		// parse cards
		c1, c2, dc = ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
		// establish hand
		hand = c1 + c2
		// determine outcome modifiers
		split, sevens_high, dealer_match = (c1 == c2), (dc >= 7), (c1 == dc || c2 == dc)
	)
	switch {
	case (hand > 21 && split) || (hand >= 12 && !sevens_high && split):
		return "P" // Split (P)
	case (hand == 21 && dealer_match) || (hand == 20) || (hand < 21 && hand >= 17) || (hand < 21 && hand >= 12 && !sevens_high && !split):
		return "S" // Stand (S)
	case (hand == 21 && !dealer_match):
		return "W" // Automatically win (W)
	case (hand >= 12 && sevens_high) || (hand <= 11):
		return "H" // Hit (H)
	default:
		return "unknown"
	}
}
