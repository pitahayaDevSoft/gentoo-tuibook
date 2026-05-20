# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbuch:PPC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Bootloader "Handbook:PPC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Bootloader/es "Manual de Gentoo: PPC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Bootloader/fr "Handbook:PPC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Bootloader/it "Handbook:PPC/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Bootloader/hu "Handbook:PPC/Installation/Bootloader/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:PPC/Installation/Bootloader/pt-br "Handbook:PPC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Bootloader/ru "Handbook:PPC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Bootloader/ta "கையேடு:PPC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Bootloader/zh-cn "手册：PPC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Bootloader/ja "ハンドブック:PPC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Bootloader/ko "Handbook:PPC/Installation/Bootloader/ko (100% translated)")

[PPC Handbook](/wiki/Handbook:PPC/pl "Handbook:PPC/pl")[Instalacja](/wiki/Handbook:PPC/Full/Installation/pl "Handbook:PPC/Full/Installation/pl")[O instalacji](/wiki/Handbook:PPC/Installation/About/pl "Handbook:PPC/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:PPC/Installation/Media/pl "Handbook:PPC/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:PPC/Installation/Networking/pl "Handbook:PPC/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:PPC/Installation/Disks/pl "Handbook:PPC/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:PPC/Installation/Stage/pl "Handbook:PPC/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:PPC/Installation/Base/pl "Handbook:PPC/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:PPC/Installation/Kernel/pl "Handbook:PPC/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:PPC/Installation/System/pl "Handbook:PPC/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:PPC/Installation/Tools/pl "Handbook:PPC/Installation/Tools/pl")Instalacja systemu rozruchowego[Finalizacja](/wiki/Handbook:PPC/Installation/Finalizing/pl "Handbook:PPC/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:PPC/Full/Working/pl "Handbook:PPC/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:PPC/Working/Portage/pl "Handbook:PPC/Working/Portage/pl")[Flagi USE](/wiki/Handbook:PPC/Working/USE/pl "Handbook:PPC/Working/USE/pl")[Funkcje portage](/wiki/Handbook:PPC/Working/Features/pl "Handbook:PPC/Working/Features/pl")[System initscript](/wiki/Handbook:PPC/Working/Initscripts/pl "Handbook:PPC/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:PPC/Working/EnvVar/pl "Handbook:PPC/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:PPC/Full/Portage/pl "Handbook:PPC/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:PPC/Portage/Files/pl "Handbook:PPC/Portage/Files/pl")[Zmienne](/wiki/Handbook:PPC/Portage/Variables/pl "Handbook:PPC/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:PPC/Portage/Branches/pl "Handbook:PPC/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:PPC/Portage/Tools/pl "Handbook:PPC/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:PPC/Portage/CustomTree/pl "Handbook:PPC/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:PPC/Portage/Advanced/pl "Handbook:PPC/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:PPC/Full/Networking/pl "Handbook:PPC/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:PPC/Networking/Introduction/pl "Handbook:PPC/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:PPC/Networking/Advanced/pl "Handbook:PPC/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:PPC/Networking/Modular/pl "Handbook:PPC/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:PPC/Networking/Wireless/pl "Handbook:PPC/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:PPC/Networking/Extending/pl "Handbook:PPC/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:PPC/Networking/Dynamic/pl "Handbook:PPC/Networking/Dynamic/pl")

[Handbook:PPC/Blocks/Bootloader/pl](/index.php?title=Handbook:PPC/Blocks/Bootloader/pl&action=edit&redlink=1 "Handbook:PPC/Blocks/Bootloader/pl (page does not exist)")

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

Po uruchomieniu świeżo zainstalowanego środowiska Gentoo, kontynuuj [Finalizowanie instalacji Gentoo](/wiki/Handbook:PPC/Installation/Finalizing "Handbook:PPC/Installation/Finalizing").

[← Installing tools](/wiki/Handbook:PPC/Installation/Tools/pl "Previous part") [Home](/wiki/Handbook:PPC "Handbook:PPC") [Finalizing →](/wiki/Handbook:PPC/Installation/Finalizing/pl "Next part")