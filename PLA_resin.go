package create_data_to_chain

import "encoding/json"

type PLAResin struct {
	Product string `json:"product"`
	Company string `json:"company"`
	SalesManager string `json:"sales_manager"`
	ProducingArea string `json:"producing_area"`
	GlassTransitionTemperatureInCentigrade float32 `json:"glass_transition_temperature_in_centigrade"`
	DensityInGramPerCubicCentimeter float32 `json:"density_in_gram_per_cubic_centimeter"`
	ElasticModulusInGigaPascal float32 `json:"elastic_modulus_in_giga_pascal"`
	TensileStrengthInGigaPascal float32 `json:"tensile_strength_in_giga_pascal"`
	ElongationPercent float32 `json:"elongation_percent"`
}

func NewPLAResin() *PLAResin {
	return &PLAResin{
		Product: "PLA resin",
	}
}

func (pr *PLAResin) SetCompany (company string) { pr.Company = company }
func (pr *PLAResin) SetSalesManager (manager string) { pr.SalesManager = manager }
func (pr *PLAResin) SetProducingArea (area string) { pr.ProducingArea = area }
func (pr *PLAResin) SetGlassTransitionTemperatureInCentigrade (temperature float32) { pr.GlassTransitionTemperatureInCentigrade = temperature }
func (pr *PLAResin) SetDensityInGramPerCubicCentimeter (density float32) { pr.DensityInGramPerCubicCentimeter = density }
func (pr *PLAResin) SetElasticModulusInGigaPascal (em float32) { pr.ElasticModulusInGigaPascal = em }
func (pr *PLAResin) SetTensileStrengthInGigaPascal (ts float32) { pr.TensileStrengthInGigaPascal = ts }
func (pr *PLAResin) SetElongationPercent (percent float32) { pr.ElongationPercent = percent }

func (pr PLAResin) ToString () string {
	bytes, _ := json.Marshal(pr)
	return string(bytes)
}