module Port
  IDENTIFIER = :PALE
  def self.get_identifier(city)
    city.chars()[0..3].join("").upcase().to_sym()
  end

  def self.get_terminal(ship_identifier)
    result = ship_identifier[0..2].to_sym()
    isResult = (result == :OIL || result == :GAS)
    return isResult ? :A : :B 
    end
end
