# Disks

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Disks/de "Handbuch:X86/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Disks "Handbook:X86/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:X86/Installation/Disks/tr "Handbook:X86/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:X86/Installation/Disks/es "Manual de Gentoo: X86/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Disks/fr "Handbook:X86/Installation/Disks/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:X86/Installation/Disks/hu "Handbook:X86/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Disks/pl "Handbook:X86/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Disks/pt-br "Manual:X86/Instalação/Discos (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Disks/cs "Handbook:X86/Installation/Disks/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Disks/ru "Handbook:X86/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Disks/ta "கையேடு:X86/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Disks/zh-cn "手册：X86/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Disks/ja "ハンドブック:X86/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Disks/ko "Handbook:X86/Installation/Disks/ko (100% translated)")

[Manuale X86](/wiki/Handbook:X86/it "Handbook:X86/it")[Installazione](/wiki/Handbook:X86/Full/Installation/it "Handbook:X86/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:X86/Installation/About/it "Handbook:X86/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:X86/Installation/Media/it "Handbook:X86/Installation/Media/it")[Configurare la rete](/wiki/Handbook:X86/Installation/Networking/it "Handbook:X86/Installation/Networking/it")Preparare i dischi[Il file stage](/wiki/Handbook:X86/Installation/Stage/it "Handbook:X86/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:X86/Installation/Base/it "Handbook:X86/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:X86/Installation/Kernel/it "Handbook:X86/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:X86/Installation/System/it "Handbook:X86/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:X86/Installation/Tools/it "Handbook:X86/Installation/Tools/it")[Impostare programma d'avvio](/wiki/Handbook:X86/Installation/Bootloader/it "Handbook:X86/Installation/Bootloader/it")[Concludere](/wiki/Handbook:X86/Installation/Finalizing/it "Handbook:X86/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:X86/Full/Working/it "Handbook:X86/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:X86/Working/Portage/it "Handbook:X86/Working/Portage/it")[Opzioni USE](/wiki/Handbook:X86/Working/USE/it "Handbook:X86/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:X86/Working/Features/it "Handbook:X86/Working/Features/it")[Sistema script di init](/wiki/Handbook:X86/Working/Initscripts/it "Handbook:X86/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:X86/Working/EnvVar/it "Handbook:X86/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:X86/Full/Portage/it "Handbook:X86/Full/Portage/it")[File e cartelle](/wiki/Handbook:X86/Portage/Files/it "Handbook:X86/Portage/Files/it")[Variabili](/wiki/Handbook:X86/Portage/Variables/it "Handbook:X86/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:X86/Portage/Branches/it "Handbook:X86/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:X86/Portage/Tools/it "Handbook:X86/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:X86/Portage/CustomTree/it "Handbook:X86/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:X86/Portage/Advanced/it "Handbook:X86/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:X86/Full/Networking/it "Handbook:X86/Full/Networking/it")[Come iniziare](/wiki/Handbook:X86/Networking/Introduction/it "Handbook:X86/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:X86/Networking/Advanced/it "Handbook:X86/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:X86/Networking/Modular/it "Handbook:X86/Networking/Modular/it")[Wireless](/wiki/Handbook:X86/Networking/Wireless/it "Handbook:X86/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:X86/Networking/Extending/it "Handbook:X86/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:X86/Networking/Dynamic/it "Handbook:X86/Networking/Dynamic/it")

## Contents

- [1Introduzione ai dispositivi a blocchi](#Introduzione_ai_dispositivi_a_blocchi)
  - [1.1Dispositivi a blocchi](#Dispositivi_a_blocchi)
  - [1.2Tabelle delle partizioni](#Tabelle_delle_partizioni)
    - [1.2.1GPT](#GPT)
    - [1.2.2MBR](#MBR)
  - [1.3Partizionamento avanzato](#Partizionamento_avanzato)
  - [1.4Schema di partizionamento predefinito](#Schema_di_partizionamento_predefinito)
- [2Progettare uno schema delle partizioni](#Progettare_uno_schema_delle_partizioni)
  - [2.1Quante partizioni e quanto grandi?](#Quante_partizioni_e_quanto_grandi.3F)
  - [2.2Riguardo lo spazio di swap?](#Riguardo_lo_spazio_di_swap.3F)
    - [2.2.1Uso di UEFI](#Uso_di_UEFI)
  - [2.3Cos'è la partizione di avvio BIOS?](#Cos.27.C3.A8_la_partizione_di_avvio_BIOS.3F)
- [3Alternativa: Uso di fdisk per partizionare il disco](#Alternativa:_Uso_di_fdisk_per_partizionare_il_disco)
  - [3.1Visualizzare lo schema delle partizioni correnti con fdisk](#Visualizzare_lo_schema_delle_partizioni_correnti_con_fdisk)
  - [3.2Rimuovere tutte le partizioni con fdisk](#Rimuovere_tutte_le_partizioni_con_fdisk)
  - [3.3Creare la partizione di avvio BIOS](#Creare_la_partizione_di_avvio_BIOS)
  - [3.4Creare la partizione di swap](#Creare_la_partizione_di_swap)
  - [3.5Creare la partizione di avvio](#Creare_la_partizione_di_avvio)
  - [3.6Creare la partizione radice](#Creare_la_partizione_radice)
  - [3.7Salvare lo schema delle partizioni](#Salvare_lo_schema_delle_partizioni)
- [4Partitioning the disk with MBR for BIOS / legacy boot](#Partitioning_the_disk_with_MBR_for_BIOS_.2F_legacy_boot)
  - [4.1Viewing the current partition layout](#Viewing_the_current_partition_layout)
  - [4.2Creating a new disklabel / removing all partitions](#Creating_a_new_disklabel_.2F_removing_all_partitions)
  - [4.3Creating the boot partition](#Creating_the_boot_partition)
  - [4.4Creating the swap partition](#Creating_the_swap_partition)
  - [4.5Creating the root partition](#Creating_the_root_partition)
  - [4.6Saving the partition layout](#Saving_the_partition_layout)
- [5Creare i file system](#Creare_i_file_system)
  - [5.1Introduzione](#Introduzione)
  - [5.2Filesystem](#Filesystem)
  - [5.3Applicare un filesystem ad una partizione](#Applicare_un_filesystem_ad_una_partizione)
    - [5.3.1Filesystem della partizione di boot su vecchi BIOS](#Filesystem_della_partizione_di_boot_su_vecchi_BIOS)
    - [5.3.2Partizioni ext4 piccole](#Partizioni_ext4_piccole)
  - [5.4Attivare la partizione di swap](#Attivare_la_partizione_di_swap)
- [6Montare la partizione di root](#Montare_la_partizione_di_root)

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

### Tabelle delle partizioni

Benché per ospitare un sistema Linux sia teoricamente possibile usare un disco grezzo e non partizionato (quando si crea un RAID btrfs per esempio), praticamente ciò non viene mai fatto. Piuttosto, i dischi vengono suddivisi in unità a blocchi più piccole e maneggevoli. Su sistemi **x86**, esse sono chiamate partizioni. Attualmente sono standardizzate due tecnologie di partizionamento: MBR e GPT.

#### GPT

La configurazione tramite _GPT (GUID Partition Table)_ usa identificatori a 64 bit per le partizioni. Lo spazio dove memorizza le informazioni sulle partizioni è molto più grande dei 512 byte dell'MBR, il ché significa che non c'è praticamente alcun limite alla quantità di partizioni definibili su un disco GPT. Inoltre, il limite per la dimensione massima di una partizione è di gran lunga maggiore (quasi 8 ZB - sì, zettabytes).

Quando l'interfaccia software del sistema, che si pone tra il sistema operativo e il firmware, è UEFI (anziché BIOS), GPT è quasi obbligatoria in quanto potrebbero sorgere problemi di compatibilità con MBR.

GPT trae anche vantaggio dalle somme di controllo (checksum) e dalla ridondanza. Porta il controllo CRC32 alla testata delle tabelle di partizione per rilevare errori ed offre un backup del segmento GPT alla fine del disco. Questo backup può essere usato per ripristinare i danni del segmento GPT corrente all'inizio del disco.

**Importante**

There are a few caveats regarding GPT:

- Using GPT on a BIOS-based computer works, but the user won't be able to dual-boot with a Microsoft Windows operating system, since Microsoft Windows refuses to boot from a GPT partition when in BIOS mode.
- Some buggy (old) motherboard firmware configured to boot in BIOS/CSM/legacy mode might also have problems with booting from GPT labeled disks.

#### MBR

La configurazione tramite _MBR (Master Boot Record)_ usa identificatori a 32 bit per il settore di avvio e per stabilire la grandezza delle partizioni. Supporta tre tipi di partizione: primaria, estesa e logica. Le partizioni primarie memorizzano le loro informazioni nel master boot record stesso - uno spazio molto piccolo (solitamente 512 byte) all'inizio del disco. A causa del poco spazio, vengono supportate solo quattro partizioni primarie (per esempio, da /dev/sda1 a /dev/sda4).

Per supportare più partizioni, una delle partizioni primarie può essere definita come partizione estesa. Questa partizione può in tal caso contenere a sua volta delle partizioni logiche (partizioni all'interno di una partizione).

**Importante**

Although still supported by most motherboard manufacturers, MBR boot sectors and their associated partitioning limitations are considered legacy. Unless working with hardware that is pre-2010, it best to partition a disk with [GUID Partition Table](https://en.wikipedia.org/wiki/GUID_Partition_Table "wikipedia:GUID Partition Table"). Readers who must proceed with setup type should knowingly acknowledge the following information:

- Most post-2010 motherboards consider using MBR boot sectors a legacy (supported, but not ideal) boot mode.
- Due to using 32-bit identifiers, partition tables in the MBR cannot address storage space that is larger than 2 TiBs in size.
- Unless an extended partition is created, MBR supports a maximum of four partitions.
- This setup does not provide a backup boot sector, so if something overwrites the partition table, all partition information will be lost.

That said, MBR and legacy BIOS boot may still used in virtualized cloud environments such as AWS.

The Handbook authors suggest using [GPT](#GPT) whenever possible for Gentoo installations.

### Partizionamento avanzato

I CD di installazione **x86** forniscono supporto per il gestore dei volumi logici (LVM). LVM accresce la flessibilità offerta dalla configurazione di partizionamento. Le istruzioni di installazione riportate di seguito si concentrano su partizioni "regolari", ma è bene sapere che anche LVM è supportato se si desidera proseguire per quella strada. Leggere l'articolo [LVM](/wiki/LVM/it "LVM/it") per ulteriori dettagli. I nuovi arrivati stiano attenti: benché LVM sia completamente supportato, va al di là dello scopo di questa guida.

Although usage is not covered in the handbook, below is a list helpful guides to get the system running:

- [Btrfs](/wiki/Btrfs/Native_System_Root_Guide "Btrfs/Native System Root Guide")
- [ZFS](/wiki/ZFS/rootfs "ZFS/rootfs")
- [LVM (Logical Volume Manager)](/wiki/LVM/it "LVM/it")

Disk encryption is also handled in the same manner:

- [Rootfs encryption](/wiki/Rootfs_encryption "Rootfs encryption")

### Schema di partizionamento predefinito

Throughout the remainder of the handbook, we will discuss and explain two cases:

1. UEFI firmware with GUID Partition Table (GPT) disk.
2. MBR DOS/legacy BIOS firmware with a MBR partition table disk.

While it is possible to mix and match boot types with certain motherboard firmware, mixing goes beyond the intention of the handbook. As previously stated, it is strongly recommended for installations on modern hardware to use UEFI boot with a GPT disklabel disk.

Per tutto il resto del manuale, verrà usato il seguento schema di partizionamento come esempio semplice di configurazione:

**Importante**

The first row of the following table contains exclusive information for _**either**_ a GPT disklabel _**or**_ a MBR DOS/legacy BIOS disklabel. When in doubt, proceed with GPT, since **x86** machines manufactured after the year 2010 generally support UEFI firmware and GPT boot sector.

Partizione
Filesystem
Dimensione
Descrizione
/dev/sda1(bootloader)
2M
Partizione di avvio BIOS
/dev/sda1ext2 (o fat32 se si utilizza UEFI)
128M
Partizione di sistema Boot/EFI
/dev/sda3(swap)
512M o maggiore
Partizione di swap
/dev/sda4ext4
Spazio rimanente del disco
Partizione radice (root)

Se ciò è sufficiente e il lettore ha scelto la configurazione GPT, si può proseguire con la sezione [Predefinito: Uso di parted per partizionare il disco](#Predefinito:_Uso_di_parted_per_partizionare_il_disco). Coloro che sono ancora interessati a MBR (ehi, capita!) e vogliono usare la configurazione d'esempio, possono proseguire con l' [Alternativa: Uso di fdisk per partizionare il disco](#Alternativa:_Uso_di_fdisk_per_partizionare_il_disco).

Sia fdisk che parted sono utilità di partizionamento. fdisk è ben noto, stabile, e raccomandato per la configurazione di partizionamento MBR, mentre parted è stata una delle prima utilità di gestione dei dispositivi a blocchi Linux a supportare le partizioni GPT. Coloro a cui piace l'interfaccia di fdisk possono usare gdisk (fdisk GPT) come alternativa a parted.

Prima di proseguire con le istruzioni di creazione, il primo insieme di sezioni descriverà con maggiori dettagli come si possono creare schemi di partizionamento e si menzioneranno alcune trappole comuni.

## Progettare uno schema delle partizioni

### Quante partizioni e quanto grandi?

Il progetto della struttura delle partizioni del disco dipende fortemente dalle esigenze del sistema e dal file system applicato al dispositivo. Se sono previsti molti utenti, allora è consigliato avere la /home su una partizione separata che aumenterà la sicurezza e renderà più facili i backup e altri tipi di manutenzione. Se Gentoo viene installato per fare da server di posta elettronica, allora /var dovrebbe essere una partizione separata in quanto tutte le email vengono memorizzate all'interno della cartella /var. I server da gioco potrebbero avere una partizione /opt separata, in quando la maggior parte dei software dei server da gioco sono installati là dentro. La ragione per queste raccomandazioni è simile a quella della cartella /home: sicurezza, backup e manutenzione.

Nella maggior parte dei casi su Gentoo, /usr e /var dovrebbero essere tenuti di dimensioni relativamente grandi. /usr ospita la maggioranza delle applicazioni disponibili sul sistema e i sorgenti del kernel Linux (in /usr/src). Di norma, /var ospita il catalogo delle ebuild Gentoo (situato in /var/db/repos/gentoo) che, a seconda del file system, in genere consuma circa 650 MiB di spazio su disco. Questo spazio stimato _esclude_ le cartelle /var/cache/distfiles e /var/cache/binpkgs, che vengono gradualmente riempite coi file sorgenti e (opzionalmente) coi pacchetti binari rispettivamente quando vengono aggiunti al sistema.

Quante partizioni e quanto grandi dipende molto dalla valutazione dei compromessi e dalla scelta dell'opzione migliore per la circostanza. Partizioni o volumi separati hanno i seguenti vantaggi:

- Scelta del filesystem dalle migliori prestazioni per ciascuna partizione o volume.
- L'intero sistema non esaurirà lo spazio, nel caso in cui uno strumento non più attivo continui a scrivere file su una partizione o un volume.
- Se necessario, i controlli del filesystem impiegheranno meno tempo, in quanto possono essere fatti in parallelo (benché questo vantaggio venga realizzato più con dischi multipli piuttosto che con partizioni multiple).
- La sicurezza può essere migliorata montando alcune partizioni o volumi in modalità di sola lettura, `nosuid` (ignora i bit setuid), `noexec` (ignora i bit eseguibili), ecc.

Tuttavia, avere partizioni multiple presenta anche certi svantaggi quali:

- Se il sistema non viene appropriatamente configurato, potrebbe avere molto spazio libero su una partizione e poco su un'altra.
- Una partizione separata per /usr/ potrebbe richiedere all'amministratore di avviare il sistema con un initramfs per montare la partizione prima che altri script all'avvio vengano eseguiti. Poichè la creazione e la manutenzione di un initramfs va al di là dello scopo di questo manuale, **consigliamo ai novizi di non usare una partizione separata per /usr/**.
- C'è anche un limite di 15 partizioni per SCSI e SATA a meno che il disco non usi una tabella GPT.

**Nota**

Le installazioni che intendono usare systemd come sistema di inizializzzione e dei servizi devono avere la cartella /usr disponibile all'avvio, o come elemento del file system di root o montata tramite un initramfs.

### Riguardo lo spazio di swap?

Consigli sulla dimensione dello spazio di swap
RAM installataSupporto alla sospensione?Supporto all' ibernazione ?
2 GB o meno4GB4GB
2 fino a 8 GBRAM amount2 \* RAM
8 fino a 64 GB8 GB minimo, 16 massimo1.5 \* RAM
64 GB o più8 GB minimoIbernazione non consigliata! Ibernazione non è consigliata per sistemi con quantità di memoria molto grandi. Anche se possibile, l' intero contenuto della memoria deve essere scritto sul disco affinchè l'ibernazione riesca. Scrivere decine di gigabytes (o peggio!) sul disco può richiedere un considerevole lasso di tempo, specialmente usando dischi meccanici. E' meglio la sospensione in questo scenario.

Non c'è un valore perfetto per la dimensione dello spazio di swap. Lo scopo dello spazio di swap è quello di fornire spazio sul disco al kernel quando la memoria dinamica interna (RAM) è sotto pressione. Uno spazio di swap permette al kernel di spostare le pagine di memoria, che probabilmente non verranno utilizzate entro breve tempo, sul disco (swap o spaginazione), liberando memoria nella RAM per l'operazione corrente. Ovviamente, se le pagine spostate sul disco sono improvvisamente necessarie, vengono rimesse nella memoria (paginazione), il che impiegherà notevolmente di più che leggere dalla RAM (dato che i dischi sono molto lenti se paragonati alla memoria interna).

Quando un sistema non esegue applicazioni che occupano molta memoria oppure ha tantissima RAM disponibile, allora è probabile che non serva molto spazio di swap. Comunque si noti che in caso di ibernazione lo spazio di swap è impiegato per memorizzare _l'intero contenuto della memoria_ (probabile su sistemi desktop o portatili piuttosto che sistemi server). Se il sistema richiede il supporto all'ibernazione, allora è necessario uno spazio di swap più grande o pari alla quantità di memoria.

Come regola generale per quantitativi di RAM minori di 4GB, la dimensione consigliata dello spazio di swap è due volte la memoria interna (RAM). Per sistemi con dischi multipli, è saggio creare una partizione di swap su ogni disco cosicchè possano essere utilizzati per operazioni di lettura/scrittura in parallelo. Più veloce il disco potrà fare lo swap, più veloce il sistema funzionerà quando dovrà accedere a dati nello spazio di swap. Quando si sceglie tra dischi a stato solido o meccanici, è meglio per le prestazioni mettere lo swap sull'hardware a stato solido.

Vale la pena notare che _file_ di swap possono essere usati come alternativa alle _partizioni_ di swap; ciò è utile soprattutto per sistemi con spazio su disco molto limitato.

#### Uso di UEFI

Quando si installa Gentoo su un sistema che utilizza UEFI per avviare il sistema operativo (invece di BIOS), allora è importante creare una Partizione di Sistema EFI (ESP). Le istruzioni per parted di seguito contengono i puntatori necessari per gestire questa operazione correttamente.

La partizione ESP _deve_ essere una variante di FAT (talvolta mostrata come _vfat_ sui sistemi Linux). Le [specifiche UEFI](http://www.uefi.org/sites/default/files/resources/UEFI%202_5.pdf) ufficiali dichiarano che i filesystem FAT12, 16 o 32 vengono riconosciuti dal firmware UEFI, benché sia raccomandato FAT32 per la ESP. Procedere con la formattazione della ESP in FAT32:

`root #` `mkfs.fat -F 32 /dev/sda1`

**Importante**

Se non viene usata una variante FAT per l'ESP, non è garantito che il firmware UEFI di sistema trovi il bootloader (o il kernel Linux) e probabilmente non sarà in grado di avviare il sistema!

### Cos'è la partizione di avvio BIOS?

Una partizione di avvio BIOS è una partizione molto piccola (da 1 a 2 MB) in cui i bootloader come GRUB2 possono inserire dati aggiuntivi se non riescono a stare nello spazio allocato (poche centinaia di byte nel caso di MBR) e se non possono stare altrove.

## Alternativa: Uso di fdisk per partizionare il disco

La seguente parte spiega come impostare le partizioni secondo l'esempio usando fdisk. Lo schema delle partizioni d'esempio menzionato prima:

Modificare lo schema di partizionamento in base alle proprie personali preferenze.

Partizione
Descrizione
/dev/sda1Partizione di avvio BIOS
/dev/sda1Partizione di avvio
/dev/sda3Partizione di swap
/dev/sda4Partizione radice (root)

### Visualizzare lo schema delle partizioni correnti con fdisk

fdisk è un famoso e potente strumento per dividere un disco in partizioni. Lanciare fdisk per il disco (nel nostro esempio usiamo /dev/sda):

`root #` `fdisk /dev/sda`

Usare il tasto `p` per visualizzare l'attuale configurazione delle partizioni sul disco:

`Command (m for help):` `p`

```
Disk /dev/sda: 240 heads, 63 sectors, 2184 cylinders
Units = cylinders of 15120 * 512 bytes

   Device Boot    Start       End    Blocks   Id  System
/dev/sda1   *         1        14    105808+  83  Linux
/dev/sda2            15        49    264600   82  Linux swap
/dev/sda3            50        70    158760   83  Linux
/dev/sda4            71      2184  15981840    5  Extended
/dev/sda5            71       209   1050808+  83  Linux
/dev/sda6           210       348   1050808+  83  Linux
/dev/sda7           349       626   2101648+  83  Linux
/dev/sda8           627       904   2101648+  83  Linux
/dev/sda9           905      2184   9676768+  83  Linux

```

Device Start End Sectors Size Type
/dev/sda1 2048 2099199 2097152 1G EFI System
/dev/sda2 2099200 10487807 8388608 4G Linux swap
/dev/sda3 10487808 1953523711 1943035904 926.5G Linux root (x86-64)

}}

Questo particolare disco è stato configurato per ospitare 7 filesystem Linux (ciascuno con una corrispondente partizione elencata come "Linux") e una partizione di swap (indicata con "Linux swap").

### Rimuovere tutte le partizioni con fdisk

Pressing the `g` key will instantly remove all existing disk partitions and create a new GPT disklabel:

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 3E56EE74-0571-462B-A992-9872E3855D75).

```

Prima rimuovere tutte le partizioni esistenti dal disco. Digitare `d` per eliminare una partizione. Per esempio, per eliminare un'esistente /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

La partizione è ora programmata per l'eliminazione. Non sarà più mostrata quando si richiede l'elenco delle partizioni ( `p`), comunque non sarà effettivamente eliminata finché i cambiamenti non saranno salvati. Ciò permette agli utenti di annullare l'operazione se è stato commesso qualche errore - in tal caso, digitare subito `q` e premere `Enter` così la partizione non sarà eliminata.

Digitare nuovamente `p` per visualizzare un elenco delle partizioni e premere `d` seguito dal numero della partizione da eliminare. Alla fine, la tabella delle partizioni sarà vuota:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.0 GB, 30005821440 bytes
240 heads, 63 sectors/track, 3876 cylinders
Units = cylinders of 15120 * 512 = 7741440 bytes

Device Boot    Start       End    Blocks   Id  System

```

Ora che la tabella delle partizioni risulta vuota, anche se solo nella memoria, siamo pronti per creare le nuove partizioni.

### Creare la partizione di avvio BIOS

**Nota**

A smaller ESP is possible but not recommended, especially given it may be shared with other OSes.

Per prima cosa si crei una piccola partizione di avvio per il BIOS. Digitare `n` per creare una nuova partizione, quindi `p` per selezionare una partizione primaria, seguito da `1` per selezionare la prima partizione primaria. Quando viene richiesto il settore di inizio, assicurarsi che inizi dal 2048 (necessario per il boot loader) e premere `Enter`. Quando viene richiesto il settore finale, digitare +2M per creare una partizione grande 2 MByte:

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 1
First sector (64-10486533532, default 64): 2048
Last sector, +sectors +size{M,K,G} (4096-10486533532, default 10486533532): +2M

```

Do you want to remove the signature? \[Y\]es/\[N\]o: Y
The signature will be removed by a write command.

}}

Segnare la partizione per gli scopi di UEFI:

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 4
Changed system type of partition 1 to 4 (BIOS boot)

```

### Creare la partizione di swap

Per creare una partizione di swap: digitare `n` per creare una nuova partizione, poi `p` per dire a fdisk di creare una partizione primaria. Digitare `3` per creare una terza partizione primaria, /dev/sda3. Quando viene richiesto il settore di inizio, premere `Enter`. Quando viene richiesto il settore finale, digitare +512M (o qualsiasi altra grandezza sia necessaria per lo spazio di swap) così da creare una partizione grande 512 MB.

### Creare la partizione di avvio

Fatto tutto questo, digitare `t` per impostare il tipo di partizione, `3` per selezionare la partizione appena creata e poi digitare _82_ per impostare il tipo di partizione "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type or alias (type L to list all): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Creare la partizione radice

Infine, per creare la partizione radice (root), digitare `n` per creare una nuova partizione, quindi `p` per indicare a fdisk di creare una partizione primaria. Poi digitare `4` per creare la quarta partizione primaria, /dev/sda4. Quando viene richiesto il settore di inizio, premere `Enter`. Quando viene richiesto il settore finale, premere `Enter` per creare una partizione che occupi il rimanente spazio su disco. Dopo aver completato questi passaggi, digitando `p` si dovrebbe vedere una tabella delle partizioni simile a questa:

`Command (m for help):` `n`

```
Partition number (3-128, default 3): 3
First sector (10487808-1953525134, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525134, default 1953523711):
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Created a new partition 3 of type 'Linux filesystem' and of size 926.5 GiB..

```

**Nota**

Setting the root partition's type to "Linux root (x86-64)" is not required and the system will function normally if it is set to the "Linux filesystem" type. This filesystem type is only necessary for cases where a bootloader that supports it (i.e. systemd-boot) is used and a fstab file is not wanted.

After creating the root partition, press `t` to set the partition type, `3` to select the partition just created, and then type in _23_ to set the partition type to "Linux Root (x86-64)".

`Command(m for help):` `t`

```
Partition number (1-3, default 3): 3
Partition type or alias (type L to list all): 23
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Changed type of partition 'Linux filesystem' to 'Linux root (x86-64)'

```

After completing these steps, pressing `p` should display a partition table that looks similar to the following:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.0 GB, 30005821440 bytes
240 heads, 63 sectors/track, 3876 cylinders
Units = cylinders of 15120 * 512 = 7741440 bytes

   Device Boot    Start       End    Blocks   Id  System
/dev/sda1             1         3      5198+  ef  EFI (FAT-12/16/32)
/dev/sda2   *         3        14    105808+  83  Linux
/dev/sda3            15        81    506520   82  Linux swap
/dev/sda4            82      3876  28690200   83  Linux

```

Device Start End Sectors Size Type
/dev/sda1 2048 2099199 2097152 1G EFI System
/dev/sda2 2099200 10487807 8388608 4G Linux swap
/dev/sda3 10487808 1953523711 1943035904 926.5G Linux root (x86-64)

Filesystem/RAID signature on partition 1 will be wiped.

}}

### Salvare lo schema delle partizioni

Per salvare la configurazione delle partizioni e uscire da fdisk, premere `w`.

`Command (m for help):` `w`

Ora che le partizioni sono state create, si deve procedere alla creazione di un filesystem su ciascuna di esse.

## Partitioning the disk with MBR for BIOS / legacy boot

The following table provides a recommended partition layout for a trivial MBR DOS / legacy BIOS boot installation. Additional partitions can be added according to personal preference or system design goals.

Device path (sysfs)
Mount point
File system
DPS UUID (PARTUUID)
Description
/dev/sda1/bootext4
N/A
MBR DOS / legacy BIOS boot partition details.
/dev/sda2N/A. Swap is not mounted to the filesystem like a device file.0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Swap partition details.
/dev/sda3/xfs
44479540-f297-41b2-9af7-d131d5f0458a
Root partition details.

Change the partition layout according to personal preference.

### Viewing the current partition layout

Fire up fdisk against the disk (in our example, we use /dev/sda):

`root #` `fdisk /dev/sda`

Use the `p` key to display the disk's current partition configuration:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

This particular disk was until now configured to house two Linux filesystems (each with a corresponding partition listed as "Linux") as well as a swap partition (listed as "Linux swap"), using a GPT table.

### Creating a new disklabel / removing all partitions

Pressing `o` will instantly remove all existing disk partitions and create a new MBR disklabel (also named DOS disklabel):

`Command (m for help):` `o`

```
Created a new DOS disklabel with disk identifier 0xf163b576.
The device contains 'gpt' signature and it will be removed by a write command. See fdisk(8) man page and --wipe option for more details.

```

Alternatively, to keep an existing DOS disklabel (see the output of `p` above), consider removing the existing partitions one by one from the disk. Press `d` to delete a partition. For instance, to delete an existing /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

The partition has now been scheduled for deletion. It will no longer show up when printing the list of partitions ( `p`, but it will not be erased until the changes have been saved. This allows users to abort the operation if a mistake was made - in that case, type `q` immediately and hit `Enter` and the partition will not be deleted.

Repeatedly press `p` to print out a partition listing and then press `d` and the number of the partition to delete it. Eventually, the partition table will be empty:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

```

The disk is now ready to create new partitions.

### Creating the boot partition

First, create a small partition which will be mounted as /boot. Press `n` to create a new partition, followed by `p` for a _primary_ partition and `1` to select the first primary partition. When prompted for the first sector, make sure it starts from 2048 (which may be needed for the boot loader) and press `Enter`. When prompted for the last sector, type +1G to create a partition 1 GB in size:

`Command (m for help):` `n`

```
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-1953525167, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525167, default 1953525167): +1G
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Created a new partition 1 of type 'Linux' and of size 1 GiB.

```

Mark the partition as bootable by pressing the `a` key and pressing `Enter`:

`Command (m for help):` `a`

```
Selected partition 1
The bootable flag on partition 1 is enabled now.

```

Note: if more than one partition is available on the disk, then the partition to be flagged as bootable will have to be selected.

### Creating the swap partition

Next, to create the swap partition, press `n` to create a new partition, then `p`, then type `2` to create the second primary partition, /dev/sda2. When prompted for the first sector, press `Enter`. When prompted for the last sector, type +4G (or any other size needed for the swap space) to create a partition 4GB in size.

`Command (m for help):` `n`

```
Partition type
   p   primary (1 primary, 0 extended, 3 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (2-4, default 2): 2
First sector (2099200-1953525167, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525167, default 1953525167): +4G

Created a new partition 2 of type 'Linux' and of size 4 GiB.

```

After all this is done, press `t` to set the partition type, `2` to select the partition just created and then type in _82_ to set the partition type to "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Hex code (type L to list all codes): 82
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Changed type of partition 'Linux' to 'Linux swap / Solaris'.

```

### Creating the root partition

Finally, to create the root partition, press `n` to create a new partition. Then press `p` and `3` to create the third primary partition, /dev/sda3. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to create a partition that takes up the rest of the remaining space on the disk:

`Command (m for help):` `n`

```
Partition type
   p   primary (2 primary, 0 extended, 2 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (3,4, default 3): 3
First sector (10487808-1953525167, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525167, default 1953525167):
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Created a new partition 3 of type 'Linux' and of size 926.5 GiB.

```

After completing these steps, pressing `p` should display a partition table that looks similar to this:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

### Saving the partition layout

Press `w` to write the partition layout and exit fdisk:

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

Now it is time to put filesystems on the partitions.

## Creare i file system

**Attenzione**

Usando un disco SSD o NVME, ha molto senso controllare eventuali aggiornamenti del firmware. Alcuni SSD Intel in particolare (600p e 6000p) necessitano dell'aggiornamento del firmware per [possibile corruzione dei dati](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) indotta dallo schema di I/O utilizzato da XFS. Il problema è a livello di firmware e non un errore nel filesystem XFS. Il programma di utilità smartctl può aiutare nella verifica del modello del dispositivo e della versione di firmware.

### Introduzione

Una volta che sono state create le partizioni, è ora di applicarci un filesystem. Nella sezione successiva vengono descritti i file system compatibili con Linux. I lettori che sanno già quale filesystem usare possono continuare con [applicare un filesystem ad una partizione](/wiki/Handbook:X86/Installation/Disks/it#Applicare_un_filesystem_ad_una_partizione "Handbook:X86/Installation/Disks/it"). Gli altri dovrebbero continuare a leggere per comprendere meglio i filesystem disponibili...

### Filesystem

Linux è compatibile con parecchie dozzine di filesystem, sebbene per molti di questi abbia senso impiegarli solo per scopi specifici. Solo alcuni filesystem potrebbero risultare stabili sull'architettura x86 - si consiglia di documentarsi sui filesystem e lo stato della loro compatibilità prima di selezionarne uno più sperimentale per partizioni importanti. **XFS è il filesystem consigliato per tutti gli scopi, su tutte le piattaforme.** Quello sotto è un elenco non esaustivo:

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

Per esempio, per avere la partizione root (/dev/sda3) con xfs come nello schema delle partizioni d'esempio, si dovrebbero usare i seguenti comandi:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda3`

#### Filesystem della partizione di boot su vecchi BIOS

I sistemi che si avviano su vecchi BIOS con una tabella delle partizioni MBR/DOS possono usare qualsiasi formato del filesystem supportato dal programma di avvio (bootloader).

Per esempio, per formattare con XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda1`

#### Partizioni ext4 piccole

Quando si usa il filesystem ext4 su una piccola partizione (minore di 8GB), il filesystem dovrebbe essere creato con le opzioni appropriate per riservare abbastanza inode. Ciò può essere specificato usando l'opzione `-T small`:

`root #` `mkfs.ext4 -T small /dev/<dispositivo>`

Facendo ciò si quadruplicherà il numero di inode per un certo filesystem poiché i suoi "bytes-per-inode" si riducono da uno ogni 16kB a uno ogni 4kB.

### Attivare la partizione di swap

mkswap è il comando che viene usato per inizializzare la partizione di swap:

`root #` `mkswap /dev/sda2`

**Nota**

Le installazioni che sono state iniziate in precedenza, ma non hanno terminato tutto il processo possono riprendere da questo punto del manuale. Si usi questo link come permalink [Iniziare da qui per installazioni da riprendere](/wiki/Handbook:X86/Installation/Disks/it#Resumed_installations_start_here "Handbook:X86/Installation/Disks/it").

Per attivare la partizione di swap, usare swapon:

`root #` `swapon /dev/sda2`

Questo passaggio di 'attivazione' è necessario solo perchè la partizione di swap è creata nuovamente all' interno dell'ambiente live. Una volta riavviato il sistema, a patto che la partizione di swap sia stata ben definita all'interno di fstab o altro meccanismo di montaggio delle partizioni, lo spazio di swap si attiverà automaticamente.

## Montare la partizione di root

In alcuni ambienti live potrebbero mancare i punti di montaggio (mount) suggeriti per la partizione di root di Gentoo (/mnt/gentoo), o punti di montaggio per partizioni aggiuntive create nella sezione di partizionamento:

`root #` `mkdir --parents /mnt/gentoo`

Continuare a creare i punti di montaggio aggiuntivi necessari per ogni partizione (personalizzata) creata durante i passaggi precedenti, usando il comando mkdir.

Con i punti di montaggio creati, è tempo di rendere le partizioni accessibili tramite il comando mount.

Montare la partizione di root:

`root #` `mount /dev/sda3 /mnt/gentoo`

Continuare montando le partizioni aggiuntive (personalizzate) come necessario usando il comando mount.

**Nota**

Se è necessario che /tmp/ risieda su una partizione separata, assicurarsi di cambiare i suoi permessi dopo averla montata:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Ciò è valido anche per /var/tmp.

Più avanti nel manuale, il filesystem proc (un'interfaccia virtuale fornita dal kernel) ed altri pseudo-filesystem del kernel verranno montati. Però prima deve essere estratto il [file stage di Gentoo](/wiki/Handbook:X86/Installation/Stage/it "Handbook:X86/Installation/Stage/it").

[← Configurare la rete](/wiki/Handbook:X86/Installation/Networking/it "Previous part") [Home](/wiki/Handbook:X86 "Handbook:X86") [Installare il file stage →](/wiki/Handbook:X86/Installation/Stage/it "Next part")