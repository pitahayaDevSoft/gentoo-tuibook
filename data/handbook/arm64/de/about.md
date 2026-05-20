# About

## Contents

- [1Einleitung](#Einleitung)
  - [1.1Willkommen](#Willkommen)
  - [1.2Wie ist die Installation strukturiert](#Wie_ist_die_Installation_strukturiert)
  - [1.3Was sind meine Optionen?](#Was_sind_meine_Optionen.3F)
  - [1.4Probleme?](#Probleme.3F)

## Einleitung

### Willkommen

Herzlich willkommen bei Gentoo. Sie sind dabei, in die Welt der Wahlmöglichkeiten und Performance einzusteigen. Bei Gentoo dreht sich vieles um Auswahlmöglichkeiten. Während der Installation von Gentoo wird Ihnen das mehrfach bewusst werden - Sie haben die Wahl, wie viele Pakete Sie selbst kompilieren, wie Sie Gentoo installieren, welchen Systemlogger Sie benutzen und vieles mehr.

Gentoo ist eine schnelle, moderne Metadistribution mit einem klaren und flexiblen Design. Gentoo ist auf einem Ökosystem freier Software gebaut und versteckt das, was unter der Haube steckt, nicht vor seinen Benutzern. Portage, das von Gentoo benutzte Paketmanagementsystem, ist in Python geschrieben, was bedeutet, dass Sie sich die Quelltexte einfach anschauen und nach Belieben verändern können. Gentoos Paketsystem benutzt den Quelltext (obwohl auch Unterstützung für vorkompilierte Pakete vorhanden ist) und die Konfiguration von Gentoo findet in normalen Textdateien statt. Mit anderen Worten: Offenheit überall.

Es ist sehr wichtig, dass jeder versteht, dass Auswahlmöglichkeiten das sind, was Gentoo ausmacht. Wir versuchen, Benutzer nicht zu Dingen zu zwingen, die diese nicht möchten. Sollte das doch vorkommen, erstellen Sie bitte einen [Bug Report](https://bugs.gentoo.org).

### Wie ist die Installation strukturiert

Die Installation von Gentoo kann als eine Prozedur von 10 Schritten gesehen werden, was den folgenden Kapiteln entspricht. Jeder Schritt führt zu einem bestimmten Ergebnis:

01. Nach Schritt 1 befindet sich der Anwender in einer funktionierenden Umgebung, aus der Gentoo installiert werden kann.
02. Nach Schritt 2 ist die Internetverbindung für die Gentoo-Installation vorbereitet.
03. Nach Schritt 3 sind die Festplatten für die Gentoo-Installation vorbereitet.
04. Nach Schritt 4 ist die Installationsumgebung vorbereitet und der Anwender ist bereit zum [chroot](/wiki/Chroot "Chroot") in die neue Umgebung.
05. Nach Schritt 5 sind die Kernpakete, die in allen Gentoo-Installationen gleich sind, installiert.
06. Nach Schritt 6 ist der Linux-Kernel installiert.
07. Nach Schritt 7 hat der Anwender die meisten Gentoo System-Konfigurationsdateien konfiguriert.
08. Nach Schritt 8 sind die notwendigen System-Tools installiert.
09. Nach Schritt 9 ist der gewählte Bootloader installiert und konfiguriert.
10. Nach Schritt 10 ist die neu installierte Gentoo Linux Umgebung bereit entdeckt zu werden.

Wenn Ihnen verschiedene Auswahlmöglichkeiten vorgestellt werden, geben wir unser Bestes, Ihnen die jeweiligen Vor- und Nachteile vorzustellen. Im weiteren Text wird zunächst eine Standardauswahl beschrieben (die im Titel durch "Standard:" gekennzeichnet ist), und anschließend die anderen Wahlmöglichkeiten (markiert durch "Alternativ:"). Die Standardauswahl ist nicht unbedingt das, was wir empfehlen. Es ist die Wahl, von der wir annehmen, dass die meisten Gentoo-Benutzer sie treffen werden.

Manchmal können Sie optionalen Schritten folgen. Solche Schritte sind als "Optional:" gekennzeichnet und nicht unbedingt notwendig, um Gentoo zu installieren. Dennoch können optionale Schritte von vorherigen Entscheidungen abhängen. Wir informieren Sie, wenn das passiert. Sowohl wenn Sie die Entscheidung treffen, als auch wenn der optionale Schritt beschrieben wird.

### Was sind meine Optionen?

Gentoo kann auf viele verschiedene Arten installiert werden. Für ARM64-Hardware beinhaltet dies in der Regel das Herunterladen eines System-Images zum DD'ing auf eMMC-, SD- oder uSD-Medien.

Dieses Dokument behandelt die Installation speziell für ARM64-Hardware. Wir versuchen, die Besonderheiten bekannter Hardware abzudecken, die wir testen und validieren konnten. ARM64-Hardware, die nicht aufgelistet ist, wird wahrscheinlich funktionieren, aber Sie müssen verstehen, was für Ihre Hardware erforderlich ist und Ihre Installation entsprechend anpassen.

**Hinweis**

Für weitere Hilfe zu den anderen Installationsmöglichkeiten, einschließlich der Nutzung von nicht-Gentoo CDs, lesen Sie bitte unseren [Leitfaden über alternative Installationsmöglichkeiten](/wiki/Installation_alternatives "Installation alternatives").

Wir bieten ebenfalls ein [Gentoo Installation Tipps & Tricks-Dokument](/wiki/Gentoo_installation_tips_and_tricks "Gentoo installation tips and tricks"), das weitere nützliche Informationen enthält.

### Probleme?

Wenn Sie ein Problem während der Installation (oder in der Dokumentation) entdecken, schauen Sie bitte in unserem [Bug-Tracking-System](https://bugs.gentoo.org), ob der Fehler bereits bekannt ist. Wenn nicht, erstellen Sie bitte einen Fehlerbericht, damit wir uns der Sache annehmen können. Haben Sie keine Angst vor den Entwicklern, denen Ihr Fehlerbericht zugeteilt wird -- für gewöhnlich essen sie keine Menschen.

Obwohl das Dokument, das Sie gerade lesen, architekturspezifisch ist, kann es Referenzen zu anderen Architekturen enthalten. Das liegt daran, dass viele Teile des Gentoo Handbuchs Textpassagen verwenden, die für alle Architekturen gleich sind (um doppelten Arbeitsaufwand zu vermeiden). Wir versuchen, solche Referenzen auf ein Minimum zu beschränken, um Missverständnisse zu vermeiden.

Wenn Sie sich nicht sicher sind, ob ein Problem ein Benutzerproblem ist (ein Fehler, den Sie trotz sorgfältiger Lektüre dieser Dokumentation machen) oder ein Softwareproblem (ein Fehler, den wir trotz sorgfältigen Tests der Installation/Dokumentation begangen haben) sollten Sie den Channel [#gentoo-de](ircs://irc.libera.chat/#gentoo-de) ([webchat](https://web.libera.chat/#gentoo-de)) im irc.libera.chat Netz besuchen. Natürlich sind Sie auch sonst willkommen, da unser Chat-Channel alle Gentoo-Themen abdeckt.

Apropos, wenn Sie eine weitere Frage hinsichtlich Gentoo haben, werfen Sie zunächst einen Blick in den Artikel [Häufig gestellte Fragen (FAQ)](/wiki/FAQ/de "FAQ/de") hier im Wiki. Sie können auch die [FAQs](https://forums.gentoo.org/viewforum.php?f=40) in unserem [Forum](https://forums.gentoo.org) lesen.

[←](/index.php?title=Handbook:ARM64/de&action=edit&redlink=1 "Previous part") [Anfang](/index.php?title=Handbook:ARM64/de&action=edit&redlink=1 "Handbook:ARM64/de (page does not exist)") [→](/index.php?title=Handbook:ARM64/Installation/Media/de&action=edit&redlink=1 "Next part")