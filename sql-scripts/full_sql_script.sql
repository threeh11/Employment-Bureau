CREATE TABLE types_activity
(
    id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) NOT NULL PRIMARY KEY,
	name_activity VARCHAR(20) NOT NULL 
);
COMMENT ON TABLE types_activity IS 'Вид деятельности ';
CREATE UNIQUE INDEX ON types_activity (name_activity);

CREATE TABLE employers(
	id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) NOT NULL PRIMARY KEY,
	name_employer VARCHAR(30) NOT NULL,
	address_employer VARCHAR(50) NOT NULL,
	number_employer VARCHAR(12) NOT NULL
);

COMMENT ON TABLE employers IS 'Работодатели';
CREATE INDEX ON employers(id_activity);

CREATE TABLE job_seekers(
	id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) NOT NULL PRIMARY KEY,
	surname_seekers VARCHAR(30) NOT NULL,
	name_seekers VARCHAR(30) NOT NULL,
	patronymic_seekers VARCHAR(30) NOT NULL,
	qualification_seekers VARCHAR(30) NOT NULL,
	id_activity INT NOT NULL,
	FOREIGN KEY (id_activity) REFERENCES types_activity (id),
	other_data VARCHAR(10000) NOT NULL,
	need_money_seekers INT NOT NULL
);

COMMENT ON TABLE job_seekers IS 'Соискатели';
CREATE INDEX ON job_seekers(id_activity);
CREATE INDEX ON job_seekers(name_seekers);
CREATE INDEX ON job_seekers(qualification_seekers);

CREATE TABLE dealings(
	id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) NOT NULL PRIMARY KEY,
	id_job_seekers INT NOT NULL,
	FOREIGN KEY (id_job_seekers) REFERENCES job_seekers (id),
	id_employer INT NOT NULL,
	FOREIGN KEY (id_employer) REFERENCES employers (id),
    id_activity INT NOT NULL,
    FOREIGN KEY (id_activity) REFERENCES types_activity (id),
    is_vacancy SMALLINT(1) NOT NULL,
    post VARCHAR(40),
	commission INT,
    description VARCHAR(10000),
    get_money INT
);

COMMENT ON TABLE dealings IS 'Сделки';
CREATE INDEX ON dealings(id_activity);
CREATE INDEX ON dealings(post);
CREATE INDEX ON dealings(get_money);

INSERT INTO types_activity (id, name_activity)
VALUES
    (1, 'Программист'),
    (2, 'Дизайнер'),
    (3, 'Маркетолог'),
    (4, 'Администратор'),
    (5, 'Финансист'),
    (6, 'Менеджер'),
    (7, 'Инженер'),
    (8, 'Архитектор'),
    (9, 'Психолог'),
    (10, 'Учитель');

INSERT INTO employers (id, name_employer, address_employer, number_employer)
VALUES
    (1, 'ООО "Рога и копыта"', 'ул. Центральная, 10', '123-45-67'),
    (2, 'ИП Иванов', 'ул. Солнечная, 5', '987-65-43'),
    (3, 'ЗАО "Строитель"', 'ул. Строительная, 15', '456-78-90'),
    (4, 'ООО "Медицина"', 'ул. Здоровая, 7', '789-01-23'),
    (5, 'ИП Петров', 'ул. Пушкина, 25', '654-32-10'),
    (6, 'ЗАО "Техника"', 'ул. Техническая, 3', '210-43-65'),
    (7, 'ООО "Коммерция"', 'ул. Торговая, 12', '876-54-32'),
    (8, 'ИП Сидоров', 'ул. Ленина, 8', '543-21-09'),
    (9, 'ООО "Стройка"', 'ул. Строительная, 9', '901-23-45'),
    (10, 'ИП Васильев', 'ул. Мира, 14', '432-10-98');

INSERT INTO job_seekers (id, surname_seekers, name_seekers, patronymic_seekers, qualification_seekers, id_activity, other_data, need_money_seekers)
VALUES
    (1, 'Иванов', 'Иван', 'Иванович', 'Высшее', 1, '...', 50000),
    (2, 'Петров', 'Петр', 'Петрович', 'Среднее', 2, '...', 30000),
    (3, 'Сидоров', 'Алексей', 'Алексеевич', 'Высшее', 3, '...', 45000),
    (4, 'Васильева', 'Елена', 'Ивановна', 'Высшее', 4, '...', 55000),
    (5, 'Медведев', 'Дмитрий', 'Сергеевич', 'Среднее', 5, '...', 25000),
    (6, 'Кузнецов', 'Игорь', 'Александрович', 'Среднее', 6, '...', 35000),
    (7, 'Попова', 'Анна', 'Николаевна', 'Высшее', 7, '...', 40000),
    (8, 'Смирнов', 'Александр', 'Владимирович', 'Высшее', 8, '...', 60000),
    (9, 'Николаев', 'Сергей', 'Алексеевич', 'Среднее', 9, '...', 20000),
    (10, 'Козлова', 'Ольга', 'Ивановна', 'Высшее', 10, '...', 35000);

INSERT INTO dealings (id, id_job_seekers, id_employer, id_activity, is_vacancy, post, commission, description, get_money)
VALUES
    (1, 1, 1, 1, 1, 'Программист', 10, '...', 50000),
    (2, 2, 3, 2, 1, 'Дизайнер', 8, '...', 30000),
    (3, 3, 2, 3, 1, 'Маркетолог', 12, '...', 45000),
    (4, 4, 4, 4, 1, 'Администратор', 6, '...', 55000),
    (5, 5, 6, 5, 1, 'Финансист', 7, '...', 25000),
    (6, 6, 5, 6, 1, 'Менеджер', 9, '...', 35000),
    (7, 7, 7, 7, 1, 'Инженер', 11, '...', 40000),
    (8, 8, 8, 8, 1, 'Архитектор', 15, '...', 60000),
    (9, 9, 9, 9, 1, 'Психолог', 13, '...', 20000),
    (10, 10, 10, 10, 1, 'Учитель', ...);