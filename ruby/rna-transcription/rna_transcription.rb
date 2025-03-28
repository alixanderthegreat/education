module Complement
  class << self
    def of_dna(input)
      return input if input.nil? || input.empty?
      result = input.chars().map() do | letter | 
        case letter
          when 'C' then 'G'
          when 'G' then 'C'
          when 'T' then 'A'
          when 'A' then 'U'
          else letter
        end
      end
      result.join()
    end
  end
end