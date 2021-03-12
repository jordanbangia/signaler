# signaler

Tell your waiting goroutines that something happened in a threadsafe way.

```go

ctx := context.Background()

s := signaler.New()

go func() {
    select {
        case <-ctx.Done():
            return
        case <-s.Subscribe():
            fmt.Println("I got the signal!")
    }
}()

time.Sleep(10 * time.Second)
s.Trigger()
```
