# 08. Testing y próximos pasos

## Qué vas a aprender

Un programa que compila no necesariamente funciona. Los tests convierten los ejemplos del curso en comprobaciones repetibles.

```bash
go test ./...
go vet ./...
go test -race ./...
```

El proyecto ya incluye pruebas unitarias para validaciones, una prueba de broadcast del hub y una prueba de integración contra PostgreSQL que se activa cuando existe `DATABASE_URL`. Todavía falta una prueba WebSocket usando dos conexiones de red reales.

## Tipos de test

| Tipo | Comprueba | Herramienta |
|---|---|---|
| Unitario | Validaciones puras | `testing` |
| HTTP | Status y JSON | `httptest` |
| Integración | SQL real | PostgreSQL |
| Concurrencia | Hub y conexiones | `-race` |

## Primer test mental

Para `CreateUser`, prepara un request válido y otro sin username. El primero debe devolver un usuario; el segundo debe devolver error sin ejecutar el INSERT.

## Estado honesto del proyecto

El MVP compila, valida membresías, persiste eventos WebSocket y tiene pruebas unitarias básicas. Para ejecutar la integración usa PostgreSQL y `DATABASE_URL`; para producción todavía necesita autenticación real y migraciones versionadas.

🎯 **Reto final:** implementa `GET /rooms/{id}/messages` con paginación.
