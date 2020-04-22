package create_data_to_chain

import (
	"encoding/hex"
	"encoding/json"
	"github.com/material-on-chain/go-owcrypt"
)

type PLAResin struct {
	Product string `json:"product"`
	Company string `json:"company"`
	SalesManager string `json:"sales_manager"`
	ProducingArea string `json:"producing_area"`
	GlassTransitionTemperatureInCentigrade float32 `json:"glass_transition_temperature"`
	DensityInGramPerCubicCentimeter float32 `json:"density"`
	ElasticModulusInGigaPascal float32 `json:"elastic_modulus"`
	TensileStrengthInGigaPascal float32 `json:"tensile_strength"`
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

func (pr PLAResin) GetDataToChain() string {
	fingerPrint := owcrypt.Hash([]byte(pr.ToString()),0, owcrypt.HASH_ALG_SHA3_256_RIPEMD160)
	var (
		body = make(map[string]interface{}, 0)
	)

	body["company"] = pr.Company
	body["finger_print"] = hex.EncodeToString(fingerPrint)

	json, _ := json.Marshal(body)

	return string(json)
}