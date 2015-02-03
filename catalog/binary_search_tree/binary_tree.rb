class Node
  def initialize(v, left, right)
    @v = v
    @left = left
    @right = right
  end

  def insert(v)
    if v < @v
      if @left.empty?
        @left = Node.new(v, Empty.new, Empty.new)
      else
        @left.insert(v)
      end
    elsif v > @v
      if @right.empty?
        @right = Node.new(v, Empty.new, Empty.new)
      else
        @right.insert(v)
      end
    end
    self
  end

  def remove(v)
    if v < @v
      if v == @left.v
        if @left.has_both?
          max = find_max_value(@left.right)
          @left.remove(max)
          @left.v = max
        else
          @left = child_or_empty(@left)
        end
      else
        @left.remove(v)
      end
    elsif v > @v
      if v == @right.v
        if @right.has_both?
          max = find_max_value(@right.right)
          @right.remove(max)
          @right.v = max
        else
          @right = child_or_empty(@right)
        end
      else
        @right.remove(v)
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

  def has_both?
    !@left.empty? && !@right.empty?
  end

  def no_child?
    @left.empty? && @right.empty?
  end

  def left_only?
    !@left.empty? && @right.empty?
  end

  def right_only?
    @left.empty? && !@right.empty?
  end

  def empty?
    self.instance_of? Empty
  end

  def inspect
    "Node(#{@v}, #{@left.inspect}, #{@right.inspect})"
  end

private

  def child_or_empty(node)
    if node.no_child?
      Empty.new
    elsif node.right_only?
      node.right
    elsif node.left_only?
      node.left
    end
  end


  def find_max_value(node)
    if node.left.empty? && node.right.empty?
      node.v
    elsif node.left.empty?
      [node.v, find_max_value(node.right)].max
    elsif node.right.empty?
      [node.v, find_max_value(node.left)].max
    else
      [node.v, find_max_value(node.left), find_max_value(node.right)].max
    end
  end
end

class Empty
  def contains(v)
    false
  end

  def empty?
    self.instance_of? Empty
  end

  def inspect
    "Empty"
  end
end

