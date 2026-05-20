# Disks

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Disks/de "Handbuch:SPARC/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Disks "Handbook:SPARC/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Disks/es "Manual de Gentoo: SPARC/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Disks/fr "Handbook:SPARC/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Disks/it "Handbook:SPARC/Installation/Disks/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Disks/hu "Handbook:SPARC/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Disks/pl "Handbook:SPARC/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Disks/pt-br "Handbook:SPARC/Installation/Disks/pt-br (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Disks/ta "கையேடு:SPARC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Disks/zh-cn "手册：SPARC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Disks/ja "ハンドブック:SPARC/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Disks/ko "Handbook:SPARC/Installation/Disks/ko (100% translated)")

[SPARC Handbook](/wiki/Handbook:SPARC/ru "Handbook:SPARC/ru")[Установка](/wiki/Handbook:SPARC/Full/Installation/ru "Handbook:SPARC/Full/Installation/ru")[Об установке](/wiki/Handbook:SPARC/Installation/About/ru "Handbook:SPARC/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:SPARC/Installation/Media/ru "Handbook:SPARC/Installation/Media/ru")[Настройка сети](/wiki/Handbook:SPARC/Installation/Networking/ru "Handbook:SPARC/Installation/Networking/ru")Подготовка дисков[Установка файла stage](/wiki/Handbook:SPARC/Installation/Stage/ru "Handbook:SPARC/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:SPARC/Installation/Base/ru "Handbook:SPARC/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:SPARC/Installation/Kernel/ru "Handbook:SPARC/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:SPARC/Installation/System/ru "Handbook:SPARC/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:SPARC/Installation/Tools/ru "Handbook:SPARC/Installation/Tools/ru")[Настройка загрузчика](/wiki/Handbook:SPARC/Installation/Bootloader/ru "Handbook:SPARC/Installation/Bootloader/ru")[Завершение](/wiki/Handbook:SPARC/Installation/Finalizing/ru "Handbook:SPARC/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:SPARC/Full/Working/ru "Handbook:SPARC/Full/Working/ru")[Введение в Portage](/wiki/Handbook:SPARC/Working/Portage/ru "Handbook:SPARC/Working/Portage/ru")[USE-флаги](/wiki/Handbook:SPARC/Working/USE/ru "Handbook:SPARC/Working/USE/ru")[Возможности Portage](/wiki/Handbook:SPARC/Working/Features/ru "Handbook:SPARC/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:SPARC/Working/Initscripts/ru "Handbook:SPARC/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:SPARC/Working/EnvVar/ru "Handbook:SPARC/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:SPARC/Full/Portage/ru "Handbook:SPARC/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:SPARC/Portage/Files/ru "Handbook:SPARC/Portage/Files/ru")[Переменные](/wiki/Handbook:SPARC/Portage/Variables/ru "Handbook:SPARC/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:SPARC/Portage/Branches/ru "Handbook:SPARC/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:SPARC/Portage/Tools/ru "Handbook:SPARC/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:SPARC/Portage/CustomTree/ru "Handbook:SPARC/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:SPARC/Portage/Advanced/ru "Handbook:SPARC/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:SPARC/Full/Networking/ru "Handbook:SPARC/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:SPARC/Networking/Introduction/ru "Handbook:SPARC/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:SPARC/Networking/Advanced/ru "Handbook:SPARC/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:SPARC/Networking/Modular/ru "Handbook:SPARC/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:SPARC/Networking/Wireless/ru "Handbook:SPARC/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:SPARC/Networking/Extending/ru "Handbook:SPARC/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:SPARC/Networking/Dynamic/ru "Handbook:SPARC/Networking/Dynamic/ru")

## Contents

- [1Введение в блочные устройства](#.D0.92.D0.B2.D0.B5.D0.B4.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B2_.D0.B1.D0.BB.D0.BE.D1.87.D0.BD.D1.8B.D0.B5_.D1.83.D1.81.D1.82.D1.80.D0.BE.D0.B9.D1.81.D1.82.D0.B2.D0.B0)
  - [1.1Блочные устройства](#.D0.91.D0.BB.D0.BE.D1.87.D0.BD.D1.8B.D0.B5_.D1.83.D1.81.D1.82.D1.80.D0.BE.D0.B9.D1.81.D1.82.D0.B2.D0.B0)
  - [1.2Таблица разделов](#.D0.A2.D0.B0.D0.B1.D0.BB.D0.B8.D1.86.D0.B0_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2)
    - [1.2.1GUID Partition Table (GPT)](#GUID_Partition_Table_.28GPT.29)
    - [1.2.2Sun partition table](#Sun_partition_table)
  - [1.3Схема разделов по умолчанию](#.D0.A1.D1.85.D0.B5.D0.BC.D0.B0_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_.D0.BF.D0.BE_.D1.83.D0.BC.D0.BE.D0.BB.D1.87.D0.B0.D0.BD.D0.B8.D1.8E)
    - [1.3.1GPT partition scheme](#GPT_partition_scheme)
    - [1.3.2Sun formatted partition scheme](#Sun_formatted_partition_scheme)
- [2Создание разделов на диске с GPT](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_.D0.BD.D0.B0_.D0.B4.D0.B8.D1.81.D0.BA.D0.B5_.D1.81_GPT)
  - [2.1Просмотр текущей разметки разделов](#.D0.9F.D1.80.D0.BE.D1.81.D0.BC.D0.BE.D1.82.D1.80_.D1.82.D0.B5.D0.BA.D1.83.D1.89.D0.B5.D0.B9_.D1.80.D0.B0.D0.B7.D0.BC.D0.B5.D1.82.D0.BA.D0.B8_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2)
  - [2.2Creating a new disklabel and removing all existing partitions](#Creating_a_new_disklabel_and_removing_all_existing_partitions)
  - [2.3Создание загрузочного раздела BIOS](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BE.D1.87.D0.BD.D0.BE.D0.B3.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0_BIOS)
  - [2.4Создание раздела подкачки](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0_.D0.BF.D0.BE.D0.B4.D0.BA.D0.B0.D1.87.D0.BA.D0.B8)
  - [2.5Создание корневого раздела](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D0.BA.D0.BE.D1.80.D0.BD.D0.B5.D0.B2.D0.BE.D0.B3.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0)
  - [2.6Сохранение разметки разделов](#.D0.A1.D0.BE.D1.85.D1.80.D0.B0.D0.BD.D0.B5.D0.BD.D0.B8.D0.B5_.D1.80.D0.B0.D0.B7.D0.BC.D0.B5.D1.82.D0.BA.D0.B8_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2)
- [3Partitioning the disk with a Sun partition table](#Partitioning_the_disk_with_a_Sun_partition_table)
  - [3.1Просмотр текущей разметки разделов](#.D0.9F.D1.80.D0.BE.D1.81.D0.BC.D0.BE.D1.82.D1.80_.D1.82.D0.B5.D0.BA.D1.83.D1.89.D0.B5.D0.B9_.D1.80.D0.B0.D0.B7.D0.BC.D0.B5.D1.82.D0.BA.D0.B8_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_2)
  - [3.2Создание нового disklabel / удаление всех разделов](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D0.BD.D0.BE.D0.B2.D0.BE.D0.B3.D0.BE_disklabel_.2F_.D1.83.D0.B4.D0.B0.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B2.D1.81.D0.B5.D1.85_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2)
  - [3.3Creating the whole disk partition](#Creating_the_whole_disk_partition)
  - [3.4Создание корневого раздела](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D0.BA.D0.BE.D1.80.D0.BD.D0.B5.D0.B2.D0.BE.D0.B3.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0_2)
  - [3.5Создание раздела подкачки](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0_.D0.BF.D0.BE.D0.B4.D0.BA.D0.B0.D1.87.D0.BA.D0.B8_2)
  - [3.6Сохранение разметки разделов](#.D0.A1.D0.BE.D1.85.D1.80.D0.B0.D0.BD.D0.B5.D0.BD.D0.B8.D0.B5_.D1.80.D0.B0.D0.B7.D0.BC.D0.B5.D1.82.D0.BA.D0.B8_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_2)
- [4Создание файловых систем](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D1.8B.D1.85_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC)
  - [4.1Введение](#.D0.92.D0.B2.D0.B5.D0.B4.D0.B5.D0.BD.D0.B8.D0.B5)
  - [4.2Файловые системы](#.D0.A4.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D1.8B.D0.B5_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B)
  - [4.3Создание файловой системы на разделе](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D0.BE.D0.B9_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B_.D0.BD.D0.B0_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B5)
    - [4.3.1Файловая система загрузочного раздела Legacy BIOS](#.D0.A4.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D0.B0.D1.8F_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D0.B0_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BE.D1.87.D0.BD.D0.BE.D0.B3.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0_Legacy_BIOS)
    - [4.3.2Маленькие разделы ext4](#.D0.9C.D0.B0.D0.BB.D0.B5.D0.BD.D1.8C.D0.BA.D0.B8.D0.B5_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D1.8B_ext4)
  - [4.4Активация раздела подкачки](#.D0.90.D0.BA.D1.82.D0.B8.D0.B2.D0.B0.D1.86.D0.B8.D1.8F_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0_.D0.BF.D0.BE.D0.B4.D0.BA.D0.B0.D1.87.D0.BA.D0.B8)
- [5Монтирование корневого раздела](#.D0.9C.D0.BE.D0.BD.D1.82.D0.B8.D1.80.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_.D0.BA.D0.BE.D1.80.D0.BD.D0.B5.D0.B2.D0.BE.D0.B3.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.B0)

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

### Таблица разделов

Although it is theoretically possible to use a raw, unpartitioned disk to house a Linux system (when creating a btrfs RAID for example), this is almost never done in practice. Instead, disk block devices are split up into smaller, more manageable block devices. On **sparc** systems, these are called partitions. There are currently two standard partitioning technologies in use: Sun and GPT; the latter is supported only on more recent systems with a sufficiently recent firmware.

#### GUID Partition Table (GPT)

_GUID Partition Table_ (GPT, таблица разделов GUID, также может называться GPT disklabel) использует 64-битные идентификаторы разделов. Место, в котором хранится информация о разделах, также гораздо больше, чем 512 байт таблицы разделов MBR (DOS disklabel), что означает, что нет почти никаких ограничений на количество разделов для диска с GPT. Также максимальный _размер_ раздела был значительно увеличен (почти 8 ЗиБ — да, зебибайт).

Также GPT использует контрольные суммы и избыточность. Он содержит контрольные суммы CRC32 для обнаружения ошибок в заголовке и таблице разделов. У GPT есть резервная таблица в конце диска. Её можно использовать для восстановления первичной таблицы GPT, которая располагается в начале диска.

GPT is only supported on Oracle SPARC machines of the T4 generation or newer. Additionally, only certain more recent firmware includes GPT support. There are several methods to check whether GPT support is available.

From the OBP prompt, execute:

`{0} ok` `cd /packages/disk-label`

`{0} ok` `.properties`

gpt

supported-labels gpt

```
                   sun
                   mbr

```

name disk-label

If gpt is included in the output, then GPT support is available. Alternatively, this can be determined from the installation media without entering OBP. Use the prtconf command from [sys-apps/sparc-utils](https://packages.gentoo.org/packages/sys-apps/sparc-utils) to access this information from userspace:

`root #` `prtconf -pv | grep -c gpt`

Or, check if the file /sys/firmware/devicetree/base/packages/disk-label/gpt exists. If none of these methods succeeds, then a firmware update is required in order to support GPT.

#### Sun partition table

Systems not manufactured by Oracle, T3 or earlier systems, or systems running an earlier firmware must use the Sun partition table type.

На системах Sun третий раздел выделен отдельно и используется как специальный слайс «whole disk». Этот раздел **не должен** содержать какой-либо файловой системы.

Пользователи, привыкшие к схеме разделов в стиле DOS, должны отметить, что метки дисков Sun не имеют «основных» или «расширенных» разделов. Вместо этого на каждом устройстве может располагаться до восьми разделов, при этом третий зарезервирован.

Авторы Руководства при установке Gentoo рекомендуют использовать [GPT](#GPT) везде, где это возможно.

### Схема разделов по умолчанию

Due to the differences in required partition layout between GPT and Sun partition tables, a single partitioning scheme is not sufficient to support all possible system requirements. Some example schemes are provided below.

#### GPT partition scheme

The following partitioning scheme will be used as an example for GPT-formatted disks:

Раздел
Файловая система
Размер
Точка монтирования
Описание
/dev/sda1(нет)
2 Мб
нет
[BIOS boot partition](/wiki/Handbook:X86/Blocks/Disks#What_is_the_BIOS_boot_partition.3F.2Fru "Handbook:X86/Blocks/Disks")/dev/sda2(swap)
Размер ОЗУ \\* 2
нет
Раздел подкачки
/dev/sda3ext4
Оставшаяся часть диска
/Корневой раздел

#### Sun formatted partition scheme

The following partitioning scheme will be used as an example for Sun-formatted disks:

Раздел
Файловая система
Размер
Точка монтирования
Описание
/dev/sda1ext4
Размер диска минус размер раздела подкачки
/Корневой раздел
/dev/sda2(swap)
Размер ОЗУ \\* 2
нет
Раздел подкачки
/dev/sda3(none)
Весь диск
нет
Раздел всего диска. Нужен для дисков с использованием таблицы разделов Sun.

**Важно**

SPARC systems using OBP version 3 or older have additional restrictions on their partitioning scheme. The root partition must be the first partition on the disk, and it may be no larger than 2 GiB. For this reason, such systems will require additional sufficiently-sized partitions for top-level directories, such as /usr, /var, /home, and other directories which would likely cause the root partition to exceed this limit. These systems are also likely to require the Sun partition table type, so do not forget to include the whole disk partition.

## Создание разделов на диске с GPT

Следующие части объяснят, как создать структуру разделов из примера для установки с GPT с использованием fdisk (пример структуры разделов приводился выше):

Раздел
Описание
/dev/sda1Загрузочный раздел
/dev/sda2Раздел подкачки
/dev/sda3Корневой раздел

Change the partition layout according the system's needs.

### Просмотр текущей разметки разделов

fdisk является популярным и мощным инструментом для создания разделов на диске. Запустите fdisk, передав в качестве параметра имя диска (в нашем примере мы используем /dev/sda):

`root #` `fdisk /dev/sda`

Нажмите на клавишу `p` для отображения текущей конфигурации разделов:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 9850A2C2-76C4-FC47-9F0B-DA60449D2413

Device     Start      End  Sectors  Size Type
/dev/sda1   2048 30547967 30545920 14.6G Linux filesystem

```

### Creating a new disklabel and removing all existing partitions

Нажмите `g`, чтобы создать новую разметку GPT на диске; это удалит все существующие разделы.

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 9850A2C2-76C4-FC47-9F0B-DA60449D2413).

```

Для диска с существующей разметкой GPT (смотрите вывод `p` выше), вы также можете удалять существующие разделы на диске. Нажмите `d` для удаления раздела. Например, чтобы удалить существующий /dev/sda1:

`Command (m for help):` `d`

```
Selected partition 1
Partition 1 has been deleted.

```

Теперь раздел отмечен для удаления. Он больше не будет отображаться в списке разделов при вводе `p`, но не будет удален, пока не будут сохранены изменения. Это даёт возможность пользователю прервать операцию, если была допущена ошибка — в этом случае сразу нажмите `q` и `Enter`, и раздел не будет удален.

Удалите все разделы, поочерёдно нажимая на `p` для вывода списка разделов, `d` и номер раздела — для удаления. В конечном счете, таблица разделов будет пуста:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 9850A2C2-76C4-FC47-9F0B-DA60449D2413

```

Теперь, когда запомненная в памяти таблица разделов пуста, мы готовы создавать разделы.

### Создание загрузочного раздела BIOS

First, create the BIOS boot partition. Type `n` to create a new partition, followed by `1` to select the first partition. When prompted for the first sector, make sure it starts from 2048 (which may be needed for the boot loader) and hit `Enter`. When prompted for the last sector, type +2M to create a partition 2 Mbyte in size:

`Command (m for help):` `n`

```
Partition number (1-128, default 1):
First sector (2048-30548062, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-30548062, default 30547967): +2M

Created a new partition 1 of type 'Linux filesystem' and of size 2 MiB.

```

Mark the partition as a BIOS boot partition:

`Command (m for help):` `t`

```
Selected partition 1
Partition type or alias (type L to list all): 4
Changed type of partition 'Linux filesystem' to 'BIOS boot'.

```

### Создание раздела подкачки

Для создания раздела подкачки введите `n`, чтобы создать новый раздел, затем введите `2` для создания второго _основного_ раздела, /dev/sda2. При появлении запроса первого сектора, введите `Enter`. При появлении запроса последнего сектора, наберите +4G (или любой другой размер, необходимый для подкачки) для создания раздела размером 4 ГиБ.

`Command (m for help):` `n`

```
Partition number (2-128, default 2):
First sector (6144-30548062, default 6144):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (6144-30548062, default 30547967): +4G

Created a new partition 2 of type 'Linux filesystem' and of size 4 GiB.

```

После этого введите `t` для выбора типа раздела, `2` для выбора только что созданного раздела и введите _19_, чтобы установить тип раздела как «Linux Swap».

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type (type L to list all types): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Создание корневого раздела

Наконец, чтобы создать корневой раздел, введите `n`, чтобы создать новый раздел. Затем введите `3`, чтобы создать третий _основной_ раздел, /dev/sda3. При запросе последнего сектора нажмите `Enter`, чтобы создать раздел, занимающий всё оставшееся доступное пространство диска. После завершения этих шагов введите `p` для вывода на экран таблицы разделов, которая должна выглядеть примерно так:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 9850A2C2-76C4-FC47-9F0B-DA60449D2413

Device       Start      End  Sectors  Size Type
/dev/sda1     2048     6143     4096    2M BIOS boot
/dev/sda2     6144  8394751  8388608    4G Linux swap
/dev/sda3  8394752 30547967 22153216 10.6G Linux filesystem

```

### Сохранение разметки разделов

Для сохранения разметки разделов и выхода из fdisk введите `w`.

`Command (m for help):` `w`

Разделы созданы, теперь настало время создать на них файловые системы.

## Partitioning the disk with a Sun partition table

The following parts explain how to create the example partition layout for a Sun partition table installation using fdisk. The example partition layout was mentioned earlier:

Раздел
Описание
/dev/sda1Корневой раздел
/dev/sda2Раздел подкачки
/dev/sda3Раздел всего диска

Change the partition layout according to personal preference. If partitioning for a system using OBP version 3 or earlier, ensure that the root partition is less than 2G in size, and additionally create partitions /dev/sda4 and onward for additional filesystems.

### Просмотр текущей разметки разделов

fdisk является популярным и мощным инструментом для создания разделов на диске. Запустите fdisk, передав в качестве параметра имя диска (в нашем примере мы используем /dev/sda):

`root #` `fdisk /dev/sda`

Нажмите на клавишу `p` для отображения текущей конфигурации разделов:

`Command (m for help):` `p`

```
Disk model: USB Flash Disk
Geometry: 64 heads, 32 sectors/track, 14916 cylinders
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: sun

Device        Start      End  Sectors  Size Id Type         Flags
/dev/sda1         0 30445567 30445568 14.5G 83 Linux native
/dev/sda2  30445568 30547967   102400   50M 82 Linux swap      u
/dev/sda3         0 30547967 30547968 14.6G  5 Whole disk

```

### Создание нового disklabel / удаление всех разделов

Type `s` to create a new Sun disklabel on the disk; this will remove all existing partitions.

`Command (m for help):` `s`

```
Created a new partition 1 of type 'Linux native' and of size 14.5 GiB.
Created a new partition 2 of type 'Linux swap' and of size 50 MiB.
Created a new partition 3 of type 'Whole disk' and of size 14.6 GiB.
Created a new Sun disklabel.

```

For an existing Sun disklabel (see the output of `p` above), alternatively consider removing the existing partitions one by one from the disk. Type `d` to delete a partition. For instance, to delete an existing /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-3, default 3): 1

Partition 1 has been deleted.

```

Теперь раздел отмечен для удаления. Он больше не будет отображаться в списке разделов при вводе `p`, но не будет удален, пока не будут сохранены изменения. Это даёт возможность пользователю прервать операцию, если была допущена ошибка — в этом случае сразу нажмите `q` и `Enter`, и раздел не будет удален.

Удалите все разделы, поочерёдно нажимая на `p` для вывода списка разделов, `d` и номер раздела — для удаления. В конечном счете, таблица разделов будет пуста:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Geometry: 64 heads, 32 sectors/track, 14916 cylinders
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: sun

```

Теперь, когда запомненная в памяти таблица разделов пуста, мы готовы создавать разделы.

### Creating the whole disk partition

First, create the whole disk partition. Type `n` to create a new partition, followed by `3` to select the third partition. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to create a partition that takes up all of the space on the disk.

`Command (m for help):` `n`

```
Partition number (1-8, default 1): 3

It is highly recommended that the third partition covers the whole disk and is of type `Whole disk'
First sector (0-30547968, default 0):
Last sector or +/-sectors or +/-size{K,M,G,T,P} (0-30547968, default 30547968):

Created a new partition 3 of type 'Whole disk' and of size 14.6 GiB.

```

fdisk will automatically set the type of such a partition to 'Whole disk', so there is no need to explicitly set the type.

### Создание корневого раздела

Next, to create the root partition, type `n` to create a new partition. Then type `1` to create the third partition, /dev/sda1. When prompted for the first sector, hit `Enter`. When prompted for the last sector, type -4G (or whatever space is required for non-root partitions). After completing these steps, typing `p` should display a partition table that looks similar to this:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Geometry: 64 heads, 32 sectors/track, 14916 cylinders
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: sun

Device     Start      End  Sectors  Size Id Type         Flags
/dev/sda1      0 22159359 22159360 10.6G 83 Linux native
/dev/sda3      0 30547967 30547968 14.6G  5 Whole disk

```

### Создание раздела подкачки

Finally, to create the swap partition, type `n` to create a new partition, then type `2` to create the second partition, /dev/sda2. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to take up the remaining space on the disk.

`Command (m for help):` `n`

```
Partition number (2,4-8, default 2):
First sector (22159360-30547968, default 22159360):
Last sector or +/-sectors or +/-size{K,M,G,T,P} (22159360-30547968, default 30547968):

Created a new partition 2 of type 'Linux native' and of size 4 GiB.

```

После этого введите `t` для выбора типа раздела, `3` для выбора только что созданного раздела и введите _82_, чтобы установить тип раздела как «Linux Swap».

`Command (m for help):` `t`

```
Partition number (1-3, default 3): 2
Hex code (type L to list all codes): 82

Changed type of partition 'Linux native' to 'Linux swap'.

```

After completing these steps, typing `p` should display a partition table that looks similar to this:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Geometry: 64 heads, 32 sectors/track, 14916 cylinders
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: sun

Device        Start      End  Sectors  Size Id Type         Flags
/dev/sda1         0 22159359 22159360 10.6G 83 Linux native
/dev/sda2  22159360 30547967  8388608    4G 82 Linux swap      u
/dev/sda3         0 30547967 30547968 14.6G  5 Whole disk

```

### Сохранение разметки разделов

Для сохранения разметки разделов и выхода из fdisk введите `w`.

`Command (m for help):` `w`

Разделы созданы, теперь настало время создать на них файловые системы.

## Создание файловых систем

**Предупреждение**

При использовании SSD или NVMe диска рекомендуется проверить наличие обновлений для его прошивки. Определённые SSD от Intel (600p и 6000p) нуждаются в обновлении прошивки для [исправления возможной потери данных](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) из-за особенностей использования I/O в XFS. Проблема проявляется на уровне прошивки и не является ошибкой самой файловой системы XFS. Программа smartctl умеет отображать модель устройства и версию прошивки.

### Введение

Теперь, когда разделы созданы, пора разместить на них файловые системы. В следующем разделе описаны различные поддерживаемые в Linux файловые системы. Те из читателей, кто уже знает, какую файловую систему будет использовать, могут продолжить с раздела [Создание файловой системы на разделе](/wiki/Handbook:SPARC/Installation/Disks/ru#Applying_a_filesystem_to_a_partition "Handbook:SPARC/Installation/Disks/ru"). Остальным стоит продолжить чтение, чтобы узнать о доступных вариантах…

### Файловые системы

Linux поддерживает несколько десятков файловых систем, хотя для большинства из них необходимы достаточно веские причины их использовать. Лишь только некоторые из них можно считать стабильными на архитектуре sparc. Рекомендуется прочитать информацию о файловых системах и об их состоянии поддержки перед тем, как останавливать свой выбор на экспериментальных. **XFS — рекомендуемая стабильная файловая система общего применения для всех платформ**. Ниже представлен неполный список файловых систем:

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

Например, чтобы отформатировать корневой раздел (/dev/sda1) в xfs при использовании структуры разделов из примера, используются следующие команды:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf /dev/sda1`

#### Файловая система загрузочного раздела Legacy BIOS

В системах, загружаемых через устаревший BIOS с использованием метки диска MBR/DOS, можно использовать любой поддерживаемый загрузчиком формат файловой системы.

Например, чтобы отформатировать раздел в формат XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf`

#### Маленькие разделы ext4

При использовании файловой системы ext4 на малых разделах (менее 8 ГиБ) её необходимо создавать с особыми параметрами для выделения достаточного количества индексных дескрипторов (inodes). Это можно сделать, указав параметр `-T small`:

`root #` `mkfs.ext4 -T small /dev/<раздел>`

Благодаря этому количество индексных дескрипторов будет увеличено вчетверо для указанной файловой системы за счёт того, что параметр «bytes-per-inode» был уменьшен с 16 до 4 кб.

### Активация раздела подкачки

Для инициализации разделов подкачки используется команда mkswap:

`root #` `mkswap /dev/sda2`

**Заметка**

Ранее начатый но незаконченный процесс установки может быть продолжен начиная с этого места Руководства. Используйте эту ссылку в качестве постоянной ссылки: [Возобновление установки начинается здесь](/wiki/Handbook:SPARC/Installation/Disks/ru#Resumed_installations_start_here "Handbook:SPARC/Installation/Disks/ru").

Чтобы активировать раздел подкачки, используйте swapon:

`root #` `swapon /dev/sda2`

Шаг «активации» необходим потому, что раздел подкачки был только что создан в живом окружении. Как только система будет перезагружена, и если раздел подкачки правильно определён в fstab или в другом средстве монтирования разделов, он будет активироваться автоматически.

## Монтирование корневого раздела

В некоторых окружениях может не хватать предлагаемой для корневого раздела Gentoo точки монтирования (/mnt/gentoo), либо точек монтирования для дополнительных разделов, созданных в разделе «Создание разделов»:

`root #` `mkdir --parents /mnt/gentoo`

По необходимости создайте с помощью команды mkdirдополнительные точки монтирования для других разделов, созданных в предыдущих шагах установки.

После создания точек монтирования настало время сделать их доступными через команду mount.

Смонтируем корневой раздел:

`root #` `mount /dev/sda1 /mnt/gentoo`

При необходимости смонтируйте дополнительные разделы с помощью команды mount.

**Заметка**

Если /tmp/ находится на отдельном разделе, не забудьте после монтирования изменить права доступа:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Это также справедливо для /var/tmp.

Позже в инструкции будут смонтированы файловая система proc (виртуальный интерфейс к ядру) и другие псевдофайловые системы ядра. Но сначала необходимо распаковать [файл стадии Gentoo](/wiki/Handbook:SPARC/Installation/Stage/ru "Handbook:SPARC/Installation/Stage/ru").

[← Настройка сети](/wiki/Handbook:SPARC/Installation/Networking/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:SPARC/ru "Handbook:SPARC/ru") [Установка Gentoo файлов установки →](/wiki/Handbook:SPARC/Installation/Stage/ru "Следующая часть")