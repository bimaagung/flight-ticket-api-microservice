CREATE TABLE IF NOT EXISTS tickets (
    "id" UUID NOT NULL PRIMARY KEY,
    "track_id" UUID NOT NULL,
    "airplane_id" UUID NOT NULL,
    "datetime" TIMESTAMP(0) WITH TIME zone NOT NULL,
    "price" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(0) WITH TIME zone,

    CONSTRAINT fk_track_id FOREIGN KEY (track_id) REFERENCES tracks (id) ON DELETE CASCADE,
    CONSTRAINT fk_airplane_id FOREIGN KEY (airplane_id) REFERENCES airplanes (id) ON DELETE CASCADE
);