package abr

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Personne struct {
	Name       string
	waitGroup  *sync.WaitGroup
	time_begin time.Time
	Glasses    bool
	Belt       bool
}

func NewPersonne(name string, wg *sync.WaitGroup) *Personne {
	return &Personne{Name: name, waitGroup: wg}
}

func (p *Personne) Start(c chan (string)) {
	go func() {
		fmt.Printf("%s a commencé à se préparer \n", p.Name)
		p.time_begin = time.Now()
		p.SePrepare(c)
		fmt.Printf("%s a passé %d secondes à se préparer \n", p.Name, time.Since(p.time_begin)*time.Second)
	}()
}

func (p *Personne) Partir() {
	go func() {
		fmt.Printf("%s a commencé à mettre ses chaussures\n", p.Name)
		p.PutShoes()
		p.waitGroup.Done()
	}()
}

func (p *Personne) SePrepare(c chan (string)) {
	p.WearGlasses()
	p.WearBelt()
	p.CloseWindow(c)
	p.CloseFan(c)
	p.waitGroup.Done()

}

func (p *Personne) WearGlasses() {
	time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second) // n := a + rand.Intn(b-a+1) => [a:b]
	p.Glasses = true
}

func (p *Personne) WearBelt() {
	time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
	p.Belt = true
}

func (p *Personne) CloseWindow(c chan (string)) {
	c <- "window"
	<-c
}

func (p *Personne) CloseFan(c chan (string)) {
	c <- "fan"
	<-c
}

func (p *Personne) PutShoes() {
	p.time_begin = time.Now()
	time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
	fmt.Printf("%s a mis %d secondes à mettre ses chaussures \n", p.Name, time.Since(p.time_begin))
}
