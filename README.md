![kurkod](https://github.com/ChainsAre2Tight/qr-decoder-utils/blob/master/.github/memes/kurkod.jpeg)
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
### Параметры

* **mode** string - основной режим работы **[image | excel | decode | mask]**

  * **image** - модуль конвертации изображения в соотвтетствие с размером пикселя кода
  * **excel** - модуль конвертации изображения в .xlsx файл, может добавлять маски при использовании флага **--include-masks**
  * **decode** - попытается декодировать содержимое кода на изображении и выводит в stdout.
    * Поддерживаемые форматы:
      * QR Version 1, 2 и 3
        - numeric
        - byte
  * **mask** - выводит указанную маску в виде .xlsx файла для дальнейшего копирования
* **--input** string - путь до изображения на ввод
    - Поддерживаемые форматы: **png** | **jpg** | **gif** (1 кадр)
* **--output** string - путь, куда записать результат (для режимов image, excel и mask), если отсутствует - будет сгенерировано рандомное имя
* **--mask** string - тип маски **[000-111]**
* **--size** int
  * в режиме **mask** -  размер матрицы с маской (по умолчанию 21x21)
  * в других режимах - пеопределяет размер кода на изображении (дебаг)
* **--include-masks** bool - дописывает листы с масками при конвертации в .xlsx файлы

### Примеры использования
* #### Конвертация .gif изображения в .xlsx с масками
```
qr_decoder.exe excel --input ./image.gif --output ./result --include-masks
```
* #### Конвертация .jpg изображения в .xlsx с переопределением размера кода
```
qr_decoder.exe excel --input ./image.jpg --output ./result --size 25
```
* #### Декодирование содержимого изображения
```
qr_decoder.exe decode --input ./image.jpg
```
* #### Создание маски 010
```
qr_decoder.exe mask --mask 010 --output ./mask010.xlsx --size 25
```
