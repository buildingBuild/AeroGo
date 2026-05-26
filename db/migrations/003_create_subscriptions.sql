CREATE TABLE IF NOT EXISTS subscriptions (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    flight_id TEXT NOT NULL REFERENCES flights(flight_id) ON DELETE CASCADE,
    active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, flight_id)
);
