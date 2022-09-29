package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/golang-graphql/datasource/electricitymap"
	"github.com/golang-graphql/graph/generated"
	"github.com/golang-graphql/graph/model"
)

// LiveCarbonIntensity is the resolver for the liveCarbonIntensity field.
func (r *queryResolver) LiveCarbonIntensity(ctx context.Context, zone string) (*model.CarbonIntensity, error) {
	param := electricitymap.TypAPIParams{Zone: zone}
	ci, err := electricitymap.GetLiveCarbonIntensity(param)
	if err != nil {
		return nil, err
	}
	return &ci, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
