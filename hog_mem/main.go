package main

import (
  "context"
  "fmt"
  "os/signal"
  "runtime"
  "sync"
  "syscall"
  "time"

  "github.com/pbnjay/memory"
  "bufio"
  "os"
)

func main() {
  fmt.Println("Hello world!")
  fmt.Print("Press 'Enter' to continue...")
  bufio.NewReader(os.Stdin).ReadBytes('\n')

  ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
  defer cancel()

  var wg sync.WaitGroup

  // Allocate all available CPUs
  numCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(numCPU)

  // Start a busy goroutine for each CPU
  for i := 0; i < numCPU; i++ {
    go func() {
      wg.Add(1)
      defer wg.Done()

      fmt.Println("Started a CPU hog")
      for ctx.Err() == nil { /* Busy loop to keep CPU busy */ }
    }()
  }

  // Start allocating memory every second
  const memChunkMB = 10
  var chunks [][]byte
  go func() {
    wg.Add(1)
    defer wg.Done()

    for ctx.Err() == nil {
      chunk := make([]byte, memChunkMB*1024*1024)

      // Prevent the memory from being optimized away
      for i := 0; i < len(chunk); i++ {
        chunk[i] = byte(i % 256)
      }
      chunks = append(chunks, chunk)

      //fmt.Printf("Total system memory: %d\n", memory.TotalMemory())
      fmt.Printf("Allocated %d MB of %d MB memory\n", memChunkMB, memory.TotalMemory()/1024/1024)
      time.Sleep(1 * time.Second)
    }
  }()

  <-ctx.Done()
  fmt.Println("Received termination signal. Initiating shutdown...")

  wg.Wait()
  fmt.Println("Shutdown complete.")
}
