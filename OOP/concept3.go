package main

import "fmt"

// Cara golang yang sebenarnya untuk memecahkan masalah di level 2
// Dengan menggunakan interfaces untuk mendefinisikan perilaku (behavior)
// Kita tidak peduli apa itu skill (fireball atau healingwave)
// Kita hanya peduli bahwa skill itu bisa digunakan (execute)

// Konsep: interfaces untuk polimorfisme (banyak bentuk, satu perilaku)

// Definisikan "kontrak" atau "perilaku"
// Skill apapun harus bisa dieksekusi dan memberi tahu kita nama dan mana costnya

// Dengan menggunakan interface IsActivable, Hero kita sekarang bisa memegang perilaku apa saja
// Dapat membuat StunSkill, BuffSkill, SummonSkill
// Tanpa harus mengubah kode di dalam struct Hero

type IsActivable interface {
	Execute(caster *HeroConcept3, target *HeroConcept3)
	GetName() string
	GetManaCost() int
}

type HeroConcept3 struct {
	Name   string
	HP     int
	Mana   int
	Skills []IsActivable
}

type DamageSkill struct {
	Name     string
	ManaCost int
	Damage   int
}

func (s DamageSkill) Execute(caster *HeroConcept3, target *HeroConcept3) {
	caster.Mana = caster.Mana - s.ManaCost
	target.HP = target.HP - s.Damage
	fmt.Printf("%s menyerang %s dengan %s, %d damage!\n", caster.Name, target.Name, s.Name, s.Damage)
}

func (s DamageSkill) GetName() string {
	return s.Name
}

func (s DamageSkill) GetManaCost() int {
	return s.ManaCost
}

type HealSkill struct {
	Name       string
	ManaCost   int
	HealAmount int
}

func (s HealSkill) Execute(caster *HeroConcept3, target *HeroConcept3) {
	caster.Mana = caster.Mana - s.ManaCost
	target.HP = target.HP + s.HealAmount
	fmt.Printf("%s menyembuhkan %s dengan %s, +%d HP!\n", caster.Name, target.Name, s.Name, s.HealAmount)
}

func (s HealSkill) GetName() string {
	return s.Name
}

func (s HealSkill) GetManaCost() int {
	return s.ManaCost
}

func (h *HeroConcept3) LearnSkill(skill IsActivable) {
	h.Skills = append(h.Skills, skill)
	fmt.Printf("%s mempelajari %s\n", h.Name, skill.GetName())
}

func (h *HeroConcept3) UseSkill(skillName string, target *HeroConcept3) {
	var foundSkill IsActivable
	for _, s := range h.Skills {
		if s.GetName() == skillName {
			foundSkill = s
			break
		}
	}

	if foundSkill == nil {
		fmt.Printf("%s tidak punya skill %s\n", h.Name, skillName)
		return
	}

	if h.Mana < foundSkill.GetManaCost() {
		fmt.Printf("Mana tidak cukup untuk %s\n", skillName)
		return
	}

	foundSkill.Execute(h, target)
	fmt.Printf("Status: %s (HP: %d, Mana: %d) | %s (HP: %d, Mana: %d)\n",
		h.Name, h.HP, h.Mana, target.Name, target.HP, target.Mana)
}

func main() {
	dazzle := &HeroConcept3{
		Name: "Dazzle",
		HP:   1500,
		Mana: 500,
	}

	axe := &HeroConcept3{
		Name: "Axe",
		HP:   800,
		Mana: 200,
	}

	poisonTouch := DamageSkill{
		Name:     "Poison Touch",
		ManaCost: 100,
		Damage:   50,
	}

	shallowGrave := HealSkill{
		Name:       "Shallow Grave",
		ManaCost:   120,
		HealAmount: 300,
	}

	dazzle.LearnSkill(poisonTouch)
	dazzle.LearnSkill(shallowGrave)

	fmt.Println("---- Pertarungan Dimulai -----")

	dazzle.UseSkill("Poison Touch", axe)
	dazzle.UseSkill("Shallow Grave", axe)
}
