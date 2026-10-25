# 04. Usuarios y perfiles

## Qué vas a aprender

Crearás una fila, leerás un listado y usarás un header temporal para representar al usuario actual.

```bash
curl -X POST http://localhost:3000/users \
  -H 'Content-Type: application/json' \
  -d '{"username":"pixel_ana","display_name":"Ana","password":"ana-password"}'
curl http://localhost:3000/users
```

## Request y respuesta

Request:

```json
{
  "username": "pixel_ana",
  "display_name": "Ana",
  "bio": "RPG"
}
```

Respuesta aproximada:

```json
{
  "status": 201,
  "data": {
    "id": 1,
    "username": "pixel_ana",
    "display_name": "Ana",
    "bio": "RPG"
  }
}
```

El `id` lo genera PostgreSQL. Después haz login y guarda el JWT: será la credencial de las siguientes operaciones.

```bash
curl -X POST http://localhost:3000/users/2/follow -H "Authorization: Bearer $TOKEN"
curl -X DELETE http://localhost:3000/users/2/follow -H "Authorization: Bearer $TOKEN"
```

🎯 **Reto:** intenta seguirte a ti mismo y observa dónde se rechaza la operación.

## Qué ocurre con un username repetido

La tabla tiene `username TEXT NOT NULL UNIQUE`. El primer usuario se crea; el segundo provoca un error de PostgreSQL. Más adelante aprenderemos a traducir ese error a `409 Conflict`.

## Checkpoint

<details><summary>¿Por qué `X-User-ID` no es autenticación?</summary>

Porque el cliente puede escribir cualquier número. Solo sirve como atajo pedagógico para identificar relaciones durante el curso.
</details>
