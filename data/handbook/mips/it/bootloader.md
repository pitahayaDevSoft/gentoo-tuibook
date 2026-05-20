# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Bootloader/de "Handbuch:MIPS/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Bootloader "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Bootloader/es "Manual de Gentoo: MIPS/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Bootloader/fr "Handbook:MIPS/Installation/Bootloader/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Bootloader/pl "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Bootloader/pt-br "Manual:MIPS/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Bootloader/ru "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Bootloader/ta "கையேடு:MIPS/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Bootloader/zh-cn "手册：MIPS/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Bootloader/ja "ハンドブック:MIPS/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Bootloader/ko "Handbook:MIPS/Installation/Bootloader/ko (100% translated)")

[Manuale MIPS](/wiki/Handbook:MIPS/it "Handbook:MIPS/it")[Installazione](/wiki/Handbook:MIPS/Full/Installation/it "Handbook:MIPS/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:MIPS/Installation/About/it "Handbook:MIPS/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:MIPS/Installation/Media/it "Handbook:MIPS/Installation/Media/it")[Configurare la rete](/wiki/Handbook:MIPS/Installation/Networking/it "Handbook:MIPS/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:MIPS/Installation/Disks/it "Handbook:MIPS/Installation/Disks/it")[Il file stage](/wiki/Handbook:MIPS/Installation/Stage/it "Handbook:MIPS/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:MIPS/Installation/Base/it "Handbook:MIPS/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:MIPS/Installation/Kernel/it "Handbook:MIPS/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:MIPS/Installation/System/it "Handbook:MIPS/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:MIPS/Installation/Tools/it "Handbook:MIPS/Installation/Tools/it")Impostare programma d'avvio[Concludere](/wiki/Handbook:MIPS/Installation/Finalizing/it "Handbook:MIPS/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:MIPS/Full/Working/it "Handbook:MIPS/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:MIPS/Working/Portage/it "Handbook:MIPS/Working/Portage/it")[Opzioni USE](/wiki/Handbook:MIPS/Working/USE/it "Handbook:MIPS/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:MIPS/Working/Features/it "Handbook:MIPS/Working/Features/it")[Sistema script di init](/wiki/Handbook:MIPS/Working/Initscripts/it "Handbook:MIPS/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:MIPS/Working/EnvVar/it "Handbook:MIPS/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:MIPS/Full/Portage/it "Handbook:MIPS/Full/Portage/it")[File e cartelle](/wiki/Handbook:MIPS/Portage/Files/it "Handbook:MIPS/Portage/Files/it")[Variabili](/wiki/Handbook:MIPS/Portage/Variables/it "Handbook:MIPS/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:MIPS/Portage/Branches/it "Handbook:MIPS/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:MIPS/Portage/Tools/it "Handbook:MIPS/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:MIPS/Portage/CustomTree/it "Handbook:MIPS/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:MIPS/Portage/Advanced/it "Handbook:MIPS/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:MIPS/Full/Networking/it "Handbook:MIPS/Full/Networking/it")[Come iniziare](/wiki/Handbook:MIPS/Networking/Introduction/it "Handbook:MIPS/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:MIPS/Networking/Advanced/it "Handbook:MIPS/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:MIPS/Networking/Modular/it "Handbook:MIPS/Networking/Modular/it")[Wireless](/wiki/Handbook:MIPS/Networking/Wireless/it "Handbook:MIPS/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:MIPS/Networking/Extending/it "Handbook:MIPS/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:MIPS/Networking/Dynamic/it "Handbook:MIPS/Networking/Dynamic/it")

[Handbook:MIPS/Blocks/Bootloader/it](/index.php?title=Handbook:MIPS/Blocks/Bootloader/it&action=edit&redlink=1 "Handbook:MIPS/Blocks/Bootloader/it (page does not exist)")

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

Una volta effettuato il riavvio nell'ambiente Gentoo appena installato, è ragionevole concludere con [Ultimare l'installazione di Gentoo](/wiki/Handbook:MIPS/Installation/Finalizing/it "Handbook:MIPS/Installation/Finalizing/it").

[← Installare gli strumenti](/wiki/Handbook:MIPS/Installation/Tools/it "Previous part") [Home](/wiki/Handbook:MIPS/it "Handbook:MIPS/it") [Concludere →](/wiki/Handbook:MIPS/Installation/Finalizing/it "Next part")