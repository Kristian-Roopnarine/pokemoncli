package pokeapi

import (
	"encoding/json"
	"errors"
	pokecache "github.com/Kristian-Roopnarine/pokemoncli/internal/pokecache"
	"io"
	"net/http"
)

const RootUrl = "https://pokeapi.co/api/v2"

type PokeApiResponse interface {
	LocationResponse |
		LocationAreaResponse
}

func Get[T PokeApiResponse](url string, cache pokecache.Cache) (T, error) {
	var pokeResponse T
	if cacheEntry, ok := cache.Data[url]; ok {
		err := json.Unmarshal(cacheEntry.Val, &pokeResponse)
		if err != nil {
			return pokeResponse, errors.New(err.Error())
		}
		return pokeResponse, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return pokeResponse, errors.New(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	cache.Add(url, body)
	if err != nil {
		return pokeResponse, errors.New(err.Error())
	}
	err = json.Unmarshal(body, &pokeResponse)
	if err != nil {
		return pokeResponse, errors.New(err.Error())
	}
	return pokeResponse, nil

}
