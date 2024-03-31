# Thumbnail    [![Go Test](https://github.com/lirprocs/thumbnail/actions/workflows/test.yml/badge.svg)](https://github.com/lirprocs/thumbnail/actions/workflows/test.yml)

## Описание
Проект thumbnail предоставляет сервер и утилиту командной строки для загрузки и кэширования миниатюр видеороликов с YouTube. Для кеширования используется SQLite.

## Установка
1. Клонируйте репозиторий
```bash
git clone https://github.com/lirprocs/thumbnail.git
```
2. Перейдите в директорию проекта:
```bash
cd thumbnail
```
3. Установите зависимости:
```bash
go mod tidy
```

## Запуск сервера
1. Перейдите в директорию проекта (Не нужно, еслу уже находитесь в ней):
```bash
cd thumbnail
```
2. Запустите сервер:
```bash
go run cmd/thumbnail/main.go 
```

## Запуск утилиты командной строки
1. Откройте второе окно терминала.
2. Перейдите в директорию проекта (Не нужно, еслу уже находитесь в ней):
```bash
cd thumbnail
```
3. Соберите утилиту:
```bash
go build -o my.exe  cmd\thumbnail_cli\main_cli.go
```
4. Запустите утилиту для загрузки миниатюры видеоролика с YouTube:
```bash
./my.exe https://www.youtube.com/watch?v=446E-r0rXHI
```
4.1. Для асинхронной загрузки миниатюр видеороликов используйте флаг --async:
```bash
./my.exe --async https://www.youtube.com/watch?v=446E-r0rXHI https://www.youtube.com/watch?v=5C_HPTJg5ek
```

## Примичание
В проекте thumbnail настроен автоматический запуск тестов при пуше изменений в репозиторий. При отправке изменений на GitHub, GitHub Actions автоматически запустит тесты для проверки работоспособности кода.
