USE go_course;  

CREATE TABLE posts (
    id int not null auto_increment, 
    title varchar(50) not null, 
    body text,
    primary key(id)
    );

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;