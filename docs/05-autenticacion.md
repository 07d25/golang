# 07. Autenticación y autorización

## Dos preguntas diferentes

- **Autenticación:** ¿quién eres?
- **Autorización:** ¿puedes hacer esta acción?

El login responde la primera. La membresía de una sala responde la segunda.

## Crear cuenta

```bash
curl -i -X POST http://localhost:3000/users \
  -H 'Content-Type: application/json' \
  -d '{"username":"pixel_ana","display_name":"Ana","password":"ana-password"}'
```

La contraseña se transforma con bcrypt antes de guardarse. La API nunca la devuelve.

## Login y JWT

```bash
curl -s -X POST http://localhost:3000/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"pixel_ana","password":"ana-password"}'
```

La respuesta contiene un token firmado. El cliente lo envía así:

```http
Authorization: Bearer eyJ...
```

El middleware valida la firma, lee el `sub` y coloca el ID en el contexto. Los handlers pueden usarlo sin confiar en un ID enviado dentro del body.

## Flujo de autorización de sala

```text
login
→ obtener JWT
→ POST /rooms/1/join
→ abrir WebSocket con identidad
→ enviar mensajes
```

🎯 **Reto:** intenta iniciar sesión con una contraseña incorrecta y observa el `401`.

## Forma de los errores

```json
{"status":400,"error":"password debe tener al menos 8 caracteres"}
```

El cliente siempre encuentra el mensaje en `error`.
