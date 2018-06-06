package robotname

import (
	"fmt"
	"math/rand"
)

var issuedName = make(map[string]bool)

type Robot struct {
	name string
}

func (robot *Robot) Name() string {
	if len(robot.name) == 0 {
		for {
			var l1 = rand.Int() % 26
			var l2 = rand.Int() % 26
			var num = rand.Int() % 1000
			var name = fmt.Sprintf("%c%c%03d", 'A'+l1, 'A'+l2, num)
			if !issuedName[name] {
				issuedName[name] = true
				robot.name = name
				break
			}
		}
	}
	return robot.name
}

func (robot *Robot) Reset() {
	robot.name = ""
}
