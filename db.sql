CREATE TABLE authors
(
  id bigserial NOT NULL,
  firstname varchar(32) NOT NULL,
  lastname varchar(32) NOT NULL
);

ALTER TABLE authors ADD CONSTRAINT pk_authors
  PRIMARY KEY (id);

CREATE TABLE books
(
  id serial NOT NULL,
  title varchar(64) NOT NULL,
  author_id bigserial NOT NULL
);

ALTER TABLE books ADD CONSTRAINT pk_books
  PRIMARY KEY (id);

ALTER TABLE books ADD CONSTRAINT fk_books_author_id
  FOREIGN KEY (author_id) REFERENCES authors (id);

