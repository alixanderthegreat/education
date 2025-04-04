class AssemblyLine
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    base = 221 * @speed
    return base unless @speed > 4

    modifier = 0.0
    modifier = 0.77 if @speed == 10 
    modifier = 0.80 if @speed == 9 
    modifier = 0.90 if @speed > 4 && @speed < 9
    result = (base * modifier)
    return result
  end

  def working_items_per_minute
    raise 'Please implement the AssemblyLine#working_items_per_minute method'
  end
end
