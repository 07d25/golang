# 13. Frontend: DevSquads con HTMX, Alpine.js y Tailwind

## La misión

Hasta aquí usábamos `curl` para hablar con la API. Ahora convertimos la
comunidad en DevSquads: un sprint de cuatro semanas para que juniors tech
terminen un proyecto, compartan avances y se acerquen a una oportunidad real.

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

## El producto en una frase

> Construí algo que puedas mostrar, acompañado por un squad.

El producto no es el chat. El chat es una herramienta. El valor está en la
misión semanal, la racha, el check-in, el proyecto terminado y la oportunidad
curada.

## Recorrido del MVP

1. La persona entra y elige un objetivo; en esta primera versión usamos
   `Backend con Go`.
2. El dashboard muestra la misión de la semana, XP, racha y progreso.
3. El usuario hace un check-in de 30 minutos.
4. Comparte un avance en el squad y recibe ánimo con likes/comentarios.
5. Revisa oportunidades curadas que enlazan a la fuente original.
6. Termina con un demo day y un proyecto para su portfolio.

La misión y el check-in son estado local de Alpine en esta etapa didáctica.
Los posts, likes y comentarios ya viajan a la API y PostgreSQL.

## El primer agente: reglas antes que magia

El frontend pide la siguiente misión a:

```http
POST /api/coach/next
```

Enviamos el objetivo y la semana actual. La implementación inicial usa
`RulesCoach`: devuelve una misión predefinida, es barata y se puede probar.
La interfaz `Coach` permite reemplazarla por un agente de IA más adelante sin
cambiar el dashboard.

El check-in se revisa en:

```http
POST /api/check-ins
```

El agente no hace el proyecto por la persona. Devuelve una siguiente acción
pequeña y concreta. Esa es la diferencia entre un asistente útil y una caja de
chat que responde cualquier cosa.

No estamos creando una SPA. La pantalla sigue siendo HTML servido por Go y
solo usamos JavaScript donde aporta una interacción concreta.

## Primer recorrido técnico

1. Abre `/register` y crea una cuenta.
2. El formulario llama a `POST /users` y luego a `POST /auth/login`.
3. Go devuelve el JWT y lo guarda en una cookie `HttpOnly`; el navegador la
   envía también al WebSocket.
4. `/app` carga publicaciones con `fetch`.
5. Al publicar, Alpine envía JSON y actualiza el squad.
6. Las oportunidades son datos curados del MVP; el botón lleva a la fuente.
7. En una sala, el navegador abre `/ws/rooms/{id}` y escucha eventos.

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

## Qué todavía no está conectado a PostgreSQL

Para mantener el primer producto pequeño, estas piezas son una maqueta
interactiva: XP, racha, misión, check-in y oportunidades. El siguiente reto
es crear las tablas `study_sprints`, `check_ins` y `challenges`.

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
