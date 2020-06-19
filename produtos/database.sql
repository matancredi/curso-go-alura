create table produtos (
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco decimal,
    quantidade integer
);

insert into produtos (nome, descricao, preco, quantidade) VALUES
('Camiseta', 'dos Beatles', 29, 10),
('Fone', 'com microfone', 99, 50);