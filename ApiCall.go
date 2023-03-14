package groupietrackers


import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"fmt"
	"os"
)

func APICall(url string, Dataform interface{}) {
	/*
	* Function who call the API and return the data
	 */
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	json.Unmarshal(data, Dataform)
}