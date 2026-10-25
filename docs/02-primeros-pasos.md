# 02. Primeros pasos

## Qué vas a aprender

Vas a arrancar PostgreSQL, iniciar Go y comprobar el primer endpoint sin modificar todavía la lógica de usuarios.

## Misión

Levantar la API y comprobar que responde.

```bash
docker compose up -d
cp .env.example .env
go run ./cmd/api
```

El archivo `.env` contiene la configuración local y no debe subirse al repositorio. El proyecto incluye `.env.example` como plantilla.

En otra terminal:

```bash
curl -i http://localhost:3000/health
```

Busca `r.Get("/health"` en `cmd/api/main.go` y sigue cómo se construye la respuesta JSON.

## Qué significa cada comando

```bash
docker compose up -d
```

Inicia el contenedor de PostgreSQL en segundo plano. La aplicación necesita que la base esté disponible antes de llamar a `database.Open`.

```bash
go run ./cmd/api
```

Compila y ejecuta el package `main` ubicado en `cmd/api`. Al iniciar, el programa abre la conexión, ejecuta la migración inicial y comienza a escuchar en `:3000`.

```bash
curl -i http://localhost:3000/health
```

`curl` actúa como cliente HTTP. `-i` muestra también los headers y el código de estado.

## Lee la respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{"status":200,"data":{"status":"ok"}}
```

El código `200` significa que la operación terminó correctamente. `Content-Type` explica que el cuerpo es JSON.

## Si algo falla

| Síntoma | Causa probable | Comprobación |
|---|---|---|
| Error de conexión | PostgreSQL no está iniciado | `docker compose ps` |
| Puerto ocupado | Otro proceso usa 3000 | Cambia `PORT` |
| Error de módulo | Dependencias no descargadas | `go mod tidy` |

## Solución del reto

La ruta `/health` llama a `write`, que envuelve cualquier dato dentro de `status` y `data`. Para añadir el nombre de la aplicación cambiarías el mapa enviado a `write`.

🎯 **Reto:** explica qué cambiarías para incluir el nombre de la aplicación en la respuesta.
