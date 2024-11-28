package handlers

import (
	"log/slog"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/requests"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete song
// @Schemes
// @Description Delete song from the database
// @Tags API
// @Accept json
// @Produce json
// @Param body body requests.DeleteSongRequest true "Request body"
// @Success 200 {string} OK
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /delete [delete]
func DeleteSongHandler(db repository.DatabaseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requests.DeleteSongRequest

		err := c.BindJSON(&req)

		if err != nil {
			slog.Error("Failed to unmarshall json.", slog.Any("error", err))
			c.String(400, "Bad request")
			return
		}

		slog.Debug("Delete request.", slog.Any("request", req))

		err = db.DeleteSong(c, req)

		if err != nil {
			slog.Error("Failed to delete song.", slog.Any("error", err), slog.Any("request", req))
			c.String(500, "Internal server error")
			return
		}

		slog.Info("Successfully deleted song from the database.", slog.Any("request", req))
		c.String(200, "OK")
	}
}
