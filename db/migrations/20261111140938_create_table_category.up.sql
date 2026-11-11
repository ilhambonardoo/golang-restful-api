-- Active: 1784534504950@@127.0.0.1@3307@golang_database_migration
CREATE TABLE category (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;