# Laboratorio completo: de usuario a chat

Este capítulo conecta todo. Ejecuta los pasos en orden y copia los IDs de las respuestas.

Este orden es importante: no puedes enviar un mensaje antes de que exista la sala y de que el usuario sea miembro.

## 1. Arrancar

```bash
docker compose up -d
go run ./cmd/api
```

## 2. Crear a Ana

```bash
curl -s -X POST http://localhost:3000/users -H 'Content-Type: application/json' -d '{"username":"pixel_ana","display_name":"Ana","password":"ana-password"}'
```

Guarda el `id` como `ANA_ID`.

## 3. Crear a Luis

```bash
curl -s -X POST http://localhost:3000/users -H 'Content-Type: application/json' -d '{"username":"pixel_luis","display_name":"Luis","password":"luis-password"}'
```

Guarda el `id` como `LUIS_ID`. En los ejemplos suponemos que son `1` y `2`.

## 4. Crear una publicación

```bash
curl -s -X POST http://localhost:3000/posts -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' -d '{"content":"Mi primera partida"}'
```

El header identifica quién publica. En una aplicación real lo resolvería una sesión o un token, no el cliente.

## 5. Crear una sala

```bash
curl -s -X POST http://localhost:3000/rooms -H 'Content-Type: application/json' -d '{"name":"RPG nocturno"}'
```

Guarda el ID de la sala. En los siguientes comandos usamos `1`.

## 6. Unir usuarios

```bash
curl -i -X POST http://localhost:3000/rooms/1/join -H "Authorization: Bearer $TOKEN"
curl http://localhost:3000/rooms/1/members
```

Ahora ambos tienen permiso para abrir el WebSocket.

## 7. Abrir dos conexiones

En dos terminales ejecuta:

```bash
websocat -H="Authorization: Bearer $TOKEN" ws://localhost:3000/ws/rooms/1
```

Escribe en la primera terminal:

```json
{"content":"¿Jugamos?"}
```

El servidor fuerza `room_id` y `user_id`, guarda el mensaje y lo difunde a los clientes de la sala.

## 8. Consultar historial

```bash
curl http://localhost:3000/rooms/1/messages
```

El mensaje sigue ahí aunque cierres las conexiones porque se guardó en PostgreSQL antes del broadcast.

## Qué aprendiste

```text
REST crea y consulta recursos.
PostgreSQL conserva el estado.
WebSocket mantiene una conversación abierta.
El hub coordina clientes.
Goroutines leen y escriben en paralelo.
Los canales transportan eventos.
Context cancela trabajo al desconectar.
```

## Reto final

Conecta un tercer cliente sin hacer join. Observa que el servidor rechaza la conexión. Después implementa un evento `user_joined`.
