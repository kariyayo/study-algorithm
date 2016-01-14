g = []

n = STDIN.gets.to_i
n.times.each {
  as = STDIN.gets.split(' ').map{|x| x.to_i }
  g[as[0]] = as.drop(2)
}

stack = []
mode = [] # nil: 未訪問, 1: 訪問済み, 2: 隣接も全て訪問済み
start_t = []
end_t = []

t = 0
for i in 1..n
  if mode[i] == nil
    t = t + 1
    v = i
    mode[v] = 1
    start_t[v] = t
    stack.push(v)
    while stack.size != 0
      t = t + 1
      adjacencies = g[v]
      has_yet = false
      for a in adjacencies
        if mode[a] == nil
          # 初訪問
          mode[a] = 1
          v = a
          start_t[v] = t
          stack.push(v)
          has_yet = true
          break
        end
      end
      if !has_yet
        # 隣接ノードが全部訪問ずみ
        mode[v] = 2
        end_t[v] = t
        stack.pop
        v = stack.last
      end
    end
  end
end

for i in 1..n
  puts "#{i} #{start_t[i]} #{end_t[i]}"
end
