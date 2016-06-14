def sum(n)
  if n<2
    return 1
  end
  return n+sum(n-1)

end

def fib(n)
  if n<2
    return 0
  end
  if n<3
    return 1
  end

  return fib(n-2)+fib(n-1)
end

def main
  start = Time.now
  puts fib(42)
  finish = Time.now
  puts "Time = #{finish-start}"
end

main
