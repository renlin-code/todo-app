CREATE TABLE users (
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE categories (
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255)
);

CREATE TABLE users_categories (
    id SERIAL NOT NULL UNIQUE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE tasks (
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255),
    done BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE categories_tasks (
    id SERIAL NOT NULL UNIQUE,
    task_id INT REFERENCES tasks(id) ON DELETE CASCADE NOT NULL,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE NOT NULL
);
