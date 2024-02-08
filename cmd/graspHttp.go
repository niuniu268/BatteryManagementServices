package cmd

import (
	`bytes`
	`fmt`
	`io`
	`io/ioutil`
	`net/http`
)

func GetHttp(url string) []byte {

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending GET request: %v", err)

	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing GET request: %v", err)
		}
	}(resp.Body)

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v", err)

	}

	return body

}

func PostHttp(url, key, value string) []byte {

	jsonData := fmt.Sprintf(`{"%s": "%s"}`, key, value)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v", err)

	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	return body

}
