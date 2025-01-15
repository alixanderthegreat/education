I am glad I am trying to do a speed run using go because that will mean that I will likely be able to cross examine the concepts of `go` against the concepts that are found in other languages. With that being said... 
> let's not get too cocky.

For instance, `%` is a modulus which returns the remainder after division... that was new for me. And speaking of percent, I almost completely forgot that when you are trying to determine the percent out of 100, you have to divide by 100...

It was fun actually using the modulus operator on this problem. I have never had to use something like that before to solve a problem. I particularly liked being able to create groups and solve the cost. I can kind of see that I could solve for cost per minute... but that's another thing entirely. 

What makes me feel silly is that the community solution to this problem was
```go
func CalculateCost(carsCount int) uint {
	return uint(carsCount/10*95000 + carsCount%10*10000)
}
```
To me, there is nothing wrong with this; it solves the test perfectly and it does it in the least lines of code. So I guess, why I feel so silly is that my solution to the test was... explicit and maybe long winded; but I sensed that it solves it in a more readable , less mathy kind of way.