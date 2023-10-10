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