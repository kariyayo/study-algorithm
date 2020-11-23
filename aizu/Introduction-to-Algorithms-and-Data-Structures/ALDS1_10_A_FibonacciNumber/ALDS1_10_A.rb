$memo = {}
$memo[0] = 1
$memo[1] = 1
def fib(n)
  if n == 0
    1
  elsif n == 1
    1
  else
    a = $memo[n - 1]
    if a.nil?
      a = fib(n - 1)
      $memo[n - 1] = a
    end
    b = $memo[n - 2]
    if b.nil?
      b = fib(n - 2)
      $memo[n - 2] = b
    end
    a + b
  end
end


n = STDIN.gets.to_i
puts fib(n)
