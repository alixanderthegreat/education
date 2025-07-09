module Blackjack
  def self.parse_card(card)
    ["joker", "one", "two", "three", "four", "five",
    "six", "seven", "eight", "nine", "ten","ace"].index(card) || 10 # jack, queen and king
  end

  def self.card_range(card1, card2)
    low = (4..11)
    mid = (12..16)
    high = (17..20)
    bj = 21
    case parse_card(card1) + parse_card(card2)
    when low then 'low'
    when mid then 'mid'
    when high then 'high'
    when bj then 'blackjack'
    end
  end

  def self.first_turn(card1, card2, dealer_card)
    raise "Please implement the Blackjack.first_turn method"
  end
end
