create database unit_test;
use unit_test;
create table users (
  id bigint unsigned primary key not null,
  name varchar(255)
);

insert into users(id, name) values (1, '张三');