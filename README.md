# Wankar Service

## Para instalar las dependencias del proyecto: 

>`go mod download`

## Para crear conexion a la Base de Datos:

Puede crear un tunel ssh a la base de datos remota o hacer una copia local de esta. Hacer una copia local es recomendable para no perder o distorsionar datos en la base de datos de produccion a causa de algun error.

ssh tunnel:
https://www.ssh.com/academy/ssh/tunneling-example

Hacer una copia local:
https://www.postgresql.org/docs/8.1/backup.html

## .env:
crear archivo `.env` en el directorio raiz del proyecto.

En este archivo las variables:

PORT = 8080

DB_HOST = localhost

DB_PORT = 5432

DB_USER = postgres

DB_PASSWORD = ???

DB_NAME = ???

## Para ejecutar el proyecto:

>`go run test/test.go`

## Para usar GraphiQL:

localhost:8080

### query ejemplo:
`{
  allCategorias {
    edges {
      node {
        id
        descripcion
        tipoCategoria {
          id
          nombre
        }
        categoriaPadre {
          id
          descripcion
          categoriaPadre {
            id
            descripcion
          }
        }
      }
    }
  }
}
`