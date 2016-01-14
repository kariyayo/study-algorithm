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

def heap_increase_key(arr, key)
  now = arr.last
  arr.push(key)
  if now != nil && key > now
    i = arr.size
    while i > 1 && arr[parent_p(i) - 1] < arr[i - 1]
      tmp = arr[i - 1]
      arr[i - 1] = arr[parent_p(i) - 1]
      arr[parent_p(i) - 1] = tmp
      i = parent_p(i)
    end
  end
end

def insert(arr, key)
  heap_increase_key(arr, key)
end

def heap_extra_max(arr)
  if arr.size < 1
    nil
  else
    max = arr[0]
    arr[0] = arr.last
    arr.delete_at(arr.size - 1)
    max_heapfy(arr, 0)
    max
  end
end


arr = []
s = STDIN.gets
while s != nil && s != "end"
  ss = s.split(' ')
  if ss[0] == "insert"
    insert(arr, ss[1].to_i)
  elsif ss[0] == "extract"
    puts heap_extra_max(arr)
  end
  s = STDIN.gets
end

