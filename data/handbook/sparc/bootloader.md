# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbuch:SPARC/Installation/Bootloader (100% translated)")
- English
- [español](/wiki/Handbook:SPARC/Installation/Bootloader/es "Manual de Gentoo: SPARC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Bootloader/fr "Handbook:SPARC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Bootloader/pl "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Bootloader/pt-br "Handbook:SPARC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Bootloader/ru "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Bootloader/ta "கையேடு:SPARC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Bootloader/zh-cn "手册：SPARC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Bootloader/ja "ハンドブック:SPARC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Bootloader/ko "Handbook:SPARC/Installation/Bootloader/ko (100% translated)")

[SPARC Handbook](/wiki/Handbook:SPARC "Handbook:SPARC")[Installation](/wiki/Handbook:SPARC/Full/Installation "Handbook:SPARC/Full/Installation")[About the installation](/wiki/Handbook:SPARC/Installation/About "Handbook:SPARC/Installation/About")[Choosing the media](/wiki/Handbook:SPARC/Installation/Media "Handbook:SPARC/Installation/Media")[Configuring the network](/wiki/Handbook:SPARC/Installation/Networking "Handbook:SPARC/Installation/Networking")[Preparing the disks](/wiki/Handbook:SPARC/Installation/Disks "Handbook:SPARC/Installation/Disks")[The stage file](/wiki/Handbook:SPARC/Installation/Stage "Handbook:SPARC/Installation/Stage")[Installing base system](/wiki/Handbook:SPARC/Installation/Base "Handbook:SPARC/Installation/Base")[Configuring the kernel](/wiki/Handbook:SPARC/Installation/Kernel "Handbook:SPARC/Installation/Kernel")[Configuring the system](/wiki/Handbook:SPARC/Installation/System "Handbook:SPARC/Installation/System")[Installing tools](/wiki/Handbook:SPARC/Installation/Tools "Handbook:SPARC/Installation/Tools")Configuring the bootloader[Finalizing](/wiki/Handbook:SPARC/Installation/Finalizing "Handbook:SPARC/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:SPARC/Full/Working "Handbook:SPARC/Full/Working")[Portage introduction](/wiki/Handbook:SPARC/Working/Portage "Handbook:SPARC/Working/Portage")[USE flags](/wiki/Handbook:SPARC/Working/USE "Handbook:SPARC/Working/USE")[Portage features](/wiki/Handbook:SPARC/Working/Features "Handbook:SPARC/Working/Features")[Initscript system](/wiki/Handbook:SPARC/Working/Initscripts "Handbook:SPARC/Working/Initscripts")[Environment variables](/wiki/Handbook:SPARC/Working/EnvVar "Handbook:SPARC/Working/EnvVar")[Working with Portage](/wiki/Handbook:SPARC/Full/Portage "Handbook:SPARC/Full/Portage")[Files and directories](/wiki/Handbook:SPARC/Portage/Files "Handbook:SPARC/Portage/Files")[Variables](/wiki/Handbook:SPARC/Portage/Variables "Handbook:SPARC/Portage/Variables")[Mixing software branches](/wiki/Handbook:SPARC/Portage/Branches "Handbook:SPARC/Portage/Branches")[Additional tools](/wiki/Handbook:SPARC/Portage/Tools "Handbook:SPARC/Portage/Tools")[Custom package repository](/wiki/Handbook:SPARC/Portage/CustomTree "Handbook:SPARC/Portage/CustomTree")[Advanced features](/wiki/Handbook:SPARC/Portage/Advanced "Handbook:SPARC/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:SPARC/Full/Networking "Handbook:SPARC/Full/Networking")[Getting started](/wiki/Handbook:SPARC/Networking/Introduction "Handbook:SPARC/Networking/Introduction")[Advanced configuration](/wiki/Handbook:SPARC/Networking/Advanced "Handbook:SPARC/Networking/Advanced")[Modular networking](/wiki/Handbook:SPARC/Networking/Modular "Handbook:SPARC/Networking/Modular")[Wireless](/wiki/Handbook:SPARC/Networking/Wireless "Handbook:SPARC/Networking/Wireless")[Adding functionality](/wiki/Handbook:SPARC/Networking/Extending "Handbook:SPARC/Networking/Extending")[Dynamic management](/wiki/Handbook:SPARC/Networking/Dynamic "Handbook:SPARC/Networking/Dynamic")

## Contents

- [1GRUB](#GRUB)
  - [1.1Emerge](#Emerge)
  - [1.2Install](#Install)
    - [1.2.1GPT](#GPT)
    - [1.2.2Sun partition table](#Sun_partition_table)
  - [1.3Configure](#Configure)
- [2SILO, the SPARC bootloader](#SILO.2C_the_SPARC_bootloader)
- [3Rebooting the system](#Rebooting_the_system)

## GRUB\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Bootloader&action=edit&section=T-1 "Edit section: ")\]

When [selecting a 64-bit profile](/wiki/Handbook:SPARC/Installation/Base#Choosing_the_right_profile "Handbook:SPARC/Installation/Base") during installation, then [GRUB](/wiki/GRUB "GRUB") is the only supported bootloader.

### Emerge\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Bootloader&action=edit&section=T-2 "Edit section: ")\]

GRUB should be correctly configured for the platform automatically based on the profile. To make it explicit, however, specify it using:

`root #` `echo 'GRUB_PLATFORMS="ieee1275"' >> /etc/portage/make.conf`

`root #` `emerge --ask --verbose sys-boot/grub`

The GRUB software has now been merged to the system, but not yet installed as a bootloader.

### Install\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Bootloader&action=edit&section=T-3 "Edit section: ")\]

#### GPT\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Bootloader&action=edit&section=T-4 "Edit section: ")\]

If the [disk is partitioned](/wiki/Handbook:SPARC/Installation/Disks#Using_fdisk_to_partition_the_disk "Handbook:SPARC/Installation/Disks") using GPT (the preferred method), then install GRUB to the BIOS boot partition. Presuming the first disk (the one where the system boots from) is /dev/sda, the following commands will do:

`root #` `grub-install --target=sparc64-ieee1275 --recheck /dev/sda`

**Tip**

To find the boot device string to enter in the firmware, use the grub-ofpathname tool. If the BIOS boot partition is the first partition on the disk, select the entire disk:

`root #` `grub-ofpathname /dev/sda`

Otherwise, explicitly select the BIOS boot partition:

`root #` `grub-ofpathname /dev/sda2`

#### Sun partition table\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Bootloader&action=edit&section=T-5 "Edit section: ")\]

If the disk is partitioned using a Sun partition table instead, GRUB must be installed using blocklists. In this mode, instead of providing the physical disk as an argument, provide the path to the partition on which /boot/grub is mounted.

`root #` `grub-install --target=sparc64-ieee1275 --recheck --force --skip-fs-probe /dev/sda1`

### Configure\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Bootloader&action=edit&section=T-6 "Edit section: ")\]

Next, generate the GRUB2 configuration based on the user configuration specified in the /etc/default/grub file and /etc/grub.d scripts. In most cases, no configuration is needed by users as GRUB2 will automatically detect which kernel to boot (the highest one available in /boot/) and what the root file system is. It is also possible to append kernel parameters in /etc/default/grub using the GRUB\_CMDLINE\_LINUX variable.

To generate the final GRUB2 configuration, run the grub-mkconfig command:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.19.3-gentoo
Found initrd image: /boot/initramfs-genkernel-sparc-6.19.3-gentoo
done

```

The output of the command must mention that at least one Linux image is found, as those are needed to boot the system. If an initramfs is used or genkernel was used to build the kernel, the correct initrd image should be detected as well. If this is not the case, go to /boot/ and check the contents using the ls command. If the files are indeed missing, go back to the kernel configuration and installation instructions.

## SILO, the SPARC bootloader\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Bootloader&action=edit&section=T-7 "Edit section: ")\]

When [selecting a 32-bit profile](/wiki/Handbook:SPARC/Installation/Base#Choosing_the_right_profile "Handbook:SPARC/Installation/Base") during installation, then [SILO](https://git.kernel.org/?p=linux/kernel/git/davem/silo.git;a=summary) (Sparc Improved boot LOader) is the only supported bootloader.

`root #` `emerge --ask sys-boot/silo`

Next create /etc/silo.conf:

`root #` `nano -w /etc/silo.conf`

Below an example silo.conf file is shown. It uses the partitioning scheme we use throughout this book, kernel-6.19.3-gentoo as kernel image and initramfs-genkernel-sparc64-6.19.3-gentoo as initramfs.

FILE **`/etc/silo.conf`** **Example configuration file**

```
partition = 1         # Boot partition (= root partition)
root = /dev/sda1      # Root partition
timeout = 150         # Wait 15 seconds before booting the default section

image = /boot/kernel-6.19.3-gentoo
  label = linux
  append = "initrd=/boot/initramfs-genkernel-sparc64-6.19.3-gentoo root=/dev/sda1"

```

When using the example silo.conf file as delivered by Portage, be sure to comment out all lines that aren't needed.

If the physical disk on which to install SILO (as bootloader) differs from the physical disk on which /etc/silo.conf resides, then first copy over /etc/silo.conf to a partition on that disk. If /boot/ is a separate partition on that disk, copy over the configuration file to /boot/ and run /sbin/silo:

`root #` `cp /etc/silo.conf /boot
`

`root #` `/sbin/silo -C /boot/silo.conf`

```
/boot/silo.conf appears to be valid

```

Otherwise just run /sbin/silo:

`root #` `/sbin/silo`

```
/etc/silo.conf appears to be valid

```

**Note**

Run silo (with parameters if necessary) again each time after updating or installing the [sys-boot/silo](https://packages.gentoo.org/packages/sys-boot/silo) package.

## Rebooting the system\[ [edit](/index.php?title=Handbook:Parts/Installation/Bootloader&action=edit&section=T-1 "Edit section: ")\]

Exit the chrooted environment and unmount all mounted partitions. Then type in that one magical command that initiates the final, true test: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Do not forget to remove the live image, otherwise it may be targeted again instead of the newly installed Gentoo system!

Once rebooted in the fresh Gentoo environment, continue to the [Finalizing the Gentoo installation](/wiki/Handbook:SPARC/Installation/Finalizing "Handbook:SPARC/Installation/Finalizing") for important information on setting up the new installation, and finally some tips on how to start a productive Gentoo Linux journey.

[← Installing tools](/wiki/Handbook:SPARC/Installation/Tools "Previous part") [Home](/wiki/Handbook:SPARC "Handbook:SPARC") [Finalizing →](/wiki/Handbook:SPARC/Installation/Finalizing "Next part")