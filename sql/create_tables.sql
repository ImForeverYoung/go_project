
-- 1. Create Schools (Faculty-level)
CREATE TABLE schools (
    school_id SERIAL PRIMARY KEY,
    school_name VARCHAR(100) NOT NULL UNIQUE,
    dean_name VARCHAR(100),
    office_location VARCHAR(100),
    established_year INT CHECK (established_year > 1900),
    contact_email VARCHAR(100)
);

-- 2. Create Majors (Linked to Schools)
CREATE TABLE majors (
    major_id SERIAL PRIMARY KEY,
    major_name VARCHAR(50) NOT NULL,
    description TEXT,
    duration_years INT CHECK (duration_years IN (2, 3, 4, 5, 6)),
    degree_type VARCHAR(15) CHECK (degree_type IN ('Bachelor', 'Master', 'PhD')),
    school_id INT REFERENCES schools(school_id)
);

-- 3. Create Groups (Cohorts of a Major)
CREATE TABLE groups (
    group_id SERIAL PRIMARY KEY,
    group_name VARCHAR(50) NOT NULL, -- e.g., "CS-2023", "Math-2024"
    major_id INT REFERENCES majors(major_id),
    intake_year INT CHECK (intake_year > 2000)
);

-- 4. Create Professors
CREATE TABLE professors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(80) NOT NULL,
    birth_date DATE,
    department VARCHAR(100),
    title VARCHAR(50),
    hire_date DATE,
    salary DECIMAL(10, 2),
    email VARCHAR(100) UNIQUE
);

-- 5. Create Students (Linked to Groups and Majors)
CREATE TABLE students (
    student_id SERIAL PRIMARY KEY,
    student_name VARCHAR(80) NOT NULL,
    gender VARCHAR(5) CHECK (gender IN ('M', 'F')),
    birth_date DATE,
    group_id INT REFERENCES groups(group_id),
    major_id INT REFERENCES majors(major_id) -- Redundant if group implies major, but kept for flexibility
);

-- 6. Create Subjects (Linked to Professors)
CREATE TABLE subjects (
    subject_id SERIAL PRIMARY KEY,
    subject_name VARCHAR(100) NOT NULL UNIQUE,
    professor_id INT REFERENCES professors(id)
);

-- 7. Create Schedules
CREATE TABLE schedules (
    schedule_id SERIAL PRIMARY KEY,
    group_id INT REFERENCES groups(group_id),
    subject_id INT REFERENCES subjects(subject_id),
    time_slot VARCHAR(50)
);

-- 8. Create Attendance
CREATE TABLE attendance (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES students(student_id),
    subject_id INT REFERENCES subjects(subject_id),
    visit_day DATE,
    visited BOOLEAN
);

-- 9. Create Users (for Auth)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);


-- DATA INSERTION

-- Schools
INSERT INTO schools (school_name, dean_name, established_year) VALUES 
('School of Engineering and Digital Sciences', 'Dr. Smith', 2010),
('School of Sciences and Humanities', 'Dr. Johnson', 2010);

-- Majors
INSERT INTO majors (major_name, school_id, duration_years, degree_type) VALUES 
('Computer Science', 1, 4, 'Bachelor'),
('Physics', 1, 4, 'Bachelor'),
('Philosophy', 2, 4, 'Bachelor'),
('History', 2, 4, 'Bachelor');

-- Professors
INSERT INTO professors (name, department, title, email) VALUES 
('Dr. Alan Turing', 'Computer Science', 'Professor', 'alan@uni.edu'),
('Dr. Richard Feynman', 'Physics', 'Professor', 'richard@uni.edu'),
('Dr. Socrates', 'Philosophy', 'Lecturer', 'soc@uni.edu');

-- Groups (Cohorts)
INSERT INTO groups (group_name, major_id, intake_year) VALUES 
('CS-2023', 1, 2023),    -- Computer Science 1st year
('Phys-2023', 2, 2023),  -- Physics 1st year
('Phil-2022', 3, 2022);  -- Philosophy 2nd year

-- Students
INSERT INTO students (student_name, gender, birth_date, group_id, major_id) VALUES 
('Amre Jumadiyev', 'M', '2003-05-15', 1, 1), -- CS student
('Dina Kairatkyzy', 'F', '2004-02-20', 1, 1), -- CS student    
('Sanzhar Aliyev', 'M', '2003-11-10', 2, 2), -- Physics student
('Aidana Aiyasheva', 'F', '2005-01-05', 2, 2), -- Physics student   
('Uliyana Mazurenko', 'F', '2002-08-30', 3, 3); -- Philosophy student    

-- Subjects
INSERT INTO subjects (subject_name, professor_id) VALUES 
('Algorithms', 1), 
('Quantum Mechanics', 2), 
('Ethics', 3), 
('World History', 3);

-- Schedules
INSERT INTO schedules (group_id, subject_id, time_slot) VALUES 
(1, 1, '09:00 - 10:30'), -- CS-2023: Algorithms
(2, 2, '10:45 - 12:15'), -- Phys-2023: Quantum Mechanics
(3, 3, '09:00 - 10:30'); -- Phil-2022: Ethics

-- Attendance
INSERT INTO attendance (student_id, subject_id, visit_day, visited) VALUES 
(1, 1, '2026-01-11', TRUE),  
(2, 1, '2026-01-11', FALSE), 
(1, 1, '2026-01-18', TRUE);  
