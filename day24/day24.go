package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type groupType struct {
	groupID     int
	groupType   string
	units       int
	hitPoints   int
	immuneTo    string
	weakTo      string
	damageLevel int
	weapon      string
	initiative  int
	boost       int
}

type combatType struct {
	attacker *groupType
	defender *groupType
}

type attackerSort []*groupType

func (s attackerSort) Len() int {
	return len(s)
}
func (s attackerSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s attackerSort) Less(i, j int) bool {
	powerI := (s[i].damageLevel + s[i].boost) * s[i].units
	powerJ := (s[j].damageLevel + s[j].boost) * s[j].units
	return powerI > powerJ || powerI == powerJ && s[i].initiative > s[j].initiative
}

type defenderSort struct {
	attacker  *groupType
	defenders *[]*groupType
}

func (s defenderSort) Len() int {
	return len(*(s.defenders))
}

func (s defenderSort) Swap(i, j int) {
	(*s.defenders)[i], (*s.defenders)[j] = (*s.defenders)[j], (*s.defenders)[i]
}

func attackPower(attacer *groupType, defender *groupType) (attackPower int) {
	weaponFactor := 1
	if strings.Contains(defender.weakTo, attacer.weapon) {
		weaponFactor = 2
	}
	if strings.Contains(defender.immuneTo, attacer.weapon) {
		weaponFactor = 0
	}
	attackPower = attacer.units * (attacer.damageLevel + attacer.boost) * weaponFactor
	return
}

func (s defenderSort) Less(i, j int) bool {

	powerI := attackPower(s.attacker, (*s.defenders)[i])
	powerJ := attackPower(s.attacker, (*s.defenders)[j])
	if powerI > powerJ {
		return true
	}
	if powerI == powerJ {
		defenderPowerI := (*s.defenders)[i].units * ((*s.defenders)[i].damageLevel + (*s.defenders)[i].boost)
		defenderPowerJ := (*s.defenders)[j].units * ((*s.defenders)[j].damageLevel + (*s.defenders)[j].boost)
		if defenderPowerI > defenderPowerJ {
			return true
		}
		if defenderPowerI == defenderPowerJ {
			return (*s.defenders)[i].initiative > (*s.defenders)[j].initiative
		}
	}
	return false
}

type combatSort []combatType

func (s combatSort) Len() int {
	return len(s)
}

func (s combatSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s combatSort) Less(i, j int) bool {
	return s[i].attacker.initiative > s[j].attacker.initiative
}

// File loading generates array of steps
func loadFile(name string) (immune []groupType, infection []groupType) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	immune = make([]groupType, 0)
	infection = make([]groupType, 0)
	var army *[]groupType
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("Immune System:|Infection:|(\\d+) units each with (\\d+) hit points( \\((weak|immune) to (\\w+[, \\w+]*)(; (weak|immune) to (\\w+[, \\w+]*))?\\))? with an attack that does (\\d+) (\\w+) damage at initiative (\\d+)")
	for scanner.Scan() {

		if i := scanner.Text(); len(i) > 0 {
			command := re.FindStringSubmatch(i)
			if command[0] == "Immune System:" {
				army = &immune
				continue
			}
			if command[0] == "Infection:" {
				army = &infection
				continue
			}
			group := groupType{}
			group.groupID = len(*army) + 1
			if army == &immune {
				group.groupType = "Immune System"
			} else {
				group.groupType = "Infection"
			}
			group.units, _ = strconv.Atoi(command[1])
			group.hitPoints, _ = strconv.Atoi(command[2])
			if command[4] == "weak" {
				group.weakTo = command[5]
			}
			if command[7] == "weak" {
				group.weakTo = command[8]
			}
			if command[4] == "immune" {
				group.immuneTo = command[5]
			}
			if command[7] == "immune" {
				group.immuneTo = command[8]
			}
			group.damageLevel, _ = strconv.Atoi(command[9])
			group.weapon = command[10]
			group.initiative, _ = strconv.Atoi(command[11])
			*army = append(*army, group)
		}
	}
	return immune, infection
}

func selectTarget(attacers, defeners *[]groupType) []combatType {
	combat := make([]combatType, 0)
	fighters := make([]*groupType, 0)
	for i := range *attacers {
		if (*attacers)[i].units > 0 {
			fighters = append(fighters, &(*attacers)[i])
		}
	}
	sort.Sort(attackerSort(fighters))
	targets := make([]*groupType, 0)
	for i := range *defeners {
		if (*defeners)[i].units > 0 {
			targets = append(targets, &(*defeners)[i])
		}
	}
	for i := range fighters {
		if len(targets) == 0 {
			break
		}
		sort.Sort(defenderSort{fighters[i], &targets})
		if attackPower(fighters[i], targets[0]) > 0 {
			combat = append(combat, combatType{fighters[i], targets[0]})
			targets = targets[1:]
		}
	}
	return combat
}

func displayCombat(combat []combatType) {
	for _, v := range combat {
		fmt.Println(v.attacker.groupType,
			"Group", v.attacker.groupID,
			"would deal defending group",
			v.defender.groupID,
			attackPower(v.attacker, v.defender), "damage")

	}
}

func attack(combat []combatType) {
	sort.Sort(combatSort(combat))
	for i := range combat {
		damege := attackPower(combat[i].attacker, combat[i].defender)
		unitsToKill := damege / combat[i].defender.hitPoints
		if combat[i].defender.units > unitsToKill {
			combat[i].defender.units -= unitsToKill
		} else {
			unitsToKill = combat[i].defender.units
			combat[i].defender.units = 0
		}
		if unitsToKill == 0 {
			continue
		}
		/*		fmt.Println(combat[i].attacker.groupType, "group", combat[i].attacker.groupID,
				"attacks defending group", combat[i].defender.groupID,
				", killing", unitsToKill, "units")
		*/
	}
	//fmt.Println()
}

func part1(immune, infection *[]groupType) (sumImmune, sumInfection int) {
	for i := 1; ; i++ {
		fmt.Println("step", i)
		combat := selectTarget(immune, infection)
		combat = append(combat, selectTarget(infection, immune)...)
		displayCombat(combat)
		if len(combat) == 0 {
			break
		}
		attack(combat)
	}
	sumImmune, sumInfection = unitSum(immune, infection)
	return
}

func unitSum(immune, infection *[]groupType) (sumImmune, sumInfection int) {
	sumImmune = 0
	for _, v := range *immune {
		sumImmune += v.units
	}
	sumInfection = 0
	for _, v := range *infection {
		sumInfection += v.units
	}
	return
}

func part2(file string) (sumImmune, sumInfection, boost int) {
	for boost = 0; ; boost++ {
		immune, infection := loadFile(file)
		for i := range immune {
			immune[i].boost = boost
		}
		for i := 1; ; i++ {
			combat := selectTarget(&immune, &infection)
			combat = append(combat, selectTarget(&infection, &immune)...)
			if len(combat) == 0 {
				break
			}
			sumImmunePre, sumInfectionPre := unitSum(&immune, &infection)
			attack(combat)
			sumImmune, sumInfection = unitSum(&immune, &infection)
			if sumImmunePre == sumImmune && sumInfectionPre == sumInfection {
				break
			}
		}
		if sumInfection == 0 && sumImmune > 0 {
			return
		}
	}
}

func main() {
	immune, infection := loadFile("input.txt")
	fmt.Println(immune)
	fmt.Println(infection)
	sumImmune, sumInfection := part1(&immune, &infection)
	fmt.Println("Immune", sumImmune, "Infection", sumInfection)
	boost := 0
	sumImmune, sumInfection, boost = part2("input.txt")
	fmt.Println("Boost", boost)
}
