# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Bootloader/de "Handbuch:Alpha/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Bootloader "Handbook:Alpha/Installation/Bootloader (100% translated)")
- español
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

[Alpha Handbook](/wiki/Handbook:Alpha "Handbook:Alpha")[Installation](/wiki/Handbook:Alpha/Full/Installation "Handbook:Alpha/Full/Installation")[About the installation](/wiki/Handbook:Alpha/Installation/About/es "Handbook:Alpha/Installation/About/es")[Choosing the media](/wiki/Handbook:Alpha/Installation/Media/es "Handbook:Alpha/Installation/Media/es")[Configuring the network](/wiki/Handbook:Alpha/Installation/Networking/es "Handbook:Alpha/Installation/Networking/es")[Preparing the disks](/wiki/Handbook:Alpha/Installation/Disks/es "Handbook:Alpha/Installation/Disks/es")[The stage file](/wiki/Handbook:Alpha/Installation/Stage/es "Handbook:Alpha/Installation/Stage/es")[Installing base system](/wiki/Handbook:Alpha/Installation/Base/es "Handbook:Alpha/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:Alpha/Installation/Kernel/es "Handbook:Alpha/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:Alpha/Installation/System/es "Handbook:Alpha/Installation/System/es")[Installing tools](/wiki/Handbook:Alpha/Installation/Tools/es "Handbook:Alpha/Installation/Tools/es")Configuring the bootloader[Finalizing](/wiki/Handbook:Alpha/Installation/Finalizing/es "Handbook:Alpha/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:Alpha/Full/Working "Handbook:Alpha/Full/Working")[Portage introduction](/wiki/Handbook:Alpha/Working/Portage/es "Handbook:Alpha/Working/Portage/es")[USE flags](/wiki/Handbook:Alpha/Working/USE/es "Handbook:Alpha/Working/USE/es")[Portage features](/wiki/Handbook:Alpha/Working/Features/es "Handbook:Alpha/Working/Features/es")[Initscript system](/wiki/Handbook:Alpha/Working/Initscripts/es "Handbook:Alpha/Working/Initscripts/es")[Environment variables](/wiki/Handbook:Alpha/Working/EnvVar/es "Handbook:Alpha/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:Alpha/Full/Portage "Handbook:Alpha/Full/Portage")[Files and directories](/wiki/Handbook:Alpha/Portage/Files/es "Handbook:Alpha/Portage/Files/es")[Variables](/wiki/Handbook:Alpha/Portage/Variables/es "Handbook:Alpha/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:Alpha/Portage/Branches/es "Handbook:Alpha/Portage/Branches/es")[Additional tools](/wiki/Handbook:Alpha/Portage/Tools/es "Handbook:Alpha/Portage/Tools/es")[Custom package repository](/wiki/Handbook:Alpha/Portage/CustomTree/es "Handbook:Alpha/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:Alpha/Portage/Advanced/es "Handbook:Alpha/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:Alpha/Full/Networking "Handbook:Alpha/Full/Networking")[Getting started](/wiki/Handbook:Alpha/Networking/Introduction/es "Handbook:Alpha/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:Alpha/Networking/Advanced/es "Handbook:Alpha/Networking/Advanced/es")[Modular networking](/wiki/Handbook:Alpha/Networking/Modular/es "Handbook:Alpha/Networking/Modular/es")[Wireless](/wiki/Handbook:Alpha/Networking/Wireless/es "Handbook:Alpha/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:Alpha/Networking/Extending/es "Handbook:Alpha/Networking/Extending/es")[Dynamic management](/wiki/Handbook:Alpha/Networking/Dynamic/es "Handbook:Alpha/Networking/Dynamic/es")

## Making a choice

Now that the kernel is configured and compiled and the necessary system configuration files are filled in correctly, it is time to install a program that will fire up the kernel when the system is started. Such a program is called a bootloader.

Several bootloaders exist for Linux/Alpha. Choose one of the supported bootloaders, not all. We document [aBoot](#Default:_Using_aboot) and [MILO](#Alternative:_Using_MILO).

## Default: Using aBoot

**Nota**

aboot only supports booting from ext2 and ext3 partitions.

First install aboot on the system

`root #` `emerge --ask sys-boot/aboot`

The next step is to make the bootdisk bootable. This will start aboot when booting the system. We make our bootdisk bootable by writing the aboot bootloader to the start of the disk.

`root #` `swriteboot -f3 /dev/sda /boot/bootlx
`

`root #` `abootconf /dev/sda 2
`

**Nota**

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

## Rebooting the system

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

Once rebooted in the fresh Gentoo environment, continue to the [Finalizing the Gentoo installation](/wiki/Handbook:Alpha/Installation/Finalizing/es "Handbook:Alpha/Installation/Finalizing/es") for important information on setting up the new installation, and finally some tips on how to start a productive Gentoo Linux journey.

[← Installing tools](/wiki/Handbook:Alpha/Installation/Tools "Previous part") [Home](/wiki/Handbook:Alpha "Handbook:Alpha") [Finalizing →](/wiki/Handbook:Alpha/Installation/Finalizing "Next part")