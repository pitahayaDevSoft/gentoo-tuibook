# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbuch:PPC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Bootloader "Handbook:PPC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Bootloader/es "Manual de Gentoo: PPC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Bootloader/fr "Handbook:PPC/Installation/Bootloader/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:PPC/Installation/Bootloader/hu "Handbook:PPC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Bootloader/pl "Handbook:PPC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Bootloader/pt-br "Handbook:PPC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Bootloader/ru "Handbook:PPC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Bootloader/ta "கையேடு:PPC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Bootloader/zh-cn "手册：PPC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Bootloader/ja "ハンドブック:PPC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Bootloader/ko "Handbook:PPC/Installation/Bootloader/ko (100% translated)")

[Manuale PPC](/wiki/Handbook:PPC/it "Handbook:PPC/it")[Installazione](/wiki/Handbook:PPC/Full/Installation/it "Handbook:PPC/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:PPC/Installation/About/it "Handbook:PPC/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:PPC/Installation/Media/it "Handbook:PPC/Installation/Media/it")[Configurare la rete](/wiki/Handbook:PPC/Installation/Networking/it "Handbook:PPC/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:PPC/Installation/Disks/it "Handbook:PPC/Installation/Disks/it")[Il file stage](/wiki/Handbook:PPC/Installation/Stage/it "Handbook:PPC/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:PPC/Installation/Base/it "Handbook:PPC/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:PPC/Installation/Kernel/it "Handbook:PPC/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:PPC/Installation/System/it "Handbook:PPC/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:PPC/Installation/Tools/it "Handbook:PPC/Installation/Tools/it")Impostare programma d'avvio[Concludere](/wiki/Handbook:PPC/Installation/Finalizing/it "Handbook:PPC/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:PPC/Full/Working/it "Handbook:PPC/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:PPC/Working/Portage/it "Handbook:PPC/Working/Portage/it")[Opzioni USE](/wiki/Handbook:PPC/Working/USE/it "Handbook:PPC/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:PPC/Working/Features/it "Handbook:PPC/Working/Features/it")[Sistema script di init](/wiki/Handbook:PPC/Working/Initscripts/it "Handbook:PPC/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:PPC/Working/EnvVar/it "Handbook:PPC/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:PPC/Full/Portage/it "Handbook:PPC/Full/Portage/it")[File e cartelle](/wiki/Handbook:PPC/Portage/Files/it "Handbook:PPC/Portage/Files/it")[Variabili](/wiki/Handbook:PPC/Portage/Variables/it "Handbook:PPC/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:PPC/Portage/Branches/it "Handbook:PPC/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:PPC/Portage/Tools/it "Handbook:PPC/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:PPC/Portage/CustomTree/it "Handbook:PPC/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:PPC/Portage/Advanced/it "Handbook:PPC/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:PPC/Full/Networking/it "Handbook:PPC/Full/Networking/it")[Come iniziare](/wiki/Handbook:PPC/Networking/Introduction/it "Handbook:PPC/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:PPC/Networking/Advanced/it "Handbook:PPC/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:PPC/Networking/Modular/it "Handbook:PPC/Networking/Modular/it")[Wireless](/wiki/Handbook:PPC/Networking/Wireless/it "Handbook:PPC/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:PPC/Networking/Extending/it "Handbook:PPC/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:PPC/Networking/Dynamic/it "Handbook:PPC/Networking/Dynamic/it")

## Contents

- [1Opzioni del programma d'avvio](#Opzioni_del_programma_d.27avvio)
- [2NewWorld Mac](#NewWorld_Mac)
  - [2.1GRUB](#GRUB)
  - [2.2Installazione](#Installazione)
  - [2.3Impostare partizione d'avvio](#Impostare_partizione_d.27avvio)
  - [2.4Configurare GRUB](#Configurare_GRUB)
- [3OldWorld Mac](#OldWorld_Mac)
  - [3.1BootX](#BootX)
- [4Pegasos](#Pegasos)
  - [4.1BootCreator](#BootCreator)
- [5Riavviare il sistema](#Riavviare_il_sistema)

## Opzioni del programma d'avvio

Ora che il kernel è configurato e compilato e i file di configurazione del sistema necessari sono stati compilati correttamente, è ora di installare un programma d'avvio del kernel quando si accenderà il sistema. Tale programma d'avvio è chiamato bootloader.

Il programma d'avvio da usare dipende dal tipo di macchina PPC.

Per macchine Apple NewWorld o IBM, bisogna selezionare grub. Le macchine Apple OldWorld hanno una opzione: BootX. I Pegasos non richiedono un programma d'avvio, ma è necessario fare l'emerge di bootcreator per creare i menù d'avvio SmartFirmware.

## NewWorld Mac

### GRUB

### Installazione

`root #` `emerge --ask sys-boot/grub`

### Impostare partizione d'avvio

Prima, preparare la partizione di avvio che è stata creata durante il passaggio [preparare i dischi](/wiki/Handbook:PPC/Installation/Disks/it "Handbook:PPC/Installation/Disks/it"). Secondo l'esempio questa partizione dovrebbe essere /dev/sda2. Facoltativamente, confermare ciò usando parted:

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

Da questi dati, la partizione 2 ha le informazioni di avvio quindi /dev/sda2 è il percorso giusto da usare.

Formattare questa partizione come [HFS](/wiki/HFS "HFS") usando il comando hformat cha fa parte del pacchetto [sys-fs/hfsutils](https://packages.gentoo.org/packages/sys-fs/hfsutils):

`root #` `dd if=/dev/zero of=/dev/sda2 bs=512`

`root #` `hformat -l bootstrap /dev/sda2`

Creare una cartella di mount per la partizione di avvio e poi farne il mount:

`root #` `mkdir /boot/NWBB
`

`root #` `mount --types hfs /dev/sda2 /boot/NWBB`

### Configurare GRUB

`root #` `grub-install --macppc-directory=/boot/NWBB /dev/sda2`

Se si installa senza errori, fare l'unmount della partizione di avvio:

`root #` `umount /boot/NWBB`

Quindi, fare il "bless" della partizione in modo che possa avviarsi:

`root #` `hmount /dev/sda2
`

`root #` `hattrib -t tbxi -c UNIX :System:Library:CoreServices:BootX
`

`root #` `hattrib -b :System:Library:CoreServices
`

`root #` `humount`

Infine, costruire il file grub.cfg:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

## OldWorld Mac

### BootX

**Importante**

BootX può essere usato solo su sistemi Apple OldWorld con MacOS classic da 7 a 9. Per macchine con meno di 32MB di RAM usare MacOS classic 7.

BootX può essere scaricato da [https://github.com/immolo/BootX/archive/refs/tags/1.2.2.tar.gz](https://github.com/immolo/BootX/archive/refs/tags/1.2.2.tar.gz)

Poichè BootX avvia Linux dall'interno di MacOS, il kernel necessiterà di essere copiato dalla partizione Linux a quella MacOS. Prima fare il mount della partizione MacOS dall'esterno della chroot. Usare mac-fdisk -l per trovare il numero della partizione MacOS, sda6 è usato come esempio qui. Una volta fatto il mount della partizione, si copierà il kernel nella Cartella di Sistema di MacOS in modo che BootX possa trovarlo.

`root #` `exit`

`cdimage ~#` `mkdir /mnt/mac
`

`cdimage ~#` `mount /dev/sda6 /mnt/mac -t hfs
`

`cdimage ~#` `cp /mnt/gentoo/usr/src/linux/vmlinux "/mnt/mac/System Folder/Linux Kernels/kernel-6.18.8-gentoo"`

Ora che si è finito di copiare il kernel, si avrà bisogno di riavviare per configurare BootX.

`cdimage ~#` `cd /
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/pts,/shm,}
`

`cdimage ~#` `umount -l /mnt/gentoo{/proc,/sys,}
`

`cdimage ~#` `umount -l /mnt/mac
`

`cdimage ~#` `reboot`

Naturalmente, non dimenticare di rimuovere il CD avviabile, altrimenti si avvierà nuovamente il CD invece di MacOS.

Una volta che la macchina si è avviata in MacOS, aprire il pannello di controllo di BootX. Quando non si usa genkernel, selezionare Options e togliere la spunta a "Use specified RAM disk". Se si usa genkernel, assicurarsi che l'initrd del genkernel venga selezionato invece di quello del CD. Se non si usa genkernel, c'è ora un opzione per specificare il disco e la partizione di root del sistema Linux. Compilare con i valori appropriati. A seconda della configurazione del kernel, potrebbe essere necessario applicare argomenti di avvio supplementari.

BootX può essere configurato per avviare Linux subito dopo l'accensione. Se ciò viene fatto, allora la macchina avvierà in MacOS prima e, durante la fase iniziale, BootX caricherà e avvierà Linux. Vedere la pagina internet di BootX per maggiori informazioni.

**Importante**

Ci si assicuri di includere nel kernel il supporto ai filesystem HFS e HFS+, altrimenti aggiornamenti o modifiche al kernel sulla partizione MacOS non saranno possibili.

## Pegasos

**Nota**

Pegasos ha anche il supporto per Grub ma attualmente ciò non è documentato in Gentoo. Si prega di aggiungerlo al wiki principale e di darne notizia nello spazio di discussione di questa pagina una volta pronto per essere trasferito qui.

### BootCreator

**Importante**

BootCreator compilerà un grazioso menù di avvio per SmartFirmware scritto in Forth per il Pegasos.

Prima assicurarsi di avere bootcreator installato sul sistema:

`root #` `emerge --ask sys-boot/bootcreator`

Ora copiare il file /etc/bootmenu.example in /etc/bootmenu/ e modificarlo per adattarlo alle necessità personali:

`root #` `cp /etc/bootmenu.example /etc/bootmenu
`

`root #` `nano -w /etc/bootmenu`

Sotto ecco un file di configurazione /etc/bootmenu completo. vmlinux e initrd dovrebbero essere sostituiti dai nomi delle immagini del kernel e dell'initrd.

FILE **`/etc/bootmenu`** **Esempio di file di configurazione di bootcreator**

```
#
# Descrizione del file di esempio per bootcreator 1.1
#

[VERSION]
1

[TITLE]
Boot Menu

[SETTINGS]
AbortOnKey = false
Timeout    = 9
Default    = 1

[SECTION]
Local HD -> Morphos      (Normal)
ide:0 boot2.img ramdebug edebugflags="logkprintf"

[SECTION]
Local HD -> Linux (Normal)
ide:0 kernel-6.18.8-gentoo video=radeonfb:1024x768@70 root=/dev/sda3

[SECTION]
Local HD -> Genkernel (Normal)
ide:0 kernel-genkernel-ppc-6.18.8-gentoo root=/dev/ram0
root=/dev/sda3 initrd=initramfs-genkernel-ppc-6.18.8-gentoo

```

Infine il menù di avvio deve essere convertito in Forth e copiato nella partizione di avvio, in modo che SmartFirmware lo possa leggere. Perciò è necessario eseguire bootcreator:

`root #` `bootcreator /etc/bootmenu /boot/menu`

**Nota**

Assicurarsi di dare un occhiata nelle impostazioni di SmartFirmware quando si riavvia, quel menù è il file che verrà caricato per impostazione predefinita.

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

Una volta effettuato il riavvio nell'ambiente Gentoo appena installato, è ragionevole concludere con [Ultimare l'installazione di Gentoo](/wiki/Handbook:PPC/Installation/Finalizing/it "Handbook:PPC/Installation/Finalizing/it").

[← Installare gli strumenti](/wiki/Handbook:PPC/Installation/Tools/it "Previous part") [Home](/wiki/Handbook:PPC/it "Handbook:PPC/it") [Concludere →](/wiki/Handbook:PPC/Installation/Finalizing/it "Next part")