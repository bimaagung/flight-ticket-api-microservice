CREATE TABLE IF NOT EXISTS airplanes (
    "id"              UUID NOT NULL PRIMARY KEY,
    "flight_code"     VARCHAR(255) NOT NULL,
    "seats"           INTEGER NOT NULL,
    "created_at"      TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"      TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"      TIMESTAMP(0) WITH TIME zone
);