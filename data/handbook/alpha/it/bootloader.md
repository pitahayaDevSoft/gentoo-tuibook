# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Bootloader/de "Handbuch:Alpha/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Bootloader "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Bootloader/es "Manual de Gentoo: Alpha/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Bootloader/fr "Manuel:Alpha/Installation/Système d'amorçage (100% translated)")
- italiano
- [magyar](/wiki/Handbook:Alpha/Installation/Bootloader/hu "Handbook:Alpha/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Bootloader/pl "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Bootloader/pt-br "Manual:Alpha/Instalação/Gerenciador de boot (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Bootloader/cs "Handbook:Alpha/Installation/Bootloader/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Bootloader/ru "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Bootloader/ta "கையேடு:Alpha/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Bootloader/zh-cn "手册：Alpha/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Bootloader/ja "ハンドブック:Alpha/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Bootloader/ko "Handbook:Alpha/Installation/Bootloader/ko (100% translated)")

[Manuale Alpha](/wiki/Handbook:Alpha/it "Handbook:Alpha/it")[Installazione](/wiki/Handbook:Alpha/Full/Installation/it "Handbook:Alpha/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:Alpha/Installation/About/it "Handbook:Alpha/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:Alpha/Installation/Media/it "Handbook:Alpha/Installation/Media/it")[Configurare la rete](/wiki/Handbook:Alpha/Installation/Networking/it "Handbook:Alpha/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:Alpha/Installation/Disks/it "Handbook:Alpha/Installation/Disks/it")[Il file stage](/wiki/Handbook:Alpha/Installation/Stage/it "Handbook:Alpha/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:Alpha/Installation/Base/it "Handbook:Alpha/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:Alpha/Installation/Kernel/it "Handbook:Alpha/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:Alpha/Installation/System/it "Handbook:Alpha/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:Alpha/Installation/Tools/it "Handbook:Alpha/Installation/Tools/it")Impostare programma d'avvio[Concludere](/wiki/Handbook:Alpha/Installation/Finalizing/it "Handbook:Alpha/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:Alpha/Full/Working/it "Handbook:Alpha/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:Alpha/Working/Portage/it "Handbook:Alpha/Working/Portage/it")[Opzioni USE](/wiki/Handbook:Alpha/Working/USE/it "Handbook:Alpha/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:Alpha/Working/Features/it "Handbook:Alpha/Working/Features/it")[Sistema script di init](/wiki/Handbook:Alpha/Working/Initscripts/it "Handbook:Alpha/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:Alpha/Working/EnvVar/it "Handbook:Alpha/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:Alpha/Full/Portage/it "Handbook:Alpha/Full/Portage/it")[File e cartelle](/wiki/Handbook:Alpha/Portage/Files/it "Handbook:Alpha/Portage/Files/it")[Variabili](/wiki/Handbook:Alpha/Portage/Variables/it "Handbook:Alpha/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:Alpha/Portage/Branches/it "Handbook:Alpha/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:Alpha/Portage/Tools/it "Handbook:Alpha/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:Alpha/Portage/CustomTree/it "Handbook:Alpha/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:Alpha/Portage/Advanced/it "Handbook:Alpha/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:Alpha/Full/Networking/it "Handbook:Alpha/Full/Networking/it")[Come iniziare](/wiki/Handbook:Alpha/Networking/Introduction/it "Handbook:Alpha/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:Alpha/Networking/Advanced/it "Handbook:Alpha/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:Alpha/Networking/Modular/it "Handbook:Alpha/Networking/Modular/it")[Wireless](/wiki/Handbook:Alpha/Networking/Wireless/it "Handbook:Alpha/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:Alpha/Networking/Extending/it "Handbook:Alpha/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:Alpha/Networking/Dynamic/it "Handbook:Alpha/Networking/Dynamic/it")

[Handbook:Alpha/Blocks/Bootloader/it](/index.php?title=Handbook:Alpha/Blocks/Bootloader/it&action=edit&redlink=1 "Handbook:Alpha/Blocks/Bootloader/it (page does not exist)")

## Riavviare il sistema

Uscire dall'ambiente con la radice cambiata (chroot) e smontare tutte le partizioni montate. Poi si inserisca quell'unico magico comando che avvia il vero test finale: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Non dimenticare di rimuovere l'immagine live, altrimenti potrebbe essere riavviata al posto del nuovo sistema Gentoo installato.

Una volta effettuato il riavvio nell'ambiente Gentoo appena installato, è ragionevole concludere con [Ultimare l'installazione di Gentoo](/wiki/Handbook:Alpha/Installation/Finalizing/it "Handbook:Alpha/Installation/Finalizing/it").

[← Installare gli strumenti](/wiki/Handbook:Alpha/Installation/Tools/it "Previous part") [Home](/wiki/Handbook:Alpha/it "Handbook:Alpha/it") [Concludere →](/wiki/Handbook:Alpha/Installation/Finalizing/it "Next part")