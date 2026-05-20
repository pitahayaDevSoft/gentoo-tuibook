# Stage

Other languages:

- Deutsch
- [English](/wiki/Handbook:X86/Installation/Stage "Handbook:X86/Installation/Stage (100% translated)")
- [Türkçe](/wiki/Handbook:X86/Installation/Stage/tr "Handbook:X86/Installation/Stage/tr (0% translated)")
- [español](/wiki/Handbook:X86/Installation/Stage/es "Manual de Gentoo: X86/Instalación/Stage (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Stage/fr "Handbook:X86/Installation/Stage/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Stage/it "Handbook:X86/Installation/Stage (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/Stage/hu "Handbook:X86/Installation/Stage/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Stage/pl "Handbook:X86/Installation/Stage (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Stage/pt-br "Manual:X86/Instalação/Stage (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Stage/cs "Handbook:X86/Installation/Stage/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Stage/ru "Handbook:X86/Installation/Stage (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Stage/ta "கையேடு:X86/நிறுவல்/நிலை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Stage/zh-cn "手册：X86/安装/安装stage3 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Stage/ja "ハンドブック:X86/インストール/Stage (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Stage/ko "Handbook:X86/Installation/Stage/ko (100% translated)")

[X86 Handbuch](/wiki/Handbook:X86/de "Handbook:X86/de")[Installation](/wiki/Handbook:X86/Full/Installation/de "Handbook:X86/Full/Installation/de")[Über die Installation](/wiki/Handbook:X86/Installation/About/de "Handbook:X86/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:X86/Installation/Media/de "Handbook:X86/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:X86/Installation/Networking/de "Handbook:X86/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:X86/Installation/Disks/de "Handbook:X86/Installation/Disks/de")Installation des Stage Archivs[Installation des Basissystems](/wiki/Handbook:X86/Installation/Base/de "Handbook:X86/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:X86/Installation/Kernel/de "Handbook:X86/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:X86/Installation/System/de "Handbook:X86/Installation/System/de")[Installation der Tools](/wiki/Handbook:X86/Installation/Tools/de "Handbook:X86/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:X86/Installation/Bootloader/de "Handbook:X86/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:X86/Installation/Finalizing/de "Handbook:X86/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:X86/Full/Working/de "Handbook:X86/Full/Working/de")[Portage-Einführung](/wiki/Handbook:X86/Working/Portage/de "Handbook:X86/Working/Portage/de")[USE-Flags](/wiki/Handbook:X86/Working/USE/de "Handbook:X86/Working/USE/de")[Portage-Features](/wiki/Handbook:X86/Working/Features/de "Handbook:X86/Working/Features/de")[Initskript-System](/wiki/Handbook:X86/Working/Initscripts/de "Handbook:X86/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:X86/Working/EnvVar/de "Handbook:X86/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:X86/Full/Portage/de "Handbook:X86/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:X86/Portage/Files/de "Handbook:X86/Portage/Files/de")[Variablen](/wiki/Handbook:X86/Portage/Variables/de "Handbook:X86/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:X86/Portage/Branches/de "Handbook:X86/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:X86/Portage/Tools/de "Handbook:X86/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:X86/Portage/CustomTree/de "Handbook:X86/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:X86/Portage/Advanced/de "Handbook:X86/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:X86/Full/Networking/de "Handbook:X86/Full/Networking/de")[Zu Beginn](/wiki/Handbook:X86/Networking/Introduction/de "Handbook:X86/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:X86/Networking/Advanced/de "Handbook:X86/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:X86/Networking/Modular/de "Handbook:X86/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:X86/Networking/Wireless/de "Handbook:X86/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:X86/Networking/Extending/de "Handbook:X86/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:X86/Networking/Dynamic/de "Handbook:X86/Networking/Dynamic/de")

## Contents

- [1Auswahl eines Stage Tar-Archivs](#Auswahl_eines_Stage_Tar-Archivs)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
  - [1.3Multilib (32 and 64-bit)](#Multilib_.2832_and_64-bit.29)
  - [1.4No-multilib (nur 64-bit)](#No-multilib_.28nur_64-bit.29)
- [2Stage Tar-Archiv herunterladen](#Stage_Tar-Archiv_herunterladen)
- [3Datum und Uhrzeit einstellen](#Datum_und_Uhrzeit_einstellen)
  - [3.1Automatisch](#Automatisch)
  - [3.2Manuell](#Manuell)
  - [3.3Browser mit grafischer Benutzeroberfläche](#Browser_mit_grafischer_Benutzeroberfl.C3.A4che)
  - [3.4Textbasierte Browser](#Textbasierte_Browser)
  - [3.5Überprüfung und Validierung](#.C3.9Cberpr.C3.BCfung_und_Validierung)
- [4Installation eines Stage Tar-Archivs](#Installation_eines_Stage_Tar-Archivs)
- [5Compiler-Optionen konfigurieren](#Compiler-Optionen_konfigurieren)
  - [5.1Einleitung](#Einleitung)
  - [5.2CFLAGS und CXXFLAGS](#CFLAGS_und_CXXFLAGS)
  - [5.3RUSTFLAGS](#RUSTFLAGS)
  - [5.4MAKEOPTS](#MAKEOPTS)
  - [5.5Los geht's!](#Los_geht.27s.21)
- [6Referenzen](#Referenzen)

### Auswahl eines Stage Tar-Archivs

**Tipp**

On supported architectures, it is recommended for users targeting a desktop (graphical) operating system environment to use a stage file with the term `desktop` within the name. These files include packages such as [llvm-core/llvm](https://packages.gentoo.org/packages/llvm-core/llvm) and [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) and USE flag tuning which will greatly improve install time.

The [stage file](/wiki/Stage_file "Stage file") acts as the seed of a Gentoo install. Stage files are generated by the [Release Engineering Team](/wiki/Project:RelEng "Project:RelEng") with [Catalyst](/wiki/Catalyst "Catalyst"). Stage files are based on specific [profiles](/wiki/Profile_(Portage) "Profile (Portage)"), and contain an almost-complete system.

When choosing a stage file, it's important to pick one with profile targets corresponding to the desired system type.

**Hinweis**

Bei einem bestehenden Gentoo System ist es technisch möglich, zwischen OpenRC und systemd zu wechseln. Solche Wechsel sind jedoch aufwändig und können nicht im Rahmen dieses Installations-Handbuchs beschrieben werden. Bevor Sie einen Stage-Tarball herunterladen, entscheiden Sie, ob OpenRC oder systemd als Ziel-Init-System verwendet werden soll und laden Sie den entsprechenden Stage-Tarball herunter.

Die meisten Anwender sollten die "advanced" Tar-Archiv Optionen NICHT verwenden. Sie sind für spezielle Software- oder Hardware-Konfigurationen gedacht.

#### OpenRC

[OpenRC](/wiki/OpenRC/de "OpenRC/de") ist ein Abhängigkeits-basiertes Init-System. Nachdem der Kernel gebootet hat, ist es ist zuständig für das Starten von System-Diensten. OpenRC ist kompatibel mit dem vom System bereitgestellten Init-Programm, das normalerweise unter /sbin/init installiert ist. OpenRC wurde unter und für Gentoo entwickelt, aber es wird auch bei einigen anderen Linux-Distributionen und BSD-Systemen verwendet.

#### systemd

systemd ist ein moderner SysV-style Init- und rc-Ersatz für Linux-Systeme. Mittlerweile wird es bei der Mehrzahl der Linux-Distibutionen als primäres Init-System verwendet. systemd wird von Gentoo vollständig unterstützt und funktioniert führ den vorgesehenen Zweck. Wenn etwas im Handbuch für einen systemd-Installationspfad zu fehlen scheint lesen Sie den [systemd Artikel](/wiki/Systemd/de "Systemd/de") _bevor_ Sie um Unterstützung bitten.

#### Multilib (32 and 64-bit)

**Hinweis**

Nicht jede Architektur hat eine Multilib-Option. Viele laufen nur mit nativem Code. Multilib wird am häufigsten auf **amd64** angewendet.

Die Wahl des richtigen Stage Tar-Archivs kann später im Installationsprozess erhebliche Mengen an Zeit einsparen, ganz besonders wenn der Zeitpunkt gekommen ist, für die [Auswahl des System-Profils](/wiki/Handbook:X86/Installation/Base/de#Choosing_the_right_profile "Handbook:X86/Installation/Base/de"). Ein "multilib" Stage Tar-Archiv ermöglicht ein System mit 64- und 32-Bit Bibliotheken, wobei nach Möglichkeit die 64-Bit Bibliotheken verwendet werden. Falls dies nicht möglich sein sollte, können die 32-Bit Bibliotheken verwendet werden. Für die meisten Installationen ist dies eine hervorragende Wahl, weil sie große Flexibilität und Anpassungsmöglichkeiten für die Zukunft ermöglicht. Auch wer in der Lage sein möchte, einfach zwischen verschiedenen Profilen zu wechseln, sollte ein "multilib" Stage Tar-Archiv wählen.

**Tipp**

Using `multilib` targets makes it easier to switch profiles later, compared to `no-multilib`

#### No-multilib (nur 64-bit)

**Warnung**

Leser, die gerade erst mit Gentoo anfangen, sollten _nicht_ einen no-multilib Tarball wählen, es sei denn, es ist absolut notwendig. Die Migration von einem "no-multilib" zu einem "multilib" System ist sehr schwierig und erfordert sehr gute Kenntnisse von Gentoo und der Low-Level Toolchain. Sogar für die Gentoo [Toolchain Entwickler](/wiki/Project:Toolchain "Project:Toolchain") ist ein solcher Wechsel nicht ganz einfach. Anders ausgedrückt: gewöhnliche Anwender, die sich für
"no-multilib" entscheiden, können nur durch eine Neu-Installation auf "multilib" wechseln.

Die Wahl eines "no-multilib" Stage Tar-Archives ermöglicht die Installation einer reinen 64-Bit Linux-Umgebung. Bitte beachten Sie, dass einige Anwendungen wie Wine, die eine 32-Bit Umgebung benötigen, dann nicht laufen werden. Ein späterer Wechsel auf eine "multilib" Umgebung ist schwierig, jedoch nicht unmöglich.

### Stage Tar-Archiv herunterladen

Before downloading the _stage file_, the current directory should be set to the location of the mount used for the install:

`root #` `cd /mnt/gentoo`

### Datum und Uhrzeit einstellen

Stage archives are generally obtained using HTTPS which requires relatively accurate system time. Clock skew can prevent downloads from working, and can cause unpredictable errors if the system time is adjusted by any considerable amount after installation.

Überprüfen Sie das System-Datum und die System-Uhrzeit mit dem Kommando date:

`root #` `date`

```
Mo 3. Okt 13:16:22 CET 2022

```

Wenn Datum/Uhrzeit um mehr als ein paar Minuten abweichen, sollte sie mit einer der folgenden Methoden aktualisiert werden.

#### Automatisch

Using [Network\_Time\_Protocol](/wiki/Network_Time_Protocol "Network Time Protocol") to correct clock skew is typically easier and more reliable than manually setting the system clock.

chronyd, part of [net-misc/chrony](https://packages.gentoo.org/packages/net-misc/chrony) can be used to update the system clock to UTC with:

`root #` `chronyd -q`

**Wichtig**

Systems without a functioning Real-Time Clock (RTC) must sync the system clock at every system start, and on regular intervals thereafter. This is also beneficial for systems with a RTC, as the battery could fail, and clock skew can accumulate.

**Warnung**

Standard NTP traffic is not authenticated, it is important to verify time data obtained from the network.

#### Manuell

When NTP access is unavailable, date can be used to manually set the system clock.

Die UTC-Zeit wird für alle Linux-Systeme empfohlen. Eine Zeitzone wird später bei der Installation definiert, wodurch die Uhr die lokale Zeit anzeigt.

Für Systeme, die keinen Zugang zu einem Zeitserver haben, kann auch der Befehl date verwendet werden, um die Systemuhr zu stellen. Dabei wird das folgende Format als Argument verwendet: `MMDDhhmmYYYY` syntax ( **M** onth, **D** ay, **h** our, **m** inute and **Y** ear).

Um beispielsweise das Datum auf den 27. November 2022, 13:39 Uhr einzustellen, geben Sie folgendes ein:

`root #` `date 112713392022`

#### Browser mit grafischer Benutzeroberfläche

Wenn Sie einen Web-Browser mit grafischer Benutzeroberfläche verwenden: gehen Sie auf die [Download Seite](https://www.gentoo.org/downloads/#other-arches) und kopieren Sie die URL des gewünschten Stage Tar-Archivs in die Zwischenablage (durch Drücken der rechten Maus-Taste und dann "Copy Link"). Gehen Sie dann in Ihr Terminal-Fenster, tippen Sie wget und kopieren Sie die URL aus der Zwischenablage. Drücken Sie `Return`, um den Download zu starten.

`root #` `wget <PASTED_STAGE_URL>`

#### Textbasierte Browser

Wenn Sie lieber in einem Terminal-Fenster arbeiten, können Sie links verwenden, einen textbasierten, menügeführten Browser. Starten Sie links[www-client/links](https://packages.gentoo.org/packages/www-client/links) und navigieren Sie zu der Gentoo Mirror-Seite:

`root #` `links https://www.gentoo.org/downloads/mirrors/`

Um einen HTTP-Proxy mit links zu verwenden, übergeben Sie die URL mit der `-http-proxy` Option:

`root #` `links -http-proxy proxy.server.com:8080 https://www.gentoo.org/downloads/mirrors/`

Neben links gibt es auch den Browser lynx[www-client/lynx](https://packages.gentoo.org/packages/www-client/lynx). Wie links ist es ein nicht-grafischer Browser, aber er ist nicht menügesteuert.

`root #` `lynx https://www.gentoo.org/downloads/mirrors/`

Wenn ein Proxy definiert werden muss, exportieren Sie die http\_proxy und/ oder ftp\_proxy Variablen:

`root #` `export http_proxy="http://proxy.server.com:port"
`

`root #` `export ftp_proxy="http://proxy.server.com:port"`

Bitte wählen Sie in der Spiegel-Liste einen Spiegel in Ihrer Nähe. Für gewöhnlich genügen HTTP Spiegel, andere Protokolle stehen aber auch zur Verfügung. Gehen Sie in das Verzeichnis releases/x86/autobuilds/. Dort werden alle verfügbaren Stage Tar-Archive angezeigt (sie können in Unterverzeichnissen gespeichert sein, benannt nach den einzelnen Sub-Architekturen). Wählen Sie eines aus und drücken Sie `d` zum Download.

Nachdem Sie das Stage Tar-Archiv erfolgreich heruntergeladen haben, können Sie die Integrität des Tar-Archivs verifizieren und den Inhalt validieren. Wie das geht, steht im [folgenden Abschnitt](/wiki/Handbook:X86/Installation/Stage/de#Verifying_and_validating "Handbook:X86/Installation/Stage/de").

Wenn Sie kein Interesse an einer Überprüfung und Validierung des Stage Tar-Archivs haben, können Sie jetzt `q` drücken, um den Browser zu beenden. Springen Sie danach zu dem Abschnitt [Stage Tar-Archiv entpacken](/wiki/Handbook:X86/Installation/Stage/de#Unpacking_the_stage_tarball "Handbook:X86/Installation/Stage/de").

#### Überprüfung und Validierung

Wie bei der minimalen Installations-CDs stehen zusätzliche Downloads zur Verfügung, mit denen das Stage Tar-Archiv überprüft und validiert werden kann. Obwohl dieser Schritt übersprungen werden kann, können diese Downloads von Anwendern genutzt werden, die die Integrität des Stage Tar-Archivs sicherstellen wollen.

`root #` `wget https://distfiles.gentoo.org/releases/`

- Eine Datei .CONTENTS, die eine Liste aller Dateien im Stage Tar-Archiv enthält.
- Eine Datei .DIGESTS, die Prüfsummen des Stage Tar-Archivs von verschiedenen Algorithmen beinhaltet.

Genau wie bei der ISO-Datei kann die kryptografische Signatur der Datei tar.xz mit gpg überprüft werden, um sicherzustellen, dass keine Manipulationen am Tarball vorgenommen wurden.

For official Gentoo live images, the [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release) package provides PGP signing keys for automated releases. The keys must first be imported into the user's session in order to be used for verification:

`root #` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

For all non-official live images which offer gpg and wget in the live environment, refer to the earlier [Verifying the downloaded files](/wiki/Handbook:X86/Installation/Media/de#Verifying_the_downloaded_files "Handbook:X86/Installation/Media/de") section.

Verify the signature of the tarball and, optionally, associated checksum files:

`root #` `gpg --verify stage3-x86-<release>-<init>.tar.xz.asc
`

`root #` `gpg --verify stage3-x86-<release>-<init>.tar.xz.DIGEST
`

`root #` `gpg --verify stage3-x86-<release>-<init>.tar.xz.sha256
`

If verification succeeds, "Good signature from" will be in the output of the previous command(s).

Die Fingerprints der OpenPGP Schlüssel, die zum Signieren der Release Medien verwendet werden, finden Sie auf der [Release Media Signatures Seite](https://www.gentoo.org/downloads/signatures/).

**Hinweis**

Die meisten Stages sind jetzt explizit [mit dem Suffix](https://www.gentoo.org/news/2021/07/20/more-downloads.html) des Init-Systemtyps (openrc oder systemd), obwohl diese bei einigen Architekturen noch fehlen können.

Verwenden Sie openssl zum Berechnen einer Prüfsumme des Stage tar-Archivs und vergleichen Sie die Ausgabe mit den Prüfsummen, die in der Datei .DIGESTS steht.

Um die SHA512-Prüfsumme mit openssl zu überprüfen:

`root #` `openssl dgst -r -sha512 stage3-x86-<release>-<init>.tar.xz`

`dgst` instructs the openssl command to use the Message Digest sub-command, `-r` prints the digest output in coreutils format, and `-sha512` selects the SHA512 digest.

Um die BLAKE2B512-Prüfsumme mit openssl zu überprüfen:

`root #` `openssl dgst -r -blake2b512 stage3-x86-<release>-<init>.tar.xz`

Vergleichen Sie die Ausgabe dieser Befehle mit dem Wert der in den .DIGESTS Dateien eingetragen ist. Die Werte müssen übereinstimmen, andernfalls ist möglicherweise die heruntergeladene Datei beschädigt (oder die DIGEST-Datei ist es).

Um den SHA256-Hash aus einer verknüpften .sha256-Datei mit dem Dienstprogramm sha256sum zu überprüfen:

`root #` `sha256sum --check stage3-x86-<release>-<init>.tar.xz.sha256`

The `--check` option instructs sha256sum to read a list of expected files and associated hashes, and then print an associated "OK" for each file that calculates correctly or a "FAILED" for files that do not.

## Installation eines Stage Tar-Archivs

Als Nächstes entpacken Sie die heruntergeladene Stufe auf das System. Verwenden Sie das Dienstprogramm tar, um wie folgt fortzufahren:

`root #` `tar xpvf stage3-*.tar.xz --xattrs-include='*.*' --numeric-owner`

Verifizieren Sie, dass Sie genau die oben angegebenen Optionen ( `xvpf`, `--xattrs-include='*.*'` und `--numeric-owner`) im Befehl verwenden. Das `x` steht für e **x** trahieren, das `p` ( **p** reserve) für den Erhalt der Dateirechte und das `f` ( **f** ile) gibt an, dass wir das auszupackende Archiv aus einer Datei lesen wollen - und nicht von der Standardeingabe. `--xattrs-include='*.*'` bedeutet, dass die erweiterten (extended) Attribute erhalten bleiben sollen. `--numeric-owner` ist erforderlich um sicherzustellen, dass die User- und Gruppen IDs der extrahierten Dateien so gesetzt werden, wie vom Gentoo Release Team definiert (und zwar auch dann, wenn abenteuerlustige Anwender bei der Installation nicht die offiziellen Live-Umgebungen verwenden).

- `x` e **x** tract, instructs tar to extract the contents of the archive.
- `p` **p** reserve permissions.
- `v` **v** erbose output.
- `f` **f** ile, provides tar with the name of the input archive.
- `--xattrs-include='*.*'` Preserves extended attributes in all namespaces stored in the archive.
- `--numeric-owner` Ensure that the user and group IDs of files being extracted from the tarball remain the same as Gentoo's release engineering team intended (even if adventurous users are not using official Gentoo live environments for the installation process).
- `-C /mnt/gentoo` Extract files to the root partition regardless of the current directory.

Nachdem nun das Stage Tar-Archiv ausgepackt ist, geht es weiter mit dem Schritt: [Compiler-Optionen konfigurieren](/wiki/Handbook:X86/Installation/Stage/de#Configuring_compile_options "Handbook:X86/Installation/Stage/de").

## Compiler-Optionen konfigurieren

### Einleitung

Um das System zu optimieren, können Variablen gesetzt werden, mit denen das Verhalten von Portage (Gentoos offiziellem Paket-Manager) beeinflusst wird. Diese Variablen können als Umgebungs-Variablen gesetzt werden (mit export), aber diese wären nicht permanent.

**Hinweis**

Technisch gesehen können Variablen über die Profil- oder rc-Dateien der [Shell](/wiki/Shell "Shell") exportiert werden, was jedoch für die grundlegende Systemadministration nicht die beste Praxis ist.

Portage liest die Datei [make.conf](/wiki/Make.conf "Make.conf") ein, wenn es läuft, was das Laufzeitverhalten abhängig von den in der Datei gespeicherten Werten verändert. Die Datei make.conf kann als die primäre Konfigurationsdatei für Portage angesehen werden, also behandeln Sie ihren Inhalt sorgfältig.

**Tipp**

Eine kommentiere Auflistung aller Variablen kann unter /mnt/gentoo/usr/share/portage/make.conf.example gefunden werden. Zusätzliche Dokumentation zu make.conf finden Sie, indem Sie man 5 make.conf ausführen.

Für eine erfolgreiche Gentoo-Installation müssen nur die unten genannten Variablen gesetzt werden.

Starten Sie einen Editor, damit Sie die Werte der Optimierungs-Variablen, die wir im Folgenden besprechen werden, ändern können. In dieser Anleitung verwenden wir den Editor nano.

`root #` `nano /mnt/gentoo/etc/portage/make.conf`

Anhand der Datei make.conf.example sollte die Syntax von make.conf erkennbar sein: Kommentar-Zeilen starten mit einem `#`, andere Zeilen nutzen die `VARIABLE="Wert"` Syntax. Im nächsten Abschnitt werden wir einige dieser Variablen besprechen.

### CFLAGS und CXXFLAGS

In den Variablen CFLAGS und CXXFLAGS
können Optimierungs-Optionen für den GCC C-Compiler und den GCC C++ Compiler definiert werden. Die in make.conf global definierten Optimierungs-Optionen gelten dann für die Installation aller Pakete. Um die maximal mögliche Performance jedes einzelnen Pakets zu erreichen, bräuchte man jedoch für jedes Programm andere Optionen - weil jedes Programm anders ist und anders optimiert werden muss. Dies ist jedoch nicht praktikabel - und deshalb werden die Optimierungs-Optionen global für alle Pakete in make.conf definiert.

In make.conf sollten Optimierungen definiert werden, mit denen ihr System schnell und stabil läuft. Definieren Sie hier keine experimentellen Werte. Zu viel Optimierung kann dazu führen, dass Programme nicht mehr gut laufen: sie stürzen ab, funktionieren nicht richtig, oder - noch schlimmer - berechnen falsche Ergebnisse.

In diesem Abschnitt werden wir nur die wichtigsten Optimierungs-Optionen erklären. Eine Übersicht über alle möglichen Optimierungs-Optionen finden Sie auf der [GCC online documentation](https://gcc.gnu.org/onlinedocs/), auf der GCC man page (man gcc) und
auf der GCC info Seite (info gcc). man und info sind jedoch nur auf einem fertig installierten Linux System verfügbar. Weiterhin enthält die Datei make.conf.example viele Beispiele und Informationen - bitte vergessen Sie nicht, sie zu lesen.

Eine wichtige Einstellung ist die
`-march=` or `-mtune=` Option, mit der der Name der Ziel-Architektur definiert wird. Mögliche Werte werden in der Datei make.conf.example beschrieben (als Kommentare). Ein häufig benutzter Wert ist _native_. Mit diesem Wert wählt der Compiler die Ziel-Architektur des Systems, auf dem er gerade läuft (also des Systems, das Sie gerade installieren).

Eine weitere wichtige Option ist `-O` (ein großer Buchstabe O, keine Null), mit der die GCC Optimierungs-Klasse definiert wird. Mögliche Werte sind s ( **s** ize, optimiert auf kleine Code-Größe), 0 (Null - keine Optimierungen), 1, 2 oder sogar 3 für Geschwindigkeits-Optimierung (wobei jede Klasse die Optimierungen der vorhergehenden Klasse und einige zusätzliche Optimierungen durchführt). `-O2` ist der empfohlene Standard-Wert. Von `-O3` ist bekannt, dass es Probleme geben wird, wenn es systemweit benutzt wird. Wir empfehlen deshalb, bei `-O2` zu bleiben.

Eine weitere häufig genutzte Option ist `-pipe` (verwende für die Kommunikation zwischen den verschiedenen Compiler-Stufen Pipes statt temporärer Dateien). Diese Option hat keine Auswirkungen auf den generierten Code, benötigt aber mehr Arbeitsspeicher. Auf Systemen mit wenig Arbeitsspeicher kann dies dazu führen, dass GCC vorzeitig abgeschossen wird. Wenn das passieren sollte, verwenden Sie diese Option nicht.

Die Verwendung von `-fomit-frame-pointer` (die dazu führt, dass der Frame Pointer in Funktionen, die keinen Frame Pointer benötigen, nicht gesetzt wird) kann erhebliche Auswirkungen auf das Debugging von Programmen haben.

Wenn Sie die CFLAGS oder die CXXFLAGS Variable definieren, sollten Sie die verschiedenen Optimierungs-Optionen in einem Wert kombinieren. Die Standard-Werte in dem von Ihnen ausgepackten Stage Tar-Archiv sollten ausreichend sein. In der folgenden Box zeigen wir ein Beispiel:

CODE **Beispiel für CFLAGS und CXXFLAGS Variablen**

```
# Compiler flags to set for all languages
COMMON_FLAGS="-O2 -march=i686 -pipe"
# Use the same settings for both variables
CFLAGS="${COMMON_FLAGS}"
CXXFLAGS="${COMMON_FLAGS}"

```

**Tipp**

Obwohl der Artikel [GCC-Optimierung](/wiki/GCC_optimization "GCC optimization") mehr Informationen darüber enthält, wie sich die verschiedenen Kompilierungsoptionen auf ein System auswirken können, ist der Artikel [Sichere CFLAGS](/wiki/Safe_CFLAGS "Safe CFLAGS") für Anfänger ein praktischerer Ort, um mit der Optimierung ihres Systems zu beginnen.

### RUSTFLAGS

Many programs are now written in Rust which has its own way of optimising. By default Rust optimizes to level 3 on all release builds unless a project changes it so this should be left as is. A full optimization list (known as codegen) that can be passed to the rust compiler is available at [https://doc.rust-lang.org/rustc/codegen-options/index.html](https://doc.rust-lang.org/rustc/codegen-options/index.html).

The most useful optimization would be to tell Rust to compile for your CPU using the following example:

DATEI **`/etc/portage/make.conf`** **RUSTFLAGS Example**

```
RUSTFLAGS="${RUSTFLAGS} -C target-cpu=native"

```

To get a list of supported CPUs in Rust run:

`root #` `rustc -C target-cpu=help`

This will show what the default and also which CPU will be selected with native.

**Hinweis**

The above command only works on desktop stage3 tarballs or after emerging [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) or [dev-lang/rust](https://packages.gentoo.org/packages/dev-lang/rust).

### MAKEOPTS

Die MAKEOPTS Variable definiert, wie viele parallele Kompilierungen bei der Installation eines Pakets stattfinden sollen. Ab Portage Version 3.0.31[\[1\]](#cite_note-1) ist das Standardverhalten von Portage, wenn ist undefiniert gelassen wird, den MAKEOPTS Wert auf die gleiche Anzahl von Threads zu setzten, die von nproc zurückgegeben wird.

Further, as of Portage 3.0.53[\[2\]](#cite_note-2), if left undefined, Portage's default behavior is to set the MAKEOPTS load-average value to the same number of threads returned by nproc.

Eine gute Wahl ist der kleinere der folgenden Werte: die Anzahl der Threads der CPU oder die Gesamtmenge des System-RAM geteilt durch 2 GiB.

**Warnung**

Eine große Anzahl von Jobs benötigt auch entsprechenden Hauptspeicher (RAM). Eine guter Schätzwert sind mindestens 2 GiB RAM pro Job ( `-j6` benötigt also _mindestens_ 12 GiB RAM). Um zu verhindern, dass mehr RAM benötigt wird als vorhanden ist, sollten Sie die Anzahl der Jobs ggf. reduzieren und an das vorhandene RAM anpassen.

**Tipp**

Wenn mehrere Emerge-Jobs parallel laufen gelassen werden
( `--jobs`), kann die Anzahl der resultierenden Jobs exponentiell ansteigen (bis zum Produkt von "make jobs" multipliziert mit "emerge jobs"). Dieses Wachstum kann kontrolliert werden durch den Einsatz einer lokalen distcc Konfiguration, die die Anzahl der Compiler-Aufrufe pro Host begrenzt.

DATEI **`/etc/portage/make.conf`** **Beispiel MAKEOPTS Deklaration**

```
# Wenn nicht definiert, setzt Portage standardmäßig den MAKEOPTS Wert auf die gleiche Anzahl von Threads die von `nproc` zurückgegeben wird.
MAKEOPTS="-j4"

```

Suchen Sie nach MAKEOPTS in man 5 make.conf für weitere Informationen.

### Los geht's!

Aktualisieren Sie die Datei /mnt/gentoo/etc/portage/make.conf nach Ihren Wünschen und speichern Sie sie (nano Benutzer drücken `Ctrl` + `o`, um Änderung zu schreiben und dann `Ctrl` + `x`, um zu beenden).

## Referenzen

1. [↑](#cite_ref-1)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2](https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2)
2. [↑](#cite_ref-2)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e](https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e)

[← Vorbereiten der Festplatte(n)](/wiki/Handbook:X86/Installation/Disks/de "Previous part") [Anfang](/wiki/Handbook:X86/de "Handbook:X86/de") [Installation des Basissystems →](/wiki/Handbook:X86/Installation/Base/de "Next part")