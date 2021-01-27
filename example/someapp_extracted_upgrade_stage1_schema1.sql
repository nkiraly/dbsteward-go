-- someapp_extracted_upgrade_stage1_schema1.sql
-- DBSteward stage 1 structure additions and modifications - generated Wed, 27 Jan 2021 15:18:31 -0500
-- Old definition: someapp_v2_composite.xml
-- New definition someapp_extracted_composite.xml

BEGIN;


-- SQL STAGE STAGE1BEFORE COMMANDS

DROP VIEW IF EXISTS public.group_list_view;

CREATE OR REPLACE FUNCTION public.destroy_session(character varying) RETURNS void
AS $_$
  DELETE FROM session_information WHERE session_id=$1;
$_$
LANGUAGE sql
VOLATILE;

ALTER FUNCTION public.destroy_session(character varying) OWNER TO pgsql;

COMMENT ON FUNCTION public.destroy_session(character varying) IS 'Deletes session data from the database';

ALTER TABLE public.user_status_list DROP CONSTRAINT user_status_list_pkey;

ALTER TABLE public.user_access_level DROP CONSTRAINT user_access_level_pkey;

ALTER TABLE public.group_list DROP CONSTRAINT group_list_pkey;

ALTER TABLE public.sql_user DROP CONSTRAINT sql_user_user_status_list_id_fkey;

ALTER TABLE public.sql_user DROP CONSTRAINT sql_user_user_access_level_id_fkey;

ALTER TABLE public.sql_user DROP CONSTRAINT sql_user_pkey;

ALTER TABLE _p_public_sql_user.partition_0 DROP CONSTRAINT sql_user_p_0_chk;

ALTER TABLE _p_public_sql_user.partition_0 DROP CONSTRAINT p0_user_status_list_id_fk;

ALTER TABLE _p_public_sql_user.partition_0 DROP CONSTRAINT p0_user_access_level_id_fk;

ALTER TABLE _p_public_sql_user.partition_0 DROP CONSTRAINT partition_0_pkey;

ALTER TABLE _p_public_sql_user.partition_1 DROP CONSTRAINT sql_user_p_1_chk;

ALTER TABLE _p_public_sql_user.partition_1 DROP CONSTRAINT p1_user_status_list_id_fk;

ALTER TABLE _p_public_sql_user.partition_1 DROP CONSTRAINT p1_user_access_level_id_fk;

ALTER TABLE _p_public_sql_user.partition_1 DROP CONSTRAINT partition_1_pkey;

ALTER TABLE _p_public_sql_user.partition_2 DROP CONSTRAINT sql_user_p_2_chk;

ALTER TABLE _p_public_sql_user.partition_2 DROP CONSTRAINT p2_user_status_list_id_fk;

ALTER TABLE _p_public_sql_user.partition_2 DROP CONSTRAINT p2_user_access_level_id_fk;

ALTER TABLE _p_public_sql_user.partition_2 DROP CONSTRAINT partition_2_pkey;

ALTER TABLE _p_public_sql_user.partition_3 DROP CONSTRAINT sql_user_p_3_chk;

ALTER TABLE _p_public_sql_user.partition_3 DROP CONSTRAINT p3_user_status_list_id_fk;

ALTER TABLE _p_public_sql_user.partition_3 DROP CONSTRAINT p3_user_access_level_id_fk;

ALTER TABLE _p_public_sql_user.partition_3 DROP CONSTRAINT partition_3_pkey;

ALTER TABLE public.session_information DROP CONSTRAINT session_information_user_id_fkey;

ALTER TABLE public.session_information DROP CONSTRAINT session_information_pkey;

ALTER TABLE public.group_list
  SET WITHOUT OIDS;

ALTER TABLE public.group_list
  /* changing from type bigserial */
  ALTER COLUMN group_list_id TYPE serial,
  ALTER COLUMN group_visible SET DEFAULT true;

ALTER TABLE public.group_list
  ADD CONSTRAINT group_list_pkey PRIMARY KEY (group_list_id);

ALTER TABLE public.user_access_level
  SET WITHOUT OIDS;

ALTER TABLE public.user_access_level
  /* changing from type int */
  ALTER COLUMN user_access_level_id TYPE integer,
  /* changing from type varchar(10) */
  ALTER COLUMN user_access_level TYPE character varying(10);

ALTER TABLE public.user_access_level
  ADD CONSTRAINT user_access_level_pkey PRIMARY KEY (user_access_level_id);

ALTER TABLE public.user_status_list
  SET WITHOUT OIDS;

ALTER TABLE public.user_status_list
  /* changing from type int */
  ALTER COLUMN user_status_list_id TYPE integer;

ALTER TABLE public.user_status_list
  ADD CONSTRAINT user_status_list_pkey PRIMARY KEY (user_status_list_id);

ALTER TABLE _p_public_sql_user.partition_0
  SET WITHOUT OIDS;

ALTER TABLE _p_public_sql_user.partition_0
  ADD COLUMN user_id serial,
  ADD COLUMN user_name character varying(40),
  ADD COLUMN password text,
  ADD COLUMN somecol1 text,
  ADD COLUMN import_id character varying(32),
  ADD COLUMN register_date timestamp with time zone,
  ADD COLUMN user_status_list_id integer,
  ADD COLUMN email text,
  ADD COLUMN user_access_level_id integer DEFAULT 1;

UPDATE _p_public_sql_user.partition_0
SET user_access_level_id = DEFAULT
WHERE user_access_level_id IS NULL;

