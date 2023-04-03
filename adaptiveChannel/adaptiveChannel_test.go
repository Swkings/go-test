package adaptiveChannel

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	ach := NewAdaptiveChannel[int](10)

	go func() {
		for i := 0; i <= 1000; i++ {
			ach.InChan <- i
		}
		close(ach.InChan)
	}()

	for v := range ach.OutChan {
		// time.Sleep(10 * time.Millisecond)
		fmt.Printf("recv element: %v, chan len: %v, buffer node num: %v, node element list len: %v\n", v, ach.GetAdaptiveChannelLen(), ach.GetBufferNodeNum(), ach.GetBufferNodeElementListLen())
	}
}
