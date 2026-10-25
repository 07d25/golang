# 10. Migraciones y evolución de la base

## Por qué separar el SQL

La conexión debe encargarse de abrir PostgreSQL y las migraciones deben describir cómo cambia el esquema. Por eso el proyecto guarda el SQL en `internal/database/migrations`.

```text
internal/database/
  database.go
  migrations/
    001_init.sql
```

`database.Migrate` encuentra los archivos `.sql`, los ordena por nombre y los ejecuta. El prefijo numérico indica el orden.

## Cómo agregar una tabla

1. Crea `002_notifications.sql`.
2. Escribe solo el cambio nuevo.
3. Ejecuta la aplicación contra una base local.
4. Comprueba que la tabla exista.
5. No modifiques una migración que ya se haya ejecutado en otros entornos.

## Reto

Diseña una migración para añadir `last_seen_at` a `users`. Antes de escribirla, explica qué ocurre con usuarios que ya existen.
