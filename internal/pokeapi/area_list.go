package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
	// "fmt"
)

func (c *Client) GetLocation(area string) (RespShallowAreas, error)  {

	url := baseURL + "/location-area/" + area

	// fmt.Printf("URL: %s", url)

	if val, ok := c.cache.Get(url); ok {
		areaResp := RespShallowAreas{}
		err := json.Unmarshal(val, &areaResp)
		if err != nil {
			return RespShallowAreas{}, err
		}
		
		return areaResp, nil
	}



	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowAreas{}, err
	}
	

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowAreas{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowAreas{}, err
	}


	areaResp := RespShallowAreas{}
	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return RespShallowAreas{}, err
	}
	
	c.cache.Add(url, dat)
	return areaResp, nil
}




