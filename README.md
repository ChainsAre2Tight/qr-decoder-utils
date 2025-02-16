# Утилита для декодирования QR кодов

## Установка
**(необходим golang версии 1.23.2)**
```
$ git clone https://github.com/ChainsAre2Tight/qr-decoder-utils
$ cd qr-decoder-utils
$ make build
```
## Использование
### Запуск
```
$ ./bin/qr_decoder
```
### Параметры

* **--mode** - режим работы, может быть value, image и excel
    * value - попытается декодировать содержимое кода на изображении и вывод в stdout.
        * Поддерживаемые форматы:
          * QR Version 1 (21x21)
            - numeric
            - byte
    * image - конвертирует выбранное изображение произвольного размера в 21x21
    * excel - конвертирует выбранное изображение в .xlsx файл
* **--input** - путь до изображения на ввод
    * поддерживаемые форматы:
      * png
      * jpg
      * gif (1 кадр)
* **--output** - путь, куда записать результат (для режимов image и excel)

## Примеры использования
### Конвертация .gif изображения в .xlsx
```
$ ./bin/qr_decoder --mode excel --input ./image.gif --output ./result
```
### Декодирование содержимого изображения
```
$ ./bin/qr_decoder --mode value --input ./image.jpg
```