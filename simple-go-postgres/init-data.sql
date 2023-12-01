-- Create the 'humanresource' database
CREATE DATABASE humanresource;

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
   -- Insert data into the 'people' table
INSERT INTO people (first_name, last_name, gender, date_of_birth)
VALUES
    ('John', 'Doe', 'Male', '1990-01-15'),
    ('Jane', 'Smith', 'Female', '1985-03-20'),
    ('Michael', 'Johnson', 'Male', '1995-07-10'),
    ('Emily', 'Brown', 'Female', '1992-09-25'),
    ('David', 'Lee', 'Male', '1988-12-05');

-- Insert data into the 'university' table
INSERT INTO university (name, location)
VALUES
    ('University of ABC', 'Cityville'),
    ('XYZ University', 'Townsville'),
    ('ABC Institute of Technology', 'Tech City');

-- Insert data into the 'department' table
INSERT INTO department (name, description)
VALUES
    ('Human Resources', 'Manage personnel and hiring'),
    ('Finance', 'Manage financial transactions'),
    ('Computer Science', 'Teaching computer science courses');

-- Insert data into the 'salary_history' table
INSERT INTO salary_history (person_id, salary, effective_date)
VALUES
    (1, 55000.00, '2022-01-01'),
    (2, 60000.00, '2022-01-01'),
    (3, 58000.00, '2022-01-01'),
    (4, 62000.00, '2022-01-01'),
    (5, 53000.00, '2022-01-01');