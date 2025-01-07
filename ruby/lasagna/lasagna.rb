class Lasagna
  EXPECTED_MINUTES_IN_OVEN = 40

  def remaining_minutes_in_oven(
    actual_minutes_in_oven # can do this, which is neat
  )
    # multi-line math
    EXPECTED_MINUTES_IN_OVEN -
      actual_minutes_in_oven
  end

  def preparation_time_in_minutes(layers)
    # assuming each layer takes you 2 minutes to prepare
    minutes_per_layer = 2
    # multiply assumption and layers
    minutes_per_layer * layers
  end

  def total_time_in_minutes(
    number_of_layers:,
    actual_minutes_in_oven: # the feild
    # interesting
  )
    preparation_time_in_minutes(number_of_layers) +
      actual_minutes_in_oven
  end
end
