class Node
  def initialize(v, left, right)
    @v = v
    @left = left
    @right = right
  end

  def insert(v)
    if v < @v
      if @left.instance_of? Empty
        @left = Node.new(v, Empty.new, Empty.new)
      else
        @left.insert(v)
      end
    elsif v > @v
      if @right.instance_of? Empty
        @right = Node.new(v, Empty.new, Empty.new)
      else
        @right.insert(v)
      end
    end
    self
  end

  def contains(v)
    if @v == v then
      true
    elsif v < @v
      @left.contains(v)
    elsif v > @v
      @right.contains(v)
    end
  end

  def to_s
    "Node(#{@v}, #{@left.to_s}, #{@right.to_s})"
  end
end

class Empty
  def contains(v)
    false
  end

  def to_s
    "Empty"
  end
end

