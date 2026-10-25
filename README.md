# Comunidad Gamer en Go

Proyecto educativo para aprender Go construyendo una red social mínima con chat en tiempo real.

Incluye usuarios, publicaciones, follows, likes, comentarios, salas, mensajes, PostgreSQL y WebSocket con `github.com/coder/websocket`.

Las rutas que modifican contenido requieren `Authorization: Bearer <token>`. El flujo de aprendizaje empieza creando una cuenta, haciendo login y usando el JWT.

## Ejecutar

```bash
docker compose up -d
cp .env.example .env
go run ./cmd/api
```

En otra terminal:

```bash
curl http://localhost:3000/health
curl -X POST http://localhost:3000/users -H 'Content-Type: application/json' -d '{"username":"pixel_ana","display_name":"Ana","password":"ana-password"}'
```

La documentación completa está en [docs/README.md](docs/README.md).

```bash
go test ./...
go vet ./...
```
