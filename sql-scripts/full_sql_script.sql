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
