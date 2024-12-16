CREATE DATABASE "DbStockManagement"

CREATE TABLE IF NOT EXISTS public.item
(
    code character varying(20) COLLATE pg_catalog."default" NOT NULL,
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    uom character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT pk_item PRIMARY KEY (code)
)

CREATE TABLE IF NOT EXISTS public.batch
(
    id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    item_code character varying(20) COLLATE pg_catalog."default" NOT NULL,
    expiry_date date NOT NULL,
    CONSTRAINT pk_batch PRIMARY KEY (id),
    CONSTRAINT fk_batch_item FOREIGN KEY (item_code)
        REFERENCES public.item (code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

CREATE TABLE IF NOT EXISTS public.entry
(
    id character varying(10) COLLATE pg_catalog."default" NOT NULL,
    tanggal date NOT NULL,
    type character varying(5) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT pk_entry PRIMARY KEY (id),
    CONSTRAINT chk_tpye CHECK (type::text = 'IN'::text OR type::text = 'OUT'::text)
)

CREATE TABLE IF NOT EXISTS public.entry_detail
(
    entry_detail_id integer NOT NULL,
    entry_id character varying(10) COLLATE pg_catalog."default" NOT NULL,
    item_code character varying(20) COLLATE pg_catalog."default" NOT NULL,
    batch_id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    expiry_date date NOT NULL,
    qty integer NOT NULL,
    CONSTRAINT pk_entry_detail PRIMARY KEY (entry_detail_id),
    CONSTRAINT fk_entry_detail_batch FOREIGN KEY (batch_id)
        REFERENCES public.batch (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_entry_detail_entry FOREIGN KEY (entry_id)
        REFERENCES public.entry (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_entry_detail_item FOREIGN KEY (item_code)
        REFERENCES public.item (code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

CREATE TABLE IF NOT EXISTS public.stock_ledger
(
    item_code character varying(20) COLLATE pg_catalog."default" NOT NULL,
    batch_id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    tanggal date NOT NULL,
    last_stock integer NOT NULL,
    qty_in integer,
    qty_out integer,
    current_stock integer NOT NULL,
    CONSTRAINT fk_stock_ledger_batch FOREIGN KEY (batch_id)
        REFERENCES public.batch (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_stock_ledger_item FOREIGN KEY (item_code)
        REFERENCES public.item (code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)