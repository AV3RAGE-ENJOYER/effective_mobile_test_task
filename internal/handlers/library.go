package handlers

import (
	"log/slog"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/requests"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/utils"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Get library info
// @Schemes
// @Description Get library info from the database
// @Tags API
// @Accept json
// @Produce json
// @Param group query string false "Group Name"
// @Param song query string false "Song Name"
// @Param release_date query string false "Song Release Date"
// @Param text query string false "Song Text"
// @Param link query string false "Song Link"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} []models.SongsInfo
// @Failure 500 {string} string "Internal server error"
// @Router /library [get]
func GetLibraryInfoHandler(db repository.DatabaseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		defaultValues := map[string]string{
			"group":        "",
			"song":         "",
			"release_date": "",
			"text":         "",
			"link":         "",
			"offset":       "0",
		}

		query := c.Request.URL.Query()

		proccessedQuery := utils.ProccessQuery(query, defaultValues)

		slog.Debug("Proccessed Query.", slog.Any("query", proccessedQuery))

		req := requests.GetInfoRequest{
			Group:       proccessedQuery["group"],
			Song:        proccessedQuery["song"],
			ReleaseDate: proccessedQuery["release_date"],
			Text:        proccessedQuery["text"],
			Link:        proccessedQuery["link"],
			Offset:      utils.ConvertToInt(proccessedQuery["offset"]),
		}

		info, err := db.GetInfo(c, req)

		if err != nil {
			slog.Error("Failed to get info.", slog.Any("error", err))
			c.String(500, "Internal server error")
			return
		}

		slog.Info("Successfully got Library Info.", slog.Any("items", len(info)))

		c.IndentedJSON(200, info)
	}
}
