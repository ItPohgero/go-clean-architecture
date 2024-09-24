CREATE TABLE
    addresses (
        id VARCHAR(100) NOT NULL,
        street VARCHAR(100) NOT NULL,
        number VARCHAR(100) NOT NULL,
        neighborhood VARCHAR(100) NOT NULL,
        city VARCHAR(100) NOT NULL,
        state VARCHAR(100) NOT NULL,
        user_id VARCHAR(100) NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        PRIMARY KEY (id),
        CONSTRAINT fk_addresses_user_id FOREIGN KEY (user_id) REFERENCES users (id)
    );