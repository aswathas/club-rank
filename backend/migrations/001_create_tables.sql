-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(50) NOT NULL CHECK (role IN ('super_admin', 'club_admin', 'tech_head')),
    club_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create clubs table
CREATE TABLE clubs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    logo TEXT,
    subdomain VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create domains table
CREATE TABLE domains (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    club_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_club
        FOREIGN KEY (club_id)
        REFERENCES clubs(id)
);

-- Create students table
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    reg_no VARCHAR(50) UNIQUE NOT NULL,
    domain_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_domain
        FOREIGN KEY (domain_id)
        REFERENCES domains(id)
        ON DELETE CASCADE,
    CONSTRAINT check_reg_no_format
        CHECK (reg_no ~ '^RA[0-9]{13}$')
);

-- Create scores table
CREATE TABLE scores (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    domain_id INT NOT NULL,
    score INT NOT NULL,
    month DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_student
        FOREIGN KEY (student_id)
        REFERENCES students(id),
    CONSTRAINT fk_domain
        FOREIGN KEY (domain_id)
        REFERENCES domains(id)
);