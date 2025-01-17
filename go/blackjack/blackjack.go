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
func FirstTurn(card1, card2, dealerCard string) (outcome string) {
	/*
		- Stand (S)
		- Hit (H)
		- Split (P)
		- Automatically win (W)
	*/var (
		p1          = ParseCard(card1)
		p2          = ParseCard(card2)
		p3          = ParseCard(dealerCard)
		hand        = p1 + p2
		split       = p1 == p2
		sevens_high = p3 >= 7
	)
	switch {
	case hand > 21 && split:
		outcome = "P" // split the two aces
		_ = p3
	case hand == 21 && !split && (p1 == p3 || p2 == p3):
		outcome = "S"
		_ = p3
	case hand == 21 && !split && (p1 > p3 || p2 > p3):
		outcome = "W"
		_ = p3
	case hand == 20 && split: // actually you "should" split to get a bigger pot... but what ever.
		outcome = "S" // stand on two "10"s
		_ = p3
	case hand == 20 && !split:
		outcome = "S" // stand on 20
		_ = p3
	case hand == 19 && !split:
		outcome = "S" // stand on 19
		_ = p3
	case hand == 18 && !split:
		outcome = "S" // stand on 18
		_ = p3
	case hand == 17 && !split:
		outcome = "S" // stand on 17
		_ = p3
	case hand == 16 && !split && !sevens_high:
		outcome = "S" // stand on 16 when dealer has less than 7
		// this seems wrong, because dealer could have an ace up his sleeve... 11 + 6 = 17...
		_ = p3
	case hand == 16 && !split && sevens_high:
		outcome = "H" // hit on 16 when dealer has a 7 (could be 17)
		_ = p3
	case hand == 15 && !split && !sevens_high:
		outcome = "S" // stand on 16 when dealer has less than 7
		// this seems wrong, because dealer could have an ace up his sleeve... 11 + 6 = 17...
		_ = p3
	case hand == 15 && !split && sevens_high: // this is almost always going to end in a bust
		outcome = "H" // hit on 16 when dealer has a 7 (could be 17)
		_ = p3
	case hand == 15 && !split:
		outcome = "H" // hit
		_ = p3
	case hand == 15 && split:
		outcome = "H" // hit
		_ = p3
	case hand == 14 && !split && !sevens_high:
		outcome = "S" // stand
		_ = p3
	case hand == 14 && !split && sevens_high:
		outcome = "H" // hit
		_ = p3
	case hand == 13 && !split && !sevens_high:
		outcome = "S" // stand
		_ = p3
	case hand == 13 && !split && sevens_high:
		outcome = "H" // hit
		_ = p3
	case hand == 12 && !split && !sevens_high:
		outcome = "S" // stand
		_ = p3
	case hand == 12 && !split && sevens_high:
		outcome = "H" // hit
		_ = p3
	case hand == 11 && !split && !sevens_high:
		outcome = "S" // stand
		_ = p3
	case hand == 11 && !split && sevens_high:
		outcome = "H" // hit
		_ = p3
	case hand <= 10:
		outcome = "H" // hit
		_ = p3
	default:
		outcome = "unknown"
	}
	return
}
