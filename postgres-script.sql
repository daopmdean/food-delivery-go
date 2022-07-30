-- SQLINES LICENSE FOR EVALUATION USE ONLY

DROP TABLE IF EXISTS carts;
CREATE TABLE carts (
  user_id int NOT NULL,
  food_id int NOT NULL,
  quantity int NOT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id,food_id)
) ;
CREATE INDEX food_id ON carts (food_id);


DROP TABLE IF EXISTS categories;
CREATE SEQUENCE categories_seq;
CREATE TABLE categories (
  id int NOT NULL DEFAULT NEXTVAL ('categories_seq'),
  name varchar(100) NOT NULL,
  description text,
  icon json DEFAULT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;


DROP TABLE IF EXISTS cities;
CREATE SEQUENCE cities_seq;
CREATE TABLE cities (
  id int NOT NULL DEFAULT NEXTVAL ('cities_seq'),
  title varchar(100) NOT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;


DROP TABLE IF EXISTS food_likes;
CREATE TABLE food_likes (
  user_id int NOT NULL,
  food_id int NOT NULL,
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id,food_id)
) ;
CREATE INDEX food_id ON food_likes (food_id);


DROP TABLE IF EXISTS food_ratings;
CREATE SEQUENCE food_ratings_seq;
CREATE TABLE food_ratings (
  id int NOT NULL DEFAULT NEXTVAL ('food_ratings_seq'),
  user_id int NOT NULL,
  food_id int NOT NULL,
  point double precision DEFAULT '0',
  comment text,
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;
CREATE INDEX food_id ON food_ratings (food_id);


DROP TABLE IF EXISTS foods;
CREATE SEQUENCE foods_seq;
CREATE TABLE foods (
  id int NOT NULL DEFAULT NEXTVAL ('foods_seq'),
  restaurant_id int NOT NULL,
  category_id int DEFAULT NULL,
  name varchar(255) NOT NULL,
  description text,
  price double precision NOT NULL,
  images json NOT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;
CREATE INDEX restaurant_id ON foods (restaurant_id);
CREATE INDEX category_id ON foods (category_id);
CREATE INDEX status ON foods (status);


DROP TABLE IF EXISTS images;
CREATE SEQUENCE images_seq;
CREATE TABLE images (
  id int NOT NULL DEFAULT NEXTVAL ('images_seq'),
  file_name varchar(100) NOT NULL,
  width int NOT NULL,
  height int NOT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;


DROP TABLE IF EXISTS order_details;
CREATE SEQUENCE order_details_seq;
CREATE TABLE order_details (
  id int NOT NULL DEFAULT NEXTVAL ('order_details_seq'),
  order_id int NOT NULL,
  food_origin json DEFAULT NULL,
  price double precision NOT NULL,
  quantity int NOT NULL,
  discount double precision DEFAULT '0',
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;
CREATE INDEX order_id ON order_details (order_id);


DROP TABLE IF EXISTS order_trackings;
CREATE TYPE state_enum AS ENUM ('waiting_for_shipper', 'preparing', 'on_the_way','delivered','cancel');
CREATE SEQUENCE order_trackings_seq;
CREATE TABLE order_trackings (
  id int NOT NULL DEFAULT NEXTVAL ('order_trackings_seq'),
  order_id int NOT NULL,
  state state_enum NOT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;
CREATE INDEX order_id ON order_trackings(order_id);


DROP TABLE IF EXISTS orders;
CREATE SEQUENCE orders_seq;
CREATE TABLE orders (
  id int NOT NULL DEFAULT NEXTVAL ('orders_seq'),
  user_id int NOT NULL,
  total_price double precision NOT NULL,
  shipper_id int DEFAULT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;
CREATE INDEX user_id ON orders (user_id);
CREATE INDEX shipper_id ON orders (shipper_id);


DROP TABLE IF EXISTS restaurant_foods;
CREATE TABLE restaurant_foods (
  restaurant_id int NOT NULL,
  food_id int NOT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (restaurant_id,food_id)
) ;
CREATE INDEX food_id ON restaurant_foods (food_id);


DROP TABLE IF EXISTS restaurant_likes;
CREATE TABLE restaurant_likes (
  restaurant_id int NOT NULL,
  user_id int NOT NULL,
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (restaurant_id,user_id)
) ;
CREATE INDEX user_id ON restaurant_likes (user_id);


DROP TABLE IF EXISTS restaurant_ratings;
CREATE SEQUENCE restaurant_ratings_seq;
CREATE TABLE restaurant_ratings (
  id int NOT NULL DEFAULT NEXTVAL ('restaurant_ratings_seq'),
  user_id int NOT NULL,
  restaurant_id int NOT NULL,
  point double precision NOT NULL DEFAULT '0',
  comment text,
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;
CREATE INDEX user_id ON restaurant_ratings (user_id);
CREATE INDEX restaurant_id ON restaurant_ratings (restaurant_id);


DROP TABLE IF EXISTS restaurants;
CREATE SEQUENCE restaurants_seq;
CREATE TABLE restaurants (
  id int NOT NULL DEFAULT NEXTVAL ('restaurants_seq'),
  owner_id int NOT NULL,
  name varchar(50) NOT NULL,
  addr varchar(255) NOT NULL,
  city_id int DEFAULT NULL,
  lat double precision DEFAULT NULL,
  lng double precision DEFAULT NULL,
  cover json NOT NULL,
  logo json NOT NULL,
  shipping_fee_per_km double precision DEFAULT '0',
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;
CREATE INDEX owner_id ON restaurants (owner_id);
CREATE INDEX city_id ON restaurants (city_id);
CREATE INDEX status ON restaurants (status);


DROP TABLE IF EXISTS user_addresses;
CREATE SEQUENCE user_addresses_seq;
CREATE TABLE user_addresses (
  id int NOT NULL DEFAULT NEXTVAL ('user_addresses_seq'),
  user_id int NOT NULL,
  city_id int NOT NULL,
  title varchar(100) DEFAULT NULL,
  icon json DEFAULT NULL,
  addr varchar(255) NOT NULL,
  lat double precision DEFAULT NULL,
  lng double precision DEFAULT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;
CREATE INDEX user_id ON user_addresses (user_id);
CREATE INDEX city_id ON user_addresses (city_id);


DROP TABLE IF EXISTS user_device_tokens;
CREATE SEQUENCE user_device_tokens_seq;
CREATE TABLE user_device_tokens (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('user_device_tokens_seq'),
  user_id int check (user_id > 0) DEFAULT NULL,
  is_production smallint DEFAULT '0',
  os enum('ios','android','web') DEFAULT 'ios' COMMENT '1: iOS, 2: Android',
  token varchar(255) DEFAULT NULL,
  status smallint unsigned NOT NULL DEFAULT '1',
  updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY `user_id` (user_id) USING BTREE,
  KEY `os` (os) USING BTREE
) ENGINE=InnoDB;


DROP TABLE IF EXISTS users;
CREATE SEQUENCE users_seq;
CREATE TABLE users (
  id int NOT NULL DEFAULT NEXTVAL ('users_seq'),
  email varchar(50) NOT NULL,
  fb_id varchar(50) DEFAULT NULL,
  gg_id varchar(50) DEFAULT NULL,
  password varchar(50) NOT NULL,
  salt varchar(50) DEFAULT NULL,
  last_name varchar(50) NOT NULL,
  first_name varchar(50) NOT NULL,
  phone varchar(20) DEFAULT NULL,
  role enum('user','admin','shipper') NOT NULL DEFAULT 'user',
  avatar json DEFAULT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY `email` (email)
) ENGINE=InnoDB;


INSERT INTO cities (id, title, status, created_at, updated_at) VALUES
(1, 'An Giang', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(2, 'Vũng Tàu', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(3, 'Bắc Giang', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(4, 'Bắc Cạn', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(5, 'Bạc Liêu', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(6, 'Bắc Ninh', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(7, 'Bến Tre', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(8, 'Bình Định', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(9, 'Bình Dương', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(10, 'Bình Phước', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(11, 'Bình Thuận', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(12, 'Cà Mau', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(13, 'Cần Thơ', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(14, 'Cao Bằng', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(15, 'Đà Nẵng', 1, '2020-05-18 06:52:21', '2020-05-18 06:52:21'),
(16, 'Đắk Lắk', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(17, 'Đắk Nông', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(18, 'Điện Biên', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(19, 'Đồng Nai', 1, '2020-05-18 06:52:21', '2020-05-18 06:52:21'),
(20, 'Đồng Tháp', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(21, 'Gia Lai', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(22, 'Hà Giang', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(23, 'Hà Nam', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(24, 'Hà Nội', 1, '2020-05-18 06:52:21', '2020-05-18 06:52:21'),
(25, 'Hà Tĩnh', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(26, 'Hải Dương', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(27, 'Hải Phòng', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(28, 'Hậu Giang', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(29, 'Hoà Bình', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(30, 'Hưng Yên', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(31, 'Khánh Hoà', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(32, 'Kiên Giang', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(33, 'Kon Tum', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(34, 'Lai Châu', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(35, 'Lâm Đồng', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(36, 'Lạng Sơn', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(37, 'Lào Cai', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(38, 'Long An', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(39, 'Nam Định', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(40, 'Nghệ An', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(41, 'Ninh Bình', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(42, 'Ninh Thuận', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(43, 'Phú Thọ', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(44, 'Phú Yên', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(45, 'Quảng Bình', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(46, 'Quảng Namm', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(47, 'Quãng Ngãi', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(48, 'Quãng Ninh', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(49, 'Quãng Trị', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(50, 'Sóc Trăng', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(51, 'Sơn La', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(52, 'Tây Ninh', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(53, 'Thái Bình', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(54, 'Thái Nguyên', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(55, 'Thanh Hoá', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(56, 'Huế', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(57, 'Tiền Giang', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(58, 'Hồ Chí Minh', 1, '2020-05-18 06:41:51', '2020-05-18 06:41:51'),
(59, 'Trà Vinh', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(60, 'Tuyên Quang', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(61, 'Vĩnh Long', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(62, 'Vĩnh Phúc', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(63, 'Yên Bái', 1, '2020-05-18 06:55:19', '2020-05-18 06:55:19');
