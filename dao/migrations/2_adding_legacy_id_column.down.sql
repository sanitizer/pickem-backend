START TRANSACTION;

ALTER TABLE app_user
  drop COLUMN
  legacyId;

ALTER TABLE competition
  drop COLUMN
  legacyId;

ALTER TABLE league
  drop COLUMN
  legacyId;

ALTER TABLE team
  drop COLUMN
  legacyId;

ALTER TABLE game_match
  drop COLUMN
  legacyId;

ALTER TABLE game_map
  DROP INDEX UC_MAP_NAME;

COMMIT;