CREATE TABLE shops
(
  id serial not null unique,
  title varchar(255) not null,
  link varchar(255),
  class_name varchar(255)
);
CREATE TABLE models
(
  id serial not null unique,
  shop_id int references shops(id) on delete cascade not null,
  title varchar(255) not null,
  price numeric not null,
  avail int default 0 not null,
  img_url varchar(255) default '' not null,
  page_url varchar(255) not null UNIQUE,
  size varchar(255) default '' not null
);
INSERT INTO shops (title, link, class_name) VALUES
('nike.com - женские', 'https://www.nike.com/', 'nike');

INSERT INTO models (shop_id, title, link, class_name) VALUES
('nike.com - женские', 'https://www.nike.com/', 'nike');