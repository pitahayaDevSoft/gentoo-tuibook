# Networking

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Networking/de "Handbuch:PPC/Installation/Netzwerk (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Networking "Handbook:PPC/Installation/Networking (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Networking/es "Manual de Gentoo: PPC/Instalación/Redes (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Networking/fr "Handbook:PPC/Installation/Networking/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Networking/it "Handbook:PPC/Installation/Networking/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Networking/hu "Handbook:PPC/Installation/Networking/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Networking/pl "Handbook:PPC/Installation/Networking (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Networking/pt-br "Handbook:PPC/Installation/Networking/pt-br (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:PPC/Installation/Networking/ta "கையேடு:PPC/நிறுவல்/வலையமைத்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Networking/zh-cn "手册：PPC/安装/配置网络 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Networking/ja "ハンドブック:PPC/インストール/ネットワーク (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Networking/ko "Handbook:PPC/Installation/Networking/ko (100% translated)")

[PPC Handbook](/wiki/Handbook:PPC/ru "Handbook:PPC/ru")[Установка](/wiki/Handbook:PPC/Full/Installation/ru "Handbook:PPC/Full/Installation/ru")[Об установке](/wiki/Handbook:PPC/Installation/About/ru "Handbook:PPC/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:PPC/Installation/Media/ru "Handbook:PPC/Installation/Media/ru")Настройка сети[Подготовка дисков](/wiki/Handbook:PPC/Installation/Disks/ru "Handbook:PPC/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:PPC/Installation/Stage/ru "Handbook:PPC/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:PPC/Installation/Base/ru "Handbook:PPC/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:PPC/Installation/Kernel/ru "Handbook:PPC/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:PPC/Installation/System/ru "Handbook:PPC/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:PPC/Installation/Tools/ru "Handbook:PPC/Installation/Tools/ru")[Настройка загрузчика](/wiki/Handbook:PPC/Installation/Bootloader/ru "Handbook:PPC/Installation/Bootloader/ru")[Завершение](/wiki/Handbook:PPC/Installation/Finalizing/ru "Handbook:PPC/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:PPC/Full/Working/ru "Handbook:PPC/Full/Working/ru")[Введение в Portage](/wiki/Handbook:PPC/Working/Portage/ru "Handbook:PPC/Working/Portage/ru")[USE-флаги](/wiki/Handbook:PPC/Working/USE/ru "Handbook:PPC/Working/USE/ru")[Возможности Portage](/wiki/Handbook:PPC/Working/Features/ru "Handbook:PPC/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:PPC/Working/Initscripts/ru "Handbook:PPC/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:PPC/Working/EnvVar/ru "Handbook:PPC/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:PPC/Full/Portage/ru "Handbook:PPC/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:PPC/Portage/Files/ru "Handbook:PPC/Portage/Files/ru")[Переменные](/wiki/Handbook:PPC/Portage/Variables/ru "Handbook:PPC/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:PPC/Portage/Branches/ru "Handbook:PPC/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:PPC/Portage/Tools/ru "Handbook:PPC/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:PPC/Portage/CustomTree/ru "Handbook:PPC/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:PPC/Portage/Advanced/ru "Handbook:PPC/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:PPC/Full/Networking/ru "Handbook:PPC/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:PPC/Networking/Introduction/ru "Handbook:PPC/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:PPC/Networking/Advanced/ru "Handbook:PPC/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:PPC/Networking/Modular/ru "Handbook:PPC/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:PPC/Networking/Wireless/ru "Handbook:PPC/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:PPC/Networking/Extending/ru "Handbook:PPC/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:PPC/Networking/Dynamic/ru "Handbook:PPC/Networking/Dynamic/ru")

## Contents

- [1Автоматическая конфигурация сети](#.D0.90.D0.B2.D1.82.D0.BE.D0.BC.D0.B0.D1.82.D0.B8.D1.87.D0.B5.D1.81.D0.BA.D0.B0.D1.8F_.D0.BA.D0.BE.D0.BD.D1.84.D0.B8.D0.B3.D1.83.D1.80.D0.B0.D1.86.D0.B8.D1.8F_.D1.81.D0.B5.D1.82.D0.B8)
  - [1.1Использование DHCP](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_DHCP)
  - [1.2Проверка сети](#.D0.9F.D1.80.D0.BE.D0.B2.D0.B5.D1.80.D0.BA.D0.B0_.D1.81.D0.B5.D1.82.D0.B8)
- [2Получение информации о сетевых интерфейсах](#.D0.9F.D0.BE.D0.BB.D1.83.D1.87.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B8.D0.BD.D1.84.D0.BE.D1.80.D0.BC.D0.B0.D1.86.D0.B8.D0.B8_.D0.BE_.D1.81.D0.B5.D1.82.D0.B5.D0.B2.D1.8B.D1.85_.D0.B8.D0.BD.D1.82.D0.B5.D1.80.D1.84.D0.B5.D0.B9.D1.81.D0.B0.D1.85)
- [3Настройка беспроводного доступа](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D0.B1.D0.B5.D1.81.D0.BF.D1.80.D0.BE.D0.B2.D0.BE.D0.B4.D0.BD.D0.BE.D0.B3.D0.BE_.D0.B4.D0.BE.D1.81.D1.82.D1.83.D0.BF.D0.B0)
  - [3.1Дополнительно: Проверка беспроводного доступа](#.D0.94.D0.BE.D0.BF.D0.BE.D0.BB.D0.BD.D0.B8.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE:_.D0.9F.D1.80.D0.BE.D0.B2.D0.B5.D1.80.D0.BA.D0.B0_.D0.B1.D0.B5.D1.81.D0.BF.D1.80.D0.BE.D0.B2.D0.BE.D0.B4.D0.BD.D0.BE.D0.B3.D0.BE_.D0.B4.D0.BE.D1.81.D1.82.D1.83.D0.BF.D0.B0)
  - [3.2Использование NetworkManager](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_NetworkManager)
  - [3.3Использование net-setup](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_net-setup)
  - [3.4Manual setup](#Manual_setup)
    - [3.4.1Manually connecting to an open network](#Manually_connecting_to_an_open_network)
- [4Настройка конкретных приложений](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D0.BA.D0.BE.D0.BD.D0.BA.D1.80.D0.B5.D1.82.D0.BD.D1.8B.D1.85_.D0.BF.D1.80.D0.B8.D0.BB.D0.BE.D0.B6.D0.B5.D0.BD.D0.B8.D0.B9)
  - [4.1Настройка веб-прокси](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D0.B2.D0.B5.D0.B1-.D0.BF.D1.80.D0.BE.D0.BA.D1.81.D0.B8)
  - [4.2Использование pppoe-setup для ADSL](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_pppoe-setup_.D0.B4.D0.BB.D1.8F_ADSL)
  - [4.3Использование PPTP](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_PPTP)
- [5Интернет и основы IP-адресации](#.D0.98.D0.BD.D1.82.D0.B5.D1.80.D0.BD.D0.B5.D1.82_.D0.B8_.D0.BE.D1.81.D0.BD.D0.BE.D0.B2.D1.8B_IP-.D0.B0.D0.B4.D1.80.D0.B5.D1.81.D0.B0.D1.86.D0.B8.D0.B8)
  - [5.1Интерфейсы и адреса](#.D0.98.D0.BD.D1.82.D0.B5.D1.80.D1.84.D0.B5.D0.B9.D1.81.D1.8B_.D0.B8_.D0.B0.D0.B4.D1.80.D0.B5.D1.81.D0.B0)
  - [5.2Сети и CIDR](#.D0.A1.D0.B5.D1.82.D0.B8_.D0.B8_CIDR)
  - [5.3Интернет](#.D0.98.D0.BD.D1.82.D0.B5.D1.80.D0.BD.D0.B5.D1.82)
  - [5.4Система доменных имён (DNS)](#.D0.A1.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D0.B0_.D0.B4.D0.BE.D0.BC.D0.B5.D0.BD.D0.BD.D1.8B.D1.85_.D0.B8.D0.BC.D1.91.D0.BD_.28DNS.29)
- [6Ручная конфигурация сети с использованием статического IP](#.D0.A0.D1.83.D1.87.D0.BD.D0.B0.D1.8F_.D0.BA.D0.BE.D0.BD.D1.84.D0.B8.D0.B3.D1.83.D1.80.D0.B0.D1.86.D0.B8.D1.8F_.D1.81.D0.B5.D1.82.D0.B8_.D1.81_.D0.B8.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5.D0.BC_.D1.81.D1.82.D0.B0.D1.82.D0.B8.D1.87.D0.B5.D1.81.D0.BA.D0.BE.D0.B3.D0.BE_IP)
  - [6.1Конфигурация адреса интерфейса](#.D0.9A.D0.BE.D0.BD.D1.84.D0.B8.D0.B3.D1.83.D1.80.D0.B0.D1.86.D0.B8.D1.8F_.D0.B0.D0.B4.D1.80.D0.B5.D1.81.D0.B0_.D0.B8.D0.BD.D1.82.D0.B5.D1.80.D1.84.D0.B5.D0.B9.D1.81.D0.B0)
  - [6.2Настройка маршрута по умолчанию](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D0.BC.D0.B0.D1.80.D1.88.D1.80.D1.83.D1.82.D0.B0_.D0.BF.D0.BE_.D1.83.D0.BC.D0.BE.D0.BB.D1.87.D0.B0.D0.BD.D0.B8.D1.8E)
  - [6.3Конфигурация DNS](#.D0.9A.D0.BE.D0.BD.D1.84.D0.B8.D0.B3.D1.83.D1.80.D0.B0.D1.86.D0.B8.D1.8F_DNS)

## Автоматическая конфигурация сети

Может быть, всё уже работает?

Если система подключена к сети Ethernet, в которой есть IPv6-маршрутизатор или DHCP-сервер, весьма вероятно, что системная конфигурация сети была выполнена автоматически. Если расширенная настройка не требуется, можно сразу перейти к [проверке подключения к Интернету](/wiki/Handbook:PPC/Installation/Networking/ru#Testing_the_network "Handbook:PPC/Installation/Networking/ru").

### Использование DHCP

DHCP (Dynamic Host Configuration Protocol) облегчает настройку сети и позволяет автоматически настраивать различные параметры, включая: IP-адрес, маску сети, маршруты, DNS-сервера, NTP-сервера и так далее.

Для работы DHCP требуется, чтобы сервер работал в том же _Layer 2_ ( _Ethernet_) сегменте, что и клиент, запрашивающий _аренду_. DHCP часто используется в RFC1918 ( _частных_) сетях, а также в сетях провайдеров, чтобы запросить публичный IP-адрес.

By default the Gentoo live media uses [NetworkManager](/wiki/NetworkManager "NetworkManager") which will automatically connect to via DHCP, if this is not the case then check the Ethernet cable for issues.

### Проверка сети

Правильная настройка маршрута _по умолчанию_ ( _default_) критически важна для подключения к Интернету. Конфигурацию маршрута можно проверить с помощью следующей команды:

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

Если маршрут _по умолчанию_ не существует, Интернет-подключение не будет работать, и понадобится дополнительная конфигурация.

Простую проверку соединения с Интернетом можно провести утилитой ping:

`root #` `ping -c 3 1.1.1.1`

**Совет**

Лучше всего начать с проверки известного IP-адреса, а не доменного имени. Это поможет отделить проблемы с DNS от общих проблем с подключением к Интернету.

Исходящий доступ по HTTPS и разрешение доменных имён в DNS можно проверить с помощью:

`root #` `curl --location gentoo.org --output /dev/null`

Если проверки curl (и другие) завершились успешно, можно продолжить установку с [подготовки дисков](/wiki/Handbook:PPC/Installation/Disks/ru "Handbook:PPC/Installation/Disks/ru").

Если curl сообщил об ошибке, но ping других серверов в Интернете работает, возможно, требуется [настроить DNS](/wiki/Handbook:PPC/Installation/Networking/ru#DNS_configuration "Handbook:PPC/Installation/Networking/ru").

Если Интернет-подключение и вовсе не было установлено, сначала [необходимо проверить информацию о сетевых интерфейсах](/wiki/Handbook:PPC/Installation/Networking/ru#Obtaining_interface_info "Handbook:PPC/Installation/Networking/ru"), затем:

- Для упрощения процесса сетевой конфигурации [можно использовать nmtui](/wiki/Handbook:PPC/Installation/Networking/ru#Using_NetworkManager "Handbook:PPC/Installation/Networking/ru").
- Возможно, понадобится [настроить некоторые приложения](/wiki/Handbook:PPC/Installation/Networking/ru#Application_specific_configuration "Handbook:PPC/Installation/Networking/ru").
- Также можно попытаться [сконфигурировать сеть вручную](/wiki/Handbook:PPC/Installation/Networking/ru#Manual_network_configuration "Handbook:PPC/Installation/Networking/ru").

## Получение информации о сетевых интерфейсах

Если сеть не работает «из коробки», необходимо выполнить дополнительные шаги, чтобы обеспечить подключение к Интернету. Как правило, в начале необходимо определить список сетевых интерфейсов узла.

Для запроса и изменения системных настроек сети можно использовать команду ip из пакета [sys-apps/iproute2](https://packages.gentoo.org/packages/sys-apps/iproute2).

Чтобы получить список сетевых интерфейсов и их связи, используйте подкоманду **link**:

`root #` `ip link`

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
4: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
```

Чтобы получить информацию об адресе устройства, используйте подкоманду **address**:

`root #` `ip address`

```
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000<pre>
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
    inet 10.0.20.77/22 brd 10.0.23.255 scope global enp1s0
       valid_lft forever preferred_lft forever
    inet6 fe80::ea40:f2ff:feac:257a/64 scope link
       valid_lft forever preferred_lft forever
```

Эта команда выводит информацию для каждого сетевого интерфейса в системе. Первым идёт порядковый номер устройства, затем его название enp1s0.

**Совет**

Если в выводе ifconfig нет никаких интерфейсов, кроме **lo** ( _loopack_), это означает, что сетевое оборудование неисправно, либо драйвер для сетевого интерфейса не был загружен в ядро. Решение обеих проблем выходит за рамки данного Руководства. В этом случае, пожалуйста, попросите о помощи в сообществе [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)).

Для единообразия и упрощения, в Руководстве главный сетевой интерфейс называется enp1s0 в качестве примера.

**Заметка**

В результате перехода на [предсказуемые имена для сетевых интерфейсов](https://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/) название интерфейса может отличаться от старого соглашения о именовании (eth0). Современные загрузочные носители Gentoo используют префиксы при именовании сетевых интерфейсов, например: eno0, ens1 или enp5s0.

## Настройка беспроводного доступа

### Дополнительно: Проверка беспроводного доступа

iw may be used to display the current wireless settings on the card, this should also show that the kernel wireless stack is working (note that the iw command is only available on the following architectures: **amd64**, **x86**, **arm**, **arm64**, **ppc**, **ppc64**, and **riscv**):

`root #` `iw dev wlp9s0 info`

```
Interface wlp9s0
	ifindex 3
	wdev 0x1
	addr 00:00:00:00:00:00
	type managed
	wiphy 0
	channel 11 (2462 MHz), width: 20 MHz (no HT), center1: 2462 MHz
	txpower 30.00 dBm

```

**Заметка**

Some wireless cards may have a device name of wlan0 or ra0 instead of wlp9s0. Run ip link to determine the correct device name.

### Использование NetworkManager

The fastest way to enable wireless connectivity in the Gentoo Live media is by using the nmtui command and following the on-screen instructions to setup the wireless network.

`root #` `nmtui`

LiveGUI users can click on the network icon in the bottom right of the taskbar to also configure WiFi.

### Использование net-setup

В случае, если автоматическая настройка сети не удалась, для упрощения конфигурации сети в _загрузочный носитель_ Gentoo включён сценарий net-setup, который также можно использовать для настройки информации о беспроводной сети и статических IP-адресов.

`root #` `killall dhcpcd`

`root #` `net-setup enp1s0`

net-setup задаст несколько вопросов о сетевом окружении и с помощью этой информации настроит wpa\_supplicant или _статическую адресацию_.

**Важно**

Рекомендуется [проверить](/wiki/Handbook:PPC/Installation/Networking/ru#Testing_the_network "Handbook:PPC/Installation/Networking/ru") статус сети после выполнения любых конфигурационных действий. В случае, если сценарий настройки не сработает, придётся [настроить сеть вручную](/wiki/Handbook:PPC/Installation/Networking/ru#Manual_network_configuration "Handbook:PPC/Installation/Networking/ru").

### Manual setup

Wireless networks may alternatively be set up with a daemon such as wpa\_supplicant or iwd. For more information on configuring wireless networking in Gentoo Linux, please read the [Wireless networking chapter](/wiki/Handbook:PPC/Networking/Wireless/ru "Handbook:PPC/Networking/Wireless/ru") in the Gentoo Handbook.}}

#### Manually connecting to an open network

For most users, there are only two settings needed to connect, the ESSID (aka wireless network name) and, optionally, the WEP key.

- First, ensure the interface is active:

`root #` `ip link set dev wlp9s0 up`

- To connect to an open network with the name _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## Настройка конкретных приложений

**Совет**

These steps are provided for users where using nmtui is not able to configure their network's needs.

Приведённые ниже способы обычно необязательны, но могут быть полезны в случаях, когда для подключения к Интернету необходимо предпринимать дополнительную конфигурацию.

### Настройка веб-прокси

Если доступ в Интернет осуществляется через веб-прокси, то информацию о каждом поддерживаемом протоколе необходимо задать Portage. Portage учитывает переменные окружения http\_proxy, ftp\_proxy и RSYNC\_PROXY во время скачивания пакетов через механизмы wget и rsync.

Некоторые текстовые веб-браузеры, например links, тоже могут использовать переменные окружения, определяющие настройки веб-прокси. В частности, для доступа по протоколу HTTPS для них потребуется определить переменную окружения https\_proxy. В Portage переменные окружения подхватываются автоматически, без дополнительных аргументов к команде, но для links могут потребоваться дополнительные параметры.

В большинстве случаев достаточно задать переменные окружения, используя имя узла сервера. В следующем примере предполагается, что прокси-сервер находится по адресу proxy.gentoo.org и доступен по порту 8080.

**Заметка**

Символы `#` в следующих командах означают комментарии. Они используются только для ясности, и их _необязательно_ писать во время запуска команд.

Чтобы задать HTTP-прокси (для HTTP- и HTTPS-трафика):

`root #` `export http_proxy="http://proxy.gentoo.org:8080" # Для Portage и Links
`

`root #` `export https_proxy="http://proxy.gentoo.org:8080" # Только для Links
`

Если для HTTP-прокси требуется аутентификация, установить имя пользователя и пароль с помощью следующей команды:

`root #` `export http_proxy="http://username:password@proxy.gentoo.org:8080" # Для Portage и Links
`

`root #` `export https_proxy="http://username:password@proxy.gentoo.org:8080" # Только для Links
`

Запустите links со следующими параметрами, чтобы включить поддержку прокси:

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

Чтобы задать FTP-прокси для Portage и/или links:

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080" # Для Portage и Links`

Запустите links со следующими параметрами для поддержки FTP-прокси:

`user $` `links -ftp-proxy ${ftp_proxy} `

Чтобы задать RSYNC-прокси для Portage:

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080" # Только для Portage; Links не поддерживает rsync-прокси`

### Использование pppoe-setup для ADSL

В состав _загрузочного носителя_ Gentoo входит сценарий pppoe-setup, упрощающий настройку ppp, в случаях, когда для доступа в Интернет используется PPPoE.

Во время настройки pppoe-setup попросит предоставить:

- Название Ethernet- _интерфейса_, подключенного к ADSL-модему.
- Имя пользователя и пароль PPPoE.
- IP-адреса DNS-сервера.
- Нужно ли использовать межсетевой экран.

`root #` `pppoe-setup
`

`root #` `pppoe-start`

В случае ошибки следует проверить учётные данные в /etc/ppp/pap-secrets или /etc/ppp/chap-secrets. Если они верны, следует проверить, правильно ли выбран Ethernet-интерфейс для PPPoE.

### Использование PPTP

Если нужна поддержка PPTP, можно использовать утилиту pptpclient, но для неё требуется настройка перед запуском.

Отредактируйте /etc/ppp/pap-secrets or /etc/ppp/chap-secrets так, чтобы файл содержал правильный логин и пароль:

`root #` `nano /etc/ppp/chap-secrets`

При необходимости подправьте /etc/ppp/options.pptp:

`root #` `nano /etc/ppp/options.pptp`

После завершения настройки запустите pptp (вместе с параметрами, которые невозможно установить в options.pptp) для подключения к серверу:

`root #` `pptp <server ipv4 address>`

## Интернет и основы IP-адресации

Если все способы выше провалились, необходимо настроить сеть вручную. Это относительно несложно, но нужно делать всё осторожно. Этот раздел призван разъяснить терминологию и ознакомить пользователя с основными концепциями сети, которые будут встречаться во время ручной конфигурации подключения к Интернету.

**Совет**

Некоторые устройства, предоставляемые провайдером, могут одновременно совмещать в себе функции _маршрутизатора_, _точки доступа_, _модема_, _DHCP-сервера_ и _DNS-сервера_. Важно отличать функции устройства от его физического представления.

### Интерфейсы и адреса

Сетевые _интерфейсы_ — это логические представления сетевых устройств. _Интерфейсу_ нужен _адрес_ для связи с другими устройствами в _сети_. _Интерфейс_ обязан иметь хотя бы один _адрес_, однако он может иметь и несколько адресов одновременно. Это особенно полезно для параллельного использования IPv4 и IPv6.

Для единообразия в данном примере предполагается, что используется интерфейс enp1s0, которому будет назначен адрес 192.168.0.2.

**Важно**

IP-адреса могут быть заданы произвольно, из-за чего разные устройства могут использовать одинаковый IP-адрес, что приведёт к _конфликту адресов_. Конфликты адресов следует избегать, используя DHCP или SLAAC.

**Совет**

В IPv6 обычно используется **S** tate **L** ess **A** ddress **A** uto **C** onfiguration (SLAAC) для настройки адреса. В большинстве случаев не рекомендуется задавать IPv6-адрес вручную. Если нужно задать конкретный суффикс адреса, можно использовать [идентификационные токены интерфейса](/wiki/IPv6_Static_Addresses_using_Tokens "IPv6 Static Addresses using Tokens").

### Сети и CIDR

Откуда устройство знает, как общаться с другими устройствами после получения адреса?

IP- _адреса_ связаны с понятием _сети_. IP- _сеть_ — это последовательный логический диапазон адресов.

_Бесклассовая адресация_ ( _Classless Inter-Domain Routing_ или _CIDR_) используется для определения размера сети.

- Значение _CIDR_, обычно начинающееся с **/**, определяет вместимость сети.

  - Формулой расчёта вместимости сети является _2 ^ (32 - CIDR)_.
  - После подсчёта вместимости сети количество доступных адресов вычисляется путем вычитания 2 из вместимости сети.
    - Первый IP-адрес сети является _сетевым адресом_, а последний — обычно _широковещательным адресом_. Эти адреса — особенные и не могут использоваться для адресации обычных узлов.

**Совет**

Наиболее часто встречающимися значениями _CIDR_ являются **/24** и **/32**, вместимость которых равна **254** узла и один узел соответственно.

_CIDR_, равный **/24**, де-факто является размером сети по умолчанию. Этому значения соответствует маска подсети _255.255.255.0_, где последние 8 бит отведены для IP-адресов узлов в этой сети.

Запись 192.168.0.2/24 может трактоваться следующим образом:

- _Адрес_ узла — 192.168.0.2
- _Сеть_ — 192.168.0.0
- Вместимость сети — **254** ( _2 ^ (32 - 24) \- 2_)

  - Диапазон возможных для использования адресов — 192.168.0.1 - 192.168.0.254
- Широковещательный адрес — 192.168.0.255
  - В большинстве случаев последний адрес сети используется как _широковещательный адрес_, но он может быть задан другим.

Используя данную конфигурацию, устройство может общаться с любым узлом в этой же _сети_ (192.168.0.0).

### Интернет

Как только устройство оказывается подключённым к сети, как оно определяет возможность взаимодействия с другими устройствами в Интернете?

Для общения с другими устройствам вне локальной _сети_ необходимо использовать _маршрутизацию_. _Маршрутизатор_ — это простое устройство сети, которое перебрасывает сетевой трафик других устройств. Термин _маршрут по умолчанию_ или _шлюз_ обычно описывает устройство в текущей сети, которое используется для доступа к внешним сетям.

**Совет**

Распространённой практикой является использование первого или последнего IP-адреса сети для назначения _шлюзу_.

Если подключённый к Интернету маршрутизатор доступен по адресу 192.168.0.1, то он может быть использован в качестве _маршрута по умолчанию_ для доступа к Интернету.

В итоге:

- Сетевые интерфейсы должны быть сконфигурированы с использованием _адреса_ и _информации о сети_, например _CIDR_.
- Доступ к локальной сети используется для подключения к _маршрутизатору_ этой сети.
- При наличии _маршрута по умолчанию_ весь сетевой трафик, предназначенный для внешних сетей, будет перенаправлен _шлюзу_, который предоставляет доступ к Интернету.

### Система доменных имён (DNS)

Запоминание IP-адресов является сложным и неудобным. Для обеспечения связи между _доменными именами_ и _IP-адресами_ была создана _система доменных имён_ ( _Domain Name System_).

В системах Linux используется файл /etc/resolv.conf для определения _серверов имён_ ( _nameservers_), которые участвуют в _разрешении (резолвинге) DNS-имён_ ( _DNS resolution_).

**Совет**

Множество _маршрутизаторов_ также функционируют как DNS-серверы; использование локального DNS-сервера улучшает приватность и ускоряет запросы благодаря кэшированию.

Множество провайдеров запускают собственные DNS-сервера, информация о которых обычно распространяется от _шлюза_ через протокол DHCP. Использование локального DNS-сервера улучшает время отклика, но большинство публичных DNS-серверов возвращают те же результаты, так что предпочтение в выборе сервера остаётся за пользователем.

## Ручная конфигурация сети с использованием статического IP

### Конфигурация адреса интерфейса

**Важно**

При ручной конфигурации IP-адреса следует учитывать локальную топологию сети. IP-адреса могут быть уже занятыми; конфликт адресов может привести к нарушению работы сети.

Чтобы сконфигурировать интерфейс enp1s0 с адресом 192.168.0.2 и CIDR /24:

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**Совет**

Начало этой команды может быть сокращено до ip a.

### Настройка маршрута по умолчанию

Конфигурация адреса и информации о сети для интерфейса задаёт маршруты связи (link), благодаря которым становится возможным взаимодействие с этим сегментом сети:

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**Совет**

Начало этой команды может быть сокращено до ip r.

Маршрут по умолчанию (default) 192.168.0.1 может быть задан следующей командой:

`root #` `ip route add default via 192.168.0.1`

### Конфигурация DNS

Информация о серверах имён обычно распространяется через DHCP, но адреса серверов можно задать вручную с помощью записей `nameserver` в /etc/resolv.conf.

**Предупреждение**

Если запущен сервис _dhcpcd_, изменения в /etc/resolv.conf не сохранятся. Проверить статус можно с помощью команды `ps x | grep dhcpcd`.

Текстовый редактор nano включён в _загрузочный носитель_ Gentoo и может быть использован для правки /etc/resolv.conf:

`root #` `nano /etc/resolv.conf`

Строки, содержащие ключевое слово `nameserver`, за которым следует IP-адрес DNS-сервера, используются в порядке определения:

ФАЙЛ **`/etc/resolv.conf`** **Использование DNS Quad9.**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

ФАЙЛ **`/etc/resolv.conf`** **Использование DNS Cloudflare.**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

Состояние DNS может быть проверено с помощью ping доменного имени:

`root #` `ping -c 3 gentoo.org`

Как только подключение [было успешно проверено](/wiki/Handbook:PPC/Installation/Networking/ru#Testing_the_network "Handbook:PPC/Installation/Networking/ru"), продолжайте с раздела [«Подготовка дисков»](/wiki/Handbook:PPC/Installation/Disks/ru "Handbook:PPC/Installation/Disks/ru").

[← Выбор подходящего источника для установки](/wiki/Handbook:PPC/Installation/Media/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:PPC/ru "Handbook:PPC/ru") [Подготовка дисков →](/wiki/Handbook:PPC/Installation/Disks/ru "Следующая часть")