# Media

Other languages:

- Deutsch
- [English](/wiki/Handbook:MIPS/Installation/Media "Handbook:MIPS/Installation/Media (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Media/es "Handbook:MIPS/Installation/Media/es (57% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Media/fr "Handbook:MIPS/Installation/Media/fr (57% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Media/it "Handbook:MIPS/Installation/Media/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Media/hu "Handbook:MIPS/Installation/Media/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Media/pl "Handbook:MIPS/Installation/Media/pl (14% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Media/pt-br "Handbook:MIPS/Installation/Media/pt-br (57% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Media/ru "Handbook:MIPS/Installation/Media/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Media/ta "Handbook:MIPS/Installation/Media/ta (57% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Media/zh-cn "Handbook:MIPS/Installation/Media/zh-cn (57% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Media/ja "Handbook:MIPS/Installation/Media/ja (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Media/ko "Handbook:MIPS/Installation/Media/ko (57% translated)")

[MIPS Handbuch](/wiki/Handbook:MIPS/de "Handbook:MIPS/de")[Installation](/wiki/Handbook:MIPS/Full/Installation/de "Handbook:MIPS/Full/Installation/de")[Über die Installation](/wiki/Handbook:MIPS/Installation/About/de "Handbook:MIPS/Installation/About/de")Auswahl des Mediums[Konfiguration des Netzwerks](/wiki/Handbook:MIPS/Installation/Networking/de "Handbook:MIPS/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:MIPS/Installation/Disks/de "Handbook:MIPS/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:MIPS/Installation/Stage/de "Handbook:MIPS/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:MIPS/Installation/Base/de "Handbook:MIPS/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:MIPS/Installation/Kernel/de "Handbook:MIPS/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:MIPS/Installation/System/de "Handbook:MIPS/Installation/System/de")[Installation der Tools](/wiki/Handbook:MIPS/Installation/Tools/de "Handbook:MIPS/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:MIPS/Installation/Bootloader/de "Handbook:MIPS/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:MIPS/Installation/Finalizing/de "Handbook:MIPS/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:MIPS/Full/Working/de "Handbook:MIPS/Full/Working/de")[Portage-Einführung](/wiki/Handbook:MIPS/Working/Portage/de "Handbook:MIPS/Working/Portage/de")[USE-Flags](/wiki/Handbook:MIPS/Working/USE/de "Handbook:MIPS/Working/USE/de")[Portage-Features](/wiki/Handbook:MIPS/Working/Features/de "Handbook:MIPS/Working/Features/de")[Initskript-System](/wiki/Handbook:MIPS/Working/Initscripts/de "Handbook:MIPS/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:MIPS/Working/EnvVar/de "Handbook:MIPS/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:MIPS/Full/Portage/de "Handbook:MIPS/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:MIPS/Portage/Files/de "Handbook:MIPS/Portage/Files/de")[Variablen](/wiki/Handbook:MIPS/Portage/Variables/de "Handbook:MIPS/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:MIPS/Portage/Branches/de "Handbook:MIPS/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:MIPS/Portage/Tools/de "Handbook:MIPS/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:MIPS/Portage/CustomTree/de "Handbook:MIPS/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:MIPS/Portage/Advanced/de "Handbook:MIPS/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:MIPS/Full/Networking/de "Handbook:MIPS/Full/Networking/de")[Zu Beginn](/wiki/Handbook:MIPS/Networking/Introduction/de "Handbook:MIPS/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:MIPS/Networking/Advanced/de "Handbook:MIPS/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:MIPS/Networking/Modular/de "Handbook:MIPS/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:MIPS/Networking/Wireless/de "Handbook:MIPS/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:MIPS/Networking/Extending/de "Handbook:MIPS/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:MIPS/Networking/Dynamic/de "Handbook:MIPS/Networking/Dynamic/de")

## Contents

- [1Hardwareanforderungen](#Hardwareanforderungen)
- [2Bemerkungen zur Installation](#Bemerkungen_zur_Installation)
- [3Netboot Übersicht](#Netboot_.C3.9Cbersicht)
  - [3.1TFTP und DHCP einrichten](#TFTP_und_DHCP_einrichten)
- [4Netboot auf SGI Stationen](#Netboot_auf_SGI_Stationen)
  - [4.1Netboot Abbild herunterladen](#Netboot_Abbild_herunterladen)
  - [4.2DHCP Konfiguration für einen SGI Client](#DHCP_Konfiguration_f.C3.BCr_einen_SGI_Client)
  - [4.3Kernel Optionen](#Kernel_Optionen)
  - [4.4Starten der Daemons](#Starten_der_Daemons)
  - [4.5Netboot der SGI Station](#Netboot_der_SGI_Station)
  - [4.6Fehlerbehebung](#Fehlerbehebung)
- [5Netboot auf Cobalt Stationen](#Netboot_auf_Cobalt_Stationen)
  - [5.1Übersicht des Netboot Verfahrens](#.C3.9Cbersicht_des_Netboot_Verfahrens)
  - [5.2Cobalt Netboot Abbild herunterladen](#Cobalt_Netboot_Abbild_herunterladen)
  - [5.3NFS Serverkonfiguration](#NFS_Serverkonfiguration)
  - [5.4DHCP Konfiguration für eine Cobalt Maschine](#DHCP_Konfiguration_f.C3.BCr_eine_Cobalt_Maschine)
  - [5.5Starten der Dienste](#Starten_der_Dienste)
  - [5.6Netboot der Cobalt Maschine](#Netboot_der_Cobalt_Maschine)
  - [5.7Fehlerbehebung](#Fehlerbehebung_2)
- [6Eine Installations-CD verwenden](#Eine_Installations-CD_verwenden)

## Hardwareanforderungen

CPU (Big Endian Port)
MIPS3, MIPS4, MIPS5 oder MIPS64-Klasse CPU
CPU (Little Endian Port)
MIPS4, MIPS5 oder MIPS64-Klasse CPU
Arbeitsspeicher
128 MB
Festplattenspeicher
3.0 GB (ohne Swap-Speicher)
Swap-Speicher
Mindestens 256 MB

Für mehr Informationen, lesen Sie den [MIPS Hardware Requirements](/wiki/MIPS/Hardware_Requirements "MIPS/Hardware Requirements") en Artikel.

## Bemerkungen zur Installation

Auf vielen Architekturen hat der Prozessor mehrere Generationen durchschritten. Die jeweils neuere Generation baut auf dem Fundament der vorhergehenden auf. MIPS ist da keine Ausnahme.Es gibt mehrere CPU Generationen unter der MIPS Architektur. Um das Netboot-Abbild Stage Tarball und die CFLAGS richtig auszuwählen, ist es notwendig sich darüber im Klaren zu sein zu welcher Familie die CPU des Systems gehört. Diese Familien werden als Befehlssatzarchitektur (Instruction Set Architecture) bezeichnet.

MIPS ISA
32/64-Bit
CPU
MIPS 1
32-Bit
R2000, R3000
MIPS 2
32-Bit
R6000
MIPS 3
64-Bit
R4000, R4400, R4600, R4700
MIPS 4
64-Bit
R5000, RM5000, RM7000 R8000, R9000, R10000, R12000, R14000, R16000
MIPS 5
4-Bit
Bisher keine
MIPS32
32-Bit
AMD Alchemy Serie, 4kc, 4km, viele andere ... Es gibt ein paar Revisionen in der MIPS32 ISA.
MIPS64
64-Bit
Broadcom SiByte SB1, 5kc, etc. ... Es gibt ein paar Revisionen in der MIPS64 ISA.

**Hinweis**

Die MIPS5 ISA Ebene wurde im Jahr 1994 von Silocon Graphics entworfen, aber eigentlich nie wirklich in einer richtigen CPU verwendet. Sie lebt aber als Teil der MIPS64 ISA weiter.

**Hinweis**

Die MIPS32 und MIPS64 ISAs sind häufig eine Quelle der Verwirrung. Die MIPS64 ISA Ebene ist eigentlich eine Obermenge der MIPS5 ISA. Als solche enthält Sie alle Befehle von MIPS5 und früherer ISAs. MIPS32 ist die 32-Bit Untermenge von MIPS64. Sie existiert, weil die meisten Anwendungen nur 32-Bit Verarbeitung erfordern.

Ein weiteres wichtiges zu erfassendes Konzept ist das der Byte-Reihenfolge (Endianness). Die Byte-Reihenfolge bezieht sich auf die Art, in der eine CPU ein Datenwort aus dem Arbeitsspeicher liest. Ein Datenwort kann entweder als Big-Endian (höherwertiges Byte bei der niedrigeren Speicheradresse) oder als Little-Endian (niederwertiges Byte bei der niedrigeren Speicheradresse) gelesen werden. Intel x86 Maschinen verwenden grundsätzlich Little-Endian, Apple und Sparc Maschinen hingegen Big-Endian. MIPS können entweder das eine oder das andere sein. Um beide Arten zu unterscheiden hängen wir zur Kennzeichnung von Little-Endian _el_ an den Architekturnamen.

Architektur
32/64-Bit
Endianness
Maschinen
mips
32-Bit
Big-Endian
Silicon Graphics
mipsel
32-Bit
Little-Endian
Cobalt Servers
mips64
64-Bit
Big-Endian
Silicon Graphics
mips64el
64-Bit
Little-Endian
Cobalt Servers

Für jene die mehr über ISAs lernen wollen, könnten die folgenden Seiten hilfreich sein:

- [Linux/MIPS Website: MIPS ISA](http://www.linux-mips.org/wiki/index.php/Instruction_Set_Architecture) en
- [Linux/MIPS Website: Endianness](http://www.linux-mips.org/wiki/index.php/Endianess) en
- [Linux/MIPS Website: Processors](http://www.linux-mips.org/wiki/index.php/Processors) en
- [Wikipedia: Instruction Set](https://en.wikipedia.org/wiki/Instruction_set) en

## Netboot Übersicht

In diesem Abschnitt handeln wir das Notwendige was zum Netzwerk-booten einer Silicon Graphics Workstation oder eines Cobalt Servers benötigt wird ab. Dies ist nur ein kurzer Leitfaden. Es ist nicht beabsichtigt vollständig zu sein. Für weitere Informationen wird empfohlen den [Diskless Nodes](/wiki/Diskless_nodes "Diskless nodes") en Artikel zu lesen.

Abhängig von der Maschine gibt es eine gewisse Menge an Hardware die zum Netboot und zur Linux-Installation benötigt wird.

- Grundsätzlich:
  - DHCP/BOAMD Alchemy Serie, 4kc, 4km, viele andere ... Es gibt ein paar Revisionen von MIPS32 ISA.OTP Server (ISC DHCPd empfohlen)
  - Geduld -- und viel davon
- Für Silicon Graphics Workstations:
  - TFTP Server (tftp-hpa empfohlen)
  - Wenn die serielle Konsole verwendet werden muss:
    - MiniDIN8 --> RS-232 serielles Kabel (nur für IP22 und IP28 Systeme benötigt)
    - Nullmodem-Kabel
    - VT100 oder ANSI kompatibles Terminal, 9600 Baud fähig
- Für Cobalt Server (NICHT der original Qube):
  - NFS Server
  - Nullmodem-Kabel
  - VT100 oder ANSI kompatibles Terminal, 115200 Baud fähig

**Hinweis**

SGI Maschinen verwenden einen Mini-DIN-8 Anschluss für die seriellen Ports. Anscheinend funktionieren Apple Modemkabel als serielle Kabel gut. Seit es Apple Maschinen mit USB und internen Modems gibt, findet man diese Kabel aber immer schwerer. Ein Verdrahtungsplan ist im [Linux/MIPS Wiki](http://www.linux-mips.org/wiki/Serial_Cable) zu finden. Die meisten Elektronikläden sollten die erforderlichen Stecker auf Lager haben.

**Hinweis**

Als Terminal kann ein echtes VT100/ANSI Terminal oder ein PC mit einer Terminal Emulation Software (wie HyperTerminal, Minicom, seyon, Telex, xc, screen -- was auch immer Sie bevorzugen) zum Einsatz kommen. Es spielt keine Rolle auf welcher Plattform diese Maschine läuft -- solange Sie einen seriellen RS-232 Anschluss hat und passende Software.

**Hinweis**

Dieser Leitfaden deckt NICHT den original Qube ab. Der original Qube Server lässt in seiner Standardkonfiguration einen seriellen Anschluss vermissen. Deshalb ist es nicht möglich Gentoo auf ihm ohne die Hilfe eines Schraubendrehers und einer Hilfs-Maschine zu installieren.

### TFTP und DHCP einrichten

Wie bereits erwähnt ist dies keine komplette Anleitung. Es ist ein Grundgerüst, das die Dinge nur ins Rollen bringt. Verwenden Sie es um ein Setup von Grund auf neu zu beginnen, oder nutzen Sie die Vorschläge zur Erweiterung eines bestehenden Setup mit Netboot-Unterstützung.

Es ist erwähnenswert, dass die genutzten Server nicht Gentoo Linux verwenden müssen. Sie können auch gut und gerne FeeBSD oder jede andere Unix-ähnliche Plattform nutzen. Dieser Leitfaden geht von der Verwendung von Gentoo Linux aus. Wenn gewünscht, ist es auch möglich TFTP/NFS auf einer separaten Maschine am DHCP Server zu betreiben.

**Warnung**

Das Gentoo/MIPS Team kann Ihnen beim Einrichten anderer Betriebssysteme als Netboot-Server nicht helfen.

Erster Schritt -- DHCP konfigurieren. Damit der ISC DHCP Daemon auf BOOTP anfragen reagiert (wie es vom SGI und Cobalt BOOTROM benötigt wird) Aktivieren Sie als erstes das dynamische BOOTP für den verwendeten Adressbereich. Dann erstellen Sie einen Eintrag für jeden Client mit Zeigern auf das Bootabbild.

`root #` `emerge --ask net-misc/dhcp`

Sobald es installiert ist, erstellen Sie die Datei /etc/dhcp/dhcpd.conf. Hier ist ein Konfigurationsgrundgerüst um loszulegen.

DATEI **`/etc/dhcp/dhcpd.conf`** **dhcpd.conf Grundgerüst**

```
# Weisen Sie dhcpd an dynamisches DNS zu deaktivieren.
# dhcpd wird ohne dies den Start verweigern.
ddns-update-style none;

# Erstellen Sie ein Subnetz:
subnet 192.168.10.0 netmask 255.255.255.0 {
  # Der Adressraum für unsere Boot-Clients. Vergessen Sie nicht die 'dynamic-bootp' Angabe!
  pool {
    range dynamic-bootp 192.168.10.1 192.168.10.254;
  }

  # DNS Server und Standard Gateway -- ersetzten Sie die Adressen passend.
  option domain-name-servers 203.1.72.96, 202.47.56.17;
  option routers 192.168.10.1;

  # Sagen Sie dem DHCP Server, dass er für das Subnetz maßgeblich ("authoritative") ist.
  authoritative;

  # Gestatten Sie die Nutzung von BOOTP in diesem Subnetz.
  allow bootp;
}

```

Mit diesem Setup kann man eine beliebige Anzahl von Clients im Subnetz-Abschnitt hinzufügen. Wir besprechen das später in dieser Anleitung.

Nächster Schritt -- TFTP Server einrichten. Es wird empfohlen tftp-hpa zu verwenden, da es der einzige bekannte TFTP Daemon ist der korrekt arbeitet. Fahren Sie wie unten gezeigt mit der Installation fort:

`root #` `emerge --ask net-ftp/tftp-hpa`

Dies erzeugt das Verzeichnis /tftproot zur Aufbewahrung der Netboot Abbilder. Verschieben Sie es an eine andere Stelle, falls notwendig. Zum Zwecke dieser Anleitung wird davon ausgegangen, dass es an der vorgegebenen Stelle verbleibt.

## Netboot auf SGI Stationen

### Netboot Abbild herunterladen

Je nach System für das die Installation gedacht ist, gibt es mehrere mögliche Images zum Download. Diese sind je nach System- und CPU-Typ benannt für den Sie kompiliert wurden. Die Maschinentypen sind wie folgt:

Codename
Maschine
IP22
Indy, \*Indigo 2, Challenge S
IP26
\*Indigo 2 Power
IP27
Origin 200, Origin 2000
IP28
\*Indigo 2 Impact
IP30
Octane
IP32
O2

**Hinweis**

_Indigo 2_ \- Es ist ein verbreiteter Fehler die folgenden Maschinen zu verwechseln:

- IRIS Indigo (IP12 mit R3000 CPU oder IP20 mit R4000 CPU, auf keiner von beiden läuft Linux)
- Indigo 2 (IP22, auf dieser läuft Linux gut)
- Indigo 2 Power (R8000-basiert, Linux läuft darauf überhaupt nicht)
- Indigo 2 Impact (IP28 R10000-basiert, sehr experimentell)

Denken Sie daran, dass dies alles unterschiedliche Maschinen sind.

Im Dateinamen bezeichnet r4k den Prozessor der R4000-Serie, r5k den R5000, rm5k den RM5200 und r10k den R10000. Die Abbilder stehen auf den [Gentoo Spiegelservern](https://www.gentoo.org/downloads/mirrors/) zur Verfügung.

### DHCP Konfiguration für einen SGI Client

Nach dem Herunterladen der Datei legen Sie die entpackte Image-Datei in das /tftproot/ Verzeichnis (verwenden Sie bzip2 -d zum dekomprimieren). Bearbeiten Sie anschließend die Datei /etc/dhcp/dhcpd.conf und fügen Sie den passenden Eintrag für den SGI Client hinzu.

DATEI **`/etc/dhcp/dhcpd.conf`** **SGI Workstation Abschnitt**

```
subnet xxx.xxx.xxx.xxx netmask xxx.xxx.xxx.xxx {
  # ... anderes gewöhnliches Zeug ...

  # SGI Workstation ... ändern Sie 'sgi' in den Hostname Ihrer SGI Maschine.
  host sgi {

    # MAC Adresse der SGI Maschine.
    # Steht normalerweise auf der Rückseite oder Bodenplatte der Maschine.
    hardware ethernet 08:00:69:08:db:77;

    # Download TFTP Server (standardmäßig derselbe wie der DHCP Server)
    next-server 192.168.10.1;

    # IP Adresse der SGI Maschine
    fixed-address 192.168.10.3;

    # Dateiname der herunterzuladenden und zu bootenden PROM Datei
    filename "/gentoo-r4k.img";
  }
}

```

### Kernel Optionen

Wir sind fast fertig, aber es gibt noch ein paar Optimierungen durchzuführen. Öffnen Sie eine Konsole mit Root-Privilegien.

Deaktivieren Sie "Path Maximum Transfer Unit". Andernfalls findet der SGI PROM den Kernel nicht:

`root #` `echo 1 > /proc/sys/net/ipv4/ip_no_pmtu_disc`

Stellen Sie den vom SGI PROM nutzbaren Port-Bereich ein:

`root #` `echo "2048 32767" > /proc/sys/net/ipv4/ip_local_port_range`

`root #` `echo "2048 32767" > /proc/sys/net/ipv4/ip_local_port_range`

Dies sollte ausreichend sein, damit der Linux Server mit dem SGI PROM zusammenspielt.

### Starten der Daemons

Starten Sie hier die Dienste.

`root #` `/etc/init.d/dhcp start
`

`root #` `/etc/init.d/in.tftpd start
`

Wenn im letzten Schritt nichts schiefgegangen ist, sollte alles zum Einschalten der Workstation und zum Fortfahren mit der Anleitung bereit sein. Falls der DHCP Server aus welchen Gründen auch immer nicht startet, versuchen Sie dhcpd auf der Kommandozeile und schauen Sie was er sagt. Wenn alles gut läuft, sollte er direkt in den Hintergrund gehen. Andernfalls zeigt er 'exiting.' genau unterhalb seiner Beschwerde.

Eine einfache Möglichkeit zum Test ob der TFTP Daemon ausgeführt wird ist den folgenden Befehl einzugeben und seine Ausgabe zu prüfen:

`root #` `netstat -al | grep ^udp`

```
udp        0      0 *:bootpc                *:*
udp        0      0 *:631                   *:*
udp        0      0 *:xdmcp                 *:*
udp        0      0 *:tftp                  *:* <-- (suchen Sie nach dieser Zeile)

```

### Netboot der SGI Station

Alles ist eingestellt, DHCP läuft, genau wie TFTP. Jetzt ist es an der Zeit die SGI Maschine zu starten. Schalten Sie das Gerät ein. Wenn "Running power-on diagnostics" auf dem Bildschirm erscheint, klicken Sie entweder auf "Stop For Maintenance" oder drücken Sie die `Escape` Taste. Ein Menü wird angezeigt, das ähnlich wie das folgende aussieht.

`Running power-on diagnostics`

```
System Maintenance Menu

1) Start System
2) Install System Software
3) Run Diagnostics
4) Recover System
5) Enter Command Monitor
Option?

```

Geben Sie `5` ein, um zum 'Command Monitor' zu gelangen. Starten Sie den BootP Vorgang:

`>>` `bootp(): root=/dev/ram0`

Ab diesem Zeitpunkt sollte die Maschine den Download des Images beginnen und ungefähr 20 Sekunden später das Booten von Linux. Wenn alles gut geht, wird eine BusyBox ash shell wie unten gezeigt gestartet und die Installation von Gentoo Linux kann weitergehen.

CODE **Wenn alles gut geht ...**

```
init started:  BusyBox v1.00-pre10 (2004.04.27-02:55+0000) multi-call binary

Gentoo Linux; http://www.gentoo.org/
 Copyright 2001-2004 Gentoo Technologies, Inc.; Distributed under the GPL

 Gentoo/MIPS Netboot for Silicon Graphics Machines
 Build Date: April 26th, 2004

 * To configure networking, do the following:

 * For Static IP:
 * /bin/net-setup <IP Address> <Gateway Address> [telnet]

 * For Dynamic IP:
 * /bin/net-setup dhcp [telnet]

 * If you would like a telnetd daemon loaded as well, pass "telnet"
 * As the final argument to /bin/net-setup.

Please press Enter to activate this console.

```

### Fehlerbehebung

Wenn die Maschine stur ist und sich weigert ihr Image herunterzuladen, kann eines von zwei Dingen zutreffen:

1. Die Anweisungen wurden nicht korrekt befolgt, oder
2. Es bedarf ein wenig sanfter Überredung. (Nein, legen Sie den Vorschlaghammer weg!)

Hier ist eine Liste der Dinge, die sie überprüfen können:

- dhcpd gibt der SGI Maschine eine IP Adresse. Es sollten sich einige Meldungen über eine BOOTP Anfrage in den Logdateien befinden. tcdump ist hier auch sinnvoll.
- Die Berechtigungen sollten im tftp Verzeichnis korrekt gesetzt sein. (Typischerweise sollte /tftproot/ für alle lesbar sein.)
- Schauen Sie in die Systemlogdateien, um zu prüfen was der tftp Server berichtet (- vielleicht Fehler).

Wenn Sie alles auf dem Server überprüft haben und Timeouts oder andere Fehler auf der SGI Maschine aufgetreten sind, geben Sie dies in die Konsole ein:

`>>` `resetenv
`

`>>` `unsetenv netaddr
`

`>>` `unsetenv dlserver
`

`>>` `init
`

`>>` `bootp(): root=/dev/ram0`

## Netboot auf Cobalt Stationen

### Übersicht des Netboot Verfahrens

Im Gegensatz zu den SGI Maschinen verwenden Cobalt Server NFS um ihren Kernel zum Booten zu übertragen. Booten Sie die Maschine durch Niederdrücken der Pfeil nach links und rechts Tasten, während Sie das Gerät einschalten. Die Maschine wird dann versuchen einen IP-Adresse über BOOTP zu erhalten und das Verzeichnis /nfsroot/ des Servers über NFS einzuhängen. Sie versucht dann die Datei vmlinux\_raq-2800.gz (abhängig vom Maschinentyp) herunterzuladen und zu booten von der Sie ausgeht, dass es eine Standard-ELF-Datei ist.

### Cobalt Netboot Abbild herunterladen

Unter [http://distfiles.gentoo.org/experimental/mips/historical/netboot/cobalt/](http://distfiles.gentoo.org/experimental/mips/historical/netboot/cobalt/) werden die nötigen Boot-Images um einen Cobalt Server hoch und ans Laufen zu bekommen verfügbar gemacht. Die Dateien haben die Bezeichnung nfsroot-KERNEL-COLO-DATE-cobalt.tar. Wählen Sie die neueste und entpacken Sie diese in / wie unten gezeigt:

`root #` `tar -C / -xvf nfsroot-2.6.13.4-1.19-20051122-cobalt.tar`

### NFS Serverkonfiguration

Weil diese Maschinen NFS zum Download seines Images verwendet, ist es notwendig /nfsroot/ auf dem Server zu exportieren. Installieren Sie das Paket [net-fs/nfs-utils](https://packages.gentoo.org/packages/net-fs/nfs-utils):

`root #` `emerge --ask net-fs/nfs-utils`

Wenn das erledigt ist, schreiben Sie das Folgende in die Datei /etc/exports.

DATEI **`/etc/exports`** **Exporting the /nfsroot directory**

```
/nfsroot      *(ro,sync)

```

Starten Sie nun den NFS Server:

`root #` `/etc/init.d/nfs start`

Wenn der NFS Server bereits zu diesem Zeitpunkt ausgeführt wurde sagen Sie ihm mit `exportfs`, dass er erneut einen Blick auf seine exports Datei werfen soll.

`root #` `exportfs -av`

### DHCP Konfiguration für eine Cobalt Maschine

Nun ist die DHCP Seite der Dinge relativ unkompliziert. Fügen Sie das Folgende der Datei /etc/dhcp/dhcpd.conf hinzu.

DATEI **`/etc/dhcp/dhcpd.conf`** **Ausschnitt für Cobalt Server**

```
subnet xxx.xxx.xxx.xxx netmask xxx.xxx.xxx.xxx {
  # ... üblicher Kram hier ...

  # Konfiguration für einen Cobalt Server
  # Einstellen des Hostnamen hier:
  host qube {
    # Pfad zum Verzeichnis nfsroot.
    # Dies ist hauptsächlich dazu da, wenn Sie die TFTP Boot-Option von CoLo nutzen
    # Es sollte nicht notwendig sein das zu ändern.
    option root-path "/nfsroot";

    # Die Ethernet MAC-Adresse des Cobalt Server
    hardware ethernet 00:10:e0:00:86:3d;

    # Server von dem das Image heruntergeladen wird
    next-server 192.168.10.1;

    # IP Adresse des Cobalt Server
    fixed-address 192.168.10.2;

    # Die Lage der Datei default.colo relative zu /nfsroot
    # Es sollte nicht notwendig sein das zu ändern.
    filename "default.colo";
  }
}

```

### Starten der Dienste

Zum Starten der Dienste geben Sie folgendes ein:

`root #` `/etc/init.d/dhcp start
`

`root #` `/etc/init.d/nfs start`

Wenn im letzten Schritt nichts schiefgegangen ist, sollte alles zum Einschalten der Workstation und zum Fortfahren mit der Anleitung bereit sein. Falls der DHCP Server aus welchen Gründen auch immer nicht startet, versuchen Sie dhcpd auf der Kommandozeile und schauen Sie was er sagt. Wenn alles gut ist, sollte er direkt in den Hintergrund gehen. Andernfalls zeigt er 'exiting.' genau unterhalb seiner Beschwerde.

### Netboot der Cobalt Maschine

Nun ist es an der Zeit die Cobalt Maschinen zu starten. Koppeln Sie das Nullmodem-Kabel an und stellen Sie das serielle terminal auf 115200 Baud, 8 Bits, keine Parität, 1 Stopbit und VT100 Emulation. Sobald das erledigt ist halten Sie die rechte und linke Pfeiltaste gedrückt während Sie das Gerät einschalten.

Das Display auf der Rückseite sollte "Net Booting" anzeigen und einige Netzwerkaktivität sollte sichtbar sein, dicht gefolgt von CoLo. Scrollen Sie auf der rückwärtigen Anzeige nach unten bis zur "Network (NFS)" Option und drücken Sie `Enter`. Beachten Sie, dass die Maschine auf der seriellen Konsole zu booten beginnt.

`...`

```
elf: 80080000 <-- 00001000 6586368t + 192624t
elf: entry 80328040
net: interface down
CPU revision is: 000028a0
FPU revision is: 000028a0
Primary instruction cache 32kB, physically tagged, 2-way, linesize 32 bytes.
Primary data cache 32kB 2-way, linesize 32 bytes.
Linux version 2.4.26-mipscvs-20040415 (root@khazad-dum) (gcc version 3.3.3...
Determined physical RAM map:
 memory: 08000000 @ 00000000 (usable)
Initial ramdisk at: 0x80392000 (3366912 bytes)
On node 0 totalpages: 32768
zone(0): 32768 pages.
zone(1): 0 pages.
zone(2): 0 pages.
Kernel command line: console=ttyS0,115200 root=/dev/ram0
Calibrating delay loop... 249.85 BogoMIPS
Memory: 122512k/131072k available (2708k kernel code, 8560k reserved, 3424k dat)

```

Eine BusyBox ash Shell wird sich wie unten gezeigt öffnen, aus der die Gentoo Linux Installation fortgesetzt werden kann.

CODE **Wenn alles gut geht ...**

```
VFS: Mounted root (ext2 filesystem) readonly.
Freeing unused kernel memory: 280k freed
init started:  BusyBox v1.00-pre10 (2004.04.27-02:55+0000) multi-call binary

Gentoo Linux; http://www.gentoo.org/
 Copyright 2001-2004 Gentoo Technologies, Inc.; Distributed under the GPL

 Gentoo/MIPS Netboot for Cobalt Microserver Machines
 Build Date: April 26th, 2004

 * To configure networking, do the following:

 * For Static IP:
 * /bin/net-setup <IP Address> <Gateway Address> [telnet]

 * For Dynamic IP:
 * /bin/net-setup dhcp [telnet]

 * If you would like a telnetd daemon loaded as well, pass "telnet"
 * As the final argument to /bin/net-setup.

Please press Enter to activate this console.

```

### Fehlerbehebung

Wenn die Maschine stur ist und sich weigert ihr Image herunterzuladen, kann eines von zwei Dingen zutreffen:

1. Die Anweisungen wurden nicht korrekt befolgt, oder
2. Es bedarf ein wenig sanfter Überredung. (Nein, legen Sie den Vorschlaghammer weg!)

Hier ist eine Liste von Dingen, die Sie überprüfen können:

- dhcpd vergibt an die Cobalt Maschine eine IP-Adresse. Achten Sie auf Meldungen zu BOOTP Anfragen in den Systemlogdateien. tcpdump ist hier ebenfalls nützlich.
- Sind die Berechtigungen im Verzeichnis /nfsroot/ richtig gesetzt? (Der Inhalt sollte für alle lesbar sein.)
- Stellen Sie sicher, dass der NFS Server läuft und das Verzeichnis /nfsroot/ exportiert. Überprüfen Sie dies durch Eingabe des Befehls exportfs -v auf dem Server.

## Eine Installations-CD verwenden

Auf Silicon Graphics Maschinen ist es möglich von einer CD zu booten um Betriebssysteme zu installieren. (So istalliert man zum Beispiel IRIX.) Seit kurzem wurden Abbilder für solche bootbaren CDs zur Installation von Gentoo möglich. Diese CDs sind dafür entworfen, um in der gleichen Weise zu funktionieren.

Momentan funktioniert die Gentoo/MIPS Live-CD nur auf den SGI Indy, Indigo 2 und O2 Workstations, die mit der R4000 und R5000 CPU Serie ausgestattet sind. In der Zukunft können aber auch andere Plattformen möglich sein.

Die Live-CD Abbilder können auf einem Gentoo Spiegelserver unter dem Verzeichnis experimental/mips/livecd/ gefunden werden.

**Warnung**

Diese CDs sind zu diesem Zeitpunkt sehr experimentell. Es ist möglich, dass sie funktionieren oder nicht funktionieren. Bitte melden Sie Erfolg oder Fehler entweder auf Bugzilla, [diesem Forumsbeitrag](https://forums.gentoo.org/viewtopic.php?t=242518) oder im #gentoo-mips IRC-Kanal.

[← Über die Installation](/wiki/Handbook:MIPS/Installation/About/de "Previous part") [Home](/wiki/Handbook:MIPS "Handbook:MIPS") [Konfiguration des Netzwerks →](/wiki/Handbook:MIPS/Installation/Networking/de "Next part")