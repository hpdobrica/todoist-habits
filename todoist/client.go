package todoist

import (
	"io"
	"net/http"
	"os"
)

var client = &http.Client{}

type TodoistClient struct {
	// client http.Client
	Token string
}

func (c *TodoistClient) get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

// func (c *TodoistClient) getObj(url string, obj interface{}) (resp interface{}, err error) {
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resBody, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Printf("client: could not read response body: %s\n", err)
// 		os.Exit(1)
// 	}
// 	fmt.Printf("client: response body: %s\n", resBody)

// 	var task TodoistTask

// 	if err := json.Unmarshal(resBody, &task); err != nil { // Parse []byte to go struct pointer
// 		fmt.Println("Can not unmarshal JSON")
// 	}

// 	return c.do(req)
// }

func (c *TodoistClient) post(url string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *TodoistClient) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+os.Getenv("TODOIST_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	return client.Do(req)

}
