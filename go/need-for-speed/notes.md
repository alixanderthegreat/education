I feel like I am tremendously familiar with structs and at this point I feel like I have learned quite a bit. It is difficult to say if I really learned anything from this other than making sure that I took the time to refactor.

This is how most people do type construction. They make each of them, one at a time. 
They also explicitly state each of the field types (eg. `int`; ||yes, we can talk about reflections later||). Here is the top rated contributes example  
```go
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}
type Track struct {
	distance int
}
```
And here is mine. I prefer to put all of my types into a single `()` group so that I can make use of the IDE's ability to  `>` roll-up the types. I also like to make use of the `,` separator to cut back on how many field type definitions there are (eg. `int` `int` `int` `int`) 
```go
type (
	Car struct {
		speed,
		battery,
		distance,
		batteryDrain int
	}
	Track struct {
		distance int
	}
)
```