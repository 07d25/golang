# 01. Variables

Una variable almacena un valor y tiene un tipo. En Go se puede declarar con `var`, indicando el tipo explícitamente, o con `:=`, dejando que el compilador infiera el tipo dentro de una función.

También se muestran constantes, valores cero, tipos enteros de tamaño fijo y `unsafe.Sizeof` para consultar el tamaño ocupado en memoria.

Ejecutar:

```bash
go run ./cmd/01-variables
```

## ¡Aprendamos jugando!

Piensa en **variables** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/01-variables
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
name := "Ana"       // Go infiere que es string.
age := 25           // Go infiere que es int.
const country = "AR" // Una constante no cambia.
fmt.Println(name, age, country)
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
