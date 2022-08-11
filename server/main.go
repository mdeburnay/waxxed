package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Pagination struct {
		Page int `json:"page"`
		Pages int `json:"pages"`
		PerPage int `json:"per_page"`
		Items int `json:"items"`
		Urls struct {
			Last string `json:"last"`
			Next string `json:"next"`
		} `json:"urls"`
	} `json:"pagination"`
	Wants []struct {
		Id int `json:"id"`
		ResourceUrl string `json:"resource_url"`
		DateAdded string `json:"date_added"`
		Rating int `json:"rating"`
		BasicInformation struct {
			Id int `json:"id"`
			MasterId int `json:"master_id"`
			MasterUrl string `json:"master_url"`
			ResourceUrl string `json:"resource_url"`
			Thumb string `json:"thumb"`
			CoverImage string `json:"cover_image"`
			Title string `json:"title"`
			Year int `json:"year"`
			Formats []struct {
				Name string `json:"name"`
				Quantity int `json:"qty"`
				Descriptions []string `json:"descriptions"`
			} `json:"formats"`
			Labels []struct {
				Name string `json:"name"`
				CatNumber string `json:"catno"`
				EntityType string `json:"entity_type"`
				EntityTypeName string `json:"entity_type_name"`
				Id int `json:"id"`
				ResourceUrl string `json:"resource_url"`
			} `json:"labels"`
			Artists []struct {
				Name string `json:"name"`
				Anv string `json:"anv"`
				Join string `json:"join"`
				Role string `json:"role"`
				Tracks int `json:"tracks"`
				Id int `json:"id"`
				ResourceUrl string `json:"resource_url"`
			} `json:"artists"`
			Genres []string `json:"genres"`
			Styles []string `json:"styles"`
		} `json:"basic_information"`
		FolderId int `json:"folder_id"`
	} `json:"wants"`
}

func main() {
	envErr := godotenv.Load("../.env")
	if(envErr != nil) {
		fmt.Println(envErr)
		os.Exit(1)
	}

	client := &http.Client{}
	discogsUrl := os.Getenv("DISCOGS_BASE_URL")
	discogsToken := os.Getenv("DISCOGS_TOKEN")

	req, err := http.NewRequest("GET", discogsUrl + "/users/maximoose95/wants", nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Authorization": {"Discogs token=" + discogsToken},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	json.Unmarshal(responseData, &response)

	var wantListItems []string


	for _, p := range response.Wants {
		wantListItems = append(wantListItems, p.BasicInformation.ResourceUrl)
	}

	fmt.Println(wantListItems)
}
