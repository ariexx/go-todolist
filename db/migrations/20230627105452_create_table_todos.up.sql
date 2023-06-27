CREATE TABLE todos(
                       id SERIAL PRIMARY KEY ,
                       title VARCHAR(255) NOT NULL,
                       description VARCHAR(255) NOT NULL,
                       user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP NULL
) ENGINE = InnoDB;