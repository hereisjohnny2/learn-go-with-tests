package main

func Hello(name string, language string) string {
  if name == "" {
    return "Hello"
  }

  switch language {
  case "es":
    return "Hola, " + name
  case "fr":
    return "Bonjour, " + name
  default:
    return "Hello, " + name
  }
}

func main() {
}
