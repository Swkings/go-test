package test

import (
	"fmt"
	"testing"
)

type Peo struct {
	Name  string
	Age   int
	State int
}

func (p *Peo) setName(name string) {
	p.Name = name
}
func (p *Peo) setAge(age int) {
	p.Age = age
}
func (p *Peo) setState(state int) {
	p.State = state
}

type S1 struct {
	f *Peo
}
type S2 struct {
	f *Peo
}
type S3 struct {
	f *Peo
}
type S4 struct {
	f *Peo
}

func (f *Peo) InitFlow() *S1 {
	return &S1{
		f: f,
	}
}

func (s *S1) SetName(name string) *S2 {
	s.f.setName(name)
	return &S2{
		f: s.f,
	}
}

func (s *S2) SetAge(age int) *S3 {
	s.f.setAge(age)
	return &S3{
		f: s.f,
	}
}

func (s *S3) SetState(state int) *S4 {
	s.f.setState(state)
	return &S4{
		f: s.f,
	}
}
func (s *S4) Value() *Peo {
	fmt.Printf("flow: %+v\n", s.f)
	return s.f
}

func TestExecStepByStep(t *testing.T) {
	f := &Peo{}
	f.InitFlow().SetName("zs").SetAge(10).SetState(1).Value().InitFlow().SetName("ls").SetAge(20).SetState(2).Value()
	f.InitFlow().SetName("").SetAge(1).SetState(1).Value()
}
