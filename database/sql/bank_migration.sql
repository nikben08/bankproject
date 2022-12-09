CREATE SEQUENCE banks_id_seq START WITH 1;
CREATE TABLE IF NOT EXISTS public.banks
(
    id bigint NOT NULL DEFAULT nextval('banks_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    CONSTRAINT banks_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.banks
    OWNER to postgres;


CREATE INDEX IF NOT EXISTS idx_banks_deleted_at
    ON public.banks USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;