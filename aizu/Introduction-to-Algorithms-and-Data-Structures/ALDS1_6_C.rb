def merge(sorted_left, sorted_right)
  arr = []
  l_p = r_p = 0
  max = 1000000000
  sorted_left.push(max)
  sorted_right.push(max)
  l_v = r_v = 0
  while l_v != max || r_v != max
    flag = true
    l_v = sorted_left[l_p].to_i
    while flag
      r_v = sorted_right[r_p].to_i
      if l_v == max && r_v == max then
        return arr
      elsif l_v < r_v then
        arr.push(l_v)
        l_p = l_p + 1
        flag = false
      else
        arr.push(r_v)
        r_p = r_p + 1
      end
    end
  end
  arr
end

def mergeSort(arr)
  if arr.size == 0 then
    arr
  elsif arr.size == 1 then
    arr
  else
    m = arr.size / 2
    left = arr.take(m)
    right = arr.drop(m)
    merge(mergeSort(left), mergeSort(right))
  end
end


def partition(arr, p, r)
  q = arr[r]
  i = p - 1 # 交換対処
  for j in p..(r - 1) # 比較対象
    if arr[j][1] <= q[1] then
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
  z
end


def quickSort(arr, p, r)
  if p < r then
    q = partition(arr, p, r)
    quickSort(arr, p, q-1)
    quickSort(arr, q+1, r)
  end
end


n = STDIN.gets.to_i

arr = []
for i in 0..(n - 1)
  c = STDIN.gets.split(' ')
  c[1] = c[1].to_i
  arr.push(c)
end


quickSort(arr, 0, arr.size - 1)


puts "####"
arr.each { |c|
  puts c.join(' ')
}

