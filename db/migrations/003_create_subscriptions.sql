CREATE TABLE subscriptions (
  user_id BIGINT NOT NULL REFERENCES users(id),
  flight_id BIGINT NOT NULL REFERENCES flights(id),
  active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY (user_id, flight_id)
);
