# Finalizing

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Finalizing/de "Handbuch:SPARC/Installation/Abschluss (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Finalizing "Handbook:SPARC/Installation/Finalizing (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Finalizing/es "Manual de Gentoo: SPARC/Instalación/Finalizar (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Finalizing/fr "Handbook:SPARC/Installation/Finalizing/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Finalizing/it "Handbook:SPARC/Installation/Finalizing/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Finalizing/hu "Handbook:SPARC/Installation/Finalizing/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Finalizing/pt-br "Handbook:SPARC/Installation/Finalizing/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Finalizing/ru "Handbook:SPARC/Installation/Finalizing (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Finalizing/ta "கையேடு:SPARC/நிறுவல்/முடித்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Finalizing/zh-cn "手册：SPARC/安装/收尾安装工作 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Finalizing/ja "ハンドブック:SPARC/インストール/ファイナライズ (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Finalizing/ko "Handbook:SPARC/Installation/Finalizing/ko (100% translated)")

[SPARC Handbook](/wiki/Handbook:SPARC/pl "Handbook:SPARC/pl")[Instalacja](/wiki/Handbook:SPARC/Full/Installation/pl "Handbook:SPARC/Full/Installation/pl")[O instalacji](/wiki/Handbook:SPARC/Installation/About/pl "Handbook:SPARC/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:SPARC/Installation/Media/pl "Handbook:SPARC/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:SPARC/Installation/Networking/pl "Handbook:SPARC/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:SPARC/Installation/Disks/pl "Handbook:SPARC/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:SPARC/Installation/Stage/pl "Handbook:SPARC/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:SPARC/Installation/Base/pl "Handbook:SPARC/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:SPARC/Installation/Kernel/pl "Handbook:SPARC/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:SPARC/Installation/System/pl "Handbook:SPARC/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:SPARC/Installation/Tools/pl "Handbook:SPARC/Installation/Tools/pl")[Instalacja systemu rozruchowego](/wiki/Handbook:SPARC/Installation/Bootloader/pl "Handbook:SPARC/Installation/Bootloader/pl")Finalizacja[Praca z Gentoo](/wiki/Handbook:SPARC/Full/Working/pl "Handbook:SPARC/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:SPARC/Working/Portage/pl "Handbook:SPARC/Working/Portage/pl")[Flagi USE](/wiki/Handbook:SPARC/Working/USE/pl "Handbook:SPARC/Working/USE/pl")[Funkcje portage](/wiki/Handbook:SPARC/Working/Features/pl "Handbook:SPARC/Working/Features/pl")[System initscript](/wiki/Handbook:SPARC/Working/Initscripts/pl "Handbook:SPARC/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:SPARC/Working/EnvVar/pl "Handbook:SPARC/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:SPARC/Full/Portage/pl "Handbook:SPARC/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:SPARC/Portage/Files/pl "Handbook:SPARC/Portage/Files/pl")[Zmienne](/wiki/Handbook:SPARC/Portage/Variables/pl "Handbook:SPARC/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:SPARC/Portage/Branches/pl "Handbook:SPARC/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:SPARC/Portage/Tools/pl "Handbook:SPARC/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:SPARC/Portage/CustomTree/pl "Handbook:SPARC/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:SPARC/Portage/Advanced/pl "Handbook:SPARC/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:SPARC/Full/Networking/pl "Handbook:SPARC/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:SPARC/Networking/Introduction/pl "Handbook:SPARC/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:SPARC/Networking/Advanced/pl "Handbook:SPARC/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:SPARC/Networking/Modular/pl "Handbook:SPARC/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:SPARC/Networking/Wireless/pl "Handbook:SPARC/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:SPARC/Networking/Extending/pl "Handbook:SPARC/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:SPARC/Networking/Dynamic/pl "Handbook:SPARC/Networking/Dynamic/pl")

## Contents

- [1User administration](#User_administration)
  - [1.1Adding a user for daily use](#Adding_a_user_for_daily_use)
  - [1.2Temporarily elevating privileges](#Temporarily_elevating_privileges)
  - [1.3Disabling root login](#Disabling_root_login)
- [2Disk cleanup](#Disk_cleanup)
  - [2.1Removing installation artifacts](#Removing_installation_artifacts)
- [3Where to go from here](#Where_to_go_from_here)
  - [3.1Additional documentation](#Additional_documentation)
  - [3.2Gentoo online](#Gentoo_online)
    - [3.2.1Forums and IRC](#Forums_and_IRC)
    - [3.2.2Mailing lists](#Mailing_lists)
    - [3.2.3Bugs](#Bugs)
    - [3.2.4Development guide](#Development_guide)
- [4Closing thoughts](#Closing_thoughts)

## User administration

### Adding a user for daily use

Working as root on a Unix/Linux system is dangerous and should be avoided as much as possible. Therefore it is strongly recommended to add one or more standard user account(s) for day-to-day use.

The groups the user is member of define what activities the user can perform. The following table lists a number of important groups:

Group
Description
audio
Enable the user account to access the audio devices.
cdrom
Enable the user account to directly access optical devices.
cron
Enable the user account to access time-based job scheduling via cron. Note: user accounts on systems running the systemd service system can use systemd timers and user service files instead of cron jobs.
floppy
Enable the user account to directly access ancient mechanical devices known as floppy drives. This group is not generally used on modern systems.
usb
Enable the user account able to access USB devices.
video
Enables the user account to access video capturing hardware and hardware acceleration.
wheel
Enables the user account able to use the su ( _substitute user_) command, which allows switching to the root account or other accounts. For single user systems that include a root account, it is a good idea to add this group for the primary standard user.

For instance, to create a user called [larry](/wiki/User:Larry "User:Larry") who is a member of the _wheel_, _users_, and _audio_ groups, log in as root first (only root can create users) and run useradd:

`Login:` `root`

```
Password: (Enter the root password)

```

When setting passwords for standard user accounts, it is good security practice to avoid using the same or a similar password as set for the root user.

Handbook authors recommended to use a password at least 16 characters in length, with a value fully unique from every other user on the system.

`root #` `useradd -m -G users,wheel,audio -s /bin/bash larry
`

`root #` `passwd larry`

```
Password: (Enter the password for larry)
Re-enter password: (Re-enter the password to verify)

```

### Temporarily elevating privileges

If a user ever needs to perform some task as root, they can use su - to temporarily receive root privileges. Another way is to use the [sudo](/wiki/Sudo "Sudo") ([app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo)) or [doas](/wiki/Doas "Doas") ([app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas)) utilities which are, if correctly configured, very secure.

### Disabling root login

**Uwaga**

Before disabling the root login, ensure that a user account is a member of the wheel group and that a method to elevate user privilege exists; otherwise root access will be locked and system administration will be impossible without performing recovery. Some common methods to elevate user privilege include: [app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo), [app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas), or systemd's run0.

To prevent possible threat actors from logging in as root, deleting the root password and/or disabling root login can help improve security.

To disable root login:

`root #` `passwd -l root`

To delete the root password and disable login:

`root #` `passwd -dl root`

## Disk cleanup

### Removing installation artifacts

With the Gentoo installation finished and the system rebooted, if everything has gone well, the stage file and other installation artifacts - such as DIGEST, CONTENT, or \*.asc (PGP signature) files - can now be safely removed.

The files are located in the / directory and can be removed with the following command:

`root #` `rm /stage3-*.tar.*`

## Where to go from here

Not sure where to go from here? There are many paths to explore... Gentoo provides its users with lots of possibilities and therefore has lots of documented (and less documented) features to explore here on the wiki and on other Gentoo related sub-domains (see the [Gentoo online](/wiki/Handbook:SPARC/Installation/Finalizing/pl#Gentoo_online "Handbook:SPARC/Installation/Finalizing/pl") section below).

### Additional documentation

It is important to note that, due to the number of choices available in Gentoo, the documentation provided by the handbook is limited in scope - it mainly focuses on the basics of getting a Gentoo system up and running and basic system management activities. The handbook intentionally excludes instructions on graphical environments, details on hardening, and other important administrative tasks. That being stated, there are more sections of the handbook to assist readers with more basic functions.

Readers should definitely take a look at the next part of the handbook entitled [Working with Gentoo](/wiki/Handbook:SPARC/Working/Portage/pl "Handbook:SPARC/Working/Portage/pl") which explains how to keep the software up to date, install additional software packages, details on USE flags, the OpenRC init system, and various other informative topics relating to managing a Gentoo system post-installation.

Apart from the handbook, readers should also feel encouraged to explore other corners of the Gentoo wiki to find additional, community-provided documentation. The Gentoo wiki team also offers a [documentation topic overview](/wiki/Main_Page#Documentation_topics "Main Page") which lists a selection of wiki articles by category. For instance, it refers to the [localization guide](/wiki/Localization/Guide "Localization/Guide") to make a system feel more at home (particularly useful for users who speak English as a second language).

The majority of users with desktop use cases will setup graphical environments in which to work natively. There are many community maintained 'meta' articles for supported [desktop environments (DEs)](/wiki/Desktop_environment "Desktop environment") and [window managers (WMs)](/wiki/Window_manager "Window manager"). Readers should be aware that each DE will require slightly different setup steps, which will lengthen add complexity to bootstrapping.

Many other [Meta articles](/wiki/Category:Meta "Category:Meta") exist to provide our readers with high level overviews of available software within Gentoo.

### Gentoo online

**Ważne**

Readers should note that all official Gentoo sites online are governed by Gentoo's [code of conduct](/wiki/Project:Council/Code_of_conduct "Project:Council/Code of conduct"). Being active in the Gentoo community is a privilege, not a right, and users should be aware that the code of conduct exists for a reason.

With the exception of the Libera.Chat hosted internet relay chat (IRC) network and the mailing lists, most Gentoo websites require an account on a per site basis in order to ask questions, open a discussion, or enter a bug.

#### Forums and IRC

Every user is welcome on our [Gentoo forums](https://forums.gentoo.org) or on one of our [internet relay chat channels](https://www.gentoo.org/get-involved/irc-channels/). It is easy to search for the forums to see if an issue experienced on a fresh Gentoo install has been discovered in the past and resolved after some feedback. The likelihood of other users experiencing the installation issues by first-time Gentoo can be surprising. It is advised users search the forums and the wiki before asking for assistance in Gentoo support channels.

#### Mailing lists

[Several mailing lists](https://www.gentoo.org/get-involved/mailing-lists/) are available to the community members who prefer to ask for support or feedback over email rather than create a user account on the forums or IRC. Users will need to follow the instructions in order to subscribe to specific mailing lists.

#### Bugs

Sometimes after reviewing the wiki, searching the forums, and seeking support in the IRC channel or mailing lists there is no known solution to a problem. Generally this is a sign to open a bug on Gentoo's [Bugzilla site](https://bugs.gentoo.org).

#### Development guide

Readers who desire to learn more about developing Gentoo can take a look at the [Development guide](https://devmanual.gentoo.org/). This guide provides instructions on writing ebuilds, working with eclasses, and provides definitions for many [general concepts](https://devmanual.gentoo.org/general-concepts/index.html) behind Gentoo development.

## Closing thoughts

Gentoo is a robust, flexible, and excellently maintained distribution. The developer community is happy to hear feedback on how to make Gentoo an even _better_ distribution.

As a reminder, any feedback for _this handbook_ should follow the guidelines detailed in the [How do I improve the Handbook?](/wiki/Handbook:Main_Page/pl#How_do_I_improve_the_Handbook.3F "Handbook:Main Page/pl") section at the beginning of the handbook.

We look forward to seeing how our users will choose to implement Gentoo to fit their unique use cases and needs.

[← Configuring the bootloader](/wiki/Handbook:SPARC/Installation/Bootloader "Previous part") [Home](/wiki/Handbook:SPARC "Handbook:SPARC") [→](/wiki/Handbook:SPARC/Working/Portage "Next part")