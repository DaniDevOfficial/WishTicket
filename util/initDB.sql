DROP TABLE IF EXISTS wishticket;
CREATE TABLE wishticket;
USE wishticket;

CREATE TABLE ticket (
    ticket_id INT AUTO_INCREMENT PRIMARY KEY, 
    title VARCHAR(255),
    description VARCHAR(255),
    creator_id INT 
)

CREATE TABLE ticket_asigned (
    ticket_id INT,
    asigned_id INT
)

CREATE TABLE ticket_status (
    ticket_id INT,
    status VARCHAR(255)
)

CREATE TABLE comment (
    comment_id INT AUTO_INCREMENT PRIMARY KEY,
    ticket_id INT,
    creator_id INT,
    content VARCHAR(255)
)

CREATE TABLE user (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    passwordHash VARCHAR(255),
)

CREATE TABLE blocked_user (
    user_id INT,
    blocked_id INT
)

INSERT INTO user (username, email, passwordHash) VALUES ('admin', 'bischof.david.db@gmail.com', 'admin')

INSERT INTO ticket (title, description, creator_id) VALUES ('Test Ticket', 'This is a test ticket', 1)

INSERT INTO ticket_asigned (ticket_id, asigned_id) VALUES (1, 1)

INSERT INTO ticket_status (ticket_id, status) VALUES (1, 'open')

INSERT INTO comment (ticket_id, creator_id, content) VALUES (1, 1, 'This is a test comment')