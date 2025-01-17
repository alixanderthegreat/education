Now, the first approach that I took for this exercise looked like this:
```go
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
	return value
}
```
But even as I was looking at it, I thought... there must be a faster way of doing this. And so I invited the bot to give me a more streamlined appraoch. This is what it came up with on first try:
```go
func ParseCard(card string) (value int) {
	cardValues := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"king":  10,
		"queen": 10,
		"jack":  10,
		"ten":   10,
	}

	// Return the value directly from the map
	return cardValues[card]
}
```
I looked over the solution for a moment and thought, it would probably be better to just to `return map[string]int` directly; to which, when prompted, the bot said the same thing. I invited the bot to come up with some other methods for determine the `int` and it has some pretty wild ideas: `Pattern Matching with First Character`, `Length-based Categorization`, `Code-Cruncher Card Parser™`(trade marked?) and `recursive poetry parser`. All of these ideas seem beyond what a person would consider, in my opinion, when taking these lessons. 

The point is to be learning about the `switch` and making various examples of `case`. 

Oof... that was annoying. I can't believe that you have to build all those different cases... but I think it was fun to be able to refactor a little bit. 
This was the first take:
```go
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
```
Then I went and talked to the bot... and I thought, you know, that sounds about right. So I started to whittle down what the case switch looked like. I would argue that refactoring is a pretty critical step. 
```go
func FirstTurn(card1, card2, dealerCard string) (outcome string) {
	/*
		- Stand (S)
		- Hit (H)
		- Split (P)
		- Automatically win (W)
	*/
	var (
		p1           = ParseCard(card1)
		p2           = ParseCard(card2)
		p3           = ParseCard(dealerCard)
		hand         = p1 + p2
		split        = p1 == p2
		sevens_high  = p3 >= 7
		dealer_match = (p1 == p3 || p2 == p3)
	)
	switch {
	case hand > 21 && split:
		outcome = "P" // split the two aces
	case hand == 21:
		if dealer_match {
			outcome = "S"
		} else {
			outcome = "W"
		}
	case hand == 20:
		// actually, you "should" split here to double your pot
		// if split {
		// 	outcome = "P" // split
		// } else {
		// 	outcome = "S" // stand
		// }
		outcome = "S"
	case hand >= 17:
		outcome = "S" // stand on 17
	case hand >= 12:
		if sevens_high {
			outcome = "H"
		} else if split {
			outcome = "P"
		} else {
			outcome = "S"
		}
	case hand <= 11:
		outcome = "H" // hit
	default:
		outcome = "unknown"
	}
}
``` 

And so I started thinking about how we have all these switching to acheive the right outcome; that's when I thought... 'what would it look like if only 1 return value was set fore each outcome'
```go
// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	/*
		- Stand (S)
		- Hit (H)
		- Split (P)
		- Automatically win (W)
	*/
	var (
		p1           = ParseCard(card1)
		p2           = ParseCard(card2)
		p3           = ParseCard(dealerCard)
		hand         = p1 + p2
		split        = p1 == p2
		sevens_high  = p3 >= 7
		dealer_match = (p1 == p3 || p2 == p3)
	)
	switch {
	case (hand > 21 && split) || (hand >= 12 && !sevens_high && split):
		return "P"
	case (hand == 21 && dealer_match) || (hand == 20) || (hand < 21 && hand >= 17) || (hand < 21 && hand >= 12 && !sevens_high && !split):
		return "S"
	case (hand == 21 && !dealer_match):
		return "W"
	case (hand >= 12 && sevens_high) || (hand <= 11):
		return "H"
	default:
		return "unknown"
	}
}
```

Now, after taking a moment to look at how the community solved the problem, I noticed that a lot of folks wanted to make some additional functions for handling the difference between one kind of hand and another kind of hand. I think that going from a very large case switch to a smaller one with some constrained bool operators... yeah, maybe I am a little overproud of my solution. I like to think that I took the time to evaluate return values from a couple different perspectives.