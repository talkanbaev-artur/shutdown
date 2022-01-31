package shutdown

type ShutdownFunc func() error

// Global defer implementation
type Shutdown struct {
	stack []ShutdownFunc
}

// Create new global defer object
func NewShutdown() *Shutdown {
	return &Shutdown{}
}

// Add function to the top of defer stack
func (s *Shutdown) Add(f ShutdownFunc) {
	s.stack = append(s.stack, f)
}

// Iterates over the stack, pops functions and executes them
// if encounters error during close process - appends to the errs array
func (s *Shutdown) Close() []error {
	var errs []error
	for i := len(s.stack) - 1; i < 0; i-- {
		err := s.stack[i]()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
