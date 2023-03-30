package boot

import (
	"sync"
	"time"
)

type compositeStarter struct {
	wg       sync.WaitGroup
	starters []BlockingStarter
}

func (app *compositeStarter) Start() {
	app.wg.Add(len(app.starters))
	for _, s := range app.starters {
		go func(start BlockingStarter) {
			defer app.wg.Done()
			start.Start()
		}(s)
		// ensure the order
		time.Sleep(10 * time.Millisecond)
	}
	app.wg.Wait()
}

type nopStarter struct{}

func (app nopStarter) Start() {}
