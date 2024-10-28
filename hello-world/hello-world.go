package hello

const (
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	punjabiHelloPrefix = "Sat Sri Akal, "
	spanishHelloPrefix = "Hola, "

	french  = "French"
	punjabi = "Punjabi"
	spanish = "Spanish"
)

func HelloWorld(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case punjabi:
		prefix = punjabiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
