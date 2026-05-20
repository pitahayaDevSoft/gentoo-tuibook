# Disks

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Disks/de "Handbuch:PPC/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Disks "Handbook:PPC/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Disks/es "Manual de Gentoo: PPC/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Disks/fr "Handbook:PPC/Installation/Disks/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:PPC/Installation/Disks/hu "Handbook:PPC/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Disks/pl "Handbook:PPC/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Disks/pt-br "Handbook:PPC/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Disks/ru "Handbook:PPC/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Disks/ta "கையேடு:PPC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Disks/zh-cn "手册：PPC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Disks/ja "ハンドブック:PPC/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Disks/ko "Handbook:PPC/Installation/Disks/ko (100% translated)")

[Manuale PPC](/wiki/Handbook:PPC/it "Handbook:PPC/it")[Installazione](/wiki/Handbook:PPC/Full/Installation/it "Handbook:PPC/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:PPC/Installation/About/it "Handbook:PPC/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:PPC/Installation/Media/it "Handbook:PPC/Installation/Media/it")[Configurare la rete](/wiki/Handbook:PPC/Installation/Networking/it "Handbook:PPC/Installation/Networking/it")Preparare i dischi[Il file stage](/wiki/Handbook:PPC/Installation/Stage/it "Handbook:PPC/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:PPC/Installation/Base/it "Handbook:PPC/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:PPC/Installation/Kernel/it "Handbook:PPC/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:PPC/Installation/System/it "Handbook:PPC/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:PPC/Installation/Tools/it "Handbook:PPC/Installation/Tools/it")[Impostare programma d'avvio](/wiki/Handbook:PPC/Installation/Bootloader/it "Handbook:PPC/Installation/Bootloader/it")[Concludere](/wiki/Handbook:PPC/Installation/Finalizing/it "Handbook:PPC/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:PPC/Full/Working/it "Handbook:PPC/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:PPC/Working/Portage/it "Handbook:PPC/Working/Portage/it")[Opzioni USE](/wiki/Handbook:PPC/Working/USE/it "Handbook:PPC/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:PPC/Working/Features/it "Handbook:PPC/Working/Features/it")[Sistema script di init](/wiki/Handbook:PPC/Working/Initscripts/it "Handbook:PPC/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:PPC/Working/EnvVar/it "Handbook:PPC/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:PPC/Full/Portage/it "Handbook:PPC/Full/Portage/it")[File e cartelle](/wiki/Handbook:PPC/Portage/Files/it "Handbook:PPC/Portage/Files/it")[Variabili](/wiki/Handbook:PPC/Portage/Variables/it "Handbook:PPC/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:PPC/Portage/Branches/it "Handbook:PPC/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:PPC/Portage/Tools/it "Handbook:PPC/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:PPC/Portage/CustomTree/it "Handbook:PPC/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:PPC/Portage/Advanced/it "Handbook:PPC/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:PPC/Full/Networking/it "Handbook:PPC/Full/Networking/it")[Come iniziare](/wiki/Handbook:PPC/Networking/Introduction/it "Handbook:PPC/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:PPC/Networking/Advanced/it "Handbook:PPC/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:PPC/Networking/Modular/it "Handbook:PPC/Networking/Modular/it")[Wireless](/wiki/Handbook:PPC/Networking/Wireless/it "Handbook:PPC/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:PPC/Networking/Extending/it "Handbook:PPC/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:PPC/Networking/Dynamic/it "Handbook:PPC/Networking/Dynamic/it")

## Contents

- [1Introduzione ai dispositivi a blocchi](#Introduzione_ai_dispositivi_a_blocchi)
  - [1.1Dispositivi a blocchi](#Dispositivi_a_blocchi)
  - [1.2Partizioni](#Partizioni)
- [2Progettare uno schema delle partizioni](#Progettare_uno_schema_delle_partizioni)
  - [2.1Quante partizioni e quanto grandi?](#Quante_partizioni_e_quanto_grandi.3F)
  - [2.2Riguardo lo spazio di swap?](#Riguardo_lo_spazio_di_swap.3F)
  - [2.3Apple NewWorld](#Apple_NewWorld)
  - [2.4Apple OldWorld](#Apple_OldWorld)
  - [2.5Pegasos](#Pegasos)
  - [2.6IBM PReP (RS/6000)](#IBM_PReP_.28RS.2F6000.29)
- [3Usare mac-fdisk (per Apple)](#Usare_mac-fdisk_.28per_Apple.29)
- [4Usare parted (per Pegasos e RS/6000)](#Usare_parted_.28per_Pegasos_e_RS.2F6000.29)
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

### Partizioni

Sebbene teoricamente sia possibile usare un intero disco per ospitare un sistema Linux, in realtà ciò non viene quasi mai fatto. Invece interi dischi di dispositivi a blocchi vengono suddivisi in dispositivi a blocchi minori e più gestibili. Sulla maggior parte dei sistemi, questi vengono chiamati partizioni.

**Nota**

Nel resto delle istruzioni di installazione, si userà lo schema di partizione dell' esempio Pegasos. Adattare alle preferenze personali.

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

### Apple NewWorld

Le macchine Apple NewWorld sono piuttosto semplici da configurare. La prima partizione è sempre una Apple Partition Map (APM). Questa partizione tiene traccia dello schema del disco. Non è possibile eliminare questa partizione. La partizione successiva dovrebbe sempre essere una partizione di avvio. Questa partizione contiene un piccolo (800KiB) filesystem HFS che contiene una copia del programma di avvio Yaboot e del suo file di configurazione. Questa partizione non è come quella della partizione /boot trovata su altre architetture. Dopo la partizione di avvio, vengono posizionati i consueti filesystem di Linux, secondo lo schema sotto. La partizione di swap è un luogo di memorizzazione temporaneo per quando il sistema esaurisce la memoria fisica. La partizione di root conterrà il filesystem sul quale sarà installato Gentoo. Per il duplice avvio, la partizione OSX può andare ovunque dopo la partizione di avvio per assicurare che yaboot si avvii prima.

**Nota**

Potrebbero esserci partizioni "Disk Driver" sul disco come Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKite Apple\_Patches. Queste vengono usate per avviare MacOS, perciò se non si ha bisogno di ciò, possono essere rimosse inizializzando il disco con l'opzione `i` di mac-fdisk. Attenzione, ciò cancellerà il disco completamente! Se in dubbio non rimuoverle.

**Nota**

Se il disco è partizionato con Utilità Disco di Apple, potrebbero esserci spazi di 128 MiB tra le partizioni che Apple riserva per "usi futuri". Posssono essere eliminate tranquillamente.

Partizione
Dimensione
Filesystem
Descrizione
/dev/sda132KiB
Nessuno
Apple Partition Map (APM).
/dev/sda2800KiB
HFS
Apple bootstrap.
/dev/sda3512 MiB
swap
Linux swap (type 0x82).
/dev/sda4Resto del disco
ext4, xfs, etc
Linux root.

### Apple OldWorld

Le macchine Apple OldWorld sono un po' più complicate da configurare. La prima partizione è sempre una Apple Partition Map (APM). Questa partizione tiene traccia dello schema del disco. Non è possibile eliminare questa partizione. Usando BootX, la configurazione sotto suppone che MacOS sia installato su un disco separato. Se non si è in questo caso, ci saranno partizioni aggiuntive per "Apple Disk Drivers" come Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKit, Apple\_Patches e l'installazione di MacOS. Usando Quik è necessario creare una partizione di avvio per contenere il kernel, a differenza degli altri metodi di avvio di Apple. Dopo la partizione di avvio, vengono posizionati i consueti filesystem di Linux, secondo lo schema sotto. La partizione di swap è un luogo di memorizzazione temporaneo per quando il sistema esaurisce la memoria fisica. La partizione di root conterrà il filesystem sul quale sarà installato Gentoo.

**Nota**

Usando una macchina OldWorld, è necessario mantenere MacOS disponibile. Lo schema qui riportato suppone che MacOS sia installato su un disco separato.

Esempio di schema di partizione per una macchina OldWorld:

Partizione
Dimensione
Filesystem
Descrizione
/dev/sda132KiB
Nessuno
Apple Partition Map (APM).
/dev/sda232MiB
ext2
Partizione di avvio di quik (solo quik).
/dev/sda3512MiB
swap
Linux swap (type 0x82).
/dev/sda4Resto del disco
ext4, xfs, etc
Linux root.

### Pegasos

Lo schema di partizione Pegasos è piuttosto semplice comparato con gli schemi Apple. La prima partizione è una partizione di avvio, che contiene i kernel da avviare insieme ad uno script Open Firmware che mostra un menù all'avvio. Dopo la partizione di avvio, vengono posizionati i consueti filesystem di Linux, secondo lo schema sotto. La partizione di swap è un luogo di memorizzazione temporaneo per quando il sistema esaurisce la memoria fisica. La partizione di root conterrà il filesystem sul quale sarà installato Gentoo.

Esempio di schema di partizione per sistemi Pegasos:

Partizione
Dimensione
Filesystem
Descrizione
/dev/sda132MiB
affs1 o ext2
Partizione di avvio.
/dev/sda2512MiB
swap
Linux swap (type 0x82).
/dev/sda3Resto del disco
ext4, xfs, etc
Linux root.

### IBM PReP (RS/6000)

L'IBM PowerPC Reference Platform (PReP) necessita di una piccola partizione di avvio PReP sulla prima partizione del disco, seguita dalle partizioni di swap e root.

Esempio di schema di partizione per l'IBM PReP:

Partizione
Dimensione
Filesystem
Descrizione
/dev/sda1800KiB
Nessuno
Partizione di avvio PReP (type 0x41).
/dev/sda2512MiB
swap
Linux swap (type 0x82).
/dev/sda3Resto del disco
ext4, xfs, etc
Linux root (type 0x83).

**Attenzione**

parted è capace di ridimensionare le partizioni incluse le HFS+. Purtroppo potrebbero presentarsi dei problemi ridimensionando filesystem HFS+ journaled, quindi, per i migliori risultati, disabilitare il journaling in MacOS X prima di ridimensionare. Si ricordi che ogni operazione di ridimensionamento è pericolosa, quindi si provi a proprio rischio! Ci si accerti di avere sempre una copia di tutti i dati prima di ridimensionare!

## Usare mac-fdisk (per Apple)

A questo punto, creare le partizioni usando mac-fdisk:

`root #` `mac-fdisk /dev/sda`

Se in precedenza è stato usato Utility Disco di Apple per lasciare spazio per Linux, prima cancellare le partizioni che potrebbero essere state create per fare spazio alla nuova installazione. Usare `d` in mac-fdisk per cancellare quella(e) partizione(i). Verrà richiesto il numero della partizione da eliminare. Di solito la prima partizione su macchine NewWorld (Apple\_partition\_map) non può essere eliminata. Per iniziare con un disco pulito, semplicente inizializzare il disco premendo `i`. Ciò cancellerà completamente il disco, perciò usare con cautela.

Secondo, creare una partizione Apple\_Bootstrap usando `b`. Verrà richiesto da quale blocco partire. Inserire il numero della prima partizione libera, seguito da una `p`. Ad esempio questa è _2p_.

**Nota**

Questa partizione non è una partizione /boot. Non viene usata affatto da Linux; non è necessario applicare un filesystem e non dovrebbe mai essere montata. Gli utenti Apple non hanno bisogno di una partizione in più per /boot.

Ora creare la partizione di swap premendo `c`. Nuovamente mac-fdisk richiederà da quale blocco inizia la partizione. poichè abbiamo usato 2 prima per creare la partizione Apple\_Bootstrap, ora inserire _3p_. Quando richiesto per la dimensione, inserire 512M (o qualsiasi dimensione si necessiti - è consigliato un minimo di 512MiB, ma 2 volte la memoria fisica è la dimensione in genere accettata). Quando richiesto per un nome, inserire _swap_.

Per creare la partizione di root, inserire `c`, seguito da _4p_ per selezionare da quale blocco la partizione di root dovrebbe iniziare. Quando richiesto per la dimensione inserire nuovamente _4p_. mac-fdisk interpreterà ciò come "Usa tutto lo spazio disponibile". Quando è richiesto il nome, inserire _root_.

Per concludere, scrivere le partizioni sul disco usando `w` e poi `q` per uscire da mac-fdisk.

**Nota**

Per essere sicuri che tutto sia a posto, eseguire mac-fdisk -l e controllare se tutte le partizioni ci sono. Se non vengono mostrate tutte le partizioni create in precedenza, o le modifiche effettuate non risultano applicate, reinizializzare le partizioni premendo `i` in mac-fdisk. Si noti che ciò ricreerà la mappa delle partizioni e quindi rimuoverà ogni partizione esistente.

## Usare parted (per Pegasos e RS/6000)

parted, il **part** ition **ed** itor, può ora manipolare le partizioni HFS+ usate da Mac OS e Mac OS X. Con questo strumento è possibile ridimensionare le partizioni Mac e creare spazio per le partizioni Linux. Nonostante ciò, l'esempio sotto descrive solo il partizionamento per le macchine Pegasos.

Per iniziare avviare parted:

`root #` `parted /dev/sda`

Se il disco non è partizionato, eseguire mklabel amiga per creare una nuova etichetta al disco.

E' possibile digitare print in ogni momento in parted per mostrare la tabella delle partizioni corrente. Per interrompere parted, premere `Ctrl` + `c`.

Se accanto a Linux, il sistema è pensato per avere anche MorphOS installato, allora creare un filesystem affs1 all'inizio del disco. 32MB dovrebbero essere più che sufficienti per memorizzare il kernel MorphOS. Con un Pegasos I, o quando Linux userà qualsiasi altro filesystem oltre ext2 o ext3, allora è necessario memorizzare anche il kernel Linux su questa partizione (il Pegasos II può avviare solo da partizioni ext2/ext3 o affs1). Per creare la partizione eseguire `mkpart primary affs1 START END` dove START ed END dovrebbero essere sostituiti con l'intervallo in megabyte (p.es. 0 32) che crea una partizione di 32MB iniziando a 0MB e finendo a 32MB. Creando una partizione ext2 o ext3 invece, sostituire ext2 o ext3 ad affs1 nel comando mkpart.

Creare due partizioni per Linux, un filesystem di root e una partizione di swap. Eseguire `mkpart primary START END` per creare ogni partizione, sostituendo START ed END coi limiti in megabyte desiderati.

In genere si consiglia di creare una partizione di swap che sia due volte più grande del quantitativo di RAM nel computer, ma almeno 512 MB è consigliata. Per creare la partizione di swap, eseguire `mkpart primary linux-swap START END` con START ed END indicanti di nuovo i limiti della partizione.

Quando si è terminato in parted digitare semplicemente `quit`.

## Creare i file system

**Attenzione**

Usando un disco SSD o NVME, ha molto senso controllare eventuali aggiornamenti del firmware. Alcuni SSD Intel in particolare (600p e 6000p) necessitano dell'aggiornamento del firmware per [possibile corruzione dei dati](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) indotta dallo schema di I/O utilizzato da XFS. Il problema è a livello di firmware e non un errore nel filesystem XFS. Il programma di utilità smartctl può aiutare nella verifica del modello del dispositivo e della versione di firmware.

### Introduzione

Una volta che sono state create le partizioni, è ora di applicarci un filesystem. Nella sezione successiva vengono descritti i file system compatibili con Linux. I lettori che sanno già quale filesystem usare possono continuare con [applicare un filesystem ad una partizione](/wiki/Handbook:PPC/Installation/Disks/it#Applicare_un_filesystem_ad_una_partizione "Handbook:PPC/Installation/Disks/it"). Gli altri dovrebbero continuare a leggere per comprendere meglio i filesystem disponibili...

### Filesystem

Linux è compatibile con parecchie dozzine di filesystem, sebbene per molti di questi abbia senso impiegarli solo per scopi specifici. Solo alcuni filesystem potrebbero risultare stabili sull'architettura ppc - si consiglia di documentarsi sui filesystem e lo stato della loro compatibilità prima di selezionarne uno più sperimentale per partizioni importanti. **XFS è il filesystem consigliato per tutti gli scopi, su tutte le piattaforme.** Quello sotto è un elenco non esaustivo:

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

Le installazioni che sono state iniziate in precedenza, ma non hanno terminato tutto il processo possono riprendere da questo punto del manuale. Si usi questo link come permalink [Iniziare da qui per installazioni da riprendere](/wiki/Handbook:PPC/Installation/Disks/it#Resumed_installations_start_here "Handbook:PPC/Installation/Disks/it").

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

Più avanti nel manuale, il filesystem proc (un'interfaccia virtuale fornita dal kernel) ed altri pseudo-filesystem del kernel verranno montati. Però prima deve essere estratto il [file stage di Gentoo](/wiki/Handbook:PPC/Installation/Stage/it "Handbook:PPC/Installation/Stage/it").

[← Configurare la rete](/wiki/Handbook:PPC/Installation/Networking/it "Previous part") [Home](/wiki/Handbook:PPC "Handbook:PPC") [Installare il file stage →](/wiki/Handbook:PPC/Installation/Stage/it "Next part")