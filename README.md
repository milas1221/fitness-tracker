# Модуль фитнес-трекера

Pet-проект на Go, демонстрирующий применение структур, методов, интерфейсов и организации кода в пакеты.  
Проект представляет собой модуль для обработки данных о тренировках (бег, ходьба) и дневной активности (прогулки), расчёта дистанции, средней скорости и потраченных калорий.  
Логика разделена по пакетам, что делает код модульным, тестируемым и расширяемым.

## Структура проекта

```
.github/workflows/       # CI (GitHub Actions)
  unit-tests.yaml        # запуск тестов
cmd/tracker/             # точка входа (main.go)
internal/                
  actioninfo/            # интерфейс и функция вывода информации
  daysteps/              # обработка прогулок
  personaldata/          # данные пользователя
  spentenergy/           # расчёты калорий, дистанции, скорости
  trainings/             # обработка тренировок
README.md                # этот файл
go.mod                   # модуль Go
```

## Пакеты и их ответственность

### `spentenergy`
Содержит математические функции для расчётов:
- `Distance(steps int, height float64) float64` – дистанция в км.
- `MeanSpeed(steps int, height float64, duration time.Duration) float64` – средняя скорость (км/ч).
- `RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error)` – калории при беге.
- `WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error)` – калории при ходьбе.

Все функции экспортируемые, используются другими пакетами.

### `personaldata`
Определяет структуру пользователя:
```go
type Personal struct {
    Name   string
    Weight float64
    Height float64
}
```
Метод `Print()` выводит данные на экран:
```
Имя: <имя>
Вес: <вес>
Рост: <рост>
```

### `trainings`
Работа с данными о тренировках (бег/ходьба).  
Структура `Training`:
```go
type Training struct {
    Steps        int
    TrainingType string
    Duration     time.Duration
    personaldata.Personal // встраивание
}
```
Методы:
- `Parse(datastring string) error` – разбор строки вида `"3456,Ходьба,3h00m"`.
- `ActionInfo() (string, error)` – формирует отчёт:
  ```
  Тип тренировки: Бег
  Длительность: 0.75 ч.
  Дистанция: 10.00 км.
  Скорость: 13.34 км/ч
  Сожгли калорий: 18621.75
  ```

### `daysteps`
Аналогично, но для дневных прогулок (без типа активности).  
Структура `DaySteps`:
```go
type DaySteps struct {
    Steps    int
    Duration time.Duration
    personaldata.Personal
}
```
Методы:
- `Parse(datastring string) error` – разбор строки `"678,0h50m"`.
- `ActionInfo() (string, error)` – отчёт:
  ```
  Количество шагов: 792.
  Дистанция составила 0.51 км.
  Вы сожгли 221.33 ккал.
  ```

### `actioninfo`
Обеспечивает единый интерфейс для обработки любых активностей.  
Интерфейс `DataParser`:
```go
type DataParser interface {
    Parse(string) error
    ActionInfo() (string, error)
}
```
Типы `Training` и `DaySteps` автоматически реализуют этот интерфейс.

Функция `Info(dataset []string, dp DataParser)`:
- принимает набор строк с данными и любой объект, удовлетворяющий `DataParser`;
- для каждой строки вызывает `Parse()`, логирует ошибки;
- если парсинг успешен, вызывает `ActionInfo()` и выводит результат.

## Пример использования (из `cmd/tracker/main.go`)

```go
package main

import (
    "yourmodule/internal/actioninfo"
    "yourmodule/internal/daysteps"
    "yourmodule/internal/personaldata"
    "yourmodule/internal/trainings"
)

func main() {
    // Данные пользователя
    person := personaldata.Personal{
        Name:   "Иван",
        Weight: 75.0,
        Height: 1.75,
    }

    // Обработка тренировок
    trainingData := []string{
        "9000,Бег,1h30m",
        "3000,Ходьба,0h45m",
    }
    training := trainings.Training{Personal: person}
    actioninfo.Info(trainingData, &training)

    // Обработка прогулок
    dayStepsData := []string{
        "15000,2h0m",
        "500,0h15m",
    }
    daySteps := daysteps.DaySteps{Personal: person}
    actioninfo.Info(dayStepsData, &daySteps)
}
```

## Установка и запуск

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/yourusername/fitness-tracker.git
   cd fitness-tracker
   ```
2. Убедитесь, что установлена Go 1.18+.
3. Соберите и запустите:
   ```bash
   go run cmd/tracker/main.go
   ```
4. Запуск тестов:
   ```bash
   go test ./...
   ```

## Требования

- Go 1.18 или новее.

## Цели проекта (для саморазвития)

- Практика декомпозиции кода на пакеты.
- Использование встраивания структур.
- Реализация интерфейсов и полиморфизма.
- Написание unit-тестов (примеры тестов есть в каждом пакете).

## Автор

Берсиров Салим
