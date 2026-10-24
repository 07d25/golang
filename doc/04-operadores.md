# 20-operadores

Aritméticos, relacionales y lógicos; `%` devuelve el resto.

## ¡Aprendamos jugando!

Piensa en **operadores** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/04-operadores
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
number := 12
isPositive := number > 0
isEven := number%2 == 0
fmt.Println(isPositive && isEven)
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
