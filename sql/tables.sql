CREATE table app_user(
	id              serial primary key,
	firstname       varchar(50),
	lastname        varchar(50),
	emailid         varchar(100),
    gender          char(1),
    active          char(1),    
	mobile          bigint,
	password        varchar(50),
	usertype        char(1),
	createddate     timestamp,
	updateddate     timestamp
);

CREATE TABLE user_address (
    id          serial primary key,
    user_id     int references app_user(id),
    add1        varchar(100),
    add2        varchar(100),
    add3        varchar(100),
    city        varchar(100),
    state       varchar(100),
    pincode     int
);

CREATE TABLE user_order (
    id          serial primary key,
    user_id     int references app_user(id),
    tax         float,
    total       int,
    grand_total int
);

CREATE TABLE order_item (
    id          serial,
    order_id    int references user_order(id),
    item        varchar(50),
    quantity    int,
    price       int
);

CREATE TABLE restaurant (
    id          serial primary key,
    name        varchar(100),
    add1        varchar(100),
    add2        varchar(100),
    city        varchar(50),
    state       varchar(50),
    pincode     int,
    picture     text
)

CREATE TABLE cuisine (
    id              serial primary key,
    restaurant_id   int references restaurant(id),
    cuisine_name    varchar(50)
)

CREATE TABLE item (
    id              serial,
    cuisine_id      int references cuisine(id),
    item_name       varchar(50),
    price           int,
    quantity        char(1)
)

