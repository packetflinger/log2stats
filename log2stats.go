package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

/**
 * Possible means of death
 */
const (
	MODBlaster = iota
	MODShotgun
	MODSuperShotgun
	MODMachinegun
	MODChaingun
	MODHandGrenade
	MODGrenadeLauncher
	MODHyperBlaster
	MODRocketLauncher
	MODRailgun
	MODBigFuckingGun
	MODFall
	MODDrown
	MODSquish
)

type Frag struct {
	Attacker     string
	Victim       string
	MeansOfDeath int
	Suicide      bool
}

/**
 * Regex patterns for matching frags
 */
type FragPatterns struct {
	BFG1         *regexp.Regexp
	BFG2         *regexp.Regexp
	BFG3         *regexp.Regexp
	Rail         *regexp.Regexp
	Rocket1      *regexp.Regexp
	Rocket2      *regexp.Regexp
	Hyper        *regexp.Regexp
	GLauncher    *regexp.Regexp
	Grenade1     *regexp.Regexp
	Grenade2     *regexp.Regexp
	Chaingun     *regexp.Regexp
	Machinegun   *regexp.Regexp
	SuperShotgun *regexp.Regexp
	Shotgun      *regexp.Regexp
	Blaster      *regexp.Regexp
}

type Stats struct {
	TotalFrags       int
	TotalSuicides    int
	PlayerTotalFrags map[string]int
}

var patterns FragPatterns

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//stats := Stats{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ParseLogLine(scanner.Text())
	}
}

func init() {
	patterns.BFG1, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) couldn't hide from (?P<Victim>.+)'s BFGn`)
	patterns.BFG2, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) was disintegrated by (?P<Victim>.+)'s BFG blast`)
	patterns.BFG3, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) saw the pretty lights from (?P<Victim>.+)'s BFG`)
	patterns.Rail, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) was railed by (?P<Victim>.+)`)
	patterns.Rocket1, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) ate (?P<Victim>.+)'s rocket`)
	patterns.Rocket2, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) almost dodged (?P<Victim>.+)'s rocket`)
	patterns.Hyper, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) was melted by (?P<Victim>.+)'s hyperblaster`)
	patterns.GLauncher, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) was popped by (?P<Victim>.+)'s grenade`)
	patterns.Grenade1, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) caught (?P<Victim>.+)'s handgrenade`)
	patterns.Grenade2, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) didn't see (?P<Victim>.+)'s handgrenade`)
	patterns.Chaingun, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) was cut in half by (?P<Victim>.+)'s chaingun`)
	patterns.Machinegun, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) was machinegunned by (?P<Victim>.+)`)
	patterns.SuperShotgun, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) was blown away by (?P<Victim>.+)'s super shotgun`)
	patterns.Shotgun, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) was gunned down by (?P<Victim>.+)`)
	patterns.Blaster, _ = regexp.Compile(`\[.+\] (?P<Attacker>.+) was blasted by (?P<Victim>.+)`)
}

func ParseLogLine(line string) {
	fmt.Println(line)
	//match := patterns.Rail.FindStringSubmatch(line)
	//if len(match) > 0 {
	//fmt.Println(match[1])
	//}
}
