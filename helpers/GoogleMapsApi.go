package helpers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"food/models"
	"io/ioutil"
	"net/http"

	"googlemaps.github.io/maps"
)

const ApiKey = ""

type DistanceMatrix struct {
	Destination_addresses []string `json:"destination_addresses"`
	Origin_addresses      []string `json:"origin_addresses"`
	Rows                  []struct {
		Elements []struct {
			Distance struct {
				Text  string  `json:"text"`
				Value float64 `json:"value"`
			} `json:"distance"`
			Duration struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration"`
			Status string `json:"status"`
		} `json:"elements"`
	} `json:"rows"`
	Status string `json:"status"`
}

func GetCoordinates(address *models.Address) (float64, float64, error) {
	c, err := maps.NewClient(maps.WithAPIKey(ApiKey))
	if err != nil {
		return 0, 0, err
	}
	r := &maps.GeocodingRequest{
		Address: fmt.Sprintf("%s %s %s %s %s", address.Neighbourhood, address.Street, address.BuildingNo, address.District, address.City),
	}
	response, err := c.Geocode(context.Background(), r)
	if err != nil {
		return 0, 0, err
	}
	lat := response[0].Geometry.Location.Lat
	lng := response[0].Geometry.Location.Lng

	return lat, lng, nil
}

func GetDistance(uaddress models.Address, raddress models.Address) (float64, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?origins=%f,%f&destinations=%f,%f&key=%s", raddress.Lat, raddress.Lng, uaddress.Lat, uaddress.Lng, ApiKey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("request is wrong ", err)
		return 0, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("response couldn't reveive ", err)
		return 0, err
	}
	defer res.Body.Close() //close the response body after this function dies.

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("response can not be readed ", err)
		return 0, err
	}
	var results DistanceMatrix
	err = json.Unmarshal(body, &results)
	if err != nil || results.Status != "OK" {
		return 0, errors.New("Request is rejected!")
	}
	return results.Rows[0].Elements[0].Distance.Value, nil
}
