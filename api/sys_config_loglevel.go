package api

import "errors"

func (c *Sys) GetLoglevel() (int, error) {
	r := c.c.NewRequest("GET", "/v1/sys/config/loglevel")
	resp, err := c.c.RawRequest(r)
	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode == 404 {
			return -1, nil
		}
	}
	if err != nil {
		return -1, err
	}

	var result map[string]interface{}
	err = resp.DecodeJSON(&result)
	if err != nil {
		return -1, err
	}

	loglevel, ok := result["loglevel"].(int)
	if !ok {
		return -1, errors.New("invalid log level")
	}

	return loglevel, nil
}
