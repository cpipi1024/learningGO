package main

const englishPreFix = "Hello, "
const spanishPreFix = "Hola, "
const franchPreFix = "Bonjour, "

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return getLanguage(language) + name
}
func getLanguage(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = spanishPreFix
	case "franch":
		prefix = franchPreFix
	default:
		prefix = englishPreFix
	}
	return prefix
}

func main() {
}
