package wait

import (
	"testing"
	"time"
)

const duration = 20 * time.Second

func TestSleep(t *testing.T) {
	start := time.Now()
	Sleep(duration)
	if time.Since(start) < duration {
		t.Error()
	}
}
