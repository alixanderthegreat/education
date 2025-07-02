> I forgot how much fun these little exercises are...

Anyway, the big learn from this one was
```rb
until x == y
    x = this(x)
end
```
why this is interesting is that I discovered is that it would loop infinitely. `x` will never equal `y`, unless `x / y = 1`. What I am trying to say here is that these until loops need to be done in a very particular way to ensure that an infinite loop is not created (unless otherwise desired).
