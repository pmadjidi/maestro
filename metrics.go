package main


type metrics struct {
	timeouts int64
	errors int64
	invalidCalls int64
	success int64
}

func newMetrics () *metrics {
	return &metrics{}
}
