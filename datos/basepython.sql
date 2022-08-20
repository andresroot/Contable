--
-- PostgreSQL database dump
--

-- Dumped from database version 12.5
-- Dumped by pg_dump version 12.5

-- Started on 2021-02-04 18:18:48

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 17597)
-- Name: auth_group; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth_group (
    id integer NOT NULL,
    name character varying(150) NOT NULL
);


ALTER TABLE public.auth_group OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 17595)
-- Name: auth_group_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auth_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_group_id_seq OWNER TO postgres;

--
-- TOC entry 3268 (class 0 OID 0)
-- Dependencies: 208
-- Name: auth_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auth_group_id_seq OWNED BY public.auth_group.id;


--
-- TOC entry 211 (class 1259 OID 17607)
-- Name: auth_group_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth_group_permissions (
    id integer NOT NULL,
    group_id integer NOT NULL,
    permission_id integer NOT NULL
);


ALTER TABLE public.auth_group_permissions OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 17605)
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auth_group_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_group_permissions_id_seq OWNER TO postgres;

--
-- TOC entry 3269 (class 0 OID 0)
-- Dependencies: 210
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auth_group_permissions_id_seq OWNED BY public.auth_group_permissions.id;


--
-- TOC entry 207 (class 1259 OID 17589)
-- Name: auth_permission; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth_permission (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    content_type_id integer NOT NULL,
    codename character varying(100) NOT NULL
);


ALTER TABLE public.auth_permission OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 17587)
-- Name: auth_permission_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auth_permission_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_permission_id_seq OWNER TO postgres;

--
-- TOC entry 3270 (class 0 OID 0)
-- Dependencies: 206
-- Name: auth_permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auth_permission_id_seq OWNED BY public.auth_permission.id;


--
-- TOC entry 213 (class 1259 OID 17615)
-- Name: auth_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth_user (
    id integer NOT NULL,
    password character varying(128) NOT NULL,
    last_login timestamp with time zone,
    is_superuser boolean NOT NULL,
    username character varying(150) NOT NULL,
    first_name character varying(150) NOT NULL,
    last_name character varying(150) NOT NULL,
    email character varying(254) NOT NULL,
    is_staff boolean NOT NULL,
    is_active boolean NOT NULL,
    date_joined timestamp with time zone NOT NULL
);


ALTER TABLE public.auth_user OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 17625)
-- Name: auth_user_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth_user_groups (
    id integer NOT NULL,
    user_id integer NOT NULL,
    group_id integer NOT NULL
);


ALTER TABLE public.auth_user_groups OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 17623)
-- Name: auth_user_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auth_user_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_user_groups_id_seq OWNER TO postgres;

--
-- TOC entry 3271 (class 0 OID 0)
-- Dependencies: 214
-- Name: auth_user_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auth_user_groups_id_seq OWNED BY public.auth_user_groups.id;


--
-- TOC entry 212 (class 1259 OID 17613)
-- Name: auth_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auth_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_user_id_seq OWNER TO postgres;

--
-- TOC entry 3272 (class 0 OID 0)
-- Dependencies: 212
-- Name: auth_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auth_user_id_seq OWNED BY public.auth_user.id;


--
-- TOC entry 217 (class 1259 OID 17633)
-- Name: auth_user_user_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth_user_user_permissions (
    id integer NOT NULL,
    user_id integer NOT NULL,
    permission_id integer NOT NULL
);


ALTER TABLE public.auth_user_user_permissions OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 17631)
-- Name: auth_user_user_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auth_user_user_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_user_user_permissions_id_seq OWNER TO postgres;

--
-- TOC entry 3273 (class 0 OID 0)
-- Dependencies: 216
-- Name: auth_user_user_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auth_user_user_permissions_id_seq OWNED BY public.auth_user_user_permissions.id;


--
-- TOC entry 221 (class 1259 OID 17734)
-- Name: contabilidad_bodega; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_bodega (
    "bodegaCodigo" character varying(2) NOT NULL,
    "bodegaNombre" character varying(50) NOT NULL
);


ALTER TABLE public.contabilidad_bodega OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 17739)
-- Name: contabilidad_centro; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_centro (
    "centroCodigo" character varying(3) NOT NULL,
    "centroNombre" character varying(50) NOT NULL
);


ALTER TABLE public.contabilidad_centro OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 17744)
-- Name: contabilidad_ciudad; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_ciudad (
    "ciudadCodigo" character varying(5) NOT NULL,
    "ciudadCodigoCiudad" character varying(3) NOT NULL,
    "ciudadCodigoDepartamento" character varying(2) NOT NULL,
    "ciudadNombre" character varying(150) NOT NULL,
    "ciudadNombreCiudad" character varying(150) NOT NULL,
    "ciudadNombreDepartamento" character varying(150) NOT NULL
);


ALTER TABLE public.contabilidad_ciudad OWNER TO postgres;

--
-- TOC entry 250 (class 1259 OID 17880)
-- Name: contabilidad_compra; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_compra (
    "compraCodigo" character varying(20) NOT NULL,
    "compraFecha" date NOT NULL,
    "compraVence" date NOT NULL,
    "compraHora" time without time zone NOT NULL,
    "compraDescuento" character varying(20) NOT NULL,
    "compraSubtotal" character varying(20) NOT NULL,
    "compraSubtotalIva19" character varying(20) NOT NULL,
    "compraSubtotalIva5" character varying(20) NOT NULL,
    "compraSubtotalIva0" character varying(20) NOT NULL,
    "compraSubtotalBase19" character varying(20) NOT NULL,
    "compraSubtotalBase5" character varying(20) NOT NULL,
    "compraSubtotalBase0" character varying(20) NOT NULL,
    "compraTotalIva" character varying(20) NOT NULL,
    "compraTotal" character varying(20) NOT NULL,
    "compraPorcentajeRetencionFuente" character varying(6) NOT NULL,
    "compraTotalRetencionFuente" character varying(20) NOT NULL,
    "compraPorcentajeRetencionIca" character varying(6) NOT NULL,
    "compraTotalRetencionIca" character varying(20) NOT NULL,
    "compraNeto" character varying(20) NOT NULL,
    "compraItems" character varying(6) NOT NULL,
    "compraFormaDePago_id" character varying(2) NOT NULL,
    "compraMedioDePago_id" character varying(3) NOT NULL,
    "compraResolucion_id" character varying(2) NOT NULL,
    "compraTerceroCodigo_id" character varying(15) NOT NULL,
    "compraVendedor_id" character varying(4) NOT NULL
);


ALTER TABLE public.contabilidad_compra OWNER TO postgres;

--
-- TOC entry 249 (class 1259 OID 17874)
-- Name: contabilidad_compradetalle; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_compradetalle (
    "compraId" integer NOT NULL,
    "compraCodigo" character varying(20) NOT NULL,
    "compraFila" bigint NOT NULL,
    "compraCantidad" character varying(20) NOT NULL,
    "compraPrecio" character varying(20) NOT NULL,
    "compraDescuento" character varying(6) NOT NULL,
    "compraMontoDescuento" character varying(20) NOT NULL,
    "compraSiGratis" character varying(2) NOT NULL,
    "compraSubtotal" character varying(20) NOT NULL,
    "compraSubtotalDescuento" character varying(20) NOT NULL,
    "compraPagina" character varying(4) NOT NULL,
    "compraBodega_id" character varying(2) NOT NULL,
    "compraProductoCodigo_id" character varying(15) NOT NULL
);


ALTER TABLE public.contabilidad_compradetalle OWNER TO postgres;

--
-- TOC entry 248 (class 1259 OID 17872)
-- Name: contabilidad_compradetalle_compraId_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."contabilidad_compradetalle_compraId_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."contabilidad_compradetalle_compraId_seq" OWNER TO postgres;

--
-- TOC entry 3274 (class 0 OID 0)
-- Dependencies: 248
-- Name: contabilidad_compradetalle_compraId_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."contabilidad_compradetalle_compraId_seq" OWNED BY public.contabilidad_compradetalle."compraId";


--
-- TOC entry 247 (class 1259 OID 17866)
-- Name: contabilidad_comprobantecabecera; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_comprobantecabecera (
    "comprobanteCodigo" integer NOT NULL,
    "comprobanteConsecutivo" bigint NOT NULL,
    "comprobantePrefijo" character varying(5) NOT NULL,
    "comprobanteResolucion" character varying(20) NOT NULL,
    "comprobanteFecha" date NOT NULL,
    "comprobanteAño" bigint NOT NULL,
    "comprobanteBloqueo" boolean NOT NULL,
    "documentoCodigo_id" character varying(2) NOT NULL
);


ALTER TABLE public.contabilidad_comprobantecabecera OWNER TO postgres;

--
-- TOC entry 246 (class 1259 OID 17864)
-- Name: contabilidad_comprobantecabecera_comprobanteCodigo_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."contabilidad_comprobantecabecera_comprobanteCodigo_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."contabilidad_comprobantecabecera_comprobanteCodigo_seq" OWNER TO postgres;

--
-- TOC entry 3275 (class 0 OID 0)
-- Dependencies: 246
-- Name: contabilidad_comprobantecabecera_comprobanteCodigo_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."contabilidad_comprobantecabecera_comprobanteCodigo_seq" OWNED BY public.contabilidad_comprobantecabecera."comprobanteCodigo";


--
-- TOC entry 245 (class 1259 OID 17858)
-- Name: contabilidad_comprobantemovimiento; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_comprobantemovimiento (
    "comprobanteMovimientoId" integer NOT NULL,
    "comprobanteConsecutivo" bigint NOT NULL,
    "comprobantePrefijo" character varying(5) NOT NULL,
    "movimientoFila" bigint NOT NULL,
    "movimientoConcepto" character varying(120) NOT NULL,
    "movimientoDocumento" character varying(20) NOT NULL,
    "movimientoFecha" timestamp with time zone NOT NULL,
    "movimientoDebito" numeric(12,2) NOT NULL,
    "movimientoCredito" numeric(12,2) NOT NULL,
    "centroCodigo_id" character varying(3) NOT NULL,
    "cuentaCodigo_id" character varying(10) NOT NULL,
    "documentoCodigo_id" character varying(2) NOT NULL,
    "terceroCodigo_id" character varying(15) NOT NULL
);


ALTER TABLE public.contabilidad_comprobantemovimiento OWNER TO postgres;

--
-- TOC entry 244 (class 1259 OID 17856)
-- Name: contabilidad_comprobantemovimiento_comprobanteMovimientoId_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."contabilidad_comprobantemovimiento_comprobanteMovimientoId_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."contabilidad_comprobantemovimiento_comprobanteMovimientoId_seq" OWNER TO postgres;

--
-- TOC entry 3276 (class 0 OID 0)
-- Dependencies: 244
-- Name: contabilidad_comprobantemovimiento_comprobanteMovimientoId_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."contabilidad_comprobantemovimiento_comprobanteMovimientoId_seq" OWNED BY public.contabilidad_comprobantemovimiento."comprobanteMovimientoId";


--
-- TOC entry 224 (class 1259 OID 17749)
-- Name: contabilidad_cuenta; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_cuenta (
    "cuentaCodigo" character varying(10) NOT NULL,
    "cuentaNombre" character varying(80) NOT NULL,
    "cuentaAuto" character varying(5) NOT NULL,
    "cuentaNivel" character varying(1) NOT NULL,
    "cuentaTercero" character varying(2) NOT NULL,
    "cuentaCentro" character varying(2) NOT NULL,
    "cuentaFactura" character varying(2) NOT NULL,
    "cuentaInforme" character varying(3) NOT NULL,
    "cuentaContra" character varying(8) NOT NULL,
    "cuentaInteres" character varying(8) NOT NULL,
    "cuentaCuota" character varying(2) NOT NULL,
    "cuentaIntereses" character varying(2) NOT NULL,
    "cuentaTipo" character varying(6) NOT NULL,
    "cuentaGrupo" character varying(2) NOT NULL,
    "cuentaContraNombre" character varying(80) NOT NULL,
    "cuentaInteresNombre" character varying(80) NOT NULL
);


ALTER TABLE public.contabilidad_cuenta OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 17756)
-- Name: contabilidad_cuota; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_cuota (
    "cuotaId" integer NOT NULL,
    "cuotaPrestamo" character varying(10) NOT NULL,
    "cuotaPlazo" character varying(10) NOT NULL,
    "cuotaInteres" character varying(10) NOT NULL,
    "cuotaFecha" timestamp with time zone NOT NULL,
    "cuotaPago" character varying(20) NOT NULL,
    "cuotaTotalInteres" character varying(20) NOT NULL,
    "cuotaTotalPago" character varying(20) NOT NULL
);


ALTER TABLE public.contabilidad_cuota OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 17754)
-- Name: contabilidad_cuota_cuotaId_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."contabilidad_cuota_cuotaId_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."contabilidad_cuota_cuotaId_seq" OWNER TO postgres;

--
-- TOC entry 3277 (class 0 OID 0)
-- Dependencies: 225
-- Name: contabilidad_cuota_cuotaId_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."contabilidad_cuota_cuotaId_seq" OWNED BY public.contabilidad_cuota."cuotaId";


--
-- TOC entry 243 (class 1259 OID 17850)
-- Name: contabilidad_cuotadetalle; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_cuotadetalle (
    id integer NOT NULL,
    "cuotaFila" bigint NOT NULL,
    "cuotaDetalleFecha" timestamp with time zone NOT NULL,
    "cuotaDetalleCuota" character varying(50) NOT NULL,
    "cuotaDetalleInteres" character varying(50) NOT NULL,
    "cuotaDetalleCapital" character varying(50) NOT NULL,
    "cuotaDetalleSaldo" character varying(50) NOT NULL,
    "cuotaCodigo_id" integer NOT NULL
);


ALTER TABLE public.contabilidad_cuotadetalle OWNER TO postgres;

--
-- TOC entry 242 (class 1259 OID 17848)
-- Name: contabilidad_cuotadetalle_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.contabilidad_cuotadetalle_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.contabilidad_cuotadetalle_id_seq OWNER TO postgres;

--
-- TOC entry 3278 (class 0 OID 0)
-- Dependencies: 242
-- Name: contabilidad_cuotadetalle_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.contabilidad_cuotadetalle_id_seq OWNED BY public.contabilidad_cuotadetalle.id;


--
-- TOC entry 227 (class 1259 OID 17762)
-- Name: contabilidad_documento; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_documento (
    "documentoCodigo" character varying(2) NOT NULL,
    "documentoNombre" character varying(50) NOT NULL,
    "documentoConsecutivo" character varying(2) NOT NULL,
    "documentoInicial" character varying(6)
);


ALTER TABLE public.contabilidad_documento OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 17767)
-- Name: contabilidad_documentoidentificacion; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_documentoidentificacion (
    "documentoIdentificacionCodigo" character varying(2) NOT NULL,
    "documentoIdentificacionNombre" character varying(100) NOT NULL
);


ALTER TABLE public.contabilidad_documentoidentificacion OWNER TO postgres;

--
-- TOC entry 241 (class 1259 OID 17840)
-- Name: contabilidad_empresa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_empresa (
    "empresaCodigo" character varying(15) NOT NULL,
    "empresaDv" character varying(1) NOT NULL,
    "empresaNombre" character varying(100) NOT NULL,
    "empresaIva" character varying(20) NOT NULL,
    "empresaReteIva" character varying(20) NOT NULL,
    "empresaDireccion" character varying(80) NOT NULL,
    "empresaActividadIca" character varying(60) NOT NULL,
    "empresaTelefono1" character varying(20) NOT NULL,
    "empresaTelefono2" character varying(20) NOT NULL,
    "empresaEmail1" character varying(100) NOT NULL,
    "empresaEmail2" character varying(100) NOT NULL,
    "empresaActiva" character varying(2) NOT NULL,
    "empresaLicencia" character varying(20) NOT NULL,
    "empresaRepresentanteDv" character varying(1) NOT NULL,
    "empresaRepresentanteNombre" character varying(100),
    "empresaContadorDv" character varying(1) NOT NULL,
    "empresaContadorNombre" character varying(100),
    "empresaRevisorDv" character varying(1) NOT NULL,
    "empresaRevisorNombre" character varying(100),
    "empresaCiudad_id" character varying(5) NOT NULL,
    "empresaContadorNit_id" character varying(15) NOT NULL,
    "empresaDocumento_id" character varying(2) NOT NULL,
    "empresaFiscal_id" character varying(10) NOT NULL,
    "empresaRegimen_id" character varying(2) NOT NULL,
    "empresaRepresentanteNit_id" character varying(15) NOT NULL,
    "empresaRevisorNit_id" character varying(15) NOT NULL,
    "empresaTipo_id" character varying(1) NOT NULL
);


ALTER TABLE public.contabilidad_empresa OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 17772)
-- Name: contabilidad_formadepago; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_formadepago (
    "formaDePagoCodigo" character varying(2) NOT NULL,
    "formaDePagoNombre" character varying(100) NOT NULL
);


ALTER TABLE public.contabilidad_formadepago OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 17777)
-- Name: contabilidad_grupo; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_grupo (
    "grupoCodigo" character varying(2) NOT NULL,
    "grupoNombre" character varying(50) NOT NULL
);


ALTER TABLE public.contabilidad_grupo OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 17782)
-- Name: contabilidad_mediodepago; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_mediodepago (
    "medioDePagoCodigo" character varying(3) NOT NULL,
    "medioDePagoNombre" character varying(100) NOT NULL
);


ALTER TABLE public.contabilidad_mediodepago OWNER TO postgres;

--
-- TOC entry 240 (class 1259 OID 17835)
-- Name: contabilidad_producto; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_producto (
    "productoCodigo" character varying(15) NOT NULL,
    "productoNombre" character varying(100) NOT NULL,
    "productoUnidad" character varying(12) NOT NULL,
    "productoIva" character varying(6) NOT NULL,
    "productoTipo" character varying(10) NOT NULL,
    "productoVenta" character varying(12) NOT NULL,
    "productoCosto" character varying(12) NOT NULL,
    "productoCantidad" character varying(12) NOT NULL,
    "productoTotal" character varying(12) NOT NULL,
    "productoSubgrupo_id" character varying(4) NOT NULL
);


ALTER TABLE public.contabilidad_producto OWNER TO postgres;

--
-- TOC entry 232 (class 1259 OID 17787)
-- Name: contabilidad_regimenfiscal; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_regimenfiscal (
    "regimenFiscalCodigo" character varying(2) NOT NULL,
    "regimenFiscalNombre" character varying(100) NOT NULL
);


ALTER TABLE public.contabilidad_regimenfiscal OWNER TO postgres;

--
-- TOC entry 239 (class 1259 OID 17830)
-- Name: contabilidad_resolucion; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_resolucion (
    "resolucionCodigo" character varying(2) NOT NULL,
    "resolucionDian" character varying(20) NOT NULL,
    "resolucionPrefijo" character varying(4) NOT NULL,
    "resolucionTipo" character varying(20) NOT NULL,
    "resolucionNombreLocal" character varying(50) NOT NULL,
    "resolucionDireccion" character varying(50) NOT NULL,
    "resolucionTelefono" character varying(20) NOT NULL,
    "resolucionInforme" character varying(50) NOT NULL,
    "resolucionFechaInicial" timestamp with time zone NOT NULL,
    "resolucionFechaFinal" timestamp with time zone NOT NULL,
    "resolucionNumeroInicial" character varying(50) NOT NULL,
    "resolucionNumeroFinal" character varying(50) NOT NULL,
    "resolucionNumeroActual" character varying(50) NOT NULL,
    "resolucionCiudad_id" character varying(5) NOT NULL
);


ALTER TABLE public.contabilidad_resolucion OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 17792)
-- Name: contabilidad_responsabilidadfiscal; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_responsabilidadfiscal (
    "responsabilidadFiscalCodigo" character varying(10) NOT NULL,
    "responsabilidadFiscalNombre" character varying(100) NOT NULL
);


ALTER TABLE public.contabilidad_responsabilidadfiscal OWNER TO postgres;

--
-- TOC entry 238 (class 1259 OID 17825)
-- Name: contabilidad_subgrupo; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_subgrupo (
    "subgrupoCodigo" character varying(4) NOT NULL,
    "subgrupoNombre" character varying(50) NOT NULL,
    "subgrupoCodigoGrupo_id" character varying(2) NOT NULL
);


ALTER TABLE public.contabilidad_subgrupo OWNER TO postgres;

--
-- TOC entry 234 (class 1259 OID 17797)
-- Name: contabilidad_tercero; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_tercero (
    "terceroCodigo" character varying(15) NOT NULL,
    "terceroDv" character varying(1) NOT NULL,
    "terceroNombre" character varying(100),
    "terceroJuridica" character varying(120) NOT NULL,
    "terceroPrimerNombre" character varying(60) NOT NULL,
    "terceroSegundoNombre" character varying(60) NOT NULL,
    "terceroPrimerApellido" character varying(60) NOT NULL,
    "terceroSegundoApellido" character varying(60) NOT NULL,
    "terceroDireccion" character varying(80) NOT NULL,
    "terceroBarrio" character varying(60) NOT NULL,
    "terceroTelefono1" character varying(20) NOT NULL,
    "terceroTelefono2" character varying(20) NOT NULL,
    "terceroEmail1" character varying(100) NOT NULL,
    "terceroEmail2" character varying(100) NOT NULL,
    "terceroContacto" character varying(60) NOT NULL,
    "terceroRut" character varying(2) NOT NULL,
    "terceroBloque" character varying(10) NOT NULL,
    "terceroPiso" character varying(10) NOT NULL,
    "terceroApartamento" character varying(10) NOT NULL,
    "terceroDescuento1" character varying(12) NOT NULL,
    "terceroDescuento2" character varying(12) NOT NULL,
    "terceroCuotaP" character varying(12) NOT NULL,
    "terceroCuota1" character varying(12) NOT NULL,
    "terceroCuota2" character varying(12) NOT NULL,
    "terceroCuota3" character varying(12) NOT NULL,
    "terceroArea" character varying(6) NOT NULL,
    "terceroFactor" character varying(5) NOT NULL,
    "terceroMatricula" character varying(20) NOT NULL,
    "terceroCatastral" character varying(20) NOT NULL,
    "terceroBanco" character varying(20) NOT NULL,
    "tercerophCodigo" character varying(15) NOT NULL,
    "tercerophDv" character varying(1) NOT NULL,
    "tercerophNombre" character varying(100),
    "terceroCiudad_id" character varying(5) NOT NULL,
    "terceroDocumento_id" character varying(2) NOT NULL,
    "terceroFiscal_id" character varying(10) NOT NULL,
    "terceroRegimen_id" character varying(2) NOT NULL,
    "terceroTipo_id" character varying(1) NOT NULL
);


ALTER TABLE public.contabilidad_tercero OWNER TO postgres;

--
-- TOC entry 235 (class 1259 OID 17805)
-- Name: contabilidad_tipoorganizacion; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_tipoorganizacion (
    "tipoOrganizacionCodigo" character varying(1) NOT NULL,
    "tipoOrganizacionNombre" character varying(100) NOT NULL
);


ALTER TABLE public.contabilidad_tipoorganizacion OWNER TO postgres;

--
-- TOC entry 237 (class 1259 OID 17815)
-- Name: contabilidad_usuario; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_usuario (
    "usuarioCodigo" character varying(4) NOT NULL,
    "usuarioDv" character varying(1) NOT NULL,
    "usuarioNombre" character varying(100),
    "usuarioTipo" character varying(50) NOT NULL,
    "usuarioClave2" character varying(15) NOT NULL,
    "usuarioClave1" character varying(15) NOT NULL,
    "usuarioCorreo1" character varying(100) NOT NULL,
    "usuarioCorreo2" character varying(100) NOT NULL,
    "usuarioNit_id" character varying(15) NOT NULL
);


ALTER TABLE public.contabilidad_usuario OWNER TO postgres;

--
-- TOC entry 236 (class 1259 OID 17810)
-- Name: contabilidad_vendedor; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contabilidad_vendedor (
    "vendedorCodigo" character varying(4) NOT NULL,
    "vendedorDv" character varying(1) NOT NULL,
    "vendedorNombre" character varying(100),
    "vendedorComision" character varying(5) NOT NULL,
    "vendedorNit_id" character varying(15) NOT NULL
);


ALTER TABLE public.contabilidad_vendedor OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 17693)
-- Name: django_admin_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.django_admin_log (
    id integer NOT NULL,
    action_time timestamp with time zone NOT NULL,
    object_id text,
    object_repr character varying(200) NOT NULL,
    action_flag smallint NOT NULL,
    change_message text NOT NULL,
    content_type_id integer,
    user_id integer NOT NULL,
    CONSTRAINT django_admin_log_action_flag_check CHECK ((action_flag >= 0))
);


ALTER TABLE public.django_admin_log OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 17691)
-- Name: django_admin_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.django_admin_log_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.django_admin_log_id_seq OWNER TO postgres;

--
-- TOC entry 3279 (class 0 OID 0)
-- Dependencies: 218
-- Name: django_admin_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.django_admin_log_id_seq OWNED BY public.django_admin_log.id;


--
-- TOC entry 205 (class 1259 OID 17579)
-- Name: django_content_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.django_content_type (
    id integer NOT NULL,
    app_label character varying(100) NOT NULL,
    model character varying(100) NOT NULL
);


ALTER TABLE public.django_content_type OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 17577)
-- Name: django_content_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.django_content_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.django_content_type_id_seq OWNER TO postgres;

--
-- TOC entry 3280 (class 0 OID 0)
-- Dependencies: 204
-- Name: django_content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.django_content_type_id_seq OWNED BY public.django_content_type.id;


--
-- TOC entry 203 (class 1259 OID 17568)
-- Name: django_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.django_migrations (
    id integer NOT NULL,
    app character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    applied timestamp with time zone NOT NULL
);


ALTER TABLE public.django_migrations OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 17566)
-- Name: django_migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.django_migrations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.django_migrations_id_seq OWNER TO postgres;

--
-- TOC entry 3281 (class 0 OID 0)
-- Dependencies: 202
-- Name: django_migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.django_migrations_id_seq OWNED BY public.django_migrations.id;


--
-- TOC entry 220 (class 1259 OID 17724)
-- Name: django_session; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.django_session (
    session_key character varying(40) NOT NULL,
    session_data text NOT NULL,
    expire_date timestamp with time zone NOT NULL
);


ALTER TABLE public.django_session OWNER TO postgres;

--
-- TOC entry 2858 (class 2604 OID 17600)
-- Name: auth_group id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_group ALTER COLUMN id SET DEFAULT nextval('public.auth_group_id_seq'::regclass);


--
-- TOC entry 2859 (class 2604 OID 17610)
-- Name: auth_group_permissions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_group_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_group_permissions_id_seq'::regclass);


--
-- TOC entry 2857 (class 2604 OID 17592)
-- Name: auth_permission id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_permission ALTER COLUMN id SET DEFAULT nextval('public.auth_permission_id_seq'::regclass);


--
-- TOC entry 2860 (class 2604 OID 17618)
-- Name: auth_user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user ALTER COLUMN id SET DEFAULT nextval('public.auth_user_id_seq'::regclass);


--
-- TOC entry 2861 (class 2604 OID 17628)
-- Name: auth_user_groups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user_groups ALTER COLUMN id SET DEFAULT nextval('public.auth_user_groups_id_seq'::regclass);


--
-- TOC entry 2862 (class 2604 OID 17636)
-- Name: auth_user_user_permissions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user_user_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_user_user_permissions_id_seq'::regclass);


--
-- TOC entry 2869 (class 2604 OID 17877)
-- Name: contabilidad_compradetalle compraId; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_compradetalle ALTER COLUMN "compraId" SET DEFAULT nextval('public."contabilidad_compradetalle_compraId_seq"'::regclass);


--
-- TOC entry 2868 (class 2604 OID 17869)
-- Name: contabilidad_comprobantecabecera comprobanteCodigo; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_comprobantecabecera ALTER COLUMN "comprobanteCodigo" SET DEFAULT nextval('public."contabilidad_comprobantecabecera_comprobanteCodigo_seq"'::regclass);


--
-- TOC entry 2867 (class 2604 OID 17861)
-- Name: contabilidad_comprobantemovimiento comprobanteMovimientoId; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_comprobantemovimiento ALTER COLUMN "comprobanteMovimientoId" SET DEFAULT nextval('public."contabilidad_comprobantemovimiento_comprobanteMovimientoId_seq"'::regclass);


--
-- TOC entry 2865 (class 2604 OID 17759)
-- Name: contabilidad_cuota cuotaId; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_cuota ALTER COLUMN "cuotaId" SET DEFAULT nextval('public."contabilidad_cuota_cuotaId_seq"'::regclass);


--
-- TOC entry 2866 (class 2604 OID 17853)
-- Name: contabilidad_cuotadetalle id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_cuotadetalle ALTER COLUMN id SET DEFAULT nextval('public.contabilidad_cuotadetalle_id_seq'::regclass);


--
-- TOC entry 2863 (class 2604 OID 17696)
-- Name: django_admin_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.django_admin_log ALTER COLUMN id SET DEFAULT nextval('public.django_admin_log_id_seq'::regclass);


--
-- TOC entry 2856 (class 2604 OID 17582)
-- Name: django_content_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.django_content_type ALTER COLUMN id SET DEFAULT nextval('public.django_content_type_id_seq'::regclass);


--
-- TOC entry 2855 (class 2604 OID 17571)
-- Name: django_migrations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.django_migrations ALTER COLUMN id SET DEFAULT nextval('public.django_migrations_id_seq'::regclass);


--
-- TOC entry 3221 (class 0 OID 17597)
-- Dependencies: 209
-- Data for Name: auth_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auth_group (id, name) FROM stdin;
\.


--
-- TOC entry 3223 (class 0 OID 17607)
-- Dependencies: 211
-- Data for Name: auth_group_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auth_group_permissions (id, group_id, permission_id) FROM stdin;
\.


--
-- TOC entry 3219 (class 0 OID 17589)
-- Dependencies: 207
-- Data for Name: auth_permission; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auth_permission (id, name, content_type_id, codename) FROM stdin;
1	Can add log entry	1	add_logentry
2	Can change log entry	1	change_logentry
3	Can delete log entry	1	delete_logentry
4	Can view log entry	1	view_logentry
5	Can add permission	2	add_permission
6	Can change permission	2	change_permission
7	Can delete permission	2	delete_permission
8	Can view permission	2	view_permission
9	Can add group	3	add_group
10	Can change group	3	change_group
11	Can delete group	3	delete_group
12	Can view group	3	view_group
13	Can add user	4	add_user
14	Can change user	4	change_user
15	Can delete user	4	delete_user
16	Can view user	4	view_user
17	Can add content type	5	add_contenttype
18	Can change content type	5	change_contenttype
19	Can delete content type	5	delete_contenttype
20	Can view content type	5	view_contenttype
21	Can add session	6	add_session
22	Can change session	6	change_session
23	Can delete session	6	delete_session
24	Can view session	6	view_session
25	Can add ciudad	7	add_ciudad
26	Can change ciudad	7	change_ciudad
27	Can delete ciudad	7	delete_ciudad
28	Can view ciudad	7	view_ciudad
29	Can add medio de pago	8	add_mediodepago
30	Can change medio de pago	8	change_mediodepago
31	Can delete medio de pago	8	delete_mediodepago
32	Can view medio de pago	8	view_mediodepago
33	Can add forma de pago	9	add_formadepago
34	Can change forma de pago	9	change_formadepago
35	Can delete forma de pago	9	delete_formadepago
36	Can view forma de pago	9	view_formadepago
37	Can add tipo organizacion	10	add_tipoorganizacion
38	Can change tipo organizacion	10	change_tipoorganizacion
39	Can delete tipo organizacion	10	delete_tipoorganizacion
40	Can view tipo organizacion	10	view_tipoorganizacion
41	Can add regimen fiscal	11	add_regimenfiscal
42	Can change regimen fiscal	11	change_regimenfiscal
43	Can delete regimen fiscal	11	delete_regimenfiscal
44	Can view regimen fiscal	11	view_regimenfiscal
45	Can add responsabilidad fiscal	12	add_responsabilidadfiscal
46	Can change responsabilidad fiscal	12	change_responsabilidadfiscal
47	Can delete responsabilidad fiscal	12	delete_responsabilidadfiscal
48	Can view responsabilidad fiscal	12	view_responsabilidadfiscal
49	Can add documento identificacion	13	add_documentoidentificacion
50	Can change documento identificacion	13	change_documentoidentificacion
51	Can delete documento identificacion	13	delete_documentoidentificacion
52	Can view documento identificacion	13	view_documentoidentificacion
53	Can add documento	14	add_documento
54	Can change documento	14	change_documento
55	Can delete documento	14	delete_documento
56	Can view documento	14	view_documento
57	Can add cuenta	15	add_cuenta
58	Can change cuenta	15	change_cuenta
59	Can delete cuenta	15	delete_cuenta
60	Can view cuenta	15	view_cuenta
61	Can add tercero	16	add_tercero
62	Can change tercero	16	change_tercero
63	Can delete tercero	16	delete_tercero
64	Can view tercero	16	view_tercero
65	Can add vendedor	17	add_vendedor
66	Can change vendedor	17	change_vendedor
67	Can delete vendedor	17	delete_vendedor
68	Can view vendedor	17	view_vendedor
69	Can add centro	18	add_centro
70	Can change centro	18	change_centro
71	Can delete centro	18	delete_centro
72	Can view centro	18	view_centro
73	Can add cuota	19	add_cuota
74	Can change cuota	19	change_cuota
75	Can delete cuota	19	delete_cuota
76	Can view cuota	19	view_cuota
77	Can add cuota detalle	20	add_cuotadetalle
78	Can change cuota detalle	20	change_cuotadetalle
79	Can delete cuota detalle	20	delete_cuotadetalle
80	Can view cuota detalle	20	view_cuotadetalle
81	Can add empresa	21	add_empresa
82	Can change empresa	21	change_empresa
83	Can delete empresa	21	delete_empresa
84	Can view empresa	21	view_empresa
85	Can add usuario	22	add_usuario
86	Can change usuario	22	change_usuario
87	Can delete usuario	22	delete_usuario
88	Can view usuario	22	view_usuario
89	Can add grupo	23	add_grupo
90	Can change grupo	23	change_grupo
91	Can delete grupo	23	delete_grupo
92	Can view grupo	23	view_grupo
93	Can add sub grupo	24	add_subgrupo
94	Can change sub grupo	24	change_subgrupo
95	Can delete sub grupo	24	delete_subgrupo
96	Can view sub grupo	24	view_subgrupo
97	Can add producto	25	add_producto
98	Can change producto	25	change_producto
99	Can delete producto	25	delete_producto
100	Can view producto	25	view_producto
101	Can add bodega	26	add_bodega
102	Can change bodega	26	change_bodega
103	Can delete bodega	26	delete_bodega
104	Can view bodega	26	view_bodega
105	Can add resolucion	27	add_resolucion
106	Can change resolucion	27	change_resolucion
107	Can delete resolucion	27	delete_resolucion
108	Can view resolucion	27	view_resolucion
109	Can add compra	28	add_compra
110	Can change compra	28	change_compra
111	Can delete compra	28	delete_compra
112	Can view compra	28	view_compra
113	Can add compra detalle	29	add_compradetalle
114	Can change compra detalle	29	change_compradetalle
115	Can delete compra detalle	29	delete_compradetalle
116	Can view compra detalle	29	view_compradetalle
117	Can add comprobante cabecera	30	add_comprobantecabecera
118	Can change comprobante cabecera	30	change_comprobantecabecera
119	Can delete comprobante cabecera	30	delete_comprobantecabecera
120	Can view comprobante cabecera	30	view_comprobantecabecera
121	Can add comprobante movimiento	31	add_comprobantemovimiento
122	Can change comprobante movimiento	31	change_comprobantemovimiento
123	Can delete comprobante movimiento	31	delete_comprobantemovimiento
124	Can view comprobante movimiento	31	view_comprobantemovimiento
\.


--
-- TOC entry 3225 (class 0 OID 17615)
-- Dependencies: 213
-- Data for Name: auth_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auth_user (id, password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined) FROM stdin;
\.


--
-- TOC entry 3227 (class 0 OID 17625)
-- Dependencies: 215
-- Data for Name: auth_user_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auth_user_groups (id, user_id, group_id) FROM stdin;
\.


--
-- TOC entry 3229 (class 0 OID 17633)
-- Dependencies: 217
-- Data for Name: auth_user_user_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auth_user_user_permissions (id, user_id, permission_id) FROM stdin;
\.


--
-- TOC entry 3233 (class 0 OID 17734)
-- Dependencies: 221
-- Data for Name: contabilidad_bodega; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_bodega ("bodegaCodigo", "bodegaNombre") FROM stdin;
\.


--
-- TOC entry 3234 (class 0 OID 17739)
-- Dependencies: 222
-- Data for Name: contabilidad_centro; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.centro ("centroCodigo", "centroNombre") FROM stdin;
\.


--
-- TOC entry 3235 (class 0 OID 17744)
-- Dependencies: 223
-- Data for Name: contabilidad_ciudad; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_ciudad ("ciudadCodigo", "ciudadCodigoCiudad", "ciudadCodigoDepartamento", "ciudadNombre", "ciudadNombreCiudad", "ciudadNombreDepartamento") FROM stdin;
\.


--
-- TOC entry 3262 (class 0 OID 17880)
-- Dependencies: 250
-- Data for Name: contabilidad_compra; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_compra ("compraCodigo", "compraFecha", "compraVence", "compraHora", "compraDescuento", "compraSubtotal", "compraSubtotalIva19", "compraSubtotalIva5", "compraSubtotalIva0", "compraSubtotalBase19", "compraSubtotalBase5", "compraSubtotalBase0", "compraTotalIva", "compraTotal", "compraPorcentajeRetencionFuente", "compraTotalRetencionFuente", "compraPorcentajeRetencionIca", "compraTotalRetencionIca", "compraNeto", "compraItems", "compraFormaDePago_id", "compraMedioDePago_id", "compraResolucion_id", "compraTerceroCodigo_id", "compraVendedor_id") FROM stdin;
\.


--
-- TOC entry 3261 (class 0 OID 17874)
-- Dependencies: 249
-- Data for Name: contabilidad_compradetalle; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_compradetalle ("compraId", "compraCodigo", "compraFila", "compraCantidad", "compraPrecio", "compraDescuento", "compraMontoDescuento", "compraSiGratis", "compraSubtotal", "compraSubtotalDescuento", "compraPagina", "compraBodega_id", "compraProductoCodigo_id") FROM stdin;
\.


--
-- TOC entry 3259 (class 0 OID 17866)
-- Dependencies: 247
-- Data for Name: contabilidad_comprobantecabecera; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_comprobantecabecera ("comprobanteCodigo", "comprobanteConsecutivo", "comprobantePrefijo", "comprobanteResolucion", "comprobanteFecha", "comprobanteAño", "comprobanteBloqueo", "documentoCodigo_id") FROM stdin;
\.


--
-- TOC entry 3257 (class 0 OID 17858)
-- Dependencies: 245
-- Data for Name: contabilidad_comprobantemovimiento; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_comprobantemovimiento ("comprobanteMovimientoId", "comprobanteConsecutivo", "comprobantePrefijo", "movimientoFila", "movimientoConcepto", "movimientoDocumento", "movimientoFecha", "movimientoDebito", "movimientoCredito", "centroCodigo_id", "cuentaCodigo_id", "documentoCodigo_id", "terceroCodigo_id") FROM stdin;
\.


--
-- TOC entry 3236 (class 0 OID 17749)
-- Dependencies: 224
-- Data for Name: contabilidad_cuenta; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_cuenta ("cuentaCodigo", "cuentaNombre", "cuentaAuto", "cuentaNivel", "cuentaTercero", "cuentaCentro", "cuentaFactura", "cuentaInforme", "cuentaContra", "cuentaInteres", "cuentaCuota", "cuentaIntereses", "cuentaTipo", "cuentaGrupo", "cuentaContraNombre", "cuentaInteresNombre") FROM stdin;
\.


--
-- TOC entry 3238 (class 0 OID 17756)
-- Dependencies: 226
-- Data for Name: contabilidad_cuota; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_cuota ("cuotaId", "cuotaPrestamo", "cuotaPlazo", "cuotaInteres", "cuotaFecha", "cuotaPago", "cuotaTotalInteres", "cuotaTotalPago") FROM stdin;
\.


--
-- TOC entry 3255 (class 0 OID 17850)
-- Dependencies: 243
-- Data for Name: contabilidad_cuotadetalle; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_cuotadetalle (id, "cuotaFila", "cuotaDetalleFecha", "cuotaDetalleCuota", "cuotaDetalleInteres", "cuotaDetalleCapital", "cuotaDetalleSaldo", "cuotaCodigo_id") FROM stdin;
\.


--
-- TOC entry 3239 (class 0 OID 17762)
-- Dependencies: 227
-- Data for Name: contabilidad_documento; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_documento ("documentoCodigo", "documentoNombre", "documentoConsecutivo", "documentoInicial") FROM stdin;
\.


--
-- TOC entry 3240 (class 0 OID 17767)
-- Dependencies: 228
-- Data for Name: contabilidad_documentoidentificacion; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_documentoidentificacion ("documentoIdentificacionCodigo", "documentoIdentificacionNombre") FROM stdin;
\.


--
-- TOC entry 3253 (class 0 OID 17840)
-- Dependencies: 241
-- Data for Name: contabilidad_empresa; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_empresa ("empresaCodigo", "empresaDv", "empresaNombre", "empresaIva", "empresaReteIva", "empresaDireccion", "empresaActividadIca", "empresaTelefono1", "empresaTelefono2", "empresaEmail1", "empresaEmail2", "empresaActiva", "empresaLicencia", "empresaRepresentanteDv", "empresaRepresentanteNombre", "empresaContadorDv", "empresaContadorNombre", "empresaRevisorDv", "empresaRevisorNombre", "empresaCiudad_id", "empresaContadorNit_id", "empresaDocumento_id", "empresaFiscal_id", "empresaRegimen_id", "empresaRepresentanteNit_id", "empresaRevisorNit_id", "empresaTipo_id") FROM stdin;
\.


--
-- TOC entry 3241 (class 0 OID 17772)
-- Dependencies: 229
-- Data for Name: contabilidad_formadepago; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_formadepago ("formaDePagoCodigo", "formaDePagoNombre") FROM stdin;
\.


--
-- TOC entry 3242 (class 0 OID 17777)
-- Dependencies: 230
-- Data for Name: contabilidad_grupo; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_grupo ("grupoCodigo", "grupoNombre") FROM stdin;
\.


--
-- TOC entry 3243 (class 0 OID 17782)
-- Dependencies: 231
-- Data for Name: contabilidad_mediodepago; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_mediodepago ("medioDePagoCodigo", "medioDePagoNombre") FROM stdin;
\.


--
-- TOC entry 3252 (class 0 OID 17835)
-- Dependencies: 240
-- Data for Name: contabilidad_producto; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_producto ("productoCodigo", "productoNombre", "productoUnidad", "productoIva", "productoTipo", "productoVenta", "productoCosto", "productoCantidad", "productoTotal", "productoSubgrupo_id") FROM stdin;
\.


--
-- TOC entry 3244 (class 0 OID 17787)
-- Dependencies: 232
-- Data for Name: contabilidad_regimenfiscal; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_regimenfiscal ("regimenFiscalCodigo", "regimenFiscalNombre") FROM stdin;
\.


--
-- TOC entry 3251 (class 0 OID 17830)
-- Dependencies: 239
-- Data for Name: contabilidad_resolucion; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_resolucion ("resolucionCodigo", "resolucionDian", "resolucionPrefijo", "resolucionTipo", "resolucionNombreLocal", "resolucionDireccion", "resolucionTelefono", "resolucionInforme", "resolucionFechaInicial", "resolucionFechaFinal", "resolucionNumeroInicial", "resolucionNumeroFinal", "resolucionNumeroActual", "resolucionCiudad_id") FROM stdin;
\.


--
-- TOC entry 3245 (class 0 OID 17792)
-- Dependencies: 233
-- Data for Name: contabilidad_responsabilidadfiscal; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_responsabilidadfiscal ("responsabilidadFiscalCodigo", "responsabilidadFiscalNombre") FROM stdin;
\.


--
-- TOC entry 3250 (class 0 OID 17825)
-- Dependencies: 238
-- Data for Name: contabilidad_subgrupo; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_subgrupo ("subgrupoCodigo", "subgrupoNombre", "subgrupoCodigoGrupo_id") FROM stdin;
\.


--
-- TOC entry 3246 (class 0 OID 17797)
-- Dependencies: 234
-- Data for Name: contabilidad_tercero; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_tercero ("terceroCodigo", "terceroDv", "terceroNombre", "terceroJuridica", "terceroPrimerNombre", "terceroSegundoNombre", "terceroPrimerApellido", "terceroSegundoApellido", "terceroDireccion", "terceroBarrio", "terceroTelefono1", "terceroTelefono2", "terceroEmail1", "terceroEmail2", "terceroContacto", "terceroRut", "terceroBloque", "terceroPiso", "terceroApartamento", "terceroDescuento1", "terceroDescuento2", "terceroCuotaP", "terceroCuota1", "terceroCuota2", "terceroCuota3", "terceroArea", "terceroFactor", "terceroMatricula", "terceroCatastral", "terceroBanco", "tercerophCodigo", "tercerophDv", "tercerophNombre", "terceroCiudad_id", "terceroDocumento_id", "terceroFiscal_id", "terceroRegimen_id", "terceroTipo_id") FROM stdin;
\.


--
-- TOC entry 3247 (class 0 OID 17805)
-- Dependencies: 235
-- Data for Name: contabilidad_tipoorganizacion; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_tipoorganizacion ("tipoOrganizacionCodigo", "tipoOrganizacionNombre") FROM stdin;
\.


--
-- TOC entry 3249 (class 0 OID 17815)
-- Dependencies: 237
-- Data for Name: contabilidad_usuario; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_usuario ("usuarioCodigo", "usuarioDv", "usuarioNombre", "usuarioTipo", "usuarioClave2", "usuarioClave1", "usuarioCorreo1", "usuarioCorreo2", "usuarioNit_id") FROM stdin;
\.


--
-- TOC entry 3248 (class 0 OID 17810)
-- Dependencies: 236
-- Data for Name: contabilidad_vendedor; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contabilidad_vendedor ("vendedorCodigo", "vendedorDv", "vendedorNombre", "vendedorComision", "vendedorNit_id") FROM stdin;
\.


--
-- TOC entry 3231 (class 0 OID 17693)
-- Dependencies: 219
-- Data for Name: django_admin_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.django_admin_log (id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id) FROM stdin;
\.


--
-- TOC entry 3217 (class 0 OID 17579)
-- Dependencies: 205
-- Data for Name: django_content_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.django_content_type (id, app_label, model) FROM stdin;
1	admin	logentry
2	auth	permission
3	auth	group
4	auth	user
5	contenttypes	contenttype
6	sessions	session
7	contabilidad	ciudad
8	contabilidad	mediodepago
9	contabilidad	formadepago
10	contabilidad	tipoorganizacion
11	contabilidad	regimenfiscal
12	contabilidad	responsabilidadfiscal
13	contabilidad	documentoidentificacion
14	contabilidad	documento
15	contabilidad	cuenta
16	contabilidad	tercero
17	contabilidad	vendedor
18	contabilidad	centro
19	contabilidad	cuota
20	contabilidad	cuotadetalle
21	contabilidad	empresa
22	contabilidad	usuario
23	contabilidad	grupo
24	contabilidad	subgrupo
25	contabilidad	producto
26	contabilidad	bodega
27	contabilidad	resolucion
28	contabilidad	compra
29	contabilidad	compradetalle
30	contabilidad	comprobantecabecera
31	contabilidad	comprobantemovimiento
\.


--
-- TOC entry 3215 (class 0 OID 17568)
-- Dependencies: 203
-- Data for Name: django_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.django_migrations (id, app, name, applied) FROM stdin;
1	contenttypes	0001_initial	2021-02-03 21:49:25.034219-05
2	auth	0001_initial	2021-02-03 21:49:25.457779-05
3	admin	0001_initial	2021-02-03 21:49:25.990492-05
4	admin	0002_logentry_remove_auto_add	2021-02-03 21:49:26.067212-05
5	admin	0003_logentry_add_action_flag_choices	2021-02-03 21:49:26.075184-05
6	contenttypes	0002_remove_content_type_name	2021-02-03 21:49:26.12916-05
7	auth	0002_alter_permission_name_max_length	2021-02-03 21:49:26.138138-05
8	auth	0003_alter_user_email_max_length	2021-02-03 21:49:26.146118-05
9	auth	0004_alter_user_username_opts	2021-02-03 21:49:26.153225-05
10	auth	0005_alter_user_last_login_null	2021-02-03 21:49:26.162192-05
11	auth	0006_require_contenttypes_0002	2021-02-03 21:49:26.164187-05
12	auth	0007_alter_validators_add_error_messages	2021-02-03 21:49:26.172166-05
13	auth	0008_alter_user_username_max_length	2021-02-03 21:49:26.220354-05
14	auth	0009_alter_user_last_name_max_length	2021-02-03 21:49:26.229334-05
15	auth	0010_alter_group_name_max_length	2021-02-03 21:49:26.238279-05
16	auth	0011_update_proxy_permissions	2021-02-03 21:49:26.260226-05
17	auth	0012_alter_user_first_name_max_length	2021-02-03 21:49:26.269192-05
18	sessions	0001_initial	2021-02-03 21:49:26.319963-05
19	contabilidad	0001_initial	2021-02-03 21:49:55.585944-05
\.


--
-- TOC entry 3232 (class 0 OID 17724)
-- Dependencies: 220
-- Data for Name: django_session; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.django_session (session_key, session_data, expire_date) FROM stdin;
\.


--
-- TOC entry 3282 (class 0 OID 0)
-- Dependencies: 208
-- Name: auth_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.auth_group_id_seq', 1, false);


--
-- TOC entry 3283 (class 0 OID 0)
-- Dependencies: 210
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.auth_group_permissions_id_seq', 1, false);


--
-- TOC entry 3284 (class 0 OID 0)
-- Dependencies: 206
-- Name: auth_permission_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.auth_permission_id_seq', 124, true);


--
-- TOC entry 3285 (class 0 OID 0)
-- Dependencies: 214
-- Name: auth_user_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.auth_user_groups_id_seq', 1, false);


--
-- TOC entry 3286 (class 0 OID 0)
-- Dependencies: 212
-- Name: auth_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.auth_user_id_seq', 1, false);


--
-- TOC entry 3287 (class 0 OID 0)
-- Dependencies: 216
-- Name: auth_user_user_permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.auth_user_user_permissions_id_seq', 1, false);


--
-- TOC entry 3288 (class 0 OID 0)
-- Dependencies: 248
-- Name: contabilidad_compradetalle_compraId_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."contabilidad_compradetalle_compraId_seq"', 1, false);


--
-- TOC entry 3289 (class 0 OID 0)
-- Dependencies: 246
-- Name: contabilidad_comprobantecabecera_comprobanteCodigo_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."contabilidad_comprobantecabecera_comprobanteCodigo_seq"', 1, false);


--
-- TOC entry 3290 (class 0 OID 0)
-- Dependencies: 244
-- Name: contabilidad_comprobantemovimiento_comprobanteMovimientoId_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."contabilidad_comprobantemovimiento_comprobanteMovimientoId_seq"', 1, false);


--
-- TOC entry 3291 (class 0 OID 0)
-- Dependencies: 225
-- Name: contabilidad_cuota_cuotaId_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."contabilidad_cuota_cuotaId_seq"', 1, false);


--
-- TOC entry 3292 (class 0 OID 0)
-- Dependencies: 242
-- Name: contabilidad_cuotadetalle_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.contabilidad_cuotadetalle_id_seq', 1, false);


--
-- TOC entry 3293 (class 0 OID 0)
-- Dependencies: 218
-- Name: django_admin_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.django_admin_log_id_seq', 1, false);


--
-- TOC entry 3294 (class 0 OID 0)
-- Dependencies: 204
-- Name: django_content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.django_content_type_id_seq', 31, true);


--
-- TOC entry 3295 (class 0 OID 0)
-- Dependencies: 202
-- Name: django_migrations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.django_migrations_id_seq', 19, true);


--
-- TOC entry 2883 (class 2606 OID 17722)
-- Name: auth_group auth_group_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_name_key UNIQUE (name);


--
-- TOC entry 2888 (class 2606 OID 17649)
-- Name: auth_group_permissions auth_group_permissions_group_id_permission_id_0cd325b0_uniq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_group_id_permission_id_0cd325b0_uniq UNIQUE (group_id, permission_id);


--
-- TOC entry 2891 (class 2606 OID 17612)
-- Name: auth_group_permissions auth_group_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_pkey PRIMARY KEY (id);


--
-- TOC entry 2885 (class 2606 OID 17602)
-- Name: auth_group auth_group_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_pkey PRIMARY KEY (id);


--
-- TOC entry 2878 (class 2606 OID 17640)
-- Name: auth_permission auth_permission_content_type_id_codename_01ab375a_uniq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_content_type_id_codename_01ab375a_uniq UNIQUE (content_type_id, codename);


--
-- TOC entry 2880 (class 2606 OID 17594)
-- Name: auth_permission auth_permission_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_pkey PRIMARY KEY (id);


--
-- TOC entry 2899 (class 2606 OID 17630)
-- Name: auth_user_groups auth_user_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_pkey PRIMARY KEY (id);


--
-- TOC entry 2902 (class 2606 OID 17664)
-- Name: auth_user_groups auth_user_groups_user_id_group_id_94350c0c_uniq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_user_id_group_id_94350c0c_uniq UNIQUE (user_id, group_id);


--
-- TOC entry 2893 (class 2606 OID 17620)
-- Name: auth_user auth_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user
    ADD CONSTRAINT auth_user_pkey PRIMARY KEY (id);


--
-- TOC entry 2905 (class 2606 OID 17638)
-- Name: auth_user_user_permissions auth_user_user_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_pkey PRIMARY KEY (id);


--
-- TOC entry 2908 (class 2606 OID 17678)
-- Name: auth_user_user_permissions auth_user_user_permissions_user_id_permission_id_14a6b632_uniq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_user_id_permission_id_14a6b632_uniq UNIQUE (user_id, permission_id);


--
-- TOC entry 2896 (class 2606 OID 17716)
-- Name: auth_user auth_user_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user
    ADD CONSTRAINT auth_user_username_key UNIQUE (username);


--
-- TOC entry 2919 (class 2606 OID 17738)
-- Name: contabilidad_bodega contabilidad_bodega_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_bodega
    ADD CONSTRAINT contabilidad_bodega_pkey PRIMARY KEY ("bodegaCodigo");


--
-- TOC entry 2922 (class 2606 OID 17743)
-- Name: contabilidad_centro contabilidad_centro_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_centro
    ADD CONSTRAINT contabilidad_centro_pkey PRIMARY KEY ("centroCodigo");


--
-- TOC entry 2925 (class 2606 OID 17748)
-- Name: contabilidad_ciudad contabilidad_ciudad_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_ciudad
    ADD CONSTRAINT contabilidad_ciudad_pkey PRIMARY KEY ("ciudadCodigo");


--
-- TOC entry 3047 (class 2606 OID 17884)
-- Name: contabilidad_compra contabilidad_compra_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_compra
    ADD CONSTRAINT contabilidad_compra_pkey PRIMARY KEY ("compraCodigo");


--
-- TOC entry 3034 (class 2606 OID 17879)
-- Name: contabilidad_compradetalle contabilidad_compradetalle_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_compradetalle
    ADD CONSTRAINT contabilidad_compradetalle_pkey PRIMARY KEY ("compraId");


--
-- TOC entry 3028 (class 2606 OID 17871)
-- Name: contabilidad_comprobantecabecera contabilidad_comprobantecabecera_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_comprobantecabecera
    ADD CONSTRAINT contabilidad_comprobantecabecera_pkey PRIMARY KEY ("comprobanteCodigo");


--
-- TOC entry 3023 (class 2606 OID 17863)
-- Name: contabilidad_comprobantemovimiento contabilidad_comprobantemovimiento_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_comprobantemovimiento
    ADD CONSTRAINT contabilidad_comprobantemovimiento_pkey PRIMARY KEY ("comprobanteMovimientoId");


--
-- TOC entry 2928 (class 2606 OID 17753)
-- Name: contabilidad_cuenta contabilidad_cuenta_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_cuenta
    ADD CONSTRAINT contabilidad_cuenta_pkey PRIMARY KEY ("cuentaCodigo");


--
-- TOC entry 2930 (class 2606 OID 17761)
-- Name: contabilidad_cuota contabilidad_cuota_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_cuota
    ADD CONSTRAINT contabilidad_cuota_pkey PRIMARY KEY ("cuotaId");


--
-- TOC entry 3014 (class 2606 OID 17855)
-- Name: contabilidad_cuotadetalle contabilidad_cuotadetalle_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_cuotadetalle
    ADD CONSTRAINT contabilidad_cuotadetalle_pkey PRIMARY KEY (id);


--
-- TOC entry 2933 (class 2606 OID 17766)
-- Name: contabilidad_documento contabilidad_documento_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_documento
    ADD CONSTRAINT contabilidad_documento_pkey PRIMARY KEY ("documentoCodigo");


--
-- TOC entry 2936 (class 2606 OID 17771)
-- Name: contabilidad_documentoidentificacion contabilidad_documentoidentificacion_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_documentoidentificacion
    ADD CONSTRAINT contabilidad_documentoidentificacion_pkey PRIMARY KEY ("documentoIdentificacionCodigo");


--
-- TOC entry 3011 (class 2606 OID 17847)
-- Name: contabilidad_empresa contabilidad_empresa_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_empresa
    ADD CONSTRAINT contabilidad_empresa_pkey PRIMARY KEY ("empresaCodigo");


--
-- TOC entry 2939 (class 2606 OID 17776)
-- Name: contabilidad_formadepago contabilidad_formadepago_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_formadepago
    ADD CONSTRAINT contabilidad_formadepago_pkey PRIMARY KEY ("formaDePagoCodigo");


--
-- TOC entry 2942 (class 2606 OID 17781)
-- Name: contabilidad_grupo contabilidad_grupo_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_grupo
    ADD CONSTRAINT contabilidad_grupo_pkey PRIMARY KEY ("grupoCodigo");


--
-- TOC entry 2945 (class 2606 OID 17786)
-- Name: contabilidad_mediodepago contabilidad_mediodepago_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_mediodepago
    ADD CONSTRAINT contabilidad_mediodepago_pkey PRIMARY KEY ("medioDePagoCodigo");


--
-- TOC entry 2989 (class 2606 OID 17839)
-- Name: contabilidad_producto contabilidad_producto_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_producto
    ADD CONSTRAINT contabilidad_producto_pkey PRIMARY KEY ("productoCodigo");


--
-- TOC entry 2947 (class 2606 OID 17791)
-- Name: contabilidad_regimenfiscal contabilidad_regimenfiscal_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_regimenfiscal
    ADD CONSTRAINT contabilidad_regimenfiscal_pkey PRIMARY KEY ("regimenFiscalCodigo");


--
-- TOC entry 2984 (class 2606 OID 17834)
-- Name: contabilidad_resolucion contabilidad_resolucion_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_resolucion
    ADD CONSTRAINT contabilidad_resolucion_pkey PRIMARY KEY ("resolucionCodigo");


--
-- TOC entry 2951 (class 2606 OID 17796)
-- Name: contabilidad_responsabilidadfiscal contabilidad_responsabilidadfiscal_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_responsabilidadfiscal
    ADD CONSTRAINT contabilidad_responsabilidadfiscal_pkey PRIMARY KEY ("responsabilidadFiscalCodigo");


--
-- TOC entry 2979 (class 2606 OID 17829)
-- Name: contabilidad_subgrupo contabilidad_subgrupo_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_subgrupo
    ADD CONSTRAINT contabilidad_subgrupo_pkey PRIMARY KEY ("subgrupoCodigo");


--
-- TOC entry 2953 (class 2606 OID 17804)
-- Name: contabilidad_tercero contabilidad_tercero_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_tercero
    ADD CONSTRAINT contabilidad_tercero_pkey PRIMARY KEY ("terceroCodigo");


--
-- TOC entry 2967 (class 2606 OID 17809)
-- Name: contabilidad_tipoorganizacion contabilidad_tipoorganizacion_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_tipoorganizacion
    ADD CONSTRAINT contabilidad_tipoorganizacion_pkey PRIMARY KEY ("tipoOrganizacionCodigo");


--
-- TOC entry 2974 (class 2606 OID 17819)
-- Name: contabilidad_usuario contabilidad_usuario_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_usuario
    ADD CONSTRAINT contabilidad_usuario_pkey PRIMARY KEY ("usuarioCodigo");


--
-- TOC entry 2969 (class 2606 OID 17814)
-- Name: contabilidad_vendedor contabilidad_vendedor_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_vendedor
    ADD CONSTRAINT contabilidad_vendedor_pkey PRIMARY KEY ("vendedorCodigo");


--
-- TOC entry 2911 (class 2606 OID 17702)
-- Name: django_admin_log django_admin_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_pkey PRIMARY KEY (id);


--
-- TOC entry 2873 (class 2606 OID 17586)
-- Name: django_content_type django_content_type_app_label_model_76bd3d3b_uniq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_app_label_model_76bd3d3b_uniq UNIQUE (app_label, model);


--
-- TOC entry 2875 (class 2606 OID 17584)
-- Name: django_content_type django_content_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_pkey PRIMARY KEY (id);


--
-- TOC entry 2871 (class 2606 OID 17576)
-- Name: django_migrations django_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.django_migrations
    ADD CONSTRAINT django_migrations_pkey PRIMARY KEY (id);


--
-- TOC entry 2915 (class 2606 OID 17731)
-- Name: django_session django_session_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.django_session
    ADD CONSTRAINT django_session_pkey PRIMARY KEY (session_key);


--
-- TOC entry 2881 (class 1259 OID 17723)
-- Name: auth_group_name_a6ea08ec_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX auth_group_name_a6ea08ec_like ON public.auth_group USING btree (name varchar_pattern_ops);


--
-- TOC entry 2886 (class 1259 OID 17660)
-- Name: auth_group_permissions_group_id_b120cbf9; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX auth_group_permissions_group_id_b120cbf9 ON public.auth_group_permissions USING btree (group_id);


--
-- TOC entry 2889 (class 1259 OID 17661)
-- Name: auth_group_permissions_permission_id_84c5c92e; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX auth_group_permissions_permission_id_84c5c92e ON public.auth_group_permissions USING btree (permission_id);


--
-- TOC entry 2876 (class 1259 OID 17646)
-- Name: auth_permission_content_type_id_2f476e4b; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX auth_permission_content_type_id_2f476e4b ON public.auth_permission USING btree (content_type_id);


--
-- TOC entry 2897 (class 1259 OID 17676)
-- Name: auth_user_groups_group_id_97559544; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX auth_user_groups_group_id_97559544 ON public.auth_user_groups USING btree (group_id);


--
-- TOC entry 2900 (class 1259 OID 17675)
-- Name: auth_user_groups_user_id_6a12ed8b; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX auth_user_groups_user_id_6a12ed8b ON public.auth_user_groups USING btree (user_id);


--
-- TOC entry 2903 (class 1259 OID 17690)
-- Name: auth_user_user_permissions_permission_id_1fbb5f2c; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX auth_user_user_permissions_permission_id_1fbb5f2c ON public.auth_user_user_permissions USING btree (permission_id);


--
-- TOC entry 2906 (class 1259 OID 17689)
-- Name: auth_user_user_permissions_user_id_a95ead1b; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX auth_user_user_permissions_user_id_a95ead1b ON public.auth_user_user_permissions USING btree (user_id);


--
-- TOC entry 2894 (class 1259 OID 17717)
-- Name: auth_user_username_6821ab7c_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX auth_user_username_6821ab7c_like ON public.auth_user USING btree (username varchar_pattern_ops);


--
-- TOC entry 2917 (class 1259 OID 17885)
-- Name: contabilidad_bodega_bodegaCodigo_9213c635_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_bodega_bodegaCodigo_9213c635_like" ON public.contabilidad_bodega USING btree ("bodegaCodigo" varchar_pattern_ops);


--
-- TOC entry 2920 (class 1259 OID 17886)
-- Name: contabilidad_centro_centroCodigo_6d6c99f1_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_centro_centroCodigo_6d6c99f1_like" ON public.contabilidad_centro USING btree ("centroCodigo" varchar_pattern_ops);


--
-- TOC entry 2923 (class 1259 OID 17887)
-- Name: contabilidad_ciudad_ciudadCodigo_8f43ef39_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_ciudad_ciudadCodigo_8f43ef39_like" ON public.contabilidad_ciudad USING btree ("ciudadCodigo" varchar_pattern_ops);


--
-- TOC entry 3035 (class 1259 OID 18105)
-- Name: contabilidad_compra_compraCodigo_82044488_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraCodigo_82044488_like" ON public.contabilidad_compra USING btree ("compraCodigo" varchar_pattern_ops);


--
-- TOC entry 3036 (class 1259 OID 18106)
-- Name: contabilidad_compra_compraFormaDePago_id_b907ff46; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraFormaDePago_id_b907ff46" ON public.contabilidad_compra USING btree ("compraFormaDePago_id");


--
-- TOC entry 3037 (class 1259 OID 18107)
-- Name: contabilidad_compra_compraFormaDePago_id_b907ff46_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraFormaDePago_id_b907ff46_like" ON public.contabilidad_compra USING btree ("compraFormaDePago_id" varchar_pattern_ops);


--
-- TOC entry 3038 (class 1259 OID 18108)
-- Name: contabilidad_compra_compraMedioDePago_id_c7bfc061; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraMedioDePago_id_c7bfc061" ON public.contabilidad_compra USING btree ("compraMedioDePago_id");


--
-- TOC entry 3039 (class 1259 OID 18109)
-- Name: contabilidad_compra_compraMedioDePago_id_c7bfc061_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraMedioDePago_id_c7bfc061_like" ON public.contabilidad_compra USING btree ("compraMedioDePago_id" varchar_pattern_ops);


--
-- TOC entry 3040 (class 1259 OID 18110)
-- Name: contabilidad_compra_compraResolucion_id_9eb8193f; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraResolucion_id_9eb8193f" ON public.contabilidad_compra USING btree ("compraResolucion_id");


--
-- TOC entry 3041 (class 1259 OID 18111)
-- Name: contabilidad_compra_compraResolucion_id_9eb8193f_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraResolucion_id_9eb8193f_like" ON public.contabilidad_compra USING btree ("compraResolucion_id" varchar_pattern_ops);


--
-- TOC entry 3042 (class 1259 OID 18112)
-- Name: contabilidad_compra_compraTerceroCodigo_id_01a13414; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraTerceroCodigo_id_01a13414" ON public.contabilidad_compra USING btree ("compraTerceroCodigo_id");


--
-- TOC entry 3043 (class 1259 OID 18113)
-- Name: contabilidad_compra_compraTerceroCodigo_id_01a13414_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraTerceroCodigo_id_01a13414_like" ON public.contabilidad_compra USING btree ("compraTerceroCodigo_id" varchar_pattern_ops);


--
-- TOC entry 3044 (class 1259 OID 18114)
-- Name: contabilidad_compra_compraVendedor_id_9e86eb27; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraVendedor_id_9e86eb27" ON public.contabilidad_compra USING btree ("compraVendedor_id");


--
-- TOC entry 3045 (class 1259 OID 18115)
-- Name: contabilidad_compra_compraVendedor_id_9e86eb27_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compra_compraVendedor_id_9e86eb27_like" ON public.contabilidad_compra USING btree ("compraVendedor_id" varchar_pattern_ops);


--
-- TOC entry 3029 (class 1259 OID 18079)
-- Name: contabilidad_compradetal_compraProductoCodigo_id_884ab25b_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compradetal_compraProductoCodigo_id_884ab25b_like" ON public.contabilidad_compradetalle USING btree ("compraProductoCodigo_id" varchar_pattern_ops);


--
-- TOC entry 3030 (class 1259 OID 18076)
-- Name: contabilidad_compradetalle_compraBodega_id_ee34fff5; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compradetalle_compraBodega_id_ee34fff5" ON public.contabilidad_compradetalle USING btree ("compraBodega_id");


--
-- TOC entry 3031 (class 1259 OID 18077)
-- Name: contabilidad_compradetalle_compraBodega_id_ee34fff5_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compradetalle_compraBodega_id_ee34fff5_like" ON public.contabilidad_compradetalle USING btree ("compraBodega_id" varchar_pattern_ops);


--
-- TOC entry 3032 (class 1259 OID 18078)
-- Name: contabilidad_compradetalle_compraProductoCodigo_id_884ab25b; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_compradetalle_compraProductoCodigo_id_884ab25b" ON public.contabilidad_compradetalle USING btree ("compraProductoCodigo_id");


--
-- TOC entry 3015 (class 1259 OID 18052)
-- Name: contabilidad_comprobante_centroCodigo_id_9e9de43a_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_comprobante_centroCodigo_id_9e9de43a_like" ON public.contabilidad_comprobantemovimiento USING btree ("centroCodigo_id" varchar_pattern_ops);


--
-- TOC entry 3016 (class 1259 OID 18054)
-- Name: contabilidad_comprobante_cuentaCodigo_id_f0261d89_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_comprobante_cuentaCodigo_id_f0261d89_like" ON public.contabilidad_comprobantemovimiento USING btree ("cuentaCodigo_id" varchar_pattern_ops);


--
-- TOC entry 3025 (class 1259 OID 18065)
-- Name: contabilidad_comprobante_documentoCodigo_id_0858616e_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_comprobante_documentoCodigo_id_0858616e_like" ON public.contabilidad_comprobantecabecera USING btree ("documentoCodigo_id" varchar_pattern_ops);


--
-- TOC entry 3017 (class 1259 OID 18056)
-- Name: contabilidad_comprobante_documentoCodigo_id_2d16dddc_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_comprobante_documentoCodigo_id_2d16dddc_like" ON public.contabilidad_comprobantemovimiento USING btree ("documentoCodigo_id" varchar_pattern_ops);


--
-- TOC entry 3018 (class 1259 OID 18058)
-- Name: contabilidad_comprobante_terceroCodigo_id_51cd7a56_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_comprobante_terceroCodigo_id_51cd7a56_like" ON public.contabilidad_comprobantemovimiento USING btree ("terceroCodigo_id" varchar_pattern_ops);


--
-- TOC entry 3026 (class 1259 OID 18064)
-- Name: contabilidad_comprobantecabecera_documentoCodigo_id_0858616e; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_comprobantecabecera_documentoCodigo_id_0858616e" ON public.contabilidad_comprobantecabecera USING btree ("documentoCodigo_id");


--
-- TOC entry 3019 (class 1259 OID 18051)
-- Name: contabilidad_comprobantemovimiento_centroCodigo_id_9e9de43a; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_comprobantemovimiento_centroCodigo_id_9e9de43a" ON public.contabilidad_comprobantemovimiento USING btree ("centroCodigo_id");


--
-- TOC entry 3020 (class 1259 OID 18053)
-- Name: contabilidad_comprobantemovimiento_cuentaCodigo_id_f0261d89; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_comprobantemovimiento_cuentaCodigo_id_f0261d89" ON public.contabilidad_comprobantemovimiento USING btree ("cuentaCodigo_id");


--
-- TOC entry 3021 (class 1259 OID 18055)
-- Name: contabilidad_comprobantemovimiento_documentoCodigo_id_2d16dddc; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_comprobantemovimiento_documentoCodigo_id_2d16dddc" ON public.contabilidad_comprobantemovimiento USING btree ("documentoCodigo_id");


--
-- TOC entry 3024 (class 1259 OID 18057)
-- Name: contabilidad_comprobantemovimiento_terceroCodigo_id_51cd7a56; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_comprobantemovimiento_terceroCodigo_id_51cd7a56" ON public.contabilidad_comprobantemovimiento USING btree ("terceroCodigo_id");


--
-- TOC entry 2926 (class 1259 OID 17888)
-- Name: contabilidad_cuenta_cuentaCodigo_5fa37d7e_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_cuenta_cuentaCodigo_5fa37d7e_like" ON public.contabilidad_cuenta USING btree ("cuentaCodigo" varchar_pattern_ops);


--
-- TOC entry 3012 (class 1259 OID 18030)
-- Name: contabilidad_cuotadetalle_cuotaCodigo_id_f3f551b8; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_cuotadetalle_cuotaCodigo_id_f3f551b8" ON public.contabilidad_cuotadetalle USING btree ("cuotaCodigo_id");


--
-- TOC entry 2931 (class 1259 OID 17889)
-- Name: contabilidad_documento_documentoCodigo_6ba8fa91_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_documento_documentoCodigo_6ba8fa91_like" ON public.contabilidad_documento USING btree ("documentoCodigo" varchar_pattern_ops);


--
-- TOC entry 2934 (class 1259 OID 17890)
-- Name: contabilidad_documentoid_documentoIdentificacionC_28b059ff_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_documentoid_documentoIdentificacionC_28b059ff_like" ON public.contabilidad_documentoidentificacion USING btree ("documentoIdentificacionCodigo" varchar_pattern_ops);


--
-- TOC entry 2993 (class 1259 OID 18009)
-- Name: contabilidad_empresa_empresaCiudad_id_2b55b473; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaCiudad_id_2b55b473" ON public.contabilidad_empresa USING btree ("empresaCiudad_id");


--
-- TOC entry 2994 (class 1259 OID 18010)
-- Name: contabilidad_empresa_empresaCiudad_id_2b55b473_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaCiudad_id_2b55b473_like" ON public.contabilidad_empresa USING btree ("empresaCiudad_id" varchar_pattern_ops);


--
-- TOC entry 2995 (class 1259 OID 18008)
-- Name: contabilidad_empresa_empresaCodigo_4a44f163_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaCodigo_4a44f163_like" ON public.contabilidad_empresa USING btree ("empresaCodigo" varchar_pattern_ops);


--
-- TOC entry 2996 (class 1259 OID 18011)
-- Name: contabilidad_empresa_empresaContadorNit_id_2aed5911; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaContadorNit_id_2aed5911" ON public.contabilidad_empresa USING btree ("empresaContadorNit_id");


--
-- TOC entry 2997 (class 1259 OID 18012)
-- Name: contabilidad_empresa_empresaContadorNit_id_2aed5911_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaContadorNit_id_2aed5911_like" ON public.contabilidad_empresa USING btree ("empresaContadorNit_id" varchar_pattern_ops);


--
-- TOC entry 2998 (class 1259 OID 18013)
-- Name: contabilidad_empresa_empresaDocumento_id_3886112f; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaDocumento_id_3886112f" ON public.contabilidad_empresa USING btree ("empresaDocumento_id");


--
-- TOC entry 2999 (class 1259 OID 18014)
-- Name: contabilidad_empresa_empresaDocumento_id_3886112f_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaDocumento_id_3886112f_like" ON public.contabilidad_empresa USING btree ("empresaDocumento_id" varchar_pattern_ops);


--
-- TOC entry 3000 (class 1259 OID 18015)
-- Name: contabilidad_empresa_empresaFiscal_id_af71a018; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaFiscal_id_af71a018" ON public.contabilidad_empresa USING btree ("empresaFiscal_id");


--
-- TOC entry 3001 (class 1259 OID 18016)
-- Name: contabilidad_empresa_empresaFiscal_id_af71a018_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaFiscal_id_af71a018_like" ON public.contabilidad_empresa USING btree ("empresaFiscal_id" varchar_pattern_ops);


--
-- TOC entry 3002 (class 1259 OID 18017)
-- Name: contabilidad_empresa_empresaRegimen_id_999609f6; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaRegimen_id_999609f6" ON public.contabilidad_empresa USING btree ("empresaRegimen_id");


--
-- TOC entry 3003 (class 1259 OID 18018)
-- Name: contabilidad_empresa_empresaRegimen_id_999609f6_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaRegimen_id_999609f6_like" ON public.contabilidad_empresa USING btree ("empresaRegimen_id" varchar_pattern_ops);


--
-- TOC entry 3004 (class 1259 OID 18019)
-- Name: contabilidad_empresa_empresaRepresentanteNit_id_5ae16bb0; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaRepresentanteNit_id_5ae16bb0" ON public.contabilidad_empresa USING btree ("empresaRepresentanteNit_id");


--
-- TOC entry 3005 (class 1259 OID 18020)
-- Name: contabilidad_empresa_empresaRepresentanteNit_id_5ae16bb0_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaRepresentanteNit_id_5ae16bb0_like" ON public.contabilidad_empresa USING btree ("empresaRepresentanteNit_id" varchar_pattern_ops);


--
-- TOC entry 3006 (class 1259 OID 18021)
-- Name: contabilidad_empresa_empresaRevisorNit_id_f33e49b7; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaRevisorNit_id_f33e49b7" ON public.contabilidad_empresa USING btree ("empresaRevisorNit_id");


--
-- TOC entry 3007 (class 1259 OID 18022)
-- Name: contabilidad_empresa_empresaRevisorNit_id_f33e49b7_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaRevisorNit_id_f33e49b7_like" ON public.contabilidad_empresa USING btree ("empresaRevisorNit_id" varchar_pattern_ops);


--
-- TOC entry 3008 (class 1259 OID 18023)
-- Name: contabilidad_empresa_empresaTipo_id_1a2044cd; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaTipo_id_1a2044cd" ON public.contabilidad_empresa USING btree ("empresaTipo_id");


--
-- TOC entry 3009 (class 1259 OID 18024)
-- Name: contabilidad_empresa_empresaTipo_id_1a2044cd_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_empresa_empresaTipo_id_1a2044cd_like" ON public.contabilidad_empresa USING btree ("empresaTipo_id" varchar_pattern_ops);


--
-- TOC entry 2937 (class 1259 OID 17891)
-- Name: contabilidad_formadepago_formaDePagoCodigo_33e1bd02_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_formadepago_formaDePagoCodigo_33e1bd02_like" ON public.contabilidad_formadepago USING btree ("formaDePagoCodigo" varchar_pattern_ops);


--
-- TOC entry 2940 (class 1259 OID 17892)
-- Name: contabilidad_grupo_grupoCodigo_4ab726ce_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_grupo_grupoCodigo_4ab726ce_like" ON public.contabilidad_grupo USING btree ("grupoCodigo" varchar_pattern_ops);


--
-- TOC entry 2943 (class 1259 OID 17893)
-- Name: contabilidad_mediodepago_medioDePagoCodigo_3ced2c79_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_mediodepago_medioDePagoCodigo_3ced2c79_like" ON public.contabilidad_mediodepago USING btree ("medioDePagoCodigo" varchar_pattern_ops);


--
-- TOC entry 2990 (class 1259 OID 17965)
-- Name: contabilidad_producto_productoCodigo_c5c950b2_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_producto_productoCodigo_c5c950b2_like" ON public.contabilidad_producto USING btree ("productoCodigo" varchar_pattern_ops);


--
-- TOC entry 2991 (class 1259 OID 17966)
-- Name: contabilidad_producto_productoSubgrupo_id_7b0f244e; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_producto_productoSubgrupo_id_7b0f244e" ON public.contabilidad_producto USING btree ("productoSubgrupo_id");


--
-- TOC entry 2992 (class 1259 OID 17967)
-- Name: contabilidad_producto_productoSubgrupo_id_7b0f244e_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_producto_productoSubgrupo_id_7b0f244e_like" ON public.contabilidad_producto USING btree ("productoSubgrupo_id" varchar_pattern_ops);


--
-- TOC entry 2948 (class 1259 OID 17894)
-- Name: contabilidad_regimenfiscal_regimenFiscalCodigo_e96d3b5e_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_regimenfiscal_regimenFiscalCodigo_e96d3b5e_like" ON public.contabilidad_regimenfiscal USING btree ("regimenFiscalCodigo" varchar_pattern_ops);


--
-- TOC entry 2985 (class 1259 OID 17958)
-- Name: contabilidad_resolucion_resolucionCiudad_id_f9f5c5fa; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_resolucion_resolucionCiudad_id_f9f5c5fa" ON public.contabilidad_resolucion USING btree ("resolucionCiudad_id");


--
-- TOC entry 2986 (class 1259 OID 17959)
-- Name: contabilidad_resolucion_resolucionCiudad_id_f9f5c5fa_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_resolucion_resolucionCiudad_id_f9f5c5fa_like" ON public.contabilidad_resolucion USING btree ("resolucionCiudad_id" varchar_pattern_ops);


--
-- TOC entry 2987 (class 1259 OID 17957)
-- Name: contabilidad_resolucion_resolucionCodigo_a52d1c14_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_resolucion_resolucionCodigo_a52d1c14_like" ON public.contabilidad_resolucion USING btree ("resolucionCodigo" varchar_pattern_ops);


--
-- TOC entry 2949 (class 1259 OID 17895)
-- Name: contabilidad_responsabil_responsabilidadFiscalCod_ba5160fa_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_responsabil_responsabilidadFiscalCod_ba5160fa_like" ON public.contabilidad_responsabilidadfiscal USING btree ("responsabilidadFiscalCodigo" varchar_pattern_ops);


--
-- TOC entry 2980 (class 1259 OID 17950)
-- Name: contabilidad_subgrupo_subgrupoCodigoGrupo_id_868275fa; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_subgrupo_subgrupoCodigoGrupo_id_868275fa" ON public.contabilidad_subgrupo USING btree ("subgrupoCodigoGrupo_id");


--
-- TOC entry 2981 (class 1259 OID 17951)
-- Name: contabilidad_subgrupo_subgrupoCodigoGrupo_id_868275fa_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_subgrupo_subgrupoCodigoGrupo_id_868275fa_like" ON public.contabilidad_subgrupo USING btree ("subgrupoCodigoGrupo_id" varchar_pattern_ops);


--
-- TOC entry 2982 (class 1259 OID 17949)
-- Name: contabilidad_subgrupo_subgrupoCodigo_c501306d_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_subgrupo_subgrupoCodigo_c501306d_like" ON public.contabilidad_subgrupo USING btree ("subgrupoCodigo" varchar_pattern_ops);


--
-- TOC entry 2954 (class 1259 OID 17917)
-- Name: contabilidad_tercero_terceroCiudad_id_bb2639dc; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroCiudad_id_bb2639dc" ON public.contabilidad_tercero USING btree ("terceroCiudad_id");


--
-- TOC entry 2955 (class 1259 OID 17918)
-- Name: contabilidad_tercero_terceroCiudad_id_bb2639dc_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroCiudad_id_bb2639dc_like" ON public.contabilidad_tercero USING btree ("terceroCiudad_id" varchar_pattern_ops);


--
-- TOC entry 2956 (class 1259 OID 17916)
-- Name: contabilidad_tercero_terceroCodigo_15b774c3_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroCodigo_15b774c3_like" ON public.contabilidad_tercero USING btree ("terceroCodigo" varchar_pattern_ops);


--
-- TOC entry 2957 (class 1259 OID 17919)
-- Name: contabilidad_tercero_terceroDocumento_id_48dc5afd; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroDocumento_id_48dc5afd" ON public.contabilidad_tercero USING btree ("terceroDocumento_id");


--
-- TOC entry 2958 (class 1259 OID 17920)
-- Name: contabilidad_tercero_terceroDocumento_id_48dc5afd_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroDocumento_id_48dc5afd_like" ON public.contabilidad_tercero USING btree ("terceroDocumento_id" varchar_pattern_ops);


--
-- TOC entry 2959 (class 1259 OID 17921)
-- Name: contabilidad_tercero_terceroFiscal_id_4c2e65ed; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroFiscal_id_4c2e65ed" ON public.contabilidad_tercero USING btree ("terceroFiscal_id");


--
-- TOC entry 2960 (class 1259 OID 17922)
-- Name: contabilidad_tercero_terceroFiscal_id_4c2e65ed_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroFiscal_id_4c2e65ed_like" ON public.contabilidad_tercero USING btree ("terceroFiscal_id" varchar_pattern_ops);


--
-- TOC entry 2961 (class 1259 OID 17923)
-- Name: contabilidad_tercero_terceroRegimen_id_228ca803; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroRegimen_id_228ca803" ON public.contabilidad_tercero USING btree ("terceroRegimen_id");


--
-- TOC entry 2962 (class 1259 OID 17924)
-- Name: contabilidad_tercero_terceroRegimen_id_228ca803_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroRegimen_id_228ca803_like" ON public.contabilidad_tercero USING btree ("terceroRegimen_id" varchar_pattern_ops);


--
-- TOC entry 2963 (class 1259 OID 17942)
-- Name: contabilidad_tercero_terceroTipo_id_38edcc50; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroTipo_id_38edcc50" ON public.contabilidad_tercero USING btree ("terceroTipo_id");


--
-- TOC entry 2964 (class 1259 OID 17943)
-- Name: contabilidad_tercero_terceroTipo_id_38edcc50_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tercero_terceroTipo_id_38edcc50_like" ON public.contabilidad_tercero USING btree ("terceroTipo_id" varchar_pattern_ops);


--
-- TOC entry 2965 (class 1259 OID 17925)
-- Name: contabilidad_tipoorganiz_tipoOrganizacionCodigo_c3d43841_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_tipoorganiz_tipoOrganizacionCodigo_c3d43841_like" ON public.contabilidad_tipoorganizacion USING btree ("tipoOrganizacionCodigo" varchar_pattern_ops);


--
-- TOC entry 2975 (class 1259 OID 17939)
-- Name: contabilidad_usuario_usuarioCodigo_d1ec702b_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_usuario_usuarioCodigo_d1ec702b_like" ON public.contabilidad_usuario USING btree ("usuarioCodigo" varchar_pattern_ops);


--
-- TOC entry 2976 (class 1259 OID 17940)
-- Name: contabilidad_usuario_usuarioNit_id_3ed4b138; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_usuario_usuarioNit_id_3ed4b138" ON public.contabilidad_usuario USING btree ("usuarioNit_id");


--
-- TOC entry 2977 (class 1259 OID 17941)
-- Name: contabilidad_usuario_usuarioNit_id_3ed4b138_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_usuario_usuarioNit_id_3ed4b138_like" ON public.contabilidad_usuario USING btree ("usuarioNit_id" varchar_pattern_ops);


--
-- TOC entry 2970 (class 1259 OID 17931)
-- Name: contabilidad_vendedor_vendedorCodigo_80146132_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_vendedor_vendedorCodigo_80146132_like" ON public.contabilidad_vendedor USING btree ("vendedorCodigo" varchar_pattern_ops);


--
-- TOC entry 2971 (class 1259 OID 17932)
-- Name: contabilidad_vendedor_vendedorNit_id_dd1a83a0; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_vendedor_vendedorNit_id_dd1a83a0" ON public.contabilidad_vendedor USING btree ("vendedorNit_id");


--
-- TOC entry 2972 (class 1259 OID 17933)
-- Name: contabilidad_vendedor_vendedorNit_id_dd1a83a0_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "contabilidad_vendedor_vendedorNit_id_dd1a83a0_like" ON public.contabilidad_vendedor USING btree ("vendedorNit_id" varchar_pattern_ops);


--
-- TOC entry 2909 (class 1259 OID 17713)
-- Name: django_admin_log_content_type_id_c4bce8eb; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX django_admin_log_content_type_id_c4bce8eb ON public.django_admin_log USING btree (content_type_id);


--
-- TOC entry 2912 (class 1259 OID 17714)
-- Name: django_admin_log_user_id_c564eba6; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX django_admin_log_user_id_c564eba6 ON public.django_admin_log USING btree (user_id);


--
-- TOC entry 2913 (class 1259 OID 17733)
-- Name: django_session_expire_date_a5c62663; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX django_session_expire_date_a5c62663 ON public.django_session USING btree (expire_date);


--
-- TOC entry 2916 (class 1259 OID 17732)
-- Name: django_session_session_key_c0390e0f_like; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX django_session_session_key_c0390e0f_like ON public.django_session USING btree (session_key varchar_pattern_ops);


--
-- TOC entry 3050 (class 2606 OID 17655)
-- Name: auth_group_permissions auth_group_permissio_permission_id_84c5c92e_fk_auth_perm; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissio_permission_id_84c5c92e_fk_auth_perm FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id) DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3049 (class 2606 OID 17650)
-- Name: auth_group_permissions auth_group_permissions_group_id_b120cbf9_fk_auth_group_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_group_id_b120cbf9_fk_auth_group_id FOREIGN KEY (group_id) REFERENCES public.auth_group(id) DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3048 (class 2606 OID 17641)
-- Name: auth_permission auth_permission_content_type_id_2f476e4b_fk_django_co; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_content_type_id_2f476e4b_fk_django_co FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id) DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3052 (class 2606 OID 17670)
-- Name: auth_user_groups auth_user_groups_group_id_97559544_fk_auth_group_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_group_id_97559544_fk_auth_group_id FOREIGN KEY (group_id) REFERENCES public.auth_group(id) DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3051 (class 2606 OID 17665)
-- Name: auth_user_groups auth_user_groups_user_id_6a12ed8b_fk_auth_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_user_id_6a12ed8b_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3054 (class 2606 OID 17684)
-- Name: auth_user_user_permissions auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id) DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3053 (class 2606 OID 17679)
-- Name: auth_user_user_permissions auth_user_user_permissions_user_id_a95ead1b_fk_auth_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_user_id_a95ead1b_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3083 (class 2606 OID 18080)
-- Name: contabilidad_compra contabilidad_compra_compraFormaDePago_id_b907ff46_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_compra
    ADD CONSTRAINT "contabilidad_compra_compraFormaDePago_id_b907ff46_fk_contabili" FOREIGN KEY ("compraFormaDePago_id") REFERENCES public.contabilidad_formadepago("formaDePagoCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3084 (class 2606 OID 18085)
-- Name: contabilidad_compra contabilidad_compra_compraMedioDePago_id_c7bfc061_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_compra
    ADD CONSTRAINT "contabilidad_compra_compraMedioDePago_id_c7bfc061_fk_contabili" FOREIGN KEY ("compraMedioDePago_id") REFERENCES public.contabilidad_mediodepago("medioDePagoCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3085 (class 2606 OID 18090)
-- Name: contabilidad_compra contabilidad_compra_compraResolucion_id_9eb8193f_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_compra
    ADD CONSTRAINT "contabilidad_compra_compraResolucion_id_9eb8193f_fk_contabili" FOREIGN KEY ("compraResolucion_id") REFERENCES public.contabilidad_resolucion("resolucionCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3086 (class 2606 OID 18095)
-- Name: contabilidad_compra contabilidad_compra_compraTerceroCodigo__01a13414_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_compra
    ADD CONSTRAINT "contabilidad_compra_compraTerceroCodigo__01a13414_fk_contabili" FOREIGN KEY ("compraTerceroCodigo_id") REFERENCES public.contabilidad_tercero("terceroCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3087 (class 2606 OID 18100)
-- Name: contabilidad_compra contabilidad_compra_compraVendedor_id_9e86eb27_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_compra
    ADD CONSTRAINT "contabilidad_compra_compraVendedor_id_9e86eb27_fk_contabili" FOREIGN KEY ("compraVendedor_id") REFERENCES public.contabilidad_vendedor("vendedorCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3081 (class 2606 OID 18066)
-- Name: contabilidad_compradetalle contabilidad_comprad_compraBodega_id_ee34fff5_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_compradetalle
    ADD CONSTRAINT "contabilidad_comprad_compraBodega_id_ee34fff5_fk_contabili" FOREIGN KEY ("compraBodega_id") REFERENCES public.contabilidad_bodega("bodegaCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3082 (class 2606 OID 18071)
-- Name: contabilidad_compradetalle contabilidad_comprad_compraProductoCodigo_884ab25b_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_compradetalle
    ADD CONSTRAINT "contabilidad_comprad_compraProductoCodigo_884ab25b_fk_contabili" FOREIGN KEY ("compraProductoCodigo_id") REFERENCES public.contabilidad_producto("productoCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3076 (class 2606 OID 18031)
-- Name: contabilidad_comprobantemovimiento contabilidad_comprob_centroCodigo_id_9e9de43a_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_comprobantemovimiento
    ADD CONSTRAINT "contabilidad_comprob_centroCodigo_id_9e9de43a_fk_contabili" FOREIGN KEY ("centroCodigo_id") REFERENCES public.contabilidad_centro("centroCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3077 (class 2606 OID 18036)
-- Name: contabilidad_comprobantemovimiento contabilidad_comprob_cuentaCodigo_id_f0261d89_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_comprobantemovimiento
    ADD CONSTRAINT "contabilidad_comprob_cuentaCodigo_id_f0261d89_fk_contabili" FOREIGN KEY ("cuentaCodigo_id") REFERENCES public.contabilidad_cuenta("cuentaCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3080 (class 2606 OID 18059)
-- Name: contabilidad_comprobantecabecera contabilidad_comprob_documentoCodigo_id_0858616e_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_comprobantecabecera
    ADD CONSTRAINT "contabilidad_comprob_documentoCodigo_id_0858616e_fk_contabili" FOREIGN KEY ("documentoCodigo_id") REFERENCES public.contabilidad_documento("documentoCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3078 (class 2606 OID 18041)
-- Name: contabilidad_comprobantemovimiento contabilidad_comprob_documentoCodigo_id_2d16dddc_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_comprobantemovimiento
    ADD CONSTRAINT "contabilidad_comprob_documentoCodigo_id_2d16dddc_fk_contabili" FOREIGN KEY ("documentoCodigo_id") REFERENCES public.contabilidad_documento("documentoCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3079 (class 2606 OID 18046)
-- Name: contabilidad_comprobantemovimiento contabilidad_comprob_terceroCodigo_id_51cd7a56_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_comprobantemovimiento
    ADD CONSTRAINT "contabilidad_comprob_terceroCodigo_id_51cd7a56_fk_contabili" FOREIGN KEY ("terceroCodigo_id") REFERENCES public.contabilidad_tercero("terceroCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3075 (class 2606 OID 18025)
-- Name: contabilidad_cuotadetalle contabilidad_cuotade_cuotaCodigo_id_f3f551b8_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_cuotadetalle
    ADD CONSTRAINT "contabilidad_cuotade_cuotaCodigo_id_f3f551b8_fk_contabili" FOREIGN KEY ("cuotaCodigo_id") REFERENCES public.contabilidad_cuota("cuotaId") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3067 (class 2606 OID 17968)
-- Name: contabilidad_empresa contabilidad_empresa_empresaCiudad_id_2b55b473_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_empresa
    ADD CONSTRAINT "contabilidad_empresa_empresaCiudad_id_2b55b473_fk_contabili" FOREIGN KEY ("empresaCiudad_id") REFERENCES public.contabilidad_ciudad("ciudadCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3068 (class 2606 OID 17973)
-- Name: contabilidad_empresa contabilidad_empresa_empresaContadorNit_i_2aed5911_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_empresa
    ADD CONSTRAINT "contabilidad_empresa_empresaContadorNit_i_2aed5911_fk_contabili" FOREIGN KEY ("empresaContadorNit_id") REFERENCES public.contabilidad_tercero("terceroCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3069 (class 2606 OID 17978)
-- Name: contabilidad_empresa contabilidad_empresa_empresaDocumento_id_3886112f_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_empresa
    ADD CONSTRAINT "contabilidad_empresa_empresaDocumento_id_3886112f_fk_contabili" FOREIGN KEY ("empresaDocumento_id") REFERENCES public.contabilidad_documentoidentificacion("documentoIdentificacionCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3070 (class 2606 OID 17983)
-- Name: contabilidad_empresa contabilidad_empresa_empresaFiscal_id_af71a018_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_empresa
    ADD CONSTRAINT "contabilidad_empresa_empresaFiscal_id_af71a018_fk_contabili" FOREIGN KEY ("empresaFiscal_id") REFERENCES public.contabilidad_responsabilidadfiscal("responsabilidadFiscalCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3071 (class 2606 OID 17988)
-- Name: contabilidad_empresa contabilidad_empresa_empresaRegimen_id_999609f6_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_empresa
    ADD CONSTRAINT "contabilidad_empresa_empresaRegimen_id_999609f6_fk_contabili" FOREIGN KEY ("empresaRegimen_id") REFERENCES public.contabilidad_regimenfiscal("regimenFiscalCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3072 (class 2606 OID 17993)
-- Name: contabilidad_empresa contabilidad_empresa_empresaRepresentante_5ae16bb0_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_empresa
    ADD CONSTRAINT "contabilidad_empresa_empresaRepresentante_5ae16bb0_fk_contabili" FOREIGN KEY ("empresaRepresentanteNit_id") REFERENCES public.contabilidad_tercero("terceroCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3073 (class 2606 OID 17998)
-- Name: contabilidad_empresa contabilidad_empresa_empresaRevisorNit_id_f33e49b7_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_empresa
    ADD CONSTRAINT "contabilidad_empresa_empresaRevisorNit_id_f33e49b7_fk_contabili" FOREIGN KEY ("empresaRevisorNit_id") REFERENCES public.contabilidad_tercero("terceroCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3074 (class 2606 OID 18003)
-- Name: contabilidad_empresa contabilidad_empresa_empresaTipo_id_1a2044cd_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_empresa
    ADD CONSTRAINT "contabilidad_empresa_empresaTipo_id_1a2044cd_fk_contabili" FOREIGN KEY ("empresaTipo_id") REFERENCES public.contabilidad_tipoorganizacion("tipoOrganizacionCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3066 (class 2606 OID 17960)
-- Name: contabilidad_producto contabilidad_product_productoSubgrupo_id_7b0f244e_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_producto
    ADD CONSTRAINT "contabilidad_product_productoSubgrupo_id_7b0f244e_fk_contabili" FOREIGN KEY ("productoSubgrupo_id") REFERENCES public.contabilidad_subgrupo("subgrupoCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3065 (class 2606 OID 17952)
-- Name: contabilidad_resolucion contabilidad_resoluc_resolucionCiudad_id_f9f5c5fa_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_resolucion
    ADD CONSTRAINT "contabilidad_resoluc_resolucionCiudad_id_f9f5c5fa_fk_contabili" FOREIGN KEY ("resolucionCiudad_id") REFERENCES public.contabilidad_ciudad("ciudadCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3064 (class 2606 OID 17944)
-- Name: contabilidad_subgrupo contabilidad_subgrup_subgrupoCodigoGrupo__868275fa_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_subgrupo
    ADD CONSTRAINT "contabilidad_subgrup_subgrupoCodigoGrupo__868275fa_fk_contabili" FOREIGN KEY ("subgrupoCodigoGrupo_id") REFERENCES public.contabilidad_grupo("grupoCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3058 (class 2606 OID 17896)
-- Name: contabilidad_tercero contabilidad_tercero_terceroCiudad_id_bb2639dc_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_tercero
    ADD CONSTRAINT "contabilidad_tercero_terceroCiudad_id_bb2639dc_fk_contabili" FOREIGN KEY ("terceroCiudad_id") REFERENCES public.contabilidad_ciudad("ciudadCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3059 (class 2606 OID 17901)
-- Name: contabilidad_tercero contabilidad_tercero_terceroDocumento_id_48dc5afd_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_tercero
    ADD CONSTRAINT "contabilidad_tercero_terceroDocumento_id_48dc5afd_fk_contabili" FOREIGN KEY ("terceroDocumento_id") REFERENCES public.contabilidad_documentoidentificacion("documentoIdentificacionCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3060 (class 2606 OID 17906)
-- Name: contabilidad_tercero contabilidad_tercero_terceroFiscal_id_4c2e65ed_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_tercero
    ADD CONSTRAINT "contabilidad_tercero_terceroFiscal_id_4c2e65ed_fk_contabili" FOREIGN KEY ("terceroFiscal_id") REFERENCES public.contabilidad_responsabilidadfiscal("responsabilidadFiscalCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3061 (class 2606 OID 17911)
-- Name: contabilidad_tercero contabilidad_tercero_terceroRegimen_id_228ca803_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_tercero
    ADD CONSTRAINT "contabilidad_tercero_terceroRegimen_id_228ca803_fk_contabili" FOREIGN KEY ("terceroRegimen_id") REFERENCES public.contabilidad_regimenfiscal("regimenFiscalCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3057 (class 2606 OID 17820)
-- Name: contabilidad_tercero contabilidad_tercero_terceroTipo_id_38edcc50_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_tercero
    ADD CONSTRAINT "contabilidad_tercero_terceroTipo_id_38edcc50_fk_contabili" FOREIGN KEY ("terceroTipo_id") REFERENCES public.contabilidad_tipoorganizacion("tipoOrganizacionCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3063 (class 2606 OID 17934)
-- Name: contabilidad_usuario contabilidad_usuario_usuarioNit_id_3ed4b138_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_usuario
    ADD CONSTRAINT "contabilidad_usuario_usuarioNit_id_3ed4b138_fk_contabili" FOREIGN KEY ("usuarioNit_id") REFERENCES public.contabilidad_tercero("terceroCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3062 (class 2606 OID 17926)
-- Name: contabilidad_vendedor contabilidad_vendedo_vendedorNit_id_dd1a83a0_fk_contabili; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contabilidad_vendedor
    ADD CONSTRAINT "contabilidad_vendedo_vendedorNit_id_dd1a83a0_fk_contabili" FOREIGN KEY ("vendedorNit_id") REFERENCES public.contabilidad_tercero("terceroCodigo") DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3055 (class 2606 OID 17703)
-- Name: django_admin_log django_admin_log_content_type_id_c4bce8eb_fk_django_co; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_content_type_id_c4bce8eb_fk_django_co FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id) DEFERRABLE INITIALLY DEFERRED;


--
-- TOC entry 3056 (class 2606 OID 17708)
-- Name: django_admin_log django_admin_log_user_id_c564eba6_fk_auth_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_user_id_c564eba6_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id) DEFERRABLE INITIALLY DEFERRED;


-- Completed on 2021-02-04 18:18:49

--
-- PostgreSQL database dump complete
--

