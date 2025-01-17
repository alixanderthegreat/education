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