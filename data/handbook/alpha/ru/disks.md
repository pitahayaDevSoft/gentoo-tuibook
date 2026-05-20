# Disks

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Disks/de "Handbuch:Alpha/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Disks "Handbook:Alpha/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Disks/tr "Handbook:Alpha/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Disks/es "Manual de Gentoo: Alpha/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Disks/fr "Manuel:Alpha/Installation/Disques (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Disks/it "Handbook:Alpha/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Disks/hu "Handbook:Alpha/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Disks/pt-br "Manual:Alpha/Instalação/Discos (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Disks/cs "Handbook:Alpha/Installation/Disks/cs (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Disks/ta "கையேடு:Alpha/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Disks/zh-cn "手册：Alpha/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Disks/ja "ハンドブック:Alpha/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Disks/ko "Handbook:Alpha/Installation/Disks/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha/ru "Handbook:Alpha/ru")[Установка](/wiki/Handbook:Alpha/Full/Installation/ru "Handbook:Alpha/Full/Installation/ru")[Об установке](/wiki/Handbook:Alpha/Installation/About/ru "Handbook:Alpha/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:Alpha/Installation/Media/ru "Handbook:Alpha/Installation/Media/ru")[Настройка сети](/wiki/Handbook:Alpha/Installation/Networking/ru "Handbook:Alpha/Installation/Networking/ru")Подготовка дисков[Установка файла stage](/wiki/Handbook:Alpha/Installation/Stage/ru "Handbook:Alpha/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:Alpha/Installation/Base/ru "Handbook:Alpha/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:Alpha/Installation/Kernel/ru "Handbook:Alpha/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:Alpha/Installation/System/ru "Handbook:Alpha/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:Alpha/Installation/Tools/ru "Handbook:Alpha/Installation/Tools/ru")[Настройка загрузчика](/wiki/Handbook:Alpha/Installation/Bootloader/ru "Handbook:Alpha/Installation/Bootloader/ru")[Завершение](/wiki/Handbook:Alpha/Installation/Finalizing/ru "Handbook:Alpha/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:Alpha/Full/Working/ru "Handbook:Alpha/Full/Working/ru")[Введение в Portage](/wiki/Handbook:Alpha/Working/Portage/ru "Handbook:Alpha/Working/Portage/ru")[USE-флаги](/wiki/Handbook:Alpha/Working/USE/ru "Handbook:Alpha/Working/USE/ru")[Возможности Portage](/wiki/Handbook:Alpha/Working/Features/ru "Handbook:Alpha/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:Alpha/Working/Initscripts/ru "Handbook:Alpha/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:Alpha/Working/EnvVar/ru "Handbook:Alpha/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:Alpha/Full/Portage/ru "Handbook:Alpha/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:Alpha/Portage/Files/ru "Handbook:Alpha/Portage/Files/ru")[Переменные](/wiki/Handbook:Alpha/Portage/Variables/ru "Handbook:Alpha/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:Alpha/Portage/Branches/ru "Handbook:Alpha/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:Alpha/Portage/Tools/ru "Handbook:Alpha/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:Alpha/Portage/CustomTree/ru "Handbook:Alpha/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:Alpha/Portage/Advanced/ru "Handbook:Alpha/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:Alpha/Full/Networking/ru "Handbook:Alpha/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:Alpha/Networking/Introduction/ru "Handbook:Alpha/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:Alpha/Networking/Advanced/ru "Handbook:Alpha/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:Alpha/Networking/Modular/ru "Handbook:Alpha/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:Alpha/Networking/Wireless/ru "Handbook:Alpha/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:Alpha/Networking/Extending/ru "Handbook:Alpha/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:Alpha/Networking/Dynamic/ru "Handbook:Alpha/Networking/Dynamic/ru")

## Contents

- [1Введение в блочные устройства](#.D0.92.D0.B2.D0.B5.D0.B4.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B2_.D0.B1.D0.BB.D0.BE.D1.87.D0.BD.D1.8B.D0.B5_.D1.83.D1.81.D1.82.D1.80.D0.BE.D0.B9.D1.81.D1.82.D0.B2.D0.B0)
  - [1.1Блочные устройства](#.D0.91.D0.BB.D0.BE.D1.87.D0.BD.D1.8B.D0.B5_.D1.83.D1.81.D1.82.D1.80.D0.BE.D0.B9.D1.81.D1.82.D0.B2.D0.B0)
  - [1.2Слайсы](#.D0.A1.D0.BB.D0.B0.D0.B9.D1.81.D1.8B)
- [2Разработка схемы разделов](#.D0.A0.D0.B0.D0.B7.D1.80.D0.B0.D0.B1.D0.BE.D1.82.D0.BA.D0.B0_.D1.81.D1.85.D0.B5.D0.BC.D1.8B_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2)
  - [2.1Сколько разделов и насколько больших?](#.D0.A1.D0.BA.D0.BE.D0.BB.D1.8C.D0.BA.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_.D0.B8_.D0.BD.D0.B0.D1.81.D0.BA.D0.BE.D0.BB.D1.8C.D0.BA.D0.BE_.D0.B1.D0.BE.D0.BB.D1.8C.D1.88.D0.B8.D1.85.3F)
  - [2.2Что по поводу пространства подкачки?](#.D0.A7.D1.82.D0.BE_.D0.BF.D0.BE_.D0.BF.D0.BE.D0.B2.D0.BE.D0.B4.D1.83_.D0.BF.D1.80.D0.BE.D1.81.D1.82.D1.80.D0.B0.D0.BD.D1.81.D1.82.D0.B2.D0.B0_.D0.BF.D0.BE.D0.B4.D0.BA.D0.B0.D1.87.D0.BA.D0.B8.3F)
- [3Использование fdisk для создания разделов диска (только SRM)](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_fdisk_.D0.B4.D0.BB.D1.8F_.D1.81.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D1.8F_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_.D0.B4.D0.B8.D1.81.D0.BA.D0.B0_.28.D1.82.D0.BE.D0.BB.D1.8C.D0.BA.D0.BE_SRM.29)
  - [3.1Определение доступных дисков](#.D0.9E.D0.BF.D1.80.D0.B5.D0.B4.D0.B5.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B4.D0.BE.D1.81.D1.82.D1.83.D0.BF.D0.BD.D1.8B.D1.85_.D0.B4.D0.B8.D1.81.D0.BA.D0.BE.D0.B2)
  - [3.2Creating a BSD disklabel](#Creating_a_BSD_disklabel)
  - [3.3Удаление всех слайсов](#.D0.A3.D0.B4.D0.B0.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B2.D1.81.D0.B5.D1.85_.D1.81.D0.BB.D0.B0.D0.B9.D1.81.D0.BE.D0.B2)
  - [3.4Создание слайса подкачки](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.81.D0.BB.D0.B0.D0.B9.D1.81.D0.B0_.D0.BF.D0.BE.D0.B4.D0.BA.D0.B0.D1.87.D0.BA.D0.B8)
  - [3.5Creating the boot slice](#Creating_the_boot_slice)
  - [3.6Создание коревого слайса](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D0.BA.D0.BE.D1.80.D0.B5.D0.B2.D0.BE.D0.B3.D0.BE_.D1.81.D0.BB.D0.B0.D0.B9.D1.81.D0.B0)
  - [3.7Сохранение разметки слайсов и выход](#.D0.A1.D0.BE.D1.85.D1.80.D0.B0.D0.BD.D0.B5.D0.BD.D0.B8.D0.B5_.D1.80.D0.B0.D0.B7.D0.BC.D0.B5.D1.82.D0.BA.D0.B8_.D1.81.D0.BB.D0.B0.D0.B9.D1.81.D0.BE.D0.B2_.D0.B8_.D0.B2.D1.8B.D1.85.D0.BE.D0.B4)
- [4Использование fdisk для создания разделов диска (только ARC/AlphaBIOS)](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_fdisk_.D0.B4.D0.BB.D1.8F_.D1.81.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D1.8F_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_.D0.B4.D0.B8.D1.81.D0.BA.D0.B0_.28.D1.82.D0.BE.D0.BB.D1.8C.D0.BA.D0.BE_ARC.2FAlphaBIOS.29)
  - [4.1Определение доступных дисков](#.D0.9E.D0.BF.D1.80.D0.B5.D0.B4.D0.B5.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B4.D0.BE.D1.81.D1.82.D1.83.D0.BF.D0.BD.D1.8B.D1.85_.D0.B4.D0.B8.D1.81.D0.BA.D0.BE.D0.B2_2)
  - [4.2Удаление всех разделов](#.D0.A3.D0.B4.D0.B0.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B2.D1.81.D0.B5.D1.85_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2)
  - [4.3Создание загрузочного раздела](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BE.D1.87.D0.BD.D0.BE.D0.B3.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0)
  - [4.4Создание раздела подкачки](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0_.D0.BF.D0.BE.D0.B4.D0.BA.D0.B0.D1.87.D0.BA.D0.B8)
  - [4.5Создание корневого раздела](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D0.BA.D0.BE.D1.80.D0.BD.D0.B5.D0.B2.D0.BE.D0.B3.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0)
  - [4.6Сохранение разметки разделов и выход](#.D0.A1.D0.BE.D1.85.D1.80.D0.B0.D0.BD.D0.B5.D0.BD.D0.B8.D0.B5_.D1.80.D0.B0.D0.B7.D0.BC.D0.B5.D1.82.D0.BA.D0.B8_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_.D0.B8_.D0.B2.D1.8B.D1.85.D0.BE.D0.B4)
- [5Создание файловых систем](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D1.8B.D1.85_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC)
  - [5.1Введение](#.D0.92.D0.B2.D0.B5.D0.B4.D0.B5.D0.BD.D0.B8.D0.B5)
  - [5.2Файловые системы](#.D0.A4.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D1.8B.D0.B5_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B)
  - [5.3Создание файловой системы на разделе](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D0.BE.D0.B9_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B_.D0.BD.D0.B0_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B5)
    - [5.3.1Файловая система загрузочного раздела Legacy BIOS](#.D0.A4.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D0.B0.D1.8F_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D0.B0_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BE.D1.87.D0.BD.D0.BE.D0.B3.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0_Legacy_BIOS)
    - [5.3.2Маленькие разделы ext4](#.D0.9C.D0.B0.D0.BB.D0.B5.D0.BD.D1.8C.D0.BA.D0.B8.D0.B5_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D1.8B_ext4)
  - [5.4Активация раздела подкачки](#.D0.90.D0.BA.D1.82.D0.B8.D0.B2.D0.B0.D1.86.D0.B8.D1.8F_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0_.D0.BF.D0.BE.D0.B4.D0.BA.D0.B0.D1.87.D0.BA.D0.B8)
- [6Монтирование корневого раздела](#.D0.9C.D0.BE.D0.BD.D1.82.D0.B8.D1.80.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_.D0.BA.D0.BE.D1.80.D0.BD.D0.B5.D0.B2.D0.BE.D0.B3.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0)

## Введение в блочные устройства

### Блочные устройства

Теперь взглянем на аспекты работы Gentoo Linux и Linux в общем, связанные с дисковой подсистемой, включая блочные устройства, разделы и файловые системы Linux. Как только основные понятия о дисках и файловых системах будут изучены, можно будет приступать к созданию разделов и файловых систем для установки.

Для начала, рассмотрим блочные устройства. Устройства SCSI и Serial ATA обозначаются как /dev/sda, /dev/sdb, /dev/sdc и так далее. На более современных компьютерах твердотельные накопители NVMe на базе PCI Express имеют дескриптор вида /dev/nvme0n1, /dev/nvme0n2 и так далее.

Следующая таблица поможет определить необходимый тип блочного устройства в системе:

Тип устройстваДескриптор устройства по умолчаниюПримечания и полезные сведения
IDE, SATA, SAS, SCSI или USB flash/dev/sdaДанный тип устройств стал доступным примерно с 2007 года и встречается до сих пор, являясь, пожалуй, самым используемым типом в Linux. Устройства могут подключаться через шины [SATA](https://en.wikipedia.org/wiki/ru:Serial_ATA "wikipedia:ru:Serial ATA"), [SCSI](https://en.wikipedia.org/wiki/ru:SCSI "wikipedia:ru:SCSI") или [USB](https://en.wikipedia.org/wiki/ru:USB "wikipedia:ru:USB") в виде устройства блочного хранилища. Например, первый раздел на первом SATA устройстве обозначается как /dev/sda1.
NVM Express (NVMe)/dev/nvme0n1Передовая на данный момент технология твердотельных накопителей. Устройства [NVMe](https://en.wikipedia.org/wiki/ru:NVM_Express "wikipedia:ru:NVM Express") подключаются к шине PCI Express и обладают наиболее быстрой скоростью передачи блочных данных, доступной на рынке. Системы образца 2014 года и новее могут иметь поддержку устройств NVMe. Первый раздел на первом NVMe устройстве обозначается как /dev/nvme0n1p1.
MMC, eMMC и SD/dev/mmcblk0Устройства [embedded MMC](https://en.wikipedia.org/wiki/ru:MultiMediaCard#eMMC "wikipedia:ru:MultiMediaCard"), SD-карты и другие типы карт памяти могут использоваться для хранения данных. Однако не все системы могут позволить загружаться с данного типа устройств. **Не рекомендуется** использовать данные устройства для установки Linux; лучше используйте их по прямому назначению — для переноса файлов. Также их можно использовать для кратковременного резервного копирования или снимков диска.
VirtIO/dev/vda[Виртуализованные](/wiki/Virtualization "Virtualization") устройства, которые можно найти при эмуляции с помощью [QEMU](/wiki/QEMU "QEMU"). Драйвер `virtio_blk` предоставляет доступ к выделенному пространству хоста для конкретноый виртуальной машины. Тем не менее, это отличный способ опробовать Gentoo на виртуальной машине.

Данные блочные устройства представляют абстрактный интерфейс к диску. Пользовательские приложения могут использовать их для взаимодействия с диском, не заботясь о том, какой это диск — SATA, SCSI или какой-либо ещё. Программа просто адресует пространство на диске как совокупность следующих друг за другом 4096-байтных (4K) блоков с произвольным доступом.

### Слайсы

Несмотря на то, что теоретически возможно использовать весь диск для размещения системы Linux, это почти никогда не делается на практике. Вместо этого, блочное устройство разбивается на меньшие, более управляемые блочные устройства. В системах Alpha они называются _слайсами_.

**Заметка**

В следующих разделах для установки будет использоваться пример для конфигурации ARC/AlphaBIOS. Измените необходимые значения под свою конфигурацию!

## Разработка схемы разделов

### Сколько разделов и насколько больших?

Расположение разделов на диске сильно зависит от потребностей системы и файловой системы (файловых систем). Если в ней будет много пользователей, рекомендуется разместить /home на отдельном разделе, что улучшит безопасность и значительно упростит резервное копирование (а также другие операции сопровождения). Если Gentoo устанавливается для использования в роли почтового сервера, следует отделить /var, так как вся почта хранится в каталоге /var. Для игровых серверов потребуется отдельный раздел /opt, так как большинство игровых серверов устанавливается туда. Причины выделения те же, что и для каталога /home: безопасность, резервное копирование и сопровождение.

В большинстве случаев /usr и /var должны быть достаточно большого размера. В /usr хранится большинство приложений, доступных системе, а также исходные коды ядра Linux (в каталоге /usr/src). По умолчанию в /var хранится репозиторий пакетов Gentoo (расположенный в /var/db/repos/gentoo), который, в зависимости от файловой системы, занимает около 650 МиБ дискового пространства. Оценка этого пространства _не включает_ каталоги /var/cache/distfiles и /var/cache/binpkgs, в которых будут скапливаться архивы исходных кодов и (не обязательно) двоичных пакетов, которые будут формироваться в самой системе.

Сколько именно и какого объёма разделов нужно системе — всё зависит от сочетания различных факторов, которые необходимо принимать во внимание. Наличие отдельных разделов или томов имеет следующие плюсы:

- Можно выбрать наиболее подходящую файловую систему для каждого раздела или тома.
- Свободное место во всей системе не закончится внезапно из-за того, что одна-единственная сбойная программа постоянно записывает файлы в раздел или том.
- Необходимая проверка файловых систем будет занимать меньше времени, так как проверка разных разделов может выполняться параллельно (хотя это это преимущество относится больше к нескольким дискам, чем к нескольким разделам).
- Можно повысить безопасность системы, монтируя часть разделов в режиме только для чтения, `nosuid` (игнорируется бит setuid), `noexec` (игнорируется бит исполнения) и так далее.

Однако у множества разделов также есть недостатки:

- Если они не настроены правильно, может получиться так, что будет огромное количество свободного места на одном разделе и нехватка на другом.
- Отдельный раздел для /usr/ может потребовать загрузки с initramfs, чтобы смонтировать раздел прежде, чем запустятся другие загрузочные сценарии. Так как сборка initramfs выходит за рамки данного руководства, **мы рекомендуем новичкам не создавать отдельный раздел для /usr/**.
- Также существует лимит в 15 разделов для SCSI и SATA, если только на диске не используются метки GPT.

**Заметка**

Если планируется использовать systemd в качестве системы инициализации и управления службами, то при загрузке системы должен быть доступен каталог /usr, либо как часть корневой файловой системы, либо смонтированный через initramfs.

### Что по поводу пространства подкачки?

Рекомендации для размера раздела подкачки
Размер ОЗУПоддержка режима приостановки?Поддержка спящего режима?
2 Гб и ниже4 Гб4 Гб
от 2 до 8 Гбразмер ОЗУ2 \* ОЗУ
от 8 до 64 Гбминимум 8 Гб, максимум 16 Гб1,5 \* ОЗУ
64 Гб и вышеминимум 8 ГбСпящий режим не рекомендуется! Спящий режим _не рекомендуется_ для систем с очень большим объёмом памяти. Хотя это и возможно, но для перевода в спящий режим необходимо записывать на диск всё содержимое памяти. Запись десятков гигабайт (и больше!) на диск может занять значительное время, особенно, когда используются обычные вращающиеся диски. В этом случае лучше использовать режим приостановки.

Не существует идеального значения для раздела подкачки. Целью пространства подкачки является предоставление дискового пространства ядру в условиях активного использования оперативной памяти. Пространство подкачки позволяет ядру переносить на диск страницы внутренней динамической (оперативной) памяти, которые не будут использоваться в ближайшее время, освобождая её (swap или page-out). Конечно, если эта память вдруг неожиданно понадобится, эти страницы должны быть помещены обратно в память (page-in), что займет намного больше времени, чем чтение с оперативной памяти (так как диски — это очень медленные устройства по сравнению с оперативной памятью).

Если на системе не требуется запускать приложения, требовательные к памяти, либо изначально доступно очень много памяти, то, скорее всего, необходимости в пространстве подкачки нет. Однако раздел подкачки также используется для сохранения _всей памяти_ в случае перехода системы в спящий режим (более вероятно на ноутбуках и десктопах, чем на серверах). Если планируется использовать этот режим, нужно пространство подкачки, равное или больше чем количеству оперативной памяти.

Как правило, рекомендуется создавать пространство подкачки для систем с оперативной памяти (ОЗУ) меньше 4 Гб, при этом размер подкачки рекомендуется должен быть в два раза больше ОЗУ. Для систем с несколькими дисками, целесообразно создать по одному разделу подкачки на каждом диске, чтобы их можно было использовать для параллельных операций чтения/записи. Чем быстрее диск может подкачивать, тем быстрее система будет работать, когда ей необходимо прочитать данные с пространства подкачки. При выборе между жестким диском и твердотельным накопителем, с точки зрения производительности лучше создать пространство подкачки на твердотельных носителях.

Стоит заметить, что в качестве альтернативы вместо _разделов_ подкачки можно использовать _файлы_ подкачки; это может быть полезно для систем с очень ограниченным размером дискового пространства.

## Использование fdisk для создания разделов диска (только SRM)

Далее будет объяснено, как создать примерную разметку слайсов для SRM:

Слайс
Описание
/dev/sda1Слайс раздела подкачки (swap)
/dev/sda2Корневой слайс (root)
/dev/sda3Весь диск (необходимо)

Измените структуру слайсов в соответствии с личными предпочтениями.

### Определение доступных дисков

Используйте следующие команды, чтобы выяснить, какие диски доступны в системе.

Для дисков IDE:

`root #` `dmesg | grep 'drive$'`

Для дисков SCSI:

`root #` `dmesg | grep 'scsi'`

Вывод команды отобразит обнаруженные диски и путь к ним в /dev/. Мы предположим, что это SCSI-диск /dev/sda.

### Creating a BSD disklabel

If the hard drive is completely blank, then first create a BSD disklabel. On alpha, this can't be done using fdisk, so using parted is the way forward here.

Now fire up parted:

`root #` `parted /dev/sda `

```
Using /dev/sda
Welcome to GNU Parted! Type 'help' to view a list of commands.

```

`(parted)` `mklabel bsd`

```
Warning: The existing disk label on /dev/sda will be destroyed and all data on this disk will be lost. Do you want to continue?
Yes/No? yes
(parted) quit
Information: You may need to update /etc/fstab.

```

Now that we have a bsd disklabel on our drive, continue creating slices. This can be done using parted or, as in the examples below, using fdisk:

`root #` `fdisk /dev/sda`

### Удаление всех слайсов

Если жёсткий диск полностью пуст, сначала создайте метку диска BSD.

`Command (m for help):` `b`

```
/dev/sda contains no disklabel.
Do you want to create a disklabel? (y/n) y
A bunch of drive-specific info will show here
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

Мы начнем с удаления всех слайсов кроме 'c'-слайса (необходим для использования меток BSD). Следующий пример показывает как удалить слайс (в примере мы используем 'a'). Повторите этот процесс, чтобы удалить остальные слайсы (все кроме слайса 'c').

Используйте `p`, чтобы просмотреть все доступные слайсы. `d` используется для удаления слайса.

`BSD disklabel command (m for help):` `p`

```
8 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        1       235*      234*    4.2BSD     1024  8192    16
  b:      235*      469*      234*      swap
  c:        1      5290*     5289*    unused        0     0
  d:      469*     2076*     1607*    unused        0     0
  e:     2076*     3683*     1607*    unused        0     0
  f:     3683*     5290*     1607*    unused        0     0
  g:      469*     1749*     1280     4.2BSD     1024  8192    16
  h:     1749*     5290*     3541*    unused        0     0

```

`BSD disklabel command (m for help):` `d`

```
Partition (a-h): a

```

После выполнения этой операции со всеми слайсами список должен показывать что-то подобное:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

### Создание слайса подкачки

В системах на базе Alpha нет необходимости создавать отдельный загрузочный слайс. Однако, нельзя использовать первый цилиндр, так как там будет размещаться образ _aboot_.

Мы создадим слайс подкачки размером в 1 ГБ, начиная с третьего цилиндра. Используйте `n`, чтобы создать новый слайс. После создания слайса мы изменим его тип на `1` (один), обозначающий подкачку.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): a
First cylinder (1-5290, default 1): 3
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): a
Hex code (type L to list codes): 1

```

После выполнения этих операций разметка должна выглядеть примерно так:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  c:        1      5290*     5289*    unused        0     0

```

### Creating the boot slice

Create a boot file system containing the kernels and boot loader configuration file (aboot.conf). aboot supports ext2 and ext3 filesystems only. Create the boot slice starting from the first cylinder after the swap slice. Use the `p` command to view where the swap slice ends. In our example, this is at 1003, making the boot slice start at 1004.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 1004
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 08

```

After these steps a layout similar to the following should be shown:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      2005      1001       ext2
  c:        1      5290*     5289*    unused        0     0

```

### Создание коревого слайса

Мы создадим корневой слайс начиная с первого цилиндра после слайса подкачки. Используйте `p`, что посмотреть, где заканчивается слайс подкачки. В нашем примере это 1003\. Создайте корневой слайс начиная с 1004.

Другая проблема заключается в том, что в fdisk обнаружена ошибка, из-за которой число доступных цилиндров на единицу больше реального числа цилиндров. Другими словами, при запросе последнего цилиндра, уменьшите номер цилиндра (в этом примере: 5290) на один.

Когда слайс создан, мы изменим тип на 8 для ext2.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 1004
Last cylinder or +size or +sizeM or +sizeK (1004-5290, default 5290): 5289

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 8

```

Итоговая разметка слайсов теперь должна выглядеть примерно так:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      2005      4286       ext2
  c:        1      5290*     5289*    unused        0     0
  d:     2006      5289      3283       ext2

```

### Сохранение разметки слайсов и выход

Выйдите из приложения fdisk, нажав `w`. Это также сохранит разметку слайсов.

`Command (m for help):` `w`

## Использование fdisk для создания разделов диска (только ARC/AlphaBIOS)

Далее будет объяснено, как создать примерную разметку разделов для ARC/AlphaBIOS:

Раздел
Описание
/dev/sda1Загрузочный раздел (boot)
/dev/sda2Раздел подкачки (swap)
/dev/sda3Корневой раздел (root)

Измените структуру разделов в соответствии с личными предпочтениями.

### Определение доступных дисков

Используйте следующие команды, чтобы выяснить, какие диски доступны в системе.

Для дисков IDE:

`root #` `dmesg | grep 'drive$'`

Для дисков SCSI:

`root #` `dmesg | grep 'scsi'`

Вывод команды отобразит обнаруженные диски и путь к ним в /dev/. Мы предположим, что это SCSI-диск /dev/sda.

Теперь запустите fdisk:

`root #` `fdisk /dev/sda`

### Удаление всех разделов

Если жёсткий диск полностью пуст, сначала создайте метку диска DOS.

`Command (m for help):` `o`

```
Building a new DOS disklabel.

```

Мы начнем с удаления всех разделов. Следующий пример показывает, как удалить раздел (в примере мы используем '1'). Повторите этот процесс, чтобы удалить остальные разделы.

Используйте `p`, чтобы просмотреть все доступные разделы. `d` используется для удаления раздела.

`command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1         478      489456   83  Linux
/dev/sda2             479        8727     8446976    5  Extended
/dev/sda5             479        1433      977904   83  Linux Swap
/dev/sda6            1434        8727     7469040   83  Linux

```

`command (m for help):` `d`

```
Partition number (1-6): 1

```

### Создание загрузочного раздела

В системах Alpha, в которых используется MILO для загрузки, необходимо создать небольшой загрузочный раздел vfat.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-8727, default 1): 1
Last cylinder or +size or +sizeM or +sizeK (1-8727, default 8727): +16M

```

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 6
Changed system type of partition 1 to 6 (FAT16)

```

### Создание раздела подкачки

Создадим раздел подкачки размером в 1 ГБ. Используйте `n`, чтобы создать новый раздел.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 2
First cylinder (17-8727, default 17): 17
Last cylinder or +size or +sizeM or +sizeK (17-8727, default 8727): +1000M

```

`Command (m for help):` `t`

```
Partition number (1-4): 2
Hex code (type L to list codes): 82
Changed system type of partition 2 to 82 (Linux swap)

```

После выполнения этих операций разметка должна выглядеть примерно так:

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap

```

### Создание корневого раздела

Теперь создадим корневой раздел. Снова, воспользуйтесь `n`.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 3
First cylinder (972-8727, default 972): 972
Last cylinder or +size or +sizeM or +sizeK (972-8727, default 8727): 8727

```

После выполнения этих операций разметка должна выглядеть примерно так:

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap
/dev/sda3             972        8727     7942144   83  Linux

```

### Сохранение разметки разделов и выход

Сохраните сделанные изменения в fdisk, нажав `w`.

`Command (m for help):` `w`

Теперь, когда разделы созданы, продолжайте с раздела [Создание файловых систем](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D1.8B.D1.85_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC).

## Создание файловых систем

**Предупреждение**

При использовании SSD или NVMe диска рекомендуется проверить наличие обновлений для его прошивки. Определённые SSD от Intel (600p и 6000p) нуждаются в обновлении прошивки для [исправления возможной потери данных](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) из-за особенностей использования I/O в XFS. Проблема проявляется на уровне прошивки и не является ошибкой самой файловой системы XFS. Программа smartctl умеет отображать модель устройства и версию прошивки.

### Введение

Теперь, когда разделы созданы, пора разместить на них файловые системы. В следующем разделе описаны различные поддерживаемые в Linux файловые системы. Те из читателей, кто уже знает, какую файловую систему будет использовать, могут продолжить с раздела [Создание файловой системы на разделе](/wiki/Handbook:Alpha/Installation/Disks/ru#Applying_a_filesystem_to_a_partition "Handbook:Alpha/Installation/Disks/ru"). Остальным стоит продолжить чтение, чтобы узнать о доступных вариантах…

### Файловые системы

Linux поддерживает несколько десятков файловых систем, хотя для большинства из них необходимы достаточно веские причины их использовать. Лишь только некоторые из них можно считать стабильными на архитектуре alpha. Рекомендуется прочитать информацию о файловых системах и об их состоянии поддержки перед тем, как останавливать свой выбор на экспериментальных. **XFS — рекомендуемая стабильная файловая система общего применения для всех платформ**. Ниже представлен неполный список файловых систем:

[XFS](/wiki/XFS/ru "XFS/ru")Файловая система с журналированием метаданных, которая поставляется с мощным набором функций и оптимизирована для масштабируемости. Она непрерывно обновляется, обрастая новым возможностями. Единственным недостатком является то, что разделы с XFS пока нельзя уменьшать (хотя и над этим ведётся работа). Примечательно, что XFS поддерживает «reflinks» и механизм «копирование при записи» (Copy-on-Write, CoW), что весьма полезно для Gentoo-систем из-за частых компиляций, которые совершает пользователь. XFS является рекомендуемой современной файловой системой общего назначения для всех платформ. Для неё требуется раздел размером не менее 300 МБ.[ext4](/wiki/Ext4/ru "Ext4/ru")Ext4 является стабильной файловой системой общего применения для всех платформ, хотя в ней отсутствуют современные возможности по типу «reflinks».[VFAT](/wiki/VFAT "VFAT")Так же известная как FAT32, поддерживается Linux, но не имеет поддержку стандартных файловых разрешений UNIX. В основном используется для взаимодействия или взаимозаменяемости с другими операционными системами (в основном Microsoft Windows и Apple macOS), но также необходима при использовании некоторых системных прошивок загрузчика (например, UEFI). Пользователям систем с UEFI должны использовать эту файловую систему для [EFI System Partition](/wiki/EFI_System_Partition/ru "EFI System Partition/ru"), чтобы иметь возможность загружаться.[btrfs](/wiki/Btrfs/ru "Btrfs/ru")Файловая система нового поколения. Предоставляет множество продвинутых функций, таких как мгновенные снимки, самовосстановление с помощью контрольных сумм, поддержка прозрачного сжатия, подтомов и интегрированный RAID. Ядра старше ветки 5.4 не обеспечивают безопасную работу btrfs, так как исправления наиболее серьёзных проблем стабильности появились только в более поздних ветках долговременной поддержки (LTS) ядра. RAID 5/6 и quota groups небезопасны на всех версиях btrfs.[f2fs](/wiki/F2fs "F2fs")Файловая система (Flash-Friendly File System) была создана Samsung для использования на NAND-накопителях. Она может быть достойным выбором при установке на microSD карту, USB-накопитель или другие накопители.[NTFS](/wiki/NTFS/ru "NTFS/ru")New Technology Filesystem является основной файловой системой для Microsoft Windows начиная с NT 3.5. Как и VFAT, она не сохраняет настройки UNIX разрешений и расширенные атрибуты, необходимые для нормальной работы BSD или Linux, поэтому в большинстве случаев её не следует использоваться в качестве файловой системы для корневого раздела. Её следует использовать _только_ для взаимодействия или обмена данными с системами Microsoft Windows (обратите внимание на акцент слова _только_).[ZFS](/wiki/ZFS "ZFS")**Важно:** Пулы ZFS можно создавать в окружении минимального установочного диска; для дополнительной информации обратитесь к странице [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs")Файловая система следующего поколения, созданная Мэттью Эхренсом и Джеффом Бонуиком. Она была создана на основе нескольких ключевых идей: администрирование хранилища должно быть простым, отказоустойчивость обеспечивается самой файловой системой, файловые системы никогда не должны быть недоступными для возможности ремонта, важны автоматизированные симуляции самых тяжёлых сценариев перед публикацией кода, а целостность данных является первостепенной.

Более подробную информацию о файловых системах можно найти в (поддерживаемой сообществом) статье [Файловая система](/wiki/Filesystem/ru "Filesystem/ru").

### Создание файловой системы на разделе

**Заметка**

Пожалуйста, не забудьте, что в конце установки (перед перезагрузкой системы) вам необходимо установить соответствующие пакеты пользовательского окружения для выбранной файловой системы. В конце процесса установки будет дополнительное напоминание об этом.

Для создания файловых систем на разделе или томе существуют пользовательские утилиты для каждого возможного типа файловой системы. Нажмите на имя файловой системы в таблице ниже для получения дополнительной информации о каждой файловой системе:

Файловая система
Команда для создания
Существует в «живом» окружении?
Пакет
[XFS](/wiki/XFS/ru "XFS/ru")mkfs.xfs да
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/ru "Ext4/ru")mkfs.ext4 да
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
mkfs.vfat да
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs/ru "Btrfs/ru")mkfs.btrfs да
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS "F2FS")mkfs.f2fs да
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS/ru "NTFS/ru")mkfs.ntfs да
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... да
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

**Важно**

Руководство рекомендует создавать новые разделы в рамках процесса установки, однако важно помнить, что выполнение команды mkfs удалит любые данные, содержащиеся в разделе. Убедитесь, что перед созданием новой файловой системы для существующих данных надлежащим образом была сделана резервная копия.

Например, чтобы отформатировать корневой раздел (/dev/sda3) в xfs при использовании структуры разделов из примера, используются следующие команды:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf /dev/sda3`

#### Файловая система загрузочного раздела Legacy BIOS

В системах, загружаемых через устаревший BIOS с использованием метки диска MBR/DOS, можно использовать любой поддерживаемый загрузчиком формат файловой системы.

Например, чтобы отформатировать раздел в формат XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf /dev/sda1`

#### Маленькие разделы ext4

При использовании файловой системы ext4 на малых разделах (менее 8 ГиБ) её необходимо создавать с особыми параметрами для выделения достаточного количества индексных дескрипторов (inodes). Это можно сделать, указав параметр `-T small`:

`root #` `mkfs.ext4 -T small /dev/<раздел>`

Благодаря этому количество индексных дескрипторов будет увеличено вчетверо для указанной файловой системы за счёт того, что параметр «bytes-per-inode» был уменьшен с 16 до 4 кб.

### Активация раздела подкачки

Для инициализации разделов подкачки используется команда mkswap:

`root #` `mkswap /dev/sda2`

**Заметка**

Ранее начатый но незаконченный процесс установки может быть продолжен начиная с этого места Руководства. Используйте эту ссылку в качестве постоянной ссылки: [Возобновление установки начинается здесь](/wiki/Handbook:Alpha/Installation/Disks/ru#Resumed_installations_start_here "Handbook:Alpha/Installation/Disks/ru").

Чтобы активировать раздел подкачки, используйте swapon:

`root #` `swapon /dev/sda2`

Шаг «активации» необходим потому, что раздел подкачки был только что создан в живом окружении. Как только система будет перезагружена, и если раздел подкачки правильно определён в fstab или в другом средстве монтирования разделов, он будет активироваться автоматически.

## Монтирование корневого раздела

В некоторых окружениях может не хватать предлагаемой для корневого раздела Gentoo точки монтирования (/mnt/gentoo), либо точек монтирования для дополнительных разделов, созданных в разделе «Создание разделов»:

`root #` `mkdir --parents /mnt/gentoo`

По необходимости создайте с помощью команды mkdirдополнительные точки монтирования для других разделов, созданных в предыдущих шагах установки.

После создания точек монтирования настало время сделать их доступными через команду mount.

Смонтируем корневой раздел:

`root #` `mount /dev/sda3 /mnt/gentoo`

При необходимости смонтируйте дополнительные разделы с помощью команды mount.

**Заметка**

Если /tmp/ находится на отдельном разделе, не забудьте после монтирования изменить права доступа:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Это также справедливо для /var/tmp.

Позже в инструкции будут смонтированы файловая система proc (виртуальный интерфейс к ядру) и другие псевдофайловые системы ядра. Но сначала необходимо распаковать [файл стадии Gentoo](/wiki/Handbook:Alpha/Installation/Stage/ru "Handbook:Alpha/Installation/Stage/ru").

[← Настройка сети](/wiki/Handbook:Alpha/Installation/Networking/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:Alpha/ru "Handbook:Alpha/ru") [Установка Gentoo файлов установки →](/wiki/Handbook:Alpha/Installation/Stage/ru "Следующая часть")