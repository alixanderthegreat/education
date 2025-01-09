I was rather surprised, although I am not quite finished yet, by how immensely powerful string manipulation can be. I know that I am walking into a difficult place by being willing to learn something that is way over my head at the moment: `regex` for one and instance `@variables` for another. 

Learning about instance variabels is very interesting in its `self`. I am thinking about whether or not the operation that are performed on `@line` would cause mutations to occur... I am wondering about if there is need to make use of memory to work from a copy... lots to think about but little experience in this regard 

`regex` on the other hand adds another level of complexity that I didn't really know was possible. I spent the better part of the day thinking about and talking with friends about the following solution; which, to be fair; really isn't mine... I think that it would have been nice to have thought like this, but it was my buddies solution - he has experience in perl regex and has a penchant for it.  
```rb
@line[/\[(.*?)\]/, 1]
```
I decided to take a moment and break down his example a little bit more, I think that it has merit beyond be a solution  
```rb
"[ERROR]: Stack overflow"[ # opens the array of the string
  / # opens the regex 
  \ # escapes the special character [
  [ # finds the first case of this character
  ( # opens the grouping to grab
  . # grabs first character , in our case E for error
  * # grabs all of the following characters, in wildcard fashion
  ? # selects the preceding character to the next match, R before ]
  ) # closes the group
  \ # escapes the special character 
  ] # the next character to match 
  / # closes the search
  , # delimits the argument 
  1 # how many iterations of the expression to find
] # closes the array
```

I then took some time to look over the kinds of things that I saw in the community to assess how they solved things versus how I my submited solution compared. I think that the comparison with the community is really helpful. I was able to look over solutions that made sense and solutions that seemed obtuse. Many of them were similar, and I suspect that the bot made many of them. I think that some people used the bot to make their own creative refactorings in order to showcase themselves as different and unique. But then again, what do I know... I can't say it is particularly wrong if it is a passing submission (or can I...).