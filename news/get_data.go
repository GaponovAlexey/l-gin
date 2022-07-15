package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type source struct {
	ID string `json:"id"`
}

type sourcesAPI struct {
	Sources []source `json:"sources"`
}

type topicAPI struct {
	Articles []Topic `json:"articles`
}

func getSources(category string) []source {
	body := getData(sourceURL(category))

	var sourceAPI sourcesAPI
	json.Unmarshal(body, &sourceAPI)

	return sourceAPI.Sources
}

func getTopics(sources []source) []Topic {
	var topics []Topic
	for _, source := range sources {
		body := getData(topicURL(source.ID))
		var topicAPI topicsAPI
		json.Unmarshal(body, &topicAPI)
		topics = append(topics, topicAPI.Articles...)
	}

	return topics
}

func sourceURL(category string) string {
	return fmt.Sprintf("https://newsapi.org/v1/sources?language=en&category=%s", category)
}

func topicURL(id string) string {
	return fmt.Sprintf("https://newsapi.org/v1/sources?language=en&category=%s", id)
}

func getData(url string) []byte {
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
