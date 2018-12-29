package adapter

// Adapter is the interface for all the collectors to fetch resources
type Adapter interface {
	Collect() error
	Plug() Plug
	Commit()
}
