package services

import (
	"encoding/json"
	"strings"

	"github.com/charlesonunze/busha-test/database"
	"github.com/charlesonunze/busha-test/model"
	"github.com/go-resty/resty/v2"
)

var (
	baseURL    = "https://swapi.dev/api"
	httpClient = resty.New().R()
)

func GetMovies() ([]model.Movie, error) {
	url := baseURL + "/films"
	movieResponse := &model.MovieResponse{}
	results := []model.Movie{}
	key := "movies_list"

	data, err := FetchFromCache(key)
	if err != nil {
		return results, err
	}

	if len(data) > 0 {
		return data, err
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return results, err
	}

	err = json.Unmarshal([]byte(resp.Body()), &movieResponse)
	if err != nil {
		return results, err
	}

	for _, m := range movieResponse.Results {
		split := strings.Split(m.Url, "films/")
		movieId := string([]rune(split[1])[0])

		count, err := GetCommentsCount(movieId)
		if err != nil {
			return results, err
		}

		m.CommentCount = count

		results = append(results, m)
	}

	err = StoreInCache(results, key)
	if err != nil {
		return results, err
	}

	return results, nil
}

func GetCommentsCount(movieId string) (int64, error) {
	db := database.DB
	var comments []model.Comment

	result := db.Where("movie_id = ?", movieId).Find(&comments)

	return result.RowsAffected, result.Error
}
