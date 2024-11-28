package repository

import (
	"context"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/requests"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository/models"
)

type DatabaseRepository interface {
	GetInfo(ctx context.Context, req requests.GetInfoRequest) ([]models.SongsInfo, error)
	GetSongId(ctx context.Context, req requests.GetSongRequest) (int, error)
	GetSong(ctx context.Context, req requests.GetSongRequest) (models.SongDetail, error)
	GetSongText(ctx context.Context, req requests.GetSongRequest) (string, error)
	GetVerse(ctx context.Context, req requests.GetVerseRequest) (string, error)
	AddSong(ctx context.Context, req requests.AddSongRequest) error
	EditSong(ctx context.Context, req requests.EditSongRequest) error
	AddVerses(ctx context.Context, req requests.AddVersesRequest) error
	DeleteSong(ctx context.Context, req requests.DeleteSongRequest) error
}
