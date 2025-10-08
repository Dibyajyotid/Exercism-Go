package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    Greet(name string) string
    LanguageName() string
}

func SayHello(name string, g Greeter) string {
    return g.Greet(name)
}

//For italian Language implementation
type Italian struct {}

//implementing the italian struct methods
//Greet method
func (Italian) Greet(name string) string {
    return "I can speak Italian: Ciao " + name+"!"
}

//LanguageName method
func (Italian) LanguageName() string {
    return "Italian"
}

//For portuguese language implementation
type Portuguese struct {}

//implementing the portuguese struct methods
// Greet method
func (Portuguese) Greet(name string) string {
    return "I can speak Portuguese: Olá " + name + "!"
}

//LanguageName method
func (Portuguese) LanguageName() string {
    return "Portuguese"
}