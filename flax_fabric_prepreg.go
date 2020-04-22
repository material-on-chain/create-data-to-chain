package create_data_to_chain

import (
	"encoding/hex"
	"encoding/json"
	"github.com/material-on-chain/go-owcrypt"
)

type FlaxFabricPrepreg struct {
	Product string `json:"product"`
	Company string `json:"company"`
	SalesManager string `json:"sales_manager"`
	Shape string `json:"shape"`
	ProcessType string `json:"process_type"`
	LayerConfiguration float32 `json:"layer_configuration"`
	DensityInGramPerCubicCentimeter float32 `json:"density"`
	LocationInBlade string `json:"location_in_blade"`
	ElasticModulusInGigaPascal float32 `json:"elastic_modulus"`
	TensileStrengthInGigaPascal float32 `json:"tensile_strength"`
	BendingStrengthInGigaPascal float32 `json:"bending_strength"`
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

func (f FlaxFabricPrepreg) GetDataToChain() string {
	fingerPrint := owcrypt.Hash([]byte(f.ToString()),0, owcrypt.HASH_ALG_SHA3_256_RIPEMD160)
	var (
		body = make(map[string]interface{}, 0)
	)

	body["company"] = f.Company
	body["finger_print"] = hex.EncodeToString(fingerPrint)

	json, _ := json.Marshal(body)

	return string(json)
}