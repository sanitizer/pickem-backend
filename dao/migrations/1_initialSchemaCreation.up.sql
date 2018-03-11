START TRANSACTION;

--done
CREATE TABLE IF NOT EXISTS app_user (
  id          INT          NOT NULL AUTO_INCREMENT,
  battleNetId VARCHAR(100),
  discordId   VARCHAR(100),
  displayName VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  email       VARCHAR(100) NOT NULL,
  gravatar    LONGTEXT,
  teamLogo    LONGTEXT,
  lastActive  TIMESTAMP    NULL DEFAULT NULL,
  PRIMARY KEY (id)
);

--done
CREATE TABLE IF NOT EXISTS team (
  id            INT NOT NULL AUTO_INCREMENT,
  name          VARCHAR(100),
  shortName     VARCHAR(3),
  logo          LONGTEXT,
  PRIMARY KEY (id)
);

--done
CREATE TABLE IF NOT EXISTS player (
  id        INT NOT NULL AUTO_INCREMENT,
  name      VARCHAR(100),
  fullName  VARCHAR(100),
  PRIMARY KEY (id),
  CONSTRAINT UC_NFN UNIQUE (name, fullName)
);

--done
CREATE TABLE IF NOT EXISTS competition (
  id          INT          NOT NULL AUTO_INCREMENT,
  winstonsId  INT          NOT NULL,
  name        VARCHAR(255) NOT NULL,
  type        VARCHAR(255)          DEFAULT 'Standard',
  organizer   VARCHAR(255),
  description LONGTEXT,
  start       DATE,
  end         DATE,
  imageRef    LONGTEXT,
  isActive    BOOLEAN               DEFAULT FALSE,
  PRIMARY KEY (id)
);

--done
CREATE TABLE IF NOT EXISTS league (
  id            INT          NOT NULL AUTO_INCREMENT,
  name          VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  description   LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  ownerId       INT          NOT NULL,
  competitionId INT          NOT NULL,
  isLocked      BOOLEAN               DEFAULT FALSE,
  isPublic      BOOLEAN               DEFAULT FALSE,
  maxUsers      INT                   DEFAULT 0,
  simpleMode    BOOLEAN               DEFAULT TRUE,
  PRIMARY KEY (id),
  FOREIGN KEY (competitionId) REFERENCES competition (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  FOREIGN KEY (ownerId) REFERENCES app_user (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  CONSTRAINT UC_ID_OWNERID_COMPETEID UNIQUE (id, competitionId, ownerId)
);

--done
CREATE TABLE IF NOT EXISTS leaderboard (
  id       INT          NOT NULL AUTO_INCREMENT,
  type     VARCHAR(255) NOT NULL,
  leagueId INT          NOT NULL,
  isActive BOOLEAN               DEFAULT FALSE,
  stage    INT,
  PRIMARY KEY (id),
  FOREIGN KEY (leagueId) REFERENCES league (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS game_map (
  id   INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(100),
  type VARCHAR(100),
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS game_match (
  id            BIGINT(20) NOT NULL AUTO_INCREMENT,
  competitionId INT        NOT NULL,
  awayTeamId    INT        NOT NULL,
  homeTeamId    INT        NOT NULL,
  isLocked      BOOLEAN             DEFAULT FALSE,
  start         DATE,
  stage         INT,
  week          INT,
  PRIMARY KEY (id),
  FOREIGN KEY (competitionId) REFERENCES competition (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  FOREIGN KEY (awayTeamId) REFERENCES team (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  FOREIGN KEY (homeTeamId) REFERENCES team (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT
);

-- compound tables
--done
CREATE TABLE IF NOT EXISTS league_user (
  id       INT NOT NULL AUTO_INCREMENT,
  userId   INT NOT NULL,
  leagueId INT NOT NULL,
  isActive BOOLEAN      DEFAULT FALSE,
  PRIMARY KEY (id),
  FOREIGN KEY (leagueId) REFERENCES league (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  FOREIGN KEY (userId) REFERENCES app_user (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS leaderboard_leagueuser_points (
  leaderboardId INT NOT NULL,
  leagueUserId  INT NOT NULL,
  points        INT DEFAULT 0,
  PRIMARY KEY (leaderboardId, leagueUserId),
  FOREIGN KEY (leagueUserId) REFERENCES league_user (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  FOREIGN KEY (leaderboardId) REFERENCES leaderboard (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS team_competition (
  teamId        INT NOT NULL,
  competitionId INT NOT NULL,
  PRIMARY KEY (teamId, competitionId),
  FOREIGN KEY (teamId) REFERENCES team (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  FOREIGN KEY (competitionId) REFERENCES competition (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- done
CREATE TABLE IF NOT EXISTS team_player_role (
  teamPlayerId   INT         NOT NULL,
  role     VARCHAR(50) NOT NULL,
  FOREIGN KEY (teamPlayerId) REFERENCES team_player (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT UC_TEAM_PLAYER_ROLE UNIQUE (teamPlayerId, role)
);

--done
CREATE TABLE IF NOT EXISTS team_player (
  id       INT         NOT NULL AUTO_INCREMENT,
  teamId   INT NOT NULL,
  playerId INT NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT UC_TEAM_PLAYER UNIQUE (teamId, playerId),
  FOREIGN KEY (teamId) REFERENCES team (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  FOREIGN KEY (playerId) REFERENCES player (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS match_map_score (
  matchId   BIGINT(20) NOT NULL,
  mapId     INT        NOT NULL,
  homeScore INT DEFAULT 0,
  awayScore INT DEFAULT 0,
  PRIMARY KEY (matchId, mapId),
  FOREIGN KEY (matchId) REFERENCES game_match (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  FOREIGN KEY (mapId) REFERENCES game_map (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS match_stat (
  matchId   BIGINT(20) NOT NULL,
  homeScore INT DEFAULT 0,
  awayScore INT DEFAULT 0,
  FOREIGN KEY (matchId) REFERENCES game_match (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS match_pick (
  id           BIGINT(20) NOT NULL AUTO_INCREMENT,
  matchId      BIGINT(20) NOT NULL,
  leagueUserId INT        NOT NULL,
  winnerTeamId INT        NOT NULL,
  loserTeamId  INT        NOT NULL,
  homeScore    INT                 DEFAULT 0,
  awayScore    INT                 DEFAULT 0,
  scorePoints  INT                 DEFAULT 0,
  winnerPoints INT                 DEFAULT 0,
  PRIMARY KEY (id),
  FOREIGN KEY (matchId) REFERENCES game_match (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  FOREIGN KEY (leagueUserId) REFERENCES league_user (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  FOREIGN KEY (winnerTeamId) REFERENCES team (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  FOREIGN KEY (loserTeamId) REFERENCES team (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS match_map_pick (
  matchPickId  BIGINT(20) NOT NULL,
  mapId        INT        NOT NULL,
  winnerTeamId INT,
  points       INT DEFAULT 0,
  PRIMARY KEY (matchPickId, mapId),
  FOREIGN KEY (matchPickId) REFERENCES match_pick (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  FOREIGN KEY (mapId) REFERENCES game_map (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT
);

COMMIT;