CREATE TABLE IF NOT EXISTS user(
                                      id TEXT PRIMARY KEY,
                                      name TEXT UNIQUE NOT NULL,
                                      position INT NOT NULL,
                                      password TEXT NOT NULL,
                                      refreshtoken TEXT
);