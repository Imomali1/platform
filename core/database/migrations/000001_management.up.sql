CREATE TABLE REGIONS
(
    region_id integer PRIMARY KEY,
    region_name varchar(64) NOT NULL UNIQUE
);

CREATE TABLE COUNTRIES
(
    country_id varchar(2) PRIMARY KEY,
    country_name varchar(64) NOT NULL UNIQUE,
    region_id integer REFERENCES REGIONS(region_id) NOT NULL
);

CREATE TABLE LOCATIONS
(
    location_id integer PRIMARY KEY,
    street_address varchar(64) NOT NULL,
    postal_code varchar(16),
    city varchar(64),
    state_province varchar(64),
    country_id varchar(2) REFERENCES COUNTRIES(country_id) NOT NULL
);

CREATE TABLE JOBS(
                     job_id varchar(16) PRIMARY KEY,
                     job_title varchar(64) NOT NULL UNIQUE,
                     min_salary numeric(8, 2) NOT NULL,
                     max_salary numeric(8, 2) NOT NULL
);

CREATE TABLE EMPLOYEES
(
    employee_id integer PRIMARY KEY,
    first_name varchar(64) NOT NULL,
    last_name varchar(64) NOT NULL,
    email varchar(64) NOT NULL UNIQUE,
    phone_integer varchar(64) NOT NULL,
    hire_date timestamp NOT NULL,
    job_id varchar(16) REFERENCES JOBS(job_id),
    salary numeric(8, 2) NOT NULL,
    commission_pct numeric(5, 2),
    manager_id integer REFERENCES EMPLOYEES(employee_id),
    department_id integer REFERENCES DEPARTMENTS(department_id)
);

CREATE TABLE DEPARTMENTS
(
    department_id integer PRIMARY KEY,
    department_name varchar(64) NOT NULL UNIQUE,
    manager_id integer REFERENCES EMPLOYEES(employee_id),
    location_id integer REFERENCES LOCATIONS(location_id)
);

CREATE TABLE JOB_HISTORY
(
    employee_id integer REFERENCES EMPLOYEES(employee_id),
    start_date timestamp,
    PRIMARY KEY(employee_id, start_date),
    end_date timestamp,
    job_id varchar(16) REFERENCES JOBS(job_id) NOT NULL,
    department_id integer REFERENCES DEPARTMENTS(department_id) NOT NULL
);