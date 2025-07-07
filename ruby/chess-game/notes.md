These trickster did two things that were rather interesting: 
1. when we call out squares in chess we say `A1`; but, the valid square option works as `1A`... that made things confusing. 
2. we were being given `"A1"` and so we had to determine that `"1"` was actually `1`... so that was tricky too. 

One thing I like about go is the static types. You can't shove `"1"` in `func` as an `arg` to a `param` of `int`. You just can't. Nope, never. Now, maybe you could take `"1"` as an `arg` for a `param` of `any`, or `interface{}`, but you would have to do a type check ...`ok := this.(int)`. 

I think you get the picture, at least I do. What was annoying was that I don't have proper error handling built into how I right ruby code yet... I need to just do that. :)
