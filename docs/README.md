# Curso: construye una comunidad gamer

La meta es aprender Go siguiendo una aplicación real desde cero. Ejecuta cada laboratorio, sigue la petición por el código y resuelve el reto.

## Ruta de aprendizaje

El backend se aprende en la rama `02-gamer-chat-api`. El frontend se aprende en
`03-gamer-chat-frontend`, partiendo de esta API ya terminada.

0. [Antes de comenzar](00-ruta.md)
1. [Mapa de conceptos](00-conceptos-completos.md)
2. [Conceptos: API REST, HTTP y WebSocket](01-conceptos.md)
3. [Primeros pasos](02-primeros-pasos.md)
4. [Arquitectura](03-arquitectura.md)
5. [REST, HTTP y JSON](03-rest-http-json.md)
6. [Usuarios y perfiles](04-usuarios.md)
7. [Autenticación](05-autenticacion.md)
8. [Red social](05-red-social.md)
9. [Salas y mensajes](06-salas.md)
10. [WebSocket, hub, goroutines y canales](07-websocket.md)
11. [Testing](08-testing.md)
12. [Laboratorio completo](09-flujo-completo.md)
13. [Migraciones](10-migraciones.md)
14. [Seguridad y límites](11-seguridad-y-limites.md)
15. [Frontend: HTMX, Alpine.js y Tailwind](13-frontend.md)
16. [Seed de usuarios de prueba](14-seed.md)

## Método

En cada capítulo: lee la misión, ejecuta el comando, predice el resultado, sigue el código, resuelve el reto y consulta la solución solo al final.

## Orden recomendado

No empieces por WebSocket. Primero construye y entiende una petición HTTP normal; después verás que WebSocket reutiliza muchas ideas conocidas y añade una conexión persistente.

```text
HTTP y JSON → PostgreSQL → relaciones → historial → WebSocket → concurrencia
```

## Convenciones

- El código usa nombres técnicos en inglés.
- La explicación y los comentarios están en español.
- Las rutas protegidas usan JWT; `X-User-ID` se utiliza internamente después de validar el token.
- Los comandos se ejecutan desde la raíz del repositorio.
