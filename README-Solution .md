<img src="https://getduna.com/svg/duna-logo.svg" width="300">

## Solución

La solución que se presenta tiene los requitos solicitados

1. Microservicio svc-aut
2. Microservicio svc-driver
3. Una base de datos relacional postgres
4. Una base de datos No relacional mongodb
5. Test para servicice del proyecto svc-driver

## Tecnologías utilizadas para implementar la solución.

* Docker
* Docker-compose
* Lenguaje: Golang
* Librerias: Fs02, testify entre otras

## Requisitos  del proyecto 
1. Contar con docker instalado, si no lo tiene, lo puede obtener [aqui](https://docs.docker.com/engine/install/).
2. Disponer de docker-compose, se puede instalar de acuerdo a lo indicado en el siguiente [link](https://docs.docker.com/compose/install/) 
3. Será más cómoda la instalación si dispone del comando Make

## Instalación  del proyecto 
1. **Descargar código fuente**

```console
git clone https://github.com/d-Una-Interviews/go-senior-challenge-lucianoBaez.git
```

2. **Compilar**

```console
 cd go-senior-challenge-lucianoBaez
 go build
```

3. **Desplegar**

Se puede desplegar utilizando comandos make:

```console
make start-solution
```
ó bien utilizando docker-compose

```console
cd svc_aut
docker build . -t svc_aut
cd ..

cd svc_aut 
docker build . -t svc_driver
cd ..

cd docker
docker-compose up -d;
```

Luego de la instalación:

```console
[luciano@localhost go-senior-challenge-lucianoBaez]$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                      NAMES
d26d46d82eaf        svc_aut             "/go/bin/svc_aut"        10 seconds ago      Up 9 seconds        0.0.0.0:8081->8081/tcp     docker_svc_auth_1
e9cc6a2f2ac5        svc_driver          "/go/bin/svc_driver"     10 seconds ago      Up 9 seconds        0.0.0.0:8082->8081/tcp     docker_svc_driver_1
04acbb808da8        postgres            "docker-entrypoint.s…"   12 seconds ago      Up 10 seconds       0.0.0.0:5432->5432/tcp     docker_db_1
7a53ca2ccee7        mongo:latest        "docker-entrypoint.s…"   12 seconds ago      Up 10 seconds       0.0.0.0:27017->27017/tcp   mongodb
```


Cuando levanta el proyecto, se crea un user, y se crean 100 drivers para poder hacer la paginación

- Información del usuario: admin/passowrd

Se crean dos bases de datos, cuyos strings de conexón quedan en:

- PostgreSql: jdbc:postgresql://localhost:5432/drivers_db, user/password
- MongoDB: mongodb://localhost:27017/drivers_mongo, admin-user/admin-password




4. **Urls**
- svc_aut   : http://localhost:8081/api/v1     
- svc_driver: http://localhost:8082/api/v1


5. **Invocación a servicios**

* Autenticación

```console
curl --location --request POST 'http://localhost:8081/api/v1/user/authenticate' \
--header 'Content-Type: application/json' \
--data-raw '{"username":"admin",
 "password":"password"}'
```


* Acceder a todos los drivers paginados

```console
curl --location --request GET 'http://localhost:8081/api/v1/drivers/limit/10/page/11' \
--header 'x-access-token: {token}
```


* Crear un conductor

```console
curl --location --request POST 'http://localhost:8081/api/v1/drivers' \
--header 'x-access-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjE4MTA2ODI4fQ.7IEurxb58eB-3axb6jb25Ib6ToYoUR8K3fSxB-NUkLw' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Name": "Mariano",
    "LastName": "Baez",
    "Email": "mariano@gmail.com",
    "Location": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z"
}'
```

* Obtener conductores a un radio X 

```console
curl --location --request GET 'http://localhost:8081/api/v1/drivers/radius/2' \
--header 'x-access-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjE4MTA2ODI4fQ.7IEurxb58eB-3axb6jb25Ib6ToYoUR8K3fSxB-NUkLw'
```

Puede importar la colección y el environment de postman incluidos en el proyecto en su cliente postman, para realizar las pruebas de manera más sencilla.

## Tests

**Ejecución de tests:** 
```console
make run-tests-svc-driver

```

ó 

```console
go test ./... -v
```

## Aclaraciones
Si bien no había tiempo establecido para la entrega, solamente puedo realizar el examen en fin de semana, intenté hacer de todo un poco de modo de que poder dejar una muestra de lo que sé. Quedo atento a cualquier consulta, modificación o comentario, gracias por la posibildad de participar en el proceso de selección



