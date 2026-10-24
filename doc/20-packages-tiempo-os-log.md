# 36-packages-tiempo-os-log

`time`, `os` y `log` para tiempo, sistema y diagnóstico.

## ¡Aprendamos jugando!

Piensa en **packages tiempo os log** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/20-packages-tiempo-os-log
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
start := time.Now()
time.Sleep(10 * time.Millisecond)
log.Println("elapsed:", time.Since(start))
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
