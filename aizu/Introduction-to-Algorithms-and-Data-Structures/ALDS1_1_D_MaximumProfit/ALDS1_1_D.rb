n = STDIN.gets.to_i

max = -2000000000
min = STDIN.gets.to_i

(2..n).each { |i|
  x = STDIN.gets.to_i

  max = if max > x - min then max else x - min end
  min = if x < min then x else min end
}
puts max

