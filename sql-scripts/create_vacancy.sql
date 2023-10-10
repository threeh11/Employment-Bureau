CREATE TABLE vacancy(
	id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) NOT NULL PRIMARY KEY,
	name_vacancy VARCHAR(30) NOT NULL,
	id_activity INT NOT NULL,
	FOREIGN KEY (id_activity) REFERENCES types_activity (id),
	description VARCHAR(10000) NOT NULL,
	get_money INT NOT NULL
);

COMMENT ON TABLE vacancy IS 'Вакансии';
CREATE INDEX ON vacancy(id_activity);
CREATE INDEX ON vacancy(name_vacancy);
CREATE INDEX ON vacancy(get_money);