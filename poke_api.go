package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"errors"
)

func getMaps(mapsResponse *getMapResponse, mainConfig *config, usePreviousUrl bool) error{
	var res *http.Response
	var err error

	res, err = fetchLocationAreas(mainConfig, usePreviousUrl)

	if err != nil {
		return err
	} 

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %v and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, mapsResponse)
	if err != nil {
		return err
	}
	return nil
}

func fetchLocationAreas(mainConfig *config, usePreviousUrl bool) (*http.Response, error) {
	if usePreviousUrl {
		if mainConfig.Previous == "" {
			return nil, errors.New("you're on the first page")
		} else {
			return http.Get(mainConfig.Previous)
		}
	} else {
		if mainConfig.Next == "" {
			return http.Get("https://pokeapi.co/api/v2/location-area")
		} else {
			return http.Get(mainConfig.Next)
		}
	}
}