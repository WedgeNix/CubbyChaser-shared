package shared

import (
	"math"
	"strconv"
)

func (s Session) String() string {
	w := int(math.Sqrt(float64(len(s.Cubbies))))
	var cubs string
	for x, cub := range s.Cubbies {
		if len(cub.OrderNumber) < 1 {
			cubs += `[     ]`
			continue
		}
		itemlen := len(cub.Items)
		y := ""
		if itemlen < 10 {
			y = " "
		}
		itemcap := cap(cub.Items)
		z := ""
		if itemcap < 10 {
			z = " "
		}
		cubs += `[` + y + strconv.Itoa(itemlen) + `/` + z + strconv.Itoa(itemcap) + `]`
		if (x+1)%w == 0 {
			cubs += "\n"
		}
	}
	return cubs
}
