# galactic-api

## database schema
CREATE TABLE spacecrafts (
    id INT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Class VARCHAR(255) NOT NULL,
    Armament JSON NOT NULL,
    Crew INT NOT NULL,
    Image VARCHAR(255),
    Value FLOAT NOT NULL,
    Status VARCHAR(255) NOT NULL
);