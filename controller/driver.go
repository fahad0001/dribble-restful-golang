package controller

import (
	"gallery-app/dataDir"
	"gallery-app/models"
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"io"
)

func InitializeApp(dataSource string) {
	dataDir.NewDBInstance(dataSource)
}

func saveImageFile(imgFileName string, imgSrcLink string) {

	response, err := http.Get(imgSrcLink)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(imgFileName + ".jpg")
	if err != nil {
		panic(err)
	}
	// Use io.Copy to just dump the response body to the file.
	_, err = io.Copy(file, response.Body)
	if err != nil {
		panic(err)
	}
	file.Close()
	fmt.Println("Success!")
}

func imageRequest(imgName string, imgSource string, accessToken string, ch chan <- *models.Photo) {
	dribbleUrl := "https://api.dribbble.com/v1/shots/" + imgName

	photoModel := new(models.Photo)
	client := &http.Client{}
	if req, err := http.NewRequest("GET", dribbleUrl, nil); err !=nil {
		panic(err)
	} else {
		req.Header.Set("Authorization", "Bearer " + accessToken)
		client.CheckRedirect = checkRedirectFunc

		if resp, err := client.Do(req); err != nil {
			panic(err)
		} else {
			defer resp.Body.Close()

			if err := json.NewDecoder(resp.Body).Decode(&photoModel); err != nil {
				panic(err)
			} else {
				photoModel.FileName = imgSource + photoModel.Title
				saveImageFile(photoModel.FileName, photoModel.DribbleLink.Normal)
				ch <- photoModel
			}
		}
	}
}

func GetDribbleData(dataSource string, imgSource string, accessToken string) {
	imageNameList := []string{"3775242-Love-Obsession", "3952963-4-8-axis", "3908535-404-Page-Exploration", "3386240-Owl-Character-Design",
		"1607412-Bear-Logo-mark", "2781214-Atom", "3156461-Boston-Lobster", "1188443-Sea", "3954699-witch-s-magic"}

	ch := make(chan *models.Photo, len(imageNameList))
	responses := []*models.Photo{}
	for _, img := range imageNameList {
		go imageRequest(img, imgSource, accessToken, ch)
	}
	for len(responses) != len(imageNameList) {
		responses = append(responses, <-ch)
	}
	for _, response := range responses {
		dataDir.LoadDataToTable(dataSource, *response)
	}
}

func checkRedirectFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", via[0].Header.Get("Authorization"))
	return nil
}