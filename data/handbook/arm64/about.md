# About

## Contents

- [1Introduction](#Introduction)
  - [1.1Welcome](#Welcome)
  - [1.2How is the installation structured](#How_is_the_installation_structured)
  - [1.3What are the options](#What_are_the_options)
  - [1.4Troubles](#Troubles)

## Introduction\[ [edit](/index.php?title=Handbook:ARM64/Installation/About&action=edit&section=1 "Edit section: Introduction")\]

### Welcome\[ [edit](/index.php?title=Handbook:ARM64/Installation/About&action=edit&section=2 "Edit section: Welcome")\]

First of all, welcome to Gentoo. You are about to enter the world of choices and performance. Gentoo is all about choices. When installing Gentoo, this is made clear several times -- users can choose how much they want to compile themselves, how to install Gentoo, what system logger to use, etc.

Gentoo is a fast, modern meta-distribution with a clean and flexible design. Gentoo is built around free software and doesn't hide from its users what is beneath the hood. Portage, the package maintenance system which Gentoo uses, is written in Python, meaning the user can easily view and modify the source code. Gentoo's packaging system uses source code (although support for precompiled packages is included too) and configuring Gentoo happens through regular text files. In other words, openness everywhere.

It is very important that everybody understands that choices are what makes Gentoo run. We try not to force users onto anything they don't like. If anyone believes otherwise, please [bugreport](https://bugs.gentoo.org) it.

### How is the installation structured\[ [edit](/index.php?title=Handbook:ARM64/Installation/About&action=edit&section=3 "Edit section: How is the installation structured")\]

The Gentoo Installation can be seen as a 10-step procedure, corresponding to the next set of chapters. Every step results in a certain state:

01. After step 1, the user is in a working environment ready to install Gentoo
02. After step 2, the Internet connection is ready to install Gentoo
03. After step 3, the hard disks are initialized to host the Gentoo installation
04. After step 4, the installation environment is prepared and the user is ready to chroot into the new environment
05. After step 5, core packages, which are the same on all Gentoo installations, are installed
06. After step 6, the Linux kernel is installed
07. After step 7, the user will have configured most of the Gentoo system configuration files
08. After step 8, the necessary system tools are installed
09. After step 9, the proper boot loader has been installed and configured
10. After step 10, the freshly installed Gentoo Linux environment is ready to be explored

Whenever a certain choice is presented, the handbook will try to explain what the pros and cons are. Although the text then continues with a default choice (identified by "Default: " in the title), the other possibilities will be documented as well (marked by "Alternative: " in the title). Do not think that the default is what Gentoo recommends. It is however what Gentoo believes most users will use.

Sometimes an optional step can be followed. Such steps are marked as "Optional: " and are therefore not needed to install Gentoo. However, some optional steps are dependent on a previous decision made. The instructions will inform the reader when this happens, both when the decision is made, and right before the optional step is described.

### What are the options\[ [edit](/index.php?title=Handbook:ARM64/Installation/About&action=edit&section=4 "Edit section: What are the options")\]

Gentoo can be installed in many different ways. For ARM64 hardware, generally this involves downloading a system image for DD'ing onto eMMC, SD, or uSD media.

This document covers the installation specifically for ARM64 hardware. We try to cover the specifics for known hardware that we've been able to test and validate. ARM64 hardware not listed will probably work however you'll need to understand what is needed for your hardware and adjust your installation accordingly.

**Note**

For help on the other installation approaches, including using non-Gentoo CDs, please read our [Alternative Installation Guide](/wiki/Installation_alternatives "Installation alternatives").

We also provide a [Gentoo Installation Tips & Tricks](/wiki/Gentoo_installation_tips_and_tricks "Gentoo installation tips and tricks") document that might be useful to read as well.

### Troubles\[ [edit](/index.php?title=Handbook:ARM64/Installation/About&action=edit&section=5 "Edit section: Troubles")\]

If a problem is found in the installation (or in the installation documentation), please visit our [bugtracking system](https://bugs.gentoo.org) and check if the bug is known. If not, please create a bugreport for it so we can take care of it. Do not be afraid of the developers who are assigned to the bugs -- they (generally) don't eat people.

Note though that, although this document is architecture-specific, it might contain references to other architectures as well. This is due to the fact that large parts of the Gentoo Handbook use source code that is common for all architectures (to avoid duplication of efforts and starvation of development resources). We will try to keep this to a minimum to avoid confusion.

If there is some uncertainty whether or not the problem is a user-problem (some error made despite having read the documentation carefully) or a software-problem (some error we made despite having tested the installation/documentation carefully) everybody is welcome to join the [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)) IRC channel. Of course, everyone is welcome otherwise too as our chat channel covers the broad Gentoo spectrum.

Speaking of which, if there are any additional questions regarding Gentoo, check out our [Frequently Asked Questions](/wiki/FAQ "FAQ"), available from the [Gentoo Wiki](/wiki/Main_Page "Main Page"). There are also [FAQs](https://forums.gentoo.org/viewforum.php?f=40) on the [Gentoo Forums](https://forums.gentoo.org).