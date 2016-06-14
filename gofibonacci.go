package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  defer timeTrack(time.Now(),"2 fibonaccis")
  var wg sync.WaitGroup

  for i := 0; i < 100; i++ {
    wg.Add(1)
    go func(){
      fmt.Printf("fib %d %d \n",i,fib(42))
      wg.Done()
    }()
  }

  wg.Wait()


}


func fib(n int)int  {
  if n < 2{
    return 0
  }
  if n < 3{
    return 1
  }
  return fib(n-1)+fib(n-2)
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s", name, elapsed)
}
