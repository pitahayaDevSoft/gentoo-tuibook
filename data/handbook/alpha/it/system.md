# System

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/System/de "Handbuch:Alpha/Installation/System (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/System "Handbook:Alpha/Installation/System (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/System/es "Manual de Gentoo: Alpha/Instalación/Sistema (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/System/fr "Handbook:Alpha/Installation/System/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:Alpha/Installation/System/hu "Handbook:Alpha/Installation/System/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/System/pl "Handbook:Alpha/Installation/System (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/System/pt-br "Manual:Alpha/Instalação/Sistema (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/System/cs "Handbook:Alpha/Installation/System/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/System/ru "Handbook:Alpha/Installation/System (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/System/ta "கையேடு:Alpha/நிறுவல்/முறைமை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/System/zh-cn "手册：Alpha/安装/配置系统 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/System/ja "ハンドブック:Alpha/インストール/システム (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/System/ko "Handbook:Alpha/Installation/System/ko (100% translated)")

[Manuale Alpha](/wiki/Handbook:Alpha/it "Handbook:Alpha/it")[Installazione](/wiki/Handbook:Alpha/Full/Installation/it "Handbook:Alpha/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:Alpha/Installation/About/it "Handbook:Alpha/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:Alpha/Installation/Media/it "Handbook:Alpha/Installation/Media/it")[Configurare la rete](/wiki/Handbook:Alpha/Installation/Networking/it "Handbook:Alpha/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:Alpha/Installation/Disks/it "Handbook:Alpha/Installation/Disks/it")[Il file stage](/wiki/Handbook:Alpha/Installation/Stage/it "Handbook:Alpha/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:Alpha/Installation/Base/it "Handbook:Alpha/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:Alpha/Installation/Kernel/it "Handbook:Alpha/Installation/Kernel/it")Configurare il sistema[Installare gli strumenti](/wiki/Handbook:Alpha/Installation/Tools/it "Handbook:Alpha/Installation/Tools/it")[Impostare programma d'avvio](/wiki/Handbook:Alpha/Installation/Bootloader/it "Handbook:Alpha/Installation/Bootloader/it")[Concludere](/wiki/Handbook:Alpha/Installation/Finalizing/it "Handbook:Alpha/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:Alpha/Full/Working/it "Handbook:Alpha/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:Alpha/Working/Portage/it "Handbook:Alpha/Working/Portage/it")[Opzioni USE](/wiki/Handbook:Alpha/Working/USE/it "Handbook:Alpha/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:Alpha/Working/Features/it "Handbook:Alpha/Working/Features/it")[Sistema script di init](/wiki/Handbook:Alpha/Working/Initscripts/it "Handbook:Alpha/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:Alpha/Working/EnvVar/it "Handbook:Alpha/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:Alpha/Full/Portage/it "Handbook:Alpha/Full/Portage/it")[File e cartelle](/wiki/Handbook:Alpha/Portage/Files/it "Handbook:Alpha/Portage/Files/it")[Variabili](/wiki/Handbook:Alpha/Portage/Variables/it "Handbook:Alpha/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:Alpha/Portage/Branches/it "Handbook:Alpha/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:Alpha/Portage/Tools/it "Handbook:Alpha/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:Alpha/Portage/CustomTree/it "Handbook:Alpha/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:Alpha/Portage/Advanced/it "Handbook:Alpha/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:Alpha/Full/Networking/it "Handbook:Alpha/Full/Networking/it")[Come iniziare](/wiki/Handbook:Alpha/Networking/Introduction/it "Handbook:Alpha/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:Alpha/Networking/Advanced/it "Handbook:Alpha/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:Alpha/Networking/Modular/it "Handbook:Alpha/Networking/Modular/it")[Wireless](/wiki/Handbook:Alpha/Networking/Wireless/it "Handbook:Alpha/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:Alpha/Networking/Extending/it "Handbook:Alpha/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:Alpha/Networking/Dynamic/it "Handbook:Alpha/Networking/Dynamic/it")

## Contents

- [1Informazioni sul filesystem](#Informazioni_sul_filesystem)
  - [1.1Etichette e UUID dei filesystem](#Etichette_e_UUID_dei_filesystem)
  - [1.2Etichette e UUID delle partizioni](#Etichette_e_UUID_delle_partizioni)
  - [1.3Riguardo fstab](#Riguardo_fstab)
  - [1.4Creare il file fstab](#Creare_il_file_fstab)
    - [1.4.1DOS/Vecchi sistemi BIOS](#DOS.2FVecchi_sistemi_BIOS)
- [2Informazioni di rete](#Informazioni_di_rete)
  - [2.1Hostname (nome del PC)](#Hostname_.28nome_del_PC.29)
    - [2.1.1Impostare l'hostname (OpenRC o systemd)](#Impostare_l.27hostname_.28OpenRC_o_systemd.29)
    - [2.1.2systemd](#systemd)
  - [2.2Rete](#Rete)
    - [2.2.1DHCP tramite dhcpcd (qualsiasi sistema di init)](#DHCP_tramite_dhcpcd_.28qualsiasi_sistema_di_init.29)
    - [2.2.2netifrc (OpenRC)](#netifrc_.28OpenRC.29)
      - [2.2.2.1Configurare la rete](#Configurare_la_rete)
      - [2.2.2.2Avvio automatico dei servizi di rete](#Avvio_automatico_dei_servizi_di_rete)
  - [2.3Il file degli host](#Il_file_degli_host)
- [3Informazioni di sistema](#Informazioni_di_sistema)
  - [3.1Password di root](#Password_di_root)
  - [3.2Configurazione del sistema di init e dell'avvio](#Configurazione_del_sistema_di_init_e_dell.27avvio)
    - [3.2.1OpenRC](#OpenRC)
    - [3.2.2systemd](#systemd_2)

## Informazioni sul filesystem

### Etichette e UUID dei filesystem

Sia MBR (BIOS) che GPT includono il supporto per le etichette e gli UUID dei _filesystem_. Questi attributi possono essere definiti in /etc/fstab come alternative per il comando mount da usare quando si tenta di trovare e montare i dispositivi a blocchi. Le etichette e gli UUID dei filesystem sono identificati dai prefissi LABEL e UUID e possono essere visionati con il comando blkid:

`root #` `blkid`

**Attenzione**

Se il filesystem dentro la partizione viene eliminato, allora i valori delle etichette e degli UUID del filesystem saranno conseguentemente alterati o rimossi.

Per la loro unicità, si raccomanda ai lettori, che stanno usando tabelle delle partizioni in stile MBR, di usare gli UUID anziché le etichette per specificare i volumi da montare nel file /etc/fstab.

**Importante**

Le UUID del filesystem su un volume LVM e le sue istantanee LVM sono identiche, quindi usare le UUID per montare volumi LVM dovrebbe essere evitato.

### Etichette e UUID delle partizioni

I sistemi con supporto alle etichette disco GPT offrono 'robuste' opzioni aggiuntive per definire le partizioni in /etc/fstab. Le etichette e gli UUID delle partizioni possono essere usate per identificare la(e) singola(e) partizione(i) del dispositivo a blocchi, senza badare a quale filesystem è stato scelto per la partizione stessa. Le etichette e gli UUID delle partizioni sono identificati dai prefissi PARTLABEL e/o PARTUUID e possono essere comodamente visionati da terminale eseguendo il comando blkid:

Il risultato per un un sistema EFI **amd64** che usa le UUID del Discoverable Partition Specification (DPS) potrebbe essere simile al seguente:

`root #` `blkid`

```
/dev/sr0: BLOCK_SIZE="2048" UUID="2023-08-28-03-54-40-00" LABEL="ISOIMAGE" TYPE="iso9660" PTTYPE="PMBR"
/dev/loop0: TYPE="squashfs"
/dev/sda2: PARTUUID="0657fd6d-a4ab-43c4-84e5-0933c84b4f4f"
/dev/sda3: PARTUUID="1cdf763a-5b4c-4dbf-99db-a056c504e8b2"
/dev/sda1: PARTUUID="c12a7328-f81f-11d2-ba4b-00a0c93ec93b"

```

Mentre non è vero sempre per le etichette delle partizioni, usare un UUID per identificare una partizione in fstab fornisce una garanzia che il programma di avvio non farà confusione quando ricercherà un certo volume, persino se il _filesystem_ cambiasse o venisse riscritto in futuro. Usare i vecchi file predefiniti dei dispositivi a blocchi (/dev/sd\*N) per definire le partizioni in fstab è rischioso per i sistemi che hanno dispositivi a blocchi SATA che vengono regolarmente aggiunti o rimossi.

La denominazione dei file dei dispositivi a blocchi dipende da un certo numero di fattori, incluso come ed in quale ordine i dischi sono collegati al sistema. Potrebbero anche apparire in un ordine differente in base a quali dispositivi vengono rilevati prima dal kernel durante la prima fase del processo di avvio. Premesso ciò, a meno che uno non intenda avere continuamente a che fare con l'ordinamento dei dischi, usare i file predefiniti per i dispositivi a blocchi è un approccio semplice e diretto.

### Riguardo fstab

Su Linux, tutte le partizioni usate dal sistema devono essere elencate in [/etc/fstab](/wiki//etc/fstab/it "/etc/fstab/it"). Questo file contiene i punti di montaggio di queste partizioni (cioè dove esse verranno viste nella struttura del filesystem), come dovranno essere montate e con quali opzioni speciali (automaticamente o meno, se gli utenti potranno montarle o meno, ecc.).

### Creare il file fstab

**Nota**

Se il sistema di init che si sta usando è systemd, se le UUID delle partizioni rispettano la Discoverable Partition Specification (DPS) come specificato in [Preparare i dischi](/wiki/Handbook:Alpha/Installation/Disks/it "Handbook:Alpha/Installation/Disks/it"), e se il sistema usa UEFI, allora si può saltare la creazione dell'fstab poichè systemd monta automaticamente le partizioni che rispettano la specifica.

Il file /etc/fstab usa una sintassi in stile tabella. Ogni linea consiste di sei campi, separati da spazio bianco (uno o più spazi, tabulazioni o un misto dei due). Ogni campo ha il suo specifico significato:

1. Il primo campo mostra il dispositivo a blocchi speciale o il filesystem in remoto che deve essere montato. Svariati tipi di identificatori di dispositivi sono disponibili per i nodi dei dispositivi a blocchi speciali, compreso percorsi a file di dispositivo, etichette e UUID dei filesystem, ed etichette e UUID di partizioni.
2. Il secondo campo mostra il punto di montaggio sul quale la partizione andrebbe montata.
3. Il terzo campo mostra il tipo di filesystem usato dalla partizione
4. Il quarto campo mostra le opzioni di montaggio usate da mount quando vuole montare la partizione. Siccome ogni filesystem ha le sue opzioni di montaggio, si incoraggiano gli amministratori di sistema a leggere la pagina del manuale di _mount_ (man mount) per un elenco completo. Opzioni multiple di montaggio vanno separate con una virgola.
5. Il quinto campo è usato da _dump_ (copia del contenuto della memoria) per determinare se la partizione necessita di un dump oppure no. Questo valore viene generalmente lasciato a `0` (zero).
6. Il sesto campo è usato da fsck per determinare l'ordine in cui i file system dovrebbero essere controllati quando il sistema non viene arrestato correttamente. Il filesystem radice (root) dovrebbe avere `1` mentre i rimanenti dovrebbero avere `2` (oppure `0` se il controllo del file system non è necessario).

**Importante**

Il file predefinito /etc/fstab fornito nei file stage di Gentoo _non_ è un file fstab valido ma piuttosto un modello che può essere usato per immetere i valori pertinenti.

`root #` `nano /etc/fstab`

#### DOS/Vecchi sistemi BIOS

Diamo un'occhiata a come scrivere le opzioni per la partizione /boot/. Questo è solo un esempio e dovrebbe essere modificato secondo le decisioni di partizionamento prese all'inizio dell'installazione.
Nell'esempio di partizionamento per **alpha**, /boot/ è solitamente la partizione /dev/sda1, con xfs consigliato come filesystem. Esso necessita di essere controllato durante l'avvio, dunque dovremmo scrivere:

FILE **`/etc/fstab`** **Un esempio per la riga di boot DOS/Legacy BIOS per /etc/fstab**

```
# Adattare per formattazioni diverse e/o partizioni aggiuntive create nel passaggio "Preparare i dischi"
/dev/sda1   /boot     ext2    defaults        0 2

```

Alcuni amministratori di sistema vogliono che la loro partizione /boot/ non sia montata automaticamente per migliorare la sicurezza del proprio sistema. Tali persone dovrebbero sostituire il `defaults` con `noauto` Ciò significa che tali utenti avranno bisogno di montare manualmente questa partizione ogni volta che vorranno usarla.

Aggiungere regole che corrispondano al precedente schema di partizionamento stabilito ed aggiungere regole per i dispositivi come i lettori CD-ROM e, naturalmente, per ogni altra partizione o disco in uso.

Di seguito un esempio molto più elaborato del file /etc/fstab:

FILE **`/etc/fstab`** **Un esempio completo di /etc/fstab per un sistema DOS/vecchio BIOS**

```
# Adattare per formattazioni diverse e/o partizioni aggiuntive create nel passaggio "Preparare i dischi"
/dev/sda1   /boot        ext2        defaults             0 2
/dev/sda2   none         swap    sw                   0 0
/dev/sda3   /            xfs     defaults,noatime     0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

[Handbook:Alpha/Blocks/Fstab/it](/index.php?title=Handbook:Alpha/Blocks/Fstab/it&action=edit&redlink=1 "Handbook:Alpha/Blocks/Fstab/it (page does not exist)")

Quando si usa `auto` nel terzo campo, si lascia che il comando mount indovini quale filesystem ci sia. Questo è raccomandato per i supporti rimovibili in quanto possono essere creati con uno tra tanti filesystem. L'opzione `user` nel quarto campo rende possibile ad un utente non-root il montaggio del CD.

Per migliorare le prestazioni, molti utenti vorranno aggiungere l'opzione di mount `noatime`, col risultato di un sistema più veloce poichè non vengono registrati gli orari di accesso (non sono comunque necessari in genere). Ciò è anche consigliato per sistemi con dischi a stato solido (SSD). Gli utenti potrebbero invece voler prendere in considerazione `lazytime`.

**Suggerimento**

A causa del degrado prestazionale, definire l'opzione di mount `discard` in /etc/fstab non è consigliato. Di solito è meglio programmare l'eliminazione dei blocchi periodicamente usando un pianificatore di azioni come cron o un temporizzatore (systemd). Vedere [operazioni di fstrim periodiche](/wiki/SSD#Periodic_fstrim_jobs "SSD") per maggiori informazioni.

Controllare due volte il file /etc/fstab, poi salvare ed uscire per continuare.

## Informazioni di rete

E' importante notare che le sezioni seguenti sono fornite per aiutare il lettore a configurare velocemente il sistema per partecipare ad una rete locale.

Per i sistemi che eseguono OpenRC, una fonte più dettagliata per l'impostazione della rete è disponibile nella sezione [advanced network configuration](/wiki/Handbook:Alpha/Networking/Introduction/it "Handbook:Alpha/Networking/Introduction/it"), che viene trattata verso la fine del manuale. Sistemi con esigenze di rete più specifiche potrebbero aver bisogno di saltare avanti e poi tornare qui per continuare con il resto dell'installazione.

Per l'impostazione più specifica della rete in systemd, si prega di esaminare la [porzione della rete](/wiki/Systemd/it#Network "Systemd/it") dell'articolo [systemd](/wiki/Systemd/it "Systemd/it").

### Hostname (nome del PC)

Una delle scelte che l'amministratore di sistema deve compiere è dare un nome al suo PC. Questo sembra piuttosto facile, ma molti utenti hanno difficoltà a trovare il nome appropriato per l'hostname. Per velocizzare le cose, si sappia che la decisione non è definitiva - può essere cambiata in seguito. Nell'esempio sottostante viene usato l'hostname _tux_.

#### Impostare l'hostname (OpenRC o systemd)

`root #` `echo tux > /etc/hostname`

#### systemd

Per impostare l'hostname di un sistema che al momento _esegue_ systemd, può essere usato il programma di utilità hostnamectl. Durante il processo di installazione comunque, il comando [systemd-firstboot](/wiki/Handbook:Alpha/Installation/System/it#Init_and_boot_configuration_systemd "Handbook:Alpha/Installation/System/it") deve essere invece usato (vedere più avanti nel manuale).

Per impostare _tux_ come hostname, si dovrà eseguire:

`root #` `hostnamectl hostname tux`

Guardare l'aiuto eseguendo hostnamectl --help o man 1 hostnamectl.

### Rete

Esistono molte opzioni disponibili per configurare le interfacce di rete. Questa sezione tratta solo alcuni metodi. Si scelga quello che sembra più adatto all'impostazione richiesta.

#### DHCP tramite dhcpcd (qualsiasi sistema di init)

La maggior parte delle reti LAN ha in funzione un server DHCP. Se ci si trova in questa eventualità, si consiglia di usare il programma dhcpcd per ottenere un indirizzo IP.

Per installarlo:

`root #` `emerge --ask net-misc/dhcpcd`

Per abilitarlo e poi avviare il servizio su sistemi OpenRC:

`root #` `rc-update add dhcpcd default
`

`root #` `rc-service dhcpcd start
`

Per abilitare il servizio su sistemi systemd:

`root #` `systemctl enable dhcpcd`

Con questi passaggi completati, la prossima volta che il sistema si avvia, dhcpcd dovrebbe ottenere un indirizzo IP dal server DHCP. Vedere l'articolo [Dhcpcd](/wiki/Dhcpcd "Dhcpcd") per maggiori dettagli

#### netifrc (OpenRC)

**Suggerimento**

Questo è un modo particolare di impostare la rete usando [Netifrc](/wiki/Netifrc "Netifrc") su OpenRC. Esistono altri metodi per configurazioni più semplici come [Dhcpcd](/wiki/Dhcpcd "Dhcpcd").

##### Configurare la rete

Durante l'installazione di Gentoo Linux, la rete era già configurata. Comunque, ciò valeva per l'ambiente live stesso e non per l'ambiente installato. Adesso, la configurazione di rete viene fatta per il sistema Gentoo Linux installato.

**Nota**

Informazioni più dettagliate sul funzionamento della rete, inclusi argomenti avanzati come accoppiamenti (bonding), interconnessioni (bridging), VLAN 802.1Q o reti wireless sono trattati nella sezione sulla [configurazione di rete avanzata](/wiki/Handbook:Alpha/Networking/Introduction/it "Handbook:Alpha/Networking/Introduction/it").

Tutte le informazioni di rete sono raccolte su /etc/conf.d/net. La sintassi è lineare ma forse non ancora intuitiva. Ma niente paura! Ogni cosa sarà spiegata di seguito. Un esempio completamente commentato che tratta molte diverse configurazioni è disponibile in /usr/share/doc/netifrc-\*/net.example.bz2.

Prima installare [net-misc/netifrc](https://packages.gentoo.org/packages/net-misc/netifrc):

`root #` `emerge --ask --noreplace net-misc/netifrc`

DHCP è usato in modo predefinito. Affinché DHCP funzioni, è necessario installare un client DHCP. Ciò viene descritto più avanti negli Strumenti di sistema necessari per l'installazione.

Se la connessione di rete necessita di essere configurata a causa di opzioni specifiche per DHCP o perché DHCP non è usato affatto, allora si apra /etc/conf.d/net:

`root #` `nano /etc/conf.d/net`

Impostare sia config\_eth0 che routes\_eth0 per inserire le informazioni sull'indirizzo IP e sull'instradamento.

**Nota**

Si presume che l'interfaccia di rete sia chiamata eth0. Comunque, ciò dipende strettamente dal sistema. È raccomandato assumere che l'interfaccia sia nominata con il nome dell'interfaccia presente all'avvio col mezzo di installazione, se quest'ultimo è abbastanza recente. Maggiori informazioni si possono trovare nella sezione [denominazione delle interfacce di rete](/wiki/Handbook:Alpha/Networking/Advanced/it#Network_interface_naming "Handbook:Alpha/Networking/Advanced/it").

FILE **`/etc/conf.d/net`** **Definizione di un IP statico**

```
config_eth0="192.168.0.2 netmask 255.255.255.0 brd 192.168.0.255"
routes_eth0="default via 192.168.0.1"

```

Per usare DHCP, definire config\_eth0:

FILE **`/etc/conf.d/net`** **Definizione di DHCP**

```
config_eth0="dhcp"

```

Si legga /usr/share/doc/netifrc-\*/net.example.bz2 per un elenco di opzioni di configurazione aggiuntive. Assicurarsi di controllare bene anche la pagina del manuale del client DHCP se è necessario impostare opzioni specifiche per DHCP.

Se il sistema ha svariate interfacce di rete, allora ripetere i passaggi precedenti per config\_eth1, config\_eth2, ecc.

Adesso salvare la configurazione ed uscire per proseguire.

##### Avvio automatico dei servizi di rete

Per avere le interfacce di rete attive all'avvio, esse vanno aggiunte al runlevel (livello di esecuzione) predefinito.

`root #` `cd /etc/init.d
`

`root #` `ln -s net.lo net.eth0
`

`root #` `rc-update add net.eth0 default`

Se il sistema ha diverse interfacce di rete, allora gli appropriati file net.\* devono essere creati come è stato fatto con net.eth0.

Se, dopo aver avviato il sistema, si scopre che il nome dell'interfaccia di rete (attualmente documentata come `eth0`) sia sbagliato, allora si eseguano i seguenti passaggi per correggerlo:

1. Aggiornare il file /etc/conf.d/net con il nome di interfaccia corretto (per esempio `enp3s0` o `enp5s0`, anziché `eth0`).
2. Creare un nuovo link simbolico (esempio /etc/init.d/net.enp3s0).
3. Rimuovere il vecchio link simbolico (rm /etc/init.d/net.eth0).
4. Aggiungere quello nuovo al runlevel (livello di esecuzione) predefinito.
5. Rimuovere quello vecchio con rc-update del net.eth0 default.

### Il file degli host

Un importante passaggio successivo può essere informare il nuovo sistema riguardo altri computer nel suo ambiente di rete. I nomi dei computer della rete possono essere definiti nel file /etc/hosts. Aggiungere qui i nomi dei computer abiliterà la risoluzione degli indirizzi IP per quei computer dei quali tale risoluzione non viene eseguita dal nameserver.

`root #` `nano /etc/hosts`

FILE **`/etc/hosts`** **Compilazione delle informazioni di rete**

```
# Questo definisce il sistema stesso e deve essere impostato
127.0.0.1     tux.homenetwork tux localhost
::1           tux.homenetwork tux localhost

# Definizione opzionale per altri sistemi sulla rete
192.168.0.5   jenny.homenetwork jenny
192.168.0.6   benny.homenetwork benny

```

Salvare ed uscire dall'editor per continuare.

## Informazioni di sistema

### Password di root

Impostare la password per l'utente root usando il comando passwd.

`root #` `passwd`

In seguito, verrà creato un account utente ordinario aggiuntivo per le operazioni quotidiane.

### Configurazione del sistema di init e dell'avvio

#### OpenRC

Quando si adopera OpenRC con Gentoo , si usa /etc/rc.conf per configurare i servizi, l'avvio e lo spegnimento del sistema. Aprire /etc/rc.conf e divertirsi con tutti i commenti nel file. Rivedere le impostazioni e cambiarle dove necessario.

`root #` `nano /etc/rc.conf`

Poi, aprire /etc/conf.d/keymaps per gestire la configurazione della tastiera. Modificare il file per configurare ed impostare correttamente la tastiera.

`root #` `nano /etc/conf.d/keymaps`

Si presti particolare attenzione alla variabile keymap. Se viene selezionata la mappa dei tasti sbagliata, quando si digita sulla tastiera verranno visualizzati strani risultati.

Infine, modificare /etc/conf.d/hwclock per impostare le opzioni dell'orologio. Modificarlo secondo le preferenze personali.

`root #` `nano /etc/conf.d/hwclock`

Se l'orologio hardware non sta usando UTC, allora è necessario impostare `clock="local"` nel file. Nel caso contrario, ciò comporterebbe disallineamenti nell'ora del sistema.

#### systemd

Innanzitutto, si consiglia di eseguire systemd-machine-id-setup e poi systemd-firstboot che faranno si che svariati componenti di sistema siano configurati correttamente per il primo avvio nel nuovo ambiente systemd. Il passaggio delle opzioni seguenti includerà l'intervento dell'utente per impostare i valori di localizzazione, fuso orario, hostname, password di root e shell di root. Verrà assegnato anche un ID macchina casuale per l'installazione:

`root #` `systemd-machine-id-setup`

`root #` `systemd-firstboot --prompt`

Successivamente gli utenti dovranno eseguire systemctl per azzerare tutti i file delle unità installate ai valori preimpostati delle regole:

**Nota**

E' possibile che venga mostrato un avviso di insuccesso dopo l'esecuzione del seguente comando. Ciò è normale, poichè il LiveCD di Gentoo usa OpenRC.

`root #` `systemctl preset-all --preset-mode=enable-only`

E' possibile eseguire l'azzeramento completo ai valori preimpostati ma ciò potrebbe azzerare anche le impostazioni di quei servizi già configurati durante l'installazione:

`root #` `systemctl preset-all`

Questi due passaggi aiuteranno ad assicurare una transizione tranquilla dall'ambiente live al primo avvio dell'installazione.

[← Configurare il kernel](/wiki/Handbook:Alpha/Installation/Kernel/it "Previous part") [Home](/wiki/Handbook:Alpha/it "Handbook:Alpha/it") [Installare gli strumenti →](/wiki/Handbook:Alpha/Installation/Tools/it "Next part")