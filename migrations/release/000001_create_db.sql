-- +goose Up
CREATE TABLE IF NOT EXISTS public.songs (
    "id" BIGINT PRIMARY KEY,
    "group" TEXT NOT NULL,
    "song" TEXT NOT NULL
);

CREATE SEQUENCE public.songs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.songs_id_seq OWNED BY public.songs.id;

ALTER TABLE ONLY public.songs ALTER COLUMN id SET DEFAULT nextval('public.songs_id_seq'::regclass);

CREATE TABLE IF NOT EXISTS public.song_details (
    "id" BIGINT PRIMARY KEY,
    "song_id" BIGINT,
    "release_date" DATE,
    "text" TEXT,
    "link" TEXT,

    FOREIGN KEY("song_id") REFERENCES public.songs("id")
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE SEQUENCE public.song_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.song_details_id_seq OWNED BY public.song_details.id;

ALTER TABLE ONLY public.song_details ALTER COLUMN id SET DEFAULT nextval('public.song_details_id_seq'::regclass);

CREATE TABLE IF NOT EXISTS public.verses (
    "id" BIGINT PRIMARY KEY,
    "song_id" BIGINT,
    "verse_num" BIGINT,
    "verse" TEXT,

    FOREIGN KEY("song_id") REFERENCES public.songs("id")
        ON UPDATE CASCADE
        ON DELETE CASCADE 
);

CREATE SEQUENCE public.verses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.verses_id_seq OWNED BY public.verses.id;

ALTER TABLE ONLY public.verses ALTER COLUMN id SET DEFAULT nextval('public.verses_id_seq'::regclass);

-- +goose Down
DROP TABLE public.song_details;
DROP TABLE public.verses;
DROP TABLE public.songs;