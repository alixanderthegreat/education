module SavingsAccount
  def self.interest_rate(balance)
    rates_of_return = [
      0.5, 1.621, 2.475
    ]
    rates_of_borrowing =[
      3.213
    ]
    case balance
    when (0..999.999)
      return rates_of_return[0]
    when (1000..4999.999)
      return rates_of_return[1]
    when (5000..)
      return rates_of_return[2]
    when (..0)
      return rates_of_borrowing[0]
    end
  end

  def self.annual_balance_update(balance)
    result = (balance * (interest_rate(balance)/100))
    return balance + result
  end

  def self.years_before_desired_balance(current_balance, desired_balance)
    count = 0
    until current_balance > desired_balance
      count += 1
      current_balance = annual_balance_update(current_balance)
    end
    return count
  end
end
