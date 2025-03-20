![kurkod](https://github.com/ChainsAre2Tight/qr-decoder-utils/blob/master/.github/memes/kurkod.jpeg)
# Утилита для декодирования QR-кодов

## Установка
###
Скачать последнюю версию из раздела Releases.

### Ручная установка
**(необходим golang версии > 1.22)**
```
$ git clone https://github.com/ChainsAre2Tight/qr-decoder-utils
$ cd qr-decoder-utils
# для Linux:
$ go build -o qr_decoder ./cmd/main.go
$ chmod +x qr_decoder
# для Windows:
> go build -o qr_decoder.exe ./cmd/main.go
```

## Использование
### Основные параметры

* **mode** string - основной режим работы **[image | excel | decode | mask]**
  * **image** - конвертирует изображение в пиксельный формат QR-кода.
  * **excel** - конвертирует изображение в `.xlsx` файл, может добавлять маски при использовании флага **--include-masks**.
  * **decode** - пытается декодировать содержимое QR-кода и выводит результат в `stdout`.
    * Поддерживаемые форматы:
      * QR Version 1, 2 и 3
        - numeric
        - byte
  * **mask** - генерирует указанную маску и сохраняет в `.xlsx` файл.

* **--input** string - путь к входному изображению.
  - Поддерживаемые форматы: **png** | **jpg** | **gif** (1 кадр).
* **--output** string - путь к выходному файлу (для режимов `image`, `excel` и `mask`). Если не указан, будет сгенерировано случайное имя.
* **--mask** string - тип маски **[000-111]**.
* **--size** int
  * В режиме **mask** - размер матрицы маски (по умолчанию 21x21).
  * В других режимах - задаёт размер QR-кода на изображении (диапазон 1-100, используется для отладки).
* **--include-masks** bool - добавляет маски в `.xlsx` файлы при конвертации.

### Примеры использования

* #### Конвертация `.gif` изображения в `.xlsx` с масками
```
qr_decoder excel --input ./image.gif --output ./result.xlsx --include-masks
```
* #### Конвертация `.jpg` изображения в `.xlsx` с переопределением размера кода
```
qr_decoder excel --input ./image.jpg --output ./result.xlsx --size 25
```
* #### Декодирование содержимого изображения
```
qr_decoder decode --input ./image.jpg
```
* #### Создание маски `010`
```
qr_decoder mask --mask 010 --output ./mask010.xlsx --size 25
```
