create database recommendation;

-------------------------

create table history
(
	message_id varchar not null,
	user_id varchar not null,
	product_id varchar not null,
	source varchar not null,
	date timestamp not null
);

create unique index history_message_id_uindex
	on history (message_id);

alter table history
	add constraint history_pk
		primary key (message_id);

-------------------------

create table bestseller
(
	product_id varchar not null,
	quantity int not null
);

create unique index bestseller_product_id_uindex
	on bestseller (product_id);

alter table bestseller
	add constraint bestseller_pk
		primary key (product_id);

---------------------------

create table products
(
	product_id varchar not null,
	category_id varchar not null
);
