It took some time to learn about the idea of 0th indexed items. I am not sure why that took so long to understand that the first `int` position of the index is the `origin`; which, counter-intuitive as it may seem, is the most logical position for the start of a list. Yes, it is the first/1st postion; but its index is at the origin. 

This was a particularly interesting problem simply on account of the need to append something onto the `slice`. I wasn't really expecting that, to be honest. I mean, I did read it; but I hadn't considered that it would be immediately relevant. 

This one was particularly easy, but I guess that the IDE really did most of the work here. 

I don't typically do cut actions; like, splits and joins and cuts. I would have normally done some kind of for loop. Not sure how I would do that here... but that's normally how I would have done it: 
```go
// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) (newSlice []int) {
	if len(slice) < index || index < 0 {
		return slice
	}
	for i, this := range slice {
		if i == index {
			continue
		}
		newSlice = append(newSlice, this)
	}
	return newSlice
}
```
But, as you might have guessed, this wouldn't be very helpful if you had millions of items that were in the slice. But I would say that I have experienced a behavior change in the way that I am making `go` code. Which, I guess, is why we do exercises. 