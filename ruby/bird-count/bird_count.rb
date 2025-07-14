class BirdCount
  def self.last_week
    [0, 2, 5, 3, 7, 8, 4]
  end

  def initialize(birds_per_day)
    @birds_array = birds_per_day
  end

  def yesterday
    @birds_array[-2]
  end

  def total
    @birds_array.sum()
  end

  def busy_days
    @birds_array.count { |count| count > 4 }
  end

  def day_without_birds?
    @birds_array.any?() { |count| count == 0 }
  end
end
