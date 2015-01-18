class Node
  def initialize(v, left, right)
    @v = v
    @left = left
    @right = right
  end

  def insert(v)
    if @v == v then
      Node.new(v, @left, @right)
    elsif v < @v
      Node.new(@v, @left.insert(v), @right)
    elsif v > @v
      Node.new(@v, @left, @right.insert(v))
    end
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
  def insert(v)
    Node.new(v, Empty.new, Empty.new)
  end

  def contains(v)
    false
  end

  def to_s
    "Empty"
  end
end

