def insertionSort(arr)
  puts arr.join(' ')
  (1..(arr.length - 1)).each {|i|
    v = arr[i]
    j = i - 1
    while j >= 0 and arr[j] > v
      arr[j + 1] = arr[j]
      j = j - 1
    end
    arr[j+1] = v
    puts arr.join(' ')
  }
end

n = STDIN.gets.to_i
arr = STDIN.gets.split(' ').map{ |x| x.to_i }
insertionSort(arr)

