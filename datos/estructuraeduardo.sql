
CREATE TABLE empleados (
    codigo character varying(15) PRIMARY KEY,
    nombre character varying(50) NOT NULL,
	fecha date NOT NULL,
	contrato character varying(2) NOT NULL
);

CREATE TABLE contratos (
    codigo character varying(2) PRIMARY KEY,
    nombre character varying(50) NOT NULL
);

