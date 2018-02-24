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

ALTER TABLE app_user
  DROP INDEX UC_LIDAU;

ALTER TABLE competition
  DROP INDEX UC_LIDC;

ALTER TABLE league
  DROP INDEX UC_LIDL;

ALTER TABLE team
  DROP INDEX UC_LIDT;

ALTER TABLE game_match
  DROP INDEX UC_LIDGM;

COMMIT;