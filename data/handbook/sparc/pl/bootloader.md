# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbuch:SPARC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Bootloader "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Bootloader/es "Manual de Gentoo: SPARC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Bootloader/fr "Handbook:SPARC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Bootloader/pt-br "Handbook:SPARC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Bootloader/ru "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Bootloader/ta "கையேடு:SPARC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Bootloader/zh-cn "手册：SPARC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Bootloader/ja "ハンドブック:SPARC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Bootloader/ko "Handbook:SPARC/Installation/Bootloader/ko (100% translated)")

[SPARC Handbook](/wiki/Handbook:SPARC/pl "Handbook:SPARC/pl")[Instalacja](/wiki/Handbook:SPARC/Full/Installation/pl "Handbook:SPARC/Full/Installation/pl")[O instalacji](/wiki/Handbook:SPARC/Installation/About/pl "Handbook:SPARC/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:SPARC/Installation/Media/pl "Handbook:SPARC/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:SPARC/Installation/Networking/pl "Handbook:SPARC/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:SPARC/Installation/Disks/pl "Handbook:SPARC/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:SPARC/Installation/Stage/pl "Handbook:SPARC/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:SPARC/Installation/Base/pl "Handbook:SPARC/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:SPARC/Installation/Kernel/pl "Handbook:SPARC/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:SPARC/Installation/System/pl "Handbook:SPARC/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:SPARC/Installation/Tools/pl "Handbook:SPARC/Installation/Tools/pl")Instalacja systemu rozruchowego[Finalizacja](/wiki/Handbook:SPARC/Installation/Finalizing/pl "Handbook:SPARC/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:SPARC/Full/Working/pl "Handbook:SPARC/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:SPARC/Working/Portage/pl "Handbook:SPARC/Working/Portage/pl")[Flagi USE](/wiki/Handbook:SPARC/Working/USE/pl "Handbook:SPARC/Working/USE/pl")[Funkcje portage](/wiki/Handbook:SPARC/Working/Features/pl "Handbook:SPARC/Working/Features/pl")[System initscript](/wiki/Handbook:SPARC/Working/Initscripts/pl "Handbook:SPARC/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:SPARC/Working/EnvVar/pl "Handbook:SPARC/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:SPARC/Full/Portage/pl "Handbook:SPARC/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:SPARC/Portage/Files/pl "Handbook:SPARC/Portage/Files/pl")[Zmienne](/wiki/Handbook:SPARC/Portage/Variables/pl "Handbook:SPARC/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:SPARC/Portage/Branches/pl "Handbook:SPARC/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:SPARC/Portage/Tools/pl "Handbook:SPARC/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:SPARC/Portage/CustomTree/pl "Handbook:SPARC/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:SPARC/Portage/Advanced/pl "Handbook:SPARC/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:SPARC/Full/Networking/pl "Handbook:SPARC/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:SPARC/Networking/Introduction/pl "Handbook:SPARC/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:SPARC/Networking/Advanced/pl "Handbook:SPARC/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:SPARC/Networking/Modular/pl "Handbook:SPARC/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:SPARC/Networking/Wireless/pl "Handbook:SPARC/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:SPARC/Networking/Extending/pl "Handbook:SPARC/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:SPARC/Networking/Dynamic/pl "Handbook:SPARC/Networking/Dynamic/pl")

[Handbook:SPARC/Blocks/Bootloader/pl](/index.php?title=Handbook:SPARC/Blocks/Bootloader/pl&action=edit&redlink=1 "Handbook:SPARC/Blocks/Bootloader/pl (page does not exist)")

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

Po uruchomieniu świeżo zainstalowanego środowiska Gentoo, kontynuuj [Finalizowanie instalacji Gentoo](/wiki/Handbook:SPARC/Installation/Finalizing "Handbook:SPARC/Installation/Finalizing").

[← Installing tools](/wiki/Handbook:SPARC/Installation/Tools/pl "Previous part") [Home](/wiki/Handbook:SPARC "Handbook:SPARC") [Finalizing →](/wiki/Handbook:SPARC/Installation/Finalizing/pl "Next part")