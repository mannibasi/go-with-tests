package hello

func HelloWorld(name string) string {
	if name == "" {
		name = "world"
	}
	return "Hello, " + name
}
