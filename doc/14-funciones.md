# 30-funciones

Funciones con varios resultados y patrón `(valor, error)`.

## ¡Aprendamos jugando!

Piensa en **funciones** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/14-funciones
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
func divide(left, right float64) (float64, error) {
if right == 0 { return 0, fmt.Errorf("division by zero") }
return left / right, nil
}
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
