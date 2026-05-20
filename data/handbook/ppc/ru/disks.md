# Disks

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Disks/de "Handbuch:PPC/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Disks "Handbook:PPC/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Disks/es "Manual de Gentoo: PPC/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Disks/fr "Handbook:PPC/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Disks/it "Handbook:PPC/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Disks/hu "Handbook:PPC/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Disks/pl "Handbook:PPC/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Disks/pt-br "Handbook:PPC/Installation/Disks/pt-br (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:PPC/Installation/Disks/ta "கையேடு:PPC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Disks/zh-cn "手册：PPC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Disks/ja "ハンドブック:PPC/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Disks/ko "Handbook:PPC/Installation/Disks/ko (100% translated)")

[PPC Handbook](/wiki/Handbook:PPC/ru "Handbook:PPC/ru")[Установка](/wiki/Handbook:PPC/Full/Installation/ru "Handbook:PPC/Full/Installation/ru")[Об установке](/wiki/Handbook:PPC/Installation/About/ru "Handbook:PPC/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:PPC/Installation/Media/ru "Handbook:PPC/Installation/Media/ru")[Настройка сети](/wiki/Handbook:PPC/Installation/Networking/ru "Handbook:PPC/Installation/Networking/ru")Подготовка дисков[Установка файла stage](/wiki/Handbook:PPC/Installation/Stage/ru "Handbook:PPC/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:PPC/Installation/Base/ru "Handbook:PPC/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:PPC/Installation/Kernel/ru "Handbook:PPC/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:PPC/Installation/System/ru "Handbook:PPC/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:PPC/Installation/Tools/ru "Handbook:PPC/Installation/Tools/ru")[Настройка загрузчика](/wiki/Handbook:PPC/Installation/Bootloader/ru "Handbook:PPC/Installation/Bootloader/ru")[Завершение](/wiki/Handbook:PPC/Installation/Finalizing/ru "Handbook:PPC/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:PPC/Full/Working/ru "Handbook:PPC/Full/Working/ru")[Введение в Portage](/wiki/Handbook:PPC/Working/Portage/ru "Handbook:PPC/Working/Portage/ru")[USE-флаги](/wiki/Handbook:PPC/Working/USE/ru "Handbook:PPC/Working/USE/ru")[Возможности Portage](/wiki/Handbook:PPC/Working/Features/ru "Handbook:PPC/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:PPC/Working/Initscripts/ru "Handbook:PPC/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:PPC/Working/EnvVar/ru "Handbook:PPC/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:PPC/Full/Portage/ru "Handbook:PPC/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:PPC/Portage/Files/ru "Handbook:PPC/Portage/Files/ru")[Переменные](/wiki/Handbook:PPC/Portage/Variables/ru "Handbook:PPC/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:PPC/Portage/Branches/ru "Handbook:PPC/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:PPC/Portage/Tools/ru "Handbook:PPC/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:PPC/Portage/CustomTree/ru "Handbook:PPC/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:PPC/Portage/Advanced/ru "Handbook:PPC/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:PPC/Full/Networking/ru "Handbook:PPC/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:PPC/Networking/Introduction/ru "Handbook:PPC/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:PPC/Networking/Advanced/ru "Handbook:PPC/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:PPC/Networking/Modular/ru "Handbook:PPC/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:PPC/Networking/Wireless/ru "Handbook:PPC/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:PPC/Networking/Extending/ru "Handbook:PPC/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:PPC/Networking/Dynamic/ru "Handbook:PPC/Networking/Dynamic/ru")

## Contents

- [1Введение в блочные устройства](#.D0.92.D0.B2.D0.B5.D0.B4.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B2_.D0.B1.D0.BB.D0.BE.D1.87.D0.BD.D1.8B.D0.B5_.D1.83.D1.81.D1.82.D1.80.D0.BE.D0.B9.D1.81.D1.82.D0.B2.D0.B0)
  - [1.1Блочные устройства](#.D0.91.D0.BB.D0.BE.D1.87.D0.BD.D1.8B.D0.B5_.D1.83.D1.81.D1.82.D1.80.D0.BE.D0.B9.D1.81.D1.82.D0.B2.D0.B0)
  - [1.2Разделы](#.D0.A0.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D1.8B)
- [2Разработка схемы разделов](#.D0.A0.D0.B0.D0.B7.D1.80.D0.B0.D0.B1.D0.BE.D1.82.D0.BA.D0.B0_.D1.81.D1.85.D0.B5.D0.BC.D1.8B_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2)
  - [2.1Сколько разделов и насколько больших?](#.D0.A1.D0.BA.D0.BE.D0.BB.D1.8C.D0.BA.D0.BE_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_.D0.B8_.D0.BD.D0.B0.D1.81.D0.BA.D0.BE.D0.BB.D1.8C.D0.BA.D0.BE_.D0.B1.D0.BE.D0.BB.D1.8C.D1.88.D0.B8.D1.85.3F)
  - [2.2Что по поводу пространства подкачки?](#.D0.A7.D1.82.D0.BE_.D0.BF.D0.BE_.D0.BF.D0.BE.D0.B2.D0.BE.D0.B4.D1.83_.D0.BF.D1.80.D0.BE.D1.81.D1.82.D1.80.D0.B0.D0.BD.D1.81.D1.82.D0.B2.D0.B0_.D0.BF.D0.BE.D0.B4.D0.BA.D0.B0.D1.87.D0.BA.D0.B8.3F)
  - [2.3Apple New World](#Apple_New_World)
  - [2.4Apple Old World](#Apple_Old_World)
  - [2.5Pegasos](#Pegasos)
  - [2.6IBM PReP (RS/6000)](#IBM_PReP_.28RS.2F6000.29)
- [3Использование mac-fdisk (Apple)](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_mac-fdisk_.28Apple.29)
- [4Использование parted (Pegasos и RS/6000)](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_parted_.28Pegasos_.D0.B8_RS.2F6000.29)
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

### Разделы

Несмотря на то, что теоретически возможно использовать весь диск для размещения системы Linux, это почти никогда не делается на практике. Вместо этого, блочное устройство разбивается на меньшие, более управляемые блочные устройства. Во многих системах они называются разделами.

**Заметка**

In the remainder of the installation instructions, we will use the Pegasos example partition layout. Adjust to personal preference.

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

### Apple New World

Apple New World machines are fairly straightforward to configure. The first partition is always an Apple Partition Map (APM). This partition keeps track of the layout of the disk. It is not possible to remove this partition. The next partition should always be a bootstrap partition. This partition contains a small (800KiB) HFS filesystem that holds a copy of the bootloader Yaboot and its configuration file. This partition is not the same as a /boot partition as found on other architectures. After the boot partition, the usual Linux filesystems are placed, according to the scheme below. The swap partition is a temporary storage place for when the system runs out of physical memory. The root partition will contain the filesystem that Gentoo is installed on. To dual boot, the OSX partition can go anywhere after the bootstrap partition to insure that yaboot starts first.

**Заметка**

There may be "Disk Driver" partitions on the disk such as Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKit, and Apple\_Patches. These are used to boot MacOS, so if there is no need for this, they can be removed by initializing the disk with mac-fdisk's `i` option. Be careful, this will completely erase the disk! If in doubt do not remove them.

**Заметка**

If the disk is partitioned with Apple's Disk Utility, there may be 128 MiB spaces between partitions which Apple reserves for "future use". These can be safely removed.

Раздел
Размер
Файловая система
Описание
/dev/sda132 КиБ
Нет.
Apple Partition Map (APM).
/dev/sda2800 КиБ
HFS
Apple bootstrap.
/dev/sda3512 МиБ
swap
Раздел подкачки Linux (тип 0x82).
/dev/sda4Оставшееся пространство диска.
ext4, xfs и т.д.
Корневой раздел Linux.

### Apple Old World

Apple Old World machines are a bit more complicated to configure. The first partition is always an Apple Partition Map (APM). This partition keeps track of the layout of the disk. It is not possible to remove this partition. When using BootX, the configuration below assumes that MacOS is installed on a separate disk. If this is not the case, there will be additional partitions for "Apple Disk Drivers" such as Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKit, Apple\_Patches and the MacOS install. When using Quik, it is necessary to create a boot partition to hold the kernel, unlike other Apple boot methods. After the boot partition, the usual Linux filesystems are placed, according to the scheme below. The swap partition is a temporary storage place for when the system runs out of physical memory. The root partition will contain the filesystem that Gentoo is installed on.

**Заметка**

When using an Old World machine, it is necessary to keep MacOS available. The layout here assumes MacOS is installed on a separate drive.

Пример схемы разбивки диска для систем Old World:

Раздел
Размер
Файловая система
Описание
/dev/sda132 КиБ
Нет.
Apple Partition Map (APM).
/dev/sda232 КиБ
ext2
Загрузочный раздел Quik (только для quik).
/dev/sda3512 МиБ
swap
Раздел подкачки Linux (тип 0x82).
/dev/sda4Оставшееся пространство диска.
ext4, xfs и т.д.
Корневой раздел Linux.

### Pegasos

The Pegasos partition layout is quite simple compared to the Apple layouts. The first partition is a boot partition, which contains kernels to be booted along with an Open Firmware script that presents a menu on boot. After the boot partition, the usual Linux filesystems are placed, according to the scheme below. The swap partition is a temporary storage place for when the system runs out of physical memory. The root partition will contain the filesystem that Gentoo is installed on.

Пример схемы разбивки диска для систем Pegasos:

Раздел
Размер
Файловая система
Описание
/dev/sda132 МиБ
affs1 или ext2
Загрузочный раздел.
/dev/sda2512 МиБ
swap
Раздел подкачки Linux (тип 0x82).
/dev/sda3Оставшееся пространство диска.
ext4, xfs и т.д.
Корневой раздел Linux.

### IBM PReP (RS/6000)

The IBM PowerPC Reference Platform (PReP) requires a small PReP boot partition on the disk's first partition, followed by the swap and root partitions.

Пример схемы разбивки диска для систем IBM PReP:

Раздел
Размер
Файловая система
Описание
/dev/sda1800 КиБ
нет
Загрузочный раздел PReP (тип 0x41).
/dev/sda2512 МиБ
swap
Раздел подкачки Linux (тип 0x82).
/dev/sda3Оставшееся пространство диска.
ext4, xfs и т.д.
Корневой раздел Linux (тип 0x83).

**Предупреждение**

parted is able to resize partitions including HFS+. Unfortunately there may be issues with resizing HFS+ journaled filesystems, so, for the best results, switch off journaling in Mac OS X before resizing. Remember that any resizing operation is dangerous, so attempt at own risk! Be sure to always have a backup of all data before resizing!

## Использование mac-fdisk (Apple)

Для начала создайте разделы с помощью mac-fdisk:

`root #` `mac-fdisk /dev/sda`

If Apple's Disk Utility was used prior to leave space for Linux, first delete the partitions that might have been created previously to make room for the new install. Use `d` in mac-fdisk to delete those partition(s). It will ask for the partition number to delete. Usually the first partition on NewWorld machines (Apple\_partition\_map) cannot be deleted. To start with a clean disk, simply initialize the disk by pressing `i`. This will completely erase the disk, so use this with caution.

Second, create an Apple\_Bootstrap partition by using `b`. It will ask for what block to start. Enter the number of the first free partition, followed by a `p`. For instance this is _2p_.

**Заметка**

This partition is not a /boot partition. It is not used by Linux at all; there is no need to place any filesystem on it and it should never be mounted. Apple users don't need an extra partition for /boot.

Now create a swap partition by pressing `c`. Again mac-fdisk will ask for what block to start this partition from. As we used 2 before to create the Apple\_Bootstrap partition, now enter _3p_. When sked for the size, enter 512M (or whatever size needed - a minimum of 512MiB is recommended, but 2 times the physical memory is the generally accepted size). When asked for a name, enter _swap_.

To create the root partition, enter `c`, followed by _4p_ to select from what block the root partition should start. When asked for the size, enter _4p_ again. mac-fdisk will interpret this as "Use all available space". When asked for the name, enter _root_.

To finish up, write the partition to the disk using `w` and `q` to quit mac-fdisk.

**Заметка**

To make sure everything is okay, run mac-fdisk -l and check whether all the partitions are there. If not all partitions created previously are shown, or the changes made are not reflected in the output, reinitialize the partitions by pressing `i` in mac-fdisk. Note that this will recreate the partition map and thus remove all existing partitions.

## Использование parted (Pegasos и RS/6000)

parted, the partition editor, can now handle HFS+ partitions used by Mac OS and Mac OS X. With this tool it is possible to resize the Mac partitions and create space for the Linux partitions. Nevertheless, the example below describes partitioning for Pegasos machines only.

Чтобы продолжить, давайте запустим parted:

`root #` `parted /dev/sda`

If the drive is unpartitioned, run mklabel amiga to create a new disklabel for the drive.

It is possible to type print at any time in parted to display the current partition table. To abort parted, press `Ctrl` + `c`.

If next to Linux, the system is also meant to have MorphOS installed, then create an affs1 filesystem at the start of the drive. 32MB should be more than enough to store the MorphOS kernel. With a Pegasos I, or when Linux will use any filesystem besides ext2 or ext3, then it is necessary to also store the Linux kernel on this partition (the Pegasos II can only boot from ext2/ext3 or affs1 partitions). To create the partition run `mkpart primary affs1 START END` where START and END should be replaced with the megabyte range (e.g. 0 32) which creates a 32 MB partition starting at 0MB and ending at 32MB. When creating an ext2 or ext3 partition instead, substitute ext2 or ext3 for affs1 in the mkpart command.

Create two partitions for Linux, one root filesystem and one swap partition. Run `mkpart primary START END` to create each partition, replacing START and END with the desired megabyte boundaries.

It is generally recommended to create a swap partition that is two times bigger than the amount of RAM in the computer, but at least 512MiB is recommended. To create the swap partition, run `mkpart primary linux-swap START END` with START and END again denoting the partition boundaries.

Просто введите `quit`, когда закончите с parted.

## Создание файловых систем

**Предупреждение**

При использовании SSD или NVMe диска рекомендуется проверить наличие обновлений для его прошивки. Определённые SSD от Intel (600p и 6000p) нуждаются в обновлении прошивки для [исправления возможной потери данных](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) из-за особенностей использования I/O в XFS. Проблема проявляется на уровне прошивки и не является ошибкой самой файловой системы XFS. Программа smartctl умеет отображать модель устройства и версию прошивки.

### Введение

Теперь, когда разделы созданы, пора разместить на них файловые системы. В следующем разделе описаны различные поддерживаемые в Linux файловые системы. Те из читателей, кто уже знает, какую файловую систему будет использовать, могут продолжить с раздела [Создание файловой системы на разделе](/wiki/Handbook:PPC/Installation/Disks/ru#Applying_a_filesystem_to_a_partition "Handbook:PPC/Installation/Disks/ru"). Остальным стоит продолжить чтение, чтобы узнать о доступных вариантах…

### Файловые системы

Linux поддерживает несколько десятков файловых систем, хотя для большинства из них необходимы достаточно веские причины их использовать. Лишь только некоторые из них можно считать стабильными на архитектуре ppc. Рекомендуется прочитать информацию о файловых системах и об их состоянии поддержки перед тем, как останавливать свой выбор на экспериментальных. **XFS — рекомендуемая стабильная файловая система общего применения для всех платформ**. Ниже представлен неполный список файловых систем:

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

Ранее начатый но незаконченный процесс установки может быть продолжен начиная с этого места Руководства. Используйте эту ссылку в качестве постоянной ссылки: [Возобновление установки начинается здесь](/wiki/Handbook:PPC/Installation/Disks/ru#Resumed_installations_start_here "Handbook:PPC/Installation/Disks/ru").

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

Позже в инструкции будут смонтированы файловая система proc (виртуальный интерфейс к ядру) и другие псевдофайловые системы ядра. Но сначала необходимо распаковать [файл стадии Gentoo](/wiki/Handbook:PPC/Installation/Stage/ru "Handbook:PPC/Installation/Stage/ru").

[← Настройка сети](/wiki/Handbook:PPC/Installation/Networking/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:PPC/ru "Handbook:PPC/ru") [Установка Gentoo файлов установки →](/wiki/Handbook:PPC/Installation/Stage/ru "Следующая часть")