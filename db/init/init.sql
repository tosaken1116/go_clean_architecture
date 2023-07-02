CREATE TYPE "content_type" AS ENUM (
  'image',
  'video',
  'gif'
);

CREATE TYPE "following_state" AS ENUM (
  'approved',
  'unapproved'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "screen_id" varchar(14) UNIQUE,
  "password_hash" varchar,
  "user_name" varchar,
  "email" varchar UNIQUE NOT NULL,
  "is_public" bool DEFAULT true,
  "icon_url" varchar,
  "header_url" varchar
);

CREATE TABLE "tweets" (
  "id" uuid PRIMARY KEY,
  "description" varchar(200),
  "tweet_user" uuid,
  "activity" int DEFAULT 0,
  "created_at" timestamp,
  "deleted_at" timestamp,
  "reply_to" uuid DEFAULT null
);

CREATE TABLE "lists" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "created_by" uuid,
  "created_at" timestamp,
  "deleted_at" timestamp,
  "is_public" bool DEFAULT true
);

CREATE TABLE "messages" (
  "id" uuid PRIMARY KEY,
  "description" varchar,
  "sender_id" uuid,
  "recipient_id" uuid,
  "created_at" timestamp
);

CREATE TABLE "media" (
  "id" uuid,
  "tweet" uuid,
  "type" content_type,
  PRIMARY KEY ("id", "tweet")
);

CREATE TABLE "like" (
  "tweet_id" uuid,
  "user_id" uuid,
  "created_at" timestamp,
  PRIMARY KEY ("tweet_id", "user_id")
);

CREATE TABLE "retweet" (
  "tweet_id" uuid,
  "user_id" uuid,
  PRIMARY KEY ("tweet_id", "user_id")
);

CREATE TABLE "list_user" (
  "list_id" uuid,
  "user_id" uuid,
  PRIMARY KEY ("list_id", "user_id")
);

CREATE TABLE "bookmark" (
  "tweet_id" uuid,
  "user_id" uuid,
  "created_at" timestamp,
  PRIMARY KEY ("tweet_id", "user_id")
);

CREATE TABLE "following" (
  "user_id" uuid,
  "target_id" uuid,
  "state" following_state,
  "created_at" timestamp,
  PRIMARY KEY ("user_id", "target_id")
);

CREATE TABLE "muting" (
  "user_id" uuid,
  "target_id" uuid,
  PRIMARY KEY ("user_id", "target_id")
);

ALTER TABLE "tweets" ADD FOREIGN KEY ("tweet_user") REFERENCES "users" ("id");

ALTER TABLE "tweets" ADD FOREIGN KEY ("reply_to") REFERENCES "tweets" ("id");

ALTER TABLE "lists" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("recipient_id") REFERENCES "users" ("id");

ALTER TABLE "media" ADD FOREIGN KEY ("tweet") REFERENCES "tweets" ("id");

ALTER TABLE "like" ADD FOREIGN KEY ("tweet_id") REFERENCES "tweets" ("id");

ALTER TABLE "like" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "retweet" ADD FOREIGN KEY ("tweet_id") REFERENCES "tweets" ("id");

ALTER TABLE "retweet" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "list_user" ADD FOREIGN KEY ("list_id") REFERENCES "lists" ("id");

ALTER TABLE "list_user" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "bookmark" ADD FOREIGN KEY ("tweet_id") REFERENCES "tweets" ("id");

ALTER TABLE "bookmark" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "following" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "following" ADD FOREIGN KEY ("target_id") REFERENCES "users" ("id");

ALTER TABLE "muting" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "muting" ADD FOREIGN KEY ("target_id") REFERENCES "users" ("id");
