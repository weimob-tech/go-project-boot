package boot

func Starter(opts ...Option) BlockingStarter {
	conf := configureFrom(opts...)

	// before setup
	conf.BeforeSetup()
	// do setup
	if conf.SetupConfig != nil {
		// setup before everything
		conf.SetupConfig()
	}
	if conf.SetupLog != nil {
		// setup after configs init.
		conf.SetupLog()
	}
	if conf.SetupContainer != nil {
		// setup weimob container
		conf.SetupContainer()
	}
	// after setup
	conf.AfterSetup()
	// setup codec
	if conf.NewCodec != nil {
		conf.NewCodec()
	}
	// setup client
	if conf.NewClient != nil {
		conf.NewClient()
	}
	// setup server
	if conf.NewServer != nil {
		conf.NewServer()
	}
	if conf.SetupServer != nil {
		conf.SetupServer()
	}
	// starter
	starters := conf.BlockingStarters
	if conf.PreStarterHook != nil {
		starters = append([]BlockingStarter{LambdaStarter(conf.PreStarterHook)}, starters...)
	}
	if conf.Server != nil {
		starters = append(starters, conf.Server)
	}
	if conf.PostStarterHook != nil {
		starters = append(starters, LambdaStarter(conf.PostStarterHook))
	}
	if len(starters) == 0 {
		return &nopStarter{}
	}
	return &compositeStarter{starters: starters}
}
