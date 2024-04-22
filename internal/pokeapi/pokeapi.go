package pokeapi

import (
	"encoding/json"
	"errors"
	pokecache "github.com/Kristian-Roopnarine/pokemoncli/internal/pokecache"
	"io"
	"net/http"
)

const RootUrl = "https://pokeapi.co/api/v2"

type PokeApiResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func Get(url string, cache pokecache.Cache) (PokeApiResponse, error) {
	if cacheEntry, ok := cache.Data[url]; ok {
		pokeResponse := PokeApiResponse{}
		err := json.Unmarshal(cacheEntry.Val, &pokeResponse)
		if err != nil {
			return PokeApiResponse{}, errors.New(err.Error())
		}
		return pokeResponse, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return PokeApiResponse{}, errors.New(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	cache.Add(url, body)
	if err != nil {
		return PokeApiResponse{}, errors.New(err.Error())
	}
	pokeResponse := PokeApiResponse{}
	err = json.Unmarshal(body, &pokeResponse)
	if err != nil {
		return PokeApiResponse{}, errors.New(err.Error())
	}
	return pokeResponse, nil

}
