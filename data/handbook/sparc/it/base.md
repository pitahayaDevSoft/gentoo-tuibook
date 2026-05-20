# Base

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Base/de "Handbuch:SPARC/Installation/Basis (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Base "Handbook:SPARC/Installation/Base (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Base/es "Manual de Gentoo: SPARC/Instalación/Base (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Base/fr "Handbook:SPARC/Installation/Base/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:SPARC/Installation/Base/hu "Handbook:SPARC/Installation/Base/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Base/pl "Handbook:SPARC/Installation/Base (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Base/pt-br "Handbook:SPARC/Installation/Base/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Base/ru "Handbook:SPARC/Installation/Base (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Base/ta "கையேடு:SPARC/நிறுவல்/அடிப்படை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Base/zh-cn "手册：SPARC/安装/安装基本系统 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Base/ja "ハンドブック:SPARC/インストール/ベース (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Base/ko "Handbook:SPARC/Installation/Base/ko (100% translated)")

[Manuale SPARC](/wiki/Handbook:SPARC/it "Handbook:SPARC/it")[Installazione](/wiki/Handbook:SPARC/Full/Installation/it "Handbook:SPARC/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:SPARC/Installation/About/it "Handbook:SPARC/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:SPARC/Installation/Media/it "Handbook:SPARC/Installation/Media/it")[Configurare la rete](/wiki/Handbook:SPARC/Installation/Networking/it "Handbook:SPARC/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:SPARC/Installation/Disks/it "Handbook:SPARC/Installation/Disks/it")[Il file stage](/wiki/Handbook:SPARC/Installation/Stage/it "Handbook:SPARC/Installation/Stage/it")Installare il sistema base[Configurare il kernel](/wiki/Handbook:SPARC/Installation/Kernel/it "Handbook:SPARC/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:SPARC/Installation/System/it "Handbook:SPARC/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:SPARC/Installation/Tools/it "Handbook:SPARC/Installation/Tools/it")[Impostare programma d'avvio](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it")[Concludere](/wiki/Handbook:SPARC/Installation/Finalizing/it "Handbook:SPARC/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:SPARC/Full/Working/it "Handbook:SPARC/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:SPARC/Working/Portage/it "Handbook:SPARC/Working/Portage/it")[Opzioni USE](/wiki/Handbook:SPARC/Working/USE/it "Handbook:SPARC/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:SPARC/Working/Features/it "Handbook:SPARC/Working/Features/it")[Sistema script di init](/wiki/Handbook:SPARC/Working/Initscripts/it "Handbook:SPARC/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:SPARC/Working/EnvVar/it "Handbook:SPARC/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:SPARC/Full/Portage/it "Handbook:SPARC/Full/Portage/it")[File e cartelle](/wiki/Handbook:SPARC/Portage/Files/it "Handbook:SPARC/Portage/Files/it")[Variabili](/wiki/Handbook:SPARC/Portage/Variables/it "Handbook:SPARC/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:SPARC/Portage/Branches/it "Handbook:SPARC/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:SPARC/Portage/Tools/it "Handbook:SPARC/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:SPARC/Portage/CustomTree/it "Handbook:SPARC/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:SPARC/Portage/Advanced/it "Handbook:SPARC/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:SPARC/Full/Networking/it "Handbook:SPARC/Full/Networking/it")[Come iniziare](/wiki/Handbook:SPARC/Networking/Introduction/it "Handbook:SPARC/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:SPARC/Networking/Advanced/it "Handbook:SPARC/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:SPARC/Networking/Modular/it "Handbook:SPARC/Networking/Modular/it")[Wireless](/wiki/Handbook:SPARC/Networking/Wireless/it "Handbook:SPARC/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:SPARC/Networking/Extending/it "Handbook:SPARC/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:SPARC/Networking/Dynamic/it "Handbook:SPARC/Networking/Dynamic/it")

## Contents

- [1Effettuare il chroot](#Effettuare_il_chroot)
  - [1.1Copiare le informazioni del DNS](#Copiare_le_informazioni_del_DNS)
  - [1.2Montare i filesystem necessari](#Montare_i_filesystem_necessari)
  - [1.3Entrare nel nuovo ambiente](#Entrare_nel_nuovo_ambiente)
  - [1.4Preparare per il programma d'avvio (bootloader)](#Preparare_per_il_programma_d.27avvio_.28bootloader.29)
    - [1.4.1DOS/Vecchi sistemi BIOS](#DOS.2FVecchi_sistemi_BIOS)
- [2Configurare Portage](#Configurare_Portage)
  - [2.1Installare un'istantanea del catalogo delle ebuild di Gentoo dal web](#Installare_un.27istantanea_del_catalogo_delle_ebuild_di_Gentoo_dal_web)
  - [2.2Facoltativo: Selezionare i mirror](#Facoltativo:_Selezionare_i_mirror)
    - [2.2.1Facoltativo: Mirror rsync](#Facoltativo:_Mirror_rsync)
  - [2.3Facoltativo: Aggiornare il catalogo delle ebuild di Gentoo](#Facoltativo:_Aggiornare_il_catalogo_delle_ebuild_di_Gentoo)
  - [2.4Leggere le novità](#Leggere_le_novit.C3.A0)
  - [2.5Scegliere il profilo corretto](#Scegliere_il_profilo_corretto)
  - [2.6Facoltativo: Aggiungere un server di pacchetti binari](#Facoltativo:_Aggiungere_un_server_di_pacchetti_binari)
    - [2.6.1Configurazione del catalogo](#Configurazione_del_catalogo)
    - [2.6.2Installare pacchetti binari](#Installare_pacchetti_binari)
  - [2.7Facoltativo: Configurare la variabile USE](#Facoltativo:_Configurare_la_variabile_USE)
    - [2.7.1CPU\_FLAGS\_\*](#CPU_FLAGS_.2A)
    - [2.7.2VIDEO\_CARDS](#VIDEO_CARDS)
  - [2.8Facoltativo: Configurare la variabile ACCEPT\_LICENSE](#Facoltativo:_Configurare_la_variabile_ACCEPT_LICENSE)
  - [2.9Facoltativo: aggiornare l'insieme @world](#Facoltativo:_aggiornare_l.27insieme_.40world)
    - [2.9.1Rimuovere pacchetti obsoleti](#Rimuovere_pacchetti_obsoleti)
- [3Ora locale](#Ora_locale)
- [4Configurare le localizzazioni](#Configurare_le_localizzazioni)
  - [4.1Generare le localizzazioni](#Generare_le_localizzazioni)
  - [4.2Selezionare le localizzazioni](#Selezionare_le_localizzazioni)
- [5Referenze](#Referenze)

## Effettuare il chroot

### Copiare le informazioni del DNS

Rimane da fare ancora una cosa prima di entrare nel nuovo ambiente, ovvero copiare le informazioni del DNS nel file /etc/resolv.conf. Ciò va fatto per assicurarsi che la rete funzioni anche dopo essere entrati nel nuovo ambiente. /etc/resolv.conf contiene i nameserver per la rete.

Per copiare queste informazioni, si raccomanda di passare l'opzione `--dereference` al comando cp. Ciò garantisce che, se /etc/resolv.conf è un collegamento simbolico, venga copiato il file indicato dal collegamento invece del collegamento simbolico stesso. Altrimenti, nel nuovo ambiente, il collegamento simbolico punterebbe ad un file non esistente (in quanto ciò a cui punta il collegamento probabilmente non sarà disponibile nel nuovo ambiente).

`root #` `cp --dereference /etc/resolv.conf /mnt/gentoo/etc/`

### Montare i filesystem necessari

**Suggerimento**

Se si sta usando il supporto di installazione di Gentoo, questo passaggio può essere sostituito semplicemente con: arch-chroot /mnt/gentoo.

In pochi istanti, la radice (root) di Linux sarà cambiata con quella della nuova posizione.

I filesystem da rendere disponibili sono:

- /proc/ è uno pseudo filesystem. Assomiglia ai file ordinari, ma in realtà è generato in tempo reale dal kernel Linux
- /sys/ è uno pseudo filesystem, come /proc/ che una volta si pensava l'avrebbe sostituito, ed è più strutturato di /proc/
- /dev/ è un filesystem normale che contiene tutti i dispositivi. E' parzialmente gestito dal gestore dei dispositivi di Linux (solitamente udev)
- /run/ è un filesystem temporaneo utilizzato per i file generati in fase di esecuzione, come i file PID o i lock

La posizione /proc/ verrà montata su /mnt/gentoo/proc/ mentre le altre saranno montate legandole alle esistenti. Ciò significa che, ad esempio, /mnt/gentoo/sys/ _sarà_ effettivamente /sys/ (è solo un secondo punto di accesso allo stesso filesystem) mentre /mnt/gentoo/proc/ sarà un nuovo montaggio (istanza per così dire) del filesystem.

`root #` `mount --types proc /proc /mnt/gentoo/proc
`

`root #` `mount --rbind /sys /mnt/gentoo/sys
`

`root #` `mount --make-rslave /mnt/gentoo/sys
`

`root #` `mount --rbind /dev /mnt/gentoo/dev
`

`root #` `mount --make-rslave /mnt/gentoo/dev
`

`root #` `mount --bind /run /mnt/gentoo/run
`

`root #` `mount --make-slave /mnt/gentoo/run
`

**Nota**

Le operazioni `--make-rslave` sono necessarie per il supporto di systemd più avanti nell'installazione.

**Attenzione**

Quando si usano mezzi di installazione non-Gentoo, quanto detto potrebbe risultare insufficiente. Alcune distribuzioni creano un link simbolico di /dev/shm su /run/shm/ che, dopo il chroot, risulta non valido. Fare un corretto montaggio tmpfs di /dev/shm/ può risolvere questo problema:

`root #` `test -L /dev/shm && rm /dev/shm && mkdir /dev/shm
`

`root #` `mount --types tmpfs --options nosuid,nodev,noexec shm /dev/shm`

Ci si assicuri anche che la modalità 1777 sia impostata:

`root #` ` chmod 1777 /dev/shm /run/shm`

### Entrare nel nuovo ambiente

Ora che tutte le partizioni sono state inizializzate e l'ambiente base installato, è tempo di entrare nel nuovo ambiente attraverso chroot. Ciò significa che la sessione cambierà la sua radice (la posizione di livello più alto alla quale si può fare accesso) da quella dell'attuale ambiente di installazione (del CD o di un altro mezzo di installazione) a quella del sistema installato (ovvero le partizioni inizializzate). Da cui il nome, _change root_ (cambiare radice) o _chroot_.

Questo cambio di radice è fatto in tre passaggi:

1. La posizione radice (root) è modificata da / (sul mezzo di installazione) a /mnt/gentoo/ (sulle partizioni) usando chroot o arch-chroot, se disponibile.
2. Alcune impostazioni (quelle su /etc/profile) sono ricaricate nella memoria usando il comando source.
3. L'attesa comandi (prompt) primaria viene modificata per ricordarci che questa sessione avviene in un ambiente chroot.

`root #` `chroot /mnt/gentoo /bin/bash
`

`root #` `source /etc/profile
`

`root #` `export PS1="(chroot) ${PS1}"`

Da questo momento, tutte le azioni sono direttamente eseguite nel nuovo ambiente Gentoo Linux.

**Suggerimento**

Se l'installazione di Gentoo venisse interrotta ovunque dopo questo punto, _dovrebbe_ essere possibile 'riprenderla' a partire da questo passaggio. Non c'è necessità di ri-partizionare i dischi nuovamente! Semplicemente controllare data e ora, [montare la partizione radice](/wiki/Handbook:SPARC/Installation/Disks/it#Mounting_the_root_partition "Handbook:SPARC/Installation/Disks/it") ed eseguire i passi di sopra iniziando con [copiare le informazioni del DNS](/wiki/Handbook:SPARC/Installation/Base/it#Copy_DNS_info "Handbook:SPARC/Installation/Base/it") per rientrare nell'ambiente funzionante. Questo è utile anche per risolvere alcuni problemi con il programma di avvio (bootloader). Ulteriori informazioni si possono trovare nell'articolo [chroot](/wiki/Chroot/it "Chroot/it").

### Preparare per il programma d'avvio (bootloader)

Una volta all'interno del nuovo ambiente, è necessario prepararlo per il programma di avvio. Sarà importante avere montata la partizione corretta quando sarà il momento di installare il programma di avvio.

#### DOS/Vecchi sistemi BIOS

Per DOS/Vecchi sistemi BIOS, il programma di avvio sarà installato nella cartella /boot, quindi montare come segue:

`root #` `mount  /boot`

## Configurare Portage

### Installare un'istantanea del catalogo delle ebuild di Gentoo dal web

Il prossimo passo consiste nell'installare un'istantanea del catalogo delle ebuild di Gentoo. Questa istantanea contiene una collezione di file che informano Portage riguardo i titoli di software disponibili (per l'installazione), quali profili può selezionare l'amministratore del sistema, elementi di novità per pacchetti o profili specifici, ecc.

L'uso di emerge-webrsync è raccomandato per chi si trova dietro un firewall restrittivo (usa i protocolli HTTP/FTP per scaricare l'istantanea) e fa risparmiare banda.

Questo preleverà l'ultima istantanea (la quale è rilasciata giorno per giorno) da uno dei distributori (mirror) di Gentoo e la installerà nel sistema:

`root #` `emerge-webrsync`

**Nota**

Durante questa operazione, emerge-webrsync potrebbe lamentarsi riguardo una posizione /var/db/repos/gentoo/ mancante. Ciò è atteso e niente di cui preoccuparsi - lo strumento creerà la posizione.

Da questo momento in avanti, Portage potrebbe menzionare che è consigliato eseguire certi aggiornamenti. Ciò avviene perché i pacchetti di sistema installati attraverso il file stage potrebbero avere disponibili delle versioni più recenti; Portage è ora a conoscenza dei nuovi pacchetti grazie all'istantanea del catalogo. Gli aggiornamenti dei pacchetti possono essere tranquillamente ignorati per ora; gli aggiornamenti possono essere posticipati fino alla conclusione dell'installazione di Gentoo.

### Facoltativo: Selezionare i mirror

Per scaricare rapidamente il codice sorgente si raccomanda di selezionare un mirror (distributore) veloce, geograficamente vicino. Portage cercherà la variabile GENTOO\_MIRRORS nel file make.conf ed utilizzerà i mirror lì elencati. È possibile navigare nella lista dei mirror di Gentoo e cercare un mirror (o più mirror) vicino alla posizione geografica del sistema (in quanto risulteranno solitamente i più veloci).

Uno strumento chiamato mirrorselect fornirà una piacevole interfaccia testuale per interrogare e selezionare più velocemente i mirror adatti. Spostarsi sui mirror scelti e premere la `Barra spazio` per selezionare uno o più mirror.

`root #` `emerge --ask --verbose --oneshot app-portage/mirrorselect`

`root #` `mirrorselect -i -o >> /etc/portage/make.conf`

In alternativa, una lista dei mirror attivi è [disponibile online](https://www.gentoo.org/downloads/mirrors/).

#### Facoltativo: Mirror rsync

Gentoo ha anche molti mirror rsync che possono velocizzare lo scaricamento dell'elenco dei pacchetti disponibili usando `emerge --sync` (spiegato dopo con maggiore dettaglio) dopo aver selezionato un mirror geograficamente vicino all'utente.

`root #` `mkdir /etc/portage/repos.conf
`

`root #` `cp /usr/share/portage/config/repos.conf /etc/portage/repos.conf/gentoo.conf
`

Selezionare un mirror vicino alla posizione geografica del sistema da [https://www.gentoo.org/support/rsync-mirrors/](https://www.gentoo.org/support/rsync-mirrors/) e sostituire il sync-uri del mirror predefinito in /etc/portage/repos.conf/gentoo.conf con quello della posizione del mirror desiderato.

FILE **`/etc/portage/repos.conf/gentoo.conf`** **Esempio di mirror rsync situato in UK**

```
[DEFAULT]
main-repo = gentoo

[gentoo]
location = /var/db/repos/gentoo
sync-type = rsync
sync-uri = rsync://rsync.uk.gentoo.org/gentoo-portage/
auto-sync = yes
sync-rsync-verify-jobs = 1
sync-rsync-verify-metamanifest = yes
sync-rsync-verify-max-age = 3
sync-openpgp-key-path = /usr/share/openpgp-keys/gentoo-release.asc
sync-openpgp-keyserver = hkps://keys.gentoo.org
sync-openpgp-key-refresh-retry-count = 40
sync-openpgp-key-refresh-retry-overall-timeout = 1200
sync-openpgp-key-refresh-retry-delay-exp-base = 2
sync-openpgp-key-refresh-retry-delay-max = 60
sync-openpgp-key-refresh-retry-delay-mult = 4
sync-webrsync-verify-signature = yes
sync-git-verify-commit-signature = true

```

### Facoltativo: Aggiornare il catalogo delle ebuild di Gentoo

È possibile aggiornare il catalogo delle ebuild di Gentoo all'ultima versione. Il precedente comando emerge-webrsync avrà installato un'istantanea molto recente (solitamente delle ultime 24 ore) dunque questo passo è decisamente facoltativo.

Supponiamo siano necessari gli ultimi aggiornamenti dei pacchetti (fino ad 1 ora prima), allora si usi emerge --sync. Questo comando userà il protocollo rsync per aggiornare il catalogo delle ebuild di Gentoo (il quale è stato prelevato prima attraverso emerge-webrsync) allo stato più recente.

`root #` `emerge --sync`

Sui terminali lenti, come alcuni framebuffer o console seriali, è raccomandato usare l'opzione `--quiet` per velocizzare il processo:

`root #` `emerge --sync --quiet`

### Leggere le novità

Quando il catalogo delle ebuild di Gentoo viene sincronizzato, Portage potrebbe scrivere messaggi informativi simili a questo:

` * IMPORTANT: 2 news items need reading for repository 'gentoo'.
`

` * Use eselect news to read news items.
`

Le novità (news items) sono state create per fornire un mezzo di comunicazione affinché i messaggi critici venissero notificati agli utenti attraverso il catalogo delle ebuild di Gentoo. Per gestirle, usare eselect news. L'applicazione eselect è un programma di utilità specifico di Gentoo che consente ad un'interfaccia comune di gestione l'amministrazione del sistema. In questo caso, a eselect è richiesto di usare il modulo `news`.

Per il modulo `news`, sono tre le operazioni più comuni:

- Con `list` viene mostrata una panoramica delle novità disponibili.
- Con `read` possono essere lette le novità.
- Con `purge` possono essere rimosse le novità una volta che siano state lette e non si potranno più rileggere.

`root #` `eselect news list
`

`root #` `eselect news read`

Ulteriori informazioni sul lettore delle novità sono disponibili attraverso la sua pagina di manuale:

`root #` `man news.eselect`

### Scegliere il profilo corretto

**Suggerimento**

I profili desktop non sono esclusivamente per gli _ambienti desktop_. Sono anche adatti per gestori di finestre essenziali come i3 o sway.

Un _profilo_ è un elemento fondamentale per un qualsiasi sistema Gentoo. Non solo specifica valori predefiniti per USE, CFLAGS, ed altre importanti variabili, ma blocca il sistema su un certo intervallo di versione di pacchetti. Queste impostazioni sono tutte mantenute dagli sviluppatori di Gentoo.

Per vedere quale profilo stia attualmente usando il sistema, eseguire eselect usando il modulo `profile`:

**Suggerimento**

Su un mezzo di installazione senza un terminale in grado di scorrere, `eselect profile list | less` può fornire un modo semplice di elencare tutti i profili disponibili insieme alla capacità di scorrerli su e giù.

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/sparc/23.0 *
  [2]   default/linux/sparc/23.0/desktop
  [3]   default/linux/sparc/23.0/desktop/gnome
  [4]   default/linux/sparc/23.0/desktop/kde

```

**Nota**

La risposta del comando è giusto un esempio e cambia con il tempo.

**Nota**

Per usare **systemd**, selezionare un profilo che ha "systemd" nel nome e viceversa, in caso contrario.

Ci sono anche sotto-profili per i desktop disponibili per alcune architetture che comprendono pacchetti software solitamente necessari per un esperienza desktop.

**Attenzione**

Gli aggiornamenti del profilo non devono essere presi alla leggera. Quando si seleziona il profilo iniziale, utilizzare il profilo corrispondente alla **stessa versione** di quello inizialmente utilizzato dal file stage (es. 23.0). Ogni nuova versione del profilo viene annunciata attraverso una novità contenente le istruzioni di migrazione; assicurarsi di seguirle attentamente prima di passare a un profilo più nuovo.

Dopo aver visto i profili disponibili per l'architettura sparc, gli utenti possono selezionare un profilo diverso per il sistema:

`root #` `eselect profile set 2`

**Nota**

Il sotto-profilo `developer` è specificamente per lo sviluppo di Gentoo Linux e non va considerato come una scelta per l'utente qualsiasi.

### Facoltativo: Aggiungere un server di pacchetti binari

Sin dal dicembre 2023, la [squadra di elaborazione delle versioni](/wiki/Project:RelEng "Project:RelEng") di Gentoo ha offerto un [server ufficiale di pacchetti binari](/wiki/Binary_package_quickstart "Binary package quickstart") (nel gergo comune detto brevemente "binhost") ad uso della comunità per reperire e installare pacchetti binari (binpkgs).[\[1\]](#cite_note-1)

Aggiungere un binhost permette a Portage di installare pacchetti con firma crittografica, già compilati. In molti casi, aggiungere un binhost diminuirà di gran lunga il tempo medio di installazione dei pacchetti ed è di grande beneficio quando si esegue Gentoo su sistemi vecchi, lenti o con poca potenza.

#### Configurazione del catalogo

La configurazione del catalogo per un binhost si trova nella cartella /etc/portage/binrepos.conf/ di Portage, la quale funziona in modo simile alla configurazione menzionata nella sezione del [catalogo delle ebuild di Gentoo](/wiki/Handbook:SPARC/Installation/Base/it#Gentoo_ebuild_repository "Handbook:SPARC/Installation/Base/it").

Quando si definisce un binhost, ci sono due aspetti importanti da considerare:

1. L'architettura e il profilo di destinazione contenuti nel valore `sync-uri` _hanno_ importanza e dovrebbero conformarsi alla rispettiva architettura del computer ( **sparc** in questo caso) e al profilo di sistema selezionato nella sezione [Scegliere il profilo corretto](/wiki/Handbook:SPARC/Installation/Base/it#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/it").
2. Selezionare un mirror veloce, geograficamente vicino di norma accorcia il tempo di scaricamento. Rivedi lo strumento mirrorselect menzionato nella sezione [selezionare i mirror](/wiki/Handbook:SPARC/Installation/Base/it#Gentoo_ebuild_repository "Handbook:SPARC/Installation/Base/it") o esamina l' [elenco online dei mirror](https://www.gentoo.org/downloads/mirrors/) dove possono essere rinvenuti i valori degli URL.


FILE **`/etc/portage/binrepos.conf/gentoo.conf`** **Esempio di binhost di tipologia CDN**

```
[gentoo]
priority = 9959
# NOTA: Bisogna adattare <arch> e <variant> in modo appropriato!
sync-uri = https://distfiles.gentoo.org/releases/<arch>/binpackages/<variant>
# esempio di sync-uri per x86-64
# sync-uri = https://distfiles.gentoo.org/releases/amd64/binpackages/23.0/x86-64/

# Introdotto in portage-3.0.74 per le scelte di verifica secondo il catalogo
verify-signature = true
# Valore predefinito con >=portage-3.0.77
location = /var/cache/binhost/gentoo

```

#### Installare pacchetti binari

Portage di norma compilerà i pacchetti dal codice sorgente. Può essere istruito perchè usi i pacchetti binari nei seguenti modi:

1. Si può passare l'opzione `--getbinpkg` durante l'invocazione del comando emerge. Questo metodo d'installazione dei pacchetti binari è utile per installare solo un particolare pacchetto binario.
2. Si può cambiare l'impostazione predefinita di Portage tramite la variabile FEATURES, che viene esposta attraverso il file /etc/portage/make.conf. Applicando questa modifica alla configurazione si indurrà Portage a interrogare l'ospite dei pacchetti binari per il(i) pacchetto(i) richiesti e ripiegare sulla compilazione locale qualora non si trovassero risultati.

Per esempio, affinchè Portage installi sempre i pacchetti binari disponibili:

FILE **`/etc/portage/make.conf`** **Configurare Portage perchè usi i pacchetti binari per impostazione predefinita**

```
# Aggiungere getbinpkg all'elenco dei valori nella variabile FEATURES
FEATURES="${FEATURES} getbinpkg"
# Per richiedere le firme
FEATURES="${FEATURES} binpkg-request-signature"

```

Si prega di eseguire per Portage anche getuto in modo da configurare il portachiavi necessario alla verifica:

`root #` `getuto`

Funzionalità di Portage aggiuntive saranno discusse nel [prossimo capitolo](/wiki/Handbook:SPARC/Working/Features/it#Portage_features "Handbook:SPARC/Working/Features/it") del manuale.

### Facoltativo: Configurare la variabile USE

USE è una delle più potenti variabili che Gentoo offre ai suoi utenti. Numerosi programmi possono essere compilati con o senza il supporto facoltativo di certi elementi. Per esempio, alcuni programmi possono essere compilati con il supporto a GTK+ o con il supporto a Qt. Altri possono essere compilati con o senza il supporto SSL. Alcuni programmi possono essere persino compilati con il supporto al framebuffer (svgalib) anziché il supporto a X11 (X-server).

La maggior parte delle distribuzioni compila i pacchetti con il supporto per più cose possibili, aumentando la dimensione dei programmi ed i tempi di avvio, per non menzionare l'enorme quantità di dipendenze. Con Gentoo, gli utenti possono definire quali opzioni per le quali un pacchetto dovrebbe essere compilato. Qui entra in gioco USE.

Nella variabile USE gli utenti definiscono le parole chiave che saranno mappate come opzioni di compilazione. Per esempio, `ssl` compilerà il supporto SSL nei programmi che lo supportano. `-X` rimuoverà il supporto al server X (notare il segno meno davanti). `gnome gtk -kde -qt5` compilerà i programmi con il supporto GNOME (e GTK+) e senza il supporto KDE (e Qt), mettendo il sistema perfettamente a punto per GNOME (se l'architettura lo supporta).

Le impostazioni predefinite per USE si trovano nei file make.defaults del profilo Gentoo usato dal sistema. Gentoo usa un (complesso) sistema di eredità per i profili, del quale non si tratterà a fondo nel processo di installazione. Il modo più facile per controllare le attuali impostazioni attive di USE è eseguire emerge --info e selezionare la linea che inizia con USE:

`root #` `emerge --info | grep ^USE`

```
USE="X acl alsa amd64 berkdb bindist bzip2 cli cracklib crypt cxx dri ..."

```

**Nota**

L'esempio soprastante è stato troncato, l'elenco effettivo dei valori di USE è molto, molto più grande.

Una descrizione completa delle opzioni disponibili per USE si può trovare nel sistema in /var/db/repos/gentoo/profiles/use.desc.

`root #` `less /var/db/repos/gentoo/profiles/use.desc`

Col comando less in esecuzione, si può scorrere il testo usando i tasti `↑` e `↓`, ed uscire premendo `q`.

Come esempio mostriamo un'impostazione di USE per un sistema basato su KDE con supporto DVD, ALSA e scrittura CD:

`root #` `nano /etc/portage/make.conf`

FILE **`/etc/portage/make.conf`** **Abilitare le opzioni per un sistema basato su KDE/Plasma con supporto per DVD, ALSA e la registrazione di CD**

```
USE="-gtk -gnome qt5 kde dvd alsa cdr"

```

Quando un valore di USE viene definito in /etc/portage/make.conf viene _aggiunto_ all'elenco delle opzioni USE del sistema. Le opzioni USE possono essere _rimosse_ globalmente aggiungendo un segno meno `-` davanti al valore nella lista. Per esempio, per disabilitare il supporto per l'ambiente grafico X, può essere impostato `-X`:

FILE **`/etc/portage/make.conf`** **Ignorare le opzioni predefinite di USE**

```
USE="-X acl alsa"

```

**Attenzione**

Sebbene sia possibile, l'opzione `-*` (che disabiliterà tutti i valori di USE tranne quelli specificati in make.conf) è _caldamente_ sconsigliata e imprudente. Gli sviluppatori di ebuild scelgono certi valori predefiniti delle opzioni USE per prevenire conflitti, aumentare la sicurezza, evitare errori ed altre ragioni. Disabilitare _tutte_ le opzioni USE annullerà il comportamento predefinito e potrebbe causare grossi problemi.

#### CPU\_FLAGS\_\*

Alcune architetture (incluse AMD64/X86, ARM, PPC) hanno una variabile [USE\_EXPAND](/wiki/USE_EXPAND "USE EXPAND") chiamata [CPU\_FLAGS\_<ARCH>](/wiki/CPU_FLAGS_X86 "CPU FLAGS X86"), nella quale <ARCH> viene sostituita dal nome dell'architettura pertinente.

**Importante**

Non ci si confonda! I sistemi **AMD64** e **X86** condividono parte dell'architetura, quindi il nome della variabile corretto per i sistemi **AMD64** è CPU\_FLAGS\_X86.

Questa è usata per configurare la compilazione da fare in uno specifico codice assembler o altre specificità, di solito scritte a mano o comunque supplementari, e **non** è lo stesso che chiedere al compilatore di produrre codice ottimizzato per una certa funzionalità della CPU (p.es. `-march=`).

Gli utenti dovrebbero impostare questa variabile in aggiunta alla configurazione desiderata delle proprie COMMON\_FLAGS.

Alcuni passaggi sono necessari per configurarla:

`root #` `emerge --ask --oneshot app-portage/cpuid2cpuflags`

Se curiosi, ispezionare il risultato manualmente:

`root #` `cpuid2cpuflags`

Poi copiare il risultato in package.use:

`root #` `echo "*/* $(cpuid2cpuflags)" > /etc/portage/package.use/00cpu-flags`

#### VIDEO\_CARDS

Normalmente un profilo già configura alcune schede video. Per molte ragioni non è l'ideale, ma un utile misura di sicurezza.

Per configurare il sistema correttamente l'utente ha bisogno prima di eliminare le schede video preimpostate con `VIDEO_CARDS: -*` poi di impostare la scheda corretta per il sistema.

FILE **`/etc/portage/package.use/00video_cards`** **Esempio**

```
*/* VIDEO_CARDS: -* <GPU DRIVER NAME>

```

Sotto è riportata una tabella che può essere usata per confrontare facilmente i vari tipi di schede video col rispettivo valore `VIDEO_CARDS`.

GPU
VIDEO\_CARDS
Nvidia Maxwell (Ufficiale) e seguenti`nvidia`Nouveau Nvidia [Elenco supportate](/wiki/NVIDIA "NVIDIA")`nouveau`AMD dalla Sea Islands`amdgpu radeonsi`ATI e vecchie AMDVedere [radeon#Supporto funzionalità](/wiki/Radeon#Feature_support "Radeon")Intel Nehalem e seguenti`intel`Intel Gen 2 e 3 [Elenco supportate](/wiki/Intel#.23Feature_support.2Fit "Intel")`intel i915`QEMU/KVM`virgl`WSL`d3d12`

Sotto alcuni esempi di impostazione corretta di package.use per _VIDEO\_CARDS_:

FILE **`/etc/portage/package.use/00video_cards`** **Esempio AMD**

```
*/* VIDEO_CARDS: -* amdgpu radeonsi

```

FILE **`/etc/portage/package.use/00video_cards`** **Esempio Intel**

```
*/* VIDEO_CARDS: -* intel

```

FILE **`/etc/portage/package.use/00video_cards`** **Esempio Nvidia**

```
*/* VIDEO_CARDS: -* nvidia

```

Dettagli per le diverse GPU si possono trovare negli articoli [AMDGPU](/wiki/AMDGPU "AMDGPU"), [Intel](/wiki/Intel "Intel"), [Nouveau (Open Source)](/wiki/Nouveau "Nouveau"), or [NVIDIA (Proprietario)](/wiki/NVIDIA "NVIDIA").

### Facoltativo: Configurare la variabile ACCEPT\_LICENSE

A partire dalla Gentoo Linux Enhancement Proposal 23 (GLEP 23), è stato creato un meccanismo per dare agli amministratori la possibilità di "regolamentare il software che installano in riferimento alla licenza... Alcuni vogliono un sistema libero da qualsiasi software che non sia approvato da OSI; altri sono semplicemente curiosi di quali licenze stanno implicitamente accettando." [\[2\]](#cite_note-2) Con la motivazione di avere un controllo più dettagliato sul tipo di software in esecuzione su un sistema Gentoo, è nata la variabile ACCEPT\_LICENSE.

Durante il processo di installazione, Portage tiene conto del(i) valore(i) impostato(i) nella variabile ACCEPT\_LICENSE per stabilire se il(i) pacchetto(i) richiesto(i) soddisfa(no) le scelte del sysadmin per una licenza ammissibile. In ciò si nasconde un problema: il catalogo delle ebuild di Gentoo è pieno di _migliaia_ di ebuilds che danno origine a [_centinaia_ di distinte licenze software](https://gitweb.gentoo.org/repo/gentoo.git/tree/licenses)... Ciò impone al sysadmin che approvi singolarmente ognuna e ogni nuova licenza software? Per fortuna no; GLEP 23 ha anche delineato una soluzione a questo problema, un concetto chiamato gruppi di licenze.

Per comodità dell'amministratore di sistema, le licenze software legalmente simili sono state raggruppate insieme - ognuna secondo il proprio tipo di similitudine. Definizioni di gruppi di licenze sono [disponibili per la visione](https://gitweb.gentoo.org/repo/gentoo.git/tree/profiles/license_groups) e sono gestite dal [progetto delle licenze di Gentoo](/wiki/Project:Licenses "Project:Licenses"). Mentre le license individuali non lo sono, i gruppi di licenze sintatticamente sono preceduti dal simbolo `@`, consentendogli di essere facilmente distinguibili nella variabile ACCEPT\_LICENSE.

Alcuni gruppi di licenze comuni comprendono:

Un elenco di licenze software raggruppate per tipo.
NomeDescrizione
`@GPL-COMPATIBLE`Licenze compatibili con GPL approvate dalla Free Software Foundation [\[a\_license 1\]](#cite_note-3)`@FSF-APPROVED`Licenze di software libero approvate dalla FSF (include `@GPL-COMPATIBLE`)
`@OSI-APPROVED`Licenze approvate dalla Open Source Initiative [\[a\_license 2\]](#cite_note-4)`@MISC-FREE`Licenze miste che sono probabilmente software libero, cioè seguono la Free Software Definition [\[a\_license 3\]](#cite_note-5) ma non sono approvate nè da FSF nè da OSI
`@FREE-SOFTWARE`Combina `@FSF-APPROVED`, `@OSI-APPROVED`, e `@MISC-FREE`.
`@FSF-APPROVED-OTHER`Licenze approvate dalla FSF per "documentazione libera" e "lavori di utilizzo pratico oltre al software e la documentazione" (font inclusi)
`@MISC-FREE-DOCS`Licenze miste per documenti liberi e altri lavori (font inclusi) che seguono la free definition [\[a\_license 4\]](#cite_note-6) ma NON sono elencate in `@FSF-APPROVED-OTHER`.
`@FREE-DOCUMENTS`Combina `@FSF-APPROVED-OTHER` e `@MISC-FREE-DOCS`.
`@FREE`Metainsieme di tutte le licenze con libertà d'uso, condivisione, modifica e condivisione delle modifiche. Combina `@FREE-SOFTWARE` e `@FREE-DOCUMENTS`.
`@BINARY-REDISTRIBUTABLE`Licenze che consentono almeno la ridistribuzione libera del software in forma binaria. Include `@FREE`.
`@EULA`Accordi di licenza che provano a sottrarre i propri diritti. Sono più restrittive delle "all-rights-reserved" o richiedono aprovazione esplicita.

1. [↑](#cite_ref-3)[https://www.gnu.org/licenses/license-list.html](https://www.gnu.org/licenses/license-list.html)
2. [↑](#cite_ref-4)[https://www.opensource.org/licenses](https://www.opensource.org/licenses)
3. [↑](#cite_ref-5)[https://www.gnu.org/philosophy/free-sw.html](https://www.gnu.org/philosophy/free-sw.html)
4. [↑](#cite_ref-6)[https://freedomdefined.org/](https://freedomdefined.org/)

Le licenze ammissibili per l'intero sistema attualmente impostate possono essere viste attraverso:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

Come visibile dal risultato, il valore predefinito è di permettere per l'installazione solo software raggruppato nella categoria `@FREE`.

Licenze specifiche o gruppi di licenze per un sistema possono essere definite nelle seguenti posizioni:

- Per l'intero sistema nel profilo selezionato - ciò imposta il valore predefinito.
- Per l'intero sistema nel file /etc/portage/make.conf. Gli amministratori di sistema scavalcano il valore predefinito del profilo in questo file.
- Per pacchetto nel file /etc/portage/package.license.
- Per pacchetto nella _cartella_ di file /etc/portage/package.license/.

La licenza predefinita del profilo per l'intero sistema viene scavalcata nel file /etc/portage/make.conf:

FILE **`/etc/portage/make.conf`** **Ammettere licenze per l'intero sistema con ACCEPT\_LICENSE**

```
# Scavalca il valore predefinito del profilo di ACCEPT_LICENSE
ACCEPT_LICENSE="-* @FREE @BINARY-REDISTRIBUTABLE"

```

Opzionalmente gli amministratori di sistema possono anche definire le licenze ammesse per pacchetto come mostrato nella seguente cartella di file di esempio. Notare che la _cartella_ package.license ha bisogno di essere creata se non dovesse già esistere:

`root #` `mkdir /etc/portage/package.license`

I dettagli della licenza software per un singolo pacchetto di Gentoo sono memorizzati nella variabile LICENSE della ebuild associata. Un pacchetto potrebbe avere una o più licenze software, quindi potrebbe essere necessario specificare licenze ammissibili multiple per un singolo pacchetto.

FILE **`/etc/portage/package.license/kernel`** **Ammettere licenze per pacchetto**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware linux-fw-redistributable
sys-firmware/intel-microcode intel-ucode

```

**Importante**

La variabile LICENSE in una ebuild è solo una indicazione per gli sviluppatori Gentoo e gli utenti. Non è una dichiarazione con valore legale, e _non c'è alcuna garanzia_ che corrisponda alla realtà. Si consiglia di non fidarsi solo dell'interpretazione dello sviluppatore di una ebuild nei confronti della licenza del pacchetto software; invece controllare il pacchetto stesso in profondità, inclusi tutti i file che sono stati installati sul sistema.

### Facoltativo: aggiornare l'insieme @world

**Suggerimento**

Se è stato selezionato un profilo per un ambiente desktop da un file stage non-desktop, il processo di aggiornamento dell'insieme @world potrebbe aumentare notevolmente la quantità di tempo necessaria per il processo di installazione. Coloro che hanno un tempo ristretto possono lavorare seguendo questa 'regola pratica': più corto è il nome del profilo, meno specifico sarà l' [insieme @world](/wiki/World_set_(Portage) "World set (Portage)") del sistema. Meno specifico sarà l'insieme @world, meno pacchetti saranno richiesti dal sistema. P.es.:

- Selezionare `default/linux/amd64/23.0` probabilmente comporterà meno pacchetti da aggiornare, mentre
- Selezionare `default/linux/amd64/23.0/desktop/gnome/systemd` probabilmente richiederà più pacchetti da installare, poichè il profilo scelto ha degli insiemi @system e @profile più ampi: le dipendenze di supporto all'ambiente desktop GNOME.

Aggiornare l' [insieme @world](/wiki/World_set_(Portage) "World set (Portage)") è _facoltativo_ ed è improbabile che effettui cambiamenti funzionali a meno che uno o più dei seguenti passaggi facoltativi sia stato effettuato:

1. E' stato selezionato un profilo _diverso_ da quello del file stage.
2. Sono state impostate opzioni USE aggiuntive per i pacchetti installati.

I lettori che stanno eseguendo una 'corsa di velocità nell'installare Gentoo' possono tranquillamente saltare l'aggiornamento dell'insieme @world fino a _dopo_ il riavvio del loro sistema nel nuovo ambiente Gentoo.

I lettori che stanno eseguendo un cammino lento possono far effettuare ora a Portage gli aggiornamenti per i pacchetti, il profilo e/o le modifiche delle opzioni USE:

`root #` `emerge --ask --verbose --update --deep --changed-use @world`

I lettori che hanno aggiunto un binhost [sopra](/wiki/Handbook:SPARC/Installation/Base/it#Optional:_Adding_a_binary_package_host "Handbook:SPARC/Installation/Base/it") possono aggiungere --getbinpkg (o -g) per fare in modo di recuperare i pacchetti dal binhost piuttosto che compilarli:

`root #` `emerge --ask --verbose --update --deep --newuse --getbinpkg @world`

#### Rimuovere pacchetti obsoleti

E' sempre importante effettuare il _depclean_ dopo gli aggiornamenti al sistema per rimuovere pacchetti obsoleti. Esamina il risultato attentamente con emerge --depclean --pretend per vedere se qualcuno dei pacchetti da eliminare debba essere invece mantenuto perchè usato personalmente. Per mantenere un pacchetto che altrimenti verrebbe eliminato, usare emerge --noreplace foo.

`root #` `emerge --ask --pretend --depclean`

Se si è soddisfatti, allora procedere col depclean reale:

`root #` `emerge --ask --depclean`

## Ora locale

**Nota**

Questo passaggio non si applica agli utenti di musl libc. Chi non sa cosa significa dovrebbe eseguire questo passaggio.

**Attenzione**

Si evitino le ore locali /usr/share/zoneinfo/Etc/GMT\* poiché i loro nomi non indicano le aree attese. Per esempio, GMT-8 è di fatto GMT+8.

Scegliere l'ora locale per il sistema. Cercare le ore locali disponibili su /usr/share/zoneinfo/:

`root #` `ls -l /usr/share/zoneinfo`

```
total 352
drwxr-xr-x 2 root root   1120 Jan  7 17:41 Africa
drwxr-xr-x 6 root root   2960 Jan  7 17:41 America
drwxr-xr-x 2 root root    280 Jan  7 17:41 Antarctica
drwxr-xr-x 2 root root     60 Jan  7 17:41 Arctic
drwxr-xr-x 2 root root   2020 Jan  7 17:41 Asia
drwxr-xr-x 2 root root    280 Jan  7 17:41 Atlantic
drwxr-xr-x 2 root root    500 Jan  7 17:41 Australia
drwxr-xr-x 2 root root    120 Jan  7 17:41 Brazil
-rw-r--r-- 1 root root   2094 Dec  3 17:19 CET
-rw-r--r-- 1 root root   2310 Dec  3 17:19 CST6CDT
drwxr-xr-x 2 root root    200 Jan  7 17:41 Canada
drwxr-xr-x 2 root root     80 Jan  7 17:41 Chile
-rw-r--r-- 1 root root   2416 Dec  3 17:19 Cuba
-rw-r--r-- 1 root root   1908 Dec  3 17:19 EET
-rw-r--r-- 1 root root    114 Dec  3 17:19 EST
-rw-r--r-- 1 root root   2310 Dec  3 17:19 EST5EDT
-rw-r--r-- 1 root root   2399 Dec  3 17:19 Egypt
-rw-r--r-- 1 root root   3492 Dec  3 17:19 Eire
drwxr-xr-x 2 root root    740 Jan  7 17:41 Etc
drwxr-xr-x 2 root root   1320 Jan  7 17:41 Europe
...

```

`root #` `ls -l /usr/share/zoneinfo/Europe/`

```
total 256
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Amsterdam
-rw-r--r-- 1 root root 1742 Dec  3 17:19 Andorra
-rw-r--r-- 1 root root 1151 Dec  3 17:19 Astrakhan
-rw-r--r-- 1 root root 2262 Dec  3 17:19 Athens
-rw-r--r-- 1 root root 3664 Dec  3 17:19 Belfast
-rw-r--r-- 1 root root 1920 Dec  3 17:19 Belgrade
-rw-r--r-- 1 root root 2298 Dec  3 17:19 Berlin
-rw-r--r-- 1 root root 2301 Dec  3 17:19 Bratislava
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Brussels
...

```

Supponendo che l'ora locale scelta sia Europe/Brussels, per selezionarla, si può creare un collegamento da questo file in zoneinfo a /etc/localtime:

`root #` `ln -sf ../usr/share/zoneinfo/Europe/Brussels /etc/localtime`

**Suggerimento**

Il percorso di destinazione con `../` all'inizio è _relativo alla posizione del collegamento_, non alla cartella corrente.

**Nota**

Si può usare un percorso assoluto per il collegamento, ma un collegamento relativo è creato anche da timedatectl di systemd ed è più compatibile con _ROOT_ alternative.

## Configurare le localizzazioni

**Nota**

Questo passaggio non si applica agli utenti di musl libc. Chi non sa cosa significa dovrebbe eseguire questo passaggio.

### Generare le localizzazioni

Un'installazione predefinita di Gentoo Linux fornisce un archivio che contiene tutte le localizzazioni supportate, per un ammontare di 500 o più. Tuttavia per un amministratore è tipico averne bisogno solo una o due. In questo caso, il file di configurazione /etc/locale.gen può essere riempito con un elenco delle localizzazioni richieste. Di norma, locale-gen leggerà questo file e compilerà solo le localizzazioni che sono specificate, risparmiando sia tempo che spazio a lungo termine.

Le localizzazioni non specificano solamente la lingua che l'utente dovrà usare per interagire con il sistema, ma anche le regole per ordinare le stringhe, mostrare le date e gli orari, ecc. Le localizzazioni _distinguono tra maiuscole e miniscole_ e devono essere rappresentate esattamente come descritte. Un elenco completo di tutte le localizzazioni disponibili si può trovare nel file /usr/share/i18n/SUPPORTED.

`root #` `nano /etc/locale.gen`

Le localizzazioni seguenti sono un esempio per avere sia l'inglese (Stati Uniti) sia il tedesco (Germania/Deutschland) con i formati di carattere corrispondenti (come UTF-8).

FILE **`/etc/locale.gen`** **Abilitare le localizzazioni US e DE con i formati dei caratteri appropriati**

```
en_US.UTF-8 UTF-8
de_DE.UTF-8 UTF-8

# Le localizzazioni non UTF-8 sono sconsigliate e dovrebbero essere usate solo in speciali circostanze.
#en_US ISO-8859-1
#de_DE ISO-8859-1

```

**Attenzione**

Molte applicazioni richiedono almeno una localizzazione UTF-8 per compilare correttamente.

Il passo successivo è eseguire il comando locale-gen. Questo comando genererà tutte le localizzazioni specificate nel file /etc/locale.gen.

`root #` `locale-gen`

Per verificare che le localizzazioni selezionate sono ora disponibili, eseguire locale -a.

Nelle installazioni con systemd, si può usare localectl, p.es. localectl set-locale ... o localectl list-locales.

### Selezionare le localizzazioni

Fatto ciò, si passa ad impostare le localizzazioni per tutto il sistema. Si userà nuovamente eselect, questa volta con il modulo `locale`.

Con eselect locale list vengono mostrate le opzioni disponibili:

`root #` `eselect locale list`

```
Available targets for the LANG variable:
  [1]  C
  [2]  C.UTF8
  [3]  POSIX
  [4]  de_DE.UTF8
  [5]  en_US.UTF8
  [ ]  (free form)

```

Con eselect locale set <NUMBER>, si può selezionare la localizzazione corretta:

`root #` `eselect locale set 2`

Manualmente, ciò può essere ottenuto tramite il file /etc/env.d/02locale, e per systemd il file /etc/locale.conf:

FILE **`/etc/env.d/02locale`** **Impostare manualmente le definizioni delle localizzazioni del sistema**

```
LANG="de_DE.UTF-8"
LC_COLLATE="C.UTF-8"

```

L'impostazione della localizzazione eviterà avvisi ed errori durante le compilazioni del kernel e del software più avanti nell'installazione.

Ora ricaricare l'ambiente:

`root #` `env-update && source /etc/profile && export PS1="(chroot) ${PS1}"`

Per assistenza aggiuntiva nel processo di selezione della localizzazione, leggere anche la [Guida alla localizazzione](/wiki/Localization/Guide/it "Localization/Guide/it") e la guida [UTF-8](/wiki/UTF-8/it "UTF-8/it").

## Referenze

1. [↑](#cite_ref-1)[https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html](https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html)
2. [↑](#cite_ref-2)[https://www.gentoo.org/glep/glep-0023.html#id7](https://www.gentoo.org/glep/glep-0023.html#id7)

[← Installare il file stage](/wiki/Handbook:SPARC/Installation/Stage/it "Previous part") [Home](/wiki/Handbook:SPARC/it "Handbook:SPARC/it") [Configurare il kernel →](/wiki/Handbook:SPARC/Installation/Kernel/it "Next part")