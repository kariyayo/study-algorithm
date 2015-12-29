def lcs(xs, ys)
  memo = []
  memo[0] = []

  xl = xs.size
  yl = ys.size

  for i in 1..xl
    memo[i] = []
#    memo[i][0] = 0
  end

#  for j in 1..yl
#    memo[0] = []
#    memo[0][j] = 0
#  end

  max = 0
  for i in 1..xl
    for j in 1..yl
      if xs[i-1] == ys[j-1]
        a = memo[i - 1][j - 1]
        memo[i][j] = if a == nil then 1 else a + 1 end
      else
        a = if memo[i - 1][j].nil? then 0 else memo[i - 1][j] end
        b = if memo[i][j - 1].nil? then 0 else memo[i][j - 1] end
        if a >= b
          memo[i][j] = a
        else
          memo[i][j] = b
        end
      end
      if max < memo[i][j]
        max = memo[i][j]
      end
    end
  end
  max
end


n = STDIN.gets.to_i

n.times.each {
  a = STDIN.gets
  b = STDIN.gets
  xs = a.split("")
  ys = b.split("")

  xs.delete_at(xs.size - 1)
  ys.delete_at(ys.size - 1)
  puts lcs(xs, ys)
}

