module Acronym
  class << self
    def abbreviate(input)

      # Let's have this result contain an empty string
      result = ""

      # Let's select the first letter of the array
      first_letter = 0

      # Iterate through the input
      input.strip().split().each() do |part|

        # Handle the case when a hyphen `-`` is used 
        if part.include?("-")

          # Strip and split the part into sub strings
          part.strip().split("-") do |sub|

            # Concatenate the first letter of the each sub string to the result
            result += sub[first_letter]
          
          end
        else # Concatenate the first letter of each part string to the result
          result += part[first_letter]
        end
      end
      
      # Modify result to uppercase
      result.upcase!()

      # Return result
      return result
    end
  alias_method :call, :abbreviate
  end
end