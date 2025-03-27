module Isogram
  class << self
    def isogram?(input)
      return true if input.nil? || input.empty?
      input = input.strip()
      input.downcase!()
      hash_map = Hash.new(0) 
      input.chars.each() do |letter|
        next if letter == " " || letter == "-" 
          if hash_map[letter] == 0
            hash_map[letter] += 1
          else
            return false
          end
        end
      return true
    end
  end
end