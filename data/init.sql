CREATE TABLE hotel (
    id SERIAL PRIMARY KEY,
    name text,
    address text,
    city text,
    reviews integer,
    rating real
);

INSERT INTO
    hotel (name, address, city, reviews, rating)
VALUES
    (
        'name',
        'address',
        'city',
        100,
        3.33
    ),
    (
        'name',
        'address',
        'city',
        100,
        3.33
    ),
    (
        'name',
        'address',
        'city',
        100,
        3.33
    );