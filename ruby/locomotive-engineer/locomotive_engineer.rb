class LocomotiveEngineer
  def self.generate_list_of_wagons(*args) = args

  def self.fix_list_of_wagons(each_wagons_id, missing_wagons)
    return each_wagons_id[2], *missing_wagons, *each_wagons_id[3..], each_wagons_id[0], each_wagons_id[1]
  end

  def self.add_missing_stops(routing, stops = {})
    return {**routing, stops: stops.values()}
  end

  def self.extend_route_information(route, more_route_information)
    return {**route, **more_route_information}
  end
end
