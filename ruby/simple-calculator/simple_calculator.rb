class SimpleCalculator
 
  UnsupportedOperation = Class.new(ArgumentError)

  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze()

  def self.calculate(first_operand, second_operand, operation)
    raise ArgumentError, "Operand must be integer" unless first_operand.is_a?(Integer) && second_operand.is_a?(Integer) 
    raise UnsupportedOperation, "Operation not allowed." unless ALLOWED_OPERATIONS.include?(operation)
    begin
      result = first_operand.public_send(operation, second_operand) 
      return "#{first_operand} #{operation} #{second_operand} = #{result}"
    rescue ZeroDivisionError
      return "Division by zero is not allowed."
    end
  end
end
