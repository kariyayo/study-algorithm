class Node

  attr_accessor:id, :parent, :child_left, :right_sibling, :depth

  def initialize(id)
    @id = id
    @parent = nil
    @child_left = nil
    @right_sibling = nil
    @depth = 0
  end

  def depth
    if @parent.nil?
      0
    else
      @parent.depth + 1
    end
  end

  def type
    if @parent == nil
      "root"
    elsif @child_left.nil?
      "leaf"
    else
      "internal node"
    end
  end

  def to_s
    parent_id =
      if @parent.nil?
        -1
      else
        @parent.id
      end
    "node #{@id}: parent = #{parent_id}, depth = #{depth}, #{type}, [#{children_to_s.join(', ')}]"
  end

  def children_to_s
    ss = []
    if @child_left.nil?
      ss
    else
      a = @child_left
      while a != nil
        ss.push(a.id.to_s)
        a = a.right_sibling
      end
      ss
    end
  end
end

n = STDIN.gets.to_i

tree = []
for i in 0..(n - 1)
  as = STDIN.gets.split(' ')
  id = as[0].to_i
  m = as[1].to_i
  children = as.drop(2).map { |x| x.to_i }

  node = tree[id]
  if node.nil?
    node = Node.new(id)
    tree[id] = node
  end

  l = 0
  children.each_with_index { |c, j|
    child_node = tree[c]
    if child_node.nil?
      child_node = Node.new(c)
      tree[c] = child_node
    end
    if j == 0
      node.child_left = child_node
    else
      tree[l].right_sibling = child_node
    end
    l = c
    tree[c].parent = node
  }
end


for i in 0..(n-1)
  puts tree[i].to_s
end

