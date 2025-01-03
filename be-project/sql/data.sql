INSERT INTO data (name, email)
VALUES ($1, $2)
RETURNING id, name, email;

SELECT id, name, email
FROM data;
