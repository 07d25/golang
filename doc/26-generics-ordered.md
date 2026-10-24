# 43-generics-ordered

Restricción `Ordered`, `~` y comparación con `<`.

## ¡Aprendamos jugando!

Piensa en **generics ordered** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/26-generics-ordered
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
type Ordered interface { ~int | ~float64 | ~string }
func Maximum[T Ordered](left, right T) T {
if left > right { return left }
return right
}
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
