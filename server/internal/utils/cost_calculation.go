package utils

import (
	"math"
	"strings"
)

type ShippingOption struct {
	Service     string  `json:"service"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
	ETD         string  `json:"etd"`
}

func EstimateShippingRates(
	originProvinceID, originCityID,
	destProvinceID, destCityID int,
	weight int,
	courier string,
) ([]ShippingOption, error) {
	var base float64

	switch {
	case originProvinceID == destProvinceID && originCityID == destCityID:
		base = 5000
	case originProvinceID == destProvinceID:
		base = 10000
	default:
		base = 15000
	}

	roundedWeight := math.Ceil(float64(weight) / 1000)

	rateTable := map[string]map[string]float64{
		"jne": {
			"OKE": 1.0,
			"REG": 1.3,
			"YES": 2.5,
		},
		"sicepat": {
			"HEMAT": 1.1,
			"REG":   1.4,
			"BEST":  2.0,
		},
	}

	etdMap := map[string]string{
		"OKE":   "4-5 days",
		"REG":   "2-3 days",
		"YES":   "1 days",
		"HEMAT": "4-5 days",
		"BEST":  "1-2 days",
	}

	descMap := map[string]string{
		"OKE":   "Ongkos Kirim Ekonomis",
		"REG":   "Layanan Reguler",
		"YES":   "Yakin Esok Sampai",
		"HEMAT": "SiCepat Hemat",
		"BEST":  "Best Express Service",
	}

	courier = strings.ToLower(courier)
	rates, ok := rateTable[courier]
	if !ok {
		return nil, nil
	}

	var result []ShippingOption
	for service, multiplier := range rates {
		cost := base + (multiplier * 10000 * roundedWeight)
		result = append(result, ShippingOption{
			Service:     service,
			Description: descMap[service],
			Cost:        cost,
			ETD:         etdMap[service],
		})
	}

	return result, nil
}
