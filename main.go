package main

var (
	Name       string
	Version    string
	BuildTime  string
	CommitHash string
	Log        Logger
)

func main() {
	NewApplication().
		WithArgumentParser(NewKingpinArgumentParser()).
		WithLogger(NewZapLogger()).
		Run()
}
