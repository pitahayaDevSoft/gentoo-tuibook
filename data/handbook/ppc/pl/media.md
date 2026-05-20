# Media

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Media/de "Handbuch:PPC/Installation/Medium (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Media "Handbook:PPC/Installation/Media (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Media/es "Manual de Gentoo: PPC/Instalación/Medio (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Media/fr "Handbook:PPC/Installation/Media/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Media/it "Handbook:PPC/Installation/Media (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Media/hu "Handbook:PPC/Installation/Media/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:PPC/Installation/Media/pt-br "Manual:PPC/Instalação/Midia (100% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Media/ru "Handbook:PPC/Installation/Media (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Media/ta "கையேடு:PPC/நிறுவல்/ஊடகம் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Media/zh-cn "手册：PPC/安装/选择安装媒介 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Media/ja "ハンドブック:PPC/インストール/メディア (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Media/ko "Handbook:PPC/Installation/Media/ko (100% translated)")

[PPC Handbook](/wiki/Handbook:PPC "Handbook:PPC")[Installation](/wiki/Handbook:PPC/Full/Installation "Handbook:PPC/Full/Installation")[About the installation](/wiki/Handbook:PPC/Installation/About/pl "Handbook:PPC/Installation/About/pl")Choosing the media[Configuring the network](/wiki/Handbook:PPC/Installation/Networking/pl "Handbook:PPC/Installation/Networking/pl")[Preparing the disks](/wiki/Handbook:PPC/Installation/Disks/pl "Handbook:PPC/Installation/Disks/pl")[The stage file](/wiki/Handbook:PPC/Installation/Stage/pl "Handbook:PPC/Installation/Stage/pl")[Installing base system](/wiki/Handbook:PPC/Installation/Base/pl "Handbook:PPC/Installation/Base/pl")[Configuring the kernel](/wiki/Handbook:PPC/Installation/Kernel/pl "Handbook:PPC/Installation/Kernel/pl")[Configuring the system](/wiki/Handbook:PPC/Installation/System/pl "Handbook:PPC/Installation/System/pl")[Installing tools](/wiki/Handbook:PPC/Installation/Tools/pl "Handbook:PPC/Installation/Tools/pl")[Configuring the bootloader](/wiki/Handbook:PPC/Installation/Bootloader/pl "Handbook:PPC/Installation/Bootloader/pl")[Finalizing](/wiki/Handbook:PPC/Installation/Finalizing/pl "Handbook:PPC/Installation/Finalizing/pl")[Working with Gentoo](/wiki/Handbook:PPC/Full/Working "Handbook:PPC/Full/Working")[Portage introduction](/wiki/Handbook:PPC/Working/Portage/pl "Handbook:PPC/Working/Portage/pl")[USE flags](/wiki/Handbook:PPC/Working/USE/pl "Handbook:PPC/Working/USE/pl")[Portage features](/wiki/Handbook:PPC/Working/Features/pl "Handbook:PPC/Working/Features/pl")[Initscript system](/wiki/Handbook:PPC/Working/Initscripts/pl "Handbook:PPC/Working/Initscripts/pl")[Environment variables](/wiki/Handbook:PPC/Working/EnvVar/pl "Handbook:PPC/Working/EnvVar/pl")[Working with Portage](/wiki/Handbook:PPC/Full/Portage "Handbook:PPC/Full/Portage")[Files and directories](/wiki/Handbook:PPC/Portage/Files/pl "Handbook:PPC/Portage/Files/pl")[Variables](/wiki/Handbook:PPC/Portage/Variables/pl "Handbook:PPC/Portage/Variables/pl")[Mixing software branches](/wiki/Handbook:PPC/Portage/Branches/pl "Handbook:PPC/Portage/Branches/pl")[Additional tools](/wiki/Handbook:PPC/Portage/Tools/pl "Handbook:PPC/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:PPC/Portage/CustomTree/pl "Handbook:PPC/Portage/CustomTree/pl")[Advanced features](/wiki/Handbook:PPC/Portage/Advanced/pl "Handbook:PPC/Portage/Advanced/pl")[OpenRC network configuration](/wiki/Handbook:PPC/Full/Networking "Handbook:PPC/Full/Networking")[Getting started](/wiki/Handbook:PPC/Networking/Introduction/pl "Handbook:PPC/Networking/Introduction/pl")[Advanced configuration](/wiki/Handbook:PPC/Networking/Advanced/pl "Handbook:PPC/Networking/Advanced/pl")[Modular networking](/wiki/Handbook:PPC/Networking/Modular/pl "Handbook:PPC/Networking/Modular/pl")[Wireless](/wiki/Handbook:PPC/Networking/Wireless/pl "Handbook:PPC/Networking/Wireless/pl")[Adding functionality](/wiki/Handbook:PPC/Networking/Extending/pl "Handbook:PPC/Networking/Extending/pl")[Dynamic management](/wiki/Handbook:PPC/Networking/Dynamic/pl "Handbook:PPC/Networking/Dynamic/pl")

## Contents

- [1Hardware requirements](#Hardware_requirements)
- [2Media instalacyjne Gentoo Linux](#Media_instalacyjne_Gentoo_Linux)
  - [2.1Minimal installation CD](#Minimal_installation_CD)
  - [2.2What are stage files?](#What_are_stage_files.3F)
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
  - [5.1Default: Booting the installation CD with yaboot](#Default:_Booting_the_installation_CD_with_yaboot)
  - [5.2Alternative: Booting the installation CD on a Pegasos system](#Alternative:_Booting_the_installation_CD_on_a_Pegasos_system)
  - [5.3Alternative: Booting the installation CD with BootX](#Alternative:_Booting_the_installation_CD_with_BootX)
  - [5.4Setting the keyboard map](#Setting_the_keyboard_map)
  - [5.5LiveGUI root access](#LiveGUI_root_access)
  - [5.6Extra hardware configuration](#Extra_hardware_configuration)
  - [5.7Optional: User accounts](#Optional:_User_accounts)
  - [5.8Optional: Viewing documentation while installing](#Optional:_Viewing_documentation_while_installing)
    - [5.8.1TTYs](#TTYs)
    - [5.8.2GNU Screen](#GNU_Screen)
  - [5.9Optional: Starting the SSH daemon](#Optional:_Starting_the_SSH_daemon)

## Hardware requirements

Gentoo can be installed on all manners of systems, so listing minimal hardware specification is more down to how long a user is willing to wait. Any reasonably modern machine can complete the install process in less than an hour, however it is not a race so, take your time, and remember it takes as long as it takes.

For a smooth install process though, the Handbook recommends at least 40GB of space for the root partition.

## Media instalacyjne Gentoo Linux

**Wskazówka**

While it's recommended to use the official Gentoo boot media when installing, it's possible to use other installation environments. However, there is no guarantee they will contain required components. If an alternate install environment is used, skip to [Preparing the disks](/wiki/Handbook:PPC/Installation/Disks/pl "Handbook:PPC/Installation/Disks/pl").

### Minimal installation CD

The _Gentoo minimal installation CD_, also known as the installcd, is a small, bootable image: a self-contained Gentoo environment. This image is maintained by [Gentoo developers](/wiki/Project:RelEng "Project:RelEng") and designed to allow any user with an Internet connection to install Gentoo. During the boot process, the hardware is detected, and appropriate drivers are automatically loaded.

Minimal Installation CD releases are named using the format:
install-<arch>-minimal-<release timestamp>.iso.

**Ważne**

The Gentoo minimal installation CD requires at least 140MB of RAM to boot.

### What are stage files?

A [stage file](/wiki/Stage_file "Stage file") is an archive which serves as the seed for a Gentoo environment.

Stage 3 files can be downloaded from releases/ppc/autobuilds/ on any of the [official Gentoo mirrors](https://www.gentoo.org/downloads/mirrors/). Stages are updated frequently and are therefore not included within official live images.

**Wskazówka**

For now, _stage files_ can be ignored. They will be described in greater detail later when they are needed

**Informacja**

Historically, Gentoo provided stages lower than 3. Over time theses type of installations have become obsolete due to better methods and no longer provided for this reason.

## Downloading

### Obtain the media

The default installation media used by Gentoo Linux are the _minimal installation CDs_, which provide a very small, bootable, Gentoo Linux environment. This environment contains the necessary tools to install Gentoo. The images themselves can be downloaded from the [downloads page](https://www.gentoo.org/downloads/) (recommended) or by manually browsing to the ISO location on one of the many [available mirrors](https://www.gentoo.org/downloads/mirrors/).

#### Navigating Gentoo mirrors

If downloading from a mirror, the minimal installation CDs can be found by:

1. Connect to the mirror, typically using a local one found at [Gentoo source mirrors](https://www.gentoo.org/downloads/mirrors/).
2. Navigate to the releases/ directory.
3. Select the directory for the relevant target architecture (such as **ppc/**).
4. Select the autobuilds/ directory.
5. For **amd64** and **x86** architectures select either the current-install-amd64-minimal/ or current-install-x86-minimal/ directory (respectively). For all other architectures navigate to the current-iso/ directory.

**Informacja**

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

- A .CONTENTS.gz file which is a gz-compressed text file listing all files available on the installation media. This file can be useful to verify if particular firmware or drivers are available on the installation media before downloading it.
- A .DIGESTS file which contains the hash of the ISO file itself, in various hashing formats/algorithms. This file can be used to verify ISO file integrity.
- A .asc file which is a cryptographic signature of the ISO file. This can be used to verify image integrity and authenticity - that the download is indeed provided by the [Gentoo Release Engineering team](/wiki/Project:RelEng "Project:RelEng"), free from tampering.

Ignore the other files available at this location for now - those will come back when the installation has proceeded further. Download the .iso file and, if verification of the download is wanted, download the .iso.asc file for the .iso file as well.

**Wskazówka**

The .DIGESTS file is only needed if the signature in the .iso.asc file is not verified.

### Verifying the downloaded files

**Wskazówka**

This is an optional step and not necessary to install Gentoo Linux. However, it is recommended as it ensures that the downloaded file is not corrupt and has indeed been provided by the [Gentoo Infrastructure team](/wiki/Project:Infrastructure "Project:Infrastructure").

The .asc file provides a cryptographic signature of the ISO. By validating it, one can make sure that
the installation file is provided by the Gentoo Release Engineering team and is intact and unmodified.

#### Microsoft Windows-based verification

To first verify the cryptographic signature, tools such as [GPG4Win](https://www.gpg4win.org/) can be used. After installation, the public keys of the Gentoo Release Engineering team need to be imported. The list of keys is available on the [signatures page](https://www.gentoo.org/downloads/signatures/). Once imported, the user can then verify the signature in the .asc file.

#### Linux based verification

On a Linux system, the most common method for verifying the cryptographic signature is to use the [app-crypt/gnupg](https://packages.gentoo.org/packages/app-crypt/gnupg) software. With this package installed, the following command can be used to verify the cryptographic signature in the .asc file.

**Wskazówka**

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

Alternatively you can use instead the [WKD](/wiki/WKD "WKD") to download the key:

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

`user $` `gpg --verify install-ppc-minimal-20240107T170309Z.iso.asc install-ppc-minimal-20240107T170309Z.iso`

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

**Informacja**

It's generally good practice to mark an imported key as trusted, once it's certain the key is trustworthy. When trusted keys are verified, gpg will not say _unknown_ and warn about the signature being untrusted.

## Writing the boot media

Of course, with just an ISO file downloaded, the Gentoo Linux installation cannot be started. The ISO file must be written to bootable media. This generally requires that the image is extracted to a filesystem, or written directly to a device.

### Writing a bootable USB

**Uwaga**

Using tools such as Ventoy to write to a USB drive can cause boot failures.

Most modern systems support booting from a USB device.

#### Writing with Linux

dd is typically available on most Linux distros, and can be used to write the Gentoo boot media to a USB drive.

##### Determining the USB device path

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

**Wskazówka**

Using the base device path, ie. sdd opposed to sdd1, is recommend as the Gentoo _boot media_ contains a full [GPT](/wiki/GPT "GPT") partition scheme.

##### Writing with dd

**Uwaga**

Be sure to check the target ( _of=target_) path before executing dd, as it will be overwritten.

With the device path (/dev/sdd) and _boot media_ install-amd64-minimal-<release timestamp>.iso ready:

`root #` `dd if=install-amd64-minimal-<release timestamp>.iso of=/dev/sdd bs=4096 status=progress && sync`

**Informacja**

**if=** specifies the _input file_, **of=** specifies the _output file_, which in this case, is a device.

**Wskazówka**

**bs=4096** is used as it speeds up transfers in most cases, **status=progress** displays transfers stats.

#### Writing with Windows

[win32diskimager](https://sourceforge.net/projects/win32diskimager/) is simple to use Windows utility to write iso images to a USB drive.

They also provide excellent documentation to assist if needed.

**Uwaga**

Only download from [https://sourceforge.net/projects/win32diskimager/](https://sourceforge.net/projects/win32diskimager/)

### Burning a disk

**See also**

A more elaborate set of instructions can be found in [CD/DVD/BD\_writing#Image\_writing](/wiki/CD/DVD/BD_writing#Image_writing.2Fpl "CD/DVD/BD writing").

#### Burning with Microsoft Windows 7 and above

Versions of Microsoft Windows 7 and above can both mount and burn ISO images to optical media without the requirement for third-party software. Simply insert a burnable disk, browse to the downloaded ISO files, right click the file in Windows Explorer, and select "Burn disk image".

#### Burning with Linux

The cdrecord utility from the package [app-cdr/cdrtools](https://packages.gentoo.org/packages/app-cdr/cdrtools) can burn ISO images on Linux.

To burn the ISO file on the CD in the /dev/sr0 device (this is the first CD device on the system - substitute with the right device file if necessary):

`user $` `cdrecord dev=/dev/sr0 install-ppc-minimal-20141204.iso`

Users that prefer a graphical user interface can use K3B, part of the [kde-apps/k3b](https://packages.gentoo.org/packages/kde-apps/k3b) package. In K3B, go to Tools and use Burn CD Image.

## Booting

### Default: Booting the installation CD with yaboot

On NewWorld machines place the Installation CD in the CD-ROM and reboot the system. When the system-start-bell sounds, simply hold down the `c` until the CD loads.

After the installation CD loaded, a boot prompt will show up at the bottom of the screen.

We provide one generic kernel, ppc32. This kernel is built with support for multiple CPUs, but it will boot on single processor machines as well.

It is possible to tweak some kernel options at this prompt. The following table lists some of the available boot options you can add:

Boot option
Description
`video`This option takes one of the following vendor-specific tags: nvidiafb, radeonfb, rivafb, atyfb, aty128, or ofonly. Follow this tag with the resolution refresh rate and color depth to use. For instance, `video=radeonfb:1280x1024@75-32` will select the ATI Radeon frame buffer at a resolution of 1280x1024 with a refresh rate of 75Hz and a color depth of 32 bits. When uncertain about what to choose, and the default doesn't work, `video=ofonly` will most certainly work.
`nol3`Disables level 3 cache on some PowerBooks (needed for at least the 17")
`dofirewire`Enables support for IEEE1394 (FireWire) devices, like external hard disks.
`dopcmcia`To use PCMCIA devices during the installation (like PCMCIA network cards) this options needs to be enabled.
`dosshd`Starts the SSH daemon. Useful for unattended installs.
`passwd=foo`Sets the value after the `=` as the root password. Use with `dosshd` for remote installs.

To use the above options, at the boot prompt, type `ppc32` followed by the desired option. In the example below, we'll force the kernel to use the Open Firmware framebuffer instead of the device specific driver.

`boot:` `ppc32 video=ofonly`

If no options are needed, just type ppc32 at this prompt, and a complete Gentoo Linux environment will be loaded from the CD.

### Alternative: Booting the installation CD on a Pegasos system

On the Pegasos simply insert the CD and at the SmartFirmware boot-prompt type `cd /boot/menu`.

`boot` `cd /boot/menu`

This will open a small bootmenu that allows users to choose between several preconfigured video configs. Any special boot options can be appended to the command-line just like with Yaboot above. For example:

`boot` `cd /boot/pegasos video=radeonfb:1280x1024@75 mem=256M`

The default kernel options (in case something goes wrong) are preconfigured with `console=ttyS0,115200 console=tty0 init=/linuxrc looptype=squashfs loop=/image.squashfs cdroot root=/dev/ram0`.

### Alternative: Booting the installation CD with BootX

With an OldWorld Mac the bootable portion of the livecd can't be used. The most simple solution is to use MacOS 9 or earlier to bootstrap into a Linux environment with a tool called BootX.

First, download BootX and unpack the archive. Copy the the BootX Extension from the unpacked archive into Extensions Folder and the BootX App Control Panel into Control Panels, both of which are located in the MacOS System Folder. Next, create a folder called "Linux Kernels" in the System folder and copy the ppc32 kernel from the CD to this folder. Finally, copy ppc32.igz from the Installation CD boot folder into the MacOS System Folder.

To prepare BootX, start the BootX App Control Panel. First select the Options dialog and check Use Specified RAM Disk and select ppc32.igz from the System Folder. Continue back to the initial screen and ensure that the ramdisk size is at least 32000. Finally, set the kernel arguments as shown below:

CODE **BootX kernel arguments**

```
cdroot root=/dev/ram0 init=linuxrc loop=image.squashfs looptype=squashfs console=tty0

```

**Informacja**

The kernel parameters in the yaboot section above are also applicable here. Append any of those options to the kernel arguments above.

Check once more to make sure the settings are correct and then save the configuration. This saves typing just in case it doesn't boot or something is missing. Press the Linux button at the top of the window. If everything goes correctly, it should boot into the Installation CD.

### Setting the keyboard map

After boot, a root ("#") prompt appears on the current console. It is possible to switch to other consoles by pressing `Alt` \+ `F2`, `Alt` \+ `F3` and `Alt` \+ `F4`. Get back to the first one by pressing `Alt` \+ `F1`. Due to the keyboard layout, it may be necessary to press `Alt` \+ `fn` \+ `F#` on Apple machines.

When installing Gentoo on a system with a non-US keyboard, use loadkeys to load the keymap for the keyboard. To list the available keymaps, execute:

`root #` `ls /usr/share/keymaps/i386`

Now load the keymap of choice:

`root #` `loadkeys be-latin1`

### LiveGUI root access

sudo has been configured to run without the need of a password on the LiveGUI as both the root and gentoo have a scrambled password.

To gain access to the superuser account, in any terminal run:

`root #` `sudo -i`

### Extra hardware configuration

When the Installation medium boots, it tries to detect all the hardware devices and loads the appropriate kernel modules to support the hardware. In the vast majority of cases, it does a very good job. However, in some cases it may not auto-load the kernel modules needed by the system. If the PCI auto-detection missed some of the system's hardware, the appropriate kernel modules have to be loaded manually.

In the next example the 8139too module (which supports certain kinds of network interfaces) is loaded:

`root #` `modprobe 8139too`

### Optional: User accounts

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

### Optional: Viewing documentation while installing

#### TTYs

To view the Gentoo handbook from a TTY during the installation, first create a user account as described above, then press `Alt` + `F2` to go to a new terminal (TTY) and login as the newly created user. Following the [principle of least privilege](https://en.wikipedia.org/wiki/Principle_of_least_privilege "wikipedia:Principle of least privilege"), it is best practice to avoid browsing the web or generally performing any task with higher privileges than necessary. The root account has full control of the system and therefore must be used sparingly.

During the installation, the links web browser can be used to browse the Gentoo handbook - of course only from the moment that the Internet connection is working.

`user $` `links https://wiki.gentoo.org/wiki/Handbook:PPC`

To go back to the original terminal, press `Alt` + `F1`.

**Wskazówka**

When booted to the Gentoo minimal or Gentoo admin environments, seven TTYs will be available. They can be switched by pressing `Alt` then a function key between `F1`- `F7`. It can be useful to switch to a new terminal when waiting for job to complete, to open documentation, etc.

#### GNU Screen

The [Screen](/wiki/Screen "Screen") utility is installed by default on official Gentoo installation media. It may be more efficient for the seasoned Linux enthusiast to use screen to view installation instructions via split panes rather than the multiple TTY method mentioned above.

### Optional: Starting the SSH daemon

To allow other users to access the system during the installation (perhaps to provide/receive support during an installation, or even do it remotely), a user account needs to be created (as was documented earlier on) and the SSH daemon needs to be started.

To fire up the SSH daemon on an OpenRC init, execute the following command:

`root #` `rc-service sshd start`

**Informacja**

If users log on to the system, they will see a message that the host key for this system needs to be confirmed (through what is called a fingerprint). This behavior is typical and can be expected for initial connections to an SSH server. However, later when the system is set up and someone logs on to the newly created system, the SSH client will warn that the host key has been changed. This is because the user now logs on to - for SSH - a different server (namely the freshly installed Gentoo system rather than the live environment that the installation is currently using). Follow the instructions given on the screen then to replace the host key on the client system.

To be able to use sshd, the network needs to function properly. Continue with the chapter on [Configuring the network](/wiki/Handbook:PPC/Installation/Networking/pl "Handbook:PPC/Installation/Networking/pl").

[← About the installation](/wiki/Handbook:PPC/Installation/About "Previous part") [Home](/wiki/Handbook:PPC "Handbook:PPC") [Configuring the network →](/wiki/Handbook:PPC/Installation/Networking "Next part")