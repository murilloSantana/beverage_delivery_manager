package repository

import (
	"beverage_delivery_manager/pdv/domain"
	"context"
)

//go:generate mockery --name PdvRepository --case=underscore --output ../../mocks

// PdvRepository it is the interface that involves functions that make the integration with storage possible
//
// HasDocument Although it is possible to add settings in mongoDB to ensure that the value of a document
// is unique for each user, it was decided to  create the HasDocument function to remove the dependency
// on external tools, in this case, if it is necessary to use another db, this business rule would still
// be protected from external modifications. The same was not done for the id because the identifier is
// dynamically generated in the application and not informed by the client
type PdvRepository interface {
	HasDocument(document string) (bool, error)
	Save(ctx context.Context, pdv domain.Pdv, generateNewID func() string) (domain.Pdv, error)
	FindByID(ID string) (domain.Pdv, error)
	FindByAddress(point domain.Point) (domain.Pdv, error)
	GenerateNewID() func() string
}
