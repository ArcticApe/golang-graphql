package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/golang-graphql/datasource/electricitymap"
	"github.com/golang-graphql/datasource/transportforlondon"
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

// LivePowerBreakdown is the resolver for the livePowerBreakdown field.
func (r *queryResolver) LivePowerBreakdown(ctx context.Context, zone string) (*model.PowerBreakdown, error) {
	param := electricitymap.TypAPIParams{Zone: zone}
	ci, err := electricitymap.GetLivePowerBreakdown(param)
	if err != nil {
		return nil, err
	}
	return &ci, nil
}

// LiveStopPointFares is the resolver for the liveStopPointFares field.
func (r *queryResolver) LiveStopPointFares(ctx context.Context) ([]*transportforlondon.StopPointFares, error) {
	ci, err := transportforlondon.GetStopPointFares()
	if err != nil {
		return nil, err
	}
	var result []*transportforlondon.StopPointFares
	for _, s := range ci {
		result = append(result, &s)
	}
	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
