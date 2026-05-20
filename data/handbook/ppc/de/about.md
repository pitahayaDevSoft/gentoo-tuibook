# About

Other languages:

- Deutsch
- [English](/wiki/Handbook:PPC/Installation/About "Handbook:PPC/Installation/About (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/About/es "Manual de Gentoo: PPC/Instalación/Acerca de (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/About/fr "Handbook:PPC/Installation/About/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/About/it "Handbook:PPC/Installation/About (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/About/hu "Handbook:PPC/Installation/About/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/About/pl "Handbook:PPC/Installation/About (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/About/pt-br "Manual:PPC/Instalação/Sobre (100% translated)")
- [русский](/wiki/Handbook:PPC/Installation/About/ru "Handbook:PPC/Installation/About (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/About/ta "கையேடு:PPC/நிறுவல்/பற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/About/zh-cn "手册：PPC/安装/关于 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/About/ja "ハンドブック:PPC/インストール/About (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/About/ko "Handbook:PPC/Installation/About/ko (100% translated)")

[PPC Handbuch](/wiki/Handbook:PPC/de "Handbook:PPC/de")[Installation](/wiki/Handbook:PPC/Full/Installation/de "Handbook:PPC/Full/Installation/de")Über die Installation[Auswahl des Mediums](/wiki/Handbook:PPC/Installation/Media/de "Handbook:PPC/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:PPC/Installation/Networking/de "Handbook:PPC/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:PPC/Installation/Disks/de "Handbook:PPC/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:PPC/Installation/Stage/de "Handbook:PPC/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:PPC/Installation/Base/de "Handbook:PPC/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:PPC/Installation/Kernel/de "Handbook:PPC/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:PPC/Installation/System/de "Handbook:PPC/Installation/System/de")[Installation der Tools](/wiki/Handbook:PPC/Installation/Tools/de "Handbook:PPC/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbook:PPC/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:PPC/Installation/Finalizing/de "Handbook:PPC/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:PPC/Full/Working/de "Handbook:PPC/Full/Working/de")[Portage-Einführung](/wiki/Handbook:PPC/Working/Portage/de "Handbook:PPC/Working/Portage/de")[USE-Flags](/wiki/Handbook:PPC/Working/USE/de "Handbook:PPC/Working/USE/de")[Portage-Features](/wiki/Handbook:PPC/Working/Features/de "Handbook:PPC/Working/Features/de")[Initskript-System](/wiki/Handbook:PPC/Working/Initscripts/de "Handbook:PPC/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:PPC/Working/EnvVar/de "Handbook:PPC/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:PPC/Full/Portage/de "Handbook:PPC/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:PPC/Portage/Files/de "Handbook:PPC/Portage/Files/de")[Variablen](/wiki/Handbook:PPC/Portage/Variables/de "Handbook:PPC/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:PPC/Portage/Branches/de "Handbook:PPC/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:PPC/Portage/Tools/de "Handbook:PPC/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:PPC/Portage/CustomTree/de "Handbook:PPC/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:PPC/Portage/Advanced/de "Handbook:PPC/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:PPC/Full/Networking/de "Handbook:PPC/Full/Networking/de")[Zu Beginn](/wiki/Handbook:PPC/Networking/Introduction/de "Handbook:PPC/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:PPC/Networking/Advanced/de "Handbook:PPC/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:PPC/Networking/Modular/de "Handbook:PPC/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:PPC/Networking/Wireless/de "Handbook:PPC/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:PPC/Networking/Extending/de "Handbook:PPC/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:PPC/Networking/Dynamic/de "Handbook:PPC/Networking/Dynamic/de")

## Contents

- [1Einleitung](#Einleitung)
  - [1.1Willkommen](#Willkommen)
    - [1.1.1Offenheit](#Offenheit)
    - [1.1.2Auswahl](#Auswahl)
    - [1.1.3Macht](#Macht)
  - [1.2Wie ist die Installation strukturiert](#Wie_ist_die_Installation_strukturiert)
    - [1.2.1Auswahl der Schritte](#Auswahl_der_Schritte)
    - [1.2.2Empfohlene Schritte](#Empfohlene_Schritte)
    - [1.2.3Optionale Schritte](#Optionale_Schritte)
    - [1.2.4Veraltete Schritte](#Veraltete_Schritte)
    - [1.2.5Standards und Alternativen](#Standards_und_Alternativen)
  - [1.3Installationsoptionen von Gentoo](#Installationsoptionen_von_Gentoo)
  - [1.4Probleme](#Probleme)

## Einleitung

### Willkommen

Willkommen bei Gentoo! Gentoo ist ein freies, auf Linux basierendes Betriebssystem, das automatisch für nahezu jede Anwendung oder Anforderung optimiert und angepasst werden kann. Es baut auf einem Ökosystem freier Software auf und verbirgt nicht, was unter der Haube läuft, vor seinen Benutzern.

#### Offenheit

Die wichtigsten Werkzeuge von Gentoo sind in einfachen Programmiersprachen geschrieben. [Portage](/wiki/Portage "Portage"), das Paketverwaltungssystem von Gentoo ist [in Python geschrieben](https://gitweb.gentoo.org/proj/portage.git). Ebuilds, die Paketdefinitionen für Portage bereitstellen [sind in Bash geschrieben](https://gitweb.gentoo.org/proj/gentoo.git). Unsere Benutzer:innen werden damit ermutigt, den Quellcode für alle Teile von Gentoo zu überprüfen, zu verändern und zu verbessern.

Standardmäßig werden Pakete nur gepatcht, wenn es notwendig ist, um Fehler zu beheben oder Interoperabilität innerhalb von Gentoo zu gewährleisten. Sie werden auf dem System installiert, indem der Quellcode von Upstream-Projekten in das Binärformat kompiliert wird (obwohl auch vorkompilierte Binärpakte unterstützt werden). Die Konfiguration von Gentoo erfolgt über Textdateien.

Aus den oben genannten und anderen Gründen ist _Offenheit_ als ein _Designprinzip_ eingebaut.

#### Auswahl

_Auswahl_ ist ein weiteres Gentoo _Designprinzip._

Bei der Installation von Gentoo wird die Auswahl im Handbuch deutlich gemacht. Systemadministratoren können zwei vollständig unterstützte Init-Systeme wählen (Gentoos eigenes [OpenRC](/wiki/OpenRC/de "OpenRC/de") und Freedesktop.org's [systemd](/wiki/Systemd/de "Systemd/de")), die Partitionstruktur für die Festplatte(n), welche Dateisysteme auf der/den Platte(n) verwendet werden sollen, ein Ziel- [Systemprofil](/wiki/Profile "Profile"), entfernen oder hinzufügen von Funktionen auf globaler (systemweiter) oder paketspezifischer Ebene über USE-Flags, Bootloader, Netzwerkverwaltungsprogramm und vieles, vieles mehr.

Als Entwicklungsphilosophie versuchen [Gentoo Autoren](https://www.gentoo.org/inside-gentoo/developers) es zu vermeiden, die Benutzer auf ein bestimmtes Systemprofil oder eine bestimmte Desktopumgebung zu zwingen. Wenn etwas im GNU/Linux Ökosystem angeboten wird, ist es wahrscheinlich auch in Gentoo verfügbar. Wenn nicht dann würden wir es gerne so sehen. Für neue Paketanfragen wird empfohlen zunächst ein Paket an [GURU](/wiki/GURU "GURU") zu senden. Sobald das Paket stabil ist und sich eine Entwickler:in bereit erklärt das neue Paket zu unterstützen, kann es dem offiziellen Gentoo Repositorium hinzugefügt werden.

#### Macht

Da Gentoo ein quellbasiertes Betriebssystem ist, kann es auf neue [Computerarchitekturen](https://en.wikipedia.org/wiki/instruction_set_architecture "wikipedia:instruction set architecture") portiert werden und alle installierten Pakete können angepasst werden. Diese Stärke zeigt ein weiteres _Designprinzip_ von Gentoo: _Macht_.

Ein Systemadministrator, der Gentoo erfolgreich installiert und angepasst hat, hat ein maßgeschneidertes Betriebssystem aus dem Quellcode kompiliert. Das gesamte Betriebssystem kann auf binärer Ebene über die Mechanismen in der Portage-Datei [make.conf](/wiki//etc/portage/make.conf/de "/etc/portage/make.conf/de") angepasst werden. Falls gewünscht, können Anpassungen auf der Basis einzelner Pakete oder Paketgruppen vorgenommen werden. Mit USE-Flags können sogar ganze Sätze von Funktionen hinzugefügt oder entfernt werden.

Es ist sehr wichtig, dass die Handbuch Lesenden verstehen, dass diese Designprinzipien Gentoo einzigartig machen. Mit den Prinzipien der großen Macht, der vielen Wahlmöglichkeiten und der extremen Offenheit, die hervorgehoben werden, sollte man bei der Verwendung von Gentoo Sorgfalt, Nachdenken und Absicht walten lassen.

### Wie ist die Installation strukturiert

Die Installation von Gentoo kann als eine Prozedur von 10 Schritten gesehen werden, was den Kapiteln 2 bis 11 entspricht. Jeder Schritt führt zu einem bestimmten Ergebnis:

SchrittErgebnis
1Die Benutzer:in befindet sich in einer funktionierenden Umgebung aus der Gentoo installiert werden kann.
2Die Internetverbindung ist für die Gentoo-Installation vorbereitet.
3Die Festplatten sind für die Gentoo-Installation vorbereitet.
4Die Installationsumgebung ist vorbereitet und die Benutzer:in ist bereit zum [chroot](/wiki/Chroot/de "Chroot/de") in die neue Umgebung.
5Die Kernpakete, die in allen Gentoo-Installationen gleich sind, sind installiert.
6Der Linux-Kernel ist installiert.
7Die Benutzer:in hat die meisten Gentoo-Systemkonfigurationsdateien konfiguriert.
8Die notwendigen System-Tools sind installiert.
9Der gewählte Bootloader ist installiert und konfiguriert.
10Die neu installierte Gentoo Linux Umgebung ist bereit entdeckt zu werden.

#### Auswahl der Schritte

Das Handbuch bietet eine überwältigende Menge an Optionen, besonders für diejenigen, die nie ein Linux ohne Installationsprogramm installiert haben.

Wichtig zu verstehen ist, dass das Handbuch im Hinblick darauf gestaltet ist, die, für eine Installation auf einer großen Bandbreite an Hardware mit unterschiedlichen Anforderungen an die Installation, notwendigen Schritte zu erklären. Daher sind viele im Handbuch dargestellten Optionen für einzelne Installationen unnötig und können übersprungen werden.

#### Empfohlene Schritte

Schritte denen " **Empfohlen:**" vorangestellt ist, sind nicht zwingend notwendig, aber in den meisten Fällen hilfreich, zum Beispiel die Installation von [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware).

#### Optionale Schritte

Abschnitte denen " **Optional:**" vorangestellt ist, sind rein fakultativ und können übersprungen werden, wenn eine einfache Installation ohne Schnörkel angestrebt wird.

Beispiele hierfür sind Compiler Flag Anpassungen, Nutzung eines vollkommen angepassten Kernels und das Deaktivieren des root Logins.

**Tipp**

Beim folgen optionaler Schritte ist es wichtig darauf zu achten, dass alle Voraussetzungen erfüllt wurden. Einige optionale Schritte hängen von anderen optionalen Schritten ab.

#### Veraltete Schritte

Gentoo existiert seit einer langen Zeit. Einige Installationsprozeduren sind im Handbuch beschrieben, weil sie einst eine größere Relevanz besaßen, aber inzwischen größtenteils veraltet sind. Anstatt diese Informationen sofort zu entfernen, werden sie vor ihrer Entfernung als **Veraltet:** markiert, da sie für einige immer noch hilfreich sein könnten. Sobald sie entfernt wurden, muß die _Historie_ verwendet werden, um jene Inhalte zu betrachten.

#### Standards und Alternativen

Wann immer das Handbuch Auswahlmöglichkeiten enthält, ist die Bestrebung das Pro und Kontra der einzelnen Möglichkeiten darzulegen.

Wenn Wahlmöglichkeiten sich gegenseitig ausschließen wird mit " **Standard:**" die am besten unterstützte oder meist gewählte Option bezeichnet, wohingegen Alternativen unter " **Alternative**" aufgeführt werden.

**Hinweis**

**Alternative** Optionen sind nicht minderwertiger als **Standard** s, aber typischerweise sind **Standard** Optionen weiter verbreitet und werden eventuell besser unterstützt.

### Installationsoptionen von Gentoo

Gentoo kann auf vielen verschiedenen Wegen installiert werden. Es kann als offizielles Gentoo Installationsmedium wie unsere bootbaren ISO-Images heruntergeladen und von dort installiert werden. Dieses Image kann auf einen USB-Stick kopiert oder aus dem Netzwerk gebootet werden. Alternativ kann Gentoo von einem nicht offiziellem Medium, wie zum Beispiel aus einer bereits installierten Distribution heraus oder von einem anderen, nicht von Gentoo herausgegebenen, bootbaren Datenträger (wie z.B. [Linux Mint](https://linuxmint.com/)) installiert werden.

Dieses Dokument beschreibt die Installation mit einem offiziellen Gentoo Installations-Datenträger, oder, in bestimmten Fällen, Netboot.

**Hinweis**

Für weitere Hilfe zu den anderen Installationsmöglichkeiten, einschließlich der Nutzung von nicht-Gentoo Boot-Medien, siehe [Leitfaden über alternative Installationsmethoden](/wiki/Installation_alternatives/de "Installation alternatives/de").

Wir bieten ebenfalls ein [Gentoo Installation Tipps & Tricks-Dokument](/wiki/Gentoo_installation_tips_and_tricks "Gentoo installation tips and tricks"), das weitere nützliche Informationen enthält.

### Probleme

Wenn ein Problem während der Installation (oder in der Dokumentation) auftritt, bitten wir darum unsere [Bug-Tracking-System](https://bugs.gentoo.org) zu besuchen und zu sehen, ob der Fehler bereits bekannt ist. Wenn nicht, bitten wir um die Erstellung eines Fehlerberichts, damit wir uns der Sache annehmen können. Keine Angst vor den Entwickler:innen, denen der Fehlerbericht zugeteilt wird -- für gewöhnlich essen sie keine Menschen.

Obwohl dieses Dokument architekturspezifisch ist, kann es Referenzen zu anderen Architekturen enthalten. Das liegt daran, dass viele Teile des Gentoo Handbuchs Textpassagen verwenden, die für alle Architekturen gleich sind (um doppelten Arbeitsaufwand zu vermeiden). Wir versuchen, solche Referenzen auf ein Minimum zu beschränken, um Missverständnisse zu vermeiden.

Sollte unklar sein, ob ein Problem ein Anwendungsproblem ist (ein Fehler, der trotz sorgfältiger Lektüre dieser Dokumentation gemacht wird) oder ein Softwareproblem (ein Fehler, den wir trotz sorgfältiger Tests der Installation/Dokumentation begangen haben) laden wir dazu ein den Channel [#gentoo-de](ircs://irc.libera.chat/#gentoo-de) ([webchat](https://web.libera.chat/#gentoo-de)) im irc.libera.chat Netz besuchen. Natürlich sind alle auch sonst willkommen, da unser Chat-Channel alle Gentoo-Themen abdeckt.

Apropos, wenn es weitere Fragen hinsichtlich Gentoo gibt, lohnt ein Blick in den Artikel [Häufig gestellte Fragen (FAQ)](/wiki/FAQ/de "FAQ/de") hier im Wiki. Es gibt außerdem noch die [FAQs](https://forums.gentoo.org/viewforum.php?f=40) in unserem [Forum](https://forums.gentoo.org).

[←](/wiki/Handbook:PPC/de "Previous part") [Anfang](/wiki/Handbook:PPC/de "Handbook:PPC/de") [Auswahl des Mediums →](/wiki/Handbook:PPC/Installation/Media/de "Next part")