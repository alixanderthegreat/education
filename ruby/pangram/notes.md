I would have never done it like this:
```rb
input.downcase().chars().uniq().sort().join().include?("abcdefghijklmnopqrstuvwxyz")
```
for me, I am not thinking... okay, we have this input... and it calls downcase, which means it is probably a string and so lets call chars to break it into an array of characters and then we call uniq on the array to trim out any of the additional letters, and then we sort them in lexicon order so that we can join the string and it should include all of these letters....     

yeah, that's not how I think. This was a lot like isogram. 

```rb
module Pangram
    def self.pangram?(input)
        return false if input.nil?() || input.empty?()
        input = input.strip()
        input.downcase!()
        for letter in "abcdefghijklmnopqrstuvwxyz".chars()
            found = false
            for char in input.chars()
                if char == letter
                    found = true
                    break
                end
            end
            if found == false
                return found
            end
        end
        return found
    end
end
```
In this, we go through each letter and check to see if, from all of the input characters, there is a match with the letter set. 

As you might guess, this was very inventive. 

like... in go, you have to think in loops; you don't have conventional uniq, or join ()
```go
package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

var pangram = "abcdefghijklmnopqrstuvwxyz"

func main() {
	input := "the 1 quick brown fox jumps over the 2 lazy dogs"
	fmt.Println(input, is_pangram(input))
}
func is_pangram(input string) bool {

	var letter_map = make(map[string]bool)

	for _, ch := range strings.ToLower(input) {

		if !unicode.IsLetter(ch) { // if it isn't a letter...
			continue
		}

		if letter_map[string(ch)] { // if it is already on the map...
			continue
		}

		letter_map[string(ch)] = true

	}

	var sorting_bin []string

	for char := range letter_map {
		sorting_bin = append(sorting_bin, char)
	}

	sort.Strings(sorting_bin)

	var result string

	for _, char := range sorting_bin {
		result += char
	}
	return strings.Contains(strings.TrimSpace(result), pangram)
}
```

Each of them are different ways of dealing with the same problem; each of them with their own way of being constructed. 
