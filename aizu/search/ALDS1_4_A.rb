def linearSearch(key, arr)
  result = false
  for i in arr
    if key == i
      result = true
      break;
    end
  end
  result
end


a_size = STDIN.gets.to_i
as = STDIN.gets.split(' ')

b_size = STDIN.gets.to_i
bs = STDIN.gets.split(' ')

n = 0
for i in bs
  if linearSearch(i, as) then
    n = n + 1
  end
end

puts n

