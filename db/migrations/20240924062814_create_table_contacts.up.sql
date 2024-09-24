CREATE TABLE
    contacts (
        id VARCHAR(100) NOT NULL,
        first_name VARCHAR(100) NOT NULL,
        last_name VARCHAR(100),
        email VARCHAR(100),
        phone VARCHAR(100),
        user_id VARCHAR(100) NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        PRIMARY KEY (id),
        CONSTRAINT fk_contacts_user_id FOREIGN KEY (user_id) REFERENCES users (id)
    );