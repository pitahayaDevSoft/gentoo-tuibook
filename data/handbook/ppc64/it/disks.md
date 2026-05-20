# Disks

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Disks/de "Handbuch:PPC64/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Disks/es "Manual de Gentoo: PPC64/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Disks/fr "Handbook:PPC64/Installation/Disks/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:PPC64/Installation/Disks/hu "Handbook:PPC64/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Disks/pl "Handbook:PPC64/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Disks/pt-br "Handbook:PPC64/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Disks/ru "Handbook:PPC64/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Disks/ta "கையேடு:PPC64/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Disks/zh-cn "手册：PPC64/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Disks/ja "ハンドブック:PPC64/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Disks/ko "Handbook:PPC64/Installation/Disks/ko (100% translated)")

[Manuale PPC64](/wiki/Handbook:PPC64/it "Handbook:PPC64/it")[Installazione](/wiki/Handbook:PPC64/Full/Installation/it "Handbook:PPC64/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:PPC64/Installation/About/it "Handbook:PPC64/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:PPC64/Installation/Media/it "Handbook:PPC64/Installation/Media/it")[Configurare la rete](/wiki/Handbook:PPC64/Installation/Networking/it "Handbook:PPC64/Installation/Networking/it")Preparare i dischi[Il file stage](/wiki/Handbook:PPC64/Installation/Stage/it "Handbook:PPC64/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:PPC64/Installation/Base/it "Handbook:PPC64/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:PPC64/Installation/Kernel/it "Handbook:PPC64/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:PPC64/Installation/System/it "Handbook:PPC64/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:PPC64/Installation/Tools/it "Handbook:PPC64/Installation/Tools/it")[Impostare programma d'avvio](/wiki/Handbook:PPC64/Installation/Bootloader/it "Handbook:PPC64/Installation/Bootloader/it")[Concludere](/wiki/Handbook:PPC64/Installation/Finalizing/it "Handbook:PPC64/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:PPC64/Full/Working/it "Handbook:PPC64/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:PPC64/Working/Portage/it "Handbook:PPC64/Working/Portage/it")[Opzioni USE](/wiki/Handbook:PPC64/Working/USE/it "Handbook:PPC64/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:PPC64/Working/Features/it "Handbook:PPC64/Working/Features/it")[Sistema script di init](/wiki/Handbook:PPC64/Working/Initscripts/it "Handbook:PPC64/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:PPC64/Working/EnvVar/it "Handbook:PPC64/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:PPC64/Full/Portage/it "Handbook:PPC64/Full/Portage/it")[File e cartelle](/wiki/Handbook:PPC64/Portage/Files/it "Handbook:PPC64/Portage/Files/it")[Variabili](/wiki/Handbook:PPC64/Portage/Variables/it "Handbook:PPC64/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:PPC64/Portage/Branches/it "Handbook:PPC64/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:PPC64/Portage/Tools/it "Handbook:PPC64/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:PPC64/Portage/CustomTree/it "Handbook:PPC64/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:PPC64/Portage/Advanced/it "Handbook:PPC64/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:PPC64/Full/Networking/it "Handbook:PPC64/Full/Networking/it")[Come iniziare](/wiki/Handbook:PPC64/Networking/Introduction/it "Handbook:PPC64/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:PPC64/Networking/Advanced/it "Handbook:PPC64/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:PPC64/Networking/Modular/it "Handbook:PPC64/Networking/Modular/it")[Wireless](/wiki/Handbook:PPC64/Networking/Wireless/it "Handbook:PPC64/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:PPC64/Networking/Extending/it "Handbook:PPC64/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:PPC64/Networking/Dynamic/it "Handbook:PPC64/Networking/Dynamic/it")

## Contents

- [1Introduzione ai dispositivi a blocchi](#Introduzione_ai_dispositivi_a_blocchi)
  - [1.1Dispositivi a blocchi](#Dispositivi_a_blocchi)
  - [1.2Partizioni e porzioni](#Partizioni_e_porzioni)
- [2Progettare uno schema delle partizioni](#Progettare_uno_schema_delle_partizioni)
  - [2.1Quante partizioni e quanto grandi?](#Quante_partizioni_e_quanto_grandi.3F)
  - [2.2Riguardo lo spazio di swap?](#Riguardo_lo_spazio_di_swap.3F)
- [3Predefinito: Usare mac-fdisk](#Predefinito:_Usare_mac-fdisk)
- [4In alternativa: Usare fdisk](#In_alternativa:_Usare_fdisk)
  - [4.1Vedere lo schema di partizione corrente](#Vedere_lo_schema_di_partizione_corrente)
  - [4.2Rimuovere tutte le partizioni](#Rimuovere_tutte_le_partizioni)
  - [4.3Creare la partizione di avvio PPC PReP](#Creare_la_partizione_di_avvio_PPC_PReP)
  - [4.4Creare la partizione di swap](#Creare_la_partizione_di_swap)
  - [4.5Creare la partizione di root](#Creare_la_partizione_di_root)
  - [4.6Salvare lo schema delle partizioni](#Salvare_lo_schema_delle_partizioni)
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

### Partizioni e porzioni

Anche se teoricamente è possibile usare l'inetro disco per ospitare un sistema Linux, in realtà ciò non viene quasi mai fatto. Invece, dispositivi a blocchi di un intero disco vengono suddivisi in dispositivi a blocchi più piccoli e più partici. Nella maggior parte dei sistemi, vengono chiamate partizioni. Altre architetture usano una tecnica simile, chiamandole _fette_.

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

## Predefinito: Usare mac-fdisk

**Importante**

Queste istruzioni sono per il sistema Apple G5.

Partizione
Descrizione
/dev/sda1Apple partition map, creata automaticamente quando il disco viene formattato con una tabella delle partizioni "mac".
/dev/sda2Partizione di avvio New World
/dev/sda3Partizione di swap
/dev/sda4Partizione di root

Avviare mac-fdisk:

`root #` `mac-fdisk /dev/sda`

Primo, cancellare le partizioni che sono state ripulite in precedenza per fare spazio alle partizioni Linux. Usare il tasto `d` in mac-fdisk per cancellare tali partizioni. Verrà richiesto il numero della partizione da cancellare.

Secondo, creare una partizione Apple\_Bootstrap premendo il tasto `b`. Verrà richiesto da quale blocco iniziare. Inserire il numero della prima partizione libera, seguito dall'immissione di `p`. Ad esempio questo è _2p_.

**Nota**

Questa non è una partizione di "boot". Non è usata affatto da Linux; non è necessario inserirvi alcun filesystem e non dovrebbe essere montata(mounted). Gli utenti PPC _non_ hanno bisogno di una partizione in più per /boot.

Ora creare una partizione di swap premendo il tasto `c`. Nuovamente mac-fdisk richiederà da quale blocco iniziare. Poichè è stato usato 2 prima per creare la partizione Apple\_Bootstrap, inserire _3p_. Quando verrà richiesta la dimensione inserire **512M** (o qualsiasi dimensione si necessiti). Quando verrà richiesto il nome, inserire _swap_ (obbligatorio).

Per creare la partizione di root, premere `c`, seguito da _4p_ per selezionare il blocco da cui iniziare la partizione di root. Quando verrà richiesta la dimensione, immettere _4p_ di nuovo. mac-fdisk interpreterà ciò come "Usare tutto lo spazio disponibile". Quando verrà richiesto il nome, inserire _root_ (obbligatorio).

Per completare, scrivere le partizioni sul disco usando `w` e poi `q` per uscire da mac-fdisk.

**Nota**

Per accertarsi che tutto sia a posto, eseguire mac-fdisk ancora una volta per verificare che tutte le partizioni siano presenti. Se una partizione è assente, o mancano alcune delle modifiche fatte, allora reinizializzare le partizioni premendo `i` in mac-fdisk. Si badi bene che ciò ricreera la mappa delle partizioni e quindi rimuoverà tutte le partizioni.

## In alternativa: Usare fdisk

**Importante**

Le istruzioni seguenti sono per sistemi IBM pSeries, iSeries, e OpenPower.

**Nota**

Quando si programma di usare un array di dischi RAID per l'installazione di Gentoo su hardware basato su POWER5, prima eseguire iprconfig per formattare i dischi nel formato Advanced Function e creare l'array del disco. Fare l'emerge di [sys-fs/iprutils](https://packages.gentoo.org/packages/sys-fs/iprutils) dopo aver completato l'installazione.

Se il sistema ha un adattatore SCSI basato su ipr, avviare il programma di utilità ipr ora.

`root #` `/etc/init.d/iprinit start`

Le parti seguenti spiegano come creare lo schema dell'esempio di partizione descritto in precedenza, cioè:

Partizione
Descrizione
/dev/sda1Partizione di avvio PPC PReP
/dev/sda2Partizione di swap
/dev/sda3Partizione di root

Cambiare o modificare lo schema delle partizioni secondo le preference personali.

### Vedere lo schema di partizione corrente

[fdisk](/wiki/Fdisk "Fdisk") è uno strumento popolare e potente per suddividere il disco in partizioni. Avviare fdisk sul disco corrente (nell'esempio si userà /dev/sda):

`root #` `fdisk /dev/sda`

```
Command (m for help)

```

Se c'è ancora uno schema di partizione AIX sul sistema, allora apparirà il seguente messaggio di errore:

`root #` `fdisk /dev/sda`

```
  There is a valid AIX label on this disk.
  Unfortunately Linux cannot handle these
  disks at the moment.  Nevertheless some
  advice:
  1. fdisk will destroy its contents on write.
  2. Be sure that this disk is NOT a still vital
     part of a volume group. (Otherwise you may
     erase the other disks as well, if unmirrored.)
  3. Before deleting this physical volume be sure
     to remove the disk logically from your AIX
     machine.  (Otherwise you become an AIXpert).

```

Non c'è da preoccuparsi, una nuova tabella delle partizioni DOS vuota può essere creata premendo `o`.

**Attenzione**

Ciò distruggerà qualsiasi versione di AIX installata!

Premere `p` per mostrare la configurazione corrente delle partizioni del disco:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          12       53266+  83  Linux
/dev/sda2              13         233      981571+  82  Linux swap
/dev/sda3             234         674     1958701+  83  Linux
/dev/sda4             675        6761    27035410+   5  Extended
/dev/sda5             675        2874     9771268+  83  Linux
/dev/sda6            2875        2919      199836   83  Linux
/dev/sda7            2920        3008      395262   83  Linux
/dev/sda8            3009        6761    16668918   83  Linux

```

Questo disco in particolare è configurato per ospitare sei filesystem Linux (ognuno con la corrispondente partizione elencata come "Linux") oltre a una partizione di swap (elencata come "Linux swap").

### Rimuovere tutte le partizioni

Prima rimuovere tutte le partizioni esistenti dal disco. Premere `d` per cancellare una partizione. Per esempio, per cancellare un esistente /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

La partizione è stata programmata per la cancellazione. Non sarà più visualizzata digitando `p`, ma non sarà cancellata fino al salvataggio delle modifiche. Se si è commesso un errore e si ha bisogno di interrompere la sessione, allora digitare `q` immediatamente e premere `Enter` e nessuna delle partizioni sarà cancellata o modificata.

Ora, supponendo che si abbia la necessità di cancellare proprio tutte le partizioni, digitare ripetutamente `p` per stampare l'elenco delle partizioni e poi digitare `d` e il numero della partizione per cancellarla. Prima o poi, la tabella delle partizioni non mostrerà più partizioni:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

Device Boot    Start       End    Blocks   Id  System

```

Ora che la tabella delle partizioni in memoria è vuota, si creino le partizioni. Si userà uno schema di partizionamento predefinito come discusso in precedenza. Naturalmente, non si seguano queste istruzioni alla lettera ma si adattino alle preferenze personali.

### Creare la partizione di avvio PPC PReP

Prima creare una piccola partizione di avvio PReP. Digitare `n` per creare una nuova partizione, poi `p` per selezionare una partizione primaria, seguito da `1` per selezionare la prima partizione primaria. Quando viene richiesto il primo cilindro, premere `Enter`. Quando
viene richiesto l'ultimo cilindro, digitare +7M per creare una partizione della dimensione di 7MB. Dopodichè, digitare `t` per impostare il tipo di partizione, `1` per selezionare la partizione appena creata e poi inserire _41_ per impostare il tipo di partizione a "PPC PReP Boot". Infine, contrassegnare la partizione PRep come avviabile.

**Nota**

La partizione PReP deve essere minore di 8 MB!

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System

```

`Command (m for help):` `n`

```
Command action
      e   extended
      p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-6761, default 1):
Using default value 1
Last cylinder or +size or +sizeM or +sizeK (1-6761, default
6761): +8M

```

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 41
Changed system type of partition 1 to 41 (PPC PReP Boot)

```

`Command (m for help):` `a`

```
Partition number (1-4): 1
Command (m for help):

```

Ora, quando si osserva nuovamente la tabella delle partizioni (attraverso `p`), dovrebbe essere mostrata la seguente informazione:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1  *            1           3       13293   41  PPC PReP Boot

```

### Creare la partizione di swap

Adesso creare la partizione di swap. Per fare ciò, digitare `n` per creare una nuova partizione, poi `p` per dire a fdisk di creare una partizione primaria. Quindi digitare `2` per creare la seconda partizione primaria, /dev/sda2 in questo caso. Quando viene richiesto il primo cilindro, premere `Enter`. Quando viene richiesto l'ultimo cilindro, digitare +512M per creare una partizione della dimensione di 512MB. Dopodichè, digitare `t` per impostare il tipo di partizione, `2` per selezionare la partizione appena creata e poi inserire _82_ per impostare il tipo di partizione a "Linux Swap". Dopo aver completato questi passaggi, digitando `p` dovrebbe apparire una tabella delle partizioni simile a questa:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1           3       13293   41  PPC PReP Boot
/dev/sda2               4         117      506331   82  Linux swap

```

### Creare la partizione di root

Infine, creare la partizione di root. Per fare ciò, digitare `n` per creare una nuova partizione, poi `p` per dire a fdisk di creare una partizione primaria. Quindi digitare `3` per creare la terza partizione primaria, /dev/sda3 in questo caso. Quando viene richiesto il primo cilindro premere `Enter`. Quando viene richiesto l'ultimo cilindro, premere `Enter` per creare una partizione che occupa il resto dello spazio rimanente sul disco. Dopo aver completato questi passaggi, digitando `p` dovrebbe apparire una tabella delle partizioni simile a questa:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1           3       13293   41  PPC PReP Boot
/dev/sda2               4         117      506331   82  Linux swap
/dev/sda3             118        6761    29509326   83  Linux

```

### Salvare lo schema delle partizioni

Per salvare lo schema delle partizioni e uscire da fdisk, digitare `w`.

`Command (m for help):` `w`

## Creare i file system

**Attenzione**

Usando un disco SSD o NVME, ha molto senso controllare eventuali aggiornamenti del firmware. Alcuni SSD Intel in particolare (600p e 6000p) necessitano dell'aggiornamento del firmware per [possibile corruzione dei dati](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) indotta dallo schema di I/O utilizzato da XFS. Il problema è a livello di firmware e non un errore nel filesystem XFS. Il programma di utilità smartctl può aiutare nella verifica del modello del dispositivo e della versione di firmware.

### Introduzione

Una volta che sono state create le partizioni, è ora di applicarci un filesystem. Nella sezione successiva vengono descritti i file system compatibili con Linux. I lettori che sanno già quale filesystem usare possono continuare con [applicare un filesystem ad una partizione](/wiki/Handbook:PPC64/Installation/Disks/it#Applicare_un_filesystem_ad_una_partizione "Handbook:PPC64/Installation/Disks/it"). Gli altri dovrebbero continuare a leggere per comprendere meglio i filesystem disponibili...

### Filesystem

Linux è compatibile con parecchie dozzine di filesystem, sebbene per molti di questi abbia senso impiegarli solo per scopi specifici. Solo alcuni filesystem potrebbero risultare stabili sull'architettura ppc64 - si consiglia di documentarsi sui filesystem e lo stato della loro compatibilità prima di selezionarne uno più sperimentale per partizioni importanti. **XFS è il filesystem consigliato per tutti gli scopi, su tutte le piattaforme.** Quello sotto è un elenco non esaustivo:

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

Le installazioni che sono state iniziate in precedenza, ma non hanno terminato tutto il processo possono riprendere da questo punto del manuale. Si usi questo link come permalink [Iniziare da qui per installazioni da riprendere](/wiki/Handbook:PPC64/Installation/Disks/it#Resumed_installations_start_here "Handbook:PPC64/Installation/Disks/it").

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

Più avanti nel manuale, il filesystem proc (un'interfaccia virtuale fornita dal kernel) ed altri pseudo-filesystem del kernel verranno montati. Però prima deve essere estratto il [file stage di Gentoo](/wiki/Handbook:PPC64/Installation/Stage/it "Handbook:PPC64/Installation/Stage/it").

[← Configurare la rete](/wiki/Handbook:PPC64/Installation/Networking/it "Previous part") [Home](/wiki/Handbook:PPC64 "Handbook:PPC64") [Installare il file stage →](/wiki/Handbook:PPC64/Installation/Stage/it "Next part")