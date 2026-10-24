# 21-condicionales

`if` y `switch` controlan el flujo sin paréntesis.

## ¡Aprendamos jugando!

Piensa en **condicionales** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/05-condicionales
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
score := 8
switch {
case score < 5: fmt.Println("reprobado")
case score == 10: fmt.Println("perfecto")
default: fmt.Println("aprobado")
}
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
