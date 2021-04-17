package domain

import (
	"encoding/json"
	"fmt"
	"io"
)

type PDV struct {
	TrandingName string       `json:"tradingName" bson:"tradingName"`
	CoverageArea MultiPolygon `json:"coverageArea" bson:"coverageArea"`
}

type MultiPolygon struct {
	Type        string           `json:"type" bson:"type"`
	Coordinates [][][][2]float64 `json:"coordinates" bson:"coordinates"`
}

func (m *MultiPolygon) UnmarshalGQL(v interface{}) error {
	value, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("invalid raw multiPolygon json")
	}

	var mp MultiPolygon

	err = json.Unmarshal(value, &mp)
	if err != nil {
		return fmt.Errorf("value is not a valid multiPolygon")
	}

	*m = mp

	return nil
}

func (m MultiPolygon) MarshalGQL(w io.Writer) {
	mp, err := json.Marshal(m)
	if err != nil {
		fmt.Println("multiPolygon marshalling fail")
		return
	}

	_, err = w.Write(mp)
	if err != nil {
		fmt.Println("multiPolygon writing fail")
		return
	}
}
