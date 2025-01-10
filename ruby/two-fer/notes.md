I will admit that I would have never figured this one out if I hadn't gone to [the github link](https://github.com/exercism/problem-specifications/issues/757) that is provided in the README.md. 

Coming from a background in go, I would have never know that you can overload arguments; [that is simply not a thing in go](https://go.dev/doc/faq#overloading). 
```go
	func(name string = "you can't do this in go"){
        if name == "" {
            name = "you have to do it inside the func"
        }
    }()
```
You have to make some kind of `if` construction once the function is opened and then move on... But in ruby, and apparently [in many other "modern" programming languages](https://stackoverflow.com/questions/19612449/default-value-in-gos-method#comment77755758_23650312), you can do the following:
```rb
  def self.two_fer(name = "you")
```
This was a real eye-opener moment for me because I had never considered having default values define in the argument prior to today. Really, this was an immensely interesting thought and I am happy that I was able to cleanly understand both the pros and cons to what this means moving forward.

And after I have submitted my results... holy cow! I had no idea that ruby has so many little nuances to it. Like... a Ternary Operator `( ? : )`, or that you can just omit parentheses `( )` when establishing arguments `method_name arg = "default"`.  My mind basically melts everytime I look at how the community of problem solvers solve problems. 

The following was my favorite solution by far. Extermely simply and does exactly what it is supposed to do. 
```rb
class TwoFer
  def self.two_fer(n = "you") = "One for #{n}, one for me."
end
```
Implicity, such a strange landscape. I guess my final take away here is that what is implicit is not well understood if you don't know what is interpretable and how it will be interpreted.  
