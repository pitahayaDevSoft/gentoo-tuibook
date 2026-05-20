# Stage

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Stage/de "Handbuch:MIPS/Installation/Stage (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Stage "Handbook:MIPS/Installation/Stage (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Stage/es "Manual de Gentoo: MIPS/Instalación/Stage (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Stage/fr "Handbook:MIPS/Installation/Stage/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:MIPS/Installation/Stage/hu "Handbook:MIPS/Installation/Stage/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Stage/pl "Handbook:MIPS/Installation/Stage (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Stage/pt-br "Manual:MIPS/Instalação/Stage (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Stage/ru "Handbook:MIPS/Installation/Stage (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Stage/ta "கையேடு:MIPS/நிறுவல்/நிலை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Stage/zh-cn "手册：MIPS/安装/安装stage3 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Stage/ja "ハンドブック:MIPS/インストール/Stage (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Stage/ko "Handbook:MIPS/Installation/Stage/ko (100% translated)")

[Manuale MIPS](/wiki/Handbook:MIPS/it "Handbook:MIPS/it")[Installazione](/wiki/Handbook:MIPS/Full/Installation/it "Handbook:MIPS/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:MIPS/Installation/About/it "Handbook:MIPS/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:MIPS/Installation/Media/it "Handbook:MIPS/Installation/Media/it")[Configurare la rete](/wiki/Handbook:MIPS/Installation/Networking/it "Handbook:MIPS/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:MIPS/Installation/Disks/it "Handbook:MIPS/Installation/Disks/it")Il file stage[Installare il sistema base](/wiki/Handbook:MIPS/Installation/Base/it "Handbook:MIPS/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:MIPS/Installation/Kernel/it "Handbook:MIPS/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:MIPS/Installation/System/it "Handbook:MIPS/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:MIPS/Installation/Tools/it "Handbook:MIPS/Installation/Tools/it")[Impostare programma d'avvio](/wiki/Handbook:MIPS/Installation/Bootloader/it "Handbook:MIPS/Installation/Bootloader/it")[Concludere](/wiki/Handbook:MIPS/Installation/Finalizing/it "Handbook:MIPS/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:MIPS/Full/Working/it "Handbook:MIPS/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:MIPS/Working/Portage/it "Handbook:MIPS/Working/Portage/it")[Opzioni USE](/wiki/Handbook:MIPS/Working/USE/it "Handbook:MIPS/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:MIPS/Working/Features/it "Handbook:MIPS/Working/Features/it")[Sistema script di init](/wiki/Handbook:MIPS/Working/Initscripts/it "Handbook:MIPS/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:MIPS/Working/EnvVar/it "Handbook:MIPS/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:MIPS/Full/Portage/it "Handbook:MIPS/Full/Portage/it")[File e cartelle](/wiki/Handbook:MIPS/Portage/Files/it "Handbook:MIPS/Portage/Files/it")[Variabili](/wiki/Handbook:MIPS/Portage/Variables/it "Handbook:MIPS/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:MIPS/Portage/Branches/it "Handbook:MIPS/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:MIPS/Portage/Tools/it "Handbook:MIPS/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:MIPS/Portage/CustomTree/it "Handbook:MIPS/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:MIPS/Portage/Advanced/it "Handbook:MIPS/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:MIPS/Full/Networking/it "Handbook:MIPS/Full/Networking/it")[Come iniziare](/wiki/Handbook:MIPS/Networking/Introduction/it "Handbook:MIPS/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:MIPS/Networking/Advanced/it "Handbook:MIPS/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:MIPS/Networking/Modular/it "Handbook:MIPS/Networking/Modular/it")[Wireless](/wiki/Handbook:MIPS/Networking/Wireless/it "Handbook:MIPS/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:MIPS/Networking/Extending/it "Handbook:MIPS/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:MIPS/Networking/Dynamic/it "Handbook:MIPS/Networking/Dynamic/it")

## Contents

- [1Scegliere un file stage](#Scegliere_un_file_stage)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
  - [1.3Multilib (32 e 64-bit)](#Multilib_.2832_e_64-bit.29)
  - [1.4No-multilib (64-bit puro)](#No-multilib_.2864-bit_puro.29)
- [2Scaricare il file stage](#Scaricare_il_file_stage)
  - [2.1Impostare data e ora](#Impostare_data_e_ora)
    - [2.1.1Automaticamente](#Automaticamente)
    - [2.1.2Manualmente](#Manualmente)
  - [2.2Browser grafici](#Browser_grafici)
  - [2.3Browser a linea di comando](#Browser_a_linea_di_comando)
  - [2.4Verificare e validare](#Verificare_e_validare)
- [3Installare un file stage](#Installare_un_file_stage)
- [4Configurare le opzioni di compilazione](#Configurare_le_opzioni_di_compilazione)
  - [4.1Introduzione](#Introduzione)
  - [4.2CFLAGS e CXXFLAGS](#CFLAGS_e_CXXFLAGS)
  - [4.3RUSTFLAGS](#RUSTFLAGS)
  - [4.4MAKEOPTS](#MAKEOPTS)
  - [4.5Pronti, partenza, via!](#Pronti.2C_partenza.2C_via.21)
- [5Referenze](#Referenze)

## Scegliere un file stage

**Suggerimento**

Sulle architetture compatibili, agli utenti che ambiscono ad un sistema operativo dall'ambiente desktop (grafico) si consiglia di usare un file stage col termine `desktop` nel nome. Questi file includono pacchetti come [llvm-core/llvm](https://packages.gentoo.org/packages/llvm-core/llvm) e [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) e ottimizzazioni delle opzione USE che miglioreranno enormemente il tempo di installazione.

Il [file stage](/wiki/Stage_file "Stage file") agisce da base per una installazione di Gentoo. I file stage sono generati dalla [squadra di elaborazione delle versioni](/wiki/Project:RelEng "Project:RelEng") con [Catalyst](/wiki/Catalyst "Catalyst"). I file stage sono basati su specifici [profili](/wiki/Profile_(Portage) "Profile (Portage)") e contengono un sistema quasi completo.

Quando si sceglie un file stage, è importante prenderne uno con propositi del profilo corrispondenti al tipo di sistema desiderato.

**Importante**

Anche se è possibile fare importanti modifiche al profilo dopo che l'installazione è stata avviata, cambiare richiede sforzo e analisi considerevoli, e va al di là dello scopo di questo manuale di installazione. Cambiare il sistema di init è difficile, ma cambiare da `no-multilib` a `multilib` richiede una conoscenza di Gentoo e degli strumenti di basso livello dettagliata.

**Suggerimento**

La maggior parte degli utenti non dovrebbe aver bisogno di usare le opzioni 'avanzate' dei tarball; queste riguardano configurazioni software o hardware atipiche o avanzate.

### OpenRC

[OpenRC](/wiki/OpenRC/it "OpenRC/it") è un sistema di init (responsabile dell'avvio dei servizi di sistema una volta avviato il kernel) basato su dipendenze che mantiene la compatibilità col programma di init fornito dal sistema, generalmente situato in /sbin/init. E' il sistema di init originale e nativo di Gentoo, ma è anche adottato da alcune altre distribuzioni Linux e sistemi BSD.

### systemd

systemd è un recente sistema di init in stile Sys-V e sostitutivo di rc per i sistemi Linux. E' usato come sistema di init principale dalla maggioranza delle distribuzioni Linux. systemd è completamente supportato da Gentoo e funziona per lo scopo prefissato. Se qualcosa sembra mancare nel manuale durante una installazione con systemd, consultare l' [articolo systemd](/wiki/Systemd/it "Systemd/it") _prima_ di richiedere assistenza.

### Multilib (32 e 64-bit)

**Nota**

Non tutte le architetture hanno una opzione multilib. Molte funzionano solo con codice nativo. Multilib è più comunemente adoperato con **amd64**.

Il profilo multilib usa le librerie a 64-bit quando possibile, e ripiega su quelle a 32-bit solo quando strettamente necessario alla compatibilità. Questa è un opzione eccellente per la maggioranza delle installazioni perchè fornisce un enorme grado di flessibilità per la personalizzazione futura.

**Suggerimento**

Usare oggetti `multilib` rende più facile cambiare profili in seguito, rispetto a `no-multilib`

### No-multilib (64-bit puro)

**Attenzione**

I lettori che stanno iniziando proprio ora con Gentoo _non_ dovrebbero scegliere una tarball no-multilib a meno che non sia assolutamente necessario. Migrare da un sistema `no-multilib` ad uno `multilib` richiede una conoscenza di Gentoo e di strumenti di basso livello molto collaudata, (ciò potrebbe far rabbrividire persino i nostri [sviluppatori di strumenti (toolchain)](/wiki/Project:Toolchain "Project:Toolchain")). Non è per i deboli di cuore e va oltre lo scopo di questa guida.

Scegliere una tarball no-multilib come base per il sistema offre un ambiente completamente a 64-bit - libero da software a 32-bit. Ciò di fatto rende la possibiltà di passaggio ad un profilo multilib gravoso, seppur ancora tecnicamente possibile.

## Scaricare il file stage

Prima di scaricare il _file stage_, la cartella corrente dovrebbe essere impostata alla posizione del punto di mount usato per l'installazione:

`root #` `cd /mnt/gentoo`

### Impostare data e ora

Gli archivi stage sono generalmente ottenuti usando HTTPS che necessita di un orario di sistema relativamente accurato. Deviazioni nell'orologio possono impedire il funzionamento degli scaricamenti, e può causare errori imprevedibili se l'orario di sistema viene corretto di una quantità rilevante dopo l'installazione.

Data e ora correnti possono essere controllate con date:

`root #` `date`

```
Lun 3 Ott 13:16:22 UTC 2021

```

Se la data e l'ora mostrate sono sbagliate di più di qualche minuto, dovrebbero essere corrette usando uno dei seguenti metodi.

#### Automaticamente

Usare [Network\_Time\_Protocol](/wiki/Network_Time_Protocol "Network Time Protocol") per correggere la deviazione dell'orologio è in genere più semplice è più affidabile che impostarlo manualmente.

chronyd, facente parte di [net-misc/chrony](https://packages.gentoo.org/packages/net-misc/chrony) può essere usato per aggiornare l'orologio di sistema a UTC con:

`root #` `chronyd -q`

**Importante**

I sistemi senza un Real-Time Clock (RTC) funzionante devono sincronizzare l'orologio di sistema ad ogni avvio, e poi ad intervalli regolari. Ciò è di beneficio anche per i sistemi con un RTC, poichè la batteria può scaricarsi, e l'orologio può accumulare ritardi.

**Attenzione**

Il traffico NTP ordinario non è autenticato, è importante verificare i dati dell'ora ottenuti dalla rete.

#### Manualmente

Quando non è disponibile l'accesso a NTP, si può usare date per impostare manualmente l'orologio di sistema.

**Nota**

L'orario UTC è raccomandato su tutti i sistemi Linux. Successivamente, viene definito un fuso orario di sistema, che modificherà della giusta differenza l'ora quando viene mostrata la data.

Per impostare l'orario viene usato il seguente formato dell' argomento: sintassi di `MMGGoommAAAA` ( **M** ese, **G** iorno, **o** ra, **m** inuto e **A** nno).

Per esempio, per impostare la data sul 3 ottobre, 13:16 nell'anno 2021, dare:

`root #` `date 100313162021`

### Browser grafici

Coloro che usano ambienti con browser di rete completamente grafici non avranno problemi a copiare l'URL del file stage dalla [sezione di download](https://www.gentoo.org/downloads/#other-arches) del sito principale. Semplicemente selezionare la scheda appropriata, cliccare con il tasto destro del mouse sul link al file stage, poi su Copia link per copiare l'indirizzo negli appunti, quindi incollare il link nel programma di utilità wget da linea di comando così da scaricare il file stage:

`root #` `wget <URL_FILE_STAGE_INCOLLATO>`

### Browser a linea di comando

I lettori più tradizionalisti o gli utenti Gentoo di 'vecchia data', che lavorano esclusivamente da linea di comando, potrebbero preferire l'uso di links([www-client/links](https://packages.gentoo.org/packages/www-client/links)), un browser non grafico e funzionante tramite menu. Per scaricare uno stage, navigare fino all'elenco dei distributori (mirror) di Gentoo così:

`root #` `links https://www.gentoo.org/downloads/mirrors/`

Per usare un proxy HTTP con links, passare l'URL con l'opzione `-http-proxy`:

`root #` `links -http-proxy proxy.server.com:8080 https://www.gentoo.org/downloads/mirrors/`

Accanto a links c'è anche il browser lynx ([www-client/lynx](https://packages.gentoo.org/packages/www-client/lynx)). Come links, si tratta di un browser non grafico ma senza menu.

`root #` `lynx https://www.gentoo.org/downloads/mirrors/`

Se è necessario definire un proxy, esportare le variabili http\_proxy e/o ftp\_proxy:

`root #` `export http_proxy="http://proxy.server.com:port"
`

`root #` `export ftp_proxy="http://proxy.server.com:port"`

Nella lista dei distributori (mirror), selezionarne uno vicino. Solitamente vanno bene i mirror HTTP, ma sono disponibili anche altri protocolli. Spostarsi nella cartella releases/mips/autobuilds/. Qui sono elencati tutti i file stage disponibili (potrebbero essere posti in sottocartelle nominate con il nome delle singole sotto architetture). Selezionarne uno e premere `d` per scaricarlo.

Appena lo scaricamento del file stage è completato, è possibile verificarne l'integrità e validarne i contenuti. Coloro che sono interessati dovrebbero leggere la [prossima sezione](/wiki/Handbook:MIPS/Installation/Stage/it#Verifying_and_validating "Handbook:MIPS/Installation/Stage/it").

Coloro che non sono interessati alla verifica e alla validazione del file stage possono chiudere il browser a linea di comando premendo `q` e possono andare direttamente alla sezione [installare il file stage](/wiki/Handbook:MIPS/Installation/Stage/it#Unpacking_the_stage_tarball "Handbook:MIPS/Installation/Stage/it").

### Verificare e validare

Come con i CD di installazione minimali, sono disponibili file aggiuntivi per verificare e validare il file stage. I file aggiuntivi sono disponibili nella cartella radice dei mirror. Navigare fino alla posizione corretta per l'architettura hardware e il profilo di sistema e scaricare i file .CONTENTS.gz, .DIGESTS e .sha256 associati.

`root #` `wget https://distfiles.gentoo.org/releases/`

- .CONTENTS.gz \- Un file compresso che contiene un elenco di tutti i file all'interno del file stage.
- .DIGESTS \- Contiene i checksum (somme di controllo) del file stage, usando parecchi algoritmi di calcolo crittografico.
- .sha256 \- Contiene un checksum SHA256 del file stage firmato con PGP. Questo file potrebbe non essere disponibile al download per tutti i file stage.

Così come per il file ISO, la firma crittografica del file tar.xz deve essere verificata usando gpg per assicurarsi che non siano avvenute manomissioni alla tarball.

Per le immagini live ufficiali di Gentoo, il pacchetto [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release) fornisce le chiavi di firma PGP per rilasci automatizzati. Le chiavi vanno prima importate nella sessione dell'utente in modo da usarle per la verifica:

`root #` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

Per tutte le immagini live non ufficiali che offrono gpg e wget nell'ambiente live, riferirsi alla sezione antecedente [verificare i file scaricati](/wiki/Handbook:MIPS/Installation/Media/it#Verifying_the_downloaded_files "Handbook:MIPS/Installation/Media/it").

Verificare la firma della tarball e, facoltativamente, i file di checksum associati:

`root #` `gpg --verify stage3-mips-<release>-<init>.tar.xz.asc stage3-mips-<release>-<init>.tar.xz
`

`root #` `gpg --output stage3-mips-<release>-<init>.tar.xz.DIGESTS.verified --verify stage3-mips-<release>-<init>.tar.xz.DIGESTS
`

`root #` `gpg --output stage3-mips-<release>-<init>.tar.xz.sha256.verified --verify stage3-mips-<release>-<init>.tar.xz.sha256
`

Se la verifica ha successo, "Good signature from" verrà mostrato dal(i) precedente(i) comando(i).

Le impronte delle chiavi PGP usate per firmare i supporti delle varie versioni possono essere trovate sulla [pagina delle firme dei supporti delle varie versioni](https://www.gentoo.org/downloads/signatures/).

**Nota**

La maggior parte degli stage ora [riporta aggiunto al nome](https://www.gentoo.org/news/2021/07/20/more-downloads.html) esplicitamente il tipo di sistema di init (openrc o systemd), sebbene alcune architetture potrebbero al momento non averlo ancora.

Strumenti crittografici e programmi di utilità come openssl, sha256sum o sha512sum possono essere usati per confrontare il loro risultato con le somme di controllo fornite dal file .DIGESTS.

Per verificare il checksum SHA512 con openssl:

`root #` `openssl dgst -r -sha512 stage3-mips-<release>-<init>.tar.xz`

`dgst` indica al comando openssl di usare il sotto-comando Message Digest, `-r` stampa il risultato del calcolo del digest nel formato coreutils e `-sha512` seleziona il formato SHA512 per il calcolo del digest.

Per verificare il checksum BLAKE2B512 con openssl:

`root #` `openssl dgst -r -blake2b512 stage3-mips-<release>-<init>.tar.xz`

Confrontare il(i) risultato(i) dei comandi di checksum con le coppie di valori nome del file e codice hash contenuti nel file .DIGESTS.verified. Le coppie di valori devono corrispondere al risultato dei comandi di checksum, altrimenti il file scaricato è corrotto e dovrebbe essere eliminato e riscaricato.

Per verificare l'hash SHA256 da un file .sha256 associato usando il programma di utilità sha256sum:

`root #` `sha256sum --check stage3-mips-<release>-<init>.tar.xz.sha256.verified`

L'opzione `--check` indica a sha256sum di leggere un elenco di file attesi e gli hash associati, e poi stampa "OK" in corrispondenza di ogni file che calcola correttamente o "FAILED" per quelli no.

## Installare un file stage

Una volta che il _file stage_ è stato scaricato e verificato, può essere estratto usando tar:

`root #` `tar xpvf stage3-*.tar.xz --xattrs-include='*.*' --numeric-owner -C /mnt/gentoo`

Prima dell'estrazione verificare le opzioni:

- `x` e **x** tract, indica a tar di estarre il contenuto dell'archivio.
- `p` **p** reserve, preserva i permessi.
- `v` **v** erbose, mostra messaggi dettagliati.
- `f` **f** ile, fornisce a tar il nome dell'archivio in ingresso.
- `--xattrs-include='*.*'` Preserva gli attributi estesi in tutti i namespace memorizzati nell'archivio.
- `--numeric-owner` Assicura che gli ID degli utenti e dei gruppi per i file estratti dalla tarball rimangano uguali a quelli pianificati dalla squadra di elaborazione della versione di Gentoo (anche se gli utenti avventurosi non stanno usando per il processo di installazione un ambiente live di Gentoo ufficiale).
- `-C /mnt/gentoo` Estrae i file nella partizione di root a prescindere dalla cartella corrente.

Ora che il file stage è stato estratto, procedere con la [configurazione delle opzioni di compilazione](/wiki/Handbook:MIPS/Installation/Stage/it#Configuring_compile_options "Handbook:MIPS/Installation/Stage/it").

## Configurare le opzioni di compilazione

### Introduzione

Per ottimizzare il sistema, è possibile impostare variabili che influenzano il comportamento di Portage, il gestore di pacchetti di Gentoo ufficialmente supportato. Tutte queste variabili possono essere impostate come variabili d'ambiente (usando export) però l'impostazione tramite export non è permanente.

**Nota**

Tecnicamente le variabili possono essere esportate tramite il profilo della [shell](/wiki/Shell "Shell") o i file rc, tuttavia ciò non è il metodo migliore per l'amministrazione di base del sistema.

Portage, quando lavora, legge nel file [make.conf](/wiki/Make.conf "Make.conf"), che ne modificherà il comportamento durante l'esecuzione a seconda dei valori lì salvati. make.conf può essere considerato il file di configurazione principale di Portage, quindi si tratti il suo contenuto attentamente.

**Suggerimento**

Si può trovare un elenco commentato di tutte le possibili variabili nel file /mnt/gentoo/usr/share/portage/config/make.conf.example. Documentazione aggiuntiva su make.conf si può trovare eseguendo man 5 make.conf.

Perchè una installazione di Gentoo riesca bene è necessario impostare solo le variabili mensionate sotto.

Si apra un editor (in questa guida si usa nano) per modificare le variabili di ottimizzazione che discuteremo a seguire.

`root #` `nano /mnt/gentoo/etc/portage/make.conf`

Dal file make.conf.example si comprende come questo dovrebbe essere strutturato: le linee commentate iniziano con `#`, le altre linee definiscono le variabili usando la sintassi `VARIABILE="valore"`. Parecchie di queste variabili saranno discusse nella prossima sezione.

### CFLAGS e CXXFLAGS

Le variabili CFLAGS e CXXFLAGS definiscono rispettivamente le opzioni di ottimizzazione per i compilatori GCC C e C++. Benché siano definite in maniera generale qui, per ottenere le migliori prestazioni si dovrebbero ottimizzare queste opzioni separatamente per ogni programma. Questo perchè ciascun programma è diverso. Tuttavia, ciò non è gestibile, per cui si definiscono queste variabili (flag) nel file make.conf.

Nel file make.conf si dovrebbero definire le opzioni di ottimizzazione che renderanno il sistema il più performante possibile in generale. Non si usino impostazioni sperimentali in questa variabile; troppe ottimizzazioni possono far si che i programmi si comportino in malo modo (blocchi, o ancora peggio, malfunzionamenti).

Il manuale non spiegherà tutte le opzioni di ottimizzazione possibili. Per comprenderle tutte, si legga il [Manuale GNU Online](https://gcc.gnu.org/onlinedocs/) o la pagina di informazioni di gcc (info gcc). Il file make.conf.example stesso contiene inoltre tantissimi esempi ed informazioni; non si dimentichi di leggerlo pure.

Una prima impostazione è l'opzione `-march=` o `-mtune=`, che specifica il nome dell'architettura di destinazione. Nel file make.conf.example sono descritte le possibili opzioni (in forma di commenti). Un valore comunemente usato è _native_ che indica al compilatore di selezionare come architettura di destinazione quella del sistema in uso (quella su cui gli utenti stanno installando Gentoo).

Un'altra opzione è `-O` (O maiuscola, non zero), che specifica la classe di ottimizzazione di gcc. Possibili classi sono s (per ottimizzare le dimensioni), 0 (zero - per non ottimizzare affatto), 1, 2 o persino 3 per ottimizzare la velocità (ogni classe ha le stesse opzioni della precedente, più alcune aggiuntive). `-O2` è l'opzione predefinita consigliata. È noto che `-O3` causa problemi se usata per l'intero sistema, quindi si consiglia di permanere su `-O2`.

Un'altra opzione di ottimizzazione popolare è `-pipe` (usa le _pipe_ \- passaggio dati - invece di file temporanei per la comunicazione tra le varie fasi della compilazione). Non ha alcun impatto sul codice generato, ma usa più memoria. Su sistemi con poca memoria, gcc potrebbe interrompere l'esecuzione. In tal caso, non si usi questa opzione.

Usare `-fomit-frame-pointer` (che non conserva il frame pointer nel registro per le funzioni che non ne hanno bisogno) potrebbe avere serie ripercussioni sulle operazioni di debug delle applicazioni.

Quando vengono definite le variabili CFLAGS e CXXFLAGS, si combinino le varie opzioni di ottimizzazione in un'unica stringa. I valori predefiniti contenuti nel file stage dovrebbero essere abbastanza buoni. Quello di seguito è solo un esempio:

CODE **Esempio delle variabili CFLAGS e CXXFLAGS**

```
# Opzioni del compilatore da impostare per tutti i linguaggi di programmazione
COMMON_FLAGS="-mabi=32 -mips4 -pipe -O2"
# Usare le stesse impostazioni per ambo le variabili
CFLAGS="${COMMON_FLAGS}"
CXXFLAGS="${COMMON_FLAGS}"

```

**Suggerimento**

Benché l'articolo [ottimizzazione di GCC](/wiki/GCC_optimization/it "GCC optimization/it") abbia più informazioni su come le varie opzioni di compilazione possano influenzare un sistema, l'articolo [CFLAGS sicure](/wiki/Safe_CFLAGS "Safe CFLAGS") potrebbe essere un posto più adatto ai principianti per iniziare ad ottimizzare i propri sistemi.

### RUSTFLAGS

Molti programmi sono ora scritti in Rust che ha il proprio modo di ottimizzare. Di norma Rust ottimizza a livello 3 su tutte le versioni compilate a meno che un progetto lo cambi, quindi dovrebbe essere lasciato così com'è. Una lista completa di ottimizzazioni (conosciuta come codegen) che può essere passata al compilatore Rust è disponibile a [https://doc.rust-lang.org/rustc/codegen-options/index.html](https://doc.rust-lang.org/rustc/codegen-options/index.html).

L'ottimizzazione più utile sarebbe di comunicare a Rust di compilare per la propria CPU usando il seguente esempio:

FILE **`/etc/portage/make.conf`** **Esempio di RUSTFLAGS**

```
RUSTFLAGS="${RUSTFLAGS} -C target-cpu=native"

```

Per avere un elenco di CPU supportate in Rust eseguire:

`root #` `rustc -C target-cpu=help`

Ciò mostrerà il valore predefinito e anche quale CPU verrà selezionata come nativa.

**Nota**

Il comando sopra riportato funziona solo su tarball stage3 desktop o dopo aver fatto l'emerge di [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) o [dev-lang/rust](https://packages.gentoo.org/packages/dev-lang/rust).

### MAKEOPTS

La variabile MAKEOPTS definisce quante compilazioni in parallelo dovrebbero avvenire quando si installa un pacchetto. Dalla versione 3.0.31 di Portage[\[1\]](#cite_note-1), se lasciata non impostata, il comportamento predefinito di Portage è di impostare il valore dei j(ob) in MAKEOPTS uguale al numero di thread restituito da nproc.

Inoltre, dalla versione 3.0.53 di Portage[\[2\]](#cite_note-2), se lasciata non impostata, il comportamento predefinito di Portage è di impostare il valore l(oad-average) in MAKEOPTS uguale al numero di thread restituito da nproc.

Una buona scelta è il minore tra: il numero dei thread della CPU, o l'ammontare totale della RAM di sistema divisa per 2 GiB.

**Attenzione**

Usare un grande numero di job può significativamente impattare il consumo di memoria. Un buon consiglio è di avere almeno 2 GiB di ram per ogni job impostato (quindi, p.es. `-j6` richiede _almeno_ 12 GiB). Per evitare di esaurire tutta la memoria, diminuire il numero dei job affinchè rientrino nella memoria disponibile.

**Suggerimento**

Quando si eseguono emerge in parallelo ( `--jobs`), il numero effettivo dei job in esecuzione può crescere esponenzialmente (fino ai job di make moltiplicati per i job di emerge). Ciò può essere aggirato eseguendo distcc in configurazione solo-localhost che limiterà il numero di istanze di compilazione per computer.

FILE **`/etc/portage/make.conf`** **Esempio di dichiarazione di MAKEOPTS**

```
# Se lasciato non impostato, il comportamento predefinito di Portage è:
# - impostare il valore dei j(ob) di MAKEOPTS uguale al numero di thread restituito da `nproc`
# - impostare il valore di l(oad-average) di MAKEOPTS leggermente superiore al numero di thread restituito da `nproc`, trattandosi di un valore attenuato
# Si prega di sostituire '4' con valori appropriati al sistema (minore tra (RAM/2GB, nr. di thread)), o lasciare non impostato.
MAKEOPTS="-j4 -l5"

```

Cercare MAKEOPTS in man 5 make.conf per maggiori dettagli.

### Pronti, partenza, via!

Aggiornare il file /mnt/gentoo/etc/portage/make.conf in base alle preferenze personali e salvarlo (gli utenti di nano dovranno premere `Ctrl` + `o` per scrivere le modifiche e poi `Ctrl` + `x` per uscire).

## Referenze

1. [↑](#cite_ref-1)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2](https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2)
2. [↑](#cite_ref-2)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e](https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e)

[← Preparare i dischi](/wiki/Handbook:MIPS/Installation/Disks/it "Previous part") [Home](/wiki/Handbook:MIPS "Handbook:MIPS") [Installare il sistema base →](/wiki/Handbook:MIPS/Installation/Base/it "Next part")