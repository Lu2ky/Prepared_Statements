DROP DATABASE IF EXISTS prepared_statements;
CREATE DATABASE prepared_statements;
USE prepared_statements;

CREATE TABLE ciudad(
	id INT PRIMARY KEY AUTO_INCREMENT,
	nombre VARCHAR(255)
); 

CREATE TABLE rol(
	id INT PRIMARY KEY AUTO_INCREMENT,
	nombre VARCHAR(255)
);

CREATE TABLE usuario(
	id INT PRIMARY KEY AUTO_INCREMENT,
	nombre VARCHAR(255),
	clave VARCHAR(255),
	ciudad_id INT,
	fecha_creacion TIMESTAMP,
	FOREIGN KEY(ciudad_id) REFERENCES ciudad(id)
);

CREATE TABLE usuario_rol(
	usuario_id INT,
	rol_id INT,
	PRIMARY KEY(usuario_id, rol_id),
	FOREIGN KEY(rol_id) REFERENCES rol(id)
);

INSERT INTO ciudad (nombre) VALUES
  ('Bogota'),
  ('Medellin'),
  ('Cali');

INSERT INTO rol (nombre) VALUES
  ('admin'),
  ('editor'),
  ('viewer');

INSERT INTO usuario (nombre, clave, ciudad_id, fecha_creacion) VALUES
  ('Ana',   SHA2('clave123', 256), 1, NOW()),
  ('Juan',   SHA2('pass456', 256),  2, NOW()),
  ('Maria',  SHA2('secret!', 256),  1, NOW()),
  ('Carlos',  SHA2('miClave9', 256), 3, NOW());

INSERT INTO usuario_rol (usuario_id, rol_id) VALUES
  (1, 1),
  (1, 2),
  (2, 2),
  (3, 3),
  (4, 1),
  (4, 3);
