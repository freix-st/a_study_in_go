package structs

import "time"

type StopWatch struct {
	start   time.Time
	total   time.Duration
	running bool
}

func (s *StopWatch) Start() {
	if !s.running {
		s.start = time.Now()
		s.running = true
	}
}
