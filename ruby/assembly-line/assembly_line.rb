class AssemblyLine
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour()
    base = 221 * @speed
    return base unless @speed > 4

    modifier = 0.0
    modifier = 0.77 if @speed == 10 
    modifier = 0.80 if @speed == 9 
    modifier = 0.90 if @speed < 9 && @speed > 4 
    result = (base * modifier)
    return result
  end

  def working_items_per_minute()
    result = (production_rate_per_hour() / 60).floor()
    return result
  end
end
