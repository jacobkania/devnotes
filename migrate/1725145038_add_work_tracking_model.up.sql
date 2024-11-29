CREATE TABLE IF NOT EXISTS work_tracking (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	start_time TIMESTAMP,
	end_time TIMESTAMP,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
