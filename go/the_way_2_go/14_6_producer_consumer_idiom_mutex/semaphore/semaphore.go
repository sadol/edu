package semaphore

type Semaphore chan interface{}

// `open' semaphore (acquire `n' resources)
func (sem Semaphore) P(n int) {
	fake := new(interface{})
	for i := 0; i < n; i++ {
		sem <- fake
	}
}

// `close' semaphore (release `n' resorces)
func (sem Semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-sem // signal (discard value from unbuffered channel)
	}
}

func NewSemaphore() Semaphore {
    newSemaphore := make(Semaphore)
    return newSemaphore
}

// mutex lock
func (s Semaphore) Lock() {
	s.P(1)
}

// mutex unlock
func (s Semaphore) Unlock() {
	s.V(1)
}

func (s Semaphore) Wait(n int) {
	s.P(n)
}

func (s Semaphore) Signal() {
	s.V(1)
}
