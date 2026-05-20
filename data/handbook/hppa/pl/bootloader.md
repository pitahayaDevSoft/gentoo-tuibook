# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbuch:HPPA/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Bootloader/es "Manual de Gentoo: HPPA/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Bootloader/it "Handbook:HPPA/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Bootloader/pt-br "Manual:HPPA/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Bootloader/ru "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Bootloader/ta "கையேடு:HPPA/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Bootloader/zh-cn "手册：HPPA/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Bootloader/ja "ハンドブック:HPPA/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko (100% translated)")

[HPPA Handbook](/wiki/Handbook:HPPA/pl "Handbook:HPPA/pl")[Instalacja](/wiki/Handbook:HPPA/Full/Installation/pl "Handbook:HPPA/Full/Installation/pl")[O instalacji](/wiki/Handbook:HPPA/Installation/About/pl "Handbook:HPPA/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:HPPA/Installation/Media/pl "Handbook:HPPA/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:HPPA/Installation/Networking/pl "Handbook:HPPA/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:HPPA/Installation/Disks/pl "Handbook:HPPA/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:HPPA/Installation/Stage/pl "Handbook:HPPA/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:HPPA/Installation/Base/pl "Handbook:HPPA/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:HPPA/Installation/Kernel/pl "Handbook:HPPA/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:HPPA/Installation/System/pl "Handbook:HPPA/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:HPPA/Installation/Tools/pl "Handbook:HPPA/Installation/Tools/pl")Instalacja systemu rozruchowego[Finalizacja](/wiki/Handbook:HPPA/Installation/Finalizing/pl "Handbook:HPPA/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:HPPA/Full/Working/pl "Handbook:HPPA/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:HPPA/Working/Portage/pl "Handbook:HPPA/Working/Portage/pl")[Flagi USE](/wiki/Handbook:HPPA/Working/USE/pl "Handbook:HPPA/Working/USE/pl")[Funkcje portage](/wiki/Handbook:HPPA/Working/Features/pl "Handbook:HPPA/Working/Features/pl")[System initscript](/wiki/Handbook:HPPA/Working/Initscripts/pl "Handbook:HPPA/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:HPPA/Working/EnvVar/pl "Handbook:HPPA/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:HPPA/Full/Portage/pl "Handbook:HPPA/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:HPPA/Portage/Files/pl "Handbook:HPPA/Portage/Files/pl")[Zmienne](/wiki/Handbook:HPPA/Portage/Variables/pl "Handbook:HPPA/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:HPPA/Portage/Branches/pl "Handbook:HPPA/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:HPPA/Portage/Tools/pl "Handbook:HPPA/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:HPPA/Portage/CustomTree/pl "Handbook:HPPA/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:HPPA/Portage/Advanced/pl "Handbook:HPPA/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:HPPA/Full/Networking/pl "Handbook:HPPA/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:HPPA/Networking/Introduction/pl "Handbook:HPPA/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:HPPA/Networking/Advanced/pl "Handbook:HPPA/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:HPPA/Networking/Modular/pl "Handbook:HPPA/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:HPPA/Networking/Wireless/pl "Handbook:HPPA/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:HPPA/Networking/Extending/pl "Handbook:HPPA/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:HPPA/Networking/Dynamic/pl "Handbook:HPPA/Networking/Dynamic/pl")

[Handbook:HPPA/Blocks/Bootloader/pl](/index.php?title=Handbook:HPPA/Blocks/Bootloader/pl&action=edit&redlink=1 "Handbook:HPPA/Blocks/Bootloader/pl (page does not exist)")

## Ponowne uruchomienie systemu

Wyjdź z środowiska chroot i odmontuj wszystkie zamontowane partycje. Następnie wpisz magiczne polecenie, które inicjuje ostateczny, prawdziwy test: reboot.

`root #` `exit`

`cdimage ~#` `cd
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`cdimage ~#` `umount -R /mnt/gentoo
`

`cdimage ~#` `reboot`

Nie zapomnij usunąć medium instalacyjnego, w przeciwnym razie zamiast nowego systemu Gentoo może zostać ponownie uruchomione medium instalacyjne.

Po uruchomieniu świeżo zainstalowanego środowiska Gentoo, kontynuuj [Finalizowanie instalacji Gentoo](/wiki/Handbook:HPPA/Installation/Finalizing "Handbook:HPPA/Installation/Finalizing").

[← Installing tools](/wiki/Handbook:HPPA/Installation/Tools/pl "Previous part") [Home](/wiki/Handbook:HPPA "Handbook:HPPA") [Finalizing →](/wiki/Handbook:HPPA/Installation/Finalizing/pl "Next part")