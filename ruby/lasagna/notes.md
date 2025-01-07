The first thing that I noticed was my background in `go`. 

I went to establish `EXPECTED_MINUTES_IN_OVEN` as a `const`, like so:
```go
const EXPECTED_MINUTES_IN_OVEN = 40
```
But `ruby` doesn't work like this...
```ruby
lasagna.rb:3:in `<class:Lasagna>': undefined method `const' for Lasagna:Class (NoMethodError)

  const EXPECTED_MINUTES_IN_OVEN = 40
  ^^^^^
Did you mean?  const_set
```
Instead, it interprets the `CAP_CASE` as a `const`, which is new to me. 

The next thing that I noticed was `return`. That you can implicitly, or explicitly, `return` values, but you can't define `return` `type` like you would in `go`; I am guessing that this has this implicit design has its advantages for speed... but it does create some down stream questions:
- What happens when you have multiple values that need to `return`?
- When not explicitly defining the value to be sent with the `return`, does the interpreter assume that the last value of the method is the one to `return`? 

> _After note:_ 
> _When using `return`, you will see that: "Redundant `return` detected"_

But with that I move on to the next thought, as it doesn't make sense to beat a dead horse on this; how about whitespace? We have already seen that the interpreter is interested in white space in the parsing of command line arguments (see: `hello-world/notes.md`), but what about in the script?
```ruby
  THIS-that
  THIS- that
  THIS -that
  THIS - that
``` 
It took a few minutes to set up the [standard ruby linter](https://marketplace.visualstudio.com/items?itemName=testdouble.vscode-standard-ruby), and install `ruby-bundler`. Once that was done though, it was interesting to watch how the linter interacts with the code. Unlike the `go` linter, that will force formatting, the `ruby` linter simply underlines the `PROBLEMS` with the code. What immediately stood out to me was that `THIS - that` doesn't actually have some kind of linting convention until we start to use multi-line operations. 
```ruby
  # surronding space missing for operator 
  THIS-
  that
  
  # surronding space missing for operator & trailing whitespace error
  THIS- 
  that

  # no errors
  THIS -
  that

  # trailing whitespace error
  THIS - 
  that
  
```
The next thing that I learned that was particularly interesting was that we can multi-line the method argument. 
```ruby
  # single line
  def what(argument) # comment

  # multi-line
  def what(
    argument # comment
  )
```
On the surface this doesn't seem to be all that helpful unless you are like me, and you like to include a tremendously verbose amount of words as commentary to the lines of code that you write. Personally, I like that I know this fact about Ruby, and that the linter really doesn't have a problem with it. That is until...
```ruby
  # multi-line, multi-arg
  def what( # cant comment
    # cant comment
    argument, # cant comment
    # cant comment
    argument # can comment
    # can comment
  )
```