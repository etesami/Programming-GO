package main

func main() {
	// Initialize the logger
	// If no file path is provided, the logger will only log to stdout
	logger := InitializeLogger("log.txt")

	// Log a message
	logger.Info().Msg("Hello World!")
	logger.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	logger.Debug().
		Str("Name", "Tom").
		Send()

	logger.Info().Msg("hello world")

	logger.Log().
		Str("foo", "bar").
		Msg("")
}
