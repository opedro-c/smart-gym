CREATE TABLE [IF NOT EXISTS] exercises(
    id INTEGER PRIMARY KEY,
    user_rf_id TEXT NOT NULL,
    started_at TEXT NOT NULL,
    finished_at TEXT NOT NULL,
    name TEXT NOT NULL,
);

CREATE TABLE [IF NOT EXISTS] series(
    id INTEGER PRIMARY KEY,
    started_at TEXT NOT NULL,
    finished_at TEXT NOT NULL,
    repetitions INTEGER NOT NULL,
    weight REAL NOT NULL,
    exercise_id INTEGER NOT NULL,
);

CREATE TABLE [IF NOT EXISTS] valid_rf_ids (
    rf_id TEXT PRIMARY KEY
);
