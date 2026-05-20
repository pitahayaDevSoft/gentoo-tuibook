# Disks

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Disks/de "Handbuch:MIPS/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Disks "Handbook:MIPS/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Disks/es "Manual de Gentoo: MIPS/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Disks/fr "Handbook:MIPS/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Disks/it "Handbook:MIPS/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Disks/hu "Handbook:MIPS/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Disks/pl "Handbook:MIPS/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Disks/pt-br "Manual:MIPS/Instalação/Discos (100% translated)")
- русский
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Disks/ta "கையேடு:MIPS/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Disks/zh-cn "手册：MIPS/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Disks/ja "ハンドブック:MIPS/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Disks/ko "Handbook:MIPS/Installation/Disks/ko (100% translated)")

[MIPS Handbook](/wiki/Handbook:MIPS/ru "Handbook:MIPS/ru")[Установка](/wiki/Handbook:MIPS/Full/Installation/ru "Handbook:MIPS/Full/Installation/ru")[Об установке](/wiki/Handbook:MIPS/Installation/About/ru "Handbook:MIPS/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:MIPS/Installation/Media/ru "Handbook:MIPS/Installation/Media/ru")[Настройка сети](/wiki/Handbook:MIPS/Installation/Networking/ru "Handbook:MIPS/Installation/Networking/ru")Подготовка дисков[Установка файла stage](/wiki/Handbook:MIPS/Installation/Stage/ru "Handbook:MIPS/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:MIPS/Installation/Base/ru "Handbook:MIPS/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:MIPS/Installation/Kernel/ru "Handbook:MIPS/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:MIPS/Installation/System/ru "Handbook:MIPS/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:MIPS/Installation/Tools/ru "Handbook:MIPS/Installation/Tools/ru")[Настройка загрузчика](/wiki/Handbook:MIPS/Installation/Bootloader/ru "Handbook:MIPS/Installation/Bootloader/ru")[Завершение](/wiki/Handbook:MIPS/Installation/Finalizing/ru "Handbook:MIPS/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:MIPS/Full/Working/ru "Handbook:MIPS/Full/Working/ru")[Введение в Portage](/wiki/Handbook:MIPS/Working/Portage/ru "Handbook:MIPS/Working/Portage/ru")[USE-флаги](/wiki/Handbook:MIPS/Working/USE/ru "Handbook:MIPS/Working/USE/ru")[Возможности Portage](/wiki/Handbook:MIPS/Working/Features/ru "Handbook:MIPS/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:MIPS/Working/Initscripts/ru "Handbook:MIPS/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:MIPS/Working/EnvVar/ru "Handbook:MIPS/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:MIPS/Full/Portage/ru "Handbook:MIPS/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:MIPS/Portage/Files/ru "Handbook:MIPS/Portage/Files/ru")[Переменные](/wiki/Handbook:MIPS/Portage/Variables/ru "Handbook:MIPS/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:MIPS/Portage/Branches/ru "Handbook:MIPS/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:MIPS/Portage/Tools/ru "Handbook:MIPS/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:MIPS/Portage/CustomTree/ru "Handbook:MIPS/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:MIPS/Portage/Advanced/ru "Handbook:MIPS/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:MIPS/Full/Networking/ru "Handbook:MIPS/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:MIPS/Networking/Introduction/ru "Handbook:MIPS/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:MIPS/Networking/Advanced/ru "Handbook:MIPS/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:MIPS/Networking/Modular/ru "Handbook:MIPS/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:MIPS/Networking/Wireless/ru "Handbook:MIPS/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:MIPS/Networking/Extending/ru "Handbook:MIPS/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:MIPS/Networking/Dynamic/ru "Handbook:MIPS/Networking/Dynamic/ru")

## Contents

- [1Введение в блочные устройства](#.D0.92.D0.B2.D0.B5.D0.B4.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B2_.D0.B1.D0.BB.D0.BE.D1.87.D0.BD.D1.8B.D0.B5_.D1.83.D1.81.D1.82.D1.80.D0.BE.D0.B9.D1.81.D1.82.D0.B2.D0.B0)
  - [1.1Блочные устройства](#.D0.91.D0.BB.D0.BE.D1.87.D0.BD.D1.8B.D0.B5_.D1.83.D1.81.D1.82.D1.80.D0.BE.D0.B9.D1.81.D1.82.D0.B2.D0.B0)
    - [1.1.1Разделы](#.D0.A0.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D1.8B)
- [2Разработка схемы разделов](#.D0.A0.D0.B0.D0.B7.D1.80.D0.B0.D0.B1.D0.BE.D1.82.D0.BA.D0.B0_.D1.81.D1.85.D0.B5.D0.BC.D1.8B_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2)
  - [2.1Сколько разделов и насколько больших?](#.D0.A1.D0.BA.D0.BE.D0.BB.D1.8C.D0.BA.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_.D0.B8_.D0.BD.D0.B0.D1.81.D0.BA.D0.BE.D0.BB.D1.8C.D0.BA.D0.BE_.D0.B1.D0.BE.D0.BB.D1.8C.D1.88.D0.B8.D1.85.3F)
  - [2.2Что по поводу пространства подкачки?](#.D0.A7.D1.82.D0.BE_.D0.BF.D0.BE_.D0.BF.D0.BE.D0.B2.D0.BE.D0.B4.D1.83_.D0.BF.D1.80.D0.BE.D1.81.D1.82.D1.80.D0.B0.D0.BD.D1.81.D1.82.D0.B2.D0.B0_.D0.BF.D0.BE.D0.B4.D0.BA.D0.B0.D1.87.D0.BA.D0.B8.3F)
- [3Использование fdisk](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_fdisk)
  - [3.1SGI машины: Создание метки диска SGI](#SGI_.D0.BC.D0.B0.D1.88.D0.B8.D0.BD.D1.8B:_.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D0.BC.D0.B5.D1.82.D0.BA.D0.B8_.D0.B4.D0.B8.D1.81.D0.BA.D0.B0_SGI)
  - [3.2Resizing the SGI volume header](#Resizing_the_SGI_volume_header)
  - [3.3Разбивка дисков Cobalt](#.D0.A0.D0.B0.D0.B7.D0.B1.D0.B8.D0.B2.D0.BA.D0.B0_.D0.B4.D0.B8.D1.81.D0.BA.D0.BE.D0.B2_Cobalt)
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

#### Разделы

Несмотря на то, что теоретически возможно использовать весь диск для размещения системы Linux, это почти никогда не делается на практике. Вместо этого, блочное устройство разбивается на меньшие, более управляемые блочные устройства. Они называются разделами.

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

## Использование fdisk

### SGI машины: Создание метки диска SGI

All disks in an SGI System require an SGI Disk Label, which serves a similar function as Sun & MS-DOS disklabels -- It stores information about the disk partitions. Creating a new SGI Disk Label will create two special partitions on the disk:

- SGI Volume Header (9th partition): This partition is important. It is where the bootloader will reside, and in some cases, it will also contain the kernel images.
- SGI Volume (11th partition): This partition is similar in purpose to the Sun Disklabel's third partition of "Whole Disk". This partition spans the entire disk, and should be left untouched. It serves no special purpose other than to assist the PROM in some undocumented fashion (or it is used by IRIX in some way).

**Предупреждение**

The SGI Volume Header must begin at cylinder 0. Failure to do so means a failure to boot from the disk.

The following is an example excerpt from an fdisk session. Read and tailor it to personal preference...

`root #` `fdisk /dev/sda`

Переключитесь в экспертный режим:

`Command (m for help):` `x`

With `m` the full menu of options is displayed:

`Expert command (m for help):` `m`

```
Command action
   b   move beginning of data in a partition
   c   change number of cylinders
   d   print the raw data in the partition table
   e   list extended partitions
   f   fix partition order
   g   create an IRIX (SGI) partition table
   h   change number of heads
   m   print this menu
   p   print the partition table
   q   quit without saving changes
   r   return to main menu
   s   change number of sectors/track
   v   verify the partition table
   w   write table to disk and exit

```

Build an SGI disk label:

`Expert command (m for help):` `g`

```
Building a new SGI disklabel. Changes will remain in memory only,
until you decide to write them. After that, of course, the previous
content will be irrecoverably lost.

```

Вернитесь в главное меню:

`Expert command (m for help):` `r`

Take a look at the current partition layout:

`Command (m for help):` `p`

```
Disk /dev/sda (SGI disk label): 64 heads, 32 sectors, 17482 cylinders
Units = cylinders of 2048 * 512 bytes

----- partitions -----
Pt#     Device  Info     Start       End   Sectors  Id  System
 9:  /dev/sda1               0         4     10240   0  SGI volhdr
11:  /dev/sda2               0     17481  35803136   6  SGI volume
----- Bootinfo -----
Bootfile: /unix
----- Directory Entries -----

```

**Заметка**

If the disk already has an existing SGI Disklabel, then fdisk will not allow the creation of a new label. There are two ways around this. One is to create a Sun or MS-DOS disklabel, write the changes to disk, and restart fdisk. The second is to overwrite the partition table with null data via the following command: `dd if=/dev/zero of=/dev/sda bs=512 count=1`

### Resizing the SGI volume header

**Важно**

This step is often needed, due to a bug in fdisk. For some reason, the volume header isn't created correctly, the end result being it starts and ends on cylinder 0. This prevents multiple partitions from being created. To get around this issue... read on.

Now that an SGI Disklabel is created, partitions may now be defined. In the above example, there are already two partitions defined. These are the special partitions mentioned above and should not normally be altered. However, for installing Gentoo, we'll need to load a bootloader, and possibly multiple kernel images (depending on system type) directly into the volume header. The volume header itself can hold up to eight images of any size, with each image allowed eight-character names.

The process of making the volume header larger isn't exactly straight-forward; there's a bit of a trick to it. One cannot simply delete and re-add the volume header due to odd fdisk behavior. In the example provided below, we'll create a 50MB Volume header in conjunction with a 50MB /boot/ partition. The actual layout of a disk may vary, but this is for illustrative purposes only.

Создайте новый раздел:

`Command (m for help):` `n`

```
Partition number (1-16): 1
First cylinder (5-8682, default 5): 51
 Last cylinder (51-8682, default 8682): 101

```

Notice how fdisk only allows Partition #1 to be re-created starting at a minimum of cylinder 5? If we attempted to delete & re-create the SGI Volume Header this way, this is the same issue we would have encountered. In our example, we want /boot/ to be 50MB, so we start it at cylinder 51 (the Volume Header needs to start at cylinder 0, remember?), and set its ending cylinder to 101, which will roughly be 50MB (+/- 1-5MB).

Удалите раздел:

`Command (m for help):` `d`

```
Partition number (1-16): 9

```

Теперь, создайте его снова:

`Command (m for help):` `n`

```
Partition number (1-16): 9
First cylinder (0-50, default 0): 0
 Last cylinder (0-50, default 50): 50

```

If unsure how to use fdisk have a look down further at the instructions for partitioning on Cobalts. The concepts are exactly the same -- just remember to leave the volume header and whole disk partitions alone.

Once this is done, create the rest of your partitions as needed. After all the partitions are laid out, make sure to set the partition ID of the swap partition to 82, which is Linux Swap. By default, it will be 83, Linux Native.

### Разбивка дисков Cobalt

On Cobalt machines, the BOOTROM expects to see a MS-DOS MBR, so partitioning the drive is relatively straightforward -- in fact, it's done the same way as done for an Intel x86 machine. However there are some things you need to bear in mind.

- Cobalt firmware will expect /dev/sda1 to be a Linux partition formatted EXT2 Revision 0. EXT2 Revision 1 partitions will NOT WORK! (The Cobalt BOOTROM only understands EXT2r0)
- The above said partition must contain a gzipped ELF image, vmlinux.gz in the root of that partition, which it loads as the kernel

For that reason, it is recommended to create a ~20MB /boot/ partition formatted EXT2r0 upon which to install CoLo & kernels. This allows the user to run a modern filesystem (EXT3 or ReiserFS) for the root filesystem.

In the example, it is assumed that /dev/sda1 is created to mount later as a /boot/ partition. To make this /, keep the PROM's expectations in mind.

So, continuing on... To create the partitions type fdisk /dev/sda at the prompt. The main commands to know are these:

КОД **List of important fdisk commands**

```
o: Wipe out old partition table, starting with an empty MS-DOS partition table
n: New Partition
t: Change Partition Type
    Use type 82 for Linux Swap, 83 for Linux FS
d: Delete a partition
p: Display (print) Partition Table
q: Quit -- leaving old partition table as is.
w: Quit -- writing partition table in the process.

```

`root #` `fdisk /dev/sda`

```
The number of cylinders for this disk is set to 19870.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
1) software that runs at boot time (e.g., old versions of LILO)
2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)

```

Start by clearing out any existing partitions:

`Command (m for help):` `o`

```
Building a new DOS disklabel. Changes will remain in memory only,
until you decide to write them. After that, of course, the previous
content won't be recoverable.


The number of cylinders for this disk is set to 19870.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
1) software that runs at boot time (e.g., old versions of LILO)
2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)
Warning: invalid flag 0x0000 of partition table 4 will be corrected by w(rite)

```

Now verify the partition table is empty using the `p` command:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

   Device Boot      Start         End      Blocks   Id  System

```

Создайте раздел /boot:

`Command (m for help):` `n`

```
Command action
   e   extended
   p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-19870, default 1):
Last cylinder or +size or +sizeM or +sizeK (1-19870, default 19870): +20M

```

When printing the partitions, notice the newly created one:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          40       20128+  83  Linux

```

Let's now create an extended partition that covers the remainder of the disk. In that extended partition, we'll create the rest (logical partitions):

`Command (m for help):` `n`

```
Command action
   e   extended
   p   primary partition (1-4)
e
Partition number (1-4): 2
First cylinder (41-19870, default 41):
Using default value 41
Last cylinder or +size or +sizeM or +sizeK (41-19870, default 19870):
Using default value 19870

```

Now we create the / partition, /usr, /var, et.

`Command (m for help):` `n`

```
Command action
   l   logical (5 or over)
   p   primary partition (1-4)
l
First cylinder (41-19870, default 41):<Press ENTER>
Using default value 41
Last cylinder or +size or +sizeM or +sizeK (41-19870, default 19870): +500M

```

Repeat this as needed.

Last but not least, the swap space. It is recommended to have at least 250MB swap,
preferrably 1GB:

`Command (m for help):` `n`

```
Command action
   l   logical (5 or over)
   p   primary partition (1-4)
l
First cylinder (17294-19870, default 17294): <Press ENTER>
Using default value 17294
Last cylinder or +size or +sizeM or +sizeK (1011-19870, default 19870): <Press ENTER>
Using default value 19870

```

When checking the partition table, everything should be ready - one thing notwithstanding.

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

Device Boot      Start         End      Blocks      ID  System
/dev/sda1               1          21       10552+  83  Linux
/dev/sda2              22       19870    10003896    5  Extended
/dev/sda5              22        1037      512032+  83  Linux
/dev/sda6            1038        5101     2048224+  83  Linux
/dev/sda7            5102        9165     2048224+  83  Linux
/dev/sda8            9166       13229     2048224+  83  Linux
/dev/sda9           13230       17293     2048224+  83  Linux
/dev/sda10          17294       19870     1298776+  83  Linux

```

Notice how #10, the swap partition is still type 83? Let's change that to the proper type:

`Command (m for help):` `t`

```
Partition number (1-10): 10
Hex code (type L to list codes): 82
Changed system type of partition 10 to 82 (Linux swap)

```

Проверьте:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

Device Boot      Start         End      Blocks      ID  System
/dev/sda1               1          21       10552+  83  Linux
/dev/sda2              22       19870    10003896    5  Extended
/dev/sda5              22        1037      512032+  83  Linux
/dev/sda6            1038        5101     2048224+  83  Linux
/dev/sda7            5102        9165     2048224+  83  Linux
/dev/sda8            9166       13229     2048224+  83  Linux
/dev/sda9           13230       17293     2048224+  83  Linux
/dev/sda10          17294       19870     1298776+  82  Linux Swap

```

We write out the new partition table:

`Command (m for help):` `w`

```
The partition table has been altered!

Calling ioctl() to re-read partition table.
Syncing disks.

```

## Создание файловых систем

**Предупреждение**

При использовании SSD или NVMe диска рекомендуется проверить наличие обновлений для его прошивки. Определённые SSD от Intel (600p и 6000p) нуждаются в обновлении прошивки для [исправления возможной потери данных](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) из-за особенностей использования I/O в XFS. Проблема проявляется на уровне прошивки и не является ошибкой самой файловой системы XFS. Программа smartctl умеет отображать модель устройства и версию прошивки.

### Введение

Теперь, когда разделы созданы, пора разместить на них файловые системы. В следующем разделе описаны различные поддерживаемые в Linux файловые системы. Те из читателей, кто уже знает, какую файловую систему будет использовать, могут продолжить с раздела [Создание файловой системы на разделе](/wiki/Handbook:MIPS/Installation/Disks/ru#Applying_a_filesystem_to_a_partition "Handbook:MIPS/Installation/Disks/ru"). Остальным стоит продолжить чтение, чтобы узнать о доступных вариантах…

### Файловые системы

Linux поддерживает несколько десятков файловых систем, хотя для большинства из них необходимы достаточно веские причины их использовать. Лишь только некоторые из них можно считать стабильными на архитектуре mips. Рекомендуется прочитать информацию о файловых системах и об их состоянии поддержки перед тем, как останавливать свой выбор на экспериментальных. **XFS — рекомендуемая стабильная файловая система общего применения для всех платформ**. Ниже представлен неполный список файловых систем:

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

Например, чтобы отформатировать корневой раздел (/dev/sda5) в xfs при использовании структуры разделов из примера, используются следующие команды:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf /dev/sda5`

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

`root #` `mkswap /dev/sda10`

**Заметка**

Ранее начатый но незаконченный процесс установки может быть продолжен начиная с этого места Руководства. Используйте эту ссылку в качестве постоянной ссылки: [Возобновление установки начинается здесь](/wiki/Handbook:MIPS/Installation/Disks/ru#Resumed_installations_start_here "Handbook:MIPS/Installation/Disks/ru").

Чтобы активировать раздел подкачки, используйте swapon:

`root #` `swapon /dev/sda10`

Шаг «активации» необходим потому, что раздел подкачки был только что создан в живом окружении. Как только система будет перезагружена, и если раздел подкачки правильно определён в fstab или в другом средстве монтирования разделов, он будет активироваться автоматически.

## Монтирование корневого раздела

В некоторых окружениях может не хватать предлагаемой для корневого раздела Gentoo точки монтирования (/mnt/gentoo), либо точек монтирования для дополнительных разделов, созданных в разделе «Создание разделов»:

`root #` `mkdir --parents /mnt/gentoo`

По необходимости создайте с помощью команды mkdirдополнительные точки монтирования для других разделов, созданных в предыдущих шагах установки.

После создания точек монтирования настало время сделать их доступными через команду mount.

Смонтируем корневой раздел:

`root #` `mount /dev/sda5 /mnt/gentoo`

При необходимости смонтируйте дополнительные разделы с помощью команды mount.

**Заметка**

Если /tmp/ находится на отдельном разделе, не забудьте после монтирования изменить права доступа:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Это также справедливо для /var/tmp.

Позже в инструкции будут смонтированы файловая система proc (виртуальный интерфейс к ядру) и другие псевдофайловые системы ядра. Но сначала необходимо распаковать [файл стадии Gentoo](/wiki/Handbook:MIPS/Installation/Stage/ru "Handbook:MIPS/Installation/Stage/ru").

[← Настройка сети](/wiki/Handbook:MIPS/Installation/Networking/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:MIPS/ru "Handbook:MIPS/ru") [Установка Gentoo файлов установки →](/wiki/Handbook:MIPS/Installation/Stage/ru "Следующая часть")