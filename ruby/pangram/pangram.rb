module Pangram
    def self.pangram?(input)
        return false if input.nil?() || input.empty?()
        input.downcase().chars().uniq().sort().join().include?("abcdefghijklmnopqrstuvwxyz")
    end
end