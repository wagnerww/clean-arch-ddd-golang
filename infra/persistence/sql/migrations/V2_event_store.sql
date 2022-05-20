CREATE TABLE clientes (
    id uuid PRIMARY KEY,
    agregado VARCHAR(255),
    agregado_id VARCHAR(255),
    acao VARCHAR(255),
    payload VARCHAR(10000),
    necessita_sincronizacao BOOLEAN,
    criado_em timestamp without time zone,
    editado_em timestamp without time zone,
    apagado_em timestamp without time zone
);