# Go Bases

Repositorio práctico en español para recorrer las lecciones 01–29 de Go. Cada ejemplo es independiente y se puede ejecutar con `go run ./cmd/<lección>`.

## Índice

| Lección | Tema | Enlaces |
|---|---|---|
| 01 | Variables | [código](cmd/01-variables) · [doc](doc/01-variables.md) |
| 02 | Conversiones de variables | [código](cmd/02-conversiones) · [doc](doc/02-conversiones.md) |
| 03 | Tipo `byte` | [código](cmd/03-byte) · [doc](doc/03-byte.md) |
| 04 | Operadores | [código](cmd/04-operadores) · [doc](doc/04-operadores.md) |
| 05 | Condicionales | [código](cmd/05-condicionales) · [doc](doc/05-condicionales.md) |
| 06 | Go Bases Tarea 1 | [código](cmd/06-tarea-1) · [doc](doc/06-tarea-1.md) |
| 07 | Vectores | [código](cmd/07-vectores) · [doc](doc/07-vectores.md) |
| 08 | Slices | [código](cmd/08-slices) · [doc](doc/08-slices.md) |
| 09 | Maps | [código](cmd/09-maps) · [doc](doc/09-maps.md) |
| 10 | `for` | [código](cmd/10-for) · [doc](doc/10-for.md) |
| 11 | Go Bases Tarea 2 | [código](cmd/11-tarea-2) · [doc](doc/11-tarea-2.md) |
| 12 | Go Bases Tarea 3 | [código](cmd/12-tarea-3) · [doc](doc/12-tarea-3.md) |
| 13 | Go Bases Tarea 4 | [código](cmd/13-tarea-4) · [doc](doc/13-tarea-4.md) |
| 14 | Funciones | [código](cmd/14-funciones) · [doc](doc/14-funciones.md) |
| 15 | Struct | [código](cmd/15-struct) · [doc](doc/15-struct.md) |
| 16 | Interface | [código](cmd/16-interface) · [doc](doc/16-interface.md) |
| 17 | Punteros | [código](cmd/17-punteros) · [doc](doc/17-punteros.md) |
| 18 | Errores | [código](cmd/18-errores) · [doc](doc/18-errores.md) |
| 19 | Go Bases Tarea 5 | [código](cmd/19-tarea-5) · [doc](doc/19-tarea-5.md) |
| 20 | `time`, `os` y `log` | [código](cmd/20-packages-tiempo-os-log) · [doc](doc/20-packages-tiempo-os-log.md) |
| 21 | `strconv`, `net` y `reflect` | [código](cmd/21-packages-strconv-net-reflect) · [doc](doc/21-packages-strconv-net-reflect.md) |
| 22 | Packages externos | [código](cmd/22-packages-externos) · [doc](doc/22-packages-externos.md) |
| 23 | Cambios en `dh-kit` | [código](cmd/23-dh-kit) · [doc](doc/23-dh-kit.md) |
| 24 | Crear nuestro propio package | [código](cmd/24-package-propio) · [doc](doc/24-package-propio.md) |
| 25 | Generics | [código](cmd/25-generics) · [doc](doc/25-generics.md) |
| 26 | Generics Ordered | [código](cmd/26-generics-ordered) · [doc](doc/26-generics-ordered.md) |
| 27 | Goroutines, canales y paralelismo | [código](cmd/27-goroutines-canales) · [doc](doc/27-goroutines-canales.md) |
| 28 | Context | [código](cmd/28-context) · [doc](doc/28-context.md) |
| 29 | Operadores binarios | [código](cmd/29-binary-operators) · [doc](doc/29-binary-operators.md) |

## Comandos útiles

```bash
go run ./cmd/25-generics
go test ./...
gofmt -w .
```

La lección 40 (“Siguiente proyecto”) es una transición del curso y no contiene código.

La numeración fue normalizada de forma continua para que el proyecto empiece en `01` y no tenga saltos entre carpetas.
