# 25-maps

Mapas clave-valor; la segunda devolución indica existencia.

## ¡Aprendamos jugando!

Piensa en **maps** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/09-maps
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
prices := map[string]float64{"bread": 1.5}
prices["coffee"] = 2.25
price, exists := prices["coffee"]
if exists { fmt.Println(price) }
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
