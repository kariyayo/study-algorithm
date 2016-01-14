def partition(arr, p, r)
  q = arr[r]
  i = p - 1 # 交換対処
  for j in p..(r - 1) # 比較対象
    if arr[j] <= q then
      i = i + 1
      tmp = arr[i]
      arr[i] = arr[j]
      arr[j] = tmp
    end
  end
  z = i + 1
  tmp = arr[z]
  arr[z] = arr[r]
  arr[r] = tmp
  arr[z] = "[#{arr[z]}]"
  puts arr.join(' ')
  z
end


n = STDIN.gets.to_i
arr = STDIN.gets.split(' ').map { |x| x.to_i }

partition(arr, 0, arr.size - 1)


