CREATE TABLE IF NOT EXISTS products (
  id VARCHAR(20) PRIMARY KEY,
  name VARCHAR(20) NOT NULL,
  description VARCHAR(80) NOT NULL,
  price DECIMAL(8, 2) NOT NULL
);

INSERT INTO products (
  id, name, description, price
) VALUES
('1', 'Arroz', 'Branco', 6.49),
('2', 'Arroz', 'Preto', 6.49),
('3', 'Feijão', 'Branco', 9.49),
('4', 'Feijão', 'Preto', 9.49),
('5', 'Macarrão', 'Ninho', 2.49),
('6', 'Macarrão', 'Penne', 2.49),
('7', 'Macarrão', 'Parafuso', 2.49),
('8', 'Macarrão', 'Espaguete', 2.49),
('9', 'Macarrão', 'Parafuso', 2.49),
('10', 'Ovos', '1/2 bandeija', 8.49),
('11', 'Ovos', 'Bandeija inteira', 16.99),
('12', 'Pão', 'Diversos', 0.49);
