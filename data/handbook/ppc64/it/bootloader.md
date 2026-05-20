# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Bootloader/de "Handbuch:PPC64/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Bootloader "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Bootloader/es "Manual de Gentoo: PPC64/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Handbook:PPC64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Bootloader/pt-br "Handbook:PPC64/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Bootloader/ru "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Bootloader/ta "கையேடு:PPC64/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Bootloader/zh-cn "手册：PPC64/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Bootloader/ja "ハンドブック:PPC64/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko (100% translated)")

[Manuale PPC64](/wiki/Handbook:PPC64/it "Handbook:PPC64/it")[Installazione](/wiki/Handbook:PPC64/Full/Installation/it "Handbook:PPC64/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:PPC64/Installation/About/it "Handbook:PPC64/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:PPC64/Installation/Media/it "Handbook:PPC64/Installation/Media/it")[Configurare la rete](/wiki/Handbook:PPC64/Installation/Networking/it "Handbook:PPC64/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:PPC64/Installation/Disks/it "Handbook:PPC64/Installation/Disks/it")[Il file stage](/wiki/Handbook:PPC64/Installation/Stage/it "Handbook:PPC64/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:PPC64/Installation/Base/it "Handbook:PPC64/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:PPC64/Installation/Kernel/it "Handbook:PPC64/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:PPC64/Installation/System/it "Handbook:PPC64/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:PPC64/Installation/Tools/it "Handbook:PPC64/Installation/Tools/it")Impostare programma d'avvio[Concludere](/wiki/Handbook:PPC64/Installation/Finalizing/it "Handbook:PPC64/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:PPC64/Full/Working/it "Handbook:PPC64/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:PPC64/Working/Portage/it "Handbook:PPC64/Working/Portage/it")[Opzioni USE](/wiki/Handbook:PPC64/Working/USE/it "Handbook:PPC64/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:PPC64/Working/Features/it "Handbook:PPC64/Working/Features/it")[Sistema script di init](/wiki/Handbook:PPC64/Working/Initscripts/it "Handbook:PPC64/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:PPC64/Working/EnvVar/it "Handbook:PPC64/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:PPC64/Full/Portage/it "Handbook:PPC64/Full/Portage/it")[File e cartelle](/wiki/Handbook:PPC64/Portage/Files/it "Handbook:PPC64/Portage/Files/it")[Variabili](/wiki/Handbook:PPC64/Portage/Variables/it "Handbook:PPC64/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:PPC64/Portage/Branches/it "Handbook:PPC64/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:PPC64/Portage/Tools/it "Handbook:PPC64/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:PPC64/Portage/CustomTree/it "Handbook:PPC64/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:PPC64/Portage/Advanced/it "Handbook:PPC64/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:PPC64/Full/Networking/it "Handbook:PPC64/Full/Networking/it")[Come iniziare](/wiki/Handbook:PPC64/Networking/Introduction/it "Handbook:PPC64/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:PPC64/Networking/Advanced/it "Handbook:PPC64/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:PPC64/Networking/Modular/it "Handbook:PPC64/Networking/Modular/it")[Wireless](/wiki/Handbook:PPC64/Networking/Wireless/it "Handbook:PPC64/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:PPC64/Networking/Extending/it "Handbook:PPC64/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:PPC64/Networking/Dynamic/it "Handbook:PPC64/Networking/Dynamic/it")

Col kernel configurato e compilato e con gli indispensabili file di configurazione del sistema correttamente compilati, adesso è ora di installare un programma che avvii il kernel quando si accende il sistema. Tale programma è chiamato boot loader.

**Nota**

Attualmente l'utilizzo di Petitboot sui sistemi Talos non è documentato in Gentoo. Si prega di aggiungere i passaggi in [TalosII#Bootloader](/wiki/TalosII#Bootloader.2Fit "TalosII") e di darne notizia in questa pagina di Discussione quando sono pronti per essere uniti al manuale.

## Contents

- [1Usare GRUB](#Usare_GRUB)
  - [1.1Installare](#Installare)
  - [1.2Hardware Mac (G5)](#Hardware_Mac_.28G5.29)
    - [1.2.1Impostare la partizione d'avvio](#Impostare_la_partizione_d.27avvio)
    - [1.2.2Impostare GRUB](#Impostare_GRUB)
  - [1.3Hardware IBM](#Hardware_IBM)
    - [1.3.1Impostare GRUB](#Impostare_GRUB_2)
    - [1.3.2File di configurazione di GRUB](#File_di_configurazione_di_GRUB)
- [2Riavviare il sistema](#Riavviare_il_sistema)

## Usare GRUB

GRUB è un programma d'avvio per macchine Linux che montano PPC64.

### Installare

`root #` `emerge --ask sys-boot/grub`

### Hardware Mac (G5)

#### Impostare la partizione d'avvio

Prima, preparare la partizione d'avvio che è stata creata durante il passaggio [preparare il disco](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks"). Seguendo l'esempio, questa partizione dovrebbe essere /dev/sda2. Facoltativamente, confermare ciò usando parted:

Sostituire /dev/sda col dispositivo corretto se necessario.

`root #` `parted /dev/sda print`

```
Model: ATA Patriot Burst El (scsi)

Disk /dev/sda: 120GB

Sector size (logical/physical): 512B/512B

Partition Table: mac

Disk Flags:

Number  Start   End     Size    File system  Name       Flags
 1      512B    32.8kB  32.3kB               Apple
 2      32.8kB  852kB   819kB   hfs          bootstrap  boot
 3      852kB   538MB   537MB   ext4         Boot
 4      538MB   54.2GB  53.7GB  ext4         Gentoo

```

Da questi risultati, la partizione 2 risulta avviabile quindi /dev/sda2 è il percorso corretto da usare.

Formattare questa partizione come [HFS](/wiki/HFS "HFS") usando il comando hformat cha fa parte del pacchetto [sys-fs/hfsutils](https://packages.gentoo.org/packages/sys-fs/hfsutils):

`root #` `dd if=/dev/zero of=/dev/sda2 bs=512`

`root #` `hformat -l bootstrap /dev/sda2`

Creare una cartella per montare la partizione di avvio e montarla:

`root #` `mkdir /boot/NWBB
`

`root #` `mount --types hfs /dev/sda2 /boot/NWBB`

#### Impostare GRUB

`root #` `grub-install --macppc-directory=/boot/NWBB /dev/sda2`

Se viene installato senza errori, smontare la partizione di avvio:

`root #` `umount /boot/NWBB`

Successivamente, effettuare il "bless" della partizione cosicchè possa fare l'avvio:

`root #` `hmount /dev/sda2
`

`root #` `hattrib -t tbxi -c UNIX :System:Library:CoreServices:BootX
`

`root #` `hattrib -b :System:Library:CoreServices
`

`root #` `humount`

Infine, compilare il file grub.cfg:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

### Hardware IBM

Impostare GRUB sull'hardware IBM è semplice come:

#### Impostare GRUB

`root #` `grub-install /dev/sda1`

**Nota**

/dev/sda1 è la partizione d'avvio PReP creata durante la fase di partizionamento

#### File di configurazione di GRUB

Infine, compilare il file grub.cfg:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

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

Una volta effettuato il riavvio nell'ambiente Gentoo appena installato, è ragionevole concludere con [Ultimare l'installazione di Gentoo](/wiki/Handbook:PPC64/Installation/Finalizing/it "Handbook:PPC64/Installation/Finalizing/it").

[← Installare gli strumenti](/wiki/Handbook:PPC64/Installation/Tools/it "Previous part") [Home](/wiki/Handbook:PPC64/it "Handbook:PPC64/it") [Concludere →](/wiki/Handbook:PPC64/Installation/Finalizing/it "Next part")