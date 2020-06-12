CREATE TABLE logs (
    id VARCHAR(36) PRIMARY KEY,
    ip VARCHAR(40) NOT NULL,
    port INT NOT NULL,
    date TIMESTAMP NOT NULL,
    resource VARCHAR(200) NOT NULL,
    method VARCHAR(5) NOT NULL,
    protocol VARCHAR(10) NOT NULL,
    status_code INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT created_log UNIQUE (ip, date, resource, method)
);

