package main

import "fmt"

type WeaponStrategy interface {
	UseWeapo() // 使用武器
}

type AK47 struct{}

func (a *AK47) UseWeapo() {
	fmt.Println("使用AK47战斗")
}

type Knife struct{}

func (a *Knife) UseWeapo() {
	fmt.Println("使用匕首战斗")
}

type Hero struct {
	w WeaponStrategy
}

func (h *Hero) SetStrategy(strategy WeaponStrategy) {
	h.w = strategy
}

func (h *Hero) Fight() {
	h.w.UseWeapo()
}

func main() {
	hero := &Hero{}
	hero.SetStrategy(&AK47{})
	hero.Fight()

	hero.SetStrategy(&Knife{})
	hero.Fight()

}
