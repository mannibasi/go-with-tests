package hello

const englishHelloPrefix = "Hello, "

func HelloWorld(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}
