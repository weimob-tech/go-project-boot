package boot

type BlockingStarter interface {
	Start()
}

type lambdaStarter struct {
	fn func()
}

func (app lambdaStarter) Start() {
	app.fn()
}

func LambdaStarter(fn func()) BlockingStarter {
	return lambdaStarter{fn: fn}
}
