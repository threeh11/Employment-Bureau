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
	id_activity INT NOT NULL,
	FOREIGN KEY (id_activity) REFERENCES types_activity (id),
	address_employer VARCHAR(50) NOT NULL,
	number_employre VARCHAR(12) NOT NULL
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

CREATE TABLE vacancy(
	id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) NOT NULL PRIMARY KEY,
	name_vacancy VARCHAR(30) NOT NULL,
	id_activity INT NOT NULL,
	FOREIGN KEY (id_activity) REFERENCES types_activity (id),
	description VARCHAR(10000) NOT NULL,
	get_money INT NOT NULL
);

CREATE TABLE dealing(
	id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) NOT NULL PRIMARY KEY,
	id_job_seekers INT NOT NULL,
	FOREIGN KEY (id_job_seekers) REFERENCES job_seekers (id),
	id_employer INT NOT NULL,
	FOREIGN KEY (id_employer) REFERENCES employers (id),
	post VARCHAR(40) NOT NULL,
	commission INT NOT NULL
);

COMMENT ON TABLE vacancy IS 'Вакансии';
CREATE INDEX ON vacancy(id_activity);
CREATE INDEX ON vacancy(name_vacancy);
CREATE INDEX ON vacancy(get_money);

-- Заполнение таблицы types_activity
INSERT INTO types_activity (name_activity) VALUES
('Activity 1'), ('Activity 2'), ('Activity 3'), ('Activity 4'), ('Activity 5'),
('Activity 6'), ('Activity 7'), ('Activity 8'), ('Activity 9'), ('Activity 10');
-- Заполнение таблицы employers
INSERT INTO employers (name_employer, id_activity, address_employer, number_employre) VALUES
('Employer 1', 1, 'Address 1', '1234567890'), ('Employer 2', 2, 'Address 2', '1234567891'),
('Employer 3', 3, 'Address 3', '1234567892'), ('Employer 4', 4, 'Address 4', '1234567893'),
('Employer 5', 5, 'Address 5', '1234567894'), ('Employer 6', 6, 'Address 6', '1234567895'),
('Employer 7', 7, 'Address 7', '1234567896'), ('Employer 8', 8, 'Address 8', '1234567897'),
('Employer 9', 9, 'Address 9', '1234567898'), ('Employer 10', 10, 'Address 10', '1234567899');
-- Заполнение таблицы job_seekers
INSERT INTO job_seekers (surname_seekers, name_seekers, patronymic_seekers, qualification_seekers,
id_activity, other_data, need_money_seekers) VALUES
('Seeker 1', 'Name 1', 'Patronymic 1', 'Qualification 1', 1, 'Other Data 1', 1000),
('Seeker 2', 'Name 2', 'Patronymic 2', 'Qualification 2', 2, 'Other Data 2', 2000),
('Seeker 3', 'Name 3', 'Patronymic 3', 'Qualification 3', 3, 'Other Data 3', 3000),
('Seeker 4', 'Name 4', 'Patronymic 4', 'Qualification 4', 4, 'Other Data 4', 4000),
('Seeker 5', 'Name 5', 'Patronymic 5', 'Qualification 5', 5, 'Other Data 5', 5000),
('Seeker 6', 'Name 6', 'Patronymic 6', 'Qualification 6', 6, 'Other Data 6', 6000),
('Seeker 7', 'Name 7', 'Patronymic 7', 'Qualification 7', 7, 'Other Data 7', 7000),
('Seeker 8', 'Name 8', 'Patronymic 8', 'Qualification 8', 8, 'Other Data 8', 8000),
('Seeker 9', 'Name 9', 'Patronymic 9', 'Qualification 9', 9, 'Other Data 9', 9000),
('Seeker 10', 'Name 10', 'Patronymic 10', 'Qualification 10', 10, 'Other Data 10', 10000);
-- Заполнение таблицы vacancy
INSERT INTO vacancy (name_vacancy, id_activity, description, get_money) VALUES
('Vacancy 1', 1, 'Description 1', 10000), ('Vacancy 2', 2, 'Description 2', 20000),
('Vacancy 3', 3, 'Description 3', 30000), ('Vacancy 4', 4, 'Description 4', 40000),
('Vacancy 5', 5, 'Description 5', 50000), ('Vacancy 6', 6, 'Description 6', 60000),
('Vacancy 7', 7, 'Description 7', 70000), ('Vacancy 8', 8, 'Description 8', 80000),
('Vacancy 9', 9, 'Description 9', 90000), ('Vacancy 10', 10, 'Description 10', 100000);
-- Заполнение таблицы dealing
INSERT INTO dealing (id_job_seekers, id_employer, post, commission) VALUES
(1, 1, 'Post 1', 1000), (2, 2, 'Post 2', 2000),
(3, 3, 'Post 3', 3000), (4, 4, 'Post 4', 4000),
(5, 5, 'Post 5', 5000), (6, 6, 'Post 6', 6000),
(7, 7, 'Post 7', 7000), (8, 8, 'Post 8', 8000),
(9, 9, 'Post 9', 9000), (10, 10, 'Post 10', 10000);