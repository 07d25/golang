# 11. Seguridad y límites

## Rate limiting

La API incluye un límite educativo de 120 solicitudes por minuto y dirección IP. Cuando se supera, responde:

```http
429 Too Many Requests
Retry-After: 60
```

El middleware vive en `internal/ratelimit`. Guarda un contador por IP protegido con un mutex y reinicia la ventana después de un minuto.

## Qué protege y qué no protege

El límite reduce abusos accidentales y ráfagas de peticiones. No reemplaza autenticación, autorización ni un sistema distribuido como Redis.

`X-User-ID` sigue siendo una identidad didáctica. Cualquier cliente puede cambiarla, por lo que no sirve para producción.

## Reto

Modifica el límite en `cmd/api/main.go`, observa el test del middleware y explica por qué un contador en memoria no funcionaría igual con varias instancias de la API.
