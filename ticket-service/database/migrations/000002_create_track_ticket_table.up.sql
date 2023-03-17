CREATE TABLE IF NOT EXISTS tracks (
    "id" UUID NOT NULL PRIMARY KEY,
    "arrival" VARCHAR(255) NOT NULL,
    "departure" VARCHAR(255) NOT NULL,
    "long_flight" INTEGER NOT NULL
);