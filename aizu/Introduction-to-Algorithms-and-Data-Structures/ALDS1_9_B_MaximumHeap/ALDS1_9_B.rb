def left_p(i)
  2 * i
end
def right_p(i)
  2 * i + 1
end
def parent_p(i)
  i / 2
end

def max_heapfy(arr, i)
  l = left_p(i)
  r = right_p(i)

  if l <= arr.size && arr[l - 1] > arr[i - 1]
    largest = l
  else
    largest = i
  end

  if r <= arr.size && arr[r - 1] > arr[largest - 1]
    largest = r
  end

  if largest != i
    tmp = arr[largest - 1]
    arr[largest - 1] = arr[i - 1]
    arr[i - 1] = tmp
    max_heapfy(arr, largest)
  end
end

def build_max_heap(arr)
  i = arr.size / 2 + 1
  while i > 0
    max_heapfy(arr, i)
    i = i - 1
  end
end


n = STDIN.gets.to_i
arr = STDIN.gets.split(' ').map { |x| x.to_i }

build_max_heap(arr)
puts ' ' + arr.join(' ')

