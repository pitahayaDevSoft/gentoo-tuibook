# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbuch:SPARC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Bootloader "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Bootloader/es "Manual de Gentoo: SPARC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Bootloader/fr "Handbook:SPARC/Installation/Bootloader/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Bootloader/pl "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Bootloader/pt-br "Handbook:SPARC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Bootloader/ru "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Bootloader/ta "கையேடு:SPARC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Bootloader/zh-cn "手册：SPARC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Bootloader/ja "ハンドブック:SPARC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Bootloader/ko "Handbook:SPARC/Installation/Bootloader/ko (100% translated)")

[Manuale SPARC](/wiki/Handbook:SPARC/it "Handbook:SPARC/it")[Installazione](/wiki/Handbook:SPARC/Full/Installation/it "Handbook:SPARC/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:SPARC/Installation/About/it "Handbook:SPARC/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:SPARC/Installation/Media/it "Handbook:SPARC/Installation/Media/it")[Configurare la rete](/wiki/Handbook:SPARC/Installation/Networking/it "Handbook:SPARC/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:SPARC/Installation/Disks/it "Handbook:SPARC/Installation/Disks/it")[Il file stage](/wiki/Handbook:SPARC/Installation/Stage/it "Handbook:SPARC/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:SPARC/Installation/Base/it "Handbook:SPARC/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:SPARC/Installation/Kernel/it "Handbook:SPARC/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:SPARC/Installation/System/it "Handbook:SPARC/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:SPARC/Installation/Tools/it "Handbook:SPARC/Installation/Tools/it")Impostare programma d'avvio[Concludere](/wiki/Handbook:SPARC/Installation/Finalizing/it "Handbook:SPARC/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:SPARC/Full/Working/it "Handbook:SPARC/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:SPARC/Working/Portage/it "Handbook:SPARC/Working/Portage/it")[Opzioni USE](/wiki/Handbook:SPARC/Working/USE/it "Handbook:SPARC/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:SPARC/Working/Features/it "Handbook:SPARC/Working/Features/it")[Sistema script di init](/wiki/Handbook:SPARC/Working/Initscripts/it "Handbook:SPARC/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:SPARC/Working/EnvVar/it "Handbook:SPARC/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:SPARC/Full/Portage/it "Handbook:SPARC/Full/Portage/it")[File e cartelle](/wiki/Handbook:SPARC/Portage/Files/it "Handbook:SPARC/Portage/Files/it")[Variabili](/wiki/Handbook:SPARC/Portage/Variables/it "Handbook:SPARC/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:SPARC/Portage/Branches/it "Handbook:SPARC/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:SPARC/Portage/Tools/it "Handbook:SPARC/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:SPARC/Portage/CustomTree/it "Handbook:SPARC/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:SPARC/Portage/Advanced/it "Handbook:SPARC/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:SPARC/Full/Networking/it "Handbook:SPARC/Full/Networking/it")[Come iniziare](/wiki/Handbook:SPARC/Networking/Introduction/it "Handbook:SPARC/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:SPARC/Networking/Advanced/it "Handbook:SPARC/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:SPARC/Networking/Modular/it "Handbook:SPARC/Networking/Modular/it")[Wireless](/wiki/Handbook:SPARC/Networking/Wireless/it "Handbook:SPARC/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:SPARC/Networking/Extending/it "Handbook:SPARC/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:SPARC/Networking/Dynamic/it "Handbook:SPARC/Networking/Dynamic/it")

[Handbook:SPARC/Blocks/Bootloader/it](/index.php?title=Handbook:SPARC/Blocks/Bootloader/it&action=edit&redlink=1 "Handbook:SPARC/Blocks/Bootloader/it (page does not exist)")

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

Una volta effettuato il riavvio nell'ambiente Gentoo appena installato, è ragionevole concludere con [Ultimare l'installazione di Gentoo](/wiki/Handbook:SPARC/Installation/Finalizing/it "Handbook:SPARC/Installation/Finalizing/it").

[← Installare gli strumenti](/wiki/Handbook:SPARC/Installation/Tools/it "Previous part") [Home](/wiki/Handbook:SPARC/it "Handbook:SPARC/it") [Concludere →](/wiki/Handbook:SPARC/Installation/Finalizing/it "Next part")