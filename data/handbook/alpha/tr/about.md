# About

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/About/de "Handbuch:Alpha/Installation/Über (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/About "Handbook:Alpha/Installation/About (100% translated)")
- Türkçe
- [español](/wiki/Handbook:Alpha/Installation/About/es "Manual de Gentoo: Alpha/Instalación/Acerca de (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/About/fr "Manuel:Alpha/Installation/À-propos (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/About/it "Handbook:Alpha/Installation/About (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/About/hu "Handbook:Alpha/Installation/About/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/About/pl "Handbook:Alpha/Installation/About (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/About/pt-br "Manual:Alpha/Instalação/Sobre (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/About/cs "Handbook:Alpha/Installation/About/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/About/ru "Handbook:Alpha/Installation/About (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/About/ta "கையேடு:Alpha/நிறுவல்/பற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/About/zh-cn "手册：Alpha/安装/关于 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/About/ja "ハンドブック:Alpha/インストール/About (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/About/ko "Handbook:Alpha/Installation/About/ko (100% translated)")

[Alpha Rehber](/wiki/Handbook:Alpha "Handbook:Alpha")[Kurulum](/wiki/Handbook:Alpha/Full/Installation "Handbook:Alpha/Full/Installation")Gentoo Linux kurulumu hakkında[Kurulum için doğru aracı seçme](/wiki/Handbook:Alpha/Installation/Media/tr "Handbook:Alpha/Installation/Media/tr")[İnternet ağı ayarlamaları](/wiki/Handbook:Alpha/Installation/Networking/tr "Handbook:Alpha/Installation/Networking/tr")[Diski hazırlama](/wiki/Handbook:Alpha/Installation/Disks/tr "Handbook:Alpha/Installation/Disks/tr")[stage3 dosyalarını indirme](/wiki/Handbook:Alpha/Installation/Stage/tr "Handbook:Alpha/Installation/Stage/tr")[Ana sistem dosyalarını yükleme](/wiki/Handbook:Alpha/Installation/Base "Handbook:Alpha/Installation/Base")[Kernel ayarlamaları](/wiki/Handbook:Alpha/Installation/Kernel "Handbook:Alpha/Installation/Kernel")[Sistem ayarlamaları](/wiki/Handbook:Alpha/Installation/System "Handbook:Alpha/Installation/System")[Sistem araçlarının yüklenmesi](/wiki/Handbook:Alpha/Installation/Tools "Handbook:Alpha/Installation/Tools")[Bootloader ayarlamaları](/wiki/Handbook:Alpha/Installation/Bootloader "Handbook:Alpha/Installation/Bootloader")[Son dokunuşlar](/wiki/Handbook:Alpha/Installation/Finalizing "Handbook:Alpha/Installation/Finalizing")[Gentoo](/wiki/Handbook:Alpha/Full/Working "Handbook:Alpha/Full/Working")[Portage paket yöneticisine giriş](/wiki/Handbook:Alpha/Working/Portage "Handbook:Alpha/Working/Portage")[USE flag ayarları](/wiki/Handbook:Alpha/Working/USE "Handbook:Alpha/Working/USE")[Portage özellikleri](/wiki/Handbook:Alpha/Working/Features "Handbook:Alpha/Working/Features")[Initscript sistemi](/wiki/Handbook:Alpha/Working/Initscripts "Handbook:Alpha/Working/Initscripts")[Ortam değişkenleri](/wiki/Handbook:Alpha/Working/EnvVar "Handbook:Alpha/Working/EnvVar")[Portage](/wiki/Handbook:Alpha/Full/Portage "Handbook:Alpha/Full/Portage")[Dosyalar ve dizinler](/wiki/Handbook:Alpha/Portage/Files "Handbook:Alpha/Portage/Files")[Değişkenler](/wiki/Handbook:Alpha/Portage/Variables "Handbook:Alpha/Portage/Variables")[Farklı dalları birlikte kullanmak](/wiki/Handbook:Alpha/Portage/Branches "Handbook:Alpha/Portage/Branches")[Ekstra araçlar](/wiki/Handbook:Alpha/Portage/Tools "Handbook:Alpha/Portage/Tools")[Özel paket depoları](/wiki/Handbook:Alpha/Portage/CustomTree "Handbook:Alpha/Portage/CustomTree")[Gelişmiş özellikler](/wiki/Handbook:Alpha/Portage/Advanced "Handbook:Alpha/Portage/Advanced")[Ağ ayarları](/wiki/Handbook:Alpha/Full/Networking "Handbook:Alpha/Full/Networking")[Başlangıç](/wiki/Handbook:Alpha/Networking/Introduction "Handbook:Alpha/Networking/Introduction")[Gelişmiş yapılandırma](/wiki/Handbook:Alpha/Networking/Advanced "Handbook:Alpha/Networking/Advanced")[Modüler ağ](/wiki/Handbook:Alpha/Networking/Modular "Handbook:Alpha/Networking/Modular")[Kablosuz ağ](/wiki/Handbook:Alpha/Networking/Wireless "Handbook:Alpha/Networking/Wireless")[Özellik ekleme](/wiki/Handbook:Alpha/Networking/Extending "Handbook:Alpha/Networking/Extending")[Dinamik ağ yönetimi](/wiki/Handbook:Alpha/Networking/Dynamic "Handbook:Alpha/Networking/Dynamic")

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

## Introduction

### Welcome

Welcome to Gentoo! Gentoo is a free operating system based on Linux that can be automatically optimized and customized for just about any application or need. It is built on an ecosystem of free software and does not hide what is running beneath the hood from its users.

#### Openness

Gentoo's premier tools are built from simple programming languages. [Portage](/wiki/Portage "Portage"), Gentoo's package maintenance system, is [written in Python](https://gitweb.gentoo.org/proj/portage.git/). Ebuilds, which provide package definitions for Portage [are written in bash](https://gitweb.gentoo.org/repo/gentoo.git). Our users are encouraged to review, modify, and enhance the source code for all parts of Gentoo.

By default, packages are only patched when necessary to fix bugs or provide interoperability within Gentoo. They are installed to the system by compiling source code provided by upstream projects into binary format (although support for precompiled binary packages is included too). Configuring Gentoo happens through text files.

For the above reasons and others: _openness_ is built-in as a _design principle_.

#### Choice

_Choice_ is another Gentoo _design principle._

When installing Gentoo, choice is made clear throughout the Handbook. System administrators can choose two fully supported init systems (Gentoo's own [OpenRC](/wiki/OpenRC "OpenRC") and Freedesktop.org's [systemd](/wiki/Systemd "Systemd")), partition structure for storage disk(s), what file systems to use on the disk(s), a target [system profile](/wiki/Profile "Profile"), remove or add features on a global (system-wide) or package specific level via USE flags, bootloader, network management utility, and much, much more.

As a development philosophy, [Gentoo's authors](https://www.gentoo.org/inside-gentoo/developers/) try to avoid forcing users onto a specific system profile or desktop environment. If something is offered in the GNU/Linux ecosystem, it's likely available in Gentoo. If not, then we'd love to see it so. For new packages, it is recommended to first submit a package to [GURU](/wiki/GURU "GURU"). Once it has matured and a Gentoo developer has volunteered to sponsor the new package, it can then be added to the official Gentoo package repository.

#### Power

Being a source-based operating system allows Gentoo to be ported onto new computer [instruction set architectures](https://en.wikipedia.org/wiki/instruction_set_architecture "wikipedia:instruction set architecture") and also allows all installed packages to be tuned. This strength surfaces another Gentoo _design principle_: _power_.

A system administrator who has successfully installed and customized Gentoo has compiled a tailored operating system from source code. The entire operating system can be tuned at a binary level via the mechanisms included in Portage's [make.conf](/wiki//etc/portage/make.conf "/etc/portage/make.conf") file. If so desired, adjustments can be made on a per-package basis, or a package group basis. In fact, entire sets of functionality can be added or removed using USE flags.

It is very important that the Handbook reader understands that these design principles are what makes Gentoo unique. With the principles of great power, many choices, and extreme openness highlighted, diligence, thought, and intentionality should be employed while using Gentoo.

### How the installation is structured

The Gentoo Installation can be seen as a 10-step procedure, corresponding to the next set of chapters. Each step results in a certain state:

StepResult
1The user is in a working environment ready to install Gentoo.
2The Internet connection is ready to install Gentoo.
3The hard disks are initialized to host the Gentoo installation.
4The installation environment is prepared and the user is ready to [chroot](/wiki/Chroot/tr "Chroot/tr") into the new environment.
5Core packages, which are the same on all Gentoo installations, are installed.
6The Linux kernel is installed.
7Most of the Gentoo system configuration files are created.
8The necessary system tools are installed.
9The proper boot loader has been installed and configured.
10The freshly installed Gentoo Linux environment is ready to be explored.

#### Deciding which steps to take

The handbook presents an overwhelming amount of options, especially for someone who has never installed Linux without an installer.

It's important to realize that the handbook is designed to describe the steps required to install on a very wide variety of hardware, with different install needs. Because of this, many options presented in the handbook are unnecessary for a particular install and can be skipped.

#### Suggested steps

Prefixed with " **Suggested:**", some steps are not strictly required, but help in most cases, such as installing [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware).

#### Optional steps

Prefixed with " **Optional:**", many sections in the handbook are purely optional, and can be skipped if the user is looking for a simple, mostly vanilla install.

Examples of this include compiler flag customization, using a totally custom kernel, and disabling root login.

**Tip**

When following optional steps, it's important to be careful that all prerequisites were satisfied. Some optional steps depend on other optional steps.

#### Deprecated steps

Gentoo has existed for a long time. Some install processes are described in the handbook because they used to be more relevant, but are now largely deprecated. Instead of immediately removing this information, as it may still be helpful for some users, it may be designated as **Deprecated:** before removal. Once removed, the _history_ must be used to view this content.

#### Defaults and alternatives

Whenever choices are presented, the handbook will try to explain the pros and cons of each choice.

If potential choices are mutually exclusive, " **Default:**" is used to designate the most supported or commonly chosen option, while alternatives are listed with " **Alternative**".

**Önemli**

**Alternative** options are not inferior to **Default** s, but **Default** options are typically more widely used and may have better support.

### Installation options for Gentoo

Gentoo can be installed in many different ways. It can be downloaded and installed from official Gentoo installation media such as our bootable ISO images. The installation media can be installed on a USB stick or accessed via a netbooted environment. Alternatively, Gentoo can be installed from non-official media such as an already installed distribution or a non-Gentoo bootable disk (such as [Linux Mint](https://linuxmint.com/)).

This document covers the installation using official Gentoo Installation media or, in certain cases, netbooting.

**Not**

For help on the other installation approaches, including using non-Gentoo bootable media, please read our [Alternative installation guide](/wiki/Installation_alternatives "Installation alternatives").

We also provide a [Gentoo installation tips and tricks](/wiki/Gentoo_installation_tips_and_tricks "Gentoo installation tips and tricks") document that might be useful.

### Troubles

If a problem is found in the installation (or in the installation documentation), please visit our [bug tracking system](https://bugs.gentoo.org) and check if the bug is known. If not, please create a bug report for it so we can take care of it. Do not be afraid of the developers who are assigned to the bugs - they (generally) don't eat people.

Although this document is architecture-specific, it may contain references to other architectures as well, because large parts of the Gentoo Handbook use text that is identical for all architectures (to avoid duplication of effort). Such references have been kept to a minimum, to avoid confusion.

If there is some uncertainty about whether or not the problem is a user-problem (some error made despite having read the documentation carefully) or a software-problem (some error we made despite having tested the installation/documentation carefully) everybody is welcome to join the [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)) channel on irc.libera.chat. Of course, everyone is welcome otherwise too as our chat channel covers the broad Gentoo spectrum.

Speaking of which, if there are any additional questions regarding Gentoo, check out the [Frequently Asked Questions](/wiki/FAQ/tr "FAQ/tr") article. There are also [FAQs](https://forums.gentoo.org/viewforum.php?f=40) on the [Gentoo Forums](https://forums.gentoo.org).

[←](/wiki/Handbook:Alpha "Previous part") [Home](/wiki/Handbook:Alpha "Handbook:Alpha") [Choosing the media →](/wiki/Handbook:Alpha/Installation/Media "Next part")