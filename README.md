![kurkod](https://github.com/ChainsAre2Tight/qr-decoder-utils/blob/master/.github/memes/kurkod.jpeg)
# Утилита для декодирования QR-кодов

## Установка
###
Скачать последнюю версию из раздела Releases.

### Ручная установка
**(необходим golang версии > 1.22)**
```bash
$ git clone https://github.com/ChainsAre2Tight/qr-decoder-utils
$ cd qr-decoder-utils
$ go run cmd/main.go
```

## Использование
### Основные параметры

* **mode** string - основной режим работы **[image | excel | decode | mask]**
  * **image** - конвертирует изображение в размер, совпадающий с размером кода.
  * **excel** - конвертирует изображение в `.xlsx` файл, может добавлять маски **QR**-кодов при использовании флага **--include-masks**.
  * **decode** - пытается декодировать содержимое кода и выводит результат в `stdout`.
    * Поддерживаемые форматы:
      * QR Version 1 - 4
        - numeric
        - byte (кодировки ISO8859-1 и UTF-8)
      * Datamatrix (ECC200, square) **Внимание**: Необходимо предоставлять размер матриицы через **--size**
        - byte
  * **mask** - генерирует указанную маску и сохраняет в `.xlsx` файл.

* **--input** string - путь к входному изображению.
  - Поддерживаемые форматы: **png** | **jpg** | **gif** (1 кадр).
* **--output** string - путь к выходному файлу (для режимов `image`, `excel` и `mask`). Если не указан, будет сгенерировано случайное имя.
* **--mask** string - тип маски **QR**-кода **[000-111]**.
* **--size** int
  * В режиме **mask** - размер матрицы маски (по умолчанию 21x21).
  * В других режимах - задаёт размер кода на изображении (диапазон 1-100).
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
