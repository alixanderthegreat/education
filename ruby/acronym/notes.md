You can obviously tell that I write `go` code. Especially when you look at the difference between my solution, and the community's:
```rb
# Acronym class
class Acronym
  def self.abbreviate(str)
    str.scan(/\b\w/).join.upcase
  end
end
```
What's interesting is that I checked speed on these different solutions, and I found that mine is a little bit faster because we aren't doing a scan. But, we do increase the file size by 10x... so we could strip the comments out... we were able to cut the weight of the code in half with out the comments and added whitespace; so we are 4x bigger codebase and we have like... a 10% faster performance.

even if we go with something that was human/bot generated...
```rb
module Acronym
  def self.abbreviate(input)
      result = input.split(/\s|-/).map() { |part| part[0] }.join()
      result.upcase!()
      return result
  end
end
```
We have a slightly larger codebase 2x from that but with roughly the same performance... 

But if the speed was more important than the actual code size (4x), mine would be the winner
```rb
module Acronym
  class << self
    def abbreviate(input)
      result = ""
      input.strip().split().each() do |part|
        if part.include?("-")
          part.strip().split("-") do |sub|
            result += sub[0] 
          end
        else 
          result += part[0]
        end
      end
      result.upcase!()
      return result
    end
  alias_method :call, :abbreviate
  end
end
```