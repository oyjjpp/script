package concurrency

import "testing"

func TestChannelConcurrency(t *testing.T) {
	channelConcurrency()
}

func TestSyncConcurrency(t *testing.T) {
	syncConcurrency()
}

func TestWithContextConcurrency(t *testing.T) {
	withContextConcurrency()
}
