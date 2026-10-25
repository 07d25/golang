# 14. Seed: datos para aprender sin escribirlos a mano

Un **seed** es un programa que carga datos conocidos para poder probar una
aplicación. En este proyecto crea solamente tres usuarios:

| Usuario | Contraseña |
| --- | --- |
| `pixel_ana` | `pixel-ana-123` |
| `ninja_luis` | `ninja-luis-123` |
| `retro_mia` | `retro-mia-123` |

## Ejecutarlo

Con PostgreSQL levantado y `.env` configurado:

```bash
go run ./cmd/seed
```

El comando también ejecuta las migraciones antes de insertar. Si lo ejecutas
dos veces, no aparecen usuarios duplicados: `username` es único y el programa
usa `ON CONFLICT` para actualizar los mismos tres perfiles.

El seed no crea posts, salas ni mensajes. Así puedes avanzar en el tutorial
viendo qué datos agrega cada acción de la interfaz.
