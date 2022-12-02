CREATE TABLE IF NOT EXISTS catalog(
                                      id TEXT PRIMARY KEY,
                                      name TEXT UNIQUE NOT NULL,
                                      count INT NOT NULL,
                                      price INT NOT NULL
);