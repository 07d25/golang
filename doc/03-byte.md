# 19-byte

`byte` es alias de `uint8`, útil para bytes binarios y UTF-8.

## ¡Aprendamos jugando!

Piensa en **byte** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/03-byte
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
text := "Go"
rawBytes := []byte(text)
fmt.Println(rawBytes)
fmt.Println(string(rawBytes))
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
