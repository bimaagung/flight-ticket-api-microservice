CREATE TABLE IF NOT EXISTS Airplanes (
    "id"                UUID NOT NULL PRIMARY KEY,
    "flight_code"       VARCHAR(255) NOT NULL,
    "seats"             INTEGER NOT NULL,
    "type"              VARCHAR(10) NOT NULL,
    "production_year"   DATE NOT NULL,
    "factory"           VARCHAR(112) NOT NULL,
    "created_at"        TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"        TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"        BOOLEAN NOT NULL DEFAULT false
);