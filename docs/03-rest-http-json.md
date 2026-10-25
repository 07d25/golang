# 03. REST, HTTP y JSON

## Qué es una API REST

Una API es un contrato entre programas. REST organiza ese contrato alrededor de recursos: usuarios, publicaciones, salas y mensajes.

| Acción | Método | Ruta |
|---|---|---|
| Crear usuario | `POST` | `/users` |
| Listar usuarios | `GET` | `/users` |
| Leer un usuario | `GET` | `/users/{id}` |
| Crear una sala | `POST` | `/rooms` |
| Leer mensajes | `GET` | `/rooms/{id}/messages` |

## Una petición por dentro

```text
Método + URL + headers + body
```

Ejemplo:

```http
POST /users
Content-Type: application/json
```

```json
{"username":"pixel_ana","display_name":"Ana","password":"ana-password"}
```

El handler usa `json.Decoder` para transformar el body en un struct Go. Luego valida, guarda y responde con `json.Encoder`.

## Status codes

```text
200 OK              lectura o actualización correcta
201 Created         recurso creado
400 Bad Request     JSON o datos inválidos
401 Unauthorized    token ausente o incorrecto
403 Forbidden       usuario sin permiso
404 Not Found       recurso inexistente
409 Conflict        duplicado
429 Too Many        demasiadas solicitudes
```

## Laboratorio

```bash
curl -i -X POST http://localhost:3000/users \
  -H 'Content-Type: application/json' \
  -d '{"username":"pixel_ana","display_name":"Ana","password":"ana-password"}'

curl -i http://localhost:3000/users/1
curl -i http://localhost:3000/users
```

🎯 **Reto:** cambia el ID a `999999` y explica por qué la API responde `404`.
