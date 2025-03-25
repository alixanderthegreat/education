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