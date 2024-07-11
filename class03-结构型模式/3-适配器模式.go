package main

import "fmt"

// 适配目标 抽象的技能
type Attack interface {
	Fight()
}

// 具体的技能
type Dabaojian struct {
}

func (d *Dabaojian) Fight() {
	fmt.Println("使用了大保健技能，将敌人击杀...")
}

type Hero struct {
	Name   string
	attack Attack
}

func (h *Hero) Skill() {
	fmt.Println(h.Name, "使用了技能:")
	h.attack.Fight()
}

// 适配者
type Poweroff struct {
}

func (p *Poweroff) Shutdown() {
	fmt.Println("电脑即将关机...")
}

// 适配器
type Adapter struct {
	Poweroff *Poweroff
}

func NewAdapter(p *Poweroff) *Adapter {
	return &Adapter{Poweroff: p}
}

func (a *Adapter) Fight() {
	a.Poweroff.Shutdown()
}

func main() {
	/**
	gailun := &Hero{"盖伦", new(Dabaojian)}
	gailun.Skill()
	*/
	gailun := &Hero{"盖伦", NewAdapter(new(Poweroff))}
	gailun.Skill()

}
