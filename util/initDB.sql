-- Drop and create the database
DROP DATABASE IF EXISTS wishticket;
CREATE DATABASE wishticket;
USE wishticket;

-- Create 'user' table
CREATE TABLE user (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- Create 'ticket' table
CREATE TABLE ticket (
    ticket_id INT AUTO_INCREMENT PRIMARY KEY, 
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    visibility ENUM('PUBLIC', 'PRIVATE') NOT NULL,
    dueDate VARCHAR(255) NOT NULL,
    creator_id INT,
    FOREIGN KEY (creator_id) REFERENCES user(user_id)
);

-- Create 'ticket_assigned' table
CREATE TABLE ticket_assigned (
    ticket_id INT,
    assigned_id INT,
    FOREIGN KEY (ticket_id) REFERENCES ticket(ticket_id),
    FOREIGN KEY (assigned_id) REFERENCES user(user_id)
);

-- Create 'ticket_status' table
CREATE TABLE ticket_status (
    ticket_id INT,
    status VARCHAR(255),
    FOREIGN KEY (ticket_id) REFERENCES ticket(ticket_id)
);

-- Create 'comment' table
CREATE TABLE comment (
    comment_id INT AUTO_INCREMENT PRIMARY KEY,
    ticket_id INT,
    creator_id INT,
    content VARCHAR(255),
    FOREIGN KEY (ticket_id) REFERENCES ticket(ticket_id),
    FOREIGN KEY (creator_id) REFERENCES user(user_id)
);

-- Create 'blocked_user' table
CREATE TABLE blocked_user (
    user_id INT,
    blocked_id INT,
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (blocked_id) REFERENCES user(user_id)
);

-- Insert sample data into 'user' table
INSERT INTO user (username, email, password_hash) 
VALUES ('admin', 'bischof.david.db@gmail.com', '$2a$10$J77xTbXaoUJmo75nZRGcQupe6grmuscQzzamo5k4s9h3NRrdNFLu6'); -- password is admin

-- Insert sample data into 'ticket' table
INSERT INTO ticket (title, description, dueDate, creator_id)
VALUES ('Test Ticket', 'This is a test ticket', 'no due date', 1);

-- Insert sample data into 'ticket_assigned' table
INSERT INTO ticket_assigned (ticket_id, assigned_id) 
VALUES (1, 1);

-- Insert sample data into 'ticket_status' table
INSERT INTO ticket_status (ticket_id, status) 
VALUES (1, 'open');

-- Insert sample data into 'comment' table
INSERT INTO comment (ticket_id, creator_id, content) 
VALUES (1, 1, 'This is a test comment');
