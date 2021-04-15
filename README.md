<img src="https://getduna.com/svg/duna-logo.svg" width="300">

# Golang Senior Challenge

## 👩‍💻 Project Overview

El objetivo de este proyecto será evaluar las habilidades del candidato. Las condiciones del challenge son presentadas solo como guias las cuales un desarrollador experimentado debe escoger como resolverlos con la tecnologia, patron o diseño de su preferencia.

Recuerda que no tenemos un timeline, reclutamos on a rolling basis, pero estaremos moviendo a personas a la siguiente fase la primera semana de enero.

## 🦶 Pasos
- Crear un repo en github con el siguiente [link](https://classroom.github.com/a/6Jhgrl0w).
- Crear un proyecto de go usando modulos y usando tu libreria http de preferencia. (d-Una usa https://github.com/gin-gonic/gin)
- Configurar una red de servicios usando docker-compose que hara disponible los siguientes servicios:
    - svc_auth: autenticará y autorizará usuarios.
        - deberá de manejar la autenticacion de conductores
    - svc_driver: expondrá los siguientes endpoints
        - obtener todas los conductores (paginado)
        - crear un nuevo conductor (manejar las credenciales de autenticacion del mismo.)
        - obtener todos los conductores basado en disponibilidad en un radio de 2km.
        - cada ruta debe autenticarse con el servicio svc_auth via un middleware (usa el esquema de autenticacion de preferencia[jwt, basic auth, custom token, etc])
    - db_driver: Usar una base de datos no relacional (e.g mongodb)
    - db_auth: Usar una base de datos relacional (e.g postgres)
- Ambos servicios deberán de considerar los siguientes principios de diseño:
    - monitoreo de registros (logs) y usar adecuadamente el nivel de alerta de logs (DEBUG, INFO, etc)
    - pruebas unitarias de handlers, repositorios o servicios.
    - manejo de codigo de estado 5xx y panics
    - al menos un patron de diseño de estructura como (DDD, MVC, etc)
    - principios REST (codigos de estados y verbos correcto asi como convecion de rutas)
    - separar la capa de presentacion y datos (requerimiento minimo)
- Al finalizar todo sube tu código al repositorio y envíanos el link

## 🎯 Puntos de evaluacion

1. Estructura de codigo
2. Codigo que sea facil de mantener y hacer pruebas.
3. Diseño REST
4. Uso de custom middlewares
5. Interaccion con base de datos relacionales
6. Interaccion con base de datos no relacionales
7. Ciclo de vida de desarrollo desde implementacion hasta lanzamiento y monitoreo.

## 🎯 Bonificaciones

- Agrega un seccion de deployment en este README.md donde expliques como manejarias el monitoreo y  lanzamiento de los servicios en premisa o con un proveedor de servicios en la nube
- Usar mocks para incrementar la cobertura de pruebas.

## 📃 Disclaimer

Esto es un desafío técnico sin proposito comercial y D-Una de ninguna manera:

* Compartirá o usará tu código
* Obligarte a realizar este desafío
* Compensarte de cualquier forma por realizar este desafío

# Buena suerte 🚀
Esta prueba se puede completar fácilmente en un lapso de 6 a 8 horas
