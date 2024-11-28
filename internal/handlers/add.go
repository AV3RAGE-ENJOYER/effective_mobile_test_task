package handlers

import (
	"log/slog"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/requests"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/utils"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository/api"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Add song
// @Schemes
// @Description Add song to the database
// @Tags API
// @Accept json
// @Produce json
// @Param body body requests.AddSongRequest true "Request body"
// @Success 200 {object} models.SongDetail
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /add [post]
func AddSongHandler(api api.ExternalApiClient, db repository.DatabaseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requests.AddSongRequest

		err := c.BindJSON(&req)

		if err != nil {
			slog.Error("Failed to unmarshall json.", slog.Any("error", err))
			c.String(400, "Bad request")
			return
		}

		slog.Debug("Add request.", slog.Any("request", req))

		err = db.AddSong(c, req)

		if err != nil {
			slog.Error("Failed to add song to the database.", slog.Any("error", err))
			c.String(500, "Internal server error")
			return
		}

		slog.Info("Successfully added song to the database.", slog.Any("request", req))

		song := requests.GetSongRequest{
			Group: req.Group,
			Song:  req.Song,
		}

		songID, err := db.GetSongId(c, song)

		if err != nil {
			slog.Error("Failed to get song from the database.", slog.Any("error", err))
		}

		paginatedText := utils.PaginateText(req.Text)

		addVersesRequest := requests.AddVersesRequest{
			SongID: songID,
			Verses: paginatedText,
		}

		err = db.AddVerses(c, addVersesRequest)

		if err != nil {
			slog.Error("Failed to add verses to the database.", slog.Any("error", err))
		}

		slog.Info("Successfully added verses to the database.", slog.Any("verses", paginatedText))

		info, err := api.GetInfo(req.Group, req.Song)

		if err != nil {
			slog.Any("Failed to get info from the External API.", slog.Any("error", err))
			c.String(500, "Internal server error")
			return
		}

		c.IndentedJSON(200, info)
	}
}
