# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbuch:HPPA/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Bootloader/es "Manual de Gentoo: HPPA/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Bootloader/it "Handbook:HPPA/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Bootloader/pl "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Bootloader/pt-br "Manual:HPPA/Instalação/Gerenciador de Boot (100% translated)")
- русский
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Bootloader/ta "கையேடு:HPPA/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Bootloader/zh-cn "手册：HPPA/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Bootloader/ja "ハンドブック:HPPA/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko (100% translated)")

[HPPA Handbook](/wiki/Handbook:HPPA/ru "Handbook:HPPA/ru")[Установка](/wiki/Handbook:HPPA/Full/Installation/ru "Handbook:HPPA/Full/Installation/ru")[Об установке](/wiki/Handbook:HPPA/Installation/About/ru "Handbook:HPPA/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:HPPA/Installation/Media/ru "Handbook:HPPA/Installation/Media/ru")[Настройка сети](/wiki/Handbook:HPPA/Installation/Networking/ru "Handbook:HPPA/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:HPPA/Installation/Disks/ru "Handbook:HPPA/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:HPPA/Installation/Stage/ru "Handbook:HPPA/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:HPPA/Installation/Base/ru "Handbook:HPPA/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:HPPA/Installation/Kernel/ru "Handbook:HPPA/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:HPPA/Installation/System/ru "Handbook:HPPA/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:HPPA/Installation/Tools/ru "Handbook:HPPA/Installation/Tools/ru")Настройка загрузчика[Завершение](/wiki/Handbook:HPPA/Installation/Finalizing/ru "Handbook:HPPA/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:HPPA/Full/Working/ru "Handbook:HPPA/Full/Working/ru")[Введение в Portage](/wiki/Handbook:HPPA/Working/Portage/ru "Handbook:HPPA/Working/Portage/ru")[USE-флаги](/wiki/Handbook:HPPA/Working/USE/ru "Handbook:HPPA/Working/USE/ru")[Возможности Portage](/wiki/Handbook:HPPA/Working/Features/ru "Handbook:HPPA/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:HPPA/Working/Initscripts/ru "Handbook:HPPA/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:HPPA/Working/EnvVar/ru "Handbook:HPPA/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:HPPA/Full/Portage/ru "Handbook:HPPA/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:HPPA/Portage/Files/ru "Handbook:HPPA/Portage/Files/ru")[Переменные](/wiki/Handbook:HPPA/Portage/Variables/ru "Handbook:HPPA/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:HPPA/Portage/Branches/ru "Handbook:HPPA/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:HPPA/Portage/Tools/ru "Handbook:HPPA/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:HPPA/Portage/CustomTree/ru "Handbook:HPPA/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:HPPA/Portage/Advanced/ru "Handbook:HPPA/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:HPPA/Full/Networking/ru "Handbook:HPPA/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:HPPA/Networking/Introduction/ru "Handbook:HPPA/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:HPPA/Networking/Advanced/ru "Handbook:HPPA/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:HPPA/Networking/Modular/ru "Handbook:HPPA/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:HPPA/Networking/Wireless/ru "Handbook:HPPA/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:HPPA/Networking/Extending/ru "Handbook:HPPA/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:HPPA/Networking/Dynamic/ru "Handbook:HPPA/Networking/Dynamic/ru")

## Установка PALO

Для платформы PA-RISC начальный загрузчик называется palo. Сначала установите этот загрузчик:

`root #` `emerge --ask sys-boot/palo`

Конфигурационный файл находится здесь /etc/palo.conf. Вот его пример:

**Предупреждение**

Данная конфигурация **должна** быть изменена после первого запуска palo! См. ниже для настройки после первого запуска.

ФАЙЛ **`/etc/palo.conf`** **Простой пример конфигурации PALO**

```
--commandline=2/kernel-6.19.1-gentoo root=/dev/sda4
--recoverykernel=/vmlinux.old
# УДАЛИТЕ эту строку после первого запуска palo!
--init-partitioned=/dev/sda
# --format-as имеет два значения, которые зависят от от того, используется ли --init-partitioned или --update-partitioned. Сохраните эту строку.
--format-as=4

```

Первая строка сообщает palo расположение ядра и параметры, с которыми оно должно загружаться. Запись `2/kernel-6.19.1-gentoo` означает, что ядро с названием /kernel-6.19.1-gentoo расположено во втором разделе. Будьте внимательны, путь к ядру относителен по отношению к загрузочному разделу, а не к корневому.

Вторая строка отображает используемое резервное ядро. Если это первая установка, и у вас нет резервного ядра (пока), пожалуйста, закомментируйте ее. Третья строка отображает раздел, на котором будет находиться palo.

При форматировании диска palo должен запускаться с определенными параметрами. В данном примере используется _ext4_ для первого раздела:

`root #` `palo --format-as=4 --init-partitioned=/dev/sda`

Как только настройка будет завершена, просто запустите команду palo.

`root #` `palo`

Для использования после установки необходимо обновить конфигурацию:

ФАЙЛ **`/etc/palo.conf`** **Simple PALO configuration example**

```
--commandline=2/kernel-6.19.1-gentoo root=/dev/sda4
--recoverykernel=/vmlinux.old
# Не удалять старый раздел, просто обновить существующий во время запуска `palo`.
--update-partitioned=/dev/sda
# --format-as имеет два значения, которые зависят от от того, используется ли --init-partitioned или --update-partitioned. Сохраните эту строку.
--format-as=4

```

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

Перезагрузившись в новое окружение Gentoo, переходите к [завершению установки Gentoo](/wiki/Handbook:HPPA/Installation/Finalizing/ru "Handbook:HPPA/Installation/Finalizing/ru").

[← Установка системных средств](/wiki/Handbook:HPPA/Installation/Tools/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:HPPA/ru "Handbook:HPPA/ru") [Завершение →](/wiki/Handbook:HPPA/Installation/Finalizing/ru "Следующая часть")