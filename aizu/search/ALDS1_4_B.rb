def binarySearch(key, arr)
  start = 0
  last = arr.size - 1

  result = false
  while (last - start) != 0
    m = (last - start) / 2 + start
    v = arr[m].to_i

    if key.to_i == v then
      result = true
      break
    elsif key.to_i < v then
      last = m - 1
    else
      start = m + 1
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
  if binarySearch(i, as) then
    n = n + 1
  end
end

puts n

