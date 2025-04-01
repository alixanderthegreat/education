module IsbnVerifier
    def self.valid?(input)
        return false if input.nil? || input.empty? 
        isbn = input.delete("-")
        return false if isbn.length != 10
        return false unless isbn[0..8].chars().all?() { |char|  char <= "9" && char >= "0" }
        return false unless ( ( isbn[9] <= "9" && isbn[9] >= "0" ) || isbn[9].upcase() == "X" )
        sum = isbn.chars().each_with_index().sum() do | digit, index |
            ( digit.upcase() == "X" && index == 9 ) ? 10 : digit.to_i * ( isbn.length() - index )
        end
        sum % 11 == 0 
    end
end