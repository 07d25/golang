# 32-interface

Interfaces definen comportamiento y se implementan implícitamente.

## ¡Aprendamos jugando!

Piensa en **interface** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/16-interface
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
type Sound interface { MakeSound() string }
type Dog struct{}
func (Dog) MakeSound() string { return "woof" }
func playSound(value Sound) { fmt.Println(value.MakeSound()) }
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
