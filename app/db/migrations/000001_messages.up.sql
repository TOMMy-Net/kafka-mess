CREATE TABLE IF NOT EXISTS messages (
	uid UUID NOT NULL UNIQUE,
	message TEXT NOT NULL,
	status INTEGER NOT NULL DEFAULT 0 CHECK(status IN (0, 1)),
	PRIMARY KEY("uid")
);
