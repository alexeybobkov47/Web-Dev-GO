create database
if not exists Blog;

CREATE TABLE
if not exists `Blog`.`Blogs`
(
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` TEXT NOT NULL,
  `description` TEXT NULL,
  PRIMARY KEY
(`id`),
  UNIQUE INDEX `id_UNIQUE`
(`id` ASC) VISIBLE);

CREATE TABLE
if not exists `Blog`.`Posts`
(
  `id` INT NOT NULL AUTO_INCREMENT,
  `blog_id` INT NOT NULL,
  `header`TEXT NOT NULL,
  `text` TEXT NOT NULL,
  PRIMARY KEY
(`id`),
  UNIQUE INDEX `id_UNIQUE`
(`id` ASC) VISIBLE);

insert into Blog.Blogs
  (name, description)
values
  ("BlogsFirst ", "First Description"),
  ("Second ", "Second Description");

insert into Blog.Posts
  (blog_id, header, text)
values
  (0, "1 post", "1 text"),
  (0, "2 post", "2 text"),
  (0, "3 post", "3 text"),
  (0, "4 post", "4 text"),
  (0, "5 post", "5 text");
  
  