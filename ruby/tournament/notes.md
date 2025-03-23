I learned so much from this experience, and I would be remiss if I said that I didn't use the bot's help. I did. There is something very tricky about how expanisve the ruby universe is, albeit the concepts in go can be seen in the way that I build. I think that I will continue to use the ideas that I learned in go to help me extrapolate further while I am building into these exercises. 

for instance. I didn't really understand that something can be a hash... and a hash is kind of like... kind of like a map... anyway. 
```rb
  data = Hash.new { |h, k| h[k] = Hash.new(0) }
```
This was something that was very different to me, but I can relate it to a map[string]map[string]int, at least that is what it is in `go`. As it is for ruby. I think that learning how to role through blocks (iterations) was helpful in this exercise. 

The other thing that was interesting was:
```rb
    header =  "%-30s | %2s | %2s | %2s | %2s | %2s\n" % [  
```
I didn't know that Printf tokens were still valid. I was surprised that you can just format a string on the fly... so to speak, instead of having to call fmt.Sprintf or what ever. 

Another thing that I learned was that I can use a `*` to spill out all the other ideas about something:
```rb
  # Update data with stats
  update_stats.call(data, update, *stats)
```
basically works the same as `stats...` as a variadic in go. So that was helpful to learn. I also learned how to make lambda functions or private functions:
```rb
update_stats = ->(data, team, wins, losses, draw, points) do
  data[team]["W"] += wins
  data[team]["L"] += losses
  data[team]["D"] += draw
  data[team]["P"] += points
end
update_stats.call(data, home, 1, 0, 0, 3)
```
When I reviewed the community solutions for this problem... oof; no wonder ruby is the corporate language, there is almost no "right" way of doing it. There is a conventionally way; and I'm not sure I could really find that unless we are saying that brevity (D.R.Y.) leads to the most conventional code. But even then...