-- Database: golangdb

-- DROP DATABASE golangdb;

CREATE DATABASE golangdb
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Portuguese_Brazil.1252'
    LC_CTYPE = 'Portuguese_Brazil.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;


    -- Table: public.pessoa

-- DROP TABLE public.pessoa;

CREATE TABLE public.pessoa
(
    name character varying(99) COLLATE pg_catalog."default",
    age integer,
    telephone character varying(99) COLLATE pg_catalog."default",
    id integer NOT NULL,
    CONSTRAINT pessoa_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.pessoa
    OWNER to postgres;