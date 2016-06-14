package main

import (
  "fmt"
  "math/rand"
  "time"
  "sync"
)

func main() {

  defer timeTrack(time.Now(),"Volados ")

  var wg sync.WaitGroup

  var a,s int

  for i := 0; i < 10000; i++ {

    wg.Add(1)
    go func ()  {
      s1 := rand.NewSource(time.Now().UnixNano())
      r1 := rand.New(s1)
      if r1.Intn(2)==1 {
        a++
      }else {
        s++
      }
      wg.Done()
    }()
  }

  fmt.Printf("%d Aguilas y %d soles \n",a,s)
  wg.Wait()

}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s\n", name, elapsed)
}
