CREATE TABLE IF NOT EXISTS url (
  id       BIGINT PRIMARY KEY,
  shortUrl VARCHAR(11) NOT NULL,
  longUrl  TEXT NOT NULL
);
