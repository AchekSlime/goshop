CREATE TABLE category
(
    id   serial       not null primary key,
    name varchar(255) not null
);

CREATE TABLE product
(
    id          serial       not null primary key,
    title       varchar(255) not null,
    price       float        not null,
    holder_name varchar(255) not null,
    category_id INTEGER REFERENCES category (id)
);
