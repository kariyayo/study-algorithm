class Queue
  def initialize
    @arr = []
    @head = 0
    @tail = 0
  end

  def enque(x)
    @arr[@tail] = x
    @tail = @tail + 1
  end

  def deque
    x = @arr[@head]
    @head = @head + 1
    x
  end

  def size
    @tail - @head
  end
end

def roundrobin(q, queue)
  finished = []
  t = 0
  while queue.size != 0
    process = queue.deque
    if (process[1].to_i <= q)
      t = t + process[1].to_i
      process[1] = 0
      finished.push("#{process[0]} #{t}")
    else
      t = t + q
      process[1] = process[1].to_i - q
      queue.enque(process)
    end
  end

  puts finished.join("\n")
end



as = STDIN.gets.split(' ')

n = as[0].to_i
q = as[1].to_i

queue = Queue.new
n.times.each {
  queue.enque(STDIN.gets.split(' '))
}

roundrobin(q, queue)

