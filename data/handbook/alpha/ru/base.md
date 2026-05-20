# Base

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Base/de "Handbuch:Alpha/Installation/Basis (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Base "Handbook:Alpha/Installation/Base (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Base/es "Manual de Gentoo: Alpha/Instalación/Base (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Base/fr "Manuel:Alpha/Installation/Base (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Base/it "Handbook:Alpha/Installation/Base/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Base/hu "Handbook:Alpha/Installation/Base/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Base/pl "Handbook:Alpha/Installation/Base (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Base/pt-br "Manual:Alpha/Instalação/Base (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Base/cs "Handbook:Alpha/Installation/Base/cs (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Base/ta "கையேடு:Alpha/நிறுவல்/அடிப்படை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Base/zh-cn "手册：Alpha/安装/安装基本系统 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Base/ja "ハンドブック:Alpha/インストール/ベース (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Base/ko "Handbook:Alpha/Installation/Base/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha/ru "Handbook:Alpha/ru")[Установка](/wiki/Handbook:Alpha/Full/Installation/ru "Handbook:Alpha/Full/Installation/ru")[Об установке](/wiki/Handbook:Alpha/Installation/About/ru "Handbook:Alpha/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:Alpha/Installation/Media/ru "Handbook:Alpha/Installation/Media/ru")[Настройка сети](/wiki/Handbook:Alpha/Installation/Networking/ru "Handbook:Alpha/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:Alpha/Installation/Disks/ru "Handbook:Alpha/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:Alpha/Installation/Stage/ru "Handbook:Alpha/Installation/Stage/ru")Установка базовой системы[Настройка ядра](/wiki/Handbook:Alpha/Installation/Kernel/ru "Handbook:Alpha/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:Alpha/Installation/System/ru "Handbook:Alpha/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:Alpha/Installation/Tools/ru "Handbook:Alpha/Installation/Tools/ru")[Настройка загрузчика](/wiki/Handbook:Alpha/Installation/Bootloader/ru "Handbook:Alpha/Installation/Bootloader/ru")[Завершение](/wiki/Handbook:Alpha/Installation/Finalizing/ru "Handbook:Alpha/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:Alpha/Full/Working/ru "Handbook:Alpha/Full/Working/ru")[Введение в Portage](/wiki/Handbook:Alpha/Working/Portage/ru "Handbook:Alpha/Working/Portage/ru")[USE-флаги](/wiki/Handbook:Alpha/Working/USE/ru "Handbook:Alpha/Working/USE/ru")[Возможности Portage](/wiki/Handbook:Alpha/Working/Features/ru "Handbook:Alpha/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:Alpha/Working/Initscripts/ru "Handbook:Alpha/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:Alpha/Working/EnvVar/ru "Handbook:Alpha/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:Alpha/Full/Portage/ru "Handbook:Alpha/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:Alpha/Portage/Files/ru "Handbook:Alpha/Portage/Files/ru")[Переменные](/wiki/Handbook:Alpha/Portage/Variables/ru "Handbook:Alpha/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:Alpha/Portage/Branches/ru "Handbook:Alpha/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:Alpha/Portage/Tools/ru "Handbook:Alpha/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:Alpha/Portage/CustomTree/ru "Handbook:Alpha/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:Alpha/Portage/Advanced/ru "Handbook:Alpha/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:Alpha/Full/Networking/ru "Handbook:Alpha/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:Alpha/Networking/Introduction/ru "Handbook:Alpha/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:Alpha/Networking/Advanced/ru "Handbook:Alpha/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:Alpha/Networking/Modular/ru "Handbook:Alpha/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:Alpha/Networking/Wireless/ru "Handbook:Alpha/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:Alpha/Networking/Extending/ru "Handbook:Alpha/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:Alpha/Networking/Dynamic/ru "Handbook:Alpha/Networking/Dynamic/ru")

## Contents

- [1Переход в изолированную среду](#.D0.9F.D0.B5.D1.80.D0.B5.D1.85.D0.BE.D0.B4_.D0.B2_.D0.B8.D0.B7.D0.BE.D0.BB.D0.B8.D1.80.D0.BE.D0.B2.D0.B0.D0.BD.D0.BD.D1.83.D1.8E_.D1.81.D1.80.D0.B5.D0.B4.D1.83)
  - [1.1Копирование информации о DNS](#.D0.9A.D0.BE.D0.BF.D0.B8.D1.80.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_.D0.B8.D0.BD.D1.84.D0.BE.D1.80.D0.BC.D0.B0.D1.86.D0.B8.D0.B8_.D0.BE_DNS)
  - [1.2Монтирование необходимых файловых систем](#.D0.9C.D0.BE.D0.BD.D1.82.D0.B8.D1.80.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_.D0.BD.D0.B5.D0.BE.D0.B1.D1.85.D0.BE.D0.B4.D0.B8.D0.BC.D1.8B.D1.85_.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2.D1.8B.D1.85_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC)
  - [1.3Переход в новое окружение](#.D0.9F.D0.B5.D1.80.D0.B5.D1.85.D0.BE.D0.B4_.D0.B2_.D0.BD.D0.BE.D0.B2.D0.BE.D0.B5_.D0.BE.D0.BA.D1.80.D1.83.D0.B6.D0.B5.D0.BD.D0.B8.D0.B5)
  - [1.4Подготовка к установке начального загрузчика](#.D0.9F.D0.BE.D0.B4.D0.B3.D0.BE.D1.82.D0.BE.D0.B2.D0.BA.D0.B0_.D0.BA_.D1.83.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B5_.D0.BD.D0.B0.D1.87.D0.B0.D0.BB.D1.8C.D0.BD.D0.BE.D0.B3.D0.BE_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D1.87.D0.B8.D0.BA.D0.B0)
    - [1.4.1DOS/Legacy BIOS systems](#DOS.2FLegacy_BIOS_systems)
- [2Настройка Portage](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_Portage)
  - [2.1Установка снимка Gentoo репозитория ebuild-файлов](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_.D1.81.D0.BD.D0.B8.D0.BC.D0.BA.D0.B0_Gentoo_.D1.80.D0.B5.D0.BF.D0.BE.D0.B7.D0.B8.D1.82.D0.BE.D1.80.D0.B8.D1.8F_ebuild-.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2)
  - [2.2Необязательно: Выбор зеркала](#.D0.9D.D0.B5.D0.BE.D0.B1.D1.8F.D0.B7.D0.B0.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE:_.D0.92.D1.8B.D0.B1.D0.BE.D1.80_.D0.B7.D0.B5.D1.80.D0.BA.D0.B0.D0.BB.D0.B0)
    - [2.2.1Optional: rsync mirrors](#Optional:_rsync_mirrors)
  - [2.3Необязательно: Обновление Gentoo репозитория ebuild-файлов](#.D0.9D.D0.B5.D0.BE.D0.B1.D1.8F.D0.B7.D0.B0.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE:_.D0.9E.D0.B1.D0.BD.D0.BE.D0.B2.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_Gentoo_.D1.80.D0.B5.D0.BF.D0.BE.D0.B7.D0.B8.D1.82.D0.BE.D1.80.D0.B8.D1.8F_ebuild-.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2)
  - [2.4Чтение новостей](#.D0.A7.D1.82.D0.B5.D0.BD.D0.B8.D0.B5_.D0.BD.D0.BE.D0.B2.D0.BE.D1.81.D1.82.D0.B5.D0.B9)
  - [2.5Выбор подходящего профиля](#.D0.92.D1.8B.D0.B1.D0.BE.D1.80_.D0.BF.D0.BE.D0.B4.D1.85.D0.BE.D0.B4.D1.8F.D1.89.D0.B5.D0.B3.D0.BE_.D0.BF.D1.80.D0.BE.D1.84.D0.B8.D0.BB.D1.8F)
  - [2.6Необязательно: Добавление узла бинарных пакетов](#.D0.9D.D0.B5.D0.BE.D0.B1.D1.8F.D0.B7.D0.B0.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE:_.D0.94.D0.BE.D0.B1.D0.B0.D0.B2.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_.D1.83.D0.B7.D0.BB.D0.B0_.D0.B1.D0.B8.D0.BD.D0.B0.D1.80.D0.BD.D1.8B.D1.85_.D0.BF.D0.B0.D0.BA.D0.B5.D1.82.D0.BE.D0.B2)
    - [2.6.1Конфигурация репозитория](#.D0.9A.D0.BE.D0.BD.D1.84.D0.B8.D0.B3.D1.83.D1.80.D0.B0.D1.86.D0.B8.D1.8F_.D1.80.D0.B5.D0.BF.D0.BE.D0.B7.D0.B8.D1.82.D0.BE.D1.80.D0.B8.D1.8F)
    - [2.6.2Installing binary packages](#Installing_binary_packages)
  - [2.7Необязательно: Настройка переменной USE](#.D0.9D.D0.B5.D0.BE.D0.B1.D1.8F.D0.B7.D0.B0.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE:_.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D0.BF.D0.B5.D1.80.D0.B5.D0.BC.D0.B5.D0.BD.D0.BD.D0.BE.D0.B9_USE)
    - [2.7.1CPU\_FLAGS\_\*](#CPU_FLAGS_.2A)
    - [2.7.2VIDEO\_CARDS](#VIDEO_CARDS)
  - [2.8Необязательно: Настройка переменной ACCEPT\_LICENSE](#.D0.9D.D0.B5.D0.BE.D0.B1.D1.8F.D0.B7.D0.B0.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE:_.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D0.BF.D0.B5.D1.80.D0.B5.D0.BC.D0.B5.D0.BD.D0.BD.D0.BE.D0.B9_ACCEPT_LICENSE)
  - [2.9Необязательно: Обновление набора @world](#.D0.9D.D0.B5.D0.BE.D0.B1.D1.8F.D0.B7.D0.B0.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE:_.D0.9E.D0.B1.D0.BD.D0.BE.D0.B2.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_.D0.BD.D0.B0.D0.B1.D0.BE.D1.80.D0.B0_.40world)
    - [2.9.1Удаление устаревших пакетов](#.D0.A3.D0.B4.D0.B0.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_.D1.83.D1.81.D1.82.D0.B0.D1.80.D0.B5.D0.B2.D1.88.D0.B8.D1.85_.D0.BF.D0.B0.D0.BA.D0.B5.D1.82.D0.BE.D0.B2)
- [3Часовой пояс](#.D0.A7.D0.B0.D1.81.D0.BE.D0.B2.D0.BE.D0.B9_.D0.BF.D0.BE.D1.8F.D1.81)
- [4Настройка локалей](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D0.BB.D0.BE.D0.BA.D0.B0.D0.BB.D0.B5.D0.B9)
  - [4.1Генерация локалей](#.D0.93.D0.B5.D0.BD.D0.B5.D1.80.D0.B0.D1.86.D0.B8.D1.8F_.D0.BB.D0.BE.D0.BA.D0.B0.D0.BB.D0.B5.D0.B9)
  - [4.2Выбор локали](#.D0.92.D1.8B.D0.B1.D0.BE.D1.80_.D0.BB.D0.BE.D0.BA.D0.B0.D0.BB.D0.B8)
- [5Примечания](#.D0.9F.D1.80.D0.B8.D0.BC.D0.B5.D1.87.D0.B0.D0.BD.D0.B8.D1.8F)

## Переход в изолированную среду

### Копирование информации о DNS

Единственное, что ещё осталось сделать перед входом в новое окружение, это скопировать информацию о DNS из файла /etc/resolv.conf. Это нужно сделать, чтобы сеть всё ещё будет работать даже после входа в новое окружение. Файл /etc/resolv.conf содержит сервера имён.

Чтобы скопировать эту информацию, рекомендуется ввести ключ `--dereference` для команды cp. Благодаря этому /etc/resolv.conf будет скопирован как файл, если является символьной ссылкой. В противном случае в новом окружении символическая ссылка будет ссылаться на несуществующий файл (так как цель ссылки, скорее всего, будет недоступна внутри нового окружения).

`root #` `cp --dereference /etc/resolv.conf /mnt/gentoo/etc/`

### Монтирование необходимых файловых систем

**Совет**

При использовании установочного носителя Gentoo этот шаг можно заменить простой командой: arch-chroot /mnt/gentoo.

Через несколько мгновений корневая система Linux будет перемещена в новое место.

Файловые системы, которые должны быть доступны:

- /proc/ — псевдофайловая система. Она выглядит как обычные файлы, но на самом деле генерируется на лету ядром Linux
- /sys/ — псевдофайловая система, как и /proc/, которую она когда-то должна была заменить, и более структурированнее, чем /proc/
- /dev/ — это обычная файловая система, которая содержит все файлы устройств. Она частично управляемется менеджером устройств Linux (обычно udev)
- /run/ — временная файловая система, которая содержит генерируемые на лету файлы по типу файлов PID или файлов блокировки (locks)

Каталог /proc/ монтируется в /mnt/gentoo/proc/, остальные — через перепривязку точки монтирования. Это означает, что, например, /mnt/gentoo/sys/ на самом деле _будет_/sys/ (это просто вторая точка входа в ту же файловую систему), тогда как /mnt/gentoo/proc/ является новой точкой монтирования (так сказать, экземпляром) файловой системы.

`root #` `mount --types proc /proc /mnt/gentoo/proc
`

`root #` `mount --rbind /sys /mnt/gentoo/sys
`

`root #` `mount --make-rslave /mnt/gentoo/sys
`

`root #` `mount --rbind /dev /mnt/gentoo/dev
`

`root #` `mount --make-rslave /mnt/gentoo/dev
`

`root #` `mount --bind /run /mnt/gentoo/run
`

`root #` `mount --make-slave /mnt/gentoo/run
`

**Заметка**

Параметр `--make-rslave` необходим для дальнейшей поддержки systemd в ходе установки.

**Предупреждение**

Если при установке используется не дистрибутив Gentoo, то этого может быть недостаточно. Некоторые дистрибутивы делают /dev/shm символьной ссылкой на /run/shm/, которая после перехода в изолированную среду станет недействительной. Создание правильного подключения /dev/shm/ в tmpfs поможет избежать этой проблемы:

`root #` `test -L /dev/shm && rm /dev/shm && mkdir /dev/shm
`

`root #` `mount --types tmpfs --options nosuid,nodev,noexec shm /dev/shm
`

Также проверьте, что права доступа установлены в 1777:

`root #` `chmod 1777 /dev/shm /run/shm`

### Переход в новое окружение

Теперь, когда все разделы инициализированы и базовое окружение установлено, настало время войти в новое установочное окружение (выполнить chroot). Это означает, что сессия изменит свой корневой каталог (наивысший каталог, в который можно перейти) из текущей установочного окружения (CD или другого установочного носителя) в систему установки (где находятся размеченные разделы).

Переход в изолированное окружение делается в три шага:

1. Изменение корневого каталога с / (который находится на установочном носителе) в /mnt/gentoo/ (на разделах диска) с помощью команды chroot или arch-chroot если доступно.
2. Загрузка в память некоторых параметров из /etc/profile с помощью команды source
3. Изменение приглашения командной строки, чтобы не забыть, что эта сессия находится в изолированном окружении.

`root #` `chroot /mnt/gentoo /bin/bash
`

`root #` `source /etc/profile
`

`root #` `export PS1="(chroot) ${PS1}"`

С этого момента все действия выполняются непосредственно в новом окружении Gentoo Linux.

**Совет**

Если установка Gentoo будет прервана где-то после этой точки, _можно_ «продолжить» установку с последнего состояния. Не нужно разбивать диск снова! Просто [смонтируйте корневой раздел](/wiki/Handbook:Alpha/Installation/Disks/ru#Mounting_the_root_partition "Handbook:Alpha/Installation/Disks/ru") снова и проделайте предыдущие шаги, начиная с [копирования информации о DNS](/wiki/Handbook:Alpha/Installation/Base/ru#Copy_DNS_info "Handbook:Alpha/Installation/Base/ru"), для повторного входа в рабочее окружение. Эти шаги подойдут и для решения проблем с загрузчиком. Больше информации можно найти в статье [chroot](/wiki/Chroot/ru "Chroot/ru").

### Подготовка к установке начального загрузчика

Now that the new environment has been entered, it is necessary to prepare the new environment for the bootloader. It will be important to have the correct partition mounted when it is time to install the bootloader.

#### DOS/Legacy BIOS systems

For DOS/Legacy BIOS systems, the bootloader will be installed into the /boot directory, therefore mount as follows:

`root #` `mount /dev/sda1 /boot`

## Настройка Portage

### Установка снимка Gentoo репозитория ebuild-файлов

Следующим шагом будет установка снимка репозитория ebuild-файлов Gentoo. Этот снимок содержит коллекцию файлов, которая сообщает Portage о доступных программах (для установки), какой профиль может выбрать системный администратор, о новостях о конкретных пакетах или профилях и так далее.

emerge-webrsync рекомендуется использовать в случаях, когда система находится за межсетевым экраном (для загрузки снимка используется только протоколы HTTP/HTTPS), а также когда необходимо снизить нагрузку канал сети. У кого нет ограничений с сетью или шириной канала, могут счастливо перейти к следующему разделу.

Команда ниже загрузит последний снимок (которые выпускаются каждый день), с одного из зеркал Gentoo, и распакует его в системе:

`root #` `emerge-webrsync`

**Заметка**

Во время этой операции, emerge-webrsync может жаловаться на отсутствие /var/db/repos/gentoo/. В этом нет ничего страшного — инструмент сам создаст этот каталог.

Начиная с этого места, Portage может попросить установить некоторые рекомендуемые обновления: некоторые системные пакеты, установленные из архива stage, могут иметь новые доступные версии. Пакетному менеджеру теперь известно о новых пакетах благодаря снимку репозитория. Обновление пакетов можно проигнорировать, этот процесс можно отложить до завершения установки Gentoo.

### Необязательно: Выбор зеркала

Для быстрой загрузки исходного кода рекомендуется выбрать быстрое, географически находящееся рядом зеркало. Portage будет искать в файле make.conf переменную GENTOO\_MIRRORS и использовать перечисленные в ней зеркала. Можно просмотреть список зеркал Gentoo и найти зеркало (или несколько зеркал), наиболее близко расположенное к месту физического расположения (чаще всего они и есть самые быстрые).

A tool called mirrorselect provides a pretty text interface to more quickly query and select suitable mirrors. Just navigate to the mirrors of choice and press `Spacebar` to select one or more mirrors.

`root #` `emerge --ask --verbose --oneshot app-portage/mirrorselect`

`root #` `mirrorselect -i -o >> /etc/portage/make.conf`

Alternatively, a list of active mirrors are [available online](https://www.gentoo.org/downloads/mirrors/).

#### Optional: rsync mirrors

Gentoo also has many rsync mirrors which can speed up downloading the available package list using `emerge --sync` (explained in more detail later) by selecting a mirror closer that is geographically closer to the user.

`root #` `mkdir /etc/portage/repos.conf
`

`root #` `cp /usr/share/portage/config/repos.conf /etc/portage/repos.conf/gentoo.conf
`

Select a mirror close to the system's location from [https://www.gentoo.org/support/rsync-mirrors/](https://www.gentoo.org/support/rsync-mirrors/) and replace the sync-uri default mirror in /etc/portage/repos.conf/gentoo.conf with the desired mirror location.

ФАЙЛ **`/etc/portage/repos.conf/gentoo.conf`** **UK-based rsync mirror example**

```
[DEFAULT]
main-repo = gentoo
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
[gentoo]
location = /var/db/repos/gentoo
sync-type = rsync
sync-uri = rsync://rsync.uk.gentoo.org/gentoo-portage/
auto-sync = yes
sync-rsync-verify-jobs = 1
sync-rsync-verify-metamanifest = yes
sync-rsync-verify-max-age = 3
sync-openpgp-key-path = /usr/share/openpgp-keys/gentoo-release.asc
sync-openpgp-keyserver = hkps://keys.gentoo.org
sync-openpgp-key-refresh-retry-count = 40
sync-openpgp-key-refresh-retry-overall-timeout = 1200
sync-openpgp-key-refresh-retry-delay-exp-base = 2
sync-openpgp-key-refresh-retry-delay-max = 60
sync-openpgp-key-refresh-retry-delay-mult = 4
sync-webrsync-verify-signature = yes
sync-git-verify-commit-signature = true

```

### Необязательно: Обновление Gentoo репозитория ebuild-файлов

Также можно обновить репозиторий ebuild-файлов Gentoo до текущего состояния. Предыдущая команда emerge-webrsync устанавливает относительно недавний снимок (обычно не старше суток), поэтому этот шаг совершенно необязателен.

Если необходимо установить последние обновления пакетов (выпущенных не более 1 часа назад), то используйте emerge --sync. Эта команда использует rsync-протокол для обновления репозитория ebuild-файлов Gentoo (которое было получено ранее с помощью emerge-webrsync) до самой свежей версии.

`root #` `emerge --sync`

На медленных терминалах (с некоторыми медленными кадровыми буферами или через последовательный порт), рекомендуется использовать параметр `--quiet` для ускорения процесса:

`root #` `emerge --sync --quiet`

### Чтение новостей

После обновления репозитория ebuild-файлов Gentoo, Portage может вывести похожие сообщения:

` * IMPORTANT: 2 news items need reading for repository 'gentoo'.
`

` * Use eselect news to read news items.
`

Новостные сообщения были созданы для обеспечения коммуникационного канала и оповещения пользователей о важных событиях через репозиторий ebuild-файлов Gentoo. Для управления оповещениями используйте команду eselect news. Приложение eselect предоставляет общий интерфейс для системного администрирования. В данном случае eselect используется совместно с модулем `news`.

Для модуля `news` есть три наиболее распространенных операций:

- `list` отображает общий список новостей.
- с помощью `read` можно прочитать какую-либо новость.
- `purge` удалит прочитанные новости, поэтому перечитать новость снова уже будет нельзя.

`root #` `eselect news list
`

`root #` `eselect news read`

Более подробную информацию о чтении новостей можно найти на странице man:

`root #` `man news.eselect`

### Выбор подходящего профиля

**Совет**

Профили desktop (для настольных систем) можно использовать не только со _средами рабочего стола_. Они также подходят и для минимальных менеджеров окон, таких как i3 или sway.

_Профиль_ — это важная часть любой системы Gentoo. Он не только определяет такие важные переменные, как USE, CFLAGS и многие другие, а также фиксирует версии для определённых пакетов. Все эти нюансы поддерживаются разработчиками Gentoo.

Для того, чтобы увидеть, какой профиль использует система на данный момент, запустите eselect с модулем `profile`:

**Совет**

On an install media without a scroll-able terminal, `eselect profile list | less` can provide an easy way to list all available profiles while providing the ability to scroll.

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/alpha/23.0 *
  [2]   default/linux/alpha/23.0/desktop
  [3]   default/linux/alpha/23.0/desktop/gnome
  [4]   default/linux/alpha/23.0/desktop/kde

```

**Заметка**

Вывод команды является только примером и может меняться время от времени.

**Заметка**

Для использования **systemd** выберите профиль, в котором содержится слово «systemd» (или его отсутствие в обратном случае).

Для некоторых архитектур есть субпрофиль для настольных систем.

**Предупреждение**

Не стоит халатно относиться к обновлениям профиля. Выбирая изначальный профиль, используйте тот, что соответствует **той же версии**, которая была использована в stage-файле (к примеру, 23.0). Каждая новая версия профиля объявляется через элемент новостей с инструкциями по миграции. Внимательно следуйте им, прежде чем перейти на новый профиль.

После просмотра доступных профилей для архитектуры alpha, пользователи могут выбрать другой системный профиль:

`root #` `eselect profile set 2`

**Заметка**

Подпрофиль `developer` сделан специально для разработки Gentoo Linux и не предназначен для использования обычными пользователями.

### Необязательно: Добавление узла бинарных пакетов

Since December 2023, Gentoo's [Release Engineering team](/wiki/Project:RelEng "Project:RelEng") has offered an [official binary package host](/wiki/Binary_package_quickstart "Binary package quickstart") (colloquially shorted to just "binhost") for use by the general community to retrieve and install binary packages (binpkgs).[\[1\]](#cite_note-1)

Adding a binary package host allows Portage to install cryptographically signed, compiled packages. In many cases, adding a binary package host will _greatly_ decrease the mean time to package installation and adds much benefit when running Gentoo on older, slower, or low power systems.

#### Конфигурация репозитория

The repository configuration for a binhost is found in Portage's /etc/portage/binrepos.conf/ directory, which functions similarly to the configuration mentioned in the [Gentoo ebuild repository](/wiki/Handbook:Alpha/Installation/Base/ru#Gentoo_ebuild_repository "Handbook:Alpha/Installation/Base/ru") section.

When defining a binary host, there are two important aspects to consider:

1. The architecture and profile targets within the `sync-uri` value _do_ matter and should align to the respective computer architecture ( **alpha** in this case) and system profile selected in the [Choosing the right profile](/wiki/Handbook:Alpha/Installation/Base/ru#Choosing_the_right_profile "Handbook:Alpha/Installation/Base/ru") section.
2. Selecting a fast, geographically close mirror will generally shorten retrieval time. Review the mirrorselect tool mentioned in the [Optional: Selecting mirrors](/wiki/Handbook:Alpha/Installation/Base/ru#Gentoo_ebuild_repository "Handbook:Alpha/Installation/Base/ru") section or review the [online list of mirrors](https://www.gentoo.org/downloads/mirrors/) where URL values can be discovered.


ФАЙЛ **`/etc/portage/binrepos.conf/gentoo.conf`** **CDN-based binary package host example**

```
[gentoo]
priority = 9959
# NOTE: Must adjust <arch> and <variant> as appropriate!
sync-uri = https://distfiles.gentoo.org/releases/<arch>/binpackages/<variant>
# x86-64 example sync-uri
# sync-uri = https://distfiles.gentoo.org/releases/amd64/binpackages/23.0/x86-64/
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
# Introduced in portage-3.0.74 for per-repo verification choices
verify-signature = true
# Default value with >=portage-3.0.77
location = /var/cache/binhost/gentoo

```

#### Installing binary packages

Portage will compile packages from code source by default. It can be instructed to use binary packages in the following ways:

1. The `--getbinpkg` option can be passed when invoking the emerge command. This method of binary package installation is useful to install only a particular binary package.
2. Changing the system's default via Portage's FEATURES variable, which is exposed through the /etc/portage/make.conf file. Applying this configuration change will cause Portage to query the binary package host for the package(s) to be requested and fall back to compiling locally when no results are found.

For example, to have Portage always install available binary packages:

ФАЙЛ **`/etc/portage/make.conf`** **Configure Portage to use binary packages by default**

```
# Appending getbinpkg to the list of values within the FEATURES variable
FEATURES="${FEATURES} getbinpkg"
# Require signatures
FEATURES="${FEATURES} binpkg-request-signature"

```

Please also run getuto for Portage to set up the necessary keyring for verification:

`root #` `getuto`

Additional Portage features will be discussed in the [the next chapter](/wiki/Handbook:Alpha/Working/Features/ru#Portage_features "Handbook:Alpha/Working/Features/ru") of the handbook.

### Необязательно: Настройка переменной USE

USE — это одна из самых мощных переменных Gentoo, доступная пользователям. Разные программы могут быть скомпилированы с или без поддержки некоторых элементов. Например, некоторые программы могут быть собраны с поддержкой GTK+ или Qt. Другие могут быть собраны с или без поддержки SSL. Некоторые программы можно даже собрать с поддержкой кадрового буфера (svgalib) вместо X11 (X-сервера).

Большинство дистрибутивов компилируют свои пакеты с поддержкой всего, что возможно, увеличивая размер и время запуска программ, не говоря уже о чрезмерных зависимостях.
Благодаря Gentoo пользователь может определить, с какими параметрами пакет должен быть скомпилирован. И здесь переменная USE вступает в игру.

В переменной USE пользователи могут определить ключевые слова, которые сказываются на параметрах сборки. Например, `ssl` компилирует SSL-поддержку в программах, которые её поддерживают. `-X` уберёт поддержку X-сервера (обратите внимание на знак минус перед X). `gnome gtk -kde -qt5` скомпилирует программы с поддержкой GNOME (и GTK+), но без поддержки KDE (и Qt), что делает систему более оптимальной для использования GNOME (если архитектура поддерживает его).

Настройки по умолчанию для USE находятся в файле make.defaults профиля Gentoo, который используется на данный момент системой. Gentoo использует сложную систему наследования для своих профилей, которая не будет подробно рассматриваться во время установки. Простой способ проверить, какие настройки используются для USE — запустить emerge --info и просмотреть строку, начинающуюся с USE:

`root #` `emerge --info | grep ^USE`

```
USE="X acl alsa amd64 berkdb bindist bzip2 cli cracklib crypt cxx dri ..."

```

**Заметка**

В приведённом выше примере список укорочен. Настоящий список USE флагов намного больше.

Полное описание всех доступных USE-флагов можно найти в файле /var/db/repos/gentoo/profiles/use.desc.

`root #` `less /var/db/repos/gentoo/profiles/use.desc`

При использовании команды less можно использовать прокрутку с помощью клавиш `↑` и `↓`, для выхода нажмите клавишу `q`.

В качестве примера мы покажем настройки USE для системы ориентированной для использования KDE с поддержкой DVD, ALSA и записи CD:

`root #` `nano /etc/portage/make.conf`

ФАЙЛ **`/etc/portage/make.conf`** **Настройки переменной USE для системы, ориентированной на использование KDE/Plasma с поддержкой DVD, ALSA и записи CD**

```
USE="-gtk -gnome qt5 kde dvd alsa cdr"

```

Если USE-флаг определён в /etc/portage/make.conf, он будет _добавлен_ (или _удалён_, если перед USE-флагом написан знак `-`) в список по умолчанию. USE флаги могут быть _удалены_ добавлением знака `-` перед флагом. Например, чтобы выключить поддержку для графических окружений X, нужно задать `-X`:

ФАЙЛ **`/etc/portage/make.conf`** **Игнорирование USE-флагов по умолчанию**

```
USE="-X acl alsa"

```

**Предупреждение**

Пусть это и возможно, но установка `-*` (что уберёт все USE флаги, кроме указанных в make.conf) не одобряется. Писатели ebuild'ов подбирают нужные флаги по умолчанию, для того, чтобы избежать конфликтов, улучшить безопасность, и по другим причинам. Удаление _всех_ USE флагов нарушит поведение по умолчанию и может вызвать серьезные проблемы.

#### CPU\_FLAGS\_\*

Некоторые архитектуры (включая AMD64/X86, ARM, PPC) имеют переменную [USE\_EXPAND](/wiki/USE_EXPAND "USE EXPAND") под названием [CPU\_FLAGS\_<ARCH>](/wiki/CPU_FLAGS_X86 "CPU FLAGS X86") (где <ARCH> заменяется на соответствующую архитектуру системы).

**Важно**

Не перепутайте! Системы **AMD64** и **X86** имеют общую архитектуру, поэтому для систем **AMD64** правильной переменной является CPU\_FLAGS\_X86.

Это используется для настройки сборки на компиляцию в определённый ассемблерный код или другие интринсики, обычно написанные вручную или каким-нибудь другим дополнительным способом, и это **не** то же самое, что попросить компилятор вывести оптимизированный код для определенной характеристики процессора (как, например, `-march=`).

Пользователям рекомендуется установить эту переменную, по желанию одновременно с настройкой COMMON\_FLAGS.

Для настройки необходимо выполнить несколько шагов:

`root #` `emerge --ask --oneshot app-portage/cpuid2cpuflags`

Если интересно, проверьте вывод вручную:

`root #` `cpuid2cpuflags`

Затем скопируйте вывод в package.use:

`root #` `echo "*/* $(cpuid2cpuflags)" > /etc/portage/package.use/00cpu-flags`

#### VIDEO\_CARDS

By default a profile already sets a few video cards. For many reasons this is not ideal, but a useful safety net.

To configure the system correctly the user needs to first unset the preset video cards with `VIDEO_CARDS: -*` then set the correct card for that system.

ФАЙЛ **`/etc/portage/package.use/00video_cards`** **Example**

```
*/* VIDEO_CARDS: -* <GPU DRIVER NAME>

```

Below is a table that can be used to easily compare the different video card types to their respective `VIDEO_CARDS` value.

GPU
VIDEO\_CARDS
(Official) Nvidia Maxwell and newer`nvidia`Nouveau Nvidia [Supported list](/wiki/NVIDIA/ru "NVIDIA/ru")`nouveau`AMD since Sea Islands`amdgpu radeonsi`ATI and older AMDSee [radeon#Feature support](/wiki/Radeon#Feature_support "Radeon")Intel Nehalem and newer`intel`Intel Gen 2 and 3 [Supported list](/wiki/Intel#.23Feature_support.2Fru "Intel")`intel i915`QEMU/KVM`virgl`WSL`d3d12`

Below is a few examples of a properly set package.use for _VIDEO\_CARDS_:

ФАЙЛ **`/etc/portage/package.use/00video_cards`** **AMD example**

```
*/* VIDEO_CARDS: -* amdgpu radeonsi

```

ФАЙЛ **`/etc/portage/package.use/00video_cards`** **Intel example**

```
*/* VIDEO_CARDS: -* intel

```

ФАЙЛ **`/etc/portage/package.use/00video_cards`** **Nvidia example**

```
*/* VIDEO_CARDS: -* nvidia

```

Details for various GPU(s) can be found at the [AMDGPU](/wiki/AMDGPU "AMDGPU"), [Intel](/wiki/Intel "Intel"), [Nouveau (Open Source)](/wiki/Nouveau/ru "Nouveau/ru"), or [NVIDIA (Proprietary)](/wiki/NVIDIA/ru "NVIDIA/ru") articles.

### Необязательно: Настройка переменной ACCEPT\_LICENSE

Starting with Gentoo Linux Enhancement Proposal 23 (GLEP 23), a mechanism was created to allow system administrators the ability to "regulate the software they install with regards to licenses... Some want a system free of any software that is not OSI-approved; others are simply curious as to what licenses they are implicitly accepting."[\[2\]](#cite_note-2) With a motivation to have more granular control over the type of software running on a Gentoo system, the ACCEPT\_LICENSE variable was born.

Portage просматривает в ACCEPT\_LICENSE, пакеты с какими лицензияи разрешены для установки. Чтобы вывести текущее системное значение, выполните команду:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

Для удобства системного администрирования схожие с точки зрения права программные лицензии объединены во множества — каждое, соответствующее их предназначению. Определения групп лицензий [доступны для просмотра](https://gitweb.gentoo.org/repo/gentoo.git/tree/profiles/license_groups) и управляются проектом [Gentoo Licenses](/wiki/Project:Licenses "Project:Licenses"). В отличие от отдельных лицензий, группы начинаются символом `@`, что позволяет четко выделять их в переменной ACCEPT\_LICENSE.

Some common license groups include:

Название группыОписание
@GPL-COMPATIBLEСовместимые с GPL лицензии, одобренные Free Software Foundation[\[a\_license 1\]](#cite_note-3).
@FSF-APPROVEDЛицензии свободного ПО, одобренные FSF (включает @GPL-COMPATIBLE).
@OSI-APPROVEDЛицензии, одобренные Open Source Initiative[\[a\_license 2\]](#cite_note-4).
@MISC-FREEРазличные лицензии, которые, вероятнее всего, тоже относятся к свободному ПО, то есть следуют Определению Свободного ПО[\[a\_license 3\]](#cite_note-5), но не одобрены ни FSF, ни OSI.
@FREE-SOFTWAREСочетание @FSF-APPROVED, @OSI-APPROVED и @MISC-FREE.
@FSF-APPROVED-OTHERОдобренные FSF лицензии для «свободной документации» и «работ для практического применения, не являющихся ПО или документацией» (включая шрифты).
@MISC-FREE-DOCSРазличные лицензии для свободных документов и прочих работ (включая шрифты), следующие определению свободного произведения[\[a\_license 4\]](#cite_note-6), но НЕ включены в @FSF-APPROVED-OTHER.
@FREE-DOCUMENTSСочетание @FSF-APPROVED-OTHER и @MISC-FREE-DOCS.
@FREEНадмножество всех лицензий, обладающих свободой использования, распространения, изменения и распространения изменений. Сочетание @FREE-SOFTWARE и @FREE-DOCUMENTS.
@BINARY-REDISTRIBUTABLEЛицензии, разрешающие по крайней мере свободное распространение ПО в двоичной форме. Включает в себя @FREE.
@EULAЛицензионные соглашения, которые пытаются отобрать ваши права. Они более строги, чем «все права защищены», или могут требовать явного согласия.

1. [↑](#cite_ref-3)[https://www.gnu.org/licenses/license-list.html](https://www.gnu.org/licenses/license-list.html)
2. [↑](#cite_ref-4)[https://www.opensource.org/licenses](https://www.opensource.org/licenses)
3. [↑](#cite_ref-5)[https://www.gnu.org/philosophy/free-sw.html](https://www.gnu.org/philosophy/free-sw.html)
4. [↑](#cite_ref-6)[https://freedomdefined.org/](https://freedomdefined.org/)

Currently set system wide acceptable license values can be viewed via:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

As visible in the output, the default value is to only allow software which has been grouped into the `@FREE` category to be installed.

Specific licenses or licenses groups for a system can be defined in the following locations:

- Для всей системы в выбранном профиле.
- Для всей системы в файле /etc/portage/make.conf.
- По-пакетно в файле /etc/portage/package.license.
- По-пакетно в _каталоге_/etc/portage/package.license/ с файлами.

По желанию переопределите принятое в системе значение по умолчанию в профилях, изменив /etc/portage/make.conf.
/etc/portage/make.conf.

ФАЙЛ **`/etc/portage/make.conf`** **Пример разрешения лицензий с помощью ACCEPT\_LICENSE в масштабах всей системы**

```
ACCEPT_LICENSE="-* @FREE @BINARY-REDISTRIBUTABLE"

```

По желанию системные администраторы могут также определить разрешаемые лицензии по-пакетно, как показано в следующем примере для каталога с файлами. Обратите внимание, что _каталог_ package.license должен быть создан, если он ещё не существует:

`root #` `mkdir /etc/portage/package.license`

Software license details for an individual Gentoo package are stored within the LICENSE variable of the associated ebuild. One package may have one or many software licenses, therefore it be necessary to specify multiple acceptable licenses for a single package.

ФАЙЛ **`/etc/portage/package.license/kernel`** **Разрешение лицензий для конкретных пакетов**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware linux-fw-redistributable
sys-firmware/intel-microcode intel-ucode

```

**Важно**

Переменная LICENSE в ebuild-файле является только ориентиром для разработчиков и пользователей Gentoo. Она не является юридически значимым заявлением и не гарантирует, что условия использования соответствуют действительности. Не стоит доверять ей безоговорочно и при необходимости следует проводить полный аудит всех файлов, которые были установлены на системе.

### Необязательно: Обновление набора @world

**Совет**

Если до этого был выбран профиль для полноценной графической оболочки, процесс установки может занять значительное время. Оценить время установки очень просто: чем короче имя профиля, тем меньше будет набор [@world](/wiki/World_set_(Portage) "World set (Portage)"); чем меньше набор @world, тем меньше пакетов системе потребуется. Другими словами:

- При выборе `default/linux/amd64/23.0` потребует обновления небольшого количества пакетов, когда как
- При выборе `default/linux/amd64/23.0/desktop/gnome/systemd` потребует обновления гораздо большего числа пакетов, так как система инициализации поменяется с OpenRC на systemd, и будут установлены пакеты рабочего стола GNOME.

Это действие _необходимо_, чтобы система могла применить какие-либо обновления с момента сборки stage3 и обновления профиля:

1. A profile target _different_ from the stage file has been selected.
2. Additional USE flags have been set for installed packages.

Readers who are performing an 'install Gentoo speed run' may safely skip @world set updates until _after_ their system has rebooted into the new Gentoo environment.

Readers who are performing a slow run can have Portage perform updates for package, profile, and/or USE flag changes at the present time:

`root #` `emerge --ask --verbose --update --deep --changed-use @world`

Readers who added a binary host [above](/wiki/Handbook:Alpha/Installation/Base/ru#Optional:_Adding_a_binary_package_host "Handbook:Alpha/Installation/Base/ru") can add --getbinpkg (or -g) in order to fetch packages from the binary host instead of compiling them:

`root #` `emerge --ask --verbose --update --deep --newuse --getbinpkg @world`

#### Удаление устаревших пакетов

It is important to always _depclean_ after system upgrades to remove obsolete packages. Review the output carefully with emerge --depclean --pretend to see if any of the to-be-cleaned packages should be kept if personally using them. To keep a package which would otherwise be depcleaned, use emerge --noreplace foo.

`root #` `emerge --ask --pretend --depclean`

If happy, then proceed with a real depclean:

`root #` `emerge --ask --depclean`

## Часовой пояс

**Заметка**

Этот раздел неприменим к пользователям стандартной библиотеки C musl. Если вы не знаете, что это такое, то к вам это не относится и вы должны следовать этому разделу.

**Предупреждение**

Старайтесь не использовать часовые пояса, начинающиеся с /usr/share/zoneinfo/Etc/GMT\*, так как их названия не отражают настоящий часовой пояс. Например, GMT-8 на самом деле является GMT+8.

Выберите часовой пояс для системы. Просмотрите список всех доступных часовых поясов в каталоге /usr/share/zoneinfo/:

`root #` `ls -l /usr/share/zoneinfo`

```
total 352
drwxr-xr-x 2 root root   1120 Jan  7 17:41 Africa
drwxr-xr-x 6 root root   2960 Jan  7 17:41 America
drwxr-xr-x 2 root root    280 Jan  7 17:41 Antarctica
drwxr-xr-x 2 root root     60 Jan  7 17:41 Arctic
drwxr-xr-x 2 root root   2020 Jan  7 17:41 Asia
drwxr-xr-x 2 root root    280 Jan  7 17:41 Atlantic
drwxr-xr-x 2 root root    500 Jan  7 17:41 Australia
drwxr-xr-x 2 root root    120 Jan  7 17:41 Brazil
-rw-r--r-- 1 root root   2094 Dec  3 17:19 CET
-rw-r--r-- 1 root root   2310 Dec  3 17:19 CST6CDT
drwxr-xr-x 2 root root    200 Jan  7 17:41 Canada
drwxr-xr-x 2 root root     80 Jan  7 17:41 Chile
-rw-r--r-- 1 root root   2416 Dec  3 17:19 Cuba
-rw-r--r-- 1 root root   1908 Dec  3 17:19 EET
-rw-r--r-- 1 root root    114 Dec  3 17:19 EST
-rw-r--r-- 1 root root   2310 Dec  3 17:19 EST5EDT
-rw-r--r-- 1 root root   2399 Dec  3 17:19 Egypt
-rw-r--r-- 1 root root   3492 Dec  3 17:19 Eire
drwxr-xr-x 2 root root    740 Jan  7 17:41 Etc
drwxr-xr-x 2 root root   1320 Jan  7 17:41 Europe
...

```

`root #` `ls -l /usr/share/zoneinfo/Europe/`

```
total 256
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Amsterdam
-rw-r--r-- 1 root root 1742 Dec  3 17:19 Andorra
-rw-r--r-- 1 root root 1151 Dec  3 17:19 Astrakhan
-rw-r--r-- 1 root root 2262 Dec  3 17:19 Athens
-rw-r--r-- 1 root root 3664 Dec  3 17:19 Belfast
-rw-r--r-- 1 root root 1920 Dec  3 17:19 Belgrade
-rw-r--r-- 1 root root 2298 Dec  3 17:19 Berlin
-rw-r--r-- 1 root root 2301 Dec  3 17:19 Bratislava
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Brussels
...

```

Предположим, что необходимо установить часовой пояс Europe/Brussels:

`root #` `ln -sf ../usr/share/zoneinfo/Europe/Brussels /etc/localtime`

**Совет**

The target path with `../` at the start is _relative to the link location_, not the current directory.

**Заметка**

An absolute path can be used for the symlink, but a relative link is also created by systemd's timedatectl and is more compatible with alternate _ROOT_ s.

## Настройка локалей

**Заметка**

Этот раздел неприменим к пользователям musl libc. Если вы не знаете, что это такое, то к вам это не относится и вы должны следовать этому разделу.

### Генерация локалей

A default installation of Gentoo Linux provides an archive that contains all supported locales, numbering 500 or more. However, it is typical for an administrator to require only one or two of these. In that case, the /etc/locale.gen configuration file may be populated with a list of the required locales. By default, locale-gen shall read this file and compile only the locales that are specified, saving both time and space in the longer term.

Локаль указывает не только язык, который использует пользователь при взаимодействии с системой, но и правила для сортировки строк, формат вывода даты и времени, и так далее. Локали являются регистрозависимыми и должны использоваться так, как они описаны. Полный список доступных локалей можно найти в файле /usr/share/i18n/SUPPORTED.

`root #` `nano /etc/locale.gen`

Следующие локали являются примером для создания английской (США) и русской (Россия) локалей с поддержкой формата символов (например, UTF-8).

ФАЙЛ **`/etc/locale.gen`** **Включение US и RU локалей с поддержкой формата символов**

```
en_US.UTF-8 UTF-8
ru_RU.UTF-8 UTF-8

```

1. Non UTF-8 locales are discouraged and should only be used in special circumstances.
2. en\_US ISO-8859-1
3. de\_DE ISO-8859-1

}}

**Предупреждение**

Для сборки многих приложений наличие хотя бы одной локали с поддержкой UTF-8 является обязательным требованием.

Далее, запустим команду locale-gen, которая создаст все перечисленные в файле /etc/locale.gen локали.

`root #` `locale-gen`

Чтобы убедится, что выбранные локали теперь доступны, запустите команду locale -a.

На установках с systemd можно использовать localectl, т.е. команды localectl set-locale ... или localectl list-locales.

### Выбор локали

Теперь настало время установить локаль для всей системы. И снова используется eselect для этого, только теперь с модулем `locale`.

Команда eselect locale list выведет список доступных локалей:

`root #` `eselect locale list`

```
Available targets for the LANG variable:
  [1]  C
  [2]  C.UTF-8
  [3]  POSIX
  [4]  de_DE.UTF-8
  [5]  en_US.UTF-8
  [ ]  (free form)

```

Команда eselect locale set <NUMBER> установит необходимую локаль:

`root #` `eselect locale set 2`

Это всё ещё можно сделать вручную с помощью файла /etc/env.d/02locale (для systemd с помощью файла /etc/locale.conf):

ФАЙЛ **`/etc/env.d/02locale`** **Ручная настройка системной локали**

```
LANG="ru_RU.UTF-8"
LC_COLLATE="C.UTF-8"

```

Установка локали предотвратит появление предупреждений и ошибок в процессе компиляции ядра и программ.

Заново перезагрузите окружение:

`root #` `env-update && source /etc/profile && export PS1="(chroot) ${PS1}"`

Для получения дополнительных рекомендаций по выбору локали см. также [руководство по локализации](/wiki/Localization/Guide/ru "Localization/Guide/ru") и статью [UTF-8](/wiki/UTF-8/ru "UTF-8/ru").

## Примечания

1. [↑](#cite_ref-1)[https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html](https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html)
2. [↑](#cite_ref-2)[https://www.gentoo.org/glep/glep-0023.html#id7](https://www.gentoo.org/glep/glep-0023.html#id7)

[← Установка Gentoo файлов установки](/wiki/Handbook:Alpha/Installation/Stage/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:Alpha/ru "Handbook:Alpha/ru") [Настройка ядра Linux →](/wiki/Handbook:Alpha/Installation/Kernel/ru "Следующая часть")