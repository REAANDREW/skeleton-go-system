package main

type Application struct {
	config Configuration
}

func NewApplication() Application {
	return Application{
		config: DefaultConfiguration(),
	}
}

func (app Application) WithArgumentParser(parserFactory ArgumentParserFactory) Application {
	var parser = parserFactory.Build(app.config)
	var config = parser.Parse()
	return Application{
		config: config,
	}
}

func (app Application) WithLogger(loggerFactory LoggerFactory) Application {
	Log = loggerFactory.Build(app.config)
	return app
}

func (app Application) Run() {
	Log.Debug("starting",
		"Version", Version,
		"BuildTime", BuildTime,
		"Hash", CommitHash)
}
