CREATE TABLE flights (
    flight_id        TEXT PRIMARY KEY,
    flight_number    TEXT NOT NULL,        -- e.g. AA123
    status           TEXT,                 -- scheduled, active, landed, cancelled
    departure_gate   TEXT,
    departure_terminal TEXT,
    delay_minutes    INT DEFAULT 0,
    scheduled_dep    TIMESTAMPTZ,
    estimated_dep    TIMESTAMPTZ,
    last_updated     TIMESTAMPTZ DEFAULT NOW()
);
