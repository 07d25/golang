# 37-packages-strconv-net-reflect

`strconv`, `net` y `reflect` para conversiones, IPs e inspección.

## ¡Aprendamos jugando!

Piensa en **packages strconv net reflect** como una herramienta de tu mochila de Go. Primero mira el ejemplo, ejecútalo y cambia un valor para descubrir qué ocurre. La mejor forma de aprender es hacer una pequeña prueba y observar el resultado.

```bash
go run ./cmd/21-packages-strconv-net-reflect
```

🎯 **Reto:** modifica el ejemplo, rompe algo a propósito y luego intenta explicar por qué apareció ese resultado.

## Ejemplo paso a paso

```go
number, err := strconv.Atoi("123")
if err != nil { log.Println(err) }
ip := net.ParseIP("127.0.0.1")
fmt.Println(number, ip, reflect.TypeOf(number))
```

Lee el ejemplo, ejecútalo y cambia un solo valor. Después vuelve a ejecutar para comprobar qué cambió.

🎯 **Reto:** modifica el ejemplo para resolver una variante del mismo problema.
