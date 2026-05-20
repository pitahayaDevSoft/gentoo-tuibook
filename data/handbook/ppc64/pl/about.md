# About

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/About/de "Handbuch:PPC64/Installation/Über (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/About "Handbook:PPC64/Installation/About (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/About/es "Manual de Gentoo: PPC64/Instalación/Acerca de (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/About/fr "Handbook:PPC64/Installation/About/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/About/it "Handbook:PPC64/Installation/About/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/About/hu "Handbook:PPC64/Installation/About/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:PPC64/Installation/About/pt-br "Manual:PPC64/Instalação/Sobre (100% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/About/ru "Handbook:PPC64/Installation/About (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/About/ta "கையேடு:PPC64/நிறுவல்/பற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/About/zh-cn "手册：PPC64/安装/关于 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/About/ja "ハンドブック:PPC64/インストール/About (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/About/ko "Handbook:PPC64/Installation/About/ko (100% translated)")

[PPC64 Handbook](/wiki/Handbook:PPC64/pl "Handbook:PPC64/pl")[Instalacja](/wiki/Handbook:PPC64/Full/Installation/pl "Handbook:PPC64/Full/Installation/pl")O instalacji[Wybór medium instalacyjnego](/wiki/Handbook:PPC64/Installation/Media/pl "Handbook:PPC64/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:PPC64/Installation/Networking/pl "Handbook:PPC64/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:PPC64/Installation/Disks/pl "Handbook:PPC64/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:PPC64/Installation/Stage/pl "Handbook:PPC64/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:PPC64/Installation/Base/pl "Handbook:PPC64/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:PPC64/Installation/Kernel/pl "Handbook:PPC64/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:PPC64/Installation/System/pl "Handbook:PPC64/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:PPC64/Installation/Tools/pl "Handbook:PPC64/Installation/Tools/pl")[Instalacja systemu rozruchowego](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader/pl")[Finalizacja](/wiki/Handbook:PPC64/Installation/Finalizing/pl "Handbook:PPC64/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:PPC64/Full/Working/pl "Handbook:PPC64/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:PPC64/Working/Portage/pl "Handbook:PPC64/Working/Portage/pl")[Flagi USE](/wiki/Handbook:PPC64/Working/USE/pl "Handbook:PPC64/Working/USE/pl")[Funkcje portage](/wiki/Handbook:PPC64/Working/Features/pl "Handbook:PPC64/Working/Features/pl")[System initscript](/wiki/Handbook:PPC64/Working/Initscripts/pl "Handbook:PPC64/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:PPC64/Working/EnvVar/pl "Handbook:PPC64/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:PPC64/Full/Portage/pl "Handbook:PPC64/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:PPC64/Portage/Files/pl "Handbook:PPC64/Portage/Files/pl")[Zmienne](/wiki/Handbook:PPC64/Portage/Variables/pl "Handbook:PPC64/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:PPC64/Portage/Branches/pl "Handbook:PPC64/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:PPC64/Portage/Tools/pl "Handbook:PPC64/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:PPC64/Portage/CustomTree/pl "Handbook:PPC64/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:PPC64/Portage/Advanced/pl "Handbook:PPC64/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:PPC64/Full/Networking/pl "Handbook:PPC64/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:PPC64/Networking/Introduction/pl "Handbook:PPC64/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:PPC64/Networking/Advanced/pl "Handbook:PPC64/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:PPC64/Networking/Modular/pl "Handbook:PPC64/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:PPC64/Networking/Wireless/pl "Handbook:PPC64/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:PPC64/Networking/Extending/pl "Handbook:PPC64/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:PPC64/Networking/Dynamic/pl "Handbook:PPC64/Networking/Dynamic/pl")

## Contents

- [1Wprowadzenie](#Wprowadzenie)
  - [1.1Witaj](#Witaj)
    - [1.1.1Openness](#Openness)
    - [1.1.2Choice](#Choice)
    - [1.1.3Power](#Power)
  - [1.2Jak zbudowana jest instalacja](#Jak_zbudowana_jest_instalacja)
    - [1.2.1Deciding which steps to take](#Deciding_which_steps_to_take)
    - [1.2.2Suggested steps](#Suggested_steps)
    - [1.2.3Optional steps](#Optional_steps)
    - [1.2.4Deprecated steps](#Deprecated_steps)
    - [1.2.5Defaults and alternatives](#Defaults_and_alternatives)
  - [1.3Opcje instalacyjne Gentoo](#Opcje_instalacyjne_Gentoo)
  - [1.4Problemy](#Problemy)

## Wprowadzenie

### Witaj

Przede wszystkim, witamy w Gentoo! Zamierzasz wkroczyć w świat wyboru i możliwości. Gentoo polega na wyborach. Podczas instalacji Gentoo, jest to wielokrotnie jasno przedstawione - użytkownicy mogą wybrać co ma zawierać ich kompilacja, jak zainstalować Gentoo, jakiego loggera systemowego użyć, etc.

#### Openness

Gentoo's premier tools are built from simple programming languages. [Portage](/wiki/Portage "Portage"), Gentoo's package maintenance system, is [written in Python](https://gitweb.gentoo.org/proj/portage.git/). Ebuilds, which provide package definitions for Portage [are written in bash](https://gitweb.gentoo.org/repo/gentoo.git). Our users are encouraged to review, modify, and enhance the source code for all parts of Gentoo.

By default, packages are only patched when necessary to fix bugs or provide interoperability within Gentoo. They are installed to the system by compiling source code provided by upstream projects into binary format (although support for precompiled binary packages is included too). Configuring Gentoo happens through text files.

For the above reasons and others: _openness_ is built-in as a _design principle_.

#### Choice

Gentoo to szybka, nowoczesna metadystrybucja o przejrzystym i elastycznym wyglądzie. Jest zbudowany w oparciu o ekosystem wolnego oprogramowania i nic nie ukrywa przed swoimi użytkownikami. Portage, system zarządzania pakietami, z którego korzysta Gentoo, został napisany w języku Python. Oznacza to, że użytkownik może łatwo przeglądać i modyfikować kod źródłowy. System pakowania Gentoo używa kodu źródłowego (dostępne jest również wsparcie dla wstępnie skompilowanych pakietów), a konfiguracja Gentoo odbywa się za pomocą zwykłych plików tekstowych. Innymi słowy, otwartość wszędzie.

When installing Gentoo, choice is made clear throughout the Handbook. System administrators can choose two fully supported init systems (Gentoo's own [OpenRC](/wiki/OpenRC "OpenRC") and Freedesktop.org's [systemd](/wiki/Systemd "Systemd")), partition structure for storage disk(s), what file systems to use on the disk(s), a target [system profile](/wiki/Profile "Profile"), remove or add features on a global (system-wide) or package specific level via USE flags, bootloader, network management utility, and much, much more.

As a development philosophy, [Gentoo's authors](https://www.gentoo.org/inside-gentoo/developers/) try to avoid forcing users onto a specific system profile or desktop environment. If something is offered in the GNU/Linux ecosystem, it's likely available in Gentoo. If not, then we'd love to see it so. For new packages, it is recommended to first submit a package to [GURU](/wiki/GURU "GURU"). Once it has matured and a Gentoo developer has volunteered to sponsor the new package, it can then be added to the official Gentoo package repository.

#### Power

Being a source-based operating system allows Gentoo to be ported onto new computer [instruction set architectures](https://en.wikipedia.org/wiki/instruction_set_architecture "wikipedia:instruction set architecture") and also allows all installed packages to be tuned. This strength surfaces another Gentoo _design principle_: _power_.

A system administrator who has successfully installed and customized Gentoo has compiled a tailored operating system from source code. The entire operating system can be tuned at a binary level via the mechanisms included in Portage's [make.conf](/wiki//etc/portage/make.conf/pl "/etc/portage/make.conf/pl") file. If so desired, adjustments can be made on a per-package basis, or a package group basis. In fact, entire sets of functionality can be added or removed using USE flags.

Jest to bardzo ważne, by każdy użytkownik zrozumiał, iż możliwość wyboru jest głównym czynnikiem istnienia Gentoo. Staramy się nie zmuszać użytkowników do robienia czegokolwiek, co im się nie podoba. Jeśli ktoś uważa inaczej, prosimy o zgłoszenie [raportu o błędzie](https://bugs.gentoo.org).

### Jak zbudowana jest instalacja

Instalacja Gentoo może być przedstawiona jako 10-etapowa procedura, odpowiadająca następnym zestawom działań. Wyniki każdego działaia w danym etapie:

EtapWynik
1Użytkownik znajduje się w środowisku roboczym, które jest gotowe do zainstalowania Gentoo.
2Połączenie internetowe jest gotowe do zainstalowania Gentoo.
3Dyski twarde są przygotowane do obsługi instalacji Gentoo.
4Środowisko instalacyjne jest przygotowane, a użytkownik jest gotowy do wykonania [chroot](/wiki/Chroot "Chroot") w nowym środowisku.
5Instalowane są podstawowe pakiety, które są takie same we wszystkich instalacjach Gentoo.
6Jądro Linuksa jest zainstalowane.
7Większość plików konfiguracyjnych systemu Gentoo jest utworzona.
8Zainstalowane są niezbędne narzędzia systemowe.
9Odpowiedni system rozruchowy został zainstalowany i skonfigurowany.
10Świeżo zainstalowane środowisko Gentoo Linux jest gotowe do eksploracji.

#### Deciding which steps to take

The handbook presents an overwhelming amount of options, especially for someone who has never installed Linux without an installer.

It's important to realize that the handbook is designed to describe the steps required to install on a very wide variety of hardware, with different install needs. Because of this, many options presented in the handbook are unnecessary for a particular install and can be skipped.

#### Suggested steps

Prefixed with " **Suggested:**", some steps are not strictly required, but help in most cases, such as installing [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware).

#### Optional steps

Prefixed with " **Optional:**", many sections in the handbook are purely optional, and can be skipped if the user is looking for a simple, mostly vanilla install.

Examples of this include compiler flag customization, using a totally custom kernel, and disabling root login.

**Wskazówka**

When following optional steps, it's important to be careful that all prerequisites were satisfied. Some optional steps depend on other optional steps.

#### Deprecated steps

Gentoo has existed for a long time. Some install processes are described in the handbook because they used to be more relevant, but are now largely deprecated. Instead of immediately removing this information, as it may still be helpful for some users, it may be designated as **Deprecated:** before removal. Once removed, the _history_ must be used to view this content.

#### Defaults and alternatives

Whenever choices are presented, the handbook will try to explain the pros and cons of each choice.

If potential choices are mutually exclusive, " **Default:**" is used to designate the most supported or commonly chosen option, while alternatives are listed with " **Alternative**".

**Ważne**

**Alternative** options are not inferior to **Default** s, but **Default** options are typically more widely used and may have better support.

### Opcje instalacyjne Gentoo

Gentoo można zainstalować na wiele różnych sposobów. Można go pobrać i zainstalować z oficjalnych nośników instalacyjnych Gentoo, takich jak nasze płyty CD i DVD. Nośnik instalacyjny można zainstalować na pamięci USB lub uzyskać do niego dostęp za pośrednictwem środowiska sieciowego. Alternatywnie, Gentoo można zainstalować z nieoficjalnych nośników, takich jak już zainstalowana dystrybucja lub dysk startowy inny niż Gentoo (np. [Knoppix](Https://www.knopper.net/knoppix/index-en.html)).

Ten dokument opisuje instalację przy użyciu oficjalnego nośnika instalacyjnego Gentoo, a w niektórych przypadkach instalację sieciową.

**Informacja**

Aby uzyskać pomoc dotyczącą innych podejść do instalacji, w tym używania płyt CD innych niż Gentoo, przeczytaj nasz [Alternatywny przewodnik instalacji](/wiki/Installation_alternatives/pl "Installation alternatives/pl").

Udostępniamy również dokument [Wskazówki i triki dotyczące instalacji Gentoo](/wiki/Gentoo_installation_tips_and_tricks/pl "Gentoo installation tips and tricks/pl"), który może okazać się przydatny.

### Problemy

Jeśli wystąpił problem w instalacji (lub dokumentacji instalacji), odwiedź nasz [system śledzenia błędów](https://bugs.gentoo.org) i sprawdź czy błąd jest już znany. Jeśli nie, utwórz raport o błędzie, abyśmy mogli się tym zająć. Nie bój się programistów, którzy są przypisani do błędów - (zazwyczaj) nie jedzą ludzi.

Chociaż ten dokument jest specyficzny dla architektury, może zawierać również odniesienia do innych architektur. Duża część Podręcznika Gentoo używa tekstu, który jest identyczny dla wszystkich architektur (aby uniknąć powielania). Takie odniesienia zostały ograniczone do minimum, aby uniknąć nieporozumień.

Jeśli istnieje niepewność, czy problem jest problemem użytkownika (błąd popełniony pomimo dokładnego przeczytania dokumentacji), czy wystąpił problem z oprogramowaniem (błąd, który popełniliśmy pomimo dokładnego przetestowania instalacji/dokumentacji), zapraszamy na międzynarodowy kanał [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)) (oraz polski kanał [#gentoo-pl](ircs://irc.libera.chat/#gentoo-pl) ([webchat](https://web.libera.chat/#gentoo-pl))) w sieci irc.freenode.net. Oczywiście wszyscy są mile widziani, ponieważ nasz kanał czatu obejmuje szerokie spektrum Gentoo.

A propos, jeśli masz dodatkowe pytania dotyczące Gentoo, zajrzyj do artykułu [Najczęściej Zadawane Pytania](/wiki/FAQ/pl "FAQ/pl"). Istnieje również wątek [Najczęściej Zadawane Pytania](https://forums.gentoo.org/viewforum.php?f=40) na [Forum Gentoo](https://forums.gentoo.org).

[←](/wiki/Handbook:PPC64 "Previous part") [Home](/wiki/Handbook:PPC64 "Handbook:PPC64") [Choosing the media →](/wiki/Handbook:PPC64/Installation/Media "Next part")