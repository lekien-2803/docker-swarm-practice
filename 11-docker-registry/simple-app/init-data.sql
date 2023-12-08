-- Create the 'humanresource' database
CREATE DATABASE humanresource;

-- -- Connect to the 'humanresource' database
-- \c humanresource

-- Create the 'people' table
CREATE TABLE people (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    gender VARCHAR(10),
    date_of_birth DATE
);

-- Create the 'university' table
CREATE TABLE university (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    location VARCHAR(255)
);

-- Create the 'department' table
CREATE TABLE department (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT
);

-- Create the 'salary_history' table
CREATE TABLE salary_history (
    id SERIAL PRIMARY KEY,
    person_id INT,
    salary DECIMAL(10, 2),
    effective_date DATE,
    FOREIGN KEY (person_id) REFERENCES people(id)
);

INSERT INTO people (first_name, last_name, gender, date_of_birth)
VALUES 
('John', 'Doe', 'Male', '1980-01-01'),
('Jane', 'Smith', 'Female', '1985-02-02'),
('Alice', 'Johnson', 'Female', '1990-03-03'),
('Bob', 'Brown', 'Male', '1975-04-04');

INSERT INTO university (name, location)
VALUES 
('Tech University', 'City A'),
('Science College', 'City B'),
('Arts Institute', 'City C'),
('Engineering Academy', 'City D');

INSERT INTO department (name, description)
VALUES 
('IT', 'Information Technology Department'),
('HR', 'Human Resources Department'),
('Finance', 'Finance and Accounting Department'),
('Sales', 'Sales and Marketing Department');

INSERT INTO salary_history (person_id, salary, effective_date)
VALUES 
(1, 50000.00, '2020-01-01'),
(2, 55000.00, '2020-02-01'),
(3, 52000.00, '2020-03-01'),
(4, 53000.00, '2020-04-01');
