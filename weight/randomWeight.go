package weight

import (
	"athena-v4/common/utils"
	"encoding/json"
	"errors"
)

type Unit struct {
	From        int
	To          int
	Probability int
}

type Calculator struct {
	Units   []*Unit
	Weights []int
}

func NewUnits(str string) ([][]int, error) {
	u := [][]int{}
	err := json.Unmarshal([]byte(str), &u)
	return u, err
}

// 0~9 50%概率，10~19 30%概率，20~39 20%概率
// input [[0,9,50],[10,19,30],[20,39,20]]
func NewCalculator(units [][]int) (*Calculator, error) {
	sum := 0
	c := Calculator{}

	for _, u := range units {
		if len(u) != 3 {
			return nil, errors.New("args error 1")
		}
		if u[0] > u[1] {
			return nil, errors.New("args error 2")
		}
		sum += u[2]
		unit := Unit{From: u[0], To: u[1], Probability: u[2]}
		c.Units = append(c.Units, &unit)
		c.Weights = append(c.Weights, u[2])
	}

	if sum != 100 {
		return nil, errors.New("args error: total probability should equal 100")
	}

	return &c, nil
}

func (this Calculator) Calculate() (idx int, value int) {
	idx = SelectOne(this.Weights...)
	hit_unit := this.Units[idx]
	value = utils.RandomRange(hit_unit.From, hit_unit.To)
	return
}

func (this Calculator) GetUnit() *Unit {
	idx := SelectOne(this.Weights...)
	hit_unit := this.Units[idx]
	return hit_unit
}

func ValidateUnits(str string) error {
	units, err := NewUnits(str)
	if err != nil {
		return err
	}
	sum := 0
	for _, u := range units {
		if len(u) != 3 {
			return errors.New("'len(u) != 3`")
		}
		if u[0] > u[1] {
			return errors.New("'u[0] > u[1]'")
		}
		sum += u[2]
	}

	if sum != 100 {
		return errors.New("total probability should equal 100")
	}

	return nil
}
