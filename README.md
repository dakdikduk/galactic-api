# galactic-api
Framework inspired by https://github.com/bxcodec/go-clean-arch for DDD (domain driven architecture)

## database schema
```
CREATE TABLE spacecrafts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    class VARCHAR(255) NOT NULL,
    armament JSON NOT NULL,
    crew INT NOT NULL,
    image VARCHAR(255) DEFAULT NULL,
    value FLOAT NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```