g = []

n = STDIN.gets.to_i
n.times.each {
  as = STDIN.gets.split(' ').map{|x| x.to_i }
  g[as[0]] = as.drop(2)
}

queue = []
mode = [] # nil: 未訪問, 1: 訪問済み, 2: 隣接も全て訪問済み
d = []

v = 1
mode[v] = 1
queue.push(v)
d[v] = 0
for i in 1..n
  while queue.size != 0
    v = queue.shift
    adjacencies = g[v]
    for a in adjacencies
      if mode[a] == nil
        # 初訪問
        mode[a] = 1
        queue.push(a)
        d[a] = d[v] + 1
      end
    end
  end
end

for i in 1..n
  puts "#{i} #{if d[i].nil? then -1 else d[i] end}"
end
