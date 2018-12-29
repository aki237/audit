package adapter

// Plug is used to denote a set of adapters that do a particular type of collection
type Plug uint

const (
	// OS plug types only fetch OS based resources like RAM, CPU Usage,
	// process usage and et.,
	OS Plug = iota + 1

	// NETWORK plug types only fetch network based resources
	NETWORK

	// LOG is used to collect all the logs of the current process
	LOG
)
