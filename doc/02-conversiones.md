# 18-conversiones

Conversiones explícitas entre tipos numéricos; pueden perder precisión.

## ¡Probémoslo!

Imagina que tienes una moneda entera y necesitas calcular con decimales: Go te obliga a cambiarla de tipo de forma explícita. Así los errores no aparecen escondidos.

```go
integerValue := 42
decimalValue := float64(integerValue)
fmt.Println(decimalValue)
```

🎯 **Reto:** convierte una temperatura entera a `float64` y calcula su mitad.

## Ejemplo paso a paso

```go
integerValue := 7
decimalValue := float64(integerValue)
result := decimalValue / 2
fmt.Println(result)
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
