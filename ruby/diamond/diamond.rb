module Diamond
  def self.make_diamond(input)
    letter = input.upcase()
    starting_position = 'A'.ord()
    letter_position = letter.ord()
    inset = ( letter_position - starting_position ) 
    tiles = ( 'A'..letter ).to_a()    
    space = ' '    
    chevron = tiles.each_with_index().map() { | tile , index |
      space_outside = space * ( inset - index )
      space_inside = space * ( ( index * 2 ) - 1 ) if index.positive?()
      line = space_outside + tile
      line += space_inside + tile if index.positive?()
      line += space_outside + "\n"
    } 
    diamond = chevron + chevron[ 0 ... -1 ].reverse()
    return diamond.join()
  end
end
