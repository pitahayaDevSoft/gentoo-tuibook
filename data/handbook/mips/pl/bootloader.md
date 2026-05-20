# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Bootloader/de "Handbuch:MIPS/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Bootloader "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Bootloader/es "Manual de Gentoo: MIPS/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Bootloader/fr "Handbook:MIPS/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Bootloader/it "Handbook:MIPS/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Bootloader/pt-br "Manual:MIPS/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Bootloader/ru "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Bootloader/ta "கையேடு:MIPS/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Bootloader/zh-cn "手册：MIPS/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Bootloader/ja "ハンドブック:MIPS/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Bootloader/ko "Handbook:MIPS/Installation/Bootloader/ko (100% translated)")

[MIPS Handbook](/wiki/Handbook:MIPS/pl "Handbook:MIPS/pl")[Instalacja](/wiki/Handbook:MIPS/Full/Installation/pl "Handbook:MIPS/Full/Installation/pl")[O instalacji](/wiki/Handbook:MIPS/Installation/About/pl "Handbook:MIPS/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:MIPS/Installation/Media/pl "Handbook:MIPS/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:MIPS/Installation/Networking/pl "Handbook:MIPS/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:MIPS/Installation/Disks/pl "Handbook:MIPS/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:MIPS/Installation/Stage/pl "Handbook:MIPS/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:MIPS/Installation/Base/pl "Handbook:MIPS/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:MIPS/Installation/Kernel/pl "Handbook:MIPS/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:MIPS/Installation/System/pl "Handbook:MIPS/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:MIPS/Installation/Tools/pl "Handbook:MIPS/Installation/Tools/pl")Instalacja systemu rozruchowego[Finalizacja](/wiki/Handbook:MIPS/Installation/Finalizing/pl "Handbook:MIPS/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:MIPS/Full/Working/pl "Handbook:MIPS/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:MIPS/Working/Portage/pl "Handbook:MIPS/Working/Portage/pl")[Flagi USE](/wiki/Handbook:MIPS/Working/USE/pl "Handbook:MIPS/Working/USE/pl")[Funkcje portage](/wiki/Handbook:MIPS/Working/Features/pl "Handbook:MIPS/Working/Features/pl")[System initscript](/wiki/Handbook:MIPS/Working/Initscripts/pl "Handbook:MIPS/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:MIPS/Working/EnvVar/pl "Handbook:MIPS/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:MIPS/Full/Portage/pl "Handbook:MIPS/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:MIPS/Portage/Files/pl "Handbook:MIPS/Portage/Files/pl")[Zmienne](/wiki/Handbook:MIPS/Portage/Variables/pl "Handbook:MIPS/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:MIPS/Portage/Branches/pl "Handbook:MIPS/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:MIPS/Portage/Tools/pl "Handbook:MIPS/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:MIPS/Portage/CustomTree/pl "Handbook:MIPS/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:MIPS/Portage/Advanced/pl "Handbook:MIPS/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:MIPS/Full/Networking/pl "Handbook:MIPS/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:MIPS/Networking/Introduction/pl "Handbook:MIPS/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:MIPS/Networking/Advanced/pl "Handbook:MIPS/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:MIPS/Networking/Modular/pl "Handbook:MIPS/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:MIPS/Networking/Wireless/pl "Handbook:MIPS/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:MIPS/Networking/Extending/pl "Handbook:MIPS/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:MIPS/Networking/Dynamic/pl "Handbook:MIPS/Networking/Dynamic/pl")

[Handbook:MIPS/Blocks/Bootloader/pl](/index.php?title=Handbook:MIPS/Blocks/Bootloader/pl&action=edit&redlink=1 "Handbook:MIPS/Blocks/Bootloader/pl (page does not exist)")

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

Po uruchomieniu świeżo zainstalowanego środowiska Gentoo, kontynuuj [Finalizowanie instalacji Gentoo](/wiki/Handbook:MIPS/Installation/Finalizing "Handbook:MIPS/Installation/Finalizing").

[← Installing tools](/wiki/Handbook:MIPS/Installation/Tools/pl "Previous part") [Home](/wiki/Handbook:MIPS "Handbook:MIPS") [Finalizing →](/wiki/Handbook:MIPS/Installation/Finalizing/pl "Next part")