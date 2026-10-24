# 33-punteros

Punteros permiten leer y modificar el valor original.

## ¡Aprendamos jugando!

Piensa en **punteros** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/17-punteros
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
func doubleValue(value *int) { *value *= 2 }
number := 7
doubleValue(&number)
fmt.Println(number)
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
