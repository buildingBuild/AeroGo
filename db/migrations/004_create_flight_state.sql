CREATE TABLE IF NOT EXISTS flight_state (
    flight_id TEXT PRIMARY KEY REFERENCES flights(flight_id) ON DELETE CASCADE,
    status TEXT,
    departure_gate TEXT,
    departure_terminal TEXT,
    delay_minutes INT NOT NULL DEFAULT 0,
    scheduled_dep TIMESTAMPTZ,
    estimated_dep TIMESTAMPTZ,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
