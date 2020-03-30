package app

import (
	_ "github.com/whosonfirst/go-whosonfirst-spatial/database/spatial"
)

import (
	"context"
	"flag"
	"github.com/whosonfirst/go-whosonfirst-spatial/database"
	"github.com/whosonfirst/go-whosonfirst-spatial/flags"
)

func NewSpatialDatabaseWithFlagSet(ctx context.Context, fl *flag.FlagSet) (database.SpatialDatabase, error) {

	spatial_uri, err := flags.StringVar(fl, "spatial-database")

	if err != nil {
		return nil, err
	}

	return database.NewSpatialDatabase(ctx, spatial_uri)
}
