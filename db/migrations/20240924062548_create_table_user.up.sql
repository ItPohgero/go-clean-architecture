CREATE TABLE
    users (
        id VARCHAR(100) NOT NULL,
        name VARCHAR(100) NOT NULL,
        password VARCHAR(100) NOT NULL,
        token VARCHAR(100) NULL,
        created_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT NOT NULL,
        updated_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT NOT NULL,
        PRIMARY KEY (id)
    );