# 34-errores

Los errores son valores que se comprueban con `err != nil`.

## ¡Aprendamos jugando!

Piensa en **errores** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/18-errores
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
value, err := squareRoot(-1)
if err != nil { fmt.Println("no se pudo calcular:", err); return }
fmt.Println(value)
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
