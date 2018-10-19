package function

import (
		"encoding/json"
)

const (
	Inside = "inside"
	Outside = "outside"
	Error = "error"
)

type Location struct {
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Request struct {
	Point *Location `json:"point"`
	Polygon []*Location `json:"polygon"`
}

// Handle a serverless request
func Handle(req []byte) string {
	request := &Request{}

	if err := json.Unmarshal(req, request); err != nil {
		return Error
	}

	inside := false
	for i, j := 0, len(request.Polygon) - 1; i < len(request.Polygon); i = i + 1 {
		if (((request.Polygon[i].Latitude <= request.Point.Latitude) && (request.Point.Latitude < request.Polygon[j].Latitude)) || ((request.Polygon[j].Latitude <= request.Point.Latitude) && (request.Point.Latitude < request.Polygon[i].Latitude))) && (request.Point.Longitude < (request.Polygon[j].Longitude - request.Polygon[i].Longitude) * (request.Point.Latitude - request.Polygon[i].Latitude) / (request.Polygon[j].Latitude - request.Polygon[i].Latitude) + request.Polygon[i].Longitude) {
			inside = !inside
		}
		j = i
	}

	if inside {
		return Inside
	} else {
		return Outside
	}
}
