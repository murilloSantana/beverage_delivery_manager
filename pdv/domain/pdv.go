package domain

import (
	logger "beverage_delivery_manager/config/log"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
)

type MultiPolygonCoordinates [][][][2]float64
type PointCoordinates []float64

type Pdv struct {
	ID           string       `json:"id" bson:"_id"`
	TradingName  string       `json:"tradingName" bson:"tradingName"`
	OwnerName    string       `json:"ownerName" bson:"ownerName"`
	Document     string       `json:"document" bson:"document"`
	CoverageArea MultiPolygon `json:"coverageArea" bson:"coverageArea"`
	Address      Point        `json:"address" bson:"address"`
}

type MultiPolygon struct {
	Type        string                  `json:"type" bson:"type"`
	Coordinates MultiPolygonCoordinates `json:"coordinates" bson:"coordinates"`
}

type Point struct {
	Type        string           `json:"type" bson:"type"`
	Coordinates PointCoordinates `json:"coordinates" bson:"coordinates"`
}

func (m *MultiPolygon) UnmarshalGQL(v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

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
	log := logger.NewLogger()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	mp, err := json.Marshal(m)
	if err != nil {
		log.Error(nil, "multiPolygon marshalling fail")
		return
	}

	_, err = w.Write(mp)
	if err != nil {
		log.Error(nil, "multiPolygon writing fail")
		return
	}
}

func (p *Point) UnmarshalGQL(v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

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
	log := logger.NewLogger()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	point, err := json.Marshal(p)
	if err != nil {
		log.Error(nil, "point marshalling fail")
		return
	}

	_, err = w.Write(point)
	if err != nil {
		log.Error(nil, "point writing fail")
		return
	}
}
