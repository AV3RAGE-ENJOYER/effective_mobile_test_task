package handlers

import (
	"log/slog"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/requests"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository"
	"github.com/gin-gonic/gin"
)

// @Summary Edit song
// @Description Edit song in the database
// @Tags API
// @Accept json
// @Produce json
// @Param body body requests.EditSongRequest true "Request body"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /edit [put]
func EditSongHandler(db repository.DatabaseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requests.EditSongRequest

		err := c.BindJSON(&req)

		if err != nil {
			slog.Error("Failed to unmarshall json.", slog.Any("error", err))
			c.String(400, "Bad request")
			return
		}

		slog.Debug("Edit request.", slog.Any("request", req))

		songDetail, err := db.GetSong(c, requests.GetSongRequest{
			Group: req.Group,
			Song:  req.Song,
		})

		if err != nil {
			slog.Error("Failed to get song details.", slog.Any("error", err))
			c.String(500, "Internal server error")
			return
		}

		defaultValues := map[string]string{
			"release_date": songDetail.ReleaseDate.Format("YYYY-MM-DD"),
			"text":         songDetail.Text,
			"link":         songDetail.Link,
		}

		if req.ReleaseDate == "" {
			req.ReleaseDate = defaultValues["release_date"]
		} else if req.Text == "" {
			req.Text = defaultValues["text"]
		} else if req.Link == "" {
			req.Link = defaultValues["link"]
		}

		err = db.EditSong(c, req)

		if err != nil {
			slog.Error("Failed to edit song.", slog.Any("error", err))
			c.String(500, "Internal server error")
			return
		}

		c.String(200, "OK")
	}
}
