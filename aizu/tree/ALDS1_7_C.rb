class Node

  attr_accessor:id, :parent, :left, :right, :depth, :height, :sibling

  def initialize(id)
    @id = id
    @parent = nil
    @left = nil
    @right = nil
    @depth = 0
    @height = 0
    @sibling = nil
  end

  def depth
    if @parent.nil?
      0
    else
      @parent.depth + 1
    end
  end

  def height
    left_height =
      if @left.nil?
        -1
      else
        @left.height
      end
    right_height =
      if @right.nil?
        -1
      else
        @right.height
      end
    if left_height > right_height
      left_height + 1
    else
      right_height + 1
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

    sibling =
      if @sibling.nil?
        -1
      else
        @sibling.id
      end

    degree = 0
    if @left != nil
      degree = degree + 1
    end
    if @right != nil
      degree = degree + 1
    end

    type =
      if @parent.nil?
        "root"
      elsif @left.nil? && @right.nil?
        "leaf"
      else
        "internal node"
      end

    "node #{@id}: parent = #{parent_id}, sibling = #{sibling}, degree = #{degree}, depth = #{depth}, height = #{height}, #{type}"
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

def preorder(node, arr)
  if node != nil
    arr.push(node.id)
    l = node.left
    preorder(l, arr)
    r = node.right
    preorder(r, arr)
  end
end

def inorder(node, arr)
  if node != nil
    l = node.left
    inorder(l, arr)
    arr.push(node.id)
    r = node.right
    inorder(r, arr)
  end
end

def postorder(node, arr)
  if node != nil
    l = node.left
    postorder(l, arr)
    r = node.right
    postorder(r, arr)
    arr.push(node.id)
  end
end


n = STDIN.gets.to_i

tree = []
for i in 0..(n - 1)
  as = STDIN.gets.split(' ')
  id = as[0].to_i
  left_id = as[1].to_i
  right_id = as[2].to_i

  node = tree[id]
  if node.nil?
    node = Node.new(id)
    tree[id] = node
  end

  if left_id != -1
    left = tree[left_id]
    if left.nil?
      left = Node.new(left_id)
      tree[left_id] = left
    end
    left.parent = node
    node.left = left
  end

  if right_id != -1
    right = tree[right_id]
    if right.nil?
      right = Node.new(right_id)
      tree[right_id] = right
    end
    right.parent = node
    node.right = right
  end

  if left_id != -1 && right_id != -1
    node.left.sibling = node.right
    node.right.sibling = node.left
  end

end

root = tree.find {|node|
  node.type == "root"
}

puts "Preorder"
as = [""]
preorder(root, as)
puts as.join(" ")

puts "Inorder"
bs = [""]
inorder(root, bs)
puts bs.join(" ")

puts "Postorder"
cs = [""]
postorder(root, cs)
puts cs.join(" ")

