CREATE TABLE IF NOT EXISTS flights (
    flight_id TEXT PRIMARY KEY,
    flight_number TEXT NOT NULL,
    status TEXT,
    departure_gate TEXT,
    departure_terminal TEXT,
    delay_minutes INT NOT NULL DEFAULT 0,
    scheduled_dep TIMESTAMPTZ,
    estimated_dep TIMESTAMPTZ,
    last_updated TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
