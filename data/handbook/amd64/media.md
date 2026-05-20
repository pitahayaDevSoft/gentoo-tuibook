# Media

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Media/de "Handbook:AMD64/Installation/Media/de (100% translated)")
- English
- [Nederlands](/wiki/Handbook:AMD64/Installation/Media/nl "Handbook:AMD64/Installation/Media/nl (100% translated)")
- [Türkçe](/wiki/Handbook:AMD64/Installation/Media/tr "Handbook:AMD64/Installation/Media/tr (0% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Media/es "Handbook:AMD64/Installation/Media/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Media/fr "Handbook:AMD64/Installation/Media/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Media/it "Handbook:AMD64/Installation/Media/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Media/hu "Handbook:AMD64/Installation/Media/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Media/pl "Handbook:AMD64/Installation/Media/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Media/pt-br "Handbook:AMD64/Installation/Media/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Media/cs "Handbook:AMD64/Installation/Media/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Media/ru "Handbook:AMD64/Installation/Media/ru (100% translated)")
- [العربية](/wiki/Handbook:AMD64/Installation/Media/ar "Handbook:AMD64/Installation/Media/ar (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Media/ta "Handbook:AMD64/Installation/Media/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Media/zh-cn "Handbook:AMD64/Installation/Media/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Media/ja "Handbook:AMD64/Installation/Media/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Media/ko "Handbook:AMD64/Installation/Media/ko (100% translated)")

[AMD64 Handbook](/wiki/Handbook:AMD64 "Handbook:AMD64")[Installation](/wiki/Handbook:AMD64/Full/Installation "Handbook:AMD64/Full/Installation")[About the installation](/wiki/Handbook:AMD64/Installation/About "Handbook:AMD64/Installation/About")Choosing the media[Configuring the network](/wiki/Handbook:AMD64/Installation/Networking "Handbook:AMD64/Installation/Networking")[Preparing the disks](/wiki/Handbook:AMD64/Installation/Disks "Handbook:AMD64/Installation/Disks")[The stage file](/wiki/Handbook:AMD64/Installation/Stage "Handbook:AMD64/Installation/Stage")[Installing base system](/wiki/Handbook:AMD64/Installation/Base "Handbook:AMD64/Installation/Base")[Configuring the kernel](/wiki/Handbook:AMD64/Installation/Kernel "Handbook:AMD64/Installation/Kernel")[Configuring the system](/wiki/Handbook:AMD64/Installation/System "Handbook:AMD64/Installation/System")[Installing tools](/wiki/Handbook:AMD64/Installation/Tools "Handbook:AMD64/Installation/Tools")[Configuring the bootloader](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader")[Finalizing](/wiki/Handbook:AMD64/Installation/Finalizing "Handbook:AMD64/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:AMD64/Full/Working "Handbook:AMD64/Full/Working")[Portage introduction](/wiki/Handbook:AMD64/Working/Portage "Handbook:AMD64/Working/Portage")[USE flags](/wiki/Handbook:AMD64/Working/USE "Handbook:AMD64/Working/USE")[Portage features](/wiki/Handbook:AMD64/Working/Features "Handbook:AMD64/Working/Features")[Initscript system](/wiki/Handbook:AMD64/Working/Initscripts "Handbook:AMD64/Working/Initscripts")[Environment variables](/wiki/Handbook:AMD64/Working/EnvVar "Handbook:AMD64/Working/EnvVar")[Working with Portage](/wiki/Handbook:AMD64/Full/Portage "Handbook:AMD64/Full/Portage")[Files and directories](/wiki/Handbook:AMD64/Portage/Files "Handbook:AMD64/Portage/Files")[Variables](/wiki/Handbook:AMD64/Portage/Variables "Handbook:AMD64/Portage/Variables")[Mixing software branches](/wiki/Handbook:AMD64/Portage/Branches "Handbook:AMD64/Portage/Branches")[Additional tools](/wiki/Handbook:AMD64/Portage/Tools "Handbook:AMD64/Portage/Tools")[Custom package repository](/wiki/Handbook:AMD64/Portage/CustomTree "Handbook:AMD64/Portage/CustomTree")[Advanced features](/wiki/Handbook:AMD64/Portage/Advanced "Handbook:AMD64/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:AMD64/Full/Networking "Handbook:AMD64/Full/Networking")[Getting started](/wiki/Handbook:AMD64/Networking/Introduction "Handbook:AMD64/Networking/Introduction")[Advanced configuration](/wiki/Handbook:AMD64/Networking/Advanced "Handbook:AMD64/Networking/Advanced")[Modular networking](/wiki/Handbook:AMD64/Networking/Modular "Handbook:AMD64/Networking/Modular")[Wireless](/wiki/Handbook:AMD64/Networking/Wireless "Handbook:AMD64/Networking/Wireless")[Adding functionality](/wiki/Handbook:AMD64/Networking/Extending "Handbook:AMD64/Networking/Extending")[Dynamic management](/wiki/Handbook:AMD64/Networking/Dynamic "Handbook:AMD64/Networking/Dynamic")

## Contents

- [1Hardware requirements](#Hardware_requirements)
- [2Gentoo Linux installation media](#Gentoo_Linux_installation_media)
  - [2.1Minimal installation CD](#Minimal_installation_CD)
  - [2.2The Gentoo LiveGUI](#The_Gentoo_LiveGUI)
  - [2.3What are stage files?](#What_are_stage_files.3F)
- [3Downloading](#Downloading)
  - [3.1Obtain the media](#Obtain_the_media)
    - [3.1.1Navigating Gentoo mirrors](#Navigating_Gentoo_mirrors)
  - [3.2Verifying the downloaded files](#Verifying_the_downloaded_files)
    - [3.2.1Microsoft Windows-based verification](#Microsoft_Windows-based_verification)
    - [3.2.2Linux based verification](#Linux_based_verification)
- [4Writing the boot media](#Writing_the_boot_media)
  - [4.1Writing a bootable USB](#Writing_a_bootable_USB)
    - [4.1.1Writing with Linux](#Writing_with_Linux)
      - [4.1.1.1Determining the USB device path](#Determining_the_USB_device_path)
      - [4.1.1.2Writing with dd](#Writing_with_dd)
    - [4.1.2Writing with Windows](#Writing_with_Windows)
  - [4.2Burning a disk](#Burning_a_disk)
    - [4.2.1Burning with Microsoft Windows 7 and above](#Burning_with_Microsoft_Windows_7_and_above)
    - [4.2.2Burning with Linux](#Burning_with_Linux)
- [5Booting](#Booting)
  - [5.1Booting the installation media](#Booting_the_installation_media)
    - [5.1.1Kernel choices](#Kernel_choices)
    - [5.1.2Hardware options](#Hardware_options)
    - [5.1.3Logical volume/device management](#Logical_volume.2Fdevice_management)
    - [5.1.4Other options](#Other_options)
  - [5.2LiveGUI root access](#LiveGUI_root_access)
  - [5.3Extra hardware configuration](#Extra_hardware_configuration)
  - [5.4Optional: User accounts](#Optional:_User_accounts)
  - [5.5Optional: Viewing documentation while installing](#Optional:_Viewing_documentation_while_installing)
    - [5.5.1TTYs](#TTYs)
    - [5.5.2GNU Screen](#GNU_Screen)
  - [5.6Optional: Starting the SSH daemon](#Optional:_Starting_the_SSH_daemon)

## Hardware requirements\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-1 "Edit section: ")\]

Gentoo can be installed on all manners of systems, so listing minimal hardware specification is more down to how long a user is willing to wait. Any reasonably modern machine can complete the install process in less than an hour, however it is not a race so, take your time, and remember it takes as long as it takes.

For a smooth install process though, the Handbook recommends at least 40GB of space for the root partition.

## Gentoo Linux installation media\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-2 "Edit section: ")\]

**Tip**

While it's recommended to use the official Gentoo boot media when installing, it's possible to use other installation environments. However, there is no guarantee they will contain required components. If an alternate install environment is used, skip to [Preparing the disks](/wiki/Handbook:AMD64/Installation/Disks "Handbook:AMD64/Installation/Disks").

### Minimal installation CD\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-3 "Edit section: ")\]

The _Gentoo minimal installation CD_, also known as the installcd, is a small, bootable image: a self-contained Gentoo environment. This image is maintained by [Gentoo developers](/wiki/Project:RelEng "Project:RelEng") and designed to allow any user with an Internet connection to install Gentoo. During the boot process, the hardware is detected, and appropriate drivers are automatically loaded.

Minimal Installation CD releases are named using the format:
install-<arch>-minimal-<release timestamp>.iso.

**Important**

The Gentoo minimal installation CD requires at least 140MB of RAM to boot.

### The Gentoo LiveGUI

The LiveGUI provides a more convenient environment for the manual installation process, using the KDE desktop environment. This allows a user to open a graphical based web browser, such as Mozilla Firefox, to make reading the handbook whilst installing in terminal more pleasant.

The LiveGUI also includes more wireless driver support than the minimal CD provides and can be more easily connected using the NetworkManager applet in KDE

**Important**

The Gentoo LiveGUI requires at least 2GB of RAM to boot.

### What are stage files?\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-5 "Edit section: ")\]

A [stage file](/wiki/Stage_file "Stage file") is an archive which serves as the seed for a Gentoo environment.

Stage 3 files can be downloaded from releases/amd64/autobuilds/ on any of the [official Gentoo mirrors](https://www.gentoo.org/downloads/mirrors/). Stages are updated frequently and are therefore not included within official live images.

**Tip**

For now, _stage files_ can be ignored. They will be described in greater detail later when they are needed

**Note**

Historically, Gentoo provided stages lower than 3. Over time theses type of installations have become obsolete due to better methods and no longer provided for this reason.

## Downloading\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-6 "Edit section: ")\]

### Obtain the media\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-7 "Edit section: ")\]

The default installation media used by Gentoo Linux are the _minimal installation CDs_, which provide a very small, bootable, Gentoo Linux environment. This environment contains the necessary tools to install Gentoo. The images themselves can be downloaded from the [downloads page](https://www.gentoo.org/downloads/) (recommended) or by manually browsing to the ISO location on one of the many [available mirrors](https://www.gentoo.org/downloads/mirrors/).

#### Navigating Gentoo mirrors\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-8 "Edit section: ")\]

If downloading from a mirror, the minimal installation CDs can be found by:

1. Connect to the mirror, typically using a local one found at [Gentoo source mirrors](https://www.gentoo.org/downloads/mirrors/).
2. Navigate to the releases/ directory.
3. Select the directory for the relevant target architecture (such as **amd64/**).
4. Select the autobuilds/ directory.
5. For **amd64** and **x86** architectures select either the current-install-amd64-minimal/ or current-install-x86-minimal/ directory (respectively). For all other architectures navigate to the current-iso/ directory.

**Note**

Some target architectures such as **arm**, **mips**, and **s390** will not have minimal install CDs. At this time the [Gentoo Release Engineering project](/wiki/Project:RelEng "Project:RelEng") does not support building .iso files for these targets.

Inside this location, the installation media file is the file with the .iso suffix. For instance, take a look at the following listing:

CODE **Example list of downloadable files at releases/amd64/autobuilds/current-install-amd64-minimal/**

```
[TXT]	install-amd64-minimal-20231112T170154Z.iso.asc	        2023-11-12 20:41        488
[TXT]	install-amd64-minimal-20231119T164701Z.iso.asc	        2023-11-19 18:41        488
[TXT]	install-amd64-minimal-20231126T163200Z.iso.asc	        2023-11-26 18:41        488
[TXT]	install-amd64-minimal-20231203T170204Z.iso.asc	        2023-12-03 18:41        488
[TXT]	install-amd64-minimal-20231210T170356Z.iso.asc	        2023-12-10 19:01        488
[TXT]	install-amd64-minimal-20231217T170203Z.iso.asc	        2023-12-17 20:01        488
[TXT]	install-amd64-minimal-20231224T164659Z.iso.asc	        2023-12-24 20:41        488
[TXT]	install-amd64-minimal-20231231T163203Z.iso.asc	        2023-12-31 19:01        488
[ ]     install-amd64-minimal-20240107T170309Z.iso              2024-01-07 20:42        466M
[ ]     install-amd64-minimal-20240107T170309Z.iso.CONTENTS.gz	2024-01-07 20:42        9.8K
[ ]     install-amd64-minimal-20240107T170309Z.iso.DIGESTS      2024-01-07 21:01        1.3K
[TXT]   install-amd64-minimal-20240107T170309Z.iso.asc	        2024-01-07 21:01        488
[ ]     install-amd64-minimal-20240107T170309Z.iso.sha256       2024-01-07 21:01        660
[TXT]	latest-install-amd64-minimal.txt                        2024-01-08 02:01        653

```

In the above example, the install-amd64-minimal-20240107T170309Z.iso file is the minimal installation CD itself. But as can be seen, other related files exist as well:

- A .CONTENTS.gz file which is a gz-compressed text file listing all files available on the installation media. This file can be useful to check if particular firmware or drivers are available on the installation media before downloading it.
- A .asc file which is a cryptographic signature of the ISO file. This is used to verify image integrity and authenticity - that the download is indeed provided by the [Gentoo Release Engineering team](/wiki/Project:RelEng "Project:RelEng"), free from tampering.

Ignore the other files available at this location for now - those will come back when the installation has proceeded further. Download the .iso file and, if verification of the download is wanted, download the .iso.asc file for the .iso file as well.

### Verifying the downloaded files\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-9 "Edit section: ")\]

This step ensures that the downloaded file is not corrupt and has indeed been provided by the [Gentoo Infrastructure team](/wiki/Project:Infrastructure "Project:Infrastructure").

The .asc file provides a cryptographic signature of the ISO. Validating it ensures
the installation file is provided by the Gentoo Release Engineering team and is intact and unmodified.

#### Microsoft Windows-based verification\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-10 "Edit section: ")\]

To first verify the cryptographic signature, tools such as [GPG4Win](https://www.gpg4win.org/) can be used. After installation, the public keys of the Gentoo Release Engineering team need to be imported. The list of keys is available on the [signatures page](https://www.gentoo.org/downloads/signatures/). Once imported, the downloaded ISO can be verified using the .asc file.

#### Linux based verification\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-11 "Edit section: ")\]

On a Linux system, the most common method for verifying the cryptographic signature is to use the [app-crypt/gnupg](https://packages.gentoo.org/packages/app-crypt/gnupg) software. With this package installed, the following command is used to verify the cryptographic signature in the .asc file.

**Important**

When importing Gentoo keys, verify that the 16-character key ID ( `BB572E0E2D182910`) matches.

Gentoo keys can be downloaded from hkps://keys.gentoo.org using fingerprints available on the [signatures page](https://www.gentoo.org/downloads/signatures/):

`user $` `gpg --keyserver hkps://keys.gentoo.org --recv-keys 13EBBDBEDE7A12775DFDB1BABB572E0E2D182910`

```
gpg: directory '/root/.gnupg' created
gpg: keybox '/root/.gnupg/pubring.kbx' created
gpg: /root/.gnupg/trustdb.gpg: trustdb created
gpg: key BB572E0E2D182910: public key "Gentoo Linux Release Engineering (Automated Weekly Release Key) <releng@gentoo.org>" imported
gpg: Total number processed: 1
gpg:               imported: 1

```

Alternatively, [WKD](/wiki/WKD "WKD") can be used to download the keys:

`user $` `gpg --auto-key-locate=clear,nodefault,wkd --locate-key releng@gentoo.org`

```
gpg: key 9E6438C817072058: public key "Gentoo Linux Release Engineering (Gentoo Linux Release Signing Key) <releng@gentoo.org>" imported
gpg: key BB572E0E2D182910: public key "Gentoo Linux Release Engineering (Automated Weekly Release Key) <releng@gentoo.org>" imported
gpg: Total number processed: 2
gpg:               imported: 2
gpg: no ultimately trusted keys found
pub   dsa1024 2004-07-20 [SC] [expires: 2025-07-01]
      D99EAC7379A850BCE47DA5F29E6438C817072058
uid           [ unknown] Gentoo Linux Release Engineering (Gentoo Linux Release Signing Key) <releng@gentoo.org>
sub   elg2048 2004-07-20 [E] [expires: 2025-07-01]

```

Or if using official Gentoo release media, import the key from /usr/share/openpgp-keys/gentoo-release.asc (provided by [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release)):

`user $` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

```
gpg: directory '/home/larry/.gnupg' created
gpg: keybox '/home/larry/.gnupg/pubring.kbx' created
gpg: key DB6B8C1F96D8BF6D: 2 signatures not checked due to missing keys
gpg: /home/larry/.gnupg/trustdb.gpg: trustdb created
gpg: key DB6B8C1F96D8BF6D: public key "Gentoo ebuild repository signing key (Automated Signing Key) <infrastructure@gentoo.org>" imported
gpg: key 9E6438C817072058: 3 signatures not checked due to missing keys
gpg: key 9E6438C817072058: public key "Gentoo Linux Release Engineering (Gentoo Linux Release Signing Key) <releng@gentoo.org>" imported
gpg: key BB572E0E2D182910: 1 signature not checked due to a missing key
gpg: key BB572E0E2D182910: public key "Gentoo Linux Release Engineering (Automated Weekly Release Key) <releng@gentoo.org>" imported
gpg: key A13D0EF1914E7A72: 1 signature not checked due to a missing key
gpg: key A13D0EF1914E7A72: public key "Gentoo repository mirrors (automated git signing key) <repomirrorci@gentoo.org>" imported
gpg: Total number processed: 4
gpg:               imported: 4
gpg: no ultimately trusted keys found
```

Next verify the cryptographic signature:

`user $` `gpg --verify install-amd64-minimal-20240107T170309Z.iso.asc install-amd64-minimal-20240107T170309Z.iso`

```
gpg: Signature made Sun 07 Jan 2024 03:01:10 PM CST
gpg:                using RSA key 534E4209AB49EEE1C19D96162C44695DB9F6043D
gpg: Good signature from "Gentoo Linux Release Engineering (Automated Weekly Release Key) <releng@gentoo.org>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: 13EB BDBE DE7A 1277 5DFD  B1BA BB57 2E0E 2D18 2910
     Subkey fingerprint: 534E 4209 AB49 EEE1 C19D  9616 2C44 695D B9F6 043D

```

To be absolutely certain that everything is valid, verify the fingerprint shown with the fingerprint on the [Gentoo signatures page](https://www.gentoo.org/downloads/signatures/).

**Note**

It's generally good practice to mark an imported key as trusted, once it's certain the key is trustworthy. When trusted keys are verified, gpg will not say _unknown_ and warn about the signature being untrusted.

## Writing the boot media\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-12 "Edit section: ")\]

Of course, with just an ISO file downloaded, the Gentoo Linux installation cannot be started. The ISO file must be written to bootable media. This generally requires that the image is extracted to a filesystem, or written directly to a device.

### Writing a bootable USB\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-13 "Edit section: ")\]

**Warning**

Using tools such as Ventoy to write to a USB drive can cause boot failures.

Most modern systems support booting from a USB device.

#### Writing with Linux\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-14 "Edit section: ")\]

dd is typically available on most Linux distros, and can be used to write the Gentoo boot media to a USB drive.

##### Determining the USB device path\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-15 "Edit section: ")\]

Before writing, the path to the desired storage device must be determined.

dmesg will display detailed information describing the storage device as it is added to the system:

`root #` `dmesg`

```
[268385.319745] sd 19:0:0:0: [sdd] 60628992 512-byte logical blocks: (31.0 GB/28.9 GiB)
```

Alternatively, lsblk can be used to display available storage devices:

`root #` `lsblk`

```
sdd           8:48   1  28.9G  0 disk
├─sdd1        8:49   1   246K  0 part
├─sdd2        8:50   1   2.8M  0 part
├─sdd3        8:51   1 463.5M  0 part
└─sdd4        8:52   1   300K  0 part

```

Once the device name has been determined, this can be added to the path prefix _/dev/_ to get the device path /dev/sdd.

**Tip**

Using the base device path, ie. sdd opposed to sdd1, is recommend as the Gentoo _boot media_ contains a full [GPT](/wiki/GPT "GPT") partition scheme.

##### Writing with dd\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-16 "Edit section: ")\]

**Warning**

Be sure to check the target ( _of=target_) path before executing dd, as it will be overwritten.

With the device path (/dev/sdd) and _boot media_ install-amd64-minimal-<release timestamp>.iso ready:

`root #` `dd if=install-amd64-minimal-<release timestamp>.iso of=/dev/sdd bs=4096 status=progress && sync`

**Note**

**if=** specifies the _input file_, **of=** specifies the _output file_, which in this case, is a device.

**Tip**

**bs=4096** is used as it speeds up transfers in most cases, **status=progress** displays transfers stats.

#### Writing with Windows\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-17 "Edit section: ")\]

[win32diskimager](https://sourceforge.net/projects/win32diskimager/) is simple to use Windows utility to write iso images to a USB drive.

They also provide excellent documentation to assist if needed.

**Warning**

Only download from [https://sourceforge.net/projects/win32diskimager/](https://sourceforge.net/projects/win32diskimager/)

### Burning a disk\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-18 "Edit section: ")\]

**See also**

A more elaborate set of instructions can be found in [CD/DVD/BD\_writing#Image\_writing](/wiki/CD/DVD/BD_writing#Image_writing "CD/DVD/BD writing").

#### Burning with Microsoft Windows 7 and above\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-19 "Edit section: ")\]

Versions of Microsoft Windows 7 and above can both mount and burn ISO images to optical media without the requirement for third-party software. Simply insert a burnable disk, browse to the downloaded ISO files, right click the file in Windows Explorer, and select "Burn disk image".

#### Burning with Linux\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-20 "Edit section: ")\]

The cdrecord utility from the package [app-cdr/cdrtools](https://packages.gentoo.org/packages/app-cdr/cdrtools) can burn ISO images on Linux.

To burn the ISO file on the CD in the /dev/sr0 device (this is the first CD device on the system - substitute with the right device file if necessary):

`user $` `cdrecord dev=/dev/sr0 install-amd64-minimal-20141204.iso`

Users that prefer a graphical user interface can use K3B, part of the [kde-apps/k3b](https://packages.gentoo.org/packages/kde-apps/k3b) package. In K3B, go to Tools and use Burn CD Image.

## Booting\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-21 "Edit section: ")\]

### Booting the installation media\[ [edit](/index.php?title=Handbook:AMD64/Blocks/Booting&action=edit&section=T-1 "Edit section: ")\]

Once the installation media is ready, it is time to boot it. Insert the media in the system, reboot, and enter the motherboard's firmware user interface. This is usually performed by pressing a keyboard key such as `DEL`, `F1`, `F10`, or `ESC` during the Power-On Self-Test (POST) process. The 'trigger' key varies depending on the system and motherboard. If it is not obvious use an internet search engine and do some research using the motherboard's model name as the search keyword. Results should be easy to determine. Once inside the motherboard's firmware menu, change the boot order so that the external bootable media (CD/DVD disks or USB drives) are tried _before_ the internal disk devices. Without this change, the system will most likely reboot to the internal disk device, ignoring the newly attached bootable media.

**Important**

When installing Gentoo on a system with an UEFI firmware interface, ensure the live image has been booted in UEFI mode. In the accidental event that DOS/legacy BIOS boot was initiated, then it will be necessary to reboot in UEFI mode before finalizing the Gentoo Linux installation.

Ensure that the installation media is inserted or plugged into the system, and reboot. A GRUB boot prompt should be shown with various boot entries. At this screen, `Enter` will begin the boot process with the default boot options. To boot the installation media with customized boot options, such as passing additional kernel parameters or the [following hardware options](#Hardware_options), highlight a boot entry, then press the `e` key to edit the boot entry. Make the necessary modification(s), then press `ctrl` + `x` or `F10` to boot the modified entry.

**Note**

In all likelihood, the default gentoo kernel, as mentioned above, without specifying any of the optional parameters will work just fine. For boot troubleshooting and expert options, continue on with this section. Otherwise, just press `Enter` and skip ahead to [Extra hardware configuration](/wiki/Handbook:AMD64/Installation/Media#Extra_hardware_configuration "Handbook:AMD64/Installation/Media").

At the boot prompt, users get the option of displaying the available kernels ( `F1`) and boot options ( `F2`). If no choice is made within 15 seconds (either displaying information or using a kernel) then the installation media will fall back to booting from disk. This allows installations to reboot and try out their installed environment without the need to remove the CD from the tray (something well appreciated for remote installations).

Specifying a kernel was mentioned. On the Minimal installation media, only two predefined kernel boot entries are provided. The default option is called gentoo. The other being the _-nofb_ variant; this disables kernel framebuffer support.

The next section displays a short overview of the available kernels and their descriptions:

#### Kernel choices\[ [edit](/index.php?title=Handbook:AMD64/Blocks/Booting&action=edit&section=T-2 "Edit section: ")\]

gentooDefault kernel with support for K8 CPUs (including NUMA support) and EM64T CPUs.gentoo-nofbSame as _gentoo_ but without framebuffer support.memtest86Test the system RAM for errors.

Alongside the kernel, boot options help in tuning the boot process further.

#### Hardware options\[ [edit](/index.php?title=Handbook:AMD64/Blocks/Booting&action=edit&section=T-3 "Edit section: ")\]

acpi=onThis loads support for ACPI and also causes the acpid daemon to be started by the CD on boot. This is only needed if the system requires ACPI to function properly. This is not required for Hyperthreading support.acpi=offCompletely disables ACPI. This is useful on some older systems and is also a requirement for using APM. This will disable any Hyperthreading support of your processor.console=XThis sets up serial console access for the CD. The first option is the device, usually ttyS0, followed by any connection options, which are comma separated. The default options are 9600,8,n,1.dmraid=XThis allows for passing options to the device-mapper RAID subsystem. Options should be encapsulated in quotes.doapmThis loads APM driver support. This also requires that `acpi=off`.dopcmciaThis loads support for PCMCIA and Cardbus hardware and also causes the pcmcia cardmgr to be started by the CD on boot. This is only required when booting from PCMCIA/Cardbus devices.doscsiThis loads support for most SCSI controllers. This is also a requirement for booting most USB devices, as they use the SCSI subsystem of the kernel.sda=strokeThis allows the user to partition the whole hard disk even when the BIOS is unable to handle large disks. This option is only used on machines with an older BIOS. Replace sda with the device that requires this option.ide=nodmaThis forces the disabling of DMA in the kernel and is required by some IDE chipsets and also by some CDROM drives. If the system is having trouble reading from the IDE CDROM, try this option. This also disables the default hdparm settings from being executed.noapicThis disables the Advanced Programmable Interrupt Controller that is present on newer motherboards. It has been known to cause some problems on older hardware.nodetectThis disables all of the autodetection done by the CD, including device autodetection and DHCP probing. This is useful for debugging a failing CD or driver.nodhcpThis disables DHCP probing on detected network cards. This is useful on networks with only static addresses.nodmraidDisables support for device-mapper RAID, such as that used for on-board IDE/SATA RAID controllers.nofirewireThis disables the loading of Firewire modules. This should only be necessary if your Firewire hardware is causing a problem with booting the CD.nogpmThis disables gpm console mouse support.nohotplugThis disables the loading of the hotplug and coldplug init scripts at boot. This is useful for debugging a failing CD or driver.nokeymapThis disables the keymap selection used to select non-US keyboard layouts.nolapicThis disables the local APIC on Uniprocessor kernels.nosataThis disables the loading of Serial ATA modules. This is used if the system is having problems with the SATA subsystem.nosmpThis disables SMP, or Symmetric Multiprocessing, on SMP-enabled kernels. This is useful for debugging SMP-related issues with certain drivers and motherboards.nosoundThis disables sound support and volume setting. This is useful for systems where sound support causes problems.nousbThis disables the autoloading of USB modules. This is useful for debugging USB issues.slowusbThis adds some extra pauses into the boot process for slow USB CDROMs, like in the IBM BladeCenter.

#### Logical volume/device management\[ [edit](/index.php?title=Handbook:AMD64/Blocks/Booting&action=edit&section=T-4 "Edit section: ")\]

dolvmThis enables support for Linux's Logical Volume Management.

#### Other options\[ [edit](/index.php?title=Handbook:AMD64/Blocks/Booting&action=edit&section=T-5 "Edit section: ")\]

debugEnables debugging code. This might get messy, as it displays a lot of data to the screen.docacheThis caches the entire runtime portion of the CD into RAM, which allows the user to umount /mnt/cdrom and mount another CDROM. This option requires that there is at least twice as much available RAM as the size of the CD.doload=XThis causes the initial ramdisk to load any module listed, as well as dependencies. Replace X with the module name. Multiple modules can be specified by a comma-separated list.dosshdStarts sshd on boot, which is useful for unattended installs.passwd=fooSets whatever follows the equals as the root password, which is required for _dosshd_ since the root password is by default scrambled.noload=XThis causes the initial ramdisk to skip the loading of a specific module that may be causing a problem. Syntax matches that of doload.nonfsDisables the starting of portmap/nfsmount on boot.noxThis causes an X-enabled LiveCD to not automatically start X, but rather, to drop to the command line instead.scandelayThis causes the CD to pause for 10 seconds during certain portions the boot process to allow for devices that are slow to initialize to be ready for use.scandelay=XThis allows the user to specify a given delay, in seconds, to be added to certain portions of the boot process to allow for devices that are slow to initialize to be ready for use. Replace X with the number of seconds to pause.

**Note**

The bootable media will check for `no*` options before `do*` options, so that options can be overridden in the exact order specified.

Now boot the media, select a kernel (if the default gentoo kernel does not suffice) and boot options. As an example, we boot the gentoo kernel, with `dopcmcia` as a kernel parameter:

`boot:` `gentoo dopcmcia`

Next the user will be greeted with a boot screen and progress bar. If the installation is done on a system with a non-US keyboard, make sure to immediately press `Alt` + `F1` to switch to verbose mode and follow the prompt. If no selection is made in 10 seconds the default (US keyboard) will be accepted and the boot process will continue. Once the boot process completes, the user is automatically logged in to the "Live" Gentoo Linux environment as the _root_ user, the super user. A root prompt is displayed on the current console, and one can switch to other consoles by pressing `Alt` + `F2`, `Alt` + `F3` and `Alt` + `F4`. Get back to the one started on by pressing `Alt` + `F1`.

### LiveGUI root access\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-22 "Edit section: ")\]

sudo has been configured to run without the need of a password on the LiveGUI as both the root and gentoo have a scrambled password.

To gain access to the superuser account, in any terminal run:

`root #` `sudo -i`

### Extra hardware configuration\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-23 "Edit section: ")\]

When the Installation medium boots, it tries to detect all the hardware devices and loads the appropriate kernel modules to support the hardware. In the vast majority of cases, it does a very good job. However, in some cases it may not auto-load the kernel modules needed by the system. If the PCI auto-detection missed some of the system's hardware, the appropriate kernel modules have to be loaded manually.

In the next example the 8139too module (which supports certain kinds of network interfaces) is loaded:

`root #` `modprobe 8139too`

### Optional: User accounts\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-24 "Edit section: ")\]

If other people need access to the installation environment, or there is need to run commands as a non-root user on the installation medium (such as to chat using irssi without root privileges for security reasons), then an additional user account needs to be created and the root password set to a strong password.

To change the root password, use the passwd utility:

`root #` `passwd`

```
New password: (Enter the new password)
Re-enter password: (Re-enter the password)

```

To create a user account, first enter their credentials, followed by the account's password. The useradd and passwd commands are used for these tasks.

In the next example, a user called _larry_ is created and is added to the `users` and `wheel` groups:

`root #` `useradd -m -G users,wheel larry
`

`root #` `passwd larry`

```
New password: (Enter larry's password)
Re-enter password: (Re-enter larry's password)

```

To switch from the (current) _root_ user to the newly created user account, use the su command:

`root #` `su - larry`

### Optional: Viewing documentation while installing\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-25 "Edit section: ")\]

#### TTYs\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-26 "Edit section: ")\]

To view the Gentoo handbook from a TTY during the installation, first create a user account as described above, then press `Alt` + `F2` to go to a new terminal (TTY) and login as the newly created user. Following the [principle of least privilege](https://en.wikipedia.org/wiki/Principle_of_least_privilege "wikipedia:Principle of least privilege"), it is best practice to avoid browsing the web or generally performing any task with higher privileges than necessary. The root account has full control of the system and therefore must be used sparingly.

During the installation, the links web browser can be used to browse the Gentoo handbook - of course only from the moment that the Internet connection is working.

`user $` `links https://wiki.gentoo.org/wiki/Handbook:AMD64`

To go back to the original terminal, press `Alt` + `F1`.

**Tip**

When booted to the Gentoo minimal or Gentoo admin environments, seven TTYs will be available. They can be switched by pressing `Alt` then a function key between `F1`- `F7`. It can be useful to switch to a new terminal when waiting for job to complete, to open documentation, etc.

#### GNU Screen\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-27 "Edit section: ")\]

The [Screen](/wiki/Screen "Screen") utility is installed by default on official Gentoo installation media. It may be more efficient for the seasoned Linux enthusiast to use screen to view installation instructions via split panes rather than the multiple TTY method mentioned above.

### Optional: Starting the SSH daemon\[ [edit](/index.php?title=Handbook:Parts/Installation/Media&action=edit&section=T-28 "Edit section: ")\]

To allow other users to access the system during the installation (perhaps to provide/receive support during an installation, or even do it remotely), a user account needs to be created (as was documented earlier on) and the SSH daemon needs to be started.

To fire up the SSH daemon on an OpenRC init, execute the following command:

`root #` `rc-service sshd start`

**Note**

If users log on to the system, they will see a message that the host key for this system needs to be confirmed (through what is called a fingerprint). This behavior is typical and can be expected for initial connections to an SSH server. However, later when the system is set up and someone logs on to the newly created system, the SSH client will warn that the host key has been changed. This is because the user now logs on to - for SSH - a different server (namely the freshly installed Gentoo system rather than the live environment that the installation is currently using). Follow the instructions given on the screen then to replace the host key on the client system.

To be able to use sshd, the network needs to function properly. Continue with the chapter on [Configuring the network](/wiki/Handbook:AMD64/Installation/Networking "Handbook:AMD64/Installation/Networking").

[← About the installation](/wiki/Handbook:AMD64/Installation/About "Previous part") [Home](/wiki/Handbook:AMD64 "Handbook:AMD64") [Configuring the network →](/wiki/Handbook:AMD64/Installation/Networking "Next part")