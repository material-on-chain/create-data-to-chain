package create_data_to_chain

import "encoding/json"

type FlaxFabricPrepreg struct {
	Product string `json:"product"`
	Company string `json:"company"`
	SalesManager string `json:"sales_manager"`
	Shape string `json:"shape"`
	ProcessType string `json:"process_type"`
	LayerConfiguration float32 `json:"layer_configuration"`
	DensityInGramPerCubicCentimeter float32 `json:"density_in_gram_per_cubic_centimeter"`
	LocationInBlade string `json:"location_in_blade"`
	ElasticModulusInGigaPascal float32 `json:"elastic_modulus_in_giga_pascal"`
	TensileStrengthInGigaPascal float32 `json:"tensile_strength_in_giga_pascal"`
	BendingStrengthInGigaPascal float32 `json:"bending_strength_in_giga_pascal"`
	PreTransaction []string `json:"pre_transaction"`
}

func NewFlaxFabricPrepreg() *FlaxFabricPrepreg {
	return &FlaxFabricPrepreg{
		Product: "flax fabric prepreg",
	}
}

func (f *FlaxFabricPrepreg) SetCompany (company string) { f.Company = company }
func (f *FlaxFabricPrepreg) SetSalesManager (manager string) { f.SalesManager = manager }
func (f *FlaxFabricPrepreg) SetShape (shape string) { f.Shape = shape }
func (f *FlaxFabricPrepreg) SetProcessType (pt string) { f.ProcessType = pt }
func (f *FlaxFabricPrepreg) SetLayerConfiguration (conf float32) { f.LayerConfiguration = conf }
func (f *FlaxFabricPrepreg) SetDensityInGramPerCubicCentimeter (density float32) { f.DensityInGramPerCubicCentimeter = density }
func (f *FlaxFabricPrepreg) SetLocationInBlade (location string) { f.LocationInBlade = location }
func (f *FlaxFabricPrepreg) SetElasticModulusInGigaPascal (em float32) { f.ElasticModulusInGigaPascal = em }
func (f *FlaxFabricPrepreg) SetTensileStrengthInGigaPascal (ts float32) { f.TensileStrengthInGigaPascal = ts }
func (f *FlaxFabricPrepreg) SetBendingStrengthInGigaPascal (bs float32) { f.BendingStrengthInGigaPascal = bs }
func (f *FlaxFabricPrepreg) SetPreTransaction(tx ...string) { f.PreTransaction = tx }


func (f FlaxFabricPrepreg) ToString () string {
	bytes, _ := json.Marshal(f)
	return string(bytes)
}