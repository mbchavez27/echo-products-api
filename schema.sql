create table products (
id serial primary key,
name varchar(200) not null,
price decimal not null,
ratings float not null,
description text not null,
category varchar(200) not null
);