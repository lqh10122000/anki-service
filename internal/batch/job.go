package batch

import "time"

func Schedule(processor *Processor) {
	ticker := time.NewTicker(100 * time.Minute)
	go func() {
		for {
			<-ticker.C
			processor.RunBatch()
		}
	}()
}
