CREATE TABLE types_activity
(
    id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) NOT NULL PRIMARY KEY,
	name_activity VARCHAR(20) NOT NULL 
);
COMMENT ON TABLE types_activity IS 'Вид деятельности ';
CREATE UNIQUE INDEX ON types_activity (name_activity);