package bruteforce

import "sync"

type BruteForce[T any] struct {
	// Threads
	threads int
	// check if current is in requirement
	Check func(T) bool
	// spawn next one, return true if no next val
	Spawn func() (T, bool)
}

// New a Bruteforce struct,
func New[T any](spawn func() (T, bool), check func(T) bool, threads int) *BruteForce[T] {
	return &BruteForce[T]{
		threads: threads,
		Check:   check,
		Spawn:   spawn,
	}
}

func (b *BruteForce[T]) Brute() (T, bool) {
	var wg sync.WaitGroup
	tasks := make(chan T, 128)
	output := make(chan T)
	// spawn workers
	for i := 0; i < b.threads; i++ {
		go func() {
			for v := range tasks {
				if b.Check(v) {
					output <- v
					wg.Done()
					return
				}
				wg.Done()
			}
		}()
	}

	for {
		select {
		case res := <-output:
			close(tasks)
			return res, true
		default:
			try, end := b.Spawn()
			if end {
				wg.Wait()
				res, ok := <-output
				return res, ok
			}
			wg.Add(1)
			tasks <- try
		}
	}
}
