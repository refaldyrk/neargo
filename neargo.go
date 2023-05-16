package neargo

import "math"

const (
	earthRadius = 6371 //Earth Radius
)

type Location struct {
	Latitude  float64
	Longitude float64
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	dLat := deg2rad(lat2 - lat1)
	dLon := deg2rad(lon2 - lon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c
	return distance
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func Nearby(yourlat, yourlon, maxDistance float64, opts []Location) []Location {
	currentLoc := Location{
		Latitude:  yourlat,
		Longitude: yourlon,
	}

	var nearestPlaces []Location
	for _, place := range opts {
		distance := haversine(currentLoc.Latitude, currentLoc.Longitude, place.Latitude, place.Longitude)
		if distance <= maxDistance {
			nearestPlaces = append(nearestPlaces, place)
		}
	}

	return nearestPlaces
}
