package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository/models"
)

type ExternalApiClient struct {
	BASE_URL string
}

func (api ExternalApiClient) GetInfo(group string, song string) (models.SongDetail, error) {
	params := url.Values{}
	params.Add("group", group)
	params.Add("song", song)

	apiURL := fmt.Sprintf("%s/info?%s", api.BASE_URL, params.Encode())

	res, err := http.Get(apiURL)

	if err != nil {
		slog.Error("Failed to get a response from the External API.", slog.Any("error", err))
		return models.SongDetail{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		slog.Error("Request executed with invalid code.", slog.Any("status_code", res.StatusCode))
		return models.SongDetail{}, errors.New("invalid status code: " + strconv.Itoa(res.StatusCode))
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		slog.Error("Failed to read a response from the External API.", slog.Any("error", err))
		return models.SongDetail{}, err
	}

	var songDetail models.SongDetail

	err = json.Unmarshal(body, &songDetail)

	if err != nil {
		slog.Error("Failed to decode a response to JSON from the External API.", slog.Any("error", err))
		return models.SongDetail{}, err
	}

	return songDetail, nil
}
