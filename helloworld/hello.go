package main

import "fmt"

const( spanish = "Spanish"
 french = "French"

 englishHelloPrefix = "Hello, "
 frenchHelloPrefix = "Bonjour, "
 spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string{
	if name == ""{
		name = "World"
	}

	if language == "Spanish"{
		return  spanishHelloPrefix + name
	}
	
	return englishHelloPrefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language{
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main(){
	fmt.Println(Hello("Jhon", "English"))
}