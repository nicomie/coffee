CREATE TABLE coffees (
  id   BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  flavor VARCHAR(50)      NOT NULL,
  acidity  SMALLINT NOT NULL
);

