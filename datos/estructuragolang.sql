
CREATE TABLE bodega (
    codigo character varying(2) PRIMARY KEY,
    nombre character varying(50) NOT NULL
);
CREATE TABLE unidaddemedida (
    codigo character varying(8) PRIMARY KEY,
    nombre character varying(50) NOT NULL
);

CREATE TABLE centro (
    codigo character varying(2) PRIMARY KEY,
    nombre character varying(50) NOT NULL
);

CREATE TABLE formadepago (
    codigo character varying(1) PRIMARY KEY,
    nombre character varying(50) NOT NULL
);


CREATE TABLE mediodepago (
    codigo character varying(3) PRIMARY KEY,
    nombre character varying(50) NOT NULL
);

CREATE TABLE ciudad (
    codigo character varying(5) PRIMARY KEY,
    codigociudad character varying(3) NOT NULL,
    codigodepartamento character varying(2) NOT NULL,
    nombre character varying(150) NOT NULL,
    nombreciudad character varying(150) NOT NULL,
    nombredepartamento character varying(150) NOT NULL
);


CREATE TABLE cuentahorizontal (
    codigo character varying(10) PRIMARY KEY,
    nombre character varying(80) NOT NULL,
    auto character varying(5) NOT NULL,
    nivel character varying(1) NOT NULL,
    tercero character varying(2) NOT NULL,
    centro character varying(2) NOT NULL,
    factura character varying(2) NOT NULL,
    informe character varying(3) NOT NULL,
    contra character varying(8) NOT NULL,
    interes character varying(8) NOT NULL,
    cuota character varying(2) NOT NULL,
    intereses character varying(2) NOT NULL,
    tipo character varying(6) NOT NULL,
    grupo character varying(2) NOT NULL,
    contranombre character varying(80) NOT NULL,
    interesnombre character varying(80) NOT NULL
);

CREATE TABLE cuenta (
    codigo character varying(10) PRIMARY KEY,
    nombre character varying(80) NOT NULL,
    auto character varying(5) NOT NULL,
    nivel character varying(1) NOT NULL,
    informe character varying(3) NOT NULL
    
);


CREATE TABLE documento (
    codigo character varying(2) PRIMARY KEY,
    nombre character varying(50) NOT NULL,
    consecutivo character varying(2) NOT NULL,
    inicial character varying(6)
);


CREATE TABLE documentoidentificacion (
    codigo character varying(2) PRIMARY KEY,
    nombre character varying(100) NOT NULL
);


CREATE TABLE grupo (
    codigo character varying(2) PRIMARY KEY,
    nombre character varying(50) NOT NULL
);


CREATE TABLE regimenfiscal (
    codigo character varying(2) PRIMARY KEY,
    nombre character varying(100) NOT NULL
);


CREATE TABLE responsabilidadfiscal (
    codigo character varying(10) PRIMARY KEY,
    nombre character varying(100) NOT NULL
);


CREATE TABLE tercero (
    codigo character varying(15) PRIMARY KEY,
    dv character varying(1) NOT NULL,
    nombre character varying(100),
    juridica character varying(120) NOT NULL,
    primernombre character varying(60) NOT NULL,
    segundonombre character varying(60) NOT NULL,
    primerapellido character varying(60) NOT NULL,
    segundoapellido character varying(60) NOT NULL,
    direccion character varying(80) NOT NULL,
    barrio character varying(60) NOT NULL,
    telefono1 character varying(20) NOT NULL,
    telefono2 character varying(20) NOT NULL,
    email1 character varying(100) NOT NULL,
    email2 character varying(100) NOT NULL,
    contacto character varying(60) NOT NULL,
    rut character varying(2) NOT NULL,
    ciudad character varying(5) NOT NULL,
    documento character varying(2) NOT NULL,
    fiscal character varying(10) NOT NULL,
    regimen character varying(2) NOT NULL,
    tipo character varying(1) NOT NULL,
	ica character varying(6) NOT NULL
);

CREATE TABLE tercerohorizontal (
    codigo character varying(15) PRIMARY KEY,
    dv character varying(1) NOT NULL,
    nombre character varying(100),
    juridica character varying(120) NOT NULL,
    primernombre character varying(60) NOT NULL,
    segundonombre character varying(60) NOT NULL,
    primerapellido character varying(60) NOT NULL,
    segundoapellido character varying(60) NOT NULL,
    direccion character varying(80) NOT NULL,
    barrio character varying(60) NOT NULL,
    telefono1 character varying(20) NOT NULL,
    telefono2 character varying(20) NOT NULL,
    email1 character varying(100) NOT NULL,
    email2 character varying(100) NOT NULL,
    contacto character varying(60) NOT NULL,
    rut character varying(2) NOT NULL,
    bloque character varying(10) NOT NULL,
    piso character varying(10) NOT NULL,
    apartamento character varying(10) NOT NULL,
    descuento1 character varying(12) NOT NULL,
    descuento2 character varying(12) NOT NULL,
    cuotap character varying(12) NOT NULL,
    cuota1 character varying(12) NOT NULL,
    cuota2 character varying(12) NOT NULL,
    cuota3 character varying(12) NOT NULL,
    area character varying(6) NOT NULL,
    factor character varying(5) NOT NULL,
    matricula character varying(20) NOT NULL,
    catastral character varying(20) NOT NULL,
    banco character varying(20) NOT NULL,
    phcodigo character varying(15) NOT NULL,
    phdv character varying(1) NOT NULL,
    phnombre character varying(100),
    ciudad character varying(5) NOT NULL,
    documento character varying(2) NOT NULL,
    fiscal character varying(10) NOT NULL,
    regimen character varying(2) NOT NULL,
    tipo character varying(1) NOT NULL,
	ica character varying(6) NOT NULL
);


CREATE TABLE tipoorganizacion (
    codigo character varying(1) PRIMARY KEY,
    nombre character varying(100) NOT NULL
);

CREATE TABLE cuota (
    cuotaId integer NOT NULL,
    cuotaPrestamo character varying(10) NOT NULL,
    cuotaPlazo character varying(10) NOT NULL,
    cuotaInteres character varying(10) NOT NULL,
    cuotaFecha timestamp with time zone NOT NULL,
    cuotaPago character varying(20) NOT NULL,
    cuotaTotalInteres character varying(20) NOT NULL,
    cuotaTotalPago character varying(20) NOT NULL
);

CREATE TABLE empresa (
    codigo character varying(15) PRIMARY KEY,
    dv character varying(1) NOT NULL,
    nombre character varying(100) NOT NULL,
    iva character varying(20) NOT NULL,
    reteIva character varying(20) NOT NULL,
    direccion character varying(80) NOT NULL,
    actividadIca character varying(60) NOT NULL,
    telefono1 character varying(20) NOT NULL,
    telefono2 character varying(20) NOT NULL,
    email1 character varying(100) NOT NULL,
    email2 character varying(100) NOT NULL,
    activa character varying(2) NOT NULL,
    licencia character varying(20) NOT NULL,
    representanteDv character varying(1) NOT NULL,
    representanteNombre character varying(100),
    contadorDv character varying(1) NOT NULL,
    contadorNombre character varying(100),
    revisorDv character varying(1) NOT NULL,
    revisorNombre character varying(100),
    ciudad character varying(5) NOT NULL,
    contadorNit character varying(15) NOT NULL,
    documento character varying(2) NOT NULL,
    fiscal character varying(10) NOT NULL,
    regimen character varying(2) NOT NULL,
    representanteNit character varying(15) NOT NULL,
    revisorNit character varying(15) NOT NULL,
    tipo character varying(1) NOT NULL
);

CREATE TABLE usuario (
    codigo character varying(2) PRIMARY KEY,
    dv character varying(1) NOT NULL,
    nombre character varying(100),
    tipo character varying(50) NOT NULL,
    clave2 character varying(15) NOT NULL,
    clave1 character varying(15) NOT NULL,
    correo1 character varying(100) NOT NULL,
    correo2 character varying(100) NOT NULL,
    nit character varying(15) NOT NULL
);

CREATE TABLE subgrupo (
    codigo character varying(4) PRIMARY KEY,
    nombre character varying(50) NOT NULL,
    grupo character varying(2) NOT NULL
);

CREATE TABLE vendedor (
    codigo character varying(2) PRIMARY KEY,
    dv character varying(1) NOT NULL,
    nombre character varying(100),
    comision character varying(5) NOT NULL,
    nit character varying(15) NOT NULL
);

CREATE TABLE producto (
    codigo character varying(15) PRIMARY KEY,
    nombre character varying(100) NOT NULL,
    unidad character varying(12) NOT NULL,
    iva character varying(6) NOT NULL,
    tipo character varying(10) NOT NULL,
    venta character varying(12) NOT NULL,
    costo character varying(12) NOT NULL,
    cantidad character varying(12) NOT NULL,
    total character varying(12) NOT NULL,
    subgrupo character varying(4) NOT NULL
);

CREATE TABLE resolucionventa (
    codigo character varying(2) PRIMARY KEY,
    numero character varying(20) NOT NULL,
    prefijo character varying(4) NOT NULL,
    tipo character varying(20) NOT NULL,
    local character varying(50) NOT NULL,
    direccion character varying(50) NOT NULL,
    telefono character varying(20) NOT NULL,
    informe character varying(50) NOT NULL,
    fechaInicial timestamp without time zone NOT NULL,
    fechaFinal timestamp without time zone NOT NULL,
    numeroInicial character varying(50) NOT NULL,
    numeroFinal character varying(50) NOT NULL,
    numeroActual character varying(50) NOT NULL,
    ciudad character varying(5) NOT NULL,
	clavetecnica character varying(200) NOT NULL,
	idesoftware character varying(200) NOT NULL,
	testid character varying(200) NOT NULL,
	pin character varying(20) NOT NULL,
	ambiente character varying(1) NOT NULL
);

CREATE TABLE resolucionsoporte (
    codigo character varying(2) PRIMARY KEY,
    numero character varying(20) NOT NULL,
    prefijo character varying(4) NOT NULL,
    tipo character varying(20) NOT NULL,
    local character varying(50) NOT NULL,
    direccion character varying(50) NOT NULL,
    telefono character varying(20) NOT NULL,
    informe character varying(50) NOT NULL,
    fechaInicial timestamp without time zone NOT NULL,
    fechaFinal timestamp without time zone NOT NULL,
    numeroInicial character varying(50) NOT NULL,
    numeroFinal character varying(50) NOT NULL,
    numeroActual character varying(50) NOT NULL,
    ciudad character varying(5) NOT NULL,
	clavetecnica character varying(200) NOT NULL,
	idesoftware character varying(200) NOT NULL,
	testid character varying(200) NOT NULL,
	pin character varying(20) NOT NULL,
	ambiente character varying(1) NOT NULL
);

CREATE TABLE venta (
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
	subtotal19 character varying(20) NOT NULL,
    subtotal5 character varying(20) NOT NULL,
    subtotal0 character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
	ret2201 character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    vendedor character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
    resolucion character varying(2) NOT NULL,
	cotizacion character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
);



CREATE TABLE ventadetalle (
    id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	fecha date NOT NULL
);


CREATE TABLE compra (
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
    subtotal19 character varying(20) NOT NULL,
	subtotal5 character varying(20) NOT NULL,
	subtotal0 character varying(20) NOT NULL,
	subtotal character varying(20) NOT NULL,
	subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    porcentajeretencionfuente character varying(6) NOT NULL,
    totalretencionfuente character varying(20) NOT NULL,
    porcentajeretencionica character varying(6) NOT NULL,
    totalretencionica character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    almacenista character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	pedido character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	centro character varying(2) NOT NULL
);


CREATE TABLE compradetalle (
	id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	fecha date NOT NULL
);


CREATE TABLE comprobante (
    documento character varying(2) NOT NULL,
	numero character varying(20) NOT NULL,
    fecha date NOT NULL,
	fechaconsignacion date NOT NULL,
	periodo character varying(4),
	licencia character varying(4),
	usuario character varying(2) NOT NULL,
	estado character varying(20),
	debito character varying(20),
	credito character varying(20)
	
);


CREATE TABLE comprobantedetalle (
    fila character varying(4),
	cuenta character varying(10),
	tercero character varying(15),
	centro character varying(2),
	factura character varying(15),
	concepto character varying(120),
	debito character varying(20),
	credito character varying(20),
	documento character varying(2) NOT NULL,
	numero character varying(20) NOT NULL,
    fecha date NOT NULL,
	fechaconsignacion date NOT NULL
		
);

CREATE TABLE almacenista (
    codigo character varying(2) PRIMARY KEY,
    dv character varying(1) NOT NULL,
    nombre character varying(100),
    nit character varying(15) NOT NULL
);


CREATE TABLE venta (
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
	subtotal19 character varying(20) NOT NULL,
    subtotal5 character varying(20) NOT NULL,
    subtotal0 character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
	ret2201 character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    vendedor character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
    resolucion character varying(2) NOT NULL,
	cotizacion character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
);



CREATE TABLE ventadetalle (
    id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	fecha date NOT NULL
);


CREATE TABLE ventaservicio (
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
	subtotal19 character varying(20) NOT NULL,
    subtotal5 character varying(20) NOT NULL,
    subtotal0 character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
	ret2201 character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    vendedor character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
    resolucion character varying(2) NOT NULL,
	cotizacion character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
);



CREATE TABLE ventaserviciodetalle (
    id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    nombreservicio character varying(500) NOT NULL,
	unidadservicio character varying(6) NOT NULL,
	codigoservicio character varying(4) NOT NULL,
	ivaservicio character varying(2) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	fecha date NOT NULL,
	
);


CREATE TABLE compra (
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
    subtotal19 character varying(20) NOT NULL,
	subtotal5 character varying(20) NOT NULL,
	subtotal0 character varying(20) NOT NULL,
	subtotal character varying(20) NOT NULL,
	subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    porcentajeretencionfuente character varying(6) NOT NULL,
    totalretencionfuente character varying(20) NOT NULL,
    porcentajeretencionica character varying(6) NOT NULL,
    totalretencionica character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    almacenista character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	pedido character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
);


CREATE TABLE compradetalle (
	id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	fecha date NOT NULL
);



CREATE TABLE cotizacion (
	codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
	subtotal19 character varying(20) NOT NULL,
    subtotal5 character varying(20) NOT NULL,
    subtotal0 character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
	ret2201 character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    vendedor character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
 	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
	    
   );

CREATE TABLE cotizaciondetalle (
	id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	fecha date NOT NULL
    
);


CREATE TABLE cotizacionservicio (
	codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
	subtotal19 character varying(20) NOT NULL,
    subtotal5 character varying(20) NOT NULL,
    subtotal0 character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
	ret2201 character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    vendedor character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
 	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
	    
   );

CREATE TABLE cotizacionserviciodetalle (
	id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
	nombreservicio character varying(500) NOT NULL,
	unidadservicio character varying(6) NOT NULL,
	codigoservicio character varying(4) NOT NULL,
	ivaservicio character varying(2) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	fecha date NOT NULL
    
);

CREATE TABLE pedido (
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
    subtotal19 character varying(20) NOT NULL,
	subtotal5 character varying(20) NOT NULL,
	subtotal0 character varying(20) NOT NULL,
	subtotal character varying(20) NOT NULL,
	subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    porcentajeretencionfuente character varying(6) NOT NULL,
    totalretencionfuente character varying(20) NOT NULL,
    porcentajeretencionica character varying(6) NOT NULL,
    totalretencionica character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    almacenista character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
);


CREATE TABLE pedidodetalle (
    id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	fecha date NOT NULL
);

CREATE TABLE pedidosoporte (
   	codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
    subtotal19 character varying(20) NOT NULL,
	subtotal5 character varying(20) NOT NULL,
	subtotal0 character varying(20) NOT NULL,
	subtotal character varying(20) NOT NULL,
	subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    porcentajeretencionfuente character varying(6) NOT NULL,
    totalretencionfuente character varying(20) NOT NULL,
    porcentajeretencionica character varying(6) NOT NULL,
    totalretencionica character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    almacenista character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
);


CREATE TABLE pedidosoportedetalle (
    id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	fecha date NOT NULL
);

CREATE TABLE pedidosoporteservicio (
	codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
    subtotal19 character varying(20) NOT NULL,
	subtotal5 character varying(20) NOT NULL,
	subtotal0 character varying(20) NOT NULL,
	subtotal character varying(20) NOT NULL,
	subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    porcentajeretencionfuente character varying(6) NOT NULL,
    totalretencionfuente character varying(20) NOT NULL,
    porcentajeretencionica character varying(6) NOT NULL,
    totalretencionica character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    almacenista character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
);


CREATE TABLE pedidosoporteserviciodetalle (
	id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
  	nombreservicio character varying(500) NOT NULL,
	unidadservicio character varying(6) NOT NULL,
	codigoservicio character varying(4) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	fecha date NOT NULL
    	
);



CREATE TABLE devolucioncompra (
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
    subtotal19 character varying(20) NOT NULL,
	subtotal5 character varying(20) NOT NULL,
	subtotal0 character varying(20) NOT NULL,
	subtotal character varying(20) NOT NULL,
	subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    porcentajeretencionfuente character varying(6) NOT NULL,
    totalretencionfuente character varying(20) NOT NULL,
    porcentajeretencionica character varying(6) NOT NULL,
    totalretencionica character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    almacenista character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	compra character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
);


CREATE TABLE devolucioncompradetalle (
    id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	fecha date NOT NULL
);

CREATE TABLE devolucionventa (
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
	subtotal19 character varying(20) NOT NULL,
    subtotal5 character varying(20) NOT NULL,
    subtotal0 character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
	ret2201 character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    vendedor character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	resolucion character varying(2) NOT NULL,
	venta character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
		
);

CREATE TABLE devolucionventadetalle (
    id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	fecha date NOT NULL
);

CREATE TABLE devolucionventaservicio (
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
	subtotal19 character varying(20) NOT NULL,
    subtotal5 character varying(20) NOT NULL,
    subtotal0 character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
	ret2201 character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    vendedor character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	resolucion character varying(2) NOT NULL,
	venta character varying(20) NOT NULL,
	tipo character varying(30) NOT NULL,
	centro character varying(2) NOT NULL
		
);

CREATE TABLE devolucionventaserviciodetalle (
    id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
	nombreservicio character varying(500) NOT NULL,
	unidadservicio character varying(6) NOT NULL,
	codigoservicio character varying(4) NOT NULL,
	ivaservicio character varying(2) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	fecha date NOT NULL
);

CREATE TABLE inventarioinicial (
	items character varying(6) NOT NULL,
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
	almacenista character varying(4) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalbase5 character varying(20) NOT NULL,
    subtotalbase0 character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
	centro character varying(2) NOT NULL
);


CREATE TABLE inventarioinicialdetalle (
    id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character varying(5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	fecha date NOT NULL
);



CREATE TABLE soporte (
    resolucion character varying(2) NOT NULL,
	codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
    subtotal19 character varying(20) NOT NULL,
	subtotal5 character varying(20) NOT NULL,
	subtotal0 character varying(20) NOT NULL,
	subtotal character varying(20) NOT NULL,
	subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    porcentajeretencionfuente character varying(6) NOT NULL,
    totalretencionfuente character varying(20) NOT NULL,
    porcentajeretencionica character varying(6) NOT NULL,
    totalretencionica character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    almacenista character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	pedidosoporte character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	centro character varying(2) NOT NULL
);


CREATE TABLE soportedetalle (
	id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	fecha date NOT NULL
);

CREATE TABLE soporteservicio (
    resolucion character varying(2) NOT NULL,
	codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
    subtotal19 character varying(20) NOT NULL,
	subtotal5 character varying(20) NOT NULL,
	subtotal0 character varying(20) NOT NULL,
	subtotal character varying(20) NOT NULL,
	subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    porcentajeretencionfuente character varying(6) NOT NULL,
    totalretencionfuente character varying(20) NOT NULL,
    porcentajeretencionica character varying(6) NOT NULL,
    totalretencionica character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    almacenista character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	pedidosoporteservicio character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	centro character varying(2) NOT NULL
);
	

CREATE TABLE soporteserviciodetalle (
	id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
  	nombreservicio character varying(500) NOT NULL,
	unidadservicio character varying(6) NOT NULL,
	codigoservicio character varying(4) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	fecha date NOT NULL
		
);


CREATE TABLE devolucionsoporte (
	resolucion character varying(2) NOT NULL,
   	codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
    subtotal19 character varying(20) NOT NULL,
	subtotal5 character varying(20) NOT NULL,
	subtotal0 character varying(20) NOT NULL,
	subtotal character varying(20) NOT NULL,
	subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    porcentajeretencionfuente character varying(6) NOT NULL,
    totalretencionfuente character varying(20) NOT NULL,
    porcentajeretencionica character varying(6) NOT NULL,
    totalretencionica character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    almacenista character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	soporte character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	centro character varying(2) NOT NULL
);


CREATE TABLE devolucionsoportedetalle (
	id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
  	nombreservicio character varying(500) NOT NULL,
	unidadservicio character varying(6) NOT NULL,
	codigoservicio character varying(4) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	fecha date NOT NULL
	
);

CREATE TABLE devolucionsoporteservicio (
	resolucion character varying(2) NOT NULL,
   	codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    vence date NOT NULL,
    hora character varying(20) NOT NULL,
    subtotal19 character varying(20) NOT NULL,
	subtotal5 character varying(20) NOT NULL,
	subtotal0 character varying(20) NOT NULL,
	subtotal character varying(20) NOT NULL,
	subtotalIva19 character varying(20) NOT NULL,
    subtotalIva5 character varying(20) NOT NULL,
    subtotalIva0 character varying(20) NOT NULL,
    subtotalBase19 character varying(20) NOT NULL,
    subtotalBase5 character varying(20) NOT NULL,
    subtotalBase0 character varying(20) NOT NULL,
    totaliva character varying(20) NOT NULL,
    total character varying(20) NOT NULL,
    porcentajeretencionfuente character varying(6) NOT NULL,
    totalretencionfuente character varying(20) NOT NULL,
    porcentajeretencionica character varying(6) NOT NULL,
    totalretencionica character varying(20) NOT NULL,
    neto character varying(20) NOT NULL,
    items character varying(6) NOT NULL,
    formaDePago character varying(2) NOT NULL,
    medioDePago character varying(3) NOT NULL,
    tercero character varying(15) NOT NULL,
    almacenista character varying(4) NOT NULL,
    descuento character varying(20) NOT NULL,
	subtotaldescuento19 character varying(20) NOT NULL,
    subtotaldescuento5 character varying(20) NOT NULL,
    subtotaldescuento0 character varying(20) NOT NULL,
	soporteservicio character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	centro character varying(2) NOT NULL
);


CREATE TABLE devolucionsoporteserviciodetalle (
	id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
    subtotal character varying(20) NOT NULL,
    pagina character varying(4) NOT NULL,
    bodega character varying(2) NOT NULL,
  	nombreservicio character varying(500) NOT NULL,
	unidadservicio character varying(6) NOT NULL,
	codigoservicio character varying(4) NOT NULL,
    descuento character varying(6) NOT NULL,
    montodescuento character varying(20) NOT NULL,
    sigratis character varying(2) NOT NULL,
    subtotaldescuento character varying(20) NOT NULL,
	tipo character varying(20) NOT NULL,
	fecha date NOT NULL
);

CREATE TABLE traslado (
    codigo character varying(20) NOT NULL,
    fecha date NOT NULL,
    items character varying(6) NOT NULL,
    almacenista character varying(4) NOT NULL,
    tipo character varying(30) NOT NULL
	
);


CREATE TABLE trasladodetalle (
	id character varying(5) NOT NULL,
    codigo character varying(20) NOT NULL,
    fila character (5) NOT NULL,
    entra character varying(20) NOT NULL,
    sale character varying(20) NOT NULL,
    bodega character varying(2) NOT NULL,
    Producto character varying(15) NOT NULL,
    tipo character varying(30) NOT NULL,
	fecha date NOT NULL
);


CREATE TABLE configuracioninventario (
    compracuenta19 character varying(10) NOT NULL,
    compranombre19 character varying(50) NOT NULL,
	compracuenta5 character varying(10) NOT NULL,
    compranombre5 character varying(50) NOT NULL,
	compracuenta0 character varying(10) NOT NULL,
    compranombre0 character varying(50) NOT NULL,
	compraiva19 character varying(10) NOT NULL,
    compranombreiva19 character varying(50) NOT NULL,
	compraiva5 character varying(10) NOT NULL,
    compranombreiva5 character varying(50) NOT NULL,
	compracuentaretfte character varying(10) NOT NULL,
    compranombreretfte character varying(50) NOT NULL,
	compracuentaretica character varying(10) NOT NULL,
    compranombreretica character varying(50) NOT NULL,
	compracuentaretotra character varying(10) NOT NULL,
    compranombreretotra character varying(50) NOT NULL,
	compracuentadescuento character varying(10) NOT NULL,
    compranombredescuento character varying(50) NOT NULL,
	compradevolucioncuenta19 character varying(10) NOT NULL,
    compradevolucionnombre19 character varying(50) NOT NULL,
	compradevolucioncuenta5 character varying(10) NOT NULL,
    compradevolucionnombre5 character varying(50) NOT NULL,
	compradevolucioncuenta0 character varying(10) NOT NULL,
    compradevolucionnombre0 character varying(50) NOT NULL,
	compradevolucioniva19 character varying(10) NOT NULL,
    compradevolucionnombreiva19 character varying(50) NOT NULL,
	compradevolucioniva5 character varying(10) NOT NULL,
    compradevolucionnombreiva5 character varying(50) NOT NULL,
	compradevolucioncuentaretfte character varying(10) NOT NULL,
    compradevolucionnombreretfte character varying(50) NOT NULL,
	compradevolucioncuentaretica character varying(10) NOT NULL,
    compradevolucionnombreretica character varying(50) NOT NULL,
	compradevolucioncuentaretotra character varying(10) NOT NULL,
    compradevolucionnombreretotra character varying(50) NOT NULL,
	compradevolucioncuentadescuento character varying(10) NOT NULL,
    compradevolucionnombredescuento character varying(50) NOT NULL,
	compracuentaproveedor character varying(10) NOT NULL,
    compranombreproveedor character varying(50) NOT NULL,
	compracuentabase character varying(10) NOT NULL,
    compracuentaporcentajeretfte character varying(10) NOT NULL,
    compracuentaporcentajeretotra character varying(10) NOT NULL,
	ventacuenta19 character varying(10) NOT NULL,
    ventanombre19 character varying(50) NOT NULL,
	ventacuenta5 character varying(10) NOT NULL,
    ventanombre5 character varying(50) NOT NULL,
	ventacuenta0 character varying(10) NOT NULL,
    ventanombre0 character varying(50) NOT NULL,
	ventaiva19 character varying(10) NOT NULL,
    ventanombreiva19 character varying(50) NOT NULL,
	ventaiva5 character varying(10) NOT NULL,
    ventanombreiva5 character varying(50) NOT NULL,
	ventacuentaret2201 character varying(10) NOT NULL,
    ventanombreret2201 character varying(50) NOT NULL,
	ventacuentadescuento character varying(10) NOT NULL,
    ventanombredescuento character varying(50) NOT NULL,
	ventadevolucioncuenta19 character varying(10) NOT NULL,
    ventadevolucionnombre19 character varying(50) NOT NULL,
	ventadevolucioncuenta5 character varying(10) NOT NULL,
    ventadevolucionnombre5 character varying(50) NOT NULL,
	ventadevolucioncuenta0 character varying(10) NOT NULL,
    ventadevolucionnombre0 character varying(50) NOT NULL,
	ventadevolucioniva19 character varying(10) NOT NULL,
    ventadevolucionnombreiva19 character varying(50) NOT NULL,
	ventadevolucioniva5 character varying(10) NOT NULL,
    ventadevolucionnombreiva5 character varying(50) NOT NULL,
	ventadevolucioncuentaret2201 character varying(10) NOT NULL,
    ventadevolucionnombreret2201 character varying(50) NOT NULL,
	ventadevolucioncuentadescuento character varying(10) NOT NULL,
    ventadevolucionnombredescuento character varying(50) NOT NULL,
	ventacuentacliente character varying(10) NOT NULL,
    ventanombrecliente character varying(50) NOT NULL,
	ventacontracuentaret2201 character varying(10) NOT NULL,
	ventacontranombreret2201 character varying(50) NOT NULL,
	ventadevolucioncontracuentaret2201 character varying(10) NOT NULL,
	ventadevolucioncontranombreret2201 character varying(50) NOT NULL,
	ventacuentaporcentajeret2201 character varying(10) NOT NULL,
	ventatipoiva character varying(10) NOT NULL,
	ventaserviciocuenta19 character varying(10) NOT NULL,
    ventaservicionombre19 character varying(50) NOT NULL,
	ventaserviciocuenta5 character varying(10) NOT NULL,
    ventaservicionombre5 character varying(50) NOT NULL,
	ventaserviciocuenta0 character varying(10) NOT NULL,
    ventaservicionombre0 character varying(50) NOT NULL,
	ventaservicioiva19 character varying(10) NOT NULL,
    ventaservicionombreiva19 character varying(50) NOT NULL,
	ventaservicioiva5 character varying(10) NOT NULL,
    ventaservicionombreiva5 character varying(50) NOT NULL,
	ventaserviciocuentaret2201 character varying(10) NOT NULL,
    ventaservicionombreret2201 character varying(50) NOT NULL,
	ventaserviciocuentadescuento character varying(10) NOT NULL,
    ventaservicionombredescuento character varying(50) NOT NULL,
	ventaserviciodevolucioncuenta19 character varying(10) NOT NULL,
    ventaserviciodevolucionnombre19 character varying(50) NOT NULL,
	ventaserviciodevolucioncuenta5 character varying(10) NOT NULL,
    ventaserviciodevolucionnombre5 character varying(50) NOT NULL,
	ventaserviciodevolucioncuenta0 character varying(10) NOT NULL,
    ventaserviciodevolucionnombre0 character varying(50) NOT NULL,
	ventaserviciodevolucioniva19 character varying(10) NOT NULL,
    ventaserviciodevolucionnombreiva19 character varying(50) NOT NULL,
	ventaserviciodevolucioniva5 character varying(10) NOT NULL,
    ventaserviciodevolucionnombreiva5 character varying(50) NOT NULL,
	ventaserviciodevolucioncuentaret2201 character varying(10) NOT NULL,
    ventaserviciodevolucionnombreret2201 character varying(50) NOT NULL,
	ventaserviciodevolucioncuentadescuento character varying(10) NOT NULL,
    ventaserviciodevolucionnombredescuento character varying(50) NOT NULL,
	ventaserviciocuentacliente character varying(10) NOT NULL,
    ventaservicionombrecliente character varying(50) NOT NULL,
	ventaserviciocontracuentaret2201 character varying(10) NOT NULL,
	ventaserviciocontranombreret2201 character varying(50) NOT NULL,
	ventaserviciodevolucioncontracuentaret2201 character varying(10) NOT NULL,
	ventaserviciodevolucioncontranombreret2201 character varying(50) NOT NULL,
	ventaserviciocuentaporcentajeret2201 character varying(10) NOT NULL,
	ventaserviciotipoiva character varying(10) NOT NULL,
	soportecuenta19 character varying(10) NOT NULL,
    soportenombre19 character varying(50) NOT NULL,
	soportecuenta5 character varying(10) NOT NULL,
    soportenombre5 character varying(50) NOT NULL,
	soportecuenta0 character varying(10) NOT NULL,
    soportenombre0 character varying(50) NOT NULL,
	soportecuentaretfte character varying(10) NOT NULL,
    soportenombreretfte character varying(50) NOT NULL,
	soportecuentaretica character varying(10) NOT NULL,
    soportenombreretica character varying(50) NOT NULL,
	soportecuentaretotra character varying(10) NOT NULL,
    soportenombreretotra character varying(50) NOT NULL,
	soportecuentadescuento character varying(10) NOT NULL,
    soportenombredescuento character varying(50) NOT NULL,
	soportedevolucioncuenta19 character varying(10) NOT NULL,
    soportedevolucionnombre19 character varying(50) NOT NULL,
	soportedevolucioncuenta5 character varying(10) NOT NULL,
    soportedevolucionnombre5 character varying(50) NOT NULL,
	soportedevolucioncuenta0 character varying(10) NOT NULL,
    soportedevolucionnombre0 character varying(50) NOT NULL,
	soportedevolucioncuentaretfte character varying(10) NOT NULL,
    soportedevolucionnombreretfte character varying(50) NOT NULL,
	soportedevolucioncuentaretica character varying(10) NOT NULL,
    soportedevolucionnombreretica character varying(50) NOT NULL,
	soportedevolucioncuentaretotra character varying(10) NOT NULL,
    soportedevolucionnombreretotra character varying(50) NOT NULL,
	soportedevolucioncuentadescuento character varying(10) NOT NULL,
    soportedevolucionnombredescuento character varying(50) NOT NULL,
	soportecuentaproveedor character varying(10) NOT NULL,
    soportenombreproveedor character varying(50) NOT NULL,
	soportecuentabase character varying(10) NOT NULL,
    soportecuentaporcentajeretfte character varying(10) NOT NULL,
    soportecuentaporcentajeretotra character varying(10) NOT NULL,
	soporteserviciocuenta19 character varying(10) NOT NULL,
    soporteservicionombre19 character varying(50) NOT NULL,
	soporteserviciocuenta5 character varying(10) NOT NULL,
    soporteservicionombre5 character varying(50) NOT NULL,
	soporteserviciocuenta0 character varying(10) NOT NULL,
    soporteservicionombre0 character varying(50) NOT NULL,
	soporteserviciocuentaretfte character varying(10) NOT NULL,
    soporteservicionombreretfte character varying(50) NOT NULL,
	soporteserviciocuentaretica character varying(10) NOT NULL,
    soporteservicionombreretica character varying(50) NOT NULL,
	soporteserviciocuentaretotra character varying(10) NOT NULL,
    soporteservicionombreretotra character varying(50) NOT NULL,
	soporteserviciocuentadescuento character varying(10) NOT NULL,
    soporteservicionombredescuento character varying(50) NOT NULL,
	soporteserviciodevolucioncuenta character varying(10) NOT NULL,
    soporteserviciodevolucionnombre character varying(50) NOT NULL,
	soporteserviciodevolucioncuentaretfte character varying(10) NOT NULL,
    soporteserviciodevolucionnombreretfte character varying(50) NOT NULL,
	soporteserviciodevolucioncuentaretica character varying(10) NOT NULL,
    soporteserviciodevolucionnombreretica character varying(50) NOT NULL,
	soporteserviciodevolucioncuentaretotra character varying(10) NOT NULL,
    soporteserviciodevolucionnombreretotra character varying(50) NOT NULL,
	soporteserviciodevolucioncuentadescuento character varying(10) NOT NULL,
    soporteserviciodevolucionnombredescuento character varying(50) NOT NULL,
	soporteserviciocuentaproveedor character varying(10) NOT NULL,
    soporteservicionombreproveedor character varying(50) NOT NULL,
	soporteserviciocuentabase character varying(10) NOT NULL,
    soporteserviciocuentaporcentajeretfte character varying(10) NOT NULL,
    soporteserviciocuentaporcentajeretotra character varying(10) NOT NULL
	
	  
);


CREATE TABLE configuracioncontabilidad (
    pagocuentaefectivo character varying(10) NOT NULL,
    pagonombreefectivo character varying(50) NOT NULL,
	pagocuentasaldoafavor character varying(10) NOT NULL,
    pagonombresaldoafavor character varying(50) NOT NULL,
	pagocuentatdebito character varying(10) NOT NULL,
    pagonombretdebito character varying(50) NOT NULL,
	pagocuentatcredito character varying(10) NOT NULL,
    pagonombretcredito character varying(50) NOT NULL,
	pagocuentaconsignacion character varying(10) NOT NULL,
    pagonombreconsignacion character varying(50) NOT NULL,
	pagocuentamayorvalor character varying(10) NOT NULL,
    pagonombremayorvalor character varying(50) NOT NULL,
	pagocuentamenorvalor character varying(10) NOT NULL,
    pagonombremenorvalor character varying(50) NOT NULL
  
);

CREATE TABLE plandecuentaniif (
    codigo character varying(8) PRIMARY KEY,
    nombre character varying(120) NOT NULL,
	nivel character varying(1) NOT NULL
);	
	
CREATE TABLE plandecuentapuc (
    codigo character varying(8) PRIMARY KEY,
    nombre character varying(120) NOT NULL,
	nivel character varying(2) NOT NULL
);	

CREATE TABLE plandecuentaempresa (
    codigo character varying(8) PRIMARY KEY,
    nombre character varying(120) NOT NULL,
	auto character varying(6) NOT NULL,
	nivel character varying(2) NOT NULL
);	

CREATE TABLE retencionenlafuente (
    codigo character varying(2) PRIMARY KEY,
	nombre character varying(120) NOT NULL,
    valor character varying(20) NOT NULL,
	porcentaje character varying(10) NOT NULL
);

CREATE TABLE comprobante (
    documento character varying(2) NOT NULL,
	numero character varying(15) NOT NULL,
    fecha date NOT NULL,
    fechaconsignacion date NOT NULL,
    debito character varying(15) NOT NULL,
    credito character varying(15) NOT NULL
);


CREATE TABLE comprobantedetalle (
    fila character varying(5) NOT NULL,
    cuenta character varying(10) NOT NULL,
    tercero character varying(15) NOT NULL,
    centro character varying(2) NOT NULL,
    concepto character varying(50) NOT NULL,
    factura character varying(15) NOT NULL,
    debito character varying(15) NOT NULL,
    credito character varying(15) NOT NULL,
    documento character varying(2) NOT NULL,
    numero character varying(15) NOT NULL,
    fecha date NOT NULL,
	fechaconsignacion date NOT NULL
);

CREATE TABLE cuentadecobro (
  	numero character varying(15) NOT NULL,
    fecha date NOT NULL,
    tercero character varying(15) NOT NULL,
    total character varying(15) NOT NULL
);


CREATE TABLE cuentadecobrodetalle (
    fila character varying(5) NOT NULL,
    numero character varying(15) NOT NULL,
	cuenta character varying(10) NOT NULL,
	descripcion character varying(50) NOT NULL,
	anterior character varying(15) NOT NULL,
	actual character varying(15) NOT NULL,
    total character varying(15) NOT NULL
    
);


CREATE TABLE inventario (
	fecha date NOT NULL,
	tipo character varying(20) NOT NULL,
    codigo character varying(20) NOT NULL,
	bodega character varying(2) NOT NULL,
	Producto character varying(15) NOT NULL,
    cantidad character varying(20) NOT NULL,
    precio character varying(20) NOT NULL,
	operacion character varying(10) NOT NULL
    
);

CREATE TABLE concepto (
    codigo character varying(2) PRIMARY KEY,
    nombre character varying(50) NOT NULL
);

CREATE TABLE financiero (
    codigo character varying(2) PRIMARY KEY,
    nombre character varying(50) NOT NULL
);