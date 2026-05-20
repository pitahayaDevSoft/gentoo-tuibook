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
- čeština
- [русский](/wiki/Handbook:Alpha/Installation/Bootloader/ru "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Bootloader/ta "கையேடு:Alpha/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Bootloader/zh-cn "手册：Alpha/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Bootloader/ja "ハンドブック:Alpha/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Bootloader/ko "Handbook:Alpha/Installation/Bootloader/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha/cs "Handbook:Alpha/cs")[Instalace](/wiki/Handbook:Alpha/Full/Installation/cs "Handbook:Alpha/Full/Installation/cs")[O instalaci](/wiki/Handbook:Alpha/Installation/About/cs "Handbook:Alpha/Installation/About/cs")[Výběr média](/wiki/Handbook:Alpha/Installation/Media/cs "Handbook:Alpha/Installation/Media/cs")[Konfigurace sítě](/wiki/Handbook:Alpha/Installation/Networking/cs "Handbook:Alpha/Installation/Networking/cs")[Příprava disků](/wiki/Handbook:Alpha/Installation/Disks/cs "Handbook:Alpha/Installation/Disks/cs")[Instalace stage3](/wiki/Handbook:Alpha/Installation/Stage/cs "Handbook:Alpha/Installation/Stage/cs")[Instalace základního systému](/wiki/Handbook:Alpha/Installation/Base/cs "Handbook:Alpha/Installation/Base/cs")[Konfigurace jádra](/wiki/Handbook:Alpha/Installation/Kernel/cs "Handbook:Alpha/Installation/Kernel/cs")[Konfigurace systému](/wiki/Handbook:Alpha/Installation/System/cs "Handbook:Alpha/Installation/System/cs")[Instalace nástrojů](/wiki/Handbook:Alpha/Installation/Tools/cs "Handbook:Alpha/Installation/Tools/cs")Konfigurace zavaděče[Dokončení](/wiki/Handbook:Alpha/Installation/Finalizing/cs "Handbook:Alpha/Installation/Finalizing/cs")[Práce s Gentoo](/wiki/Handbook:Alpha/Full/Working/cs "Handbook:Alpha/Full/Working/cs")[Úvod do Portage](/wiki/Handbook:Alpha/Working/Portage/cs "Handbook:Alpha/Working/Portage/cs")[Přepínače USE](/wiki/Handbook:Alpha/Working/USE/cs "Handbook:Alpha/Working/USE/cs")[Funkce portage](/wiki/Handbook:Alpha/Working/Features/cs "Handbook:Alpha/Working/Features/cs")[Systém initskriptů](/wiki/Handbook:Alpha/Working/Initscripts/cs "Handbook:Alpha/Working/Initscripts/cs")[Proměnné prostředí](/wiki/Handbook:Alpha/Working/EnvVar/cs "Handbook:Alpha/Working/EnvVar/cs")[Práce s Portage](/wiki/Handbook:Alpha/Full/Portage/cs "Handbook:Alpha/Full/Portage/cs")[Soubory a adresáře](/wiki/Handbook:Alpha/Portage/Files/cs "Handbook:Alpha/Portage/Files/cs")[Proměnné](/wiki/Handbook:Alpha/Portage/Variables/cs "Handbook:Alpha/Portage/Variables/cs")[Mísení softwarových větví](/wiki/Handbook:Alpha/Portage/Branches/cs "Handbook:Alpha/Portage/Branches/cs")[Doplňkové nástroje](/wiki/Handbook:Alpha/Portage/Tools/cs "Handbook:Alpha/Portage/Tools/cs")[Vlastní strom Portage](/wiki/Handbook:Alpha/Portage/CustomTree/cs "Handbook:Alpha/Portage/CustomTree/cs")[Pokročilé funkce](/wiki/Handbook:Alpha/Portage/Advanced/cs "Handbook:Alpha/Portage/Advanced/cs")[Konfigurace sítě](/wiki/Handbook:Alpha/Full/Networking/cs "Handbook:Alpha/Full/Networking/cs")[Začínáme](/wiki/Handbook:Alpha/Networking/Introduction/cs "Handbook:Alpha/Networking/Introduction/cs")[Pokročilá konfigurace](/wiki/Handbook:Alpha/Networking/Advanced/cs "Handbook:Alpha/Networking/Advanced/cs")[Modulární síťování](/wiki/Handbook:Alpha/Networking/Modular/cs "Handbook:Alpha/Networking/Modular/cs")[Bezdrátové sítě](/wiki/Handbook:Alpha/Networking/Wireless/cs "Handbook:Alpha/Networking/Wireless/cs")[Přidání funkcí](/wiki/Handbook:Alpha/Networking/Extending/cs "Handbook:Alpha/Networking/Extending/cs")[Dynamická správa](/wiki/Handbook:Alpha/Networking/Dynamic/cs "Handbook:Alpha/Networking/Dynamic/cs")

[Handbook:Alpha/Blocks/Bootloader/cs](/index.php?title=Handbook:Alpha/Blocks/Bootloader/cs&action=edit&redlink=1 "Handbook:Alpha/Blocks/Bootloader/cs (page does not exist)")

## Restartování systému

Opusťte prostředí chroot a odpojte všechny připojené oddíly. Po té napište kouzelný příkaz, který zahájí konečný, opravdový test: reboot.

`root #` `exit`

`cdimage ~#` `cd
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`cdimage ~#` `umount -R /mnt/gentoo
`

`cdimage ~#` `reboot`

Samozřejmě nezapomeňte vyjmout bootovací CD, jinak může dojít opět ke startu z CD namísto nově instalovaného systému Gentoo.

Jakmile jste zrestartovali do čerstvě nainstalovaného prostředí Gentoo, dokončete vše podle [Dokončení instalace Gentoo](/wiki/Handbook:Alpha/Installation/Finalizing "Handbook:Alpha/Installation/Finalizing").

[← Instalace nástrojů](/wiki/Handbook:Alpha/Installation/Tools/cs "Previous part") [Home](/wiki/Handbook:Alpha/cs "Handbook:Alpha/cs") [Dokončení →](/wiki/Handbook:Alpha/Installation/Finalizing/cs "Next part")