CREATE TABLE clientes (
    id uuid PRIMARY KEY,
    nome VARCHAR(255),
    email VARCHAR(255),
    ativo BOOLEAN,
    telefone INT,
    cidade VARCHAR(255),
    logradouro VARCHAR(255),
    numero INT,
    criado_em timestamp without time zone,
    editado_em timestamp without time zone,
    apagado_em timestamp without time zone
);