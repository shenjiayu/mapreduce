package src

import (
	"strconv"
	"strings"
)

type Mapper struct{}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (this *Mapper) Map(line string) (string, float64) {
	values := strings.Split(line, "\t")
	//film_id := values[1]
	rating, _ := strconv.ParseFloat(values[2], 64)
	//timestamp := values[3]
	return values[0], rating
}
