package main

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"

	"./goub"
)

func main() {
	appId := "e8132af91b61997f4fd9ba74e1dcb8c05aae0ee7ea47f96a48731310b25edb61"
	secret := "4d150fe1ce9fb8a05225cfdee65bf9339b885689b306444d80961d5ecd5652a0"
	callbackURL := "http://llocalhost:8000/"

	conf := goub.MakeConfig(appId, secret, callbackURL)
	// fmt.Println(goub.GetOauthCodeLink(conf, "state", oauth2.AccessTypeOffline))

	// code := "32b8a0e54cfc090a5f2518a2ce1a6408bc07eb7125d9a654ca1cbb9aa9e92db3"
	// token, err := goub.GetTokenFromCode(conf, code)

	token := &oauth2.Token{
		AccessToken: "a854ffda6a4a31ebb6c830b35b58734e1e8224caaa2836585a0598c6d2cd0966",
	}

	client, err := goub.NewClient(conf, token)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := client.NewRequest("GET", "api/v2/likes/by_channel?channel_id=1164719", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(req)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("")
	fmt.Println(string(body))

}
