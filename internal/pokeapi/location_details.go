package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) LocationDetails(area string) (RespDeepLocations, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		detailsResp := RespDeepLocations{}
		err := json.Unmarshal(val, &detailsResp)
		if err != nil {
			return RespDeepLocations{}, err
		}

		return detailsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocations{}, err
	}

	detailsResp := RespDeepLocations{}
	err = json.Unmarshal(dat, &detailsResp)
	if err != nil {
		return RespDeepLocations{}, err
	}

	c.cache.Add(url, dat)
	return detailsResp, nil
}
