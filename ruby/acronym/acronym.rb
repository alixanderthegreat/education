module Acronym
  class << self
    def abbreviate(input)
      result = ""
      first_letter = 0
      input.strip().split().each() do |part|
        if part.include?("-")
          part.strip().split("-") do |sub|
            result += sub[first_letter] 
          end
          result += part[first_letter]
        end
      end
      result.upcase!()
      return result
    end
  alias_method :call, :abbreviate
  end
end