CREATE TABLE logs (
    id VARCHAR(36) PRIMARY KEY,
    ip VARCHAR(40) NOT NULL,
    date DATE NOT NULL,
    resource TEXT NOT NULL,
    method VARCHAR(5) NOT NULL,
    protocol VARCHAR(10) NOT NULL,
    status_code INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT created_log UNIQUE (ip, date, resource, method)
);

