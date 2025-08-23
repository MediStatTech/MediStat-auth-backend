---Test
CREATE TABLE staff_roles (
    role_id SERIAL PRIMARY KEY,
    code    VARCHAR(50) UNIQUE NOT NULL,   -- doctor, nurse, administrator
    name    VARCHAR(100) NOT NULL          -- Человекочитаемое название
);


CREATE TABLE personal (
    staff_id     BIGSERIAL PRIMARY KEY,
    full_name    VARCHAR(255) NOT NULL,
    role_id      INT NOT NULL REFERENCES staff_roles(role_id),
    department   VARCHAR(100),
    phone        VARCHAR(50),
    email        VARCHAR(100) UNIQUE,
    hired_at     DATE,
    dismissed_at DATE,
    created_at   TIMESTAMP NOT NULL DEFAULT now(),
    updated_at   TIMESTAMP NOT NULL DEFAULT now()
);
