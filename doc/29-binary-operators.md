# 46-binary-operators

AND, OR, XOR, complemento y desplazamientos de bits.

## ¡Aprendamos jugando!

Piensa en **binary operators** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/29-binary-operators
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
left, right := uint8(0b1100), uint8(0b1010)
fmt.Printf("AND: %04b\n", left&right)
fmt.Printf("OR: %04b\n", left|right)
fmt.Printf("XOR: %04b\n", left^right)
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
