# 29-tarea-4

Recorrer ventas y acumular un total.

## ¡Aprendamos jugando!

Piensa en **tarea 4** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/13-tarea-4
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
sales := map[string]float64{"january": 120, "february": 180}
total := 0.0
for _, sale := range sales { total += sale }
fmt.Println(total)
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
