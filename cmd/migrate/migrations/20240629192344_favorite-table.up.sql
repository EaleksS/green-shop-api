CREATE TABLE IF NOT EXISTS favorite (
  id SERIAL PRIMARY KEY,
	userId INT NOT NULL,
	productId INT NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);