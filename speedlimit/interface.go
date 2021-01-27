package speedlimit

// Limiter ...
type Limiter interface {
	CanAdd() bool
}
