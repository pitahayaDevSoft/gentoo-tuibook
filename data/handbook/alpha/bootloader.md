# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Bootloader/de "Handbuch:Alpha/Installation/Bootloader (100% translated)")
- English
- [español](/wiki/Handbook:Alpha/Installation/Bootloader/es "Manual de Gentoo: Alpha/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Bootloader/fr "Manuel:Alpha/Installation/Système d'amorçage (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Bootloader/it "Handbook:Alpha/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Bootloader/hu "Handbook:Alpha/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Bootloader/pl "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Bootloader/pt-br "Manual:Alpha/Instalação/Gerenciador de boot (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Bootloader/cs "Handbook:Alpha/Installation/Bootloader/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Bootloader/ru "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Bootloader/ta "கையேடு:Alpha/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Bootloader/zh-cn "手册：Alpha/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Bootloader/ja "ハンドブック:Alpha/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Bootloader/ko "Handbook:Alpha/Installation/Bootloader/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha "Handbook:Alpha")[Installation](/wiki/Handbook:Alpha/Full/Installation "Handbook:Alpha/Full/Installation")[About the installation](/wiki/Handbook:Alpha/Installation/About "Handbook:Alpha/Installation/About")[Choosing the media](/wiki/Handbook:Alpha/Installation/Media "Handbook:Alpha/Installation/Media")[Configuring the network](/wiki/Handbook:Alpha/Installation/Networking "Handbook:Alpha/Installation/Networking")[Preparing the disks](/wiki/Handbook:Alpha/Installation/Disks "Handbook:Alpha/Installation/Disks")[The stage file](/wiki/Handbook:Alpha/Installation/Stage "Handbook:Alpha/Installation/Stage")[Installing base system](/wiki/Handbook:Alpha/Installation/Base "Handbook:Alpha/Installation/Base")[Configuring the kernel](/wiki/Handbook:Alpha/Installation/Kernel "Handbook:Alpha/Installation/Kernel")[Configuring the system](/wiki/Handbook:Alpha/Installation/System "Handbook:Alpha/Installation/System")[Installing tools](/wiki/Handbook:Alpha/Installation/Tools "Handbook:Alpha/Installation/Tools")Configuring the bootloader[Finalizing](/wiki/Handbook:Alpha/Installation/Finalizing "Handbook:Alpha/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:Alpha/Full/Working "Handbook:Alpha/Full/Working")[Portage introduction](/wiki/Handbook:Alpha/Working/Portage "Handbook:Alpha/Working/Portage")[USE flags](/wiki/Handbook:Alpha/Working/USE "Handbook:Alpha/Working/USE")[Portage features](/wiki/Handbook:Alpha/Working/Features "Handbook:Alpha/Working/Features")[Initscript system](/wiki/Handbook:Alpha/Working/Initscripts "Handbook:Alpha/Working/Initscripts")[Environment variables](/wiki/Handbook:Alpha/Working/EnvVar "Handbook:Alpha/Working/EnvVar")[Working with Portage](/wiki/Handbook:Alpha/Full/Portage "Handbook:Alpha/Full/Portage")[Files and directories](/wiki/Handbook:Alpha/Portage/Files "Handbook:Alpha/Portage/Files")[Variables](/wiki/Handbook:Alpha/Portage/Variables "Handbook:Alpha/Portage/Variables")[Mixing software branches](/wiki/Handbook:Alpha/Portage/Branches "Handbook:Alpha/Portage/Branches")[Additional tools](/wiki/Handbook:Alpha/Portage/Tools "Handbook:Alpha/Portage/Tools")[Custom package repository](/wiki/Handbook:Alpha/Portage/CustomTree "Handbook:Alpha/Portage/CustomTree")[Advanced features](/wiki/Handbook:Alpha/Portage/Advanced "Handbook:Alpha/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:Alpha/Full/Networking "Handbook:Alpha/Full/Networking")[Getting started](/wiki/Handbook:Alpha/Networking/Introduction "Handbook:Alpha/Networking/Introduction")[Advanced configuration](/wiki/Handbook:Alpha/Networking/Advanced "Handbook:Alpha/Networking/Advanced")[Modular networking](/wiki/Handbook:Alpha/Networking/Modular "Handbook:Alpha/Networking/Modular")[Wireless](/wiki/Handbook:Alpha/Networking/Wireless "Handbook:Alpha/Networking/Wireless")[Adding functionality](/wiki/Handbook:Alpha/Networking/Extending "Handbook:Alpha/Networking/Extending")[Dynamic management](/wiki/Handbook:Alpha/Networking/Dynamic "Handbook:Alpha/Networking/Dynamic")

## Making a choice\[ [edit](/index.php?title=Handbook:Alpha/Blocks/Bootloader&action=edit&section=T-1 "Edit section: ")\]

Now that the kernel is configured and compiled and the necessary system configuration files are filled in correctly, it is time to install a program that will fire up the kernel when the system is started. Such a program is called a bootloader.

Several bootloaders exist for Linux/Alpha. Choose one of the supported bootloaders, not all. We document [aBoot](#Default:_Using_aboot) and [MILO](#Alternative:_Using_MILO).

## Default: Using aBoot\[ [edit](/index.php?title=Handbook:Alpha/Blocks/Bootloader&action=edit&section=T-2 "Edit section: ")\]

**Note**

aboot only supports booting from ext2 and ext3 partitions.

First install aboot on the system

`root #` `emerge --ask sys-boot/aboot`

The next step is to make the bootdisk bootable. This will start aboot when booting the system. We make our bootdisk bootable by writing the aboot bootloader to the start of the disk.

`root #` `swriteboot -f3 /dev/sda /boot/bootlx
`

`root #` `abootconf /dev/sda 2
`

**Note**

When using a different partitioning scheme than the one used throughout this chapter, the commands need to be changed accordingly. Please read the appropriate manual pages (man 8 swriteboot and man 8 abootconf). Also, if the root filesystem is ran using the JFS filesystem, make sure it gets mounted read-only at first by adding ro as a kernel option.

Although aboot is now installed, we still need to write a configuration file for it. aboot only requires one line for each configuration, so we can do this:

`root #` `echo '0:2/boot/vmlinux.gz root=/dev/sda3' > /etc/aboot.conf`

If, while building the Linux kernel, an initramfs was build as well to boot from, then it is necessary to change the configuration by referring to this initramfs file and telling the initramfs where the real root device is at:

`root #` `echo '0:2/boot/vmlinux.gz initrd=/boot/initramfs-genkernel-alpha-6.19.1-gentoo root=/dev/sda3' > /etc/aboot.conf`

Additionally, it is possible to make Gentoo boot automatically by setting up some SRM variables. Try setting these variables from Linux, but it may be easier to do so from the SRM console itself.

`root #` `cd /proc/srm_environment/named_variables
`

`root #` `echo -n 0 > boot_osflags
`

`root #` `echo -n '' > boot_file
`

`root #` `echo -n 'BOOT' > auto_action
`

`root #` `echo -n 'dkc100' > bootdef_dev`

Of course substitute dkc100 with whatever the boot device is.

To get in the SRM console again in the future (to recover the Gentoo install, play with some variables, or whatever), just hit `Ctrl+C` to abort the automatic loading process.

When installing using a serial console, don't forget to include the serial console boot flag in aboot.conf. See /etc/aboot.conf.example for some further information.

Aboot is now configured and ready to use. Continue with [Rebooting the system](#Rebooting_the_system).

Now continue with [Rebooting the system](#Rebooting_the_system).

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

Once rebooted in the fresh Gentoo environment, continue to the [Finalizing the Gentoo installation](/wiki/Handbook:Alpha/Installation/Finalizing "Handbook:Alpha/Installation/Finalizing") for important information on setting up the new installation, and finally some tips on how to start a productive Gentoo Linux journey.

[← Installing tools](/wiki/Handbook:Alpha/Installation/Tools "Previous part") [Home](/wiki/Handbook:Alpha "Handbook:Alpha") [Finalizing →](/wiki/Handbook:Alpha/Installation/Finalizing "Next part")