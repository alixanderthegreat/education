I had to follow this article in order to complete the first of the functions to implement: https://yourbasic.org/golang/generate-number-random-range/
I was suprised that I wasn't able to use something like this before; at least, not with respect to the likelihood of landing a 0 versus not being able to.  

The next function was a little more difficult, generating wand energy... had to follow https://pkg.go.dev/math/rand#NormFloat64 to distribute the output within the desiredStdDev and desiredMean.

This swap anon function is interesting. I will admit that I always think of the sorter function as it pertains to slices, but the swap... that's an intesting thought. 
```go
rand.Shuffle(len(x), func(i, j int) {
	x[i], x[j] = x[j], x[i]
})
```

After reviewing the community, I was suprised to find that Generate Wand energy is simply takin the upward bound 12... and then times it by a number between 0.0 -> 1.0... That was definately more elegant than my solution
