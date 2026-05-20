# Finalizing

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Finalizing/de "Handbuch:X86/Installation/Abschluss (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Finalizing "Handbook:X86/Installation/Finalizing (100% translated)")
- [español](/wiki/Handbook:X86/Installation/Finalizing/es "Manual de Gentoo: X86/Instalación/Finalización (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Finalizing/fr "Handbook:X86/Installation/Finalizing/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Finalizing/it "Handbook:X86/Installation/Finalizing (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/Finalizing/hu "Handbook:X86/Installation/Finalizing/hu (100% translated)")
- polski
- [português](/wiki/Handbook:X86/Installation/Finalizing/pt "Manual:X86/Instalação/Finalizando (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Finalizing/pt-br "Manual:X86/Instalação/Finalizando (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Finalizing/cs "Handbook:X86/Installation/Finalizing/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Finalizing/ru "Handbook:X86/Installation/Finalizing (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Finalizing/ta "கையேடு:X86/நிறுவல்/முடித்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Finalizing/zh-cn "手册：X86/安装/收尾安装工作 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Finalizing/ja "ハンドブック:X86/インストール/インストールの締めくくり (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Finalizing/ko "Handbook:X86/Installation/Finalizing/ko (100% translated)")

[X86 Handbook](/wiki/Handbook:X86/pl "Handbook:X86/pl")[Instalacja](/wiki/Handbook:X86/Full/Installation/pl "Handbook:X86/Full/Installation/pl")[O instalacji](/wiki/Handbook:X86/Installation/About/pl "Handbook:X86/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:X86/Installation/Media/pl "Handbook:X86/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:X86/Installation/Networking/pl "Handbook:X86/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:X86/Installation/Disks/pl "Handbook:X86/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:X86/Installation/Stage/pl "Handbook:X86/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:X86/Installation/Base/pl "Handbook:X86/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:X86/Installation/Kernel/pl "Handbook:X86/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:X86/Installation/System/pl "Handbook:X86/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:X86/Installation/Tools/pl "Handbook:X86/Installation/Tools/pl")[Instalacja systemu rozruchowego](/wiki/Handbook:X86/Installation/Bootloader/pl "Handbook:X86/Installation/Bootloader/pl")Finalizacja[Praca z Gentoo](/wiki/Handbook:X86/Full/Working/pl "Handbook:X86/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:X86/Working/Portage/pl "Handbook:X86/Working/Portage/pl")[Flagi USE](/wiki/Handbook:X86/Working/USE/pl "Handbook:X86/Working/USE/pl")[Funkcje portage](/wiki/Handbook:X86/Working/Features/pl "Handbook:X86/Working/Features/pl")[System initscript](/wiki/Handbook:X86/Working/Initscripts/pl "Handbook:X86/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:X86/Working/EnvVar/pl "Handbook:X86/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:X86/Full/Portage/pl "Handbook:X86/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:X86/Portage/Files/pl "Handbook:X86/Portage/Files/pl")[Zmienne](/wiki/Handbook:X86/Portage/Variables/pl "Handbook:X86/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:X86/Portage/Branches/pl "Handbook:X86/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:X86/Portage/Tools/pl "Handbook:X86/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:X86/Portage/CustomTree/pl "Handbook:X86/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:X86/Portage/Advanced/pl "Handbook:X86/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:X86/Full/Networking/pl "Handbook:X86/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:X86/Networking/Introduction/pl "Handbook:X86/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:X86/Networking/Advanced/pl "Handbook:X86/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:X86/Networking/Modular/pl "Handbook:X86/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:X86/Networking/Wireless/pl "Handbook:X86/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:X86/Networking/Extending/pl "Handbook:X86/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:X86/Networking/Dynamic/pl "Handbook:X86/Networking/Dynamic/pl")

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

Not sure where to go from here? There are many paths to explore... Gentoo provides its users with lots of possibilities and therefore has lots of documented (and less documented) features to explore here on the wiki and on other Gentoo related sub-domains (see the [Gentoo online](/wiki/Handbook:X86/Installation/Finalizing/pl#Gentoo_online "Handbook:X86/Installation/Finalizing/pl") section below).

### Additional documentation

It is important to note that, due to the number of choices available in Gentoo, the documentation provided by the handbook is limited in scope - it mainly focuses on the basics of getting a Gentoo system up and running and basic system management activities. The handbook intentionally excludes instructions on graphical environments, details on hardening, and other important administrative tasks. That being stated, there are more sections of the handbook to assist readers with more basic functions.

Readers should definitely take a look at the next part of the handbook entitled [Working with Gentoo](/wiki/Handbook:X86/Working/Portage/pl "Handbook:X86/Working/Portage/pl") which explains how to keep the software up to date, install additional software packages, details on USE flags, the OpenRC init system, and various other informative topics relating to managing a Gentoo system post-installation.

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

[← Configuring the bootloader](/wiki/Handbook:X86/Installation/Bootloader "Previous part") [Home](/wiki/Handbook:X86 "Handbook:X86") [→](/wiki/Handbook:X86/Working/Portage "Next part")