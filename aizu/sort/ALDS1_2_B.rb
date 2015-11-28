def selectionSort(arr)
  n = 0
  for i in 0..(arr.size - 1)
    min_i = i
    for j in i..(arr.size - 1)
      if arr[j] < arr[min_i]
        min_i = j
      end
    end
    if i != min_i
      x = arr[min_i]
      arr[min_i] = arr[i]
      arr[i] = x
      n = n + 1
    end
  end
  puts arr.join(' ')
  puts n
end

n= STDIN.gets.to_i
arr = STDIN.gets.split(' ').map{|x| x.to_i}
selectionSort(arr)

