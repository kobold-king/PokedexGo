package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// This is where we send the GET
func (c *Client) ReadWriteLocations(url *string) (Locations, error) {
	newURL := baseURL + "/location-area"
	if url != nil {
		newURL = *url
	}
	// 1. Send the Get request
	resp, err := http.Get(newURL)
	if err != nil {
		return Locations{}, err
	}

	defer resp.Body.Close()

	// success check
	if resp.StatusCode != http.StatusOK {
		return Locations{}, err
	}

	// Read the entire response body into a byte slice.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}

	//put into stuct
	locationsResp := Locations{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return Locations{}, err
	}
	return locationsResp, nil
}
