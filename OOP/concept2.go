package main

import "fmt"

// Kekurangan level 2:
// Kaku, struct skill hanya bisa melakukan satu hal: memberikan damage
// Bagaimana jika ada skill yang heal, memberi stun, atau memberi buff
// Hal yang sama berlaku pada item. Bagaimana jika ada item yang passive?

type SkillConcept2 struct {
	Nama     string
	ManaCost int
	Damage   int
}

type ItemConcept2 struct {
	Nama string
	Cost int
}

// Hero sekarang memiliki skill dan item
type HeroConcept2 struct {
	Nama   string
	HP     int
	Mana   int
	Skills []SkillConcept2
	Items  []ItemConcept2
}

// Method untuk menambah skill/mempelajari skill
func (h *HeroConcept2) LearnSkill(s SkillConcept2) {
	h.Skills = append(h.Skills, s)
	fmt.Printf("%s mempelajari skill: %s\n", h.Nama, s.Nama)
}

// Method untuk membeli item
func (h *HeroConcept2) AddItem(i ItemConcept2) {
	h.Items = append(h.Items, i)
	fmt.Printf("%s membeli item: %s\n", h.Nama, i.Nama)
}

// Method untuk menggunakan skill
func (h *HeroConcept2) UseSkill(skillname string, target *HeroConcept2) {
	// Mencari skill di dalam daftar skill hero
	var foundSkill *SkillConcept2

	for i, s := range h.Skills {
		if s.Nama == skillname {
			foundSkill = &h.Skills[i]
			break
		}
	}

	// Cek skill
	if foundSkill == nil {
		fmt.Printf("%s tidak punya skill %s!\n", h.Nama, skillname)
		return
	}

	// Cek mana
	if h.Mana < foundSkill.ManaCost {
		fmt.Printf("%s tidak punya cukup mana untuk %s!\n", h.Nama, foundSkill.Nama)
		return
	}

	// Gunakan skill
	h.Mana = h.Mana - foundSkill.ManaCost
	target.HP = target.HP - foundSkill.Damage
	fmt.Printf("%s menggunakan %s ke %s, memberikan %d damage!\n", h.Nama, foundSkill.Nama, target.Nama, foundSkill.Damage)
	fmt.Printf("Sisa mana: %s: %d | Sisa HP %s: %d\n", h.Nama, h.Mana, target.Nama, target.HP)
}

func main() {
	lina := &HeroConcept2{
		Nama: "Lina",
		HP:   1200,
		Mana: 600,
	}

	axe := &HeroConcept2{
		Nama: "Axe",
		HP:   2000,
		Mana: 500,
	}

	dragonSlave := SkillConcept2{
		Nama:     "Dragon Slave",
		ManaCost: 200,
		Damage:   300,
	}

	lightStrike := SkillConcept2{
		Nama:     "Light Strike",
		ManaCost: 500,
		Damage:   600,
	}

	lina.LearnSkill(dragonSlave)
	lina.LearnSkill(lightStrike)

	lina.UseSkill("Dragon Slave", axe)
	lina.UseSkill("Light Strike", axe)

}
