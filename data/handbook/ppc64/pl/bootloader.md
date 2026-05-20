# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Bootloader/de "Handbuch:PPC64/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Bootloader "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Bootloader/es "Manual de Gentoo: PPC64/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Bootloader/it "Handbook:PPC64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Handbook:PPC64/Installation/Bootloader/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Bootloader/pt-br "Handbook:PPC64/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Bootloader/ru "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Bootloader/ta "கையேடு:PPC64/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Bootloader/zh-cn "手册：PPC64/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Bootloader/ja "ハンドブック:PPC64/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko (100% translated)")

[PPC64 Handbook](/wiki/Handbook:PPC64/pl "Handbook:PPC64/pl")[Instalacja](/wiki/Handbook:PPC64/Full/Installation/pl "Handbook:PPC64/Full/Installation/pl")[O instalacji](/wiki/Handbook:PPC64/Installation/About/pl "Handbook:PPC64/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:PPC64/Installation/Media/pl "Handbook:PPC64/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:PPC64/Installation/Networking/pl "Handbook:PPC64/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:PPC64/Installation/Disks/pl "Handbook:PPC64/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:PPC64/Installation/Stage/pl "Handbook:PPC64/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:PPC64/Installation/Base/pl "Handbook:PPC64/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:PPC64/Installation/Kernel/pl "Handbook:PPC64/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:PPC64/Installation/System/pl "Handbook:PPC64/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:PPC64/Installation/Tools/pl "Handbook:PPC64/Installation/Tools/pl")Instalacja systemu rozruchowego[Finalizacja](/wiki/Handbook:PPC64/Installation/Finalizing/pl "Handbook:PPC64/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:PPC64/Full/Working/pl "Handbook:PPC64/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:PPC64/Working/Portage/pl "Handbook:PPC64/Working/Portage/pl")[Flagi USE](/wiki/Handbook:PPC64/Working/USE/pl "Handbook:PPC64/Working/USE/pl")[Funkcje portage](/wiki/Handbook:PPC64/Working/Features/pl "Handbook:PPC64/Working/Features/pl")[System initscript](/wiki/Handbook:PPC64/Working/Initscripts/pl "Handbook:PPC64/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:PPC64/Working/EnvVar/pl "Handbook:PPC64/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:PPC64/Full/Portage/pl "Handbook:PPC64/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:PPC64/Portage/Files/pl "Handbook:PPC64/Portage/Files/pl")[Zmienne](/wiki/Handbook:PPC64/Portage/Variables/pl "Handbook:PPC64/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:PPC64/Portage/Branches/pl "Handbook:PPC64/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:PPC64/Portage/Tools/pl "Handbook:PPC64/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:PPC64/Portage/CustomTree/pl "Handbook:PPC64/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:PPC64/Portage/Advanced/pl "Handbook:PPC64/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:PPC64/Full/Networking/pl "Handbook:PPC64/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:PPC64/Networking/Introduction/pl "Handbook:PPC64/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:PPC64/Networking/Advanced/pl "Handbook:PPC64/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:PPC64/Networking/Modular/pl "Handbook:PPC64/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:PPC64/Networking/Wireless/pl "Handbook:PPC64/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:PPC64/Networking/Extending/pl "Handbook:PPC64/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:PPC64/Networking/Dynamic/pl "Handbook:PPC64/Networking/Dynamic/pl")

[Handbook:PPC64/Blocks/Bootloader/pl](/index.php?title=Handbook:PPC64/Blocks/Bootloader/pl&action=edit&redlink=1 "Handbook:PPC64/Blocks/Bootloader/pl (page does not exist)")

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

Po uruchomieniu świeżo zainstalowanego środowiska Gentoo, kontynuuj [Finalizowanie instalacji Gentoo](/wiki/Handbook:PPC64/Installation/Finalizing "Handbook:PPC64/Installation/Finalizing").

[← Installing tools](/wiki/Handbook:PPC64/Installation/Tools/pl "Previous part") [Home](/wiki/Handbook:PPC64 "Handbook:PPC64") [Finalizing →](/wiki/Handbook:PPC64/Installation/Finalizing/pl "Next part")