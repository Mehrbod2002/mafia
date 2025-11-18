-- Migration
-- 001_init.up.sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  phone VARCHAR(20) UNIQUE NOT NULL,
  otp VARCHAR(6),
  otp_expires TIMESTAMPTZ,
  role VARCHAR(20) DEFAULT 'user',
  status VARCHAR(20) DEFAULT 'active',
  created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE profiles (
  id SERIAL PRIMARY KEY,
  user_id INTEGER UNIQUE REFERENCES users(id),
  name VARCHAR(50),
  avatar TEXT,
  gender VARCHAR(10),
  age INTEGER,
  level INTEGER DEFAULT 1,
  wins INTEGER DEFAULT 0,
  losses INTEGER DEFAULT 0
);

CREATE TABLE wallets (
  id SERIAL PRIMARY KEY,
  user_id INTEGER UNIQUE REFERENCES users(id),
  coins INTEGER DEFAULT 100,
  diamonds INTEGER DEFAULT 10
);

CREATE TABLE game_rooms (
  id SERIAL PRIMARY KEY,
  code VARCHAR(10) UNIQUE NOT NULL,
  type VARCHAR(20) NOT NULL,
  host_id INTEGER REFERENCES users(id),
  status VARCHAR(20) DEFAULT 'waiting',
  phase VARCHAR(20) DEFAULT 'night',
  day_count INTEGER DEFAULT 0,
  winner VARCHAR(20),
  created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE room_players (
  game_room_id INTEGER REFERENCES game_rooms(id) ON DELETE CASCADE,
  user_id INTEGER REFERENCES users(id),
  PRIMARY KEY (game_room_id, user_id)
);

CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) UNIQUE NOT NULL,
  team VARCHAR(20),
  max_count INTEGER
);