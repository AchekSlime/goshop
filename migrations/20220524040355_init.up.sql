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

CREATE TABLE cart
(
    product_id INTEGER REFERENCES product (id),
    user_id    integer
);

CREATE TABLE story
(
    id            serial not null primary key,
    user_id       integer,
    creation_date date
);

CREATE TABLE story_product
(
    story_id   INTEGER REFERENCES story (id),
    product_id INTEGER REFERENCES product (id)
);