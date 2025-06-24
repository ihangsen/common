package redis

import (
	"context"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/utils/trans"
)

func GeoAdd(key string, latitude float64, longitude float64, member uint64) {
	catch.Try(client.Do(context.Background(), "geoadd", key, latitude, longitude, member).Err())
}

func GeoDist(key, member1, member2 string) float64 {
	return catch.Try1(client.Do(context.Background(), "geodist", key, member1, member2).Float64())
}

func GeoRem(key string, members vec.Vec[string]) {
	catch.Try(client.ZRem(context.Background(), key, members).Err())
}

type GPos struct {
	Longitude, Latitude float64
}

func GeoPos(key string, members vec.Vec[uint64]) vec.Vec[GPos] {
	args := make([]any, 0, 2+len(members))
	args = append(args, "geopos", key)
	for _, member := range members {
		args = append(args, member)
	}
	result := catch.Try1(client.Do(context.Background(), args...).Slice())
	vector := vec.New[GPos](len(result))
	for _, item := range result {
		v := item.([]any)
		vector.Append(GPos{
			Longitude: v[0].(float64),
			Latitude:  v[1].(float64),
		})
	}
	return vector
}

type GDist struct {
	Name string
	Dist float64
}

func GeoSearch(key string, longitude float64, latitude float64, meter uint32) vec.Vec[GDist] {
	result := catch.Try1(client.Do(
		context.Background(), "geosearch", key, "FROMLONLAT", longitude,
		latitude, "BYRADIUS", meter, "m", "asc", "COUNT", 1, "WITHDIST").
		Slice())
	vector := vec.New[GDist](len(result))
	for _, item := range result {
		v := item.([]any)
		vector.Append(GDist{
			Name: v[0].(string),
			Dist: trans.Str2F64(v[1].(string)),
		})
	}
	return vector
}
