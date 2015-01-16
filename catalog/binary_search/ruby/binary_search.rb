# x: int, xs: int[] ordered asc
def contains(x, xs)
  if xs.length == 0 then
    false
  elsif xs.length == 1 then
    if xs[0] == x then
      true
    else
      false
    end
  else
    n = xs.length / 2
    if xs[n] == x then
      true
    elsif x < xs[n] then
      contains(x, xs.take(n))
    else
      contains(x, xs.drop(n))
    end
  end
end

