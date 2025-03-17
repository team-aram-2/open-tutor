DROP TABLE IF EXISTS users;
CREATE TABLE users (
  "user_id" TEXT NOT NULL PRIMARY KEY,
  "email" TEXT NOT NULL UNIQUE,
  "signed_up_at" TIMESTAMP DEFAULT NOW(),
  "first_name" TEXT NOT NULL,
  "last_name" TEXT NOT NULL,
  "account_locked" BOOLEAN DEFAULT 'f',
  "password_hash" TEXT NOT NULL
);
COMMENT ON TABLE "users" IS 'Base User object containing shared details needed for all users.';

DROP TABLE IF EXISTS tutors;
CREATE TABLE tutors (
  "user_id" TEXT NOT NULL PRIMARY KEY,
  "total_hours" INT DEFAULT 0,
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);
COMMENT ON TABLE "tutors" IS 'Tutor object that extends user object with tutor specific information.';

-- Drop existing tables if they exist
DROP TABLE IF EXISTS academic_categories CASCADE;
-- Create the academic_categories table
CREATE TABLE academic_categories (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL UNIQUE
);
COMMENT ON TABLE academic_categories IS 'List of standard academic categories.';
COMMENT ON COLUMN academic_categories.id IS 'Unique identifier for the academic category.';
COMMENT ON COLUMN academic_categories.name IS 'Name of the academic category.';
-- Insert predefined academic categories
INSERT INTO academic_categories (id, name) VALUES
  (gen_random_uuid(), 'Mathematics'),
  (gen_random_uuid(), 'Science'),
  (gen_random_uuid(), 'Physics'),
  (gen_random_uuid(), 'Chemistry'),
  (gen_random_uuid(), 'Biology'),
  (gen_random_uuid(), 'Computer Science'),
  (gen_random_uuid(), 'Engineering'),
  (gen_random_uuid(), 'Medicine'),
  (gen_random_uuid(), 'Humanities'),
  (gen_random_uuid(), 'History'),
  (gen_random_uuid(), 'Literature'),
  (gen_random_uuid(), 'Philosophy'),
  (gen_random_uuid(), 'Linguistics'),
  (gen_random_uuid(), 'Arts'),
  (gen_random_uuid(), 'Music'),
  (gen_random_uuid(), 'Visual Arts'),
  (gen_random_uuid(), 'Performing Arts'),
  (gen_random_uuid(), 'Economics'),
  (gen_random_uuid(), 'Business'),
  (gen_random_uuid(), 'Psychology'),
  (gen_random_uuid(), 'Sociology'),
  (gen_random_uuid(), 'Political Science'),
  (gen_random_uuid(), 'Law'),
  (gen_random_uuid(), 'Education'),
  (gen_random_uuid(), 'Environmental Science'),
  (gen_random_uuid(), 'Geography'),
  (gen_random_uuid(), 'Anthropology'),
  (gen_random_uuid(), 'Astronomy');

-- Drop existing tables if they exist
DROP TABLE IF EXISTS available_skills CASCADE;
-- Create the available_skills table with a foreign key to academic_categories
CREATE TABLE available_skills (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  category_id UUID NOT NULL,
  FOREIGN KEY (category_id) REFERENCES academic_categories(id) ON DELETE CASCADE
);

COMMENT ON TABLE available_skills IS 'Table defining available skills, each associated with a category.';
COMMENT ON COLUMN available_skills.id IS 'Unique identifier for the skill.';
COMMENT ON COLUMN available_skills.title IS 'Title or name of the skill.';
COMMENT ON COLUMN available_skills.description IS 'Description of the skill.';
COMMENT ON COLUMN available_skills.category_id IS 'Reference to the academic category that this skill belongs to.';

DROP TABLE IF EXISTS tutor_skills;

CREATE TABLE tutor_skills (
  skill_id UUID NOT NULL,
  tutor_id UUID NOT NULL,
  PRIMARY KEY (skill_id, tutor_id),
  FOREIGN KEY (skill_id) REFERENCES available_skills(id) ON DELETE CASCADE,
  FOREIGN KEY (tutor_id) REFERENCES tutors(user_id) ON DELETE CASCADE
);

COMMENT ON COLUMN tutor_skills.skill_id IS 'Unique identifier for the skill.';
COMMENT ON COLUMN tutor_skills.tutor_id IS 'Tutor UUID.';
COMMENT ON COLUMN tutor_skills.validated IS 'If tutor skill is validated.';

CREATE TABLE questions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  skill_id UUID NOT NULL,
  question TEXT NOT NULL,
  correct_answers TEXT[] NOT NULL,
  FOREIGN KEY (skill_id) REFERENCES available_skills(id) ON DELETE CASCADE
);

CREATE TABLE quiz_attempts (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  tutor_id UUID NOT NULL,
  skill_id UUID NOT NULL,
  passed BOOLEAN DEFAULT FALSE,
  attempted_at TIMESTAMP DEFAULT NOW(),
  completed_at TIMESTAMP DEFAULT NULL,
  FOREIGN KEY (tutor_id) REFERENCES tutors(user_id) ON DELETE CASCADE,
  FOREIGN KEY (skill_id) REFERENCES available_skills(id) ON DELETE CASCADE
);

CREATE TABLE quiz_attempt_questions (
  quiz_attempt_id UUID NOT NULL,
  question_id UUID NOT NULL,
  PRIMARY KEY (quiz_attempt_id, question_id),
  FOREIGN KEY (quiz_attempt_id) REFERENCES quiz_attempts(id) ON DELETE CASCADE,
  FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS meetings;
CREATE TABLE meetings (
  "id" TEXT NOT NULL PRIMARY KEY,
  "tutor_id" TEXT NOT NULL,
  "student_id" TEXT NOT NULL,
  "date_created" TIMESTAMP NOT NULL,
  "start_at" TIMESTAMP NOT NULL,
  "end_at" TIMESTAMP NOT NULL,
  "zoom_join_link" TEXT NOT NULL,
  "zoom_host_link" TEXT NOT NULL,
  "created_at" TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (student_id) REFERENCES users(user_id),
  FOREIGN KEY (tutor_id) REFERENCES tutors(user_id)
);
COMMENT ON COLUMN meetings.id IS 'Meeting unique id, uuid';

DROP TABLE IF EXISTS messages;
CREATE TABLE messages (
  "id" TEXT NOT NULL PRIMARY KEY,
  "sent_at" TIMESTAMP NOT NULL,
  "origin_id" TEXT NOT NULL,
  "recipient_id" TEXT NOT NULL,
  "message" TEXT NOT NULL,
  FOREIGN KEY (origin_id) REFERENCES users(user_id),
  FOREIGN KEY (recipient_id) REFERENCES users(user_id)
);
COMMENT ON COLUMN messages.id IS 'Unique identifier for the message.';
COMMENT ON COLUMN messages.sent_at IS 'Date time the message was sent.';
COMMENT ON COLUMN messages.origin_id IS 'Unique identifier for the originID for the message.';
COMMENT ON COLUMN messages.recipient_id IS 'UserId of the recipient.';
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
COMMENT ON COLUMN message_attachments.message_id IS 'Friendly name of the attachment.';
COMMENT ON COLUMN message_attachments.message_id IS 'Mimetype of the attachment.';
COMMENT ON COLUMN message_attachments.message_id IS 'URI for of the attachment data.';

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