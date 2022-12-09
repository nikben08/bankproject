CREATE SEQUENCE interests_id_seq START WITH 1;
CREATE TABLE IF NOT EXISTS public.interests
(
    id bigint NOT NULL DEFAULT nextval('interests_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    bank_id bigint,
    interest numeric,
    time_option bigint,
    credit_type bigint,
    CONSTRAINT interests_pkey PRIMARY KEY (id),
    CONSTRAINT fk_banks_interest FOREIGN KEY (bank_id)
        REFERENCES public.banks (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.interests
    OWNER to postgres;

CREATE INDEX IF NOT EXISTS idx_interests_deleted_at
    ON public.interests USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;