# Disks

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Disks/de "Handbuch:HPPA/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Disks "Handbook:HPPA/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Disks/es "Manual de Gentoo: HPPA/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Disks/fr "Handbook:HPPA/Installation/Disks/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:HPPA/Installation/Disks/hu "Handbook:HPPA/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Disks/pl "Handbook:HPPA/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Disks/pt-br "Manual:HPPA/Instalação/Discos (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Disks/ru "Handbook:HPPA/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Disks/ta "கையேடு:HPPA/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Disks/zh-cn "手册：HPPA/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Disks/ja "ハンドブック:HPPA/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Disks/ko "Handbook:HPPA/Installation/Disks/ko (100% translated)")

[Manuale HPPA](/wiki/Handbook:HPPA/it "Handbook:HPPA/it")[Installazione](/wiki/Handbook:HPPA/Full/Installation/it "Handbook:HPPA/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:HPPA/Installation/About/it "Handbook:HPPA/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:HPPA/Installation/Media/it "Handbook:HPPA/Installation/Media/it")[Configurare la rete](/wiki/Handbook:HPPA/Installation/Networking/it "Handbook:HPPA/Installation/Networking/it")Preparare i dischi[Il file stage](/wiki/Handbook:HPPA/Installation/Stage/it "Handbook:HPPA/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:HPPA/Installation/Base/it "Handbook:HPPA/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:HPPA/Installation/Kernel/it "Handbook:HPPA/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:HPPA/Installation/System/it "Handbook:HPPA/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:HPPA/Installation/Tools/it "Handbook:HPPA/Installation/Tools/it")[Impostare programma d'avvio](/wiki/Handbook:HPPA/Installation/Bootloader/it "Handbook:HPPA/Installation/Bootloader/it")[Concludere](/wiki/Handbook:HPPA/Installation/Finalizing/it "Handbook:HPPA/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:HPPA/Full/Working/it "Handbook:HPPA/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:HPPA/Working/Portage/it "Handbook:HPPA/Working/Portage/it")[Opzioni USE](/wiki/Handbook:HPPA/Working/USE/it "Handbook:HPPA/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:HPPA/Working/Features/it "Handbook:HPPA/Working/Features/it")[Sistema script di init](/wiki/Handbook:HPPA/Working/Initscripts/it "Handbook:HPPA/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:HPPA/Working/EnvVar/it "Handbook:HPPA/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:HPPA/Full/Portage/it "Handbook:HPPA/Full/Portage/it")[File e cartelle](/wiki/Handbook:HPPA/Portage/Files/it "Handbook:HPPA/Portage/Files/it")[Variabili](/wiki/Handbook:HPPA/Portage/Variables/it "Handbook:HPPA/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:HPPA/Portage/Branches/it "Handbook:HPPA/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:HPPA/Portage/Tools/it "Handbook:HPPA/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:HPPA/Portage/CustomTree/it "Handbook:HPPA/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:HPPA/Portage/Advanced/it "Handbook:HPPA/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:HPPA/Full/Networking/it "Handbook:HPPA/Full/Networking/it")[Come iniziare](/wiki/Handbook:HPPA/Networking/Introduction/it "Handbook:HPPA/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:HPPA/Networking/Advanced/it "Handbook:HPPA/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:HPPA/Networking/Modular/it "Handbook:HPPA/Networking/Modular/it")[Wireless](/wiki/Handbook:HPPA/Networking/Wireless/it "Handbook:HPPA/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:HPPA/Networking/Extending/it "Handbook:HPPA/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:HPPA/Networking/Dynamic/it "Handbook:HPPA/Networking/Dynamic/it")

## Contents

- [1Introduzione ai dispositivi a blocchi](#Introduzione_ai_dispositivi_a_blocchi)
  - [1.1Dispositivi a blocchi](#Dispositivi_a_blocchi)
- [2Creare i file system](#Creare_i_file_system)
  - [2.1Introduzione](#Introduzione)
  - [2.2Filesystem](#Filesystem)
  - [2.3Applicare un filesystem ad una partizione](#Applicare_un_filesystem_ad_una_partizione)
    - [2.3.1Filesystem della partizione di boot su vecchi BIOS](#Filesystem_della_partizione_di_boot_su_vecchi_BIOS)
    - [2.3.2Partizioni ext4 piccole](#Partizioni_ext4_piccole)
  - [2.4Attivare la partizione di swap](#Attivare_la_partizione_di_swap)
- [3Montare la partizione di root](#Montare_la_partizione_di_root)

## Introduzione ai dispositivi a blocchi

### Dispositivi a blocchi

Si approfondiranno ora gli aspetti relativi ai dischi di Gentoo Linux e Linux in generale, compresi i dispositivi a blocchi, le partizioni e i filesystem di Linux. Appena tutto sarà chiaro in merito ai dischi, allora potranno essere stabiliti partizioni e filesystem per l'installazione.

Per iniziare, dare un'occhiata ai dispositivi a blocchi. I dischi SCSI e Serial ATA sono entrambi contrassegnati con gli appellativi di dispositivo del tipo: /dev/sda, /dev/sdb, /dev/sdc, etc. Sulle macchine più recenti i dischi a stato solido NVME basati su PCI Express hanno appellativi di dispositivo del tipo /dev/nvme0n1, /dev/nvme0n2, etc.

La tabella seguente aiuterà i lettori a stabilire dove trovare un certo tipo di dispositivo a blocchi nel sistema:

Tipo di dispositivoAppellativo di dispositivo predefinitoNote editoriali e considerazioni
IDE, SATA, SAS, SCSI, or USB flash/dev/sdaPresenti sull'hardware all'incirca dal 2007 ad oggi, questi appellativi di dispositivo sono forse i più comunenmente usati su Linux. Questi tipi di dispositivi possono essere connessi tramite il bus [SATA](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI"), [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB") come blocco di memorizzazione. Ad esempio, la prima partizione sul primo dispositivo SATA è chiamata /dev/sda1.
NVM Express (NVMe)/dev/nvme0n1Ultimi arrivati nella tecnologia a stato solido, i dischi [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") sono connessi al bus PCI Express ed hanno la maggiore velocità di trasferimento sul mercato. I sistemi all'incirca dal 2014 in poi potrebbero avere il supporto per l'hardware NVME. La prima partizione sul primo dispositivo NVME è chiamata /dev/nvme0n1p1.
MMC, eMMC, and SD/dev/mmcblk0I dispositivi [MMC integrati](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard"), le schede SD e altri tipi di schede di memoria possono essere utili per memorizzare dati. Detto questo molti sistemi potrebbero non consentire l'avvio da questo tipo di dispositivi. Si consiglia di **non** usare questi dispositivi per installazioni di Linux in attività; invece si consideri di usarli per trasferire file che è lo scopo principale per cui sono stati progettati. In alternativa questo tipo di memoria potrebbe essere utile per copie di breve periodo di file o istantanee del sistema.
VirtIO/dev/vdaI dispositivi [virtualizzati](/wiki/Virtualization "Virtualization") si trovano solo all'interno dell'emulatore virtuale [QEMU](/wiki/QEMU "QEMU"). Il driver `virtio_blk` fornisce questo accesso per lo spazio disco allocato dal server di questa macchina virtuale. Detto questo, è un ottimo modo per provare Gentoo all'interno di una macchina virtuale.

I dispositivi a blocchi di cui sopra rappresentano un'interfaccia astratta del disco. I programmi dell'utente possono usare questi dispositivi a blocchi per interagire col disco senza doversi preoccupare del fatto che i dischi siano SATA, SCSI o altro. Il programma può semplicemente indirizzare lo spazio sul disco come un insieme di blocchi contigui e accessibili in modo casuale da 4096 byte (4K).

[Handbook:HPPA/Blocks/Disks/it](/index.php?title=Handbook:HPPA/Blocks/Disks/it&action=edit&redlink=1 "Handbook:HPPA/Blocks/Disks/it (page does not exist)")

## Creare i file system

**Attenzione**

Usando un disco SSD o NVME, ha molto senso controllare eventuali aggiornamenti del firmware. Alcuni SSD Intel in particolare (600p e 6000p) necessitano dell'aggiornamento del firmware per [possibile corruzione dei dati](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) indotta dallo schema di I/O utilizzato da XFS. Il problema è a livello di firmware e non un errore nel filesystem XFS. Il programma di utilità smartctl può aiutare nella verifica del modello del dispositivo e della versione di firmware.

### Introduzione

Una volta che sono state create le partizioni, è ora di applicarci un filesystem. Nella sezione successiva vengono descritti i file system compatibili con Linux. I lettori che sanno già quale filesystem usare possono continuare con [applicare un filesystem ad una partizione](/wiki/Handbook:HPPA/Installation/Disks/it#Applicare_un_filesystem_ad_una_partizione "Handbook:HPPA/Installation/Disks/it"). Gli altri dovrebbero continuare a leggere per comprendere meglio i filesystem disponibili...

### Filesystem

Linux è compatibile con parecchie dozzine di filesystem, sebbene per molti di questi abbia senso impiegarli solo per scopi specifici. Solo alcuni filesystem potrebbero risultare stabili sull'architettura hppa - si consiglia di documentarsi sui filesystem e lo stato della loro compatibilità prima di selezionarne uno più sperimentale per partizioni importanti. **XFS è il filesystem consigliato per tutti gli scopi, su tutte le piattaforme.** Quello sotto è un elenco non esaustivo:

[XFS](/wiki/XFS/it "XFS/it")Filesystem con journaling dei metadata fornito di un robusto insieme di funzionalità ed ottimizzato per la scalabilità. E' stato aggiornato continuamente per includere funzionalità moderne. L'unico lato negativo è che le partizioni non possono essere ancora rimpicciolite, anche se ci si sta lavorando. XFS in particolare supporta reflinks e Copy on Write (CoW) che è particularmente utile nei sistemi Gentoo per via del numero di compilazioni che l'utente porta a termine. XFS è il filesystem moderno consigliato per tutti gli scopi su tutte le piattaforme. Necessita di una partizione di almeno 300MB.[ext4](/wiki/Ext4/it "Ext4/it")Ext4 è un filesystem affidabile, per tutti gli scopi su tutte le piattaforme, sebbene manchi di alcune funzionalità moderne come reflinks.[VFAT](/wiki/VFAT "VFAT")Anche conosciuto come FAT32, è compatibile con Linux ma non supporta le impostazioni dei permessi ordinari di UNIX. E' usato principalmente per interoperabilità/scambio con altri sistemi operativi (Microsoft Windows o MacOS di Apple) ma è anche una necessità per alcuni programmi di avvio del sistema nel firmware (come UEFI). Gli utenti di sistemi UEFI avranno bisogno di una [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition") formattata in VFAT per poter fare l'avvio.[btrfs](/wiki/Btrfs/it "Btrfs/it")Filesystem di generazione più nuova. Fornisce funzionalità avanzate come possibiltà di fare istantanee, auto-riparazione tramite checksum, compressione trasparente, sottovolumi e RAID integrato. I kernel di versione precedente alla 5.4.y non è garantito che siano sicuri da usare con btrfs nella quotidianità perchè delle soluzioni a problemi gravi sono presenti solo nelle versioni più recenti dei rami LTS del kernel. Il RAID 5/6 e i gruppi limite sono pericolosi su tutte le versioni di btrfs.[F2FS](/wiki/F2FS/it "F2FS/it")Il Flash-Friendly File System è stato inizialmente creato da Samsung per l'uso con le memorie flash NAND. E' una scelta discreta quando si installa Gentoo su schede microSD, dischi USB o altri dispositivi di memorizzazione basati sulle memorie flash.[NTFS](/wiki/NTFS "NTFS")Questo "New Technology" filesystem è il filesystem di riferimento di Microsoft Windows sin da Windows NT 3.1. In modo simile a VFAT, non memorizza le impostazioni dei permessi UNIX o gli attributi estesi necessari per BSD o Linux per funzionare adeguatamente, quindi non dovrebbe essere usato come filesystem di root nella maggior parte dei casi. Dovrebbe essere usato _solo_ per interoperabilità o scambio di dati coi sistemi Microsoft Windows (notare l'enfasi su _solo_).[ZFS](/wiki/ZFS "ZFS")**Importante:** I pool ZFS si possono creare sul CD di installazione minimale, per ulteriori informazioni, riferirsi a [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs")Filesystem di prossima generazione creato da Matthew Ahrens e Jeff Bonwick. E' stato progettato intorno ad alcune idee chiave: l' amministrazione della memoria dovrebbe essere semplice, la ridondanza dovrebbe essere gestita dal filesystem, i filesystem non dovrebbero mai essere posti offline per le riparazioni, le simulazioni automatiche degli scenari nei casi peggiori prima di distribuire il codice sono importanti e l'integrità dei dati è essenziale.

Informazioni più estese sui filesystem si possono trovare nell' [articolo filesystem](/wiki/Filesystem "Filesystem") curato dalla comunità.

### Applicare un filesystem ad una partizione

**Nota**

Si prega di accertarsi di installare (con emerge) il pacchetto con le utilità dello spazio utente pertinente prima di riavviare. Ci sarà un promemoria per farlo in prossimità della conclusione del processo di installazione.

Per creare un filesystem su una partizione o su un volume, ci sono programmi di utilità per lo spazio utente per ogni possibile filesystem. Cliccare sul nome del filesystem nella tabella sottostante per informazioni aggiuntive su ciascun filesystem:

Filesystem
Comando di creazione
All'interno dell'ambiente live?
Pacchetto
[XFS](/wiki/XFS/it "XFS/it")mkfs.xfs Si
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/it "Ext4/it")mkfs.ext4 Si
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
mkfs.vfat Si
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs/it "Btrfs/it")mkfs.btrfs Si
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS/it "F2FS/it")mkfs.f2fs Si
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs Si
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... Si
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

**Importante**

Il manuale consiglia nuove partizioni come parte del processo d'installazione, ma è importante notare che eseguire qualsiasi comando mkfs cancellerà tutti i dati all'interno della partizione. Quando necessario, ci si assicuri che qualsiasi dato esistente nella partizione venga opportunamente copiato prima di creare il filesystem nuovo.

Per esempio, per avere la partizione root (/dev/sda4) con xfs come nello schema delle partizioni d'esempio, si dovrebbero usare i seguenti comandi:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda4`

#### Filesystem della partizione di boot su vecchi BIOS

I sistemi che si avviano su vecchi BIOS con una tabella delle partizioni MBR/DOS possono usare qualsiasi formato del filesystem supportato dal programma di avvio (bootloader).

Per esempio, per formattare con XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda2`

#### Partizioni ext4 piccole

Quando si usa il filesystem ext4 su una piccola partizione (minore di 8GB), il filesystem dovrebbe essere creato con le opzioni appropriate per riservare abbastanza inode. Ciò può essere specificato usando l'opzione `-T small`:

`root #` `mkfs.ext4 -T small /dev/<dispositivo>`

Facendo ciò si quadruplicherà il numero di inode per un certo filesystem poiché i suoi "bytes-per-inode" si riducono da uno ogni 16kB a uno ogni 4kB.

### Attivare la partizione di swap

mkswap è il comando che viene usato per inizializzare la partizione di swap:

`root #` `mkswap /dev/sda3`

**Nota**

Le installazioni che sono state iniziate in precedenza, ma non hanno terminato tutto il processo possono riprendere da questo punto del manuale. Si usi questo link come permalink [Iniziare da qui per installazioni da riprendere](/wiki/Handbook:HPPA/Installation/Disks/it#Resumed_installations_start_here "Handbook:HPPA/Installation/Disks/it").

Per attivare la partizione di swap, usare swapon:

`root #` `swapon /dev/sda3`

Questo passaggio di 'attivazione' è necessario solo perchè la partizione di swap è creata nuovamente all' interno dell'ambiente live. Una volta riavviato il sistema, a patto che la partizione di swap sia stata ben definita all'interno di fstab o altro meccanismo di montaggio delle partizioni, lo spazio di swap si attiverà automaticamente.

## Montare la partizione di root

In alcuni ambienti live potrebbero mancare i punti di montaggio (mount) suggeriti per la partizione di root di Gentoo (/mnt/gentoo), o punti di montaggio per partizioni aggiuntive create nella sezione di partizionamento:

`root #` `mkdir --parents /mnt/gentoo`

Continuare a creare i punti di montaggio aggiuntivi necessari per ogni partizione (personalizzata) creata durante i passaggi precedenti, usando il comando mkdir.

Con i punti di montaggio creati, è tempo di rendere le partizioni accessibili tramite il comando mount.

Montare la partizione di root:

`root #` `mount /dev/sda4 /mnt/gentoo`

Continuare montando le partizioni aggiuntive (personalizzate) come necessario usando il comando mount.

**Nota**

Se è necessario che /tmp/ risieda su una partizione separata, assicurarsi di cambiare i suoi permessi dopo averla montata:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Ciò è valido anche per /var/tmp.

Più avanti nel manuale, il filesystem proc (un'interfaccia virtuale fornita dal kernel) ed altri pseudo-filesystem del kernel verranno montati. Però prima deve essere estratto il [file stage di Gentoo](/wiki/Handbook:HPPA/Installation/Stage/it "Handbook:HPPA/Installation/Stage/it").

[← Configurare la rete](/wiki/Handbook:HPPA/Installation/Networking/it "Previous part") [Home](/wiki/Handbook:HPPA "Handbook:HPPA") [Installare il file stage →](/wiki/Handbook:HPPA/Installation/Stage/it "Next part")