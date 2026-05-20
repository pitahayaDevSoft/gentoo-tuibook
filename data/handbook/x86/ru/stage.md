# Stage

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Stage/de "Handbuch:X86/Installation/Stage (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Stage "Handbook:X86/Installation/Stage (100% translated)")
- [Türkçe](/wiki/Handbook:X86/Installation/Stage/tr "Handbook:X86/Installation/Stage/tr (0% translated)")
- [español](/wiki/Handbook:X86/Installation/Stage/es "Manual de Gentoo: X86/Instalación/Stage (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Stage/fr "Handbook:X86/Installation/Stage/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Stage/it "Handbook:X86/Installation/Stage (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/Stage/hu "Handbook:X86/Installation/Stage/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Stage/pl "Handbook:X86/Installation/Stage (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Stage/pt-br "Manual:X86/Instalação/Stage (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Stage/cs "Handbook:X86/Installation/Stage/cs (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:X86/Installation/Stage/ta "கையேடு:X86/நிறுவல்/நிலை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Stage/zh-cn "手册：X86/安装/安装stage3 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Stage/ja "ハンドブック:X86/インストール/Stage (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Stage/ko "Handbook:X86/Installation/Stage/ko (100% translated)")

[X86 Handbook](/wiki/Handbook:X86/ru "Handbook:X86/ru")[Установка](/wiki/Handbook:X86/Full/Installation/ru "Handbook:X86/Full/Installation/ru")[Об установке](/wiki/Handbook:X86/Installation/About/ru "Handbook:X86/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:X86/Installation/Media/ru "Handbook:X86/Installation/Media/ru")[Настройка сети](/wiki/Handbook:X86/Installation/Networking/ru "Handbook:X86/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:X86/Installation/Disks/ru "Handbook:X86/Installation/Disks/ru")Установка файла stage[Установка базовой системы](/wiki/Handbook:X86/Installation/Base/ru "Handbook:X86/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:X86/Installation/Kernel/ru "Handbook:X86/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:X86/Installation/System/ru "Handbook:X86/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:X86/Installation/Tools/ru "Handbook:X86/Installation/Tools/ru")[Настройка загрузчика](/wiki/Handbook:X86/Installation/Bootloader/ru "Handbook:X86/Installation/Bootloader/ru")[Завершение](/wiki/Handbook:X86/Installation/Finalizing/ru "Handbook:X86/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:X86/Full/Working/ru "Handbook:X86/Full/Working/ru")[Введение в Portage](/wiki/Handbook:X86/Working/Portage/ru "Handbook:X86/Working/Portage/ru")[USE-флаги](/wiki/Handbook:X86/Working/USE/ru "Handbook:X86/Working/USE/ru")[Возможности Portage](/wiki/Handbook:X86/Working/Features/ru "Handbook:X86/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:X86/Working/Initscripts/ru "Handbook:X86/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:X86/Working/EnvVar/ru "Handbook:X86/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:X86/Full/Portage/ru "Handbook:X86/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:X86/Portage/Files/ru "Handbook:X86/Portage/Files/ru")[Переменные](/wiki/Handbook:X86/Portage/Variables/ru "Handbook:X86/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:X86/Portage/Branches/ru "Handbook:X86/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:X86/Portage/Tools/ru "Handbook:X86/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:X86/Portage/CustomTree/ru "Handbook:X86/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:X86/Portage/Advanced/ru "Handbook:X86/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:X86/Full/Networking/ru "Handbook:X86/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:X86/Networking/Introduction/ru "Handbook:X86/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:X86/Networking/Advanced/ru "Handbook:X86/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:X86/Networking/Modular/ru "Handbook:X86/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:X86/Networking/Wireless/ru "Handbook:X86/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:X86/Networking/Extending/ru "Handbook:X86/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:X86/Networking/Dynamic/ru "Handbook:X86/Networking/Dynamic/ru")

## Contents

- [1Выбор архива stage](#.D0.92.D1.8B.D0.B1.D0.BE.D1.80_.D0.B0.D1.80.D1.85.D0.B8.D0.B2.D0.B0_stage)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
  - [1.3Multilib (32 и 64 бит)](#Multilib_.2832_.D0.B8_64_.D0.B1.D0.B8.D1.82.29)
  - [1.4No-multilib (чистый 64-bit)](#No-multilib_.28.D1.87.D0.B8.D1.81.D1.82.D1.8B.D0.B9_64-bit.29)
- [2Скачивание архива stage](#.D0.A1.D0.BA.D0.B0.D1.87.D0.B8.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_.D0.B0.D1.80.D1.85.D0.B8.D0.B2.D0.B0_stage)
  - [2.1Установка даты и времени](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_.D0.B4.D0.B0.D1.82.D1.8B_.D0.B8_.D0.B2.D1.80.D0.B5.D0.BC.D0.B5.D0.BD.D0.B8)
    - [2.1.1Автоматическая настройка](#.D0.90.D0.B2.D1.82.D0.BE.D0.BC.D0.B0.D1.82.D0.B8.D1.87.D0.B5.D1.81.D0.BA.D0.B0.D1.8F_.D0.BD.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0)
    - [2.1.2Ручная настройка](#.D0.A0.D1.83.D1.87.D0.BD.D0.B0.D1.8F_.D0.BD.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0)
  - [2.2Графические веб-браузеры](#.D0.93.D1.80.D0.B0.D1.84.D0.B8.D1.87.D0.B5.D1.81.D0.BA.D0.B8.D0.B5_.D0.B2.D0.B5.D0.B1-.D0.B1.D1.80.D0.B0.D1.83.D0.B7.D0.B5.D1.80.D1.8B)
  - [2.3Веб-браузер в командной строке](#.D0.92.D0.B5.D0.B1-.D0.B1.D1.80.D0.B0.D1.83.D0.B7.D0.B5.D1.80_.D0.B2_.D0.BA.D0.BE.D0.BC.D0.B0.D0.BD.D0.B4.D0.BD.D0.BE.D0.B9_.D1.81.D1.82.D1.80.D0.BE.D0.BA.D0.B5)
  - [2.4Проверка и валидация](#.D0.9F.D1.80.D0.BE.D0.B2.D0.B5.D1.80.D0.BA.D0.B0_.D0.B8_.D0.B2.D0.B0.D0.BB.D0.B8.D0.B4.D0.B0.D1.86.D0.B8.D1.8F)
- [3Установка файла stage](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_.D1.84.D0.B0.D0.B9.D0.BB.D0.B0_stage)
- [4Настройка параметров компиляции](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D0.BF.D0.B0.D1.80.D0.B0.D0.BC.D0.B5.D1.82.D1.80.D0.BE.D0.B2_.D0.BA.D0.BE.D0.BC.D0.BF.D0.B8.D0.BB.D1.8F.D1.86.D0.B8.D0.B8)
  - [4.1Введение](#.D0.92.D0.B2.D0.B5.D0.B4.D0.B5.D0.BD.D0.B8.D0.B5)
  - [4.2CFLAGS и CXXFLAGS](#CFLAGS_.D0.B8_CXXFLAGS)
  - [4.3RUSTFLAGS](#RUSTFLAGS)
  - [4.4MAKEOPTS](#MAKEOPTS)
  - [4.5На старт, внимание, марш!](#.D0.9D.D0.B0_.D1.81.D1.82.D0.B0.D1.80.D1.82.2C_.D0.B2.D0.BD.D0.B8.D0.BC.D0.B0.D0.BD.D0.B8.D0.B5.2C_.D0.BC.D0.B0.D1.80.D1.88.21)
- [5Примечания](#.D0.9F.D1.80.D0.B8.D0.BC.D0.B5.D1.87.D0.B0.D0.BD.D0.B8.D1.8F)

## Выбор архива stage

**Совет**

On supported architectures, it is recommended for users targeting a desktop (graphical) operating system environment to use a stage file with the term `desktop` within the name. These files include packages such as [llvm-core/llvm](https://packages.gentoo.org/packages/llvm-core/llvm) and [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) and USE flag tuning which will greatly improve install time.

The [stage file](/wiki/Stage_file "Stage file") acts as the seed of a Gentoo install. Stage files are generated by the [Release Engineering Team](/wiki/Project:RelEng "Project:RelEng") with [Catalyst](/wiki/Catalyst "Catalyst"). Stage files are based on specific [profiles](/wiki/Profile_(Portage) "Profile (Portage)"), and contain an almost-complete system.

When choosing a stage file, it's important to pick one with profile targets corresponding to the desired system type.

**Заметка**

Хотя переключение между основными профилями после завершения установки возможно, это требует значительных усилий и осознанного решения, что выходит за рамки данного руководства по установке. Переключение между системами инициализации довольно сложно, а переключение с `no-multilib` на `multilib` требует больших познаний в устройстве Gentoo и низкоуровневых деталях работы системы сборки.

**Совет**

Большинству пользователей не следует использовать «продвинутые» варианты архивов; они предназначены для нетипичных или расширенных программных или аппаратных конфигураций.

### OpenRC

[OpenRC](/wiki/OpenRC/ru "OpenRC/ru") — это система инициализации на основе зависимостей (ответственна за загрузку системных сервисов после загрузки ядра), которая поддерживает совместимость с системной программой для загрузки (обычно расположенная в /sbin/init). Она является основной и оригинальной в Gentoo, но она также используется в некоторых других дистрибутивах Linux и системах BSD.

### systemd

systemd — это современная система инициализации для систем на базе Linux. Она используется как основная система инициализации в большинстве дистрибутивов Linux. systemd полностью поддерживается в Gentoo и работает по своему прямому назначению. Если вам кажется, что чего–то не хватает в Руководстве для поддержки установки с systemd, просмотрите [статью systemd](/wiki/Systemd/ru "Systemd/ru") _перед тем_, как попросить поддержки.

### Multilib (32 и 64 бит)

**Заметка**

Не все архитектуры имеют поддержку multilib. Много архитектур выполняют код в «родном» режиме. Multilib по большей части относится к **amd64**.

В профиле multilib предпочтение отдаётся 64-битным библиотекам, но, если необходимо обеспечение совместимости, есть возможность использовать 32-битные версии. Это великолепный выбор для большинства установок, так как он обеспечивает большую гибкость конфигурации в будущем.

**Совет**

По сравнению с `no-multilib`, использование `multilib` значительно облегчает последующее переключение на другие профили.

### No-multilib (чистый 64-bit)

**Предупреждение**

Имейте в виду, пользователям, которые только начинают знакомиться с Gentoo, следует _избегать_ использование архива no-multilib (если только это не продиктовано другими соображениями). Миграция с `no-multilib` на `multilib` потребует чрезвычайно хорошего знания Gentoo и наличия набора инструментов разработки более низкого уровня (от которых наших [разработчиков Toolchain](/wiki/Project:Toolchain "Project:Toolchain") может даже бросить дрожь). Данный процесс не для слабонервных и выходит за рамки данного руководства.

Выбор no‐multilib архива в качестве основы для системы позволяет получить полностью 64‐битное окружение операционной системы — без 32-битных программ. Это существенно усложняет возможность перехода на профили `multilib`, но технически это по‐прежнему осуществимо.

## Скачивание архива stage

Before downloading the _stage file_, the current directory should be set to the location of the mount used for the install:

`root #` `cd /mnt/gentoo`

### Установка даты и времени

Архивы стадий обычно скачиваются по протоколу HTTPS, для которого важно относительно точное системное время. Неточные часы могут стать причиной невозможности скачиваний и привести к непредсказуемым ошибкам, если системное время было значительно изменено после установки.

Текущую дату и время можно проверить с помощью команды date:

`root #` `date`

```
Mon Oct  3 13:16:22 PDT 2021

```

Если время и дата отличаются более чем на несколько минут, их следует обновить, используя один из следующих способов.

#### Автоматическая настройка

Использование [NTP](/wiki/NTP "NTP") для коррекции часов обычно гораздо проще и надёжнее, чем ручная установка системных часов.

Для установки системного времени в UTC можно воспользоваться chronyd из пакета [net-misc/chrony](https://packages.gentoo.org/packages/net-misc/chrony):

`root #` `chronyd -q`

**Важно**

Системы без функционирующих часов реального времени (Real-Time Clock, RTC) должны синхронизировать системное время при каждом запуске системы, а затем — регулярно в течение некоторых интервалов. Это также полезно и для систем с RTC, так как встроенная батарейка может сесть, а отклонение часов может накапливаться.

**Предупреждение**

Стандартный трафик NTP не аутентифицируется, важно проверять данные о времени, полученные через сеть.

#### Ручная настройка

Если NTP недоступен, для ручной настройки системных часов можно использовать date.

**Заметка**

Для всех систем Linux рекомендуется использовать время UTC. Позже будет установлен часовой пояс, вносящий соответствующее смещение при отображении даты.

Для установки времени используется следующий формат аргументов: code>MMDDhhmmYYYY ( **M** — месяц, **D** — день, **h** — час, **m** — минута и **Y** — год).

Например, чтобы установить дату на 13:16 3 октября 2021 года, выполните:

`root #` `date 100313162021`

### Графические веб-браузеры

У пользователей, использующих среду с полноценными веб-браузерами, не будет никаких проблем с копированием URL файла stage из [раздела загрузки](https://www.gentoo.org/downloads/#other-arches) главного веб-сайта. Просто выберите подходящую вкладку, щёлкните правой кнопкой по ссылке файла stage, выберите Копировать ссылку, скопировав её в буфер обмена. Затем вставьте ссылку в командной строке после команды wget для скачивания файла:

`root #` `wget <PASTED_STAGE_FILE_URL>`

### Веб-браузер в командной строке

Более опытные пользователи или «старики» Gentoo, которые работают исключительно из командной строки, могут воспользоваться links ([www-client/links](https://packages.gentoo.org/packages/www-client/links)) — консольным веб-браузером на основе меню. Чтобы загрузить файл архива stage, просмотрите список зеркал Gentoo:

`root #` `links https://www.gentoo.org/downloads/mirrors/`

Чтобы использовать HTTP-прокси в links, введите URL с параметром `-http-proxy`:

`root #` `links -http-proxy proxy.server.com:8080 https://www.gentoo.org/downloads/mirrors/`

Наряду с links так же есть браузер lynx ([www-client/lynx](https://packages.gentoo.org/packages/www-client/lynx)). Как и links, он не имеет графического интерфейса, но у него нет меню.

`root #` `lynx https://www.gentoo.org/downloads/mirrors/`

Если прокси нужно сохранить, экспортируйте переменные http\_proxy и/или ftp\_proxy:

`root #` `export http_proxy="http://proxy.server.com:port"
`

`root #` `export ftp_proxy="http://proxy.server.com:port"`

В списке зеркал выберите зеркало, которое находится рядом. Обычно достаточно HTTP-зеркала, но другие протоколы также доступны. Перейдите в каталог releases/x86/autobuilds/. Там отображаются все доступные stage-файлы (они могут находиться в подкаталогах с названиями отдельных субархитектур). Выберите нужный и нажмите `d` для скачивания.

После завершения скачивания можно проверить целостность и достоверность содержимого файла stage. Если вам это интересно, перейдите к [следующему разделу](/wiki/Handbook:X86/Installation/Stage/ru#Verifying_and_validating "Handbook:X86/Installation/Stage/ru").

Тем, кому не интересно проверять архив stage, могут закрыть браузер в командной строке с помощью клавиши `q` и сразу перейти к разделу [Распаковка файла stage](/wiki/Handbook:X86/Installation/Stage/ru#Unpacking_the_stage_file "Handbook:X86/Installation/Stage/ru").

### Проверка и валидация

Как и в случае с минимальными установочными компакт-дисками, доступно несколько файлов, необходимых для проверки и валидации файла stage. Хотя этот шаг может быть пропущен, эти файлы могут пригодиться пользователям, которым важна целостность только что скачанных файлов. Упомянутые дополнительные файлы находятся в том же каталоге, что и связанные stage-файлы. Перейдите в каталог, соответствующий аппаратной архитектуре и профилю системы, и загрузите соответствующие файлы с суффиксами .CONTENTS.gz, .DIGESTS и .sha256.

`root #` `wget https://distfiles.gentoo.org/releases/`

- .CONTENTS.gz — сжатый файл, содержащий список всех файлов в файле stage.
- .DIGESTS — содержит контрольные суммы файла stage, вычисленные с использованием нескольких криптографических хеш-функций.
- .sha256 — содержит контрольную сумму файла stage, подписанную с помощью PGP. Этот файл может быть недоступен для некоторых файлов stage.

Как и в случае с ISO-файлом, можно проверить криптографическую подпись файла tar.xz с помощью команды gpg, чтобы убедиться, что архив не был подделан.

Для официальных образов Gentoo предоставляется пакет [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release), в котором содержатся PGP-ключи, которыми подписываются автоматические выпуски. Чтобы сверить ключи, их необходимо сначала импортировать в сессию пользователя:

`root #` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

Для неофициальных образов, которые содержат утилиты gpg и wget, см. предыдущий раздел [Проверка скачанных файлов](/wiki/Handbook:X86/Installation/Media/ru#Verifying_the_downloaded_files "Handbook:X86/Installation/Media/ru").

Чтобы проверить подпись архива и (опционально) связанный файл с контрольными суммами:

`root #` `gpg --verify stage3-x86-<release>-<init>.tar.xz.asc stage3-x86-<release>-<init>.tar.xz
`

`root #` `gpg --output stage3-x86-<release>-<init>.tar.xz.DIGESTS.verified --verify stage3-x86-<release>-<init>.tar.xz.DIGESTS
`

`root #` `gpg --output stage3-x86-<release>-<init>.tar.xz.sha256.verified --verify stage3-x86-<release>-<init>.tar.xz.sha256
`

Если верификация пройдёт успешно, в выводе предыдущих команд появится сообщение «Good signature from».

Отпечатки ключей OpenPGP, используемые для подписи релизов, можно найти на [странице сигнатур](https://www.gentoo.org/downloads/signatures/).

**Заметка**

Многие stage архивы теперь явно [содержат суффикс](https://www.gentoo.org/news/2021/07/20/more-downloads.html) с типом системы инициализации (openrc или systemd), хотя у некоторых архитектур они могут пока отсутствовать.

Для сравнения контрольных сумм реальных данных с контрольными суммами из файла .DIGESTS можно использовать криптографические инструменты и утилиты, такие как openssl, sha256sum или sha512sum.

Чтобы проверить контрольную сумму SHA512 с помощью openssl:

`root #` `openssl dgst -r -sha512 stage3-x86-<release>-<init>.tar.xz`

Флаг `dgst` указывает команде openssl, что используется подкоманда Message Digest (Хеш-сумма Сообщения), `-r` выводит хеш-суммы в формате coreutils, а `-sha512` выбирает алгоритм SHA512.

Чтобы проверить контрольную сумму BLAKE2B512 с помощью openssl:

`root #` `openssl dgst -r -blake2b512 stage3-x86-<release>-<init>.tar.xz`

Сравните полученные результаты с парами из хэш-суммы и имён файлов из .DIGESTS.verified. Если эти значения не совпадают с полученными, это означает, что загруженный файл поврежден и его следует отбраковать и загрузить заново.

Чтобы проверить хеш SHA256 с помощью утилиты sha256sum, используя связанный файл .sha256, запустите эту команду:

`root #` `sha256sum --check stage3-x86-<release>-<init>.tar.xz.sha256.verified`

Опция `--check` указывает sha256sum, что из полученного файла необходимо считать список пар из ожидаемых файлов и их хешей, а затем вычислять реальные значения и выводить для каждого файла либо «OK», если хеш-суммы совпали для этого файла, либо «FAILED», если нет.

## Установка файла stage

После скачивания и проверки _файл stage_ необходимо распаковать с помощью tar:

`root #` `tar xpvf stage3-*.tar.xz --xattrs-include='*.*' --numeric-owner -C /mnt/gentoo`

Перед распаковкой ознакомьтесь с ее параметрами:

- `x` e **x** tract, instructs tar to extract the contents of the archive.
- `p` **p** reserve permissions.
- `v` **v** erbose output.
- `f` **f** ile, provides tar with the name of the input archive.
- `--xattrs-include='*.*'` Preserves extended attributes in all namespaces stored in the archive.
- `--numeric-owner` Ensure that the user and group IDs of files being extracted from the tarball remain the same as Gentoo's release engineering team intended (even if adventurous users are not using official Gentoo live environments for the installation process).
- `-C /mnt/gentoo` Extract files to the root partition regardless of the current directory.

Теперь, когда stage распакован, перейдём к [настройке параметров компиляции](/wiki/Handbook:X86/Installation/Stage/ru#Configuring_compile_options "Handbook:X86/Installation/Stage/ru").

## Настройка параметров компиляции

### Введение

Для оптимизации системы можно установить переменные, которые влияют на поведение Portage, официально поддерживаемого менеджера пакетов Gentoo. Все эти переменные могут быть установлены как переменные окружения (с помощью export), но установка через export не будет постоянной.

**Заметка**

Технически, переменные можно экспортировать через профиль [shell](/wiki/Shell "Shell") или файлы rc, однако это не лучшая практика для базового администрирования системы.

При запуске Portage читает файл [make.conf](/wiki/Make.conf "Make.conf"), который изменяет поведение во время выполнения в зависимости от значений, записанных в этом файле. make.conf можно считать основным конфигурационным файлом для Portage, поэтому будьте внимательнее с его содержимым.

**Совет**

Cписок всех возможных переменных с комментариями можно найти в /mnt/gentoo/usr/share/portage/config/make.conf.example. Дополнительная документация по файлу make.conf доступна через справочную страницу man 5 make.conf.

Для успешной установки Gentoo необходимо установить только те переменные, которые указаны ниже.

Запустите редактор (в этом руководстве мы используем nano) для изменения параметров оптимизации, о которых написано далее.

`root #` `nano /mnt/gentoo/etc/portage/make.conf`

В файле make.conf.example показано, как файл должен быть структурирован: строки комментариев начинаются с `#`, другие строки описывают переменные вида `ПЕРЕМЕННАЯ="содержание"`. Некоторые из этих переменных мы обсудим в следующем разделе.

### CFLAGS и CXXFLAGS

Переменные CFLAGS и CXXFLAGS определяют параметры оптимизации для компиляторов GCC C и C++ соответственно. Хотя они и указаны здесь, для достижения максимальной производительности можно было бы указать флаги оптимизации для каждой программы отдельно. Причина этого в том, что все программы различны. Но этим тяжело управлять, следовательно, запишем эти переменные в make.conf файл.

В make.conf следует указывать параметры оптимизации, которые сделают систему наиболее отзывчивой в целом. Не нужно использовать экспериментальные настройки; излишняя оптимизация может привести к непредсказуемому поведению программ (аварийному завершению, или ещё хуже, к неправильной работе).

Руководство пользователя не описывает все возможные параметры оптимизации. За более подробной информацией обратитесь к [Документации GNU](https://gcc.gnu.org/onlinedocs/) или к инфо-странице gcc (info gcc). Сам файл make.conf.example тоже содержит множество примеров и информации, так что не забудьте прочитать его.

Первым параметром обычно является флаг `-march=` или `-mtune=`, который указывает имя целевой архитектуры. Возможные варианты описаны в файле make.conf.example (в комментариях). Обычно используется значение _native_, который сообщает компилятору, чтобы он использовал целевую архитектуру существующей системы (той, на которую будет установлена Gentoo).

Второй параметр оптимизации — это флаг `-О` (это заглавная буква О, а не ноль), который определяет класс оптимизации для gcc. Возможные классы: s (оптимизация по размеру), 0 (ноль — без оптимизации), 1, 2 или даже 3 для более лучшей оптимизация по скорости (в каждый класс входят все флаги предыдущего, и некоторые дополнительные). `-O2` является рекомендованным значением по умолчанию. `-O3` может вызывать проблемы при глобальном использовании на уровне системы, так что мы рекомендуем придерживаться `-O2`.

Ещё одним популярным флагом оптимизации является `-pipe` (использование конвейера вместо временных файлов для взаимодействия между различными стадиями компиляции). Это не имеет никакого влияния на сгенерированный код, при этом использует больше памяти. В системах с небольшим объемом памяти gcc может аварийно завершиться из-за нехватки памяти. В этом случае не используйте этот флаг.

Использование `-fomit-frame-pointer` (не хранить указатель фрейма в регистре для функций, которым он не нужен) может привести к серьезным последствиям во время отладки приложений.

Определение переменных CFLAGS и CXXFLAGS позволяет комбинировать несколько флагов оптимизации в одной строке. Значений по умолчанию, содержащихся в файле stage3, обычно более чем достаточно. Ниже приведён пример конфигурации:

КОД **Пример для переменных CFLAGS и CXXFLAGS**

```
# Флаги компилятора, используемые для всех языков
COMMON_FLAGS="-O2 -march=i686 -pipe"
# Используйте те же настройки для обеих переменных
CFLAGS="${COMMON_FLAGS}"
CXXFLAGS="${COMMON_FLAGS}"

```

**Совет**

Хотя [руководство по оптимизации GCC](/wiki/GCC_optimization/ru "GCC optimization/ru") имеет больше информации о том, как различные параметры компиляции могут повлиять на систему, статья [Safe CFLAGS](/wiki/Safe_CFLAGS "Safe CFLAGS") будет более полезной для начинающих пользователей, желающих оптимизировать свою систему.

### RUSTFLAGS

Many programs are now written in Rust which has its own way of optimising. By default Rust optimizes to level 3 on all release builds unless a project changes it so this should be left as is. A full optimization list (known as codegen) that can be passed to the rust compiler is available at [https://doc.rust-lang.org/rustc/codegen-options/index.html](https://doc.rust-lang.org/rustc/codegen-options/index.html).

The most useful optimization would be to tell Rust to compile for your CPU using the following example:

ФАЙЛ **`/etc/portage/make.conf`** **Пример RUSTFLAGS**

```
RUSTFLAGS="${RUSTFLAGS} -C target-cpu=native"

```

To get a list of supported CPUs in Rust run:

`root #` `rustc -C target-cpu=help`

This will show what the default and also which CPU will be selected with native.

**Заметка**

The above command only works on desktop stage3 tarballs or after emerging [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) or [dev-lang/rust](https://packages.gentoo.org/packages/dev-lang/rust).

### MAKEOPTS

Переменная MAKEOPTS определяет, сколько параллельных процессов компиляции должно запускаться при установке пакета. На момент Portage версии 3.0.31[\[1\]](#cite_note-1), если переменная не определена, Portage выставит параметр jobs в MAKEOPTS на то же количество потоков, которое возвращает nproc.

Кроме того, начиная с Portage версии 3.0.53[\[2\]](#cite_note-2), если переменная не определена, Portage по такой же логике (используя вывод nproc) выставит параметр load-average в MAKEOPTS.

Лучше всего выбрать наименьшее из следующих значений: количество потоков у процессора или общий объем ОЗУ системы, разделённый на 2 ГиБ.

**Предупреждение**

Использование большого количества процессов может значительно повлиять на потребление памяти. Хорошая рекомендация — имейте не менее 2 Гб свободной оперативной памяти на каждый поток (так, например, для `-j6` потребуется _не менее_ 12 ГиБ). Чтобы избежать нехватки памяти, уменьшите количество процессов.

**Совет**

Когда вы используете параллельный emerge ( `--jobs`), количество потоков может вырасти в разы (до количества make потоков, помноженное на количество emerge потоков). Это можно обойти, запустив distcc только для localhost, с конфигурацией, которая ограничит количество экземпляров компилятора на хосте.

ФАЙЛ **`/etc/portage/make.conf`** **Пример объявления переменной MAKEOPTS**

```
# Если переменная не определена, то по умолчанию Portage:
# - устанавливает значение параметра jobs в переменной MAKEOPTS равным количеству потоков, получая информацию от `nproc`
# - устанавливает значение параметра load-average в переменной MAKEOPTS чуть больше количества потоков, получая информацию от `nproc`, из-за того, что это затухающее значение
# Пожалуйста, замените число '4' на соответствующее характеристикам системы (руководствуясь формулой min(ОЗУ/2ГиБ, кол-во потоков))
# или оставьте не-заданным.
MAKEOPTS="-j4 -l5"

```

Для получения более подробной информации прочтите о переменной MAKEOPTS в man 5 make.conf.

### На старт, внимание, марш!

Обновите файл /mnt/gentoo/etc/portage/make.conf в соответствии с личными предпочтениями и сохраните изменения (в nano нужно нажать `Ctrl` + `o`, чтобы записать изменения, и затем `Ctrl` + `x` для выхода).

## Примечания

1. [↑](#cite_ref-1)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2](https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2)
2. [↑](#cite_ref-2)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e](https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e)

[← Подготовка дисков](/wiki/Handbook:X86/Installation/Disks/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:X86/ru "Handbook:X86/ru") [Установка базовой системы Gentoo →](/wiki/Handbook:X86/Installation/Base/ru "Следующая часть")