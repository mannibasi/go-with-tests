package hello

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const frenchHelloPrefix = "Bonjour, "
const french = "French"
const punjabiHelloPrefix = "Sat Sri Akal, "
const punjabi = "Punjabi"

func HelloWorld(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case punjabi:
		prefix = punjabiHelloPrefix
	}

	return prefix + name
}
