package domain

import (
	"encoding/json"
	"fmt"
	"io"
)

//TODO configure marshal/unmarshal different from buit-in (it is already known that there are more performant alternatives)

type PDV struct {
	ID           string       `json:"id" bson:"_id"`
	TrandingName string       `json:"tradingName" bson:"tradingName"`
	OwnerName    string       `json:"ownerName" bson:"ownerName"`
	Document     string       `json:"document" bson:"document"`
	CoverageArea MultiPolygon `json:"coverageArea" bson:"coverageArea"`
	Address      Point        `json:"address" bson:"address"`
}

type MultiPolygon struct {
	Type        string           `json:"type" bson:"type"`
	Coordinates [][][][2]float64 `json:"coordinates" bson:"coordinates"`
}

type Point struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

//TODO look for alternatives so as not to need to marshal the value received
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

func (p *Point) UnmarshalGQL(v interface{}) error {
	value, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("invalid raw point json")
	}

	var point Point

	err = json.Unmarshal(value, &point)
	if err != nil {
		return fmt.Errorf("value is not a valid point")
	}

	*p = point

	return nil
}

func (p Point) MarshalGQL(w io.Writer) {
	point, err := json.Marshal(p)
	if err != nil {
		fmt.Println("point marshalling fail")
		return
	}

	_, err = w.Write(point)
	if err != nil {
		fmt.Println("point writing fail")
		return
	}
}
