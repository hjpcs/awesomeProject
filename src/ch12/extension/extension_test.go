package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("say")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("say", host)
}

type Dog struct {
	//p *Pet
	Pet
}

func (d *Dog) Speak() {
	//d.p.Speak()
	fmt.Println("dog say wang!")
}

//func (d *Dog) SpeakTo(host string) {
//	//d.p.SpeakTo(host)
//	d.Speak()
//	fmt.Println("/", host)
//}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("hohoho")
}
