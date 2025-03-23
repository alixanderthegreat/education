module Tournament
    class << self
        def tally(input)

          # Scrub the input
          input = input.strip()

          match_results = [
            "W",    # - W: Matches Won
            "D",    # - D: Matches Drawn (Tied)
            "L",    # - L: Matches Lost
            "P"     # - P: Points
          ]

          # Build a header for the result
          header =  "%-30s | %2s | %2s | %2s | %2s | %2s\n" % [  
            "Team", 
            "MP",   # - MP: Matches Played
            *match_results
          ]

          # Return only the header if the input is nil or empty
          return header if input.nil? || input.empty?

          # Fill the result with the header
          result = header

          # Init data to contain sums
          data = Hash.new { |h, k| h[k] = Hash.new(0) }

          # Make private lambda to update stats
          update_stats = ->(data, team, wins, losses, draw, points) do
            # update the data map
            data[team]["W"] += wins
            data[team]["L"] += losses
            data[team]["D"] += draw
            data[team]["P"] += points
          end
          
          # Iterate over the input
          input.each_line() do |line|

            # Gather the parts from the line 
            home, away, outcome = line.strip().split(";")

            # If you are here...
            [home, away].each() { |team| # teams have played the game
              data[team]["MP"] += 1 # and each has played a match
            }

            # Define stat allocations
            win = [ 1, 0, 0, 3 ]
            loss = [ 0, 1, 0, 0 ]
            draw = [ 0, 0, 1, 1 ]

            # Build outcome map
            outcomes = {
              'win'  => { home: win, away: loss },
              'loss' => { home: loss, away: win },
              'draw' => { home: draw, away: draw }
            }

            # Check if the outcome is valid, and retrieve the corresponding points
            if outcomes.key?(outcome)

              # For each of the outcomes keys, eg :home , :away
              outcomes[outcome].each do |team, stats|

                # Define update operator
                update = team == :home ? home : away

                # Update data with stats
                update_stats.call(data, update, *stats)

              end

            else # Return an error
              raise "Invalid outcome: #{outcome}. Please check the input."
            end
          end

          # Sort the data based on points
          data = data.sort_by() { |team, stats| [-stats["P"], team] }

          # Now build the formatted result
          data.each() do |team, stats|
              result += "%-30s | %2d | %2d | %2d | %2d | %2d\n" % [
                  team, stats["MP"], stats["W"], stats["D"], stats["L"], stats["P"]
          ]
          end
        return result
      end
    alias_method :call, :tally
  end
end
