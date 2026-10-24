# 27-tarea-2

Filtrar pares usando el resto de la división.

## ¡Aprendamos jugando!

Piensa en **tarea 2** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/11-tarea-2
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
numbers := []int{2, 7, 4, 9, 6}
for _, number := range numbers {
if number%2 == 0 && number > 3 { fmt.Println(number) }
}
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
