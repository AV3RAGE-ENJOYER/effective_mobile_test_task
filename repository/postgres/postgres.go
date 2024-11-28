package postgres

import (
	"context"
	"log/slog"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/requests"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/utils"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	Pool *pgxpool.Pool
}

var _ repository.DatabaseRepository = &PostgresDB{}

func NewPostgresDB(ctx context.Context, DB_URL string) (*PostgresDB, error) {
	pool, err := pgxpool.New(ctx, DB_URL)

	if err != nil {
		return &PostgresDB{}, err
	}

	postgres := &PostgresDB{pool}
	return postgres, nil
}

func (db *PostgresDB) GetInfo(ctx context.Context, req requests.GetInfoRequest) ([]models.SongsInfo, error) {
	args := pgx.NamedArgs{
		"group":        utils.AddPercentageToString(req.Group),
		"song":         utils.AddPercentageToString(req.Song),
		"release_date": utils.AddPercentageToString(req.ReleaseDate),
		"text":         utils.AddPercentageToString(req.Text),
		"link":         utils.AddPercentageToString(req.Link),
	}

	q := `SELECT 
		s.group, s.song, 
		sd.release_date, sd.link
		FROM songs AS s 
		JOIN song_details AS sd 
		ON s.id = sd.song_id
		WHERE s.group ILIKE @group
		AND s.song ILIKE @song
		AND TO_CHAR(sd.release_date, 'YYYY-MM-DD') ILIKE @release_date
		AND sd.text ILIKE @text
		AND sd.link ILIKE @link`

	rows, err := db.Pool.Query(ctx, q, args)

	if err != nil {
		slog.Error("Failed to get info.", slog.Any("error", err))
		return []models.SongsInfo{}, err
	}

	var info []models.SongsInfo
	var currentRow models.SongsInfo

	for rows.Next() {
		err := rows.Scan(
			&currentRow.Group,
			&currentRow.Song,
			&currentRow.ReleaseDate,
			&currentRow.Link,
		)

		currentRow.Offset = req.Offset

		if err != nil {
			slog.Error("Failed to scan row.", slog.Any("error", err))
			return []models.SongsInfo{}, err
		}

		text, err := db.GetVerse(
			ctx,
			requests.GetVerseRequest{
				Group:  currentRow.Group,
				Song:   currentRow.Song,
				Offset: req.Offset,
			},
		)

		if err != nil {
			slog.Error(
				"Failed to get text.",
				slog.Any("error", err),
				slog.Any("offset", req.Offset),
			)
			return []models.SongsInfo{}, err
		}

		currentRow.Text = text

		info = append(info, currentRow)
	}

	return info, nil
}

func (db *PostgresDB) GetSongId(ctx context.Context, req requests.GetSongRequest) (int, error) {
	args := pgx.NamedArgs{
		"group": req.Group,
		"song":  req.Song,
	}

	q := `SELECT id FROM songs 
		WHERE "group" = @group
		AND "song" = @song`

	row := db.Pool.QueryRow(ctx, q, args)

	var id int
	err := row.Scan(&id)

	if err != nil {
		slog.Error("Failed to scan row.", slog.Any("error", err))
		return 0, err
	}

	return id, nil
}

func (db *PostgresDB) GetSong(ctx context.Context, req requests.GetSongRequest) (models.SongDetail, error) {
	args := pgx.NamedArgs{
		"group": req.Group,
		"song":  req.Song,
	}

	q := `SELECT sd.release_date, sd.text, sd.link FROM songs AS s
		JOIN song_details AS sd ON s.id=sd.song_id 
		WHERE s.group = @group
		AND s.song = @song`

	row := db.Pool.QueryRow(ctx, q, args)

	var songDetail models.SongDetail
	err := row.Scan(
		&songDetail.ReleaseDate,
		&songDetail.Text,
		&songDetail.Link,
	)

	if err != nil {
		slog.Error("Failed to scan row.", slog.Any("error", err))
		return models.SongDetail{}, err
	}

	return songDetail, nil
}

func (db *PostgresDB) GetSongText(ctx context.Context, req requests.GetSongRequest) (string, error) {
	args := pgx.NamedArgs{
		"group": req.Group,
		"song":  req.Song,
	}

	q := `SELECT sd.text FROM song_details AS sd
		JOIN songs AS s ON s.id=sd.song_id 
		WHERE s.group = @group
		AND s.song = @song`

	row := db.Pool.QueryRow(ctx, q, args)

	var text string
	err := row.Scan(&text)

	if err != nil {
		slog.Error("Failed to scan row.", slog.Any("error", err))
		return "", err
	}

	return text, nil
}

func (db *PostgresDB) AddSong(ctx context.Context, req requests.AddSongRequest) error {
	tx, err := db.Pool.Begin(ctx)

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	args := pgx.NamedArgs{
		"group":        req.Group,
		"song":         req.Song,
		"release_date": req.ReleaseDate,
		"text":         req.Text,
		"link":         req.Link,
	}

	q := `WITH add_user AS (
		INSERT INTO 
		songs("group", "song")
		VALUES (@group, @song)
		RETURNING id)

		INSERT INTO 
		song_details("song_id", "release_date", "text", "link")
		SELECT id, @release_date, @text, @link FROM add_user`

	_, err = tx.Exec(
		ctx,
		q,
		args,
	)

	return err
}

func (db *PostgresDB) EditSong(ctx context.Context, req requests.EditSongRequest) error {
	tx, err := db.Pool.Begin(ctx)

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	args := pgx.NamedArgs{
		"group":        req.Group,
		"song":         req.Song,
		"release_date": req.ReleaseDate,
		"text":         req.Text,
		"link":         req.Link,
	}

	q := `UPDATE song_details AS sd 
		SET "release_date" = @release_date,
		"text" = @text,
		"link" = @link
		FROM songs AS s
		WHERE s.id = sd.song_id
		AND s.group = @group
		AND s.song = @song`

	_, err = tx.Exec(
		ctx,
		q,
		args,
	)

	if err != nil {
		slog.Error("Failed to edit song details.", slog.Any("error", err))

		return err
	}

	return nil
}

func (db *PostgresDB) AddVerses(ctx context.Context, req requests.AddVersesRequest) error {
	tx, err := db.Pool.Begin(ctx)

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	q := `INSERT INTO 
		verses("song_id", "verse_num", "verse")
		VALUES(@song_id, @verse_num, @verse)`

	for verseNum, verse := range req.Verses {
		args := pgx.NamedArgs{
			"song_id":   req.SongID,
			"verse_num": verseNum,
			"verse":     verse,
		}

		_, err := tx.Exec(
			ctx,
			q,
			args,
		)

		if err != nil {
			slog.Error(
				"Failed to add verse to the database.",
				slog.Any("error", err),
				slog.Any("verse_num", verseNum),
				slog.Any("verse", verse),
			)

			return err
		}

	}

	return nil
}

func (db *PostgresDB) GetVerse(ctx context.Context, req requests.GetVerseRequest) (string, error) {
	if req.Offset != 0 {
		args := pgx.NamedArgs{
			"group":  req.Group,
			"song":   req.Song,
			"offset": req.Offset,
		}

		q := `SELECT v.verse FROM verses AS v 
			JOIN songs AS s ON v.song_id = s.id 
			WHERE s.group = @group 
			AND s.song = @song 
			AND v.verse_num = @offset`

		row := db.Pool.QueryRow(ctx, q, args)

		var verse string

		err := row.Scan(&verse)

		if err == pgx.ErrNoRows {
			slog.Debug("No verse.", slog.Any("error", err), slog.Any("offset", req.Offset))
			return "", nil
		} else if err != nil {
			slog.Error("Failed to scan row.", slog.Any("error", err))
			return "", err
		}

		return verse, nil
	}

	fullTextReq := requests.GetSongRequest{
		Group: req.Group,
		Song:  req.Song,
	}

	text, err := db.GetSongText(ctx, fullTextReq)

	if err != nil {
		slog.Error("Failed to get song full text.", slog.Any("error", err))
		return "", err
	}

	return text, nil
}

func (db *PostgresDB) DeleteSong(ctx context.Context, req requests.DeleteSongRequest) error {
	tx, err := db.Pool.Begin(ctx)

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	args := pgx.NamedArgs{
		"group": req.Group,
		"song":  req.Song,
	}

	q := `DELETE FROM songs AS s
		WHERE s.group=@group AND s.song=@song`

	_, err = tx.Exec(
		ctx,
		q,
		args,
	)

	return err
}
