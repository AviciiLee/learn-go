package _interface

import (
	"fmt"
	"testing"
)

type Programer interface {
	WriteHelloWorld() string
}

type GoProgramer struct {
	language string
}

type FakeProgramer struct {
	language string
}

func (p *GoProgramer) WriteHelloWorld() string {
	return "go hello world"
}

func (p *FakeProgramer) WriteHelloWorld() string {
	return "fake hello world"
}

func introduce(p Programer) {
	fmt.Println("programer hello world: ", p.WriteHelloWorld())
}

func TestInterface(t *testing.T) {
	var goP = &GoProgramer{}
	var fakeP = new(FakeProgramer)
	introduce(goP)
	introduce(fakeP)
}
