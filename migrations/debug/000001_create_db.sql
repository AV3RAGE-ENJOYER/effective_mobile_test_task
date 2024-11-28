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

-- INSERT TEST DATA

INSERT INTO public.songs("group", "song") VALUES('Eminem', 'Stan');
INSERT INTO public.songs("group", "song") VALUES('Schoolboy Q', 'Collard Greens');
INSERT INTO public.songs("group", "song") VALUES('Kendrick Lamar', 'Swimming Pools');

INSERT INTO public.song_details("song_id", "release_date", "text", "link") 
    VALUES(1, '2000-11-20', 
    'My tea''s gone cold, I''m wondering why I\nGot out of bed at all\nThe morning rain clouds up my window (Window)\nAnd I can''t see at all\nAnd even if I could, it''d all be grey\nBut your picture on my wall\nIt reminds me that it''s not so bad, it''s not so bad (Bad)\n\nDear Slim, I wrote you, but you still ain''t callin\nI left my cell, my pager, and my home phone at the bottom\nI sent two letters back in autumn, you must not''ve got ''em\nThere prob''ly was a problem at the post office or somethin''\nSometimes I scribble addresses too sloppy when I jot ''em\nBut anyways, fuck it, what''s been up, man? How''s your daughter?',
    'https://www.youtube.com/watch?v=gOMhN-hfMtY');
INSERT INTO public.verses("song_id", "verse_num", "verse")
    VALUES(1, 1,
    'My tea''s gone cold, I''m wondering why I\nGot out of bed at all\nThe morning rain clouds up my window (Window)\nAnd I can''t see at all\nAnd even if I could, it''d all be grey\nBut your picture on my wall\nIt reminds me that it''s not so bad, it''s not so bad (Bad)');
INSERT INTO public.verses("song_id", "verse_num", "verse")
    VALUES(1, 2,
    'Dear Slim, I wrote you, but you still ain''t callin''\nI left my cell, my pager, and my home phone at the bottom\nI sent two letters back in autumn, you must not''ve got ''em\nThere prob''ly was a problem at the post office or somethin''\nSometimes I scribble addresses too sloppy when I jot ''em\nBut anyways, fuck it, what''s been up, man? How''s your daughter?');

INSERT INTO public.song_details("song_id", "release_date", "text", "link") 
    VALUES(2, '2013-06-13', 
    'Yeah, yeah, yeah\nOh, oh\nYo, yo\n\nOh, oh, luxury\nChidi-ching-ching could buy anything, cop that\nOh, oh, collard greens\nThree degrees low, make it hot for me, drop that\nOh, oh, down with that shit\nDrink this, smoke this, get down with the shit, aye\nOh, oh, down with the shit\nCop this, pop this, down with the shit\n\nSmoke this, drink this, straight to my liver\nWatch this, no tick, yeah, I''m the nigga\nGang rap, X-mas, smoke, shots I deliver\nFaded, Vegas, might sponsor the killer\nShake it, break it, hot-hot for the winter\nDrop it, cop it, eyes locked on your inner object\nRock it, blast-blast, new beginnings\nLovely, pinky how not I remember, fiendin''',
    'https://www.youtube.com/watch?v=_L2vJEb6lVE');

INSERT INTO public.verses("song_id", "verse_num", "verse")
    VALUES(2, 1,
    'Yeah, yeah, yeah\nOh, oh\nYo, yo');
INSERT INTO public.verses("song_id", "verse_num", "verse")
    VALUES(2, 2,
    'Oh, oh, luxury\nChidi-ching-ching could buy anything, cop that\nOh, oh, collard greens\nThree degrees low, make it hot for me, drop that\nOh, oh, down with that shit\nDrink this, smoke this, get down with the shit, aye\nOh, oh, down with the shit\nCop this, pop this, down with the shit');
INSERT INTO public.verses("song_id", "verse_num", "verse")
    VALUES(2, 3,
    'Smoke this, drink this, straight to my liver\nWatch this, no tick, yeah, I''m the nigga\nGang rap, X-mas, smoke, shots I deliver\nFaded, Vegas, might sponsor the killer\nShake it, break it, hot-hot for the winter\nDrop it, cop it, eyes locked on your inner object\nRock it, blast-blast, new beginnings\nLovely, pinky how not I remember, fiendin''');

INSERT INTO public.song_details("song_id", "release_date", "text", "link") 
    VALUES(3, '31-07-2012', 
    'Pour up (drank), head shot (drank)\nSit down (drank), stand up (drank)\nPass out (drank), wake up (drank)\nFaded (drank), faded (drank)\n\nNow I done grew up ''round some people livin'' their life in bottles\nGranddaddy had the golden flask\nBackstroke every day in Chicago\nSome people like the way it feels\nSome people wanna kill their sorrows\nSome people wanna fit in with the popular, that was my problem\n\nI was in the dark room, loud tunes, lookin'' to make a vow soon\nThat I''ma get fucked up, fillin'' up my cup, I see the crowd mood\nChangin'' by the minute and the record on repeat\nTook a sip, then another sip, then somebody said to me',
    'https://www.youtube.com/watch?v=B5YNiCfWC3A');
INSERT INTO public.verses("song_id", "verse_num", "verse")
    VALUES(3, 1,
    'Pour up (drank), head shot (drank)\nSit down (drank), stand up (drank)\nPass out (drank), wake up (drank)\nFaded (drank), faded (drank)');
INSERT INTO public.verses("song_id", "verse_num", "verse")
    VALUES(3, 2,
    'Now I done grew up ''round some people livin'' their life in bottles\nGranddaddy had the golden flask\nBackstroke every day in Chicago\nSome people like the way it feels\nSome people wanna kill their sorrows\nSome people wanna fit in with the popular, that was my problem');
INSERT INTO public.verses("song_id", "verse_num", "verse")
    VALUES(3, 3,
    'I was in the dark room, loud tunes, lookin'' to make a vow soon\nThat I''ma get fucked up, fillin'' up my cup, I see the crowd mood\nChangin'' by the minute and the record on repeat\nTook a sip, then another sip, then somebody said to me');

-- +goose Down
DROP TABLE public.song_details;
DROP TABLE public.verses;
DROP TABLE public.songs;