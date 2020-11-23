def to_key(s)
  s.gsub!('A', '1')
  s.gsub!('C', '2')
  s.gsub!('G', '3')
  s.gsub!('T', '4')
  s.to_i
end
def hash(key, i)
  (h1(key) + i * h2(key)) % 1046527
end
def h1(key)
  key % 1046527
end
def h2(key)
  1 + (key % (1046527 - 1))
end

class Dict
  def initialize
    @arr = []
  end
  def insert(s)
    key = to_key(s)
    i = 0
    loop do
      h = hash(key, i)
      if @arr[h] == s
        return
      elsif @arr[h].nil?
        @arr[h] = s
        return
      end
      i = i + 1
    end
  end
  def find(s)
    key = to_key(s)
    i = 0
    loop do
      h = hash(key, i)
      if @arr[h] == s
        return true
      elsif @arr[h].nil?
        return false
      end
      i = i + 1
    end
    false
  end
end

n = STDIN.gets.to_i

dict = Dict.new
n.times.each {
  ss = STDIN.gets.split(' ')
  if ss[0] == 'insert' then
    dict.insert(ss[1])
  else
    if dict.find(ss[1]) then
      puts 'yes'
    else
      puts 'no'
    end
  end
}

