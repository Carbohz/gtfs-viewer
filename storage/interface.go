package storage

import (
	"context"
	"github.com/Carbohz/gtfs-viewer/model"
)

type GtfsDataStorage interface {
	SaveRoute(ctx context.Context, m model.Route) error
}