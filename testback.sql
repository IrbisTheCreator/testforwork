toc.dat                                                                                             0000600 0004000 0002000 00000002326 15000440000 0014421 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        PGDMP                       }            testdb    17.2    17.2     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false         �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false         �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false         �           1262    16723    testdb    DATABASE     z   CREATE DATABASE testdb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';
    DROP DATABASE testdb;
                     postgres    false         �          0    16725    users 
   TABLE DATA           8   COPY public.users (user_id, email, refresh) FROM stdin;
    public               postgres    false    218       4844.dat �           0    0    users_user_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.users_user_id_seq', 10, true);
          public               postgres    false    217                                                                                                                                                                                                                                                                                                                  4844.dat                                                                                            0000600 0004000 0002000 00000001033 15000440000 0014231 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        1	alex.petrov@example.com	\N
6	anna.fedorova@platform.co	\N
7	ivan.nikitin@cloud.app	\N
8	elena.sokolova@data.tech	\N
9	pavel.morozov@digital.ai	\N
10	ekaterina.pavlova@network.biz	\N
3	sergey.smirnov@domain.org	$2a$10$47nSA2uaeO.i7.hrVNAXE.7aZr3M8hJSNAeaMaBHwiu4vrFQjhNs.
4	olga.kuznetsova@service.net	$2a$10$BDVVTFJ00aFZx2IhY39EpOmT/LqScpYGEH9SmEdar5DIXYlU30GlS
2	maria.ivanova@mail.com	$2a$10$4aTueWMo80C.mQ97TKRy4udDVe3jU9zyNnbrU5TiiqwteNnXRZI9a
5	dmitry.volkov@web.io	$2a$10$cCzBj7YAfcZZaXKyjI37kORN5ZC39uY1oV0X3z2fPCXa2samu56Gq
\.


                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     restore.sql                                                                                         0000600 0004000 0002000 00000003222 15000440000 0015342 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        --
-- NOTE:
--
-- File paths need to be edited. Search for $$PATH$$ and
-- replace it with the path to the directory containing
-- the extracted data files.
--
--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2
-- Dumped by pg_dump version 17.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE testdb;
--
-- Name: testdb; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE testdb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';


ALTER DATABASE testdb OWNER TO postgres;

\connect testdb

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (user_id, email, refresh) FROM stdin;
\.
COPY public.users (user_id, email, refresh) FROM '$$PATH$$/4844.dat';

--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_user_id_seq', 10, true);


--
-- PostgreSQL database dump complete
--

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              