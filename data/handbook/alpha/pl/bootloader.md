# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Bootloader/de "Handbuch:Alpha/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Bootloader "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Bootloader/es "Manual de Gentoo: Alpha/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Bootloader/fr "Manuel:Alpha/Installation/Système d'amorçage (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Bootloader/it "Handbook:Alpha/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Bootloader/hu "Handbook:Alpha/Installation/Bootloader/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Bootloader/pt-br "Manual:Alpha/Instalação/Gerenciador de boot (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Bootloader/cs "Handbook:Alpha/Installation/Bootloader/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Bootloader/ru "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Bootloader/ta "கையேடு:Alpha/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Bootloader/zh-cn "手册：Alpha/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Bootloader/ja "ハンドブック:Alpha/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Bootloader/ko "Handbook:Alpha/Installation/Bootloader/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha/pl "Handbook:Alpha/pl")[Instalacja](/wiki/Handbook:Alpha/Full/Installation/pl "Handbook:Alpha/Full/Installation/pl")[O instalacji](/wiki/Handbook:Alpha/Installation/About/pl "Handbook:Alpha/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:Alpha/Installation/Media/pl "Handbook:Alpha/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:Alpha/Installation/Networking/pl "Handbook:Alpha/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:Alpha/Installation/Stage/pl "Handbook:Alpha/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:Alpha/Installation/Base/pl "Handbook:Alpha/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:Alpha/Installation/Kernel/pl "Handbook:Alpha/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:Alpha/Installation/System/pl "Handbook:Alpha/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:Alpha/Installation/Tools/pl "Handbook:Alpha/Installation/Tools/pl")Instalacja systemu rozruchowego[Finalizacja](/wiki/Handbook:Alpha/Installation/Finalizing/pl "Handbook:Alpha/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:Alpha/Full/Working/pl "Handbook:Alpha/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:Alpha/Working/Portage/pl "Handbook:Alpha/Working/Portage/pl")[Flagi USE](/wiki/Handbook:Alpha/Working/USE/pl "Handbook:Alpha/Working/USE/pl")[Funkcje portage](/wiki/Handbook:Alpha/Working/Features/pl "Handbook:Alpha/Working/Features/pl")[System initscript](/wiki/Handbook:Alpha/Working/Initscripts/pl "Handbook:Alpha/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:Alpha/Working/EnvVar/pl "Handbook:Alpha/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:Alpha/Full/Portage/pl "Handbook:Alpha/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:Alpha/Portage/Files/pl "Handbook:Alpha/Portage/Files/pl")[Zmienne](/wiki/Handbook:Alpha/Portage/Variables/pl "Handbook:Alpha/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:Alpha/Portage/Branches/pl "Handbook:Alpha/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:Alpha/Portage/Tools/pl "Handbook:Alpha/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:Alpha/Portage/CustomTree/pl "Handbook:Alpha/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:Alpha/Portage/Advanced/pl "Handbook:Alpha/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:Alpha/Full/Networking/pl "Handbook:Alpha/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:Alpha/Networking/Introduction/pl "Handbook:Alpha/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:Alpha/Networking/Advanced/pl "Handbook:Alpha/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:Alpha/Networking/Modular/pl "Handbook:Alpha/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:Alpha/Networking/Wireless/pl "Handbook:Alpha/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:Alpha/Networking/Extending/pl "Handbook:Alpha/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:Alpha/Networking/Dynamic/pl "Handbook:Alpha/Networking/Dynamic/pl")

[Handbook:Alpha/Blocks/Bootloader/pl](/index.php?title=Handbook:Alpha/Blocks/Bootloader/pl&action=edit&redlink=1 "Handbook:Alpha/Blocks/Bootloader/pl (page does not exist)")

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

Po uruchomieniu świeżo zainstalowanego środowiska Gentoo, kontynuuj [Finalizowanie instalacji Gentoo](/wiki/Handbook:Alpha/Installation/Finalizing "Handbook:Alpha/Installation/Finalizing").

[← Installing tools](/wiki/Handbook:Alpha/Installation/Tools/pl "Previous part") [Home](/wiki/Handbook:Alpha "Handbook:Alpha") [Finalizing →](/wiki/Handbook:Alpha/Installation/Finalizing/pl "Next part")