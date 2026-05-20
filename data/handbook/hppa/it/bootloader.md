# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbuch:HPPA/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Bootloader/es "Manual de Gentoo: HPPA/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Bootloader/pl "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Bootloader/pt-br "Manual:HPPA/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Bootloader/ru "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Bootloader/ta "கையேடு:HPPA/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Bootloader/zh-cn "手册：HPPA/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Bootloader/ja "ハンドブック:HPPA/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko (100% translated)")

[Manuale HPPA](/wiki/Handbook:HPPA/it "Handbook:HPPA/it")[Installazione](/wiki/Handbook:HPPA/Full/Installation/it "Handbook:HPPA/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:HPPA/Installation/About/it "Handbook:HPPA/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:HPPA/Installation/Media/it "Handbook:HPPA/Installation/Media/it")[Configurare la rete](/wiki/Handbook:HPPA/Installation/Networking/it "Handbook:HPPA/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:HPPA/Installation/Disks/it "Handbook:HPPA/Installation/Disks/it")[Il file stage](/wiki/Handbook:HPPA/Installation/Stage/it "Handbook:HPPA/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:HPPA/Installation/Base/it "Handbook:HPPA/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:HPPA/Installation/Kernel/it "Handbook:HPPA/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:HPPA/Installation/System/it "Handbook:HPPA/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:HPPA/Installation/Tools/it "Handbook:HPPA/Installation/Tools/it")Impostare programma d'avvio[Concludere](/wiki/Handbook:HPPA/Installation/Finalizing/it "Handbook:HPPA/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:HPPA/Full/Working/it "Handbook:HPPA/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:HPPA/Working/Portage/it "Handbook:HPPA/Working/Portage/it")[Opzioni USE](/wiki/Handbook:HPPA/Working/USE/it "Handbook:HPPA/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:HPPA/Working/Features/it "Handbook:HPPA/Working/Features/it")[Sistema script di init](/wiki/Handbook:HPPA/Working/Initscripts/it "Handbook:HPPA/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:HPPA/Working/EnvVar/it "Handbook:HPPA/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:HPPA/Full/Portage/it "Handbook:HPPA/Full/Portage/it")[File e cartelle](/wiki/Handbook:HPPA/Portage/Files/it "Handbook:HPPA/Portage/Files/it")[Variabili](/wiki/Handbook:HPPA/Portage/Variables/it "Handbook:HPPA/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:HPPA/Portage/Branches/it "Handbook:HPPA/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:HPPA/Portage/Tools/it "Handbook:HPPA/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:HPPA/Portage/CustomTree/it "Handbook:HPPA/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:HPPA/Portage/Advanced/it "Handbook:HPPA/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:HPPA/Full/Networking/it "Handbook:HPPA/Full/Networking/it")[Come iniziare](/wiki/Handbook:HPPA/Networking/Introduction/it "Handbook:HPPA/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:HPPA/Networking/Advanced/it "Handbook:HPPA/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:HPPA/Networking/Modular/it "Handbook:HPPA/Networking/Modular/it")[Wireless](/wiki/Handbook:HPPA/Networking/Wireless/it "Handbook:HPPA/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:HPPA/Networking/Extending/it "Handbook:HPPA/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:HPPA/Networking/Dynamic/it "Handbook:HPPA/Networking/Dynamic/it")

[Handbook:HPPA/Blocks/Bootloader/it](/index.php?title=Handbook:HPPA/Blocks/Bootloader/it&action=edit&redlink=1 "Handbook:HPPA/Blocks/Bootloader/it (page does not exist)")

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

Una volta effettuato il riavvio nell'ambiente Gentoo appena installato, è ragionevole concludere con [Ultimare l'installazione di Gentoo](/wiki/Handbook:HPPA/Installation/Finalizing/it "Handbook:HPPA/Installation/Finalizing/it").

[← Installare gli strumenti](/wiki/Handbook:HPPA/Installation/Tools/it "Previous part") [Home](/wiki/Handbook:HPPA/it "Handbook:HPPA/it") [Concludere →](/wiki/Handbook:HPPA/Installation/Finalizing/it "Next part")