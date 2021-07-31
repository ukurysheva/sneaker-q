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
  price float not null,
  avail int,
  img_url varchar(255),
  page_url varchar(255) not null UNIQUE,
  size varchar(255)
);
INSERT INTO shops (title, link, class_name) VALUES
('nike.com - женские', 'https://www.nike.com/ru/w/womens-shoes-5e1x6zy7ok', 'nike');