# System

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/System/de "Handbook:AMD64/Installation/System/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/System "Handbook:AMD64/Installation/System (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/System/es "Handbook:AMD64/Installation/System/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/System/fr "Handbook:AMD64/Installation/System/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/System/it "Handbook:AMD64/Installation/System/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/System/hu "Handbook:AMD64/Installation/System/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/System/pl "Handbook:AMD64/Installation/System/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/System/pt-br "Handbook:AMD64/Installation/System/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/System/cs "Handbook:AMD64/Installation/System/cs (100% translated)")
- русский
- [தமிழ்](/wiki/Handbook:AMD64/Installation/System/ta "Handbook:AMD64/Installation/System/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/System/zh-cn "Handbook:AMD64/Installation/System/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/System/ja "Handbook:AMD64/Installation/System/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/System/ko "Handbook:AMD64/Installation/System/ko (100% translated)")

[AMD64 Handbook](/wiki/Handbook:AMD64/ru "Handbook:AMD64/ru")[Установка](/wiki/Handbook:AMD64/Full/Installation/ru "Handbook:AMD64/Full/Installation/ru")[Об установке](/wiki/Handbook:AMD64/Installation/About/ru "Handbook:AMD64/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:AMD64/Installation/Media/ru "Handbook:AMD64/Installation/Media/ru")[Настройка сети](/wiki/Handbook:AMD64/Installation/Networking/ru "Handbook:AMD64/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:AMD64/Installation/Disks/ru "Handbook:AMD64/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:AMD64/Installation/Stage/ru "Handbook:AMD64/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:AMD64/Installation/Base/ru "Handbook:AMD64/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:AMD64/Installation/Kernel/ru "Handbook:AMD64/Installation/Kernel/ru")Настройка системы[Установка системных утилит](/wiki/Handbook:AMD64/Installation/Tools/ru "Handbook:AMD64/Installation/Tools/ru")[Настройка загрузчика](/wiki/Handbook:AMD64/Installation/Bootloader/ru "Handbook:AMD64/Installation/Bootloader/ru")[Завершение](/wiki/Handbook:AMD64/Installation/Finalizing/ru "Handbook:AMD64/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:AMD64/Full/Working/ru "Handbook:AMD64/Full/Working/ru")[Введение в Portage](/wiki/Handbook:AMD64/Working/Portage/ru "Handbook:AMD64/Working/Portage/ru")[USE-флаги](/wiki/Handbook:AMD64/Working/USE/ru "Handbook:AMD64/Working/USE/ru")[Возможности Portage](/wiki/Handbook:AMD64/Working/Features/ru "Handbook:AMD64/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:AMD64/Working/Initscripts/ru "Handbook:AMD64/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:AMD64/Working/EnvVar/ru "Handbook:AMD64/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:AMD64/Full/Portage/ru "Handbook:AMD64/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:AMD64/Portage/Files/ru "Handbook:AMD64/Portage/Files/ru")[Переменные](/wiki/Handbook:AMD64/Portage/Variables/ru "Handbook:AMD64/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:AMD64/Portage/Branches/ru "Handbook:AMD64/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:AMD64/Portage/Tools/ru "Handbook:AMD64/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:AMD64/Portage/CustomTree/ru "Handbook:AMD64/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:AMD64/Portage/Advanced/ru "Handbook:AMD64/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:AMD64/Full/Networking/ru "Handbook:AMD64/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:AMD64/Networking/Introduction/ru "Handbook:AMD64/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:AMD64/Networking/Advanced/ru "Handbook:AMD64/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:AMD64/Networking/Modular/ru "Handbook:AMD64/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:AMD64/Networking/Wireless/ru "Handbook:AMD64/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:AMD64/Networking/Extending/ru "Handbook:AMD64/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:AMD64/Networking/Dynamic/ru "Handbook:AMD64/Networking/Dynamic/ru")

## Contents

- [1Информация о файловой системе](#.D0.98.D0.BD.D1.84.D0.BE.D1.80.D0.BC.D0.B0.D1.86.D0.B8.D1.8F_.D0.BE_.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D0.BE.D0.B9_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D0.B5)
  - [1.1Метки файловых систем и UUID](#.D0.9C.D0.B5.D1.82.D0.BA.D0.B8_.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D1.8B.D1.85_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC_.D0.B8_UUID)
  - [1.2Метки разделов и UUID](#.D0.9C.D0.B5.D1.82.D0.BA.D0.B8_.D1.80.D0.B0.D0.B7.D0.B4.D0.B5.D0.BB.D0.BE.D0.B2_.D0.B8_UUID)
  - [1.3О файле fstab](#.D0.9E_.D1.84.D0.B0.D0.B9.D0.BB.D0.B5_fstab)
  - [1.4Создание файла fstab](#.D0.A1.D0.BE.D0.B7.D0.B4.D0.B0.D0.BD.D0.B8.D0.B5_.D1.84.D0.B0.D0.B9.D0.BB.D0.B0_fstab)
    - [1.4.1Системы DOS/Legacy BIOS](#.D0.A1.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B_DOS.2FLegacy_BIOS)
    - [1.4.2Системы UEFI](#.D0.A1.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B_UEFI)
    - [1.4.3DPS UEFI PARTUUID](#DPS_UEFI_PARTUUID)
- [2Информация о сети](#.D0.98.D0.BD.D1.84.D0.BE.D1.80.D0.BC.D0.B0.D1.86.D0.B8.D1.8F_.D0.BE_.D1.81.D0.B5.D1.82.D0.B8)
  - [2.1Имя хоста](#.D0.98.D0.BC.D1.8F_.D1.85.D0.BE.D1.81.D1.82.D0.B0)
    - [2.1.1Установка имени хоста (OpenRC или systemd)](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_.D0.B8.D0.BC.D0.B5.D0.BD.D0.B8_.D1.85.D0.BE.D1.81.D1.82.D0.B0_.28OpenRC_.D0.B8.D0.BB.D0.B8_systemd.29)
    - [2.1.2systemd](#systemd)
  - [2.2Сеть](#.D0.A1.D0.B5.D1.82.D1.8C)
    - [2.2.1DHCP через dhcpcd (любая система инициализации)](#DHCP_.D1.87.D0.B5.D1.80.D0.B5.D0.B7_dhcpcd_.28.D0.BB.D1.8E.D0.B1.D0.B0.D1.8F_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D0.B0_.D0.B8.D0.BD.D0.B8.D1.86.D0.B8.D0.B0.D0.BB.D0.B8.D0.B7.D0.B0.D1.86.D0.B8.D0.B8.29)
    - [2.2.2netifrc (OpenRC)](#netifrc_.28OpenRC.29)
      - [2.2.2.1Настройка сети](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D1.81.D0.B5.D1.82.D0.B8)
      - [2.2.2.2Автоматический запуск сетевого подключения при загрузке системы](#.D0.90.D0.B2.D1.82.D0.BE.D0.BC.D0.B0.D1.82.D0.B8.D1.87.D0.B5.D1.81.D0.BA.D0.B8.D0.B9_.D0.B7.D0.B0.D0.BF.D1.83.D1.81.D0.BA_.D1.81.D0.B5.D1.82.D0.B5.D0.B2.D0.BE.D0.B3.D0.BE_.D0.BF.D0.BE.D0.B4.D0.BA.D0.BB.D1.8E.D1.87.D0.B5.D0.BD.D0.B8.D1.8F_.D0.BF.D1.80.D0.B8_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BA.D0.B5_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B)
  - [2.3Файл hosts](#.D0.A4.D0.B0.D0.B9.D0.BB_hosts)
- [3Системная информация](#.D0.A1.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D0.BD.D0.B0.D1.8F_.D0.B8.D0.BD.D1.84.D0.BE.D1.80.D0.BC.D0.B0.D1.86.D0.B8.D1.8F)
  - [3.1Пароль суперпользователя](#.D0.9F.D0.B0.D1.80.D0.BE.D0.BB.D1.8C_.D1.81.D1.83.D0.BF.D0.B5.D1.80.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D1.82.D0.B5.D0.BB.D1.8F)
  - [3.2Настройка инициализации и загрузки](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D0.B8.D0.BD.D0.B8.D1.86.D0.B8.D0.B0.D0.BB.D0.B8.D0.B7.D0.B0.D1.86.D0.B8.D0.B8_.D0.B8_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BA.D0.B8)
    - [3.2.1OpenRC](#OpenRC)
    - [3.2.2systemd](#systemd_2)

## Информация о файловой системе

### Метки файловых систем и UUID

И MBR (BIOS), и GPT поддерживают как метки (labels), так и UUID _файловой системы_. Эти свойства могут быть определены в /etc/fstab в качестве альтернативы для команды mount для определения блочного устройства. Такие свойства используются при попытке найти и примонтировать блочные устройства. Метки и UUID файловой системы определяются через префиксы LABEL и UUID. Их можно посмотреть командой blkid:

`root #` `blkid`

**Предупреждение**

Если файловая система внутри раздела будет полностью затёрта (wipe), то значение меток и UUID файловой системы также будут изменены или удалены.

Для обеспечения уникальности читателям, использующим таблицу разделов в стиле MBR, рекомендуется использовать UUID вместо меток для определения монтируемых томов в /etc/fstab.

**Важно**

UUID файловых систем на разделе LVM и соотносящихся LVM снапшотах одинаковы, поэтому лучше избегать использование UUID для монтирования разделов LVM.

### Метки разделов и UUID

Для систем с поддержкой меток дисков GPT есть несколько более надёжных вариантов для определения разделов в /etc/fstab. Метки разделов и UUID разделов могут быть использованы для идентификации разделов блочного устройства, независимо от того, какая файловая система была выбрана для самого раздела. Метки и UUID раздела определяются через префиксы PARTLABEL и/или PARTUUID. Их можно увидеть в терминале с помощью команды blkid.

Вывод для системы **amd64** EFI, использующей UUID с Discoverable Partition Specification, может быть следующим:

`root #` `blkid`

```
/dev/sr0: BLOCK_SIZE="2048" UUID="2023-08-28-03-54-40-00" LABEL="ISOIMAGE" TYPE="iso9660" PTTYPE="PMBR"
/dev/loop0: TYPE="squashfs"
/dev/sda2: PARTUUID="0657fd6d-a4ab-43c4-84e5-0933c84b4f4f"
/dev/sda3: PARTUUID="1cdf763a-5b4c-4dbf-99db-a056c504e8b2"
/dev/sda1: PARTUUID="c12a7328-f81f-11d2-ba4b-00a0c93ec93b"

```

Хотя это не всегда верно для меток разделов, использование UUID для идентификации раздела в fstab обеспечивает гарантию того, что загрузчик не собьётся при поиске определённого тома, даже если _файловая система_ будет изменена или перезаписана в будущем. Использование по умолчанию старых файлов блочных устройств (/dev/sd\*N) для определения разделов в fstab будет рискованно в системах, в которых регулярно добавляются и удаляются блочные устройства SATA.

Именование файлов блочных устройств зависит от ряда факторов, включая то, как и в каком порядке диски подключены в системе. Они могут отображаться в другом порядке, в зависимости от того, какое из устройств обнаруживается ядром первым в начале загрузки. При этом, если системный администратор не намерен постоянно переключать жесткие диски, использование файлов блочных устройств по умолчанию является простым и удобным подходом.

### О файле fstab

В Linux все разделы, используемые системой, должны быть записаны в файле [/etc/fstab](/wiki//etc/fstab/ru "/etc/fstab/ru"). Этот файл содержит информацию о точках монтирования разделов (где они должны быть видны в структуре файловой системы), как они должны быть подключены, а также специальные параметры (автоматическое подключение или нет, может ли пользователь их подключать или нет и так далее).

### Создание файла fstab

**Заметка**

Если в качестве системы инициализации используется systemd, UUID разделов соответствуют Discoverable Partition Specification (как это представлено в разделе [Подготовка дисков](/wiki/Handbook:AMD64/Installation/Disks/ru "Handbook:AMD64/Installation/Disks/ru")) и система использует UEFI, то создание fstab можно пропустить, так как systemd автоматически монтирует все разделы, соответствующие данной спецификации.

В файле /etc/fstab используется синтаксис, напоминающий таблицу. Каждая строка состоит из шести полей, которые разделены пропусками (пробелами, отступами или смесь этого). Каждое поле имеет своё значение:

1. Первое поле содержит блочное устройство (или удалённую файловую систему), которое следует примонтировать. Для экземпляров блочных устройств возможно использование различных идентификаторов, включая путь к устройству, метки файловой системы, метки раздела и UUID
2. Второе поле содержит точку монтирования, к которой следует монтировать раздел.
3. Третье поле содержит тип файловой системы, используемой разделом.
4. Четвёртое поле содержит параметры, используемые командой mount во время монтирования. Так как у каждой файловой системы могут быть собственные уникальные параметры, рекомендуется прочитать man-страницу команды mount (man mount), чтобы получить полный список всех возможных параметров. Параметры монтирования разделяются запятыми.
5. Пятое поле используется командой dump для определения того, нуждается ли раздел в дампе или нет. Обычно это поле содержит `0` (ноль).
6. Шестое поле используется командой fsck для определения порядка проведения проверки ошибок файловой системы, если система была отключена некорректно. Для корневой файловой системы необходимо указывать `1`, для остальных — `2` (или `0`, если проверка не требуется вовсе).

**Важно**

Файл /etc/fstab, который предоставляется в файлах stage Gentoo по умолчанию, _не является_ валидным файлом fstab, а представлен в качестве шаблона, который может быть использован для ввода актуальных значений.

`root #` `nano /etc/fstab`

#### Системы DOS/Legacy BIOS

Давайте посмотрим, как записать настройки для /boot раздела. Это просто пример, поэтому запись необходимо изменить в соответствии с ранее выбранной схемой разделов.
В **amd64** примере, /boot является обычным /dev/sda1 разделом с рекомендуемой файловой системой xfs. Необходимо проверять его во время загрузки, поэтому мы запишем следующее:

ФАЙЛ **`/etc/fstab`** **Пример строки DOS/Legacy BIOS для /etc/fstab**

```
# Исправьте все различия в форматировании и/или других разделов, созданных на этапе "Подготовка дисков".
/dev/sda1   /boot     xfs    defaults        0 2

```

Из соображения безопасности некоторые системные администраторы могут не захотеть автоматически монтировать раздел /boot. Для этого следует заменить `defaults` на `noauto`. Это будет означать, что раздел придётся монтировать каждый раз, когда понадобится его использовать.

Добавьте правила, которые соответствуют ранее запланированной схеме разметки диска, а также правила для таких устройств, как компакт-диски, и других устройств (если они есть в системе).

Ниже приведён более подробный пример файла /etc/fstab:

ФАЙЛ **`/etc/fstab`** **Полный пример /etc/fstab для систем DOS/Legacy BIOS**

```
# Исправьте все различия в форматировании и/или добавьте дополнительные разделы, созданные на этапе «Подготовка дисков».
/dev/sda1   /boot        xfs    defaults    0 2
/dev/sda2   none         swap    sw                   0 0
/dev/sda3   /            xfs    defaults,noatime              0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

#### Системы UEFI

Ниже представлен пример файла /etc/fstab для системы, загружающейся через прошивку UEFI:

ФАЙЛ **`/etc/fstab`** **Полный пример /etc/fstab для систем UEFI**

```
# Исправьте все различия в форматировании и/или добавьте дополнительные разделы, созданные на этапе «Подготовка дисков».
/dev/sda1   /efi        vfat    umask=0077,tz=UTC     0 2
/dev/sda2   none         swap    sw                   0 0
/dev/sda3   /            xfs    defaults,noatime              0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

#### DPS UEFI PARTUUID

Ниже представлен пример файла /etc/fstab для диска, отформатированного с дисковыми меткам GPT и набором Discoverable Partition Specification (DPS) UUID-ов, назначенных для прошивки UEFI:

ФАЙЛ **`/etc/fstab`** **Пример DPS PARTUUID fstab**

```
# Скорректируйте с учетом отличий форматирования и дополнительных разделов, созданных на шаге «Подготовка дисков».
# В данном примере демонстрируются дисковые метки GPT с набором Discoverable Partition Specification (DPS) UUID:
PARTUUID=c12a7328-f81f-11d2-ba4b-00a0c93ec93b   /efi        vfat    umask=0077,tz=UTC            0 2
PARTUUID=0657fd6d-a4ab-43c4-84e5-0933c84b4f4f   none        swap    sw                           0 0
PARTUUID=4f68bce3-e8cd-4db1-96e7-fbcaf984b709   /           xfs     defaults,noatime             0 1

```

При использовании `auto` в третьем поле команда mount попытается автоматически определить тип файловой системы. Это рекомендуется для отсоединяемых устройств, которые могут использовать разные файловые системы. Параметр `user` в четвертом поле позволяет монтировать компакт-диски обычными пользователями.

Чтобы улучшить производительность, большинству пользователей может понадобиться добавить параметр монтирования `noatime`, который в итоге приводит более быстрому отклику системы, так как при этом не регистрируются отметки о доступе к файлам (что не всегда необходимо). Этот параметр также рекомендуется использовать в системах с твёрдотельными накопителями (SSD). Также вместо этого параметра можно использовать `lazytime`.

**Совет**

Из–за ухудшения производительности, не рекомендуется устанавливать опцию монтирования `discard` в /etc/fstab. Вместо этого планируйте удаление блоков на периодической основе с помощью планировщика заданий, такого, как cron, или с помощью таймера (systemd). Смотрите статью [Periodic fstrim jobs](/wiki/SSD#Periodic_fstrim_jobs "SSD") для более подробной информации.

Дважды проверьте файл /etc/fstab, сохраните его и выйдите из редактора, чтобы продолжить дальше.

## Информация о сети

Важно учесть, что следующие разделы приведены для того, чтобы помочь читателю быстро настроить свою систему для использования в локальной сети.

Для систем с OpenRC, более подробное руководство по настройке сети доступно в разделе [Расширенная настройка сети](/wiki/Handbook:AMD64/Networking/Introduction/ru "Handbook:AMD64/Networking/Introduction/ru"), который находится ближе к концу руководства. Системы с более специфическими сетевыми потребностями могут пропустить этот раздел, а затем вернуться сюда, чтобы продолжить установку.

Для более конкретной настройки сети systemd, пожалуйста, обратитесь к разделу [Сеть](/wiki/Systemd/ru#Network "Systemd/ru") статьи [systemd](/wiki/Systemd/ru "Systemd/ru").

### Имя хоста

Первое решение, которое предстоит принять администратору системы, это как назвать его/её компьютер. Кажется, что это является довольно лёгким решением, но многие пользователи испытывают трудности с поиском подходящего имени для своего компьютера. Чтобы не мешкать слишком долго, выберите любое имя — его можно будет сменить позже. Например, в приведённом ниже примере используется имя хоста _tux_.

#### Установка имени хоста (OpenRC или systemd)

`root #` `echo tux > /etc/hostname`

#### systemd

Чтобы установить имя хоста для системы с _уже запущенным_ systemd, можно воспользоваться утилитой hostnamectl. Однако, во время установки вместо неё следует использовать команду [systemd-firstboot](/wiki/Handbook:AMD64/Installation/System/ru#Init_and_boot_configuration_systemd "Handbook:AMD64/Installation/System/ru") (см. ниже).

Например, чтобы установить имя хоста на «tux», необходимо запустить:

`root #` `hostnamectl hostname tux`

Просмотрите помощь, используя команду hostnamectl --help или man 1 hostnamectl.

### Сеть

Существует _много_ способов настройки сетевых интерфейсов. Этот раздел покрывает только некоторые из них. Выберите тот, которые кажется более подходящим для вашей установки.

#### DHCP через dhcpcd (любая система инициализации)

В большинстве локальных сетей работает сервер DHCP. В этом случае для получения IP-адреса рекомендуется использовать программу dhcpcd.

Чтобы установить:

`root #` `emerge --ask net-misc/dhcpcd`

Чтобы включить и затем запустить сервис на системах с OpenRC:

`root #` `rc-update add dhcpcd default
`

`root #` `rc-service dhcpcd start
`

Чтобы включить сервис на системах с systemd:

`root #` `systemctl enable dhcpcd`

После выполнения этих шагов при следующей загрузке системы dhcpcd должен получить IP-адрес от DHCP-сервера. Подробнее см. в статье [Dhcpcd](/wiki/Dhcpcd/ru "Dhcpcd/ru").

#### netifrc (OpenRC)

**Совет**

Это один конкретный способ настройки сети, используя [Netifrc](/wiki/Netifrc "Netifrc") с OpenRC. Существуют и другие способы для более простой настройки, такие как [Dhcpcd](/wiki/Dhcpcd/ru "Dhcpcd/ru").

##### Настройка сети

Во время установки Gentoo Linux сеть была уже настроена, однако она была настроена для самого установочного окружения, а не для системы. Сейчас мы устраним это упущение.

**Заметка**

Больше информации о настройке сети, в том числе об объединении интерфейсов, создании мостов, настройке 802.1Q VLAN и беспроводной сети, рассматриваются в разделе [Настройка сети](/wiki/Handbook:AMD64/Networking/Introduction/ru "Handbook:AMD64/Networking/Introduction/ru").

Все настройки сети собраны в файле /etc/conf.d/net. В нём используется простой, но, возможно, пока ещё непонятный синтаксис. Не беспокойтесь! Обо всём мы расскажем далее. Полностью документированные примеры, описывающие множество разных конфигураций, доступны в /usr/share/doc/netifrc-\*/net.example.bz2.

Сначала установите [net-misc/netifrc](https://packages.gentoo.org/packages/net-misc/netifrc):

`root #` `emerge --ask --noreplace net-misc/netifrc`

По умолчанию используется DHCP. Но для того, чтобы он заработал, необходимо установить DHCP-клиент. Это будет описано далее в разделе «Установка необходимым системных пакетов».

Если сетевое соединение требует дополнительной настройки DHCP или вовсе не использует DHCP, тогда откройте /etc/conf.d/net:

`root #` `nano /etc/conf.d/net`

Настройте оба параметра config\_eth0 и routes\_eth0, введя информацию о IP-адресе и информацию о маршрутизации:

**Заметка**

Мы предполагаем, что сетевой интерфейс будет называться eth0, однако это во многом зависит от системы. Будем считать, что интерфейс будет называться так же, как он назывался при загрузке с установочного носителя, если установочный носитель достаточно свежий. Больше информации можно найти в разделе [Именование сетевых интерфейсов](/wiki/Handbook:AMD64/Networking/Advanced/ru#Network_interface_naming "Handbook:AMD64/Networking/Advanced/ru").

ФАЙЛ **`/etc/conf.d/net`** **Настройка статического IP-адреса**

```
config_eth0="192.168.0.2 netmask 255.255.255.0 brd 192.168.0.255"
routes_eth0="default via 192.168.0.1"

```

Для использования DHCP настройте config\_eth0:

ФАЙЛ **`/etc/conf.d/net`** **Настройка для работы DHCP**

```
config_eth0="dhcp"

```

Для получения списка дополнительных настроек прочтите /usr/share/doc/netifrc-\*/net.example.bz2. Не забудьте также прочитать man-страницу для DHCP-клиента, если требуется сделать дополнительные настройки.

Если в системе имеются несколько сетевых интерфейсов, то повторите предыдущие шаги для config\_eth1, config\_eth2, и так далее.

Теперь сохраните настройки и выйдите из редактора, чтобы продолжить далее.

##### Автоматический запуск сетевого подключения при загрузке системы

Для того, чтобы сетевые интерфейсы начинали работать во время загрузки системы, их необходимо добавить к уровню запуска по умолчанию.

`root #` `cd /etc/init.d
`

`root #` `ln -s net.lo net.eth0
`

`root #` `rc-update add net.eth0 default`

Если в системе есть несколько сетевых интерфейсов, то соответствующие файлы net.\* должны быть созданы также, как мы сделали это для net.eth0.

Если после загрузки системы выяснилось, что имя сетевого интерфейса (которое в настоящее время указано как `eth0`) было неверным, выполните следующие действия для исправления:

1. Измените настройки в файле /etc/conf.d/net, используя правильное название интерфейса (например, `enp3s0` или `enp5s0` вместо `eth0`).
2. Создайте новую символьную ссылку (например, /etc/init.d/net.enp3s0).
3. Удалите старую символьную ссылку (rm /etc/init.d/net.eth0).
4. Добавьте новую в уровень запуска по умолчанию.
5. Удалите старую с помощью rc-update del net.eth0 default.

### Файл hosts

Следующим важным шагом является оповещение других узлов в сетевом окружении о новой системе. Сетевые имена хостов определяются в файле /etc/hosts. Добавление имени хоста позволит разрешать имя в IP-адрес для хостов, которых нет в сервере имён.

`root #` `nano /etc/hosts`

ФАЙЛ **`/etc/hosts`** **Внесение сетевой информации**

```
# Это обязательные настройки для текущей системы
127.0.0.1     tux.homenetwork tux localhost
::1           tux.homenetwork tux localhost

# Необязательные определения для других систем в сети
192.168.0.5   jenny.homenetwork jenny
192.168.0.6   benny.homenetwork benny

```

Сохраните и закройте текстовый редактор для продолжения.

## Системная информация

### Пароль суперпользователя

Изменить пароль суперпользователя (с именем root) можно с помощью команды passwd.

`root #` `passwd`

Позже будет создан обычный пользователь для повседневных задач.

### Настройка инициализации и загрузки

#### OpenRC

Если вы используете OpenRC, эта система инициализации использует /etc/rc.conf для настройки сервисов, запуска и остановки системы. Откройте /etc/rc.conf и прочтите комментарии в файле. Проверьте настройки и измените их при необходимости.

`root #` `nano /etc/rc.conf`

Далее, откройте /etc/conf.d/keymaps для настройки раскладки клавиатуры. Отредактируйте файл и выберите нужную раскладку.

`root #` `nano /etc/conf.d/keymaps`

Соблюдайте особую осторожность с переменной keymap. Если выбрать неправильный раскладку, то может получится странный результат при печати текста.

Наконец, отредактируйте /etc/conf.d/hwclock чтобы установить параметры часов. Отредактируйте его в соответствии с личными предпочтениями.

`root #` `nano /etc/conf.d/hwclock`

Если аппаратные часы не настроены на время UTC, то в файле необходимо установить `clock="local"`. В противном случае система может отображать неправильное время.

#### systemd

Сначала рекомендуется запустить systemd-machine-id-setup, а затем systemd-firstboot, которые подготовят различные компоненты системы к первой загрузке в новой среде systemd. При передаче следующих опций пользователю будет предложено установить локаль, часовой пояс, имя хоста, пароль от root и пользовательскую оболочку для root. После этого, установке будет присвоен случайный идентификатор машины:

`root #` `systemd-machine-id-setup`

`root #` `systemd-firstboot --prompt`

Затем следует запустить systemctl, чтобы сбросить все установленные файлы устройств на предустановленные значения правил:

**Заметка**

При выполнении данной команды может отобразиться предупреждение о сбое. Это ожидаемое поведение, так как Gentoo LiveCD использует OpenRC.

`root #` `systemctl preset-all --preset-mode=enable-only`

Также можно запустить полное изменение предустановленных значений, но это может сбросить все службы, которые уже были настроены во время процесса:

`root #` `systemctl preset-all`

Эти два шага помогут обеспечить плавный переход от среды установщика к первой загрузке системы.

[← Настройка ядра Linux](/wiki/Handbook:AMD64/Installation/Kernel/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:AMD64/ru "Handbook:AMD64/ru") [Установка системных средств →](/wiki/Handbook:AMD64/Installation/Tools/ru "Следующая часть")