package graphql

import (
	"context"

	"github.com/iot-for-tillgenglighet/api-snowdepth/pkg/database"
	"github.com/iot-for-tillgenglighet/api-snowdepth/pkg/models"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Entity() EntityResolver {
	return &entityResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type entityResolver struct{ *Resolver }

func (r *entityResolver) FindDeviceByID(ctx context.Context, id string) (*Device, error) {
	return &Device{ID: id}, nil
}

type mutationResolver struct{ *Resolver }

func convertDatabaseRecordToGQL(measurement *models.Snowdepth) *Snowdepth {
	if measurement != nil {
		depth := &Snowdepth{
			From: &Origin{
				Pos: &WGS84Position{
					Lat: measurement.Latitude,
					Lon: measurement.Longitude,
				},
			},
			When:  measurement.Timestamp,
			Depth: float64(measurement.Depth),
		}

		if len(measurement.Device) == 0 {
			depth.Manual = &[]bool{true}[0] // <- You may Google that little nugget of beauty ...
		} else {
			depth.From.Device = &Device{ID: measurement.Device}
		}

		return depth
	}

	return nil
}

func (r *mutationResolver) AddSnowdepthMeasurement(ctx context.Context, input NewSnowdepthMeasurement) (*Snowdepth, error) {
	measurement, err := database.AddManualSnowdepthMeasurement(input.Pos.Lat, input.Pos.Lon, input.Depth)
	return convertDatabaseRecordToGQL(measurement), err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Snowdepths(ctx context.Context) ([]*Snowdepth, error) {
	depths, err := database.GetLatestSnowdepths()

	if err != nil {
		panic("Failed to query latest snowdepths: " + err.Error())
	}

	depthcount := len(depths)

	if depthcount == 0 {
		return []*Snowdepth{}, nil
	}

	gqldepths := make([]*Snowdepth, 0, depthcount)

	for _, v := range depths {
		gqldepths = append(gqldepths, convertDatabaseRecordToGQL(&v))
	}

	return gqldepths, nil
}

