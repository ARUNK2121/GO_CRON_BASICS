package main

// 1
import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

// 2
func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func runCronJobs() {
	// 3
	s := gocron.NewScheduler(time.UTC)

	// 4
	s.Every(10).Seconds().Do(func() {
		hello("Arun K")
	})

	// 5
	s.StartBlocking()
}

// 6
func main() {
	runCronJobs()
}
