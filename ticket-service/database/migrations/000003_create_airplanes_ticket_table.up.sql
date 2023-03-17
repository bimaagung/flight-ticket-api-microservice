CREATE TABLE IF NOT EXISTS airplanes (
    id              VARCHAR(36) NOT NULL PRIMARY KEY,
    flight_code     VARCHAR(255) NOT NULL,
    seats           INTEGER NOT NULL
);