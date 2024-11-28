package handlers

import (
	"log/slog"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/requests"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/utils"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// @BasePath /api/v1

// @Summary Get song text
// @Schemes
// @Description Get song text from the database
// @Tags API
// @Accept json
// @Produce json
// @Param group query string true "Group name"
// @Param song query string true "Song name"
// @Param offset query int true "Offset for pagination"
// @Success 200 {string} string "Text of the song"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Song not found"
// @Failure 500 {string} string "Internal server error"
// @Router /text [get]
func GetTextHandler(db repository.DatabaseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()

		slog.Debug("Get text request.", slog.Any("query", query))

		if query["group"] != nil && query["song"] != nil && query["offset"] != nil {
			req := requests.GetVerseRequest{
				Group:  query["group"][0],
				Song:   query["song"][0],
				Offset: utils.ConvertToInt(query["offset"][0]),
			}

			verse, err := db.GetVerse(c, req)

			if err == pgx.ErrNoRows {
				c.String(404, "Song not found")
				return
			} else if err != nil {
				c.String(500, "Internal server error")
				return
			}

			slog.Info("Successfuly got verse from the database", slog.Any("request", req), slog.Any("verse", verse))

			c.String(200, verse)
			return
		}

		c.String(400, "Bad request")
	}
}
