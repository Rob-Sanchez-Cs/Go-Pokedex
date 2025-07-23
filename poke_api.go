package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Rob-Sanchez-Cs/Go-Pokedex/internal/pokecache"
)

func getMaps(mapsResponse *getMapResponse, mainConfig *config, usePreviousUrl bool, cache *pokecache.Cache) error{
	var res *http.Response
	var err error

	apiURL := determineUrlForArea(mainConfig, usePreviousUrl)
	if apiURL == "" {
		return errors.New("you're on the first page")
	}

	cacheEntry, found := cache.Get(apiURL)
	if found {
		err = json.Unmarshal(cacheEntry, mapsResponse)
		if err != nil {
			return err
		}
		return nil
	}

	res, err = fetchLocationAreas(apiURL)

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
	cache.Add(apiURL, body)
	return nil
}

func fetchLocationAreas(url string) (*http.Response, error) {
	return http.Get(url)
}

func determineUrlForArea(mainConfig *config, usePreviousUrl bool) string {
	if usePreviousUrl {
		if mainConfig.Previous == "" {
			return ""
		} else {
			return mainConfig.Previous
		}
	} else {
		if mainConfig.Next == "" {
			return "https://pokeapi.co/api/v2/location-area"
		} else {
			return mainConfig.Next
		}
	}
}

func exploreLocation(exploreResponse *getExploreResponse,cache *pokecache.Cache, parameter string) error {
	var res *http.Response
	var err error

	exploreApiUrl := constructExploreUrl(parameter)

	cacheEntry, found := cache.Get(exploreApiUrl)
	if found {
		err = json.Unmarshal(cacheEntry, exploreResponse)
		if err != nil {
			return err
		}
		return nil
	}

	res, err = http.Get(exploreApiUrl)

	if err != nil {
		return err
	} 

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %v and body: %s", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, exploreResponse)
	if err != nil {
		return err
	}
	cache.Add(exploreApiUrl, body)
	return nil

}

func constructExploreUrl(parameter string) string{
	return "https://pokeapi.co/api/v2/location-area/" + parameter
}