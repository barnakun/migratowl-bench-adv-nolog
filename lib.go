package lib

// Greet returns a greeting for name in the given language ("en" or "hu").
func Greet(name, lang string) string {
	if lang == "hu" {
		return "szia, " + name
	}
	return "hello, " + name
}

// Version returns the library version.
func Version() string { return "1.1.0" }

// Farewell says goodbye.
func Farewell(name string) string { return "bye, " + name }
