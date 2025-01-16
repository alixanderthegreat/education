I know that I am actually getting better:
```go
return (kind == "car" || kind == "truck")
```
I would have never made something like this had I not taken these exercises. 

And another benefit to this exercises that I can see is that I have learned things about `go` that I wouldn't have otherwise known:
```go
a_comes_before_b := "apple" < "banana" // true
b_comes_before_a := "apple" > "banana" // false
```
I didn't know that you could lexiconographically compare strings... 🤯 

And I am noticing a pattern with the `if else if` stack; it goes upward bound, lower bound, inner bounds:
```go
	if age < 3 {
		price = originalPrice * .8
	} else if age >= 10 {
		price = originalPrice * .5
	} else {
		price = originalPrice * .7
	}
```
It was interesting to see how other devs were solving problems; I imagine that many of them had some experience in robust solutions. I also noticed who the popular guy is; lol. 