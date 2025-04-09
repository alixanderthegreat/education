module RotationalCipher
    class << self
        LOWERCIPHER = ('a'..'z').to_a().join()
        UPPERCIPHER = LOWERCIPHER.upcase()
        def rotate(input, key)
            message = []
            return input if key.zero?() || key.nil?()
            begin
                for char in input.each_char()
                    case char
                    when "A".."Z"
                        message.append(UPPERCIPHER[(UPPERCIPHER.index(char) + key) % 26])
                    when "a".."z"
                        message.append(LOWERCIPHER[(LOWERCIPHER.index(char) + key) % 26])
                    else 
                        message.append(char)
                    end
                end
            rescue => e 
                puts e.message
            end
            result = message.join()
            return result
        end
    end
end