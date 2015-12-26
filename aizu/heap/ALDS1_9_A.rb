def left_p(i)
  2 * i
end
def right_p(i)
  2 * i + 1
end
def parent_p(i)
  i / 2
end


n = STDIN.gets.to_i
arr = STDIN.gets.split(' ')

arr.each_with_index { |a, i|
  id = i + 1
  parent_p = parent_p(id)
  left_p = left_p(id)
  right_p = right_p(id)

  parent_s =
    if parent_p != 0 && arr.size >= parent_p
      "parent key = #{arr[parent_p - 1]}, "
    else
      ""
    end

  left_s =
    if arr.size >= left_p
      "left key = #{arr[left_p - 1]}, "
    else
      ""
    end

  right_s =
    if arr.size >= right_p
      "right key = #{arr[right_p - 1]}, "
    else
      ""
    end

  puts "node #{id}: key = #{a}, #{parent_s}#{left_s}#{right_s}"
}

