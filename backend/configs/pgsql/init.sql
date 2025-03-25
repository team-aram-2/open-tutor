DROP TABLE IF EXISTS users;
CREATE TABLE users (
  "user_id" TEXT NOT NULL PRIMARY KEY,
  "email" TEXT NOT NULL UNIQUE,
  "signed_up_at" TIMESTAMP DEFAULT NOW(),
  "first_name" TEXT NOT NULL,
  "last_name" TEXT NOT NULL,
  "account_locked" BOOLEAN DEFAULT 'f',
  "password_hash" TEXT NOT NULL,
  "stripe_customer_id" TEXT UNIQUE
);
COMMENT ON TABLE "users" IS 'Base User object containing shared details needed for all users.';

DROP TABLE IF EXISTS tutors;
CREATE TABLE tutors (
  "user_id" TEXT NOT NULL PRIMARY KEY,
  "total_hours" INT DEFAULT 0,
  "hourly_rate" INT DEFAULT 10,
  "stripe_account_id" TEXT UNIQUE,
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);
COMMENT ON TABLE "tutors" IS 'Tutor object that extends user object with tutor specific information.';

DROP TABLE IF EXISTS available_skills;
CREATE TABLE available_skills (
  "id" TEXT NOT NULL PRIMARY KEY,
  "title" TEXT NOT NULL,
  "description" TEXT NOT NULL
);
COMMENT ON COLUMN available_skills.id IS 'Unique identifier for the skill, uuid.';
COMMENT ON COLUMN available_skills.title IS 'title/name for the skill';
COMMENT ON COLUMN available_skills.description IS 'Description for the skill';

DROP TABLE IF EXISTS tutor_skills;
CREATE TABLE tutor_skills (
  skill_id TEXT NOT NULL,
  tutor_id TEXT NOT NULL,
  validated BOOLEAN DEFAULT FALSE,
  PRIMARY KEY (skill_id, tutor_id),
  FOREIGN KEY (skill_id) REFERENCES available_skills(id) ON DELETE CASCADE,
  FOREIGN KEY (tutor_id) REFERENCES tutors(user_id) ON DELETE CASCADE
);

COMMENT ON COLUMN tutor_skills.skill_id IS 'Unique identifier for the skill.';
COMMENT ON COLUMN tutor_skills.tutor_id IS 'Tutor UUID.';
COMMENT ON COLUMN tutor_skills.validated IS 'If tutor skill is validated.';

DROP TABLE IF EXISTS meetings;
CREATE TABLE meetings (
  "id" TEXT NOT NULL PRIMARY KEY,
  "tutor_id" TEXT NOT NULL,
  "student_id" TEXT NOT NULL,
  "start_at" TIMESTAMP NOT NULL,
  "end_at" TIMESTAMP NOT NULL,
  "zoom_join_link" TEXT NOT NULL,
  "zoom_host_link" TEXT NOT NULL,
  "created_at" TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (student_id) REFERENCES users(user_id),
  FOREIGN KEY (tutor_id) REFERENCES tutors(user_id)
);
COMMENT ON COLUMN meetings.id IS 'Meeting unique id, uuid';

DROP TABLE IF EXISTS conversations;
CREATE TABLE conversations (
  "id" TEXT NOT NULL PRIMARY KEY,
  "user_ids" TEXT[] NOT NULL
);
CREATE INDEX idx_user_ids ON conversations USING GIN (user_ids);

DROP TABLE IF EXISTS messages;
CREATE TABLE messages (
  "id" TEXT NOT NULL PRIMARY KEY,
  "sent_at" TIMESTAMP NOT NULL,
  "origin_id" TEXT NOT NULL,
  "conversation_id" TEXT NOT NULL,
  "message" TEXT NOT NULL,
  FOREIGN KEY (origin_id) REFERENCES users(user_id),
  FOREIGN KEY (conversation_id) REFERENCES conversations(id)
);
COMMENT ON COLUMN messages.id IS 'Unique identifier for the message.';
COMMENT ON COLUMN messages.sent_at IS 'Date time the message was sent.';
COMMENT ON COLUMN messages.origin_id IS 'Unique identifier for the originID for the message.';
COMMENT ON COLUMN messages.message IS 'Message text content.';

DROP TABLE IF EXISTS message_attachments;
CREATE TABLE message_attachments (
  "attachment_id" TEXT NOT NULL PRIMARY KEY,
  "message_id" TEXT NOT NULL,
  "file_name" TEXT NOT NULL,
  "mime_type" TEXT NOT NULL,
  "url" TEXT NOT NULL,
  FOREIGN KEY (message_id) REFERENCES messages(id)
);
COMMENT ON COLUMN message_attachments.message_id IS 'Unique identifier for the message the attachement belongs to.';
COMMENT ON COLUMN message_attachments.file_name IS 'Friendly name of the attachment.';
COMMENT ON COLUMN message_attachments.mime_type IS 'Mimetype of the attachment.';
COMMENT ON COLUMN message_attachments.url IS 'URI for of the attachment data.';

DROP TYPE IF EXISTS user_type;
CREATE TYPE user_type AS ENUM ('tutor', 'student');

DROP TABLE IF EXISTS ratings;
CREATE TABLE ratings (
  "id" TEXT NOT NULL PRIMARY KEY,
  "rating_type" user_type DEFAULT NULL,
  "user_id" TEXT DEFAULT NULL,
  "reviewer_user_id" TEXT DEFAULT NULL,
  "meeting_id" TEXT DEFAULT NULL,
  "professionalism" INT DEFAULT NULL,
  "knowledge" INT DEFAULT NULL,
  "communication" INT DEFAULT NULL,
  "punctuality" INT DEFAULT NULL,
  "overall" INT NOT NULL,
  "comment" TEXT DEFAULT NULL,
  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (reviewer_user_id) REFERENCES users(user_id),
  FOREIGN KEY (meeting_id) REFERENCES meetings(id)
);

DROP TABLE IF EXISTS key_pairs;
CREATE TABLE key_pairs (
  "id" TEXT NOT NULL PRIMARY KEY,
  "public_key" TEXT NOT NULL,
  "private_key" TEXT NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION array_sort (arr TEXT[]) RETURNS TEXT[] IMMUTABLE AS
$$
  SELECT array_agg(x ORDER BY x) FROM UNNEST(arr) x;
$$ LANGUAGE SQL;

CREATE UNIQUE INDEX unique_ids ON conversations (array_sort(user_ids));