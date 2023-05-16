package neargo_test

import (
	"testing"

	"github.com/refaldyrk/neargo"
)

func TestNeargo(t *testing.T) {
	listLocation := []neargo.Location{
		{Latitude: 51.5007, Longitude: -0.1246},   // London Eye
		{Latitude: 51.5162, Longitude: -0.0945},   // Tower Bridge
		{Latitude: 51.5194, Longitude: -0.1269},   // British Museum
		{Latitude: 51.4975, Longitude: -0.1753},   // Buckingham Palace
		{Latitude: 51.5234, Longitude: -0.1589},   // Oxford Street
		{Latitude: 40.7128, Longitude: -74.0060},  // New York City
		{Latitude: 51.5074, Longitude: -0.1278},   // London
		{Latitude: -33.8679, Longitude: 151.2073}, // Sydney
		{Latitude: 35.6895, Longitude: 139.6917},  // Tokyo
		{Latitude: -22.9068, Longitude: -43.1729}, //Rio De Janiero
	}

	nearbyLocation := neargo.Nearby(51.5074, -0.1278, 5.0, listLocation)

	//Test
	t.Log(nearbyLocation) //[{51.5007 -0.1246} {51.5162 -0.0945} {51.5194 -0.1269} {51.4975 -0.1753} {51.5234 -0.1589} {51.5074 -0.1278}]

}
