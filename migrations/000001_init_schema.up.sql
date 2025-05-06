-- Создаём таблицы для пользователей, столов и бронирований

CREATE TABLE users (
    id SERIAL PRIMARY KEY,   -- Уникальный идентификатор пользователя
    name TEXT,               -- Имя пользователя
    email TEXT UNIQUE        -- Электронная почта пользователя, уникальная
);

CREATE TABLE desks (
    id SERIAL PRIMARY KEY,   -- Уникальный идентификатор стола
    location TEXT,           -- Местоположение стола в офисе
    capacity INT             -- Вместимость стола (например, количество человек)
);

CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,   -- Уникальный идентификатор бронирования
    user_id INT REFERENCES users(id),   -- Внешний ключ на пользователя
    desk_id INT REFERENCES desks(id),   -- Внешний ключ на стол
    date DATE NOT NULL,       -- Дата бронирования
    UNIQUE (desk_id, date)    -- Гарантируем, что стол не будет забронирован дважды в одну дату
);