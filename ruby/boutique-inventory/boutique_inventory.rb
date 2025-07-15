class BoutiqueInventory
  def initialize(items)
    @items = items
  end

  def item_names
    @items.map() { |item|
      item[:name]
    }.flatten().sort()
  end

  def cheap
    @items.select() { |item|
      item[:price] < 30 
    }
  end

  def out_of_stock
    @items.select() { |item|
      item[:quantity_by_size].empty?
    }
  end

  def stock_for_item(name)
    item = @items.find { |item| 
      item[:name] == name
    }
    item[:quantity_by_size].select { |_,v|
        v > 0
    }
  end

  def total_stock
    @items.empty? ? 0 : @items.map { |item|
      item[:quantity_by_size].sum { | _,v| v.nil? ? next : v }
    }.sum()
  end

  private
  attr_reader :items
end
