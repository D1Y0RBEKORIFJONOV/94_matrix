CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS books(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255),
    author VARCHAR(255),
    publisher_date TIMESTAMP,
    isb VARCHAR(255)
);