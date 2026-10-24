# 35-tarea-5

Función variádica que calcula un promedio y valida entrada vacía.

## ¡Aprendamos jugando!

Piensa en **tarea 5** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/19-tarea-5
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
func calculateAverage(grades ...float64) (float64, error) {
if len(grades) == 0 { return 0, fmt.Errorf("no grades") }
total := 0.0
for _, grade := range grades { total += grade }
return total / float64(len(grades)), nil
}
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
