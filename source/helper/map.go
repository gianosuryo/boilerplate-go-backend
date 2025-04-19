package helper

import "math"

type MapCoordinate struct {
	Latitude  float64
	Longitude float64
}

// DistanceResult holds distances in different units
type DistanceResult struct {
	Kilometers float64
	Miles      float64
}

func CalculateDistance(p1, p2 MapCoordinate) DistanceResult {
	// Earth's radius in kilometers
	const earthRadius = 6371.0
	const kmToMiles = 0.621371

	// Convert latitude and longitude from degrees to radians
	lat1 := degreesToRadians(p1.Latitude)
	lon1 := degreesToRadians(p1.Longitude)
	lat2 := degreesToRadians(p2.Latitude)
	lon2 := degreesToRadians(p2.Longitude)

	// Differences in coordinates
	dLat := lat2 - lat1
	dLon := lon2 - lon1

	// Haversine formula
	a := math.Pow(math.Sin(dLat/2), 2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Pow(math.Sin(dLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate the distances
	distanceKm := earthRadius * c
	distanceMiles := distanceKm * kmToMiles

	return DistanceResult{
		Kilometers: distanceKm,
		Miles:      distanceMiles,
	}
}

// degreesToRadians converts degrees to radians
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
