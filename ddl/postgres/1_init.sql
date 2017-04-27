-- +migrate Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE classroom_role AS ENUM (
    'student',
    'moderator',
    'teacher',
    'administrator'
);

CREATE TABLE classes (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(50) DEFAULT 'An Unnamed Class'::character varying NOT NULL,
    current_unit uuid
);

CREATE TABLE local_identities (
    user_id uuid NOT NULL,
    password text NOT NULL
);

CREATE TABLE multiple_choice_answers (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    question_id uuid NOT NULL,
    answer text NOT NULL,
    correct boolean DEFAULT false NOT NULL
);

CREATE TABLE multiple_choice_questions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    class_id uuid NOT NULL,
    author_id uuid,
    question text NOT NULL,
    partial_credit boolean DEFAULT false NOT NULL,
    "timestamp" timestamp with time zone DEFAULT now() NOT NULL,
    unit_id uuid NOT NULL
);

CREATE TABLE multiple_choice_responses (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    question_id uuid,
    author_id uuid NOT NULL,
    response_id uuid NOT NULL
);


CREATE TABLE profiles (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    class_id uuid NOT NULL,
    role classroom_role DEFAULT 'student'::classroom_role NOT NULL
);

CREATE TABLE question_ratings (
    author_id uuid NOT NULL,
    question_id uuid NOT NULL,
    vote boolean NOT NULL
);

CREATE TABLE short_answer_questions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    class_id uuid NOT NULL,
    author_id uuid,
    prompt text NOT NULL,
    minimum_length integer,
    maximum_length integer,
    unit_id uuid NOT NULL,
    "timestamp" timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT length_check CHECK (((minimum_length < maximum_length) OR (minimum_length = 0))),
    CONSTRAINT short_answer_check CHECK (((maximum_length > minimum_length) OR (maximum_length > 0)))
);

CREATE TABLE short_answer_responses (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    question_id uuid,
    author_id uuid NOT NULL,
    response text
);

CREATE TABLE true_false_questions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    class_id uuid NOT NULL,
    author_id uuid,
    question text NOT NULL,
    correct_answer boolean NOT NULL,
    unit_id uuid NOT NULL,
    "timestamp" timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE true_false_responses (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    question_id uuid NOT NULL,
    author_id uuid,
    response boolean NOT NULL
);

CREATE TABLE units (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    class_id uuid NOT NULL,
    name text,
    weight integer DEFAULT 0 NOT NULL
);

CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    email character varying(320) NOT NULL
);

ALTER TABLE ONLY classes
    ADD CONSTRAINT classes_pkey PRIMARY KEY (id);

ALTER TABLE ONLY local_identities
    ADD CONSTRAINT local_identities_pkey PRIMARY KEY (user_id);

ALTER TABLE ONLY multiple_choice_answers
    ADD CONSTRAINT multiple_choice_answers_pkey PRIMARY KEY (id);

ALTER TABLE ONLY multiple_choice_questions
    ADD CONSTRAINT multiple_choice_questions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY multiple_choice_responses
    ADD CONSTRAINT multiple_choice_responses_pkey PRIMARY KEY (id);

ALTER TABLE ONLY profiles
    ADD CONSTRAINT profiles_pkey PRIMARY KEY (id);

ALTER TABLE ONLY question_ratings
    ADD CONSTRAINT ratings_question_id_author_id_key UNIQUE (question_id, author_id);

ALTER TABLE ONLY short_answer_questions
    ADD CONSTRAINT short_answer_questions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY short_answer_responses
    ADD CONSTRAINT short_answer_responses_pkey PRIMARY KEY (id);

ALTER TABLE ONLY true_false_questions
    ADD CONSTRAINT true_false_questions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY true_false_responses
    ADD CONSTRAINT true_false_responses_pkey PRIMARY KEY (id);

ALTER TABLE ONLY units
    ADD CONSTRAINT units_pkey PRIMARY KEY (id);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY classes
    ADD CONSTRAINT classes_current_unit_fkey FOREIGN KEY (current_unit) REFERENCES units(id) ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY multiple_choice_answers
    ADD CONSTRAINT multiple_choice_answers_question_id_fkey FOREIGN KEY (question_id) REFERENCES multiple_choice_questions(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY multiple_choice_questions
    ADD CONSTRAINT multiple_choice_author_fkey FOREIGN KEY (author_id) REFERENCES profiles(id) ON DELETE SET NULL;

ALTER TABLE ONLY multiple_choice_responses
    ADD CONSTRAINT multiple_choice_author_fkey FOREIGN KEY (author_id) REFERENCES profiles(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY multiple_choice_questions
    ADD CONSTRAINT multiple_choice_class_fkey FOREIGN KEY (class_id) REFERENCES classes(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY multiple_choice_responses
    ADD CONSTRAINT multiple_choice_question_fkey FOREIGN KEY (question_id) REFERENCES multiple_choice_questions(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY multiple_choice_responses
    ADD CONSTRAINT multiple_choice_response_fkey FOREIGN KEY (response_id) REFERENCES multiple_choice_answers(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY multiple_choice_questions
    ADD CONSTRAINT multiple_choice_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES units(id) ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY profiles
    ADD CONSTRAINT profiles_class_fkey FOREIGN KEY (class_id) REFERENCES classes(id) ON DELETE CASCADE;

ALTER TABLE ONLY profiles
    ADD CONSTRAINT profiles_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY question_ratings
    ADD CONSTRAINT ratings_rater_fkey FOREIGN KEY (author_id) REFERENCES profiles(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY short_answer_questions
    ADD CONSTRAINT short_answer_author_fkey FOREIGN KEY (author_id) REFERENCES profiles(id) ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY short_answer_responses
    ADD CONSTRAINT short_answer_author_fkey FOREIGN KEY (author_id) REFERENCES profiles(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY short_answer_questions
    ADD CONSTRAINT short_answer_class_fkey FOREIGN KEY (class_id) REFERENCES classes(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY short_answer_responses
    ADD CONSTRAINT short_answer_question_fkey FOREIGN KEY (question_id) REFERENCES short_answer_questions(id) ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY short_answer_questions
    ADD CONSTRAINT short_answer_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES units(id) ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY true_false_responses
    ADD CONSTRAINT true_false_author_fkey FOREIGN KEY (author_id) REFERENCES true_false_questions(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY true_false_questions
    ADD CONSTRAINT true_false_author_id_fkey FOREIGN KEY (author_id) REFERENCES profiles(id) ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY true_false_questions
    ADD CONSTRAINT true_false_class_id_fkey FOREIGN KEY (class_id) REFERENCES classes(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY true_false_responses
    ADD CONSTRAINT true_false_question_fkey FOREIGN KEY (question_id) REFERENCES true_false_questions(id) ON DELETE CASCADE;

ALTER TABLE ONLY true_false_questions
    ADD CONSTRAINT true_false_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES units(id) ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY units
    ADD CONSTRAINT units_class_id_fkey FOREIGN KEY (class_id) REFERENCES classes(id) ON UPDATE CASCADE ON DELETE CASCADE;