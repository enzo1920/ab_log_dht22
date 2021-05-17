# ab_log_inserter

##Create db and table

```
 CREATE DATABASE ab_log_db;

CREATE TABLE weather (
    w_id serial not null primary key,
    temp_val float  NOT NULL,
	hum_val float  NOT NULL,
    w_date  timestamp default NULL
);
```
