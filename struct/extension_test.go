package _struct

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" " + host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Print("wang")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" " + host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Grace")
}