# Разлика в сумите

Напишете функция, която намира разликата между квадрата на сборa и сбора на квадратите на първите `n` естествени числа:

```go
func SquareSumDifference(n uint64) uint64
```

Ето и един пример. Сбора на квадратите на първите десет (`n` = 10) числа е:

```
1² + 2² + 3² + ... + 10² = 385
```

Квадрата на сборa на първите десет числа е:

```
(1 + 2 + 3 + ... + 10)² = 55² = 3025
```

Значи `SquareSumDifference(10)` трябва да върне `3025 - 385 = 2640`.

Не забравяйте, че вашата функция трябва да отговаря точно на дефиницията в условието и трябва да се намира в пакета `main`. Кодът, който предавате, трябва да е форматиран с `go fmt`.
