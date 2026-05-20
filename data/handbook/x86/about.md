# About

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/About/de "Handbuch:X86/Installation/Über (100% translated)")
- English
- [Türkçe](/wiki/Handbook:X86/Installation/About/tr "Handbook:X86/Installation/About/tr (0% translated)")
- [español](/wiki/Handbook:X86/Installation/About/es "Manual:X86/Instalación/AcercaDe (100% translated)")
- [français](/wiki/Handbook:X86/Installation/About/fr "Handbook:X86/Installation/About/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/About/it "Handbook:X86/Installation/About (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/About/hu "Handbook:X86/Installation/About/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/About/pl "Handbook:X86/Installation/About (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/About/pt-br "Manual:X86/Instalação/Sobre (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/About/cs "Handbook:X86/Installation/About/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/About/ru "Handbook:X86/Installation/About (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/About/ta "கையேடு:X86/நிறுவல்/பற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/About/zh-cn "手册：X86/安装/关于 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/About/ja "ハンドブック:X86/インストール/About (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/About/ko "Handbook:X86/Installation/About/ko (100% translated)")

[X86 Handbook](/wiki/Handbook:X86 "Handbook:X86")[Installation](/wiki/Handbook:X86/Full/Installation "Handbook:X86/Full/Installation")About the installation[Choosing the media](/wiki/Handbook:X86/Installation/Media "Handbook:X86/Installation/Media")[Configuring the network](/wiki/Handbook:X86/Installation/Networking "Handbook:X86/Installation/Networking")[Preparing the disks](/wiki/Handbook:X86/Installation/Disks "Handbook:X86/Installation/Disks")[The stage file](/wiki/Handbook:X86/Installation/Stage "Handbook:X86/Installation/Stage")[Installing base system](/wiki/Handbook:X86/Installation/Base "Handbook:X86/Installation/Base")[Configuring the kernel](/wiki/Handbook:X86/Installation/Kernel "Handbook:X86/Installation/Kernel")[Configuring the system](/wiki/Handbook:X86/Installation/System "Handbook:X86/Installation/System")[Installing tools](/wiki/Handbook:X86/Installation/Tools "Handbook:X86/Installation/Tools")[Configuring the bootloader](/wiki/Handbook:X86/Installation/Bootloader "Handbook:X86/Installation/Bootloader")[Finalizing](/wiki/Handbook:X86/Installation/Finalizing "Handbook:X86/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:X86/Full/Working "Handbook:X86/Full/Working")[Portage introduction](/wiki/Handbook:X86/Working/Portage "Handbook:X86/Working/Portage")[USE flags](/wiki/Handbook:X86/Working/USE "Handbook:X86/Working/USE")[Portage features](/wiki/Handbook:X86/Working/Features "Handbook:X86/Working/Features")[Initscript system](/wiki/Handbook:X86/Working/Initscripts "Handbook:X86/Working/Initscripts")[Environment variables](/wiki/Handbook:X86/Working/EnvVar "Handbook:X86/Working/EnvVar")[Working with Portage](/wiki/Handbook:X86/Full/Portage "Handbook:X86/Full/Portage")[Files and directories](/wiki/Handbook:X86/Portage/Files "Handbook:X86/Portage/Files")[Variables](/wiki/Handbook:X86/Portage/Variables "Handbook:X86/Portage/Variables")[Mixing software branches](/wiki/Handbook:X86/Portage/Branches "Handbook:X86/Portage/Branches")[Additional tools](/wiki/Handbook:X86/Portage/Tools "Handbook:X86/Portage/Tools")[Custom package repository](/wiki/Handbook:X86/Portage/CustomTree "Handbook:X86/Portage/CustomTree")[Advanced features](/wiki/Handbook:X86/Portage/Advanced "Handbook:X86/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:X86/Full/Networking "Handbook:X86/Full/Networking")[Getting started](/wiki/Handbook:X86/Networking/Introduction "Handbook:X86/Networking/Introduction")[Advanced configuration](/wiki/Handbook:X86/Networking/Advanced "Handbook:X86/Networking/Advanced")[Modular networking](/wiki/Handbook:X86/Networking/Modular "Handbook:X86/Networking/Modular")[Wireless](/wiki/Handbook:X86/Networking/Wireless "Handbook:X86/Networking/Wireless")[Adding functionality](/wiki/Handbook:X86/Networking/Extending "Handbook:X86/Networking/Extending")[Dynamic management](/wiki/Handbook:X86/Networking/Dynamic "Handbook:X86/Networking/Dynamic")

## Contents

- [1Introduction](#Introduction)
  - [1.1Welcome](#Welcome)
    - [1.1.1Openness](#Openness)
    - [1.1.2Choice](#Choice)
    - [1.1.3Power](#Power)
  - [1.2How the installation is structured](#How_the_installation_is_structured)
    - [1.2.1Deciding which steps to take](#Deciding_which_steps_to_take)
    - [1.2.2Suggested steps](#Suggested_steps)
    - [1.2.3Optional steps](#Optional_steps)
    - [1.2.4Deprecated steps](#Deprecated_steps)
    - [1.2.5Defaults and alternatives](#Defaults_and_alternatives)
  - [1.3Installation options for Gentoo](#Installation_options_for_Gentoo)
  - [1.4Troubles](#Troubles)

## Introduction\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-1 "Edit section: ")\]

### Welcome\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-2 "Edit section: ")\]

Welcome to Gentoo! Gentoo is a free operating system based on Linux that can be automatically optimized and customized for just about any application or need. It is built on an ecosystem of free software and does not hide what is running beneath the hood from its users.

#### Openness\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-3 "Edit section: ")\]

Gentoo's premier tools are built from simple programming languages. [Portage](/wiki/Portage "Portage"), Gentoo's package maintenance system, is [written in Python](https://gitweb.gentoo.org/proj/portage.git/). Ebuilds, which provide package definitions for Portage [are written in bash](https://gitweb.gentoo.org/repo/gentoo.git). Our users are encouraged to review, modify, and enhance the source code for all parts of Gentoo.

By default, packages are only patched when necessary to fix bugs or provide interoperability within Gentoo. They are installed to the system by compiling source code provided by upstream projects into binary format (although support for precompiled binary packages is included too). Configuring Gentoo happens through text files.

For the above reasons and others: _openness_ is built-in as a _design principle_.

#### Choice\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-4 "Edit section: ")\]

_Choice_ is another Gentoo _design principle._

When installing Gentoo, choice is made clear throughout the Handbook. System administrators can choose two fully supported init systems (Gentoo's own [OpenRC](/wiki/OpenRC "OpenRC") and Freedesktop.org's [systemd](/wiki/Systemd "Systemd")), partition structure for storage disk(s), what file systems to use on the disk(s), a target [system profile](/wiki/Profile "Profile"), remove or add features on a global (system-wide) or package specific level via USE flags, bootloader, network management utility, and much, much more.

As a development philosophy, [Gentoo's authors](https://www.gentoo.org/inside-gentoo/developers/) try to avoid forcing users onto a specific system profile or desktop environment. If something is offered in the GNU/Linux ecosystem, it's likely available in Gentoo. If not, then we'd love to see it so. For new packages, it is recommended to first submit a package to [GURU](/wiki/GURU "GURU"). Once it has matured and a Gentoo developer has volunteered to sponsor the new package, it can then be added to the official Gentoo package repository.

#### Power\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-5 "Edit section: ")\]

Being a source-based operating system allows Gentoo to be ported onto new computer [instruction set architectures](https://en.wikipedia.org/wiki/instruction_set_architecture "wikipedia:instruction set architecture") and also allows all installed packages to be tuned. This strength surfaces another Gentoo _design principle_: _power_.

A system administrator who has successfully installed and customized Gentoo has compiled a tailored operating system from source code. The entire operating system can be tuned at a binary level via the mechanisms included in Portage's [make.conf](/wiki//etc/portage/make.conf "/etc/portage/make.conf") file. If so desired, adjustments can be made on a per-package basis, or a package group basis. In fact, entire sets of functionality can be added or removed using USE flags.

It is very important that the Handbook reader understands that these design principles are what makes Gentoo unique. With the principles of great power, many choices, and extreme openness highlighted, diligence, thought, and intentionality should be employed while using Gentoo.

### How the installation is structured\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-6 "Edit section: ")\]

The Gentoo Installation can be seen as a 10-step procedure, corresponding to the next set of chapters. Each step results in a certain state:

StepResult
1The user is in a working environment ready to install Gentoo.
2The Internet connection is ready to install Gentoo.
3The hard disks are initialized to host the Gentoo installation.
4The installation environment is prepared and the user is ready to [chroot](/wiki/Chroot "Chroot") into the new environment.
5Core packages, which are the same on all Gentoo installations, are installed.
6The Linux kernel is installed.
7Most of the Gentoo system configuration files are created.
8The necessary system tools are installed.
9The proper boot loader has been installed and configured.
10The freshly installed Gentoo Linux environment is ready to be explored.

#### Deciding which steps to take\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-7 "Edit section: ")\]

The handbook presents an overwhelming amount of options, especially for someone who has never installed Linux without an installer.

It's important to realize that the handbook is designed to describe the steps required to install on a very wide variety of hardware, with different install needs. Because of this, many options presented in the handbook are unnecessary for a particular install and can be skipped.

#### Suggested steps\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-8 "Edit section: ")\]

Prefixed with " **Suggested:**", some steps are not strictly required, but help in most cases, such as installing [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware).

#### Optional steps\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-9 "Edit section: ")\]

Prefixed with " **Optional:**", many sections in the handbook are purely optional, and can be skipped if the user is looking for a simple, mostly vanilla install.

Examples of this include compiler flag customization, using a totally custom kernel, and disabling root login.

**Tip**

When following optional steps, it's important to be careful that all prerequisites were satisfied. Some optional steps depend on other optional steps.

#### Deprecated steps\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-10 "Edit section: ")\]

Gentoo has existed for a long time. Some install processes are described in the handbook because they used to be more relevant, but are now largely deprecated. Instead of immediately removing this information, as it may still be helpful for some users, it may be designated as **Deprecated:** before removal. Once removed, the _history_ must be used to view this content.

#### Defaults and alternatives\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-11 "Edit section: ")\]

Whenever choices are presented, the handbook will try to explain the pros and cons of each choice.

If potential choices are mutually exclusive, " **Default:**" is used to designate the most supported or commonly chosen option, while alternatives are listed with " **Alternative**".

**Important**

**Alternative** options are not inferior to **Default** s, but **Default** options are typically more widely used and may have better support.

### Installation options for Gentoo\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-12 "Edit section: ")\]

Gentoo can be installed in many different ways. It can be downloaded and installed from official Gentoo installation media such as our bootable ISO images. The installation media can be installed on a USB stick or accessed via a netbooted environment. Alternatively, Gentoo can be installed from non-official media such as an already installed distribution or a non-Gentoo bootable disk (such as [Linux Mint](https://linuxmint.com/)).

This document covers the installation using official Gentoo Installation media or, in certain cases, netbooting.

**Note**

For help on the other installation approaches, including using non-Gentoo bootable media, please read our [Alternative installation guide](/wiki/Installation_alternatives "Installation alternatives").

We also provide a [Gentoo installation tips and tricks](/wiki/Gentoo_installation_tips_and_tricks "Gentoo installation tips and tricks") document that might be useful.

### Troubles\[ [edit](/index.php?title=Handbook:Parts/Installation/About&action=edit&section=T-13 "Edit section: ")\]

If a problem is found in the installation (or in the installation documentation), please visit our [bug tracking system](https://bugs.gentoo.org) and check if the bug is known. If not, please create a bug report for it so we can take care of it. Do not be afraid of the developers who are assigned to the bugs - they (generally) don't eat people.

Although this document is architecture-specific, it may contain references to other architectures as well, because large parts of the Gentoo Handbook use text that is identical for all architectures (to avoid duplication of effort). Such references have been kept to a minimum, to avoid confusion.

If there is some uncertainty about whether or not the problem is a user-problem (some error made despite having read the documentation carefully) or a software-problem (some error we made despite having tested the installation/documentation carefully) everybody is welcome to join the [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)) channel on irc.libera.chat. Of course, everyone is welcome otherwise too as our chat channel covers the broad Gentoo spectrum.

Speaking of which, if there are any additional questions regarding Gentoo, check out the [Frequently Asked Questions](/wiki/FAQ "FAQ") article. There are also [FAQs](https://forums.gentoo.org/viewforum.php?f=40) on the [Gentoo Forums](https://forums.gentoo.org).

[←](/wiki/Handbook:X86 "Previous part") [Home](/wiki/Handbook:X86 "Handbook:X86") [Choosing the media →](/wiki/Handbook:X86/Installation/Media "Next part")