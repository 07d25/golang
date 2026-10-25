# 13. Frontend: HTMX, Alpine.js y Tailwind

## La misión

Hasta aquí usábamos `curl` para hablar con la API. Ahora construiremos una
interfaz real sin esconder la arquitectura detrás de un framework enorme.

```text
navegador → HTML + Alpine → fetch/HTMX → API Go → PostgreSQL
                                      ↘ WebSocket → hub → otros navegadores
```

## Qué aporta cada herramienta

- **HTML** describe la estructura que la persona ve.
- **Tailwind CSS** aplica estilos usando clases pequeñas y explícitas.
- **Alpine.js** mantiene estado local: sala activa, formulario y errores.
- **HTMX** pide un fragmento HTML al servidor y reemplaza solo una parte de la
  pantalla.
- **WebSocket** mantiene una conexión abierta para recibir mensajes al instante.

No estamos creando una SPA. La pantalla sigue siendo HTML servido por Go y
solo usamos JavaScript donde aporta una interacción concreta.

## Primer recorrido

1. Abre `/register` y crea una cuenta.
2. El formulario llama a `POST /users` y luego a `POST /auth/login`.
3. Go devuelve el JWT y lo guarda en una cookie `HttpOnly`; el navegador la
   envía también al WebSocket.
4. `/app` carga publicaciones y salas con `fetch`.
5. Al publicar, Alpine envía JSON y vuelve a consultar el feed.
6. En una sala, el navegador abre `/ws/rooms/{id}` y escucha eventos.

## HTMX en acción

En la tarjeta “Aprende mientras usas”, el botón tiene:

```html
<button hx-get="/ui/lesson"
        hx-target="#lesson"
        hx-swap="innerHTML">
  Ver cómo funciona
</button>
```

Al hacer clic, HTMX pide HTML a Go y lo inserta en `#lesson`. No hay que
construir manualmente un `div`, buscarlo y asignarle `innerHTML`.

## Alpine en acción

```html
<form @submit.prevent="createPost()">
  <textarea x-model="newPost"></textarea>
</form>
```

`x-model` conecta el campo con el estado y `@submit.prevent` evita la recarga
tradicional. La función `createPost` vive en `web/static/app.js`.

## Reto

Añade un contador de mensajes no leídos usando Alpine. Después cambia el
fragmento HTMX para explicar qué diferencia hay entre reemplazar `innerHTML` y
reemplazar un elemento completo.

## Ejecutar

```bash
docker compose up -d
cp .env.example .env
go run ./cmd/api
```

Visita <http://localhost:3000>. Tailwind, Alpine y HTMX se cargan desde CDN
para mantener el primer frontend pequeño y fácil de inspeccionar.
