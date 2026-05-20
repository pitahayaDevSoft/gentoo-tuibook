# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Bootloader/de "Handbuch:Alpha/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Bootloader "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Bootloader/es "Manual de Gentoo: Alpha/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Bootloader/fr "Manuel:Alpha/Installation/Système d'amorçage (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Bootloader/it "Handbook:Alpha/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Bootloader/hu "Handbook:Alpha/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Bootloader/pl "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Bootloader/pt-br "Manual:Alpha/Instalação/Gerenciador de boot (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Bootloader/cs "Handbook:Alpha/Installation/Bootloader/cs (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Bootloader/ta "கையேடு:Alpha/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Bootloader/zh-cn "手册：Alpha/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Bootloader/ja "ハンドブック:Alpha/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Bootloader/ko "Handbook:Alpha/Installation/Bootloader/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha/ru "Handbook:Alpha/ru")[Установка](/wiki/Handbook:Alpha/Full/Installation/ru "Handbook:Alpha/Full/Installation/ru")[Об установке](/wiki/Handbook:Alpha/Installation/About/ru "Handbook:Alpha/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:Alpha/Installation/Media/ru "Handbook:Alpha/Installation/Media/ru")[Настройка сети](/wiki/Handbook:Alpha/Installation/Networking/ru "Handbook:Alpha/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:Alpha/Installation/Disks/ru "Handbook:Alpha/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:Alpha/Installation/Stage/ru "Handbook:Alpha/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:Alpha/Installation/Base/ru "Handbook:Alpha/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:Alpha/Installation/Kernel/ru "Handbook:Alpha/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:Alpha/Installation/System/ru "Handbook:Alpha/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:Alpha/Installation/Tools/ru "Handbook:Alpha/Installation/Tools/ru")Настройка загрузчика[Завершение](/wiki/Handbook:Alpha/Installation/Finalizing/ru "Handbook:Alpha/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:Alpha/Full/Working/ru "Handbook:Alpha/Full/Working/ru")[Введение в Portage](/wiki/Handbook:Alpha/Working/Portage/ru "Handbook:Alpha/Working/Portage/ru")[USE-флаги](/wiki/Handbook:Alpha/Working/USE/ru "Handbook:Alpha/Working/USE/ru")[Возможности Portage](/wiki/Handbook:Alpha/Working/Features/ru "Handbook:Alpha/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:Alpha/Working/Initscripts/ru "Handbook:Alpha/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:Alpha/Working/EnvVar/ru "Handbook:Alpha/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:Alpha/Full/Portage/ru "Handbook:Alpha/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:Alpha/Portage/Files/ru "Handbook:Alpha/Portage/Files/ru")[Переменные](/wiki/Handbook:Alpha/Portage/Variables/ru "Handbook:Alpha/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:Alpha/Portage/Branches/ru "Handbook:Alpha/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:Alpha/Portage/Tools/ru "Handbook:Alpha/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:Alpha/Portage/CustomTree/ru "Handbook:Alpha/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:Alpha/Portage/Advanced/ru "Handbook:Alpha/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:Alpha/Full/Networking/ru "Handbook:Alpha/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:Alpha/Networking/Introduction/ru "Handbook:Alpha/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:Alpha/Networking/Advanced/ru "Handbook:Alpha/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:Alpha/Networking/Modular/ru "Handbook:Alpha/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:Alpha/Networking/Wireless/ru "Handbook:Alpha/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:Alpha/Networking/Extending/ru "Handbook:Alpha/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:Alpha/Networking/Dynamic/ru "Handbook:Alpha/Networking/Dynamic/ru")

## Делаем выбор

Сейчас когда ядро настроено и скомпилировано и подготовлены необходимые системные файлы, наступило время установить программу, которая будет запускать ядро при запуске системы. Такая программа называется начальным загрузчиком.

Существуют несколько загрузчиков для ОС Linux на платформе Alpha. Вам нужно выбрать один. У вас есть выбор между [aBoot](#Default:_Using_aboot) и [MILO](#Alternative:_Using_MILO).

## По умолчанию: Использование aboot

**Заметка**

aboot поддерживает загрузку только с разделов ext2 и ext3.

Во-первых установим программу aboot на нашу систему.

`root #` `emerge --ask sys-boot/aboot`

Следующий шаг \- это сделать наш диск загрузочным. Эта процедура будет запускать программу aboot, когда система начнет загружаться. Мы сделаем это записав программу aboot в самое начало диска.

`root #` `swriteboot -f3 /dev/sda /boot/bootlx
`

`root #` `abootconf /dev/sda 2
`

**Заметка**

Если используется другая схема разбивки жесткого диска чем используется в этой книге, соответственно необходимо изменить команды. Пожалуйста прочитайте соответствующие страницы в документации (man 8 swriteboot и man 8 abootconf). Также если используется файловая система JFS в качестве корневой файловой файловой системы, убедитесь что изначально она монтируется только-для-чтения, добавив ro к параметрам ядра.

Хотя aboot теперь установлен, нам все еще нужно написать для него файл конфигурации. aboot требует только одной строки на каждую конфигурацию, поэтому мы можем сделать следующее:

`root #` `echo '0:2/boot/vmlinux.gz root=/dev/sda3' > /etc/aboot.conf`

Если в процессе компиляции ядра Linux вы решили включить initramfs для загрузки, то вам нужно изменить файл конфигурации, чтобы указать этот файл initramfs, а также сказать initramfs, где на самом деле находится устройство с корневым разделом.

`root #` `echo '0:2/boot/vmlinux.gz initrd=/boot/initramfs-genkernel-alpha-6.19.1-gentoo root=/dev/sda3' > /etc/aboot.conf`

В дополнение заметим, что вы можете загружать систему Gentoo автоматически установкой некоторых переменных SRM. Вы можете попробовать установить эти переменные из среды Linux, но будет намного проще выполнить это из консоли SRM.

`root #` `cd /proc/srm_environment/named_variables
`

`root #` `echo -n 0 > boot_osflags
`

`root #` `echo -n '' > boot_file
`

`root #` `echo -n 'BOOT' > auto_action
`

`root #` `echo -n 'dkc100' > bootdef_dev`

Естественно, замените dkc100 на ваше загрузочное устройство

Если вам потребуется опять попасть в программу SRM консоли (чтобы восстановить систему, проверить переменные или еще зачем-то), просто нажмите `Ctrl+C` для прерывания процесса аавтоматической загрузки.

Если происходит установка с использованием последовательной консоли, не забудьте включить флаг загрузки последовательной консоли в файле aboot.conf. Для примера как это сделать просмотрите файл /etc/aboot.conf.example.

Программа загрузчика Aboot настроена и готова к работе. Продолжайте по ссылке [Перезагрузка системы](#Rebooting_the_system).

Теперь продолжайте с раздела [Перезагрузка системы](#Rebooting_the_system).

## Перезагрузка системы

Выйдите из изолированной среды и размонтируйте все смонтированные разделы. Затем введите ту самую волшебную команду, которая запускает последний, настоящий тест: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Не забудьте извлечь загрузочный компакт-диск, иначе он может загрузиться снова вместо новой системы Gentoo!

Перезагрузившись в новое окружение Gentoo, переходите к [завершению установки Gentoo](/wiki/Handbook:Alpha/Installation/Finalizing/ru "Handbook:Alpha/Installation/Finalizing/ru").

[← Установка системных средств](/wiki/Handbook:Alpha/Installation/Tools/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:Alpha/ru "Handbook:Alpha/ru") [Завершение →](/wiki/Handbook:Alpha/Installation/Finalizing/ru "Следующая часть")