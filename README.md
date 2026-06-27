# ReSport - Sistema de Comercio Electronico

**Autor:** Diego Medina  
**Materia:** Programacion Orientada a Objetos  
**Universidad:** Universidad Internacional del Ecuador
**Fecha:** Junio 2026  

---

## Descripcion General

ReSport es un sistema de comercio electrónico orientado a la venta de productos deportivos y de rehabilitacion. El sistema fue desarrollado en el lenguaje de programacion Go aplicando los cuatro principios fundamentales de la Programacion Orientada a Objetos: encapsulación, herencia por composición, polimorfismo mediante interfaces y abstracción mediante el patron Repository.

El sistema permite gestionar un catalogo de productos, registrar clientes con autenticación segura, administrar un carrito de compras y procesar pedidos con metodos de pago simulados. Además cuenta con un panel administrativo completo accesible desde el navegador.

---

## Objetivo del Proyecto

Desarrollar un sistema de gestion de comercio electrónico que implemente los principios de Programación Orientada a Objetos, integrando gestion de productos, autenticación de usuarios con JWT, carrito de compras y procesamiento de pedidos, mediante el lenguaje Go con arquitectura limpia por capas.

---

## Tecnologías Utilizadas

**Backend**

- Go 1.23
- Gin v1.10 (framework web HTTP)
- GORM v1.25 (ORM para Go)
- glebarez/sqlite (driver SQLite puro de Go)
- golang-jwt/jwt v5 (autenticación con JSON Web Tokens)
- golang.org/x/crypto/bcrypt (hash de contraseñas)
- joho/godotenv (variables de entorno)

**Frontend**

- HTML5
- Tailwind CSS via CDN
- JavaScript (fetch API)

**Herramientas**

- Git y GitHub (control de versiones)
- Visual Studio Code (IDE)
- DB Browser for SQLite (inspección de base de datos)

---

## Arquitectura del Sistema

El proyecto sigue una arquitectura limpia por capas donde cada capa tiene una responsabilidad específica y solo se comunica con la capa inmediatamente inferior. Esta organización permite que cambiar una capa no afecte a las demas.

```
ReSport/
  cmd/
    server/
      main.go          Punto de entrada del servidor HTTP en puerto 8081
      seed/
        main.go        Script para crear el administrador inicial
  internal/
    config/            Carga de variables de entorno desde .env
    database/          Conexion y migracion automatica de SQLite
    models/            Entidades del dominio con principios de POO
    repository/        Acceso a datos mediante interfaces (abstraccion)
    services/          Logica de negocio del sistema
    handlers/          Controladores HTTP que reciben y responden peticiones
    middleware/        Validacion de tokens JWT para rutas protegidas
  templates/           7 paginas HTML con Tailwind CSS
  static/              Archivos estaticos
  .env                 Variables de entorno (no incluido en el repositorio)
  go.mod               Modulo y dependencias del proyecto
  go.sum               Hashes de integridad de dependencias
```

La correspondencia con el patron MVC es la siguiente: los modelos son las entidades en internal/models, las vistas son las paginas HTML en templates, y los controladores son los handlers en internal/handlers.

---

## Principios de POO Aplicados

**Encapsulacion**

El modelo Producto encapsula la lógica de verificación de stock en el método VerificarStock. El modelo Carrito encapsula el cálculo del total en CalcularTotal. Los repositorios tienen el campo db en minúscula, lo que lo hace privado al paquete.

**Herencia por Composicion**

Go no tiene herencia clasica. Se utiliza struct embedding. La estructura Persona contiene los atributos comunes (Nombre, Email, Telefono) y es embebida por Cliente y Administrador, quienes heredan esos campos automaticamente.

**Polimorfismo**

La interface MetodoPago define el contrato con el metodo Procesar. Las estructuras PagoTarjeta y PagoTransferencia implementan esa interface de forma independiente. El servicio de pedidos trabaja con la interface sin conocer el tipo concreto.

**Abstraccion**

El patron Repository define interfaces como IProductoRepository que especifican que operaciones existen sin revelar como se implementan. Los servicios trabajan con las interfaces, no con las implementaciones concretas.

---

## Requisitos Previos

- Go 1.23 
- Git instalado

No se requiere instalar ninguna base de datos. SQLite esta embebido en el proyecto.

---

## Instalación y Ejecución

**Clonar el repositorio**

```bash
git clone https://github.com/DiegoM739/ReSport.git
cd ReSport
```

**Instalar dependencias**

```bash
go mod tidy
```

**Crear el archivo de variables de entorno**

Crear un archivo llamado .env en la raiz del proyecto con el siguiente contenido:

```
PORT=8081
DB_PATH=resport.db
JWT_SECRET=diegomedina95
JWT_EXPIRATION_HOURS=24
ENV=development
```

**Crear el administrador inicial**

```bash
go run cmd/server/seed/main.go
```

Este comando crea el administrador con las credenciales:

- Email: admin@resport.com
- Password: admin123
- Rol: superadmin

**Iniciar el servidor**

```bash
go run cmd/server/main.go
```

El servidor inicia en http://localhost:8081

---

## Credenciales de Prueba

**Administrador**

- Email: admin@resport.com
- Password: admin123

**Cliente de prueba**

- Email: # ReSport - Sistema de Comercio Electronico

**Autor:** Diego Medina  
**Materia:** Programacion Orientada a Objetos  
**Universidad:** Universidad Internacional del Ecuador, Sede Ambato  
**Semestre:** Tercero  
**Fecha:** Junio 2026  
**Repositorio:** https://github.com/DiegoM739/ReSport

---

## Descripcion General

ReSport es un sistema de comercio electronico orientado a la venta de productos deportivos y de rehabilitacion. El sistema fue desarrollado en el lenguaje de programacion Go aplicando los cuatro principios fundamentales de la Programacion Orientada a Objetos: encapsulacion, herencia por composicion, polimorfismo mediante interfaces y abstraccion mediante el patron Repository.

El sistema permite gestionar un catalogo de productos, registrar clientes con autenticacion segura, administrar un carrito de compras y procesar pedidos con metodos de pago simulados. Cuenta con un panel administrativo completo accesible desde el navegador.

---

## Objetivo del Proyecto

Desarrollar un sistema de gestion de comercio electronico que implemente los principios de Programacion Orientada a Objetos, integrando gestion de productos, autenticacion de usuarios con JWT, carrito de compras y procesamiento de pedidos, mediante el lenguaje Go con arquitectura limpia por capas.

---

## Tecnologias Utilizadas

**Backend**

- Go 1.23
- Gin v1.10 (framework web HTTP)
- GORM v1.25 (ORM para Go)
- glebarez/sqlite (driver SQLite pure-Go)
- golang-jwt/jwt v5 (autenticacion con JSON Web Tokens)
- golang.org/x/crypto/bcrypt (hash de contrasenas)
- joho/godotenv (variables de entorno)

**Frontend**

- HTML5
- Tailwind CSS via CDN
- JavaScript (fetch API)

**Herramientas**

- Git y GitHub (control de versiones)
- Visual Studio Code (IDE)
- DB Browser for SQLite (inspeccion de base de datos)

---

## Arquitectura del Sistema

El proyecto sigue una arquitectura limpia por capas donde cada capa tiene una responsabilidad especifica y solo se comunica con la capa inmediatamente inferior. Esta organizacion permite que cambiar una capa no afecte a las demas.

```
ReSport/
  cmd/
    server/
      main.go          Punto de entrada del servidor HTTP en puerto 8081
      seed/
        main.go        Script para crear el administrador inicial
  internal/
    config/            Carga de variables de entorno desde .env
    database/          Conexion y migracion automatica de SQLite
    models/            Entidades del dominio con principios de POO
    repository/        Acceso a datos mediante interfaces (abstraccion)
    services/          Logica de negocio del sistema
    handlers/          Controladores HTTP que reciben y responden peticiones
    middleware/        Validacion de tokens JWT para rutas protegidas
  templates/           7 paginas HTML con Tailwind CSS
  static/              Archivos estaticos
  .env                 Variables de entorno (no incluido en el repositorio)
  go.mod               Modulo y dependencias del proyecto
  go.sum               Hashes de integridad de dependencias
```

La correspondencia con el patron MVC es la siguiente: los modelos son las entidades en internal/models, las vistas son las paginas HTML en templates, y los controladores son los handlers en internal/handlers.

---

## Principios de POO Aplicados

**Encapsulacion**

El modelo Producto encapsula la logica de verificacion de stock en el metodo VerificarStock. El modelo Carrito encapsula el calculo del total en CalcularTotal. Los repositorios tienen el campo db en minuscula, lo que lo hace privado al paquete.

**Herencia por Composicion**

Go no tiene herencia clasica. Se utiliza struct embedding. La estructura Persona contiene los atributos comunes (Nombre, Email, Telefono) y es embebida por Cliente y Administrador, quienes heredan esos campos automaticamente.

**Polimorfismo**

La interface MetodoPago define el contrato con el metodo Procesar. Las estructuras PagoTarjeta y PagoTransferencia implementan esa interface de forma independiente. El servicio de pedidos trabaja con la interface sin conocer el tipo concreto.

**Abstraccion**

El patron Repository define interfaces como IProductoRepository que especifican que operaciones existen sin revelar como se implementan. Los servicios trabajan con las interfaces, no con las implementaciones concretas.

---

## Requisitos Previos

- Go 1.23 o superior instalado
- Git instalado

No se requiere instalar ninguna base de datos. SQLite esta embebido en el proyecto.

---

## Instalacion y Ejecucion

**Clonar el repositorio**

```bash
git clone https://github.com/DiegoM739/ReSport.git
cd ReSport
```

**Instalar dependencias**

```bash
go mod tidy
```

**Crear el archivo de variables de entorno**

Crear un archivo llamado .env en la raiz del proyecto con el siguiente contenido:

```
PORT=8081
DB_PATH=resport.db
JWT_SECRET=resport_secret_key_2026
JWT_EXPIRATION_HOURS=24
ENV=development
```

**Crear el administrador inicial**

```bash
go run cmd/server/seed/main.go
```

Este comando crea el administrador con las credenciales:

- Email: admin@resport.com
- Password: admin123
- Rol: superadmin

**Iniciar el servidor**

```bash
go run cmd/server/main.go
```

El servidor inicia en http://localhost:8081

---

## Credenciales de Prueba

**Administrador**

- Email: admin@resport.com
- Password: admin123

**Cliente de prueba**

- Email: diegomedina95@gmail.com
- Password: diegomedina95
- O crear un nuevo perfil

---

## Endpoints de la API REST

El sistema expone 15 endpoints REST con serializacion JSON. Los endpoints protegidos requieren el header Authorization con el formato: Bearer token.

**Endpoints publicos**

| Metodo | Ruta | Descripcion |
|--------|------|-------------|
| GET | /productos | Lista todos los productos del catalogo |
| GET | /productos/:id | Obtiene un producto por ID |
| POST | /productos | Crea un nuevo producto |
| PUT | /productos/:id | Actualiza un producto existente |
| DELETE | /productos/:id | Elimina un producto |
| POST | /auth/registro | Registra un nuevo cliente |
| POST | /auth/login | Autentica un cliente y devuelve JWT |
| POST | /admin/login | Autentica un administrador y devuelve JWT |
| GET | /health | Verifica que el servidor esta activo |

**Endpoints protegidos por token de cliente**

| Metodo | Ruta | Descripcion |
|--------|------|-------------|
| GET | /clientes/perfil | Obtiene el perfil del cliente autenticado |
| PUT | /clientes/perfil | Actualiza datos del perfil |
| GET | /carrito | Obtiene el carrito con items y total |
| POST | /carrito/items | Agrega un producto al carrito |
| PUT | /carrito/items/:id | Modifica la cantidad de un item |
| DELETE | /carrito/items/:id | Elimina un item del carrito |
| DELETE | /carrito | Vacia todo el carrito |
| POST | /pedidos | Confirma el carrito como pedido y procesa el pago |
| GET | /pedidos | Lista todos los pedidos del cliente |
| GET | /pedidos/:id | Obtiene un pedido especifico |

**Endpoints protegidos por token de administrador**

| Metodo | Ruta | Descripcion |
|--------|------|-------------|
| GET | /admin/dashboard | Accede al dashboard de administracion |

**Vistas HTML**

| Ruta | Pagina |
|------|--------|
| / | Pagina de inicio |
| /catalogo | Catalogo de productos |
| /login | Inicio de sesion de cliente |
| /registro | Registro de nuevo cliente |
| /admin | Panel de administrador |
| /carrito-view | Carrito de compras |
| /pedidos-view | Historial de pedidos |

---

## Funcionalidades del Sistema

**Gestion de Productos**

Operaciones CRUD completas sobre el catalogo. Control de stock con validacion antes de cada compra. Soporte para productos fisicos y digitales.

**Gestion de Clientes**

Registro con validacion de datos y encriptacion de contrasena con bcrypt. Autenticacion mediante tokens JWT con expiracion de 24 horas. Gestion de perfil del cliente.

**Carrito de Compras**

Carrito independiente por cliente. Agregar, modificar cantidad y eliminar productos. Calculo automatico del total. Vaciado completo del carrito.

**Procesamiento de Pedidos**

Conversion del carrito en pedido formal. Procesamiento de pago con tarjeta o transferencia bancaria. Transaccion atomica que garantiza consistencia: si algo falla, se revierte todo. Reduccion automatica del stock al confirmar el pedido.

**Panel Administrativo**

Acceso restringido por JWT con rol de administrador. Listado de productos con stock destacado en rojo cuando es bajo. Edicion de nombre, descripcion, precio y stock desde el navegador. Creacion de nuevos productos. Eliminacion con confirmacion.

---

## Base de Datos

El sistema utiliza SQLite como base de datos relacional embebida. GORM gestiona la migracion automatica de tablas al iniciar el servidor. Las tablas del sistema son: administradors, clientes, productos, carritos, items de carrito, pedidos, items de pedido y direccions.

---

## Visualizacion del Futuro

La arquitectura limpia implementada en ReSport esta disenada para escalar. Las evoluciones naturales del sistema incluyen la integracion con pasarelas de pago reales como Stripe o PayPal, el despliegue en plataformas en la nube con PostgreSQL como base de datos, la construccion de una aplicacion movil con React Native que consuma la misma API REST, y la implementacion de busqueda avanzada de productos con Elasticsearch.

Las tecnologias utilizadas en este proyecto, Go y su ecosistema de librerias, son adoptadas en produccion por empresas como Google, Uber y Dropbox. El conocimiento de Clean Architecture y los patrones Repository y Service aplicados aqui son transferibles a cualquier lenguaje y contexto profesional.

---

## Estructura del Proyecto Academico

Este proyecto integra los contenidos de las cuatro unidades de la materia de Programacion Orientada a Objetos:

- Unidad 1: Clases, objetos, atributos y metodos aplicados en los modelos del dominio
- Unidad 2: Herencia por composicion mediante struct embedding en Go
- Unidad 3: Interfaces, polimorfismo, encapsulacion y manejo de errores
- Unidad 4: Servicios web REST, serializacion JSON y arquitectura por capas

---

## Licencia

Proyecto academico desarrollado para la materia de Programacion Orientada a Objetos en la Universidad Internacional del Ecuador. Junio 2026.
- Password: password123

---

## Endpoints de la API REST

El sistema expone 15 endpoints REST con serializacion JSON. Los endpoints protegidos requieren el header Authorization con el formato: Bearer token.

**Endpoints publicos**

| Metodo | Ruta | Descripcion |
|--------|------|-------------|
| GET | /productos | Lista todos los productos del catalogo |
| GET | /productos/:id | Obtiene un producto por ID |
| POST | /productos | Crea un nuevo producto |
| PUT | /productos/:id | Actualiza un producto existente |
| DELETE | /productos/:id | Elimina un producto |
| POST | /auth/registro | Registra un nuevo cliente |
| POST | /auth/login | Autentica un cliente y devuelve JWT |
| POST | /admin/login | Autentica un administrador y devuelve JWT |
| GET | /health | Verifica que el servidor esta activo |

**Endpoints protegidos por token de cliente**

| Metodo | Ruta | Descripcion |
|--------|------|-------------|
| GET | /clientes/perfil | Obtiene el perfil del cliente autenticado |
| PUT | /clientes/perfil | Actualiza datos del perfil |
| GET | /carrito | Obtiene el carrito con items y total |
| POST | /carrito/items | Agrega un producto al carrito |
| PUT | /carrito/items/:id | Modifica la cantidad de un item |
| DELETE | /carrito/items/:id | Elimina un item del carrito |
| DELETE | /carrito | Vacia todo el carrito |
| POST | /pedidos | Confirma el carrito como pedido y procesa el pago |
| GET | /pedidos | Lista todos los pedidos del cliente |
| GET | /pedidos/:id | Obtiene un pedido especifico |

**Endpoints protegidos por token de administrador**

| Metodo | Ruta | Descripcion |
|--------|------|-------------|
| GET | /admin/dashboard | Accede al dashboard de administracion |

**Vistas HTML**

| Ruta | Pagina |
|------|--------|
| / | Pagina de inicio |
| /catalogo | Catalogo de productos |
| /login | Inicio de sesion de cliente |
| /registro | Registro de nuevo cliente |
| /admin | Panel de administrador |
| /carrito-view | Carrito de compras |
| /pedidos-view | Historial de pedidos |

---

## Funcionalidades del Sistema

**Gestion de Productos**

Operaciones CRUD completas sobre el catalogo. Control de stock con validacion antes de cada compra. Soporte para productos fisicos y digitales.

**Gestion de Clientes**

Registro con validacion de datos y encriptacion de contrasena con bcrypt. Autenticacion mediante tokens JWT con expiracion de 24 horas. Gestion de perfil del cliente.

**Carrito de Compras**

Carrito independiente por cliente. Agregar, modificar cantidad y eliminar productos. Calculo automatico del total. Vaciado completo del carrito.

**Procesamiento de Pedidos**

Conversion del carrito en pedido formal. Procesamiento de pago con tarjeta o transferencia bancaria. Transaccion atomica que garantiza consistencia: si algo falla, se revierte todo. Reduccion automatica del stock al confirmar el pedido.

**Panel Administrativo**

Acceso restringido por JWT con rol de administrador. Listado de productos con stock destacado en rojo cuando es bajo. Edicion de nombre, descripcion, precio y stock desde el navegador. Creacion de nuevos productos. Eliminacion con confirmacion.

---

## Base de Datos

El sistema utiliza SQLite como base de datos relacional embebida. GORM gestiona la migracion automatica de tablas al iniciar el servidor. Las tablas del sistema son: administradors, clientes, productos, carritos, items de carrito, pedidos, items de pedido y direccions.

---

## Visualizacion del Futuro

La arquitectura limpia implementada en ReSport esta disenada para escalar. Las evoluciones naturales del sistema incluyen la integracion con pasarelas de pago reales como Stripe o PayPal, el despliegue en plataformas en la nube con PostgreSQL como base de datos, la construccion de una aplicacion movil con React Native que consuma la misma API REST, y la implementacion de busqueda avanzada de productos con Elasticsearch.

Las tecnologias utilizadas en este proyecto, Go y su ecosistema de librerias, son adoptadas en produccion por empresas como Google, Uber y Dropbox. El conocimiento de Clean Architecture y los patrones Repository y Service aplicados aqui son transferibles a cualquier lenguaje y contexto profesional.

---

## Estructura del Proyecto Academico

Este proyecto integra los contenidos de las unidades de la materia de Programacion Orientada a Objetos:

- Clases, objetos, atributos y metodos aplicados en los modelos del dominio
- Herencia por composicion mediante struct embedding en Go
- Interfaces, polimorfismo, encapsulacion y manejo de errores
- Servicios web REST, serializacion JSON y arquitectura por capas
