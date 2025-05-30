ALTER TABLE IF EXISTS list
ADD COLUMN IF NOT EXISTS position smallint DEFAULT 1 CONSTRAINT valid_position CHECK (
    position > 0
    AND position < 10
  );