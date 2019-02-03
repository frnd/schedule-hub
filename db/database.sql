--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: schedulehub; Type: DATABASE; Schema: -; Owner: postgres
--

\connect schedulehub

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


CREATE FUNCTION created_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

BEGIN
	NEW.updated_at = EXTRACT(EPOCH FROM NOW());
	NEW.created_at = EXTRACT(EPOCH FROM NOW());
    RETURN NEW;
END;

$$;


ALTER FUNCTION public.created_at_column() OWNER TO postgres;

--
-- TOC entry 190 (class 1255 OID 36646)
-- Name: update_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION update_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

BEGIN
    NEW.updated_at = EXTRACT(EPOCH FROM NOW());
    RETURN NEW;
END;

$$;


ALTER FUNCTION public.update_at_column() OWNER TO postgres;


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: employee; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--

CREATE TABLE employee (
    id integer NOT NULL,
    name character varying,
    updated_at integer,
    created_at integer
);


ALTER TABLE employee OWNER TO postgres;

--
-- Name: employee_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE employee_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE employee_id_seq OWNER TO postgres;

--
-- Name: employee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE employee_id_seq OWNED BY employee.id;

--
-- Name: project; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--

CREATE TABLE project (
    id integer NOT NULL,
    key character varying,
    name character varying,
    description text,
    updated_at integer,
    created_at integer
);


ALTER TABLE project OWNER TO postgres;

--
-- Name: project_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE project_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE project_id_seq OWNER TO postgres;

--
-- Name: project_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE project_id_seq OWNED BY project.id;

--
-- Name: project; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--

CREATE TABLE assignment (
    project_id integer NOT NULL,
    employee_id integer NOT NULL,
    date integer,
    pct integer NOT NULL,
    updated_at integer,
    created_at integer
);


--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--

CREATE TABLE "user" (
    id integer NOT NULL,
    email character varying,
    password character varying,
    name character varying,
    updated_at integer,
    created_at integer
);


ALTER TABLE "user" OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_id_seq OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY project ALTER COLUMN id SET DEFAULT nextval('project_id_seq'::regclass);

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY employee ALTER COLUMN id SET DEFAULT nextval('employee_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY "user" (id, email, password, name, updated_at, created_at) FROM stdin;
\.


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_id_seq', 1, false);


--
-- Name: project_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY project
    ADD CONSTRAINT project_id PRIMARY KEY (id),
    ADD CONSTRAINT project_key UNIQUE (key);

--
-- Name: employee_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY employee
    ADD CONSTRAINT employee_id PRIMARY KEY (id);

--
-- Name: assignment_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY assignment
    ADD CONSTRAINT assignment_pk PRIMARY KEY (project_id, employee_id, "date");

--
-- Name: user_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_id PRIMARY KEY (id);


--
-- Name: assignment_project_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY assignment
    ADD CONSTRAINT assignment_project_id FOREIGN KEY (project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE CASCADE;

--
-- Name: assignment_employee_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY assignment
    ADD CONSTRAINT assignment_employee_id FOREIGN KEY (employee_id) REFERENCES employee(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- TOC entry 2284 (class 2620 OID 36647)
-- Name: project create_project_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_project_created_at BEFORE INSERT ON project FOR EACH ROW EXECUTE PROCEDURE created_at_column();

-- TOC entry 2284 (class 2620 OID 36647)
-- Name: project create_employee_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_employee_created_at BEFORE INSERT ON employee FOR EACH ROW EXECUTE PROCEDURE created_at_column();

--
-- TOC entry 2286 (class 2620 OID 36653)
-- Name: user create_user_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_user_created_at BEFORE INSERT ON "user" FOR EACH ROW EXECUTE PROCEDURE created_at_column();

--
-- TOC entry 2286 (class 2620 OID 36653)
-- Name: user create_user_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_assignment_created_at BEFORE INSERT ON assignment FOR EACH ROW EXECUTE PROCEDURE created_at_column();


--
-- TOC entry 2285 (class 2620 OID 36648)
-- Name: project update_project_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_project_updated_at BEFORE UPDATE ON project FOR EACH ROW EXECUTE PROCEDURE update_at_column();

--
-- TOC entry 2285 (class 2620 OID 36648)
-- Name: project update_employee_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_employee_updated_at BEFORE UPDATE ON employee FOR EACH ROW EXECUTE PROCEDURE update_at_column();


--
-- TOC entry 2285 (class 2620 OID 36648)
-- Name: project update_employee_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_assignment_updated_at BEFORE UPDATE ON assignment FOR EACH ROW EXECUTE PROCEDURE update_at_column();

--
-- TOC entry 2287 (class 2620 OID 36654)
-- Name: user update_user_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_user_updated_at BEFORE UPDATE ON "user" FOR EACH ROW EXECUTE PROCEDURE update_at_column();

--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--
