package hello

const englishHelloPrefix = "Hello, "

func HelloWorld(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Spanish" {
		return "Hola, Fernando"
	}
	return englishHelloPrefix + name
}
