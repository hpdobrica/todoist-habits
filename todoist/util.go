package todoist

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func resIntoObj[T any](resp *http.Response, obj *T) error {
	resBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close() // move resp body close out of here TODO
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resBody, obj); err != nil {
		return err
	}

	return nil
}
