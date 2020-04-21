package create_data_to_chain

import (
	"fmt"
	"testing"
)

func TestFlaxFiber_ToString(t *testing.T) {
	ff := NewFlaxFiber()

	ff.SetCompany("chunlong")
	ff.SetSalasManager("xxx")
	ff.SetProducingArea("Jiangxi,China")
	ff.SetAverageLengthInMillimeter(18)
	ff.SetDensityInGramPerCubicCentimeter(1.5)
	ff.SetElasticModulusInGigaPascal(6.5)
	ff.SetTensileStrengthInGigaPascal(1.2)
	ff.SetElongationPercent(1.4)

	fmt.Println(ff.ToString())
}
