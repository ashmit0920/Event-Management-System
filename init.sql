CREATE TABLE attendees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    roll VARCHAR(100) NOT NULL,
    qr_data VARCHAR(100)
);
