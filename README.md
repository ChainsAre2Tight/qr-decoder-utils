# Утилита для декодирования QR кодов

## Установка
###
Скачать последнюю версию из раздела Releases
### Ручная установка
**(необходим golang версии > 1.22)**
```
$ git clone https://github.com/ChainsAre2Tight/qr-decoder-utils
$ cd qr-decoder-utils
для linux:
$ go build -o qr_decoder.exe ./cmd/main.go
$ chmod +x qr_decoder.exe
для windows:
> go build -o qr_decoder.exe ./cmd/main.go
```
## Использование
### Запуск
```
$ ./qr_decoder
> qr_decoder.exe
```
### Параметры

* **mode** string - основной режим работы **[convert | decode | mask]**

  * **convert** - модуль конвертирования, переводит изображение в excel таблицу или изменяет его размеры
  * **decode** - попытается декодировать содержимое кода на изображении и выводит в stdout.
    * Поддерживаемые форматы:
      * QR Version 1 (21x21)
        - numeric
        - byte
  * **mask** - выводит указанную маску в виде .xlsx файла для дальнейшего копирования

* **--format** string - режим работы модуля convert **[image | excel]**
    
    * image - конвертирует выбранное изображение произвольного размера в 21x21
    * excel - конвертирует выбранное изображение в .xlsx файл
* **--input** string - путь до изображения на ввод
    * Поддерживаемые форматы:
      * png
      * jpg
      * gif (1 кадр)
* **--output** string - путь, куда записать результат (для режимов image, excel и mask)
* **--mask** string - тип маски **[000-111]**
* **--size** int- размер матрицы с маской (по умолчанию 21x21)

## Примеры использования
* #### Конвертация .gif изображения в .xlsx
```
  > qr_decoder.exe convert --format excel --input ./image.gif --output ./result
```
* #### Декодирование содержимого изображения
```
  > qr_decoder.exe decode --input ./image.jpg
```
* #### Создание маски 010
```
  > qr_decoder.exe mask --mask 010 --output ./mask010.xlsx
```
