package create_data_to_chain

import (
	"encoding/hex"
	"encoding/json"
	owcrypt "github.com/material-on-chain/go-owcrypt"
)

// parameters of flax fiber
type FlaxFiber struct {
	Product string `json:"product"`
	Company string `json:"company"`
	SalesManager string `json:"sales_manager"`
	ProducingArea string `json:"producing_area"`
	AverageLengthInMillimeter float32 `json:"average_length"`
	DensityInGramPerCubicCentimeter float32 `json:"density"`
	ElasticModulusInGigaPascal float32 `json:"elastic_modulus"`
	TensileStrengthInGigaPascal float32 `json:"tensile_strength"`
	ElongationPercent float32 `json:"elongation_percent"`
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

func (f FlaxFiber) ToString () string {
	bytes, _ := json.Marshal(f)
	return string(bytes)
}

func (f FlaxFiber) GetDataToChain() string {
	fingerPrint := owcrypt.Hash([]byte(f.ToString()),0, owcrypt.HASH_ALG_SHA3_256_RIPEMD160)
	var (
		body = make(map[string]interface{}, 0)
	)

	body["company"] = f.Company
	body["finger_print"] = hex.EncodeToString(fingerPrint)

	json, _ := json.Marshal(body)

	return string(json)
}