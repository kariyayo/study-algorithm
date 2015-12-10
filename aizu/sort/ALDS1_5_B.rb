$n = 0

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
      $n = $n + 1
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

n = STDIN.gets.to_i
arr = STDIN.gets.split(' ')
puts mergeSort(arr).join(' ')
puts $n

