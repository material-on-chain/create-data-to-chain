package create_data_to_chain

import "encoding/json"

// parameters of flax fiber
type FlaxFiber struct {
	Product string `json:"product"`
	Company string `json:"company"`
	SalesManager string `json:"sales_manager"`
	ProducingArea string `json:"producing_area"`
	AverageLengthInMillimeter float32 `json:"average_length_in_millimeter"`
	DensityInGramPerCubicCentimeter float32 `json:"density_in_gram_per_cubic_centimeter"`
	ElasticModulusInGigaPascal float32 `json:"elastic_modulus_in_giga_pascal"`
	TensileStrengthInGigaPascal float32 `json:"tensile_strength_in_giga_pascal"`
	ElongationPercent float32 `json:"elongation_percent"`
	PreTransaction []string `json:"pre_transaction"`
}

func NewFlaxFiber () *FlaxFiber {
	return  &FlaxFiber{
		Product: "flax fiber",
	}
}

func (f *FlaxFiber) SetCompany(company string) { f.Company = company }
func (f *FlaxFiber) SetSalasManager (manager string) { f.SalesManager = manager }
func (f *FlaxFiber) SetProducingArea (area string) { f.ProducingArea = area }
func (f *FlaxFiber) SetAverageLengthInMillimeter (length float32) {f.AverageLengthInMillimeter = length }
func (f *FlaxFiber) SetDensityInGramPerCubicCentimeter (density float32) { f.DensityInGramPerCubicCentimeter = density }
func (f *FlaxFiber) SetElasticModulusInGigaPascal (em float32) { f.ElasticModulusInGigaPascal = em }
func (f *FlaxFiber) SetTensileStrengthInGigaPascal (ts float32) { f.TensileStrengthInGigaPascal = ts }
func (f *FlaxFiber) SetElongationPercent (percent float32) { f.ElongationPercent = percent }
func (f *FlaxFiber) SetPreTransaction(tx ...string) { f.PreTransaction = tx }

func (f FlaxFiber) ToString () string {
	bytes, _ := json.Marshal(f)
	return string(bytes)
}