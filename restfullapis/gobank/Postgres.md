## start a postgres instance

---

$ docker run --name some-postgres -e POSTGRES_PASSWORD=postgres -p host:docker_port -d postgres
$ docker run --name some-postgres -e POSTGRES_PASSWORD=postgres -p 5433:5432 -d postgres

cmd>> netstat -ano | findstr :5433
cmd>> docker exec -it some-postgres psql -U postgres

psql cmd:
    - CREATE database gobank;
    - \c gobank
    - \dt

// Create Test Table
CREATE TABLE test_users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);

// Insert Values - Use single quotes 'string' not -> "string" [""-> invalid in sql]
INSERT INTO test_users VALUES (1, 'Abhishek', 'abhishek@gmail.com');
INSERT INTO test_users (id, name) VALUES (1, 'Abhi');

SELECT * FROM test_user;
