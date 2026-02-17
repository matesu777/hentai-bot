package media

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type response struct{
	Url string `json:"url"`
}

func GetApi(typer string, category string)(response, error){
	var resObject response
	res, err := http.Get("https://api.waifu.pics/" + typer + "/" + category)
	if err != nil{
		fmt.Println(err)
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println(err)
	}
	json.Unmarshal(resData, &resObject)
	return response{resObject.Url}, err
}