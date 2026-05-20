# Finalizing

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Finalizing/de "Handbuch:Alpha/Installation/Abschluss (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Finalizing "Handbook:Alpha/Installation/Finalizing (100% translated)")
- español
- [français](/wiki/Handbook:Alpha/Installation/Finalizing/fr "Manuel:Alpha/Installation/Finalisation (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Finalizing/it "Handbook:Alpha/Installation/Finalizing/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Finalizing/hu "Handbook:Alpha/Installation/Finalizing/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Finalizing/pl "Handbook:Alpha/Installation/Finalizing (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Finalizing/pt-br "Manual:Alpha/Instalação/Finalizando (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Finalizing/cs "Handbook:Alpha/Installation/Finalizing/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Finalizing/ru "Handbook:Alpha/Installation/Finalizing (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Finalizing/ta "கையேடு:Alpha/நிறுவல்/முடித்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Finalizing/zh-cn "手册：Alpha/安装/收尾安装工作 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Finalizing/ja "ハンドブック:Alpha/インストール/ファイナライズ (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Finalizing/ko "Handbook:Alpha/Installation/Finalizing/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha "Handbook:Alpha")[Installation](/wiki/Handbook:Alpha/Full/Installation "Handbook:Alpha/Full/Installation")[About the installation](/wiki/Handbook:Alpha/Installation/About/es "Handbook:Alpha/Installation/About/es")[Choosing the media](/wiki/Handbook:Alpha/Installation/Media/es "Handbook:Alpha/Installation/Media/es")[Configuring the network](/wiki/Handbook:Alpha/Installation/Networking/es "Handbook:Alpha/Installation/Networking/es")[Preparing the disks](/wiki/Handbook:Alpha/Installation/Disks/es "Handbook:Alpha/Installation/Disks/es")[The stage file](/wiki/Handbook:Alpha/Installation/Stage/es "Handbook:Alpha/Installation/Stage/es")[Installing base system](/wiki/Handbook:Alpha/Installation/Base/es "Handbook:Alpha/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:Alpha/Installation/Kernel/es "Handbook:Alpha/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:Alpha/Installation/System/es "Handbook:Alpha/Installation/System/es")[Installing tools](/wiki/Handbook:Alpha/Installation/Tools/es "Handbook:Alpha/Installation/Tools/es")[Configuring the bootloader](/wiki/Handbook:Alpha/Installation/Bootloader/es "Handbook:Alpha/Installation/Bootloader/es")Finalizing[Working with Gentoo](/wiki/Handbook:Alpha/Full/Working "Handbook:Alpha/Full/Working")[Portage introduction](/wiki/Handbook:Alpha/Working/Portage/es "Handbook:Alpha/Working/Portage/es")[USE flags](/wiki/Handbook:Alpha/Working/USE/es "Handbook:Alpha/Working/USE/es")[Portage features](/wiki/Handbook:Alpha/Working/Features/es "Handbook:Alpha/Working/Features/es")[Initscript system](/wiki/Handbook:Alpha/Working/Initscripts/es "Handbook:Alpha/Working/Initscripts/es")[Environment variables](/wiki/Handbook:Alpha/Working/EnvVar/es "Handbook:Alpha/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:Alpha/Full/Portage "Handbook:Alpha/Full/Portage")[Files and directories](/wiki/Handbook:Alpha/Portage/Files/es "Handbook:Alpha/Portage/Files/es")[Variables](/wiki/Handbook:Alpha/Portage/Variables/es "Handbook:Alpha/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:Alpha/Portage/Branches/es "Handbook:Alpha/Portage/Branches/es")[Additional tools](/wiki/Handbook:Alpha/Portage/Tools/es "Handbook:Alpha/Portage/Tools/es")[Custom package repository](/wiki/Handbook:Alpha/Portage/CustomTree/es "Handbook:Alpha/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:Alpha/Portage/Advanced/es "Handbook:Alpha/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:Alpha/Full/Networking "Handbook:Alpha/Full/Networking")[Getting started](/wiki/Handbook:Alpha/Networking/Introduction/es "Handbook:Alpha/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:Alpha/Networking/Advanced/es "Handbook:Alpha/Networking/Advanced/es")[Modular networking](/wiki/Handbook:Alpha/Networking/Modular/es "Handbook:Alpha/Networking/Modular/es")[Wireless](/wiki/Handbook:Alpha/Networking/Wireless/es "Handbook:Alpha/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:Alpha/Networking/Extending/es "Handbook:Alpha/Networking/Extending/es")[Dynamic management](/wiki/Handbook:Alpha/Networking/Dynamic/es "Handbook:Alpha/Networking/Dynamic/es")

## Contents

- [1Portage maintenance](#Portage_maintenance)
- [2User administration](#User_administration)
  - [2.1Adding a user for daily use](#Adding_a_user_for_daily_use)
  - [2.2Temporarily elevating privileges](#Temporarily_elevating_privileges)
  - [2.3Disabling root login](#Disabling_root_login)
- [3Disk cleanup](#Disk_cleanup)
  - [3.1Removing installation artifacts](#Removing_installation_artifacts)
- [4Where to go from here](#Where_to_go_from_here)
  - [4.1Additional documentation](#Additional_documentation)
  - [4.2Gentoo online](#Gentoo_online)
    - [4.2.1Forums and IRC](#Forums_and_IRC)
    - [4.2.2Mailing lists](#Mailing_lists)
    - [4.2.3Bugs](#Bugs)
    - [4.2.4Development guide](#Development_guide)
- [5Closing thoughts](#Closing_thoughts)

## Portage maintenance

Portage by default preserves copies of downloaded files, on local storage: source tarballs in /var/cache/distfiles and binary packages in /var/cache/binhost/gentoo. If an update downloads a newer version of these files, the earlier versions are still preserved.

It's good practice on a Gentoo system to keep the current versions of these source tarballs and binary packages, to be able to recover systems, in case of file corruption for example (after a certain amount of time, there isn't always a guarantee that these files will still be available for download).

Keeping the older, non current, versions of these files however can needlessly consume storage space, adding up over time, and potentially wasting space, if the older versions are not required.

Gentoo provides the [gentoolkit](/wiki/Gentoolkit/es "Gentoolkit/es") suite of tools to help with system administration, which notably provides the [eclean](/wiki/Eclean/es "Eclean/es") tool to easily remove old versions of downloaded source tarballs and binary packages, to prevent storage requirements from growing indefinitely.

First install [app-portage/gentoolkit](https://packages.gentoo.org/packages/app-portage/gentoolkit):

`root #` `emerge --ask app-portage/gentoolkit`

For cleaning old source tarballs, use:

`root #` `eclean-dist`

For cleaning old binpkgs, use:

`root #` `eclean-pkg`

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

If a user ever needs to perform some task as root, they can use su - to temporarily receive root privileges. Another way is to use the [sudo](/wiki/Sudo/es "Sudo/es") ([app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo)) or [doas](/wiki/Doas "Doas") ([app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas)) utilities which are, if correctly configured, very secure.

### Disabling root login

**Advertencia**

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

Not sure where to go from here? There are many paths to explore... Gentoo provides its users with lots of possibilities and therefore has lots of documented (and less documented) features to explore here on the wiki and on other Gentoo related sub-domains (see the [Gentoo online](/wiki/Handbook:Alpha/Installation/Finalizing/es#Gentoo_online "Handbook:Alpha/Installation/Finalizing/es") section below).

### Additional documentation

It is important to note that, due to the number of choices available in Gentoo, the documentation provided by the handbook is limited in scope - it mainly focuses on the basics of getting a Gentoo system up and running and basic system management activities. The handbook intentionally excludes instructions on graphical environments, details on hardening, and other important administrative tasks. That being stated, there are more sections of the handbook to assist readers with more basic functions.

Readers should definitely take a look at the next part of the handbook entitled [Working with Gentoo](/wiki/Handbook:Alpha/Working/Portage/es "Handbook:Alpha/Working/Portage/es") which explains how to keep the software up to date, install additional software packages, details on USE flags, the OpenRC init system, and various other informative topics relating to managing a Gentoo system post-installation.

Apart from the handbook, readers should also feel encouraged to explore other corners of the Gentoo wiki to find additional, community-provided documentation. The Gentoo wiki team also offers a [documentation topic overview](/wiki/Main_Page#Documentation_topics "Main Page") which lists a selection of wiki articles by category. For instance, it refers to the [localization guide](/wiki/Localization/Guide/es "Localization/Guide/es") to make a system feel more at home (particularly useful for users who speak English as a second language).

The majority of users with desktop use cases will setup graphical environments in which to work natively. There are many community maintained 'meta' articles for supported [desktop environments (DEs)](/wiki/Desktop_environment "Desktop environment") and [window managers (WMs)](/wiki/Window_manager "Window manager"). Readers should be aware that each DE will require slightly different setup steps, which will lengthen add complexity to bootstrapping.

Many other [Meta articles](/wiki/Category:Meta "Category:Meta") exist to provide our readers with high level overviews of available software within Gentoo.

### Gentoo online

**Importante**

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

As a reminder, any feedback for _this handbook_ should follow the guidelines detailed in the [How do I improve the Handbook?](/wiki/Handbook:Main_Page/es#How_do_I_improve_the_Handbook.3F "Handbook:Main Page/es") section at the beginning of the handbook.

We look forward to seeing how our users will choose to implement Gentoo to fit their unique use cases and needs.

[← Configuring the bootloader](/wiki/Handbook:Alpha/Installation/Bootloader "Previous part") [Home](/wiki/Handbook:Alpha "Handbook:Alpha") [→](/wiki/Handbook:Alpha/Working/Portage "Next part")