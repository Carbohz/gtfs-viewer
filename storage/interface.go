package storage

import "context"

type GtfsDataStorage interface {
	SaveRoute(ctx context.Context, m model.Route) error
}