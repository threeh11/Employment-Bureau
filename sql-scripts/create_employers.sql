CREATE TABLE employers(
	id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) NOT NULL PRIMARY KEY,
	name_employers VARCHAR(30) NOT NULL,
	id_activity INT NOT NULL,
	FOREIGN KEY (id_activity) REFERENCES types_activity (id),
	address_employers VARCHAR(50) NOT NULL,
	number_employres VARCHAR(12) NOT NULL
);

COMMENT ON TABLE employers IS 'Работодатели';
CREATE INDEX ON employers(id_activity)