package requests

type GetInfoRequest struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
	Offset      int    `json:"offset"`
}

type GetSongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type GetVerseRequest struct {
	Group  string `json:"group"`
	Song   string `json:"song"`
	Offset int    `json:"offset"`
}

type AddSongRequest struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type EditSongRequest struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type AddVersesRequest struct {
	SongID int            `json:"song_id"`
	Verses map[int]string `json:"verses"`
}

type DeleteSongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}
