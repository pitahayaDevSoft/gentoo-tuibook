# Tools

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Tools/de "Handbuch:X86/Installation/Tools (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Tools "Handbook:X86/Installation/Tools (100% translated)")
- [español](/wiki/Handbook:X86/Installation/Tools/es "Manual de Gentoo: X86/Instalación/Herramientas (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Tools/fr "Handbook:X86/Installation/Tools/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:X86/Installation/Tools/hu "Handbook:X86/Installation/Tools/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Tools/pl "Handbook:X86/Installation/Tools (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Tools/pt-br "Manual:X86/Instalação/Ferramentas (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Tools/cs "Handbook:X86/Installation/Tools/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Tools/ru "Handbook:X86/Installation/Tools (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Tools/ta "கையேடு:X86/நிறுவல்/கருவிகள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Tools/zh-cn "手册：X86/安装/安装系统工具 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Tools/ja "ハンドブック:X86/インストール/ツールのインストール (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Tools/ko "Handbook:X86/Installation/Tools/ko (100% translated)")

[Manuale X86](/wiki/Handbook:X86/it "Handbook:X86/it")[Installazione](/wiki/Handbook:X86/Full/Installation/it "Handbook:X86/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:X86/Installation/About/it "Handbook:X86/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:X86/Installation/Media/it "Handbook:X86/Installation/Media/it")[Configurare la rete](/wiki/Handbook:X86/Installation/Networking/it "Handbook:X86/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:X86/Installation/Disks/it "Handbook:X86/Installation/Disks/it")[Il file stage](/wiki/Handbook:X86/Installation/Stage/it "Handbook:X86/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:X86/Installation/Base/it "Handbook:X86/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:X86/Installation/Kernel/it "Handbook:X86/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:X86/Installation/System/it "Handbook:X86/Installation/System/it")Installare gli strumenti[Impostare programma d'avvio](/wiki/Handbook:X86/Installation/Bootloader/it "Handbook:X86/Installation/Bootloader/it")[Concludere](/wiki/Handbook:X86/Installation/Finalizing/it "Handbook:X86/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:X86/Full/Working/it "Handbook:X86/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:X86/Working/Portage/it "Handbook:X86/Working/Portage/it")[Opzioni USE](/wiki/Handbook:X86/Working/USE/it "Handbook:X86/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:X86/Working/Features/it "Handbook:X86/Working/Features/it")[Sistema script di init](/wiki/Handbook:X86/Working/Initscripts/it "Handbook:X86/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:X86/Working/EnvVar/it "Handbook:X86/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:X86/Full/Portage/it "Handbook:X86/Full/Portage/it")[File e cartelle](/wiki/Handbook:X86/Portage/Files/it "Handbook:X86/Portage/Files/it")[Variabili](/wiki/Handbook:X86/Portage/Variables/it "Handbook:X86/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:X86/Portage/Branches/it "Handbook:X86/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:X86/Portage/Tools/it "Handbook:X86/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:X86/Portage/CustomTree/it "Handbook:X86/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:X86/Portage/Advanced/it "Handbook:X86/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:X86/Full/Networking/it "Handbook:X86/Full/Networking/it")[Come iniziare](/wiki/Handbook:X86/Networking/Introduction/it "Handbook:X86/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:X86/Networking/Advanced/it "Handbook:X86/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:X86/Networking/Modular/it "Handbook:X86/Networking/Modular/it")[Wireless](/wiki/Handbook:X86/Networking/Wireless/it "Handbook:X86/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:X86/Networking/Extending/it "Handbook:X86/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:X86/Networking/Dynamic/it "Handbook:X86/Networking/Dynamic/it")

## Contents

- [1Logger _(registratore degli eventi)_ di sistema](#Logger_.28registratore_degli_eventi.29_di_sistema)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
- [2Servizio cron](#Servizio_cron)
  - [2.1OpenRC](#OpenRC_2)
    - [2.1.1Predefinito: cronie](#Predefinito:_cronie)
    - [2.1.2In alternativa: dcron](#In_alternativa:_dcron)
    - [2.1.3In alternativa: fcron](#In_alternativa:_fcron)
    - [2.1.4In alternativa: bcron](#In_alternativa:_bcron)
    - [2.1.5Utenti modprobed-db](#Utenti_modprobed-db)
  - [2.2systemd](#systemd_2)
    - [2.2.1Utenti modprobed-db](#Utenti_modprobed-db_2)
- [3Facoltativo: Indicizzazione dei file](#Facoltativo:_Indicizzazione_dei_file)
- [4Facoltativo: Accesso remoto alla shell (riga di comando)](#Facoltativo:_Accesso_remoto_alla_shell_.28riga_di_comando.29)
  - [4.1OpenRC](#OpenRC_3)
  - [4.2systemd](#systemd_3)
- [5Facoltativo: Completamento in shell](#Facoltativo:_Completamento_in_shell)
  - [5.1Bash](#Bash)
- [6Consigliato: Sincronizzazione dell'ora](#Consigliato:_Sincronizzazione_dell.27ora)
  - [6.1chrony](#chrony)
    - [6.1.1OpenRC](#OpenRC_4)
    - [6.1.2systemd](#systemd_4)
  - [6.2systemd-timesyncd](#systemd-timesyncd)
- [7Strumenti per i filesystem](#Strumenti_per_i_filesystem)

## Logger _(registratore degli eventi)_ di sistema

### OpenRC

Alcuni strumenti risultano assenti sull'archivio stage3, dato che numerosi pacchetti forniscono le stesse funzionalità. Spetta all'utente scegliere quali installare.

Il primo strumento su cui decidere è il meccanismo di logging _(registrazione degli eventi)_ di sistema. Unix e Linux hanno un'eccellente storia sulla capacità di annotazione degli eventi - se necessario, qualsiasi cosa succede sul sistema può essere registrato in un file di log _(registro degli eventi)_ di sistema.

Gentoo offre numerose utilità di registrazione di sistema. Alcune di queste includono:

- [sysklogd](/wiki/Sysklogd "Sysklogd") ([app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd)) \- Offre l'insieme tradizionale dei servizi di logging. La configurazione predefinita per il logging funziona bene così com'è, il che rende il pacchetto una buona opzione per i principianti.
- [syslog-ng](/wiki/Syslog-ng "Syslog-ng") ([app-admin/syslog-ng](https://packages.gentoo.org/packages/app-admin/syslog-ng)) \- Un logger di sistema avanzato. Richiede una configurazione aggiuntiva per qualsiasi cosa che va oltre il logging su un unico grande file. Gli utenti più esperti possono scegliere questo pacchetto per il suo potenziale di logging; si presti attenzione alla configurazione aggiuntiva, la quale è necessaria per qualsiasi tipo di logging intelligente.
- [metalog](/wiki/Metalog "Metalog") ([app-admin/metalog](https://packages.gentoo.org/packages/app-admin/metalog)) \- Un logger di sistema ampiamente configurabile.

Potrebbero esserci anche altri programmi di utilità per il logging di sistema disponibili attraverso il catalogo delle ebuild di Gentoo, poichè il numero di pacchetti disponibili aumenta giorno per giorno.

**Suggerimento**

Se si intende usare syslog-ng, si raccomanda di installare e configurare [logrotate](/wiki/Logrotate "Logrotate"). syslog-ng non fornisce alcun meccanismo di rotazione per i file di log. Comunque le versioni più recenti (>= 2.0) di sysklogd effettuano la propria rotazione del log.

Per installare il logger di sistema scelto, eseguire l'emerge. Su OpenRC, aggiungerlo al livello di esecuzione (runlevel) predefinito usando rc-update. Il seguente esempio installa e attiva [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) come programma di utilità per il syslog di sistema:

`root #` `emerge --ask app-admin/sysklogd`

`root #` `rc-update add sysklogd default`

### systemd

Mentre per i sistemi che fanno uso di OpenRC viene presentata una selezione di meccanismi di logging, systemd include un logger di sistema integrato chiamato servizio **systemd-journald**. Il servizio systemd-journald è capace di manipolare la maggior parte delle funzionalità di logging di sistema delineate nella precedente sezione logger di sistema. Ciò per dire, che la maggioranza delle installazioni che eseguono systemd come gestore del sistema e dei servizi può _serenamente saltare_ l'aggiunta di un programma di utilità per il syslog.

Vedere man journalctl per maggiori dettagli sull'uso di journalctl per richiedere ed esaminare i log di sistema.

Per diverse ragioni, come nel caso di inoltro dei log di sistema ad un computer centrale, può essere importante includere meccanismi di logging di sistema _ridondanti_ su un sistema che fa uso di systemd. Questa è una circostanza insolita per l'utenza tipica del manuale e considerata un caso d'uso avanzato. Perciò non è trattato dal manuale.

## Servizio cron

### OpenRC

Sebbene sia opzionale e non richiesto su tutti i sistemi, è saggio installare un servizio cron.

Un servizio cron esegue comandi a intervalli programmati. Gli intervalli potrebbero essere giornalieri, settimanali o mensili, una volta ogni martedí, una volta ogni due settimane etc. Un amministratore di sistema saggio regolerà il servizio cron per automatizzare le operazioni di manutenzione ripetitive del sistema.

Tutti i servizi cron supportano alti livelli di dettaglio per le operazioni pianificate, e di solito includono la capacità di inviare email o altre forme di notifica se un operazione pianificata non finisce come ci si aspetta.

Gentoo offre numerosi possibili servizi cron, tra cui:

- [cronie](/wiki/Cron/it#cronie "Cron/it") ([sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)) \- cronie si basa sul cron originale ed ha miglioramenti di sicurezza e configurazione come la capacità di usare PAM e SELinux.
- [dcron](/wiki/Cron/it#dcron_.28Dillon.27s_Cron.29 "Cron/it") ([sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron)) \- Questo servizio cron leggero ambisce ad essere semplice e sicuro, con funzioni appena sufficienti perché resti utile.
- [fcron](/wiki/Cron/it#fcron "Cron/it") ([sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron)) \- Un pianificatore di comandi con funzionalità ampliate rispetto a cron e anacron.
- [bcron](/wiki/Cron/it#bcron "Cron/it") ([sys-process/bcron](https://packages.gentoo.org/packages/sys-process/bcron)) \- Un sistema cron più giovane progettato con in testa operazioni sicure. Per fare ciò, il sistema si divide in svariati programmi separati, ognuno responsabile di una operazione separata, con comunicazioni tra le parti controllate severamente.

#### Predefinito: cronie

Il seguente esempio usa [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie) (da non confondere col servizio [NTP](/wiki/Network_Time_Protocol "Network Time Protocol") chiamato [chrony](/wiki/Chrony "Chrony")):

`root #` `emerge --ask sys-process/cronie`

Aggiungere cronie al runlevel di sistema predefinito, che lo avvierà automaticamente all'accensione:

`root #` `rc-update add cronie default`

#### In alternativa: dcron

`root #` `emerge --ask sys-process/dcron`

Se dcron è l'agente cron da fare andare avanti, deve essere eseguito un comando di inizializzazione aggiuntivo:

`root #` `crontab /etc/crontab`

#### In alternativa: fcron

`root #` `emerge --ask sys-process/fcron`

Se fcron è il gestore selezionato delle operazioni pianificate, un passaggio di emerge aggiuntivo è necessario:

`root #` `emerge --config sys-process/fcron`

#### In alternativa: bcron

bcron è un agente cron più giovane con separazione dei privilegi integrata.

`root #` `emerge --ask sys-process/bcron`

#### Utenti modprobed-db

Se l'uso di [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) è stato scelto come opzione per compilare manualmente il kernel.

Ora è il momento di impostare crontab per eseguire periodicamente la scansione del sistema alla ricerca dell'hardware utilizzato.

FILE **`/etc/crontab`** **Eseguire modprobed-db una volta ogni 6 ore**

```
0 */6 * * *     /usr/bin/modprobed-db store &> /dev/null

```

In una data successiva di almeno una settimana, si prega di visitare la sezione di compilazione del kernel dell'articolo [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fit "Modprobed-db") per completare l'impostazione.

### systemd

In modo simile al logging di sistema, i sistemi che fanno uso di systemd includono il supporto alle operazioni pianificate pronto all'uso in forma di _temporizzatori_. I temporizzatori systemd possono funzionare a livello di sistema o di utente e possono includere le stesse funzionalità che un servizio cron classico fornirebbe. A meno che non siano necessarie funzionalità ridondanti, installare un pianificatore di operazioni aggiuntivo come un servizio cron non è di solito necessario e può essere saltato tranquillamente.

#### Utenti modprobed-db

Se l'uso di [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) è stato scelto come opzione per compilare manualmente il kernel.

Ora è il momento di impostare un temporizzatore systemd per eseguire periodicamente la scansione del sistema alla ricerca dell'hardware utilizzato.

`root #` `systemctl --user enable modprobed-db`

In una data successiva di almeno una settimana, si prega di visitare la sezione di compilazione del kernel dell'articolo [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fit "Modprobed-db") per completare l'impostazione.

## Facoltativo: Indicizzazione dei file

Allo scopo di indicizzare il filesystem per avere capacità di localizzazione dei file più veloci, installare [mlocate](/wiki/Mlocate "Mlocate"):

`root #` `emerge --ask sys-apps/mlocate`

## Facoltativo: Accesso remoto alla shell (riga di comando)

**Suggerimento**

La configurazione predefinita di openssh non consente all'utente root di effettuare l'accesso da remoto. Si prega di [creare un utente non-root](/wiki/FAQ/it#How_do_I_add_a_normal_user.3F "FAQ/it") e configurarlo adeguatamente per consentirgli l'accesso dopo l'installazione se necessario o modificare /etc/ssh/sshd\_config per consentire l'accesso all'utente root.

Per avere la possibilità di accedere al sistema da remoto dopo l'installazione, sshd deve essere configurato per avviarsi all'accensione.

Per dettagli più approfonditi sulla configurazione di SSH, si faccia riferimento all'articolo [SSH](/wiki/SSH/it "SSH/it").

### OpenRC

Con OpenRC, per aggiungere lo script di init di sshd al runlevel predefinito:

`root #` `rc-update add sshd default`

Se è richiesto un accesso da console seriale (che è possibile in caso di server remoti), deve essere configurato agetty.

Togliere il commento alla sezione console seriale in /etc/inittab:

`root #` `nano -w /etc/inittab`

```
# SERIAL CONSOLES
s0:12345:respawn:/sbin/agetty 9600 ttyS0 vt100
s1:12345:respawn:/sbin/agetty 9600 ttyS1 vt100

```

### systemd

Per abilitare il server SSH, eseguire:

`root #` `systemctl enable sshd`

Per abilitare il supporto alla console seriale, eseguire:

`root #` `systemctl enable getty@tty1.service`

## Facoltativo: Completamento in shell

### Bash

Bash è la shell predefinita per i sistemi Gentoo, e quindi installare le estensioni di completamento può aiutare a gestire il sistema con efficacia e comodità. Il pacchetto [app-shells/bash-completion](https://packages.gentoo.org/packages/app-shells/bash-completion) installerà i completamenti disponibili per comandi di Gentoo specifici, così come per molti altri comandi comuni e programmi di utilità:

`root #` `emerge --ask app-shells/bash-completion`

Dopo l'installazione, il completamento in bash per comandi specifici può essere gestito attraverso eselect. Vedere la [sezione integrazioni al completamento in shell](/wiki/Bash#Shell_completion_integrations.2Fit "Bash") dell'articolo bash per maggiori dettagli.

## Consigliato: Sincronizzazione dell'ora

**Importante**

I sistemi senza un [Real-Time Clock (RTC)](/wiki/System_time#Software_clock_vs_Hardware_clock "System time") funzionante devono impostare l' [ora di sistema](/wiki/System_time "System time") ad ogni avvio del sistema, e in seguito a intervalli regolari.

E' importante usare metodi di sincronizzazione dell'orologio del sistema che fanno uso di server per l'orario su Internet. Ciò è spesso obbligatorio per sistemi senza RTC, ma è anche vantaggioso per sistemi che ce l'hanno, poichè la batteria potrebbe scaricarsi, e potrebbero accumularsi disallineamenti dell'orologio.

L'orologio solitamente è sincronizzato attraverso il [Network Time Protocol](/wiki/Network_Time_Protocol "Network Time Protocol"), con un implementazione come quella di [chrony](/wiki/Chrony "Chrony").

### chrony

Installare chrony (da non confondere col servizio cron chiamato [cronie](/wiki/Cron/it#cronie "Cron/it")):

`root #` `emerge --ask net-misc/chrony`

Vedere l'articolo [chrony](/wiki/Chrony "Chrony") per ulteriori informazioni, ad esempio se fossero necessarie configurazioni più avanzate.

#### OpenRC

Aggiungere chronyd al runlevel predefinito per avviarlo all'accensione, in modo che l'ora venga sincronizzata automaticamente:

`root #` `rc-update add chronyd default`

#### systemd

Avviare chronyd all'accensione, per avere l'ora sincronizzata automaticamente:

`root #` `systemctl enable chronyd.service`

### systemd-timesyncd

In alternativa, gli utenti di systemd possono voler usare il più semplice client SNTP systemd-timesyncd che viene installato in modo predefinito e abilitato con:

`root #` `systemctl enable systemd-timesyncd.service`

## Strumenti per i filesystem

In base al filesystem usato, si può aver bisogno di installare i programmi di utilità del filesystem necessari (per controllare l'integrità dei filesystem, (ri)formattare filesystem, ecc.). Si noti che gli strumenti dello spazio utente per ext4 ([sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)) sono già installati come parte dell' [insieme @system](/wiki/System_set_(Portage) "System set (Portage)").

La seguente tabella elenca gli strumenti da installare se un certo strumento del filesystem fosse richiesto nell'ambiente installato.

Filesystem
Package
[XFS](/wiki/XFS/it "XFS/it")[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/it "Ext4/it")[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[Btrfs](/wiki/Btrfs/it "Btrfs/it")[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS/it "F2FS/it")[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

Si consiglia di installare il pacchetto [sys-block/io-scheduler-udev-rules](https://packages.gentoo.org/packages/sys-block/io-scheduler-udev-rules) per un comportamento corretto del pianificatore con p.es. i dispositivi nvme:

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

**Suggerimento**

Per maggiori informazioni sui file system in Gentoo vedere l' [articolo sui file system](/wiki/Filesystem "Filesystem").

Ora si prosegua con [Configurare il programma d'avvio](/wiki/Handbook:X86/Installation/Bootloader/it "Handbook:X86/Installation/Bootloader/it").

[← Configurare il sistema](/wiki/Handbook:X86/Installation/System/it "Previous part") [Home](/wiki/Handbook:X86 "Handbook:X86") [Configurare programma d'avvio →](/wiki/Handbook:X86/Installation/Bootloader/it "Next part")