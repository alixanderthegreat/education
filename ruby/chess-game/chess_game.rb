module Chess
  RANKS = (1..8).freeze()
  FILES = ('A'..'H').freeze()

  def self.valid_square?(rank, file)
    ( RANKS.include?(rank) && FILES.include?(file) )
  end

  def self.nick_name(first_name, last_name)
    (first_name[..1] + last_name[( last_name.length() - 2 )..]).upcase()
  end

  def self.move_message(first_name, last_name, square)
    nickname = nick_name(first_name, last_name)
    valid_square?(square[1].to_i(), square[0]) ? 
    "#{nickname} moved to #{square}" : 
    "#{nickname} attempted to move to #{square}, but that is not a valid square"
  end
end
