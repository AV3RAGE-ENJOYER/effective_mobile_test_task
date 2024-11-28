package handlers

import (
	"log/slog"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/requests"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// @BasePath /api/v1

// @Summary Get song info
// @Schemes
// @Description Get song info from the database
// @Tags API
// @Accept json
// @Produce json
// @Param group query string true "Group Name"
// @Param song query string true "Song Name"
// @Success 200 {object} models.SongDetail
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Page not found"
// @Failure 500 {string} string "Internal server error"
// @Router /info [get]
func GetInfoHandler(db repository.DatabaseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()

		slog.Debug("Get info request.", slog.Any("query", query))

		if query["group"] != nil && query["song"] != nil {
			songDetail, err := db.GetSong(c, requests.GetSongRequest{
				Group: query["group"][0],
				Song:  query["song"][0],
			})

			if err == pgx.ErrNoRows {
				slog.Error("No song found.", slog.Any("error", err))
				c.String(404, "Page not found")
				return
			} else if err != nil {
				slog.Error("Failed to get song from the database.", slog.Any("error", err))
				c.String(500, "Internal server error")
				return
			}

			c.IndentedJSON(200, songDetail)
			return
		}

		c.String(400, "Bad request")
	}
}
