module Blackjack
  def self.parse_card(card)
    ["joker", "one", "two", "three", "four", "five",
    "six", "seven", "eight", "nine", "ten","ace"].index(card) || 10 # jack, queen and king
  end

  def self.card_range(card1, card2)
    raise "Please implement the Blackjack.card_range method"
  end

  def self.first_turn(card1, card2, dealer_card)
    raise "Please implement the Blackjack.first_turn method"
  end
end
