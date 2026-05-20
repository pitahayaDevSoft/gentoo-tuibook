# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbuch:HPPA/Installation/Bootloader (100% translated)")
- English
- [español](/wiki/Handbook:HPPA/Installation/Bootloader/es "Manual de Gentoo: HPPA/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Bootloader/it "Handbook:HPPA/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Bootloader/pl "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Bootloader/pt-br "Manual:HPPA/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Bootloader/ru "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Bootloader/ta "கையேடு:HPPA/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Bootloader/zh-cn "手册：HPPA/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Bootloader/ja "ハンドブック:HPPA/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko (100% translated)")

[HPPA Handbook](/wiki/Handbook:HPPA "Handbook:HPPA")[Installation](/wiki/Handbook:HPPA/Full/Installation "Handbook:HPPA/Full/Installation")[About the installation](/wiki/Handbook:HPPA/Installation/About "Handbook:HPPA/Installation/About")[Choosing the media](/wiki/Handbook:HPPA/Installation/Media "Handbook:HPPA/Installation/Media")[Configuring the network](/wiki/Handbook:HPPA/Installation/Networking "Handbook:HPPA/Installation/Networking")[Preparing the disks](/wiki/Handbook:HPPA/Installation/Disks "Handbook:HPPA/Installation/Disks")[The stage file](/wiki/Handbook:HPPA/Installation/Stage "Handbook:HPPA/Installation/Stage")[Installing base system](/wiki/Handbook:HPPA/Installation/Base "Handbook:HPPA/Installation/Base")[Configuring the kernel](/wiki/Handbook:HPPA/Installation/Kernel "Handbook:HPPA/Installation/Kernel")[Configuring the system](/wiki/Handbook:HPPA/Installation/System "Handbook:HPPA/Installation/System")[Installing tools](/wiki/Handbook:HPPA/Installation/Tools "Handbook:HPPA/Installation/Tools")Configuring the bootloader[Finalizing](/wiki/Handbook:HPPA/Installation/Finalizing "Handbook:HPPA/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:HPPA/Full/Working "Handbook:HPPA/Full/Working")[Portage introduction](/wiki/Handbook:HPPA/Working/Portage "Handbook:HPPA/Working/Portage")[USE flags](/wiki/Handbook:HPPA/Working/USE "Handbook:HPPA/Working/USE")[Portage features](/wiki/Handbook:HPPA/Working/Features "Handbook:HPPA/Working/Features")[Initscript system](/wiki/Handbook:HPPA/Working/Initscripts "Handbook:HPPA/Working/Initscripts")[Environment variables](/wiki/Handbook:HPPA/Working/EnvVar "Handbook:HPPA/Working/EnvVar")[Working with Portage](/wiki/Handbook:HPPA/Full/Portage "Handbook:HPPA/Full/Portage")[Files and directories](/wiki/Handbook:HPPA/Portage/Files "Handbook:HPPA/Portage/Files")[Variables](/wiki/Handbook:HPPA/Portage/Variables "Handbook:HPPA/Portage/Variables")[Mixing software branches](/wiki/Handbook:HPPA/Portage/Branches "Handbook:HPPA/Portage/Branches")[Additional tools](/wiki/Handbook:HPPA/Portage/Tools "Handbook:HPPA/Portage/Tools")[Custom package repository](/wiki/Handbook:HPPA/Portage/CustomTree "Handbook:HPPA/Portage/CustomTree")[Advanced features](/wiki/Handbook:HPPA/Portage/Advanced "Handbook:HPPA/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:HPPA/Full/Networking "Handbook:HPPA/Full/Networking")[Getting started](/wiki/Handbook:HPPA/Networking/Introduction "Handbook:HPPA/Networking/Introduction")[Advanced configuration](/wiki/Handbook:HPPA/Networking/Advanced "Handbook:HPPA/Networking/Advanced")[Modular networking](/wiki/Handbook:HPPA/Networking/Modular "Handbook:HPPA/Networking/Modular")[Wireless](/wiki/Handbook:HPPA/Networking/Wireless "Handbook:HPPA/Networking/Wireless")[Adding functionality](/wiki/Handbook:HPPA/Networking/Extending "Handbook:HPPA/Networking/Extending")[Dynamic management](/wiki/Handbook:HPPA/Networking/Dynamic "Handbook:HPPA/Networking/Dynamic")

## Installing PALO\[ [edit](/index.php?title=Handbook:HPPA/Blocks/Bootloader&action=edit&section=T-1 "Edit section: ")\]

On the PA-RISC platform, the boot loader is called palo. First merge this bootloader to the system:

`root #` `emerge --ask sys-boot/palo`

The configuration file will be found at /etc/palo.conf. Below is a sample configuration:

**Warning**

This configuration **must** be changed after running palo for the first time! See below for after first setup.

FILE **`/etc/palo.conf`** **Simple PALO configuration example**

```
--commandline=2/kernel-6.19.1-gentoo root=/dev/sda4
--recoverykernel=/vmlinux.old
# DELETE this line after running palo for the first time!
--init-partitioned=/dev/sda
# --format-as has two meanings depending on whether --init-partitioned or --update-partitioned is used. Keep this line.
--format-as=4

```

The first line tells palo the location of the kernel and which boot parameters it must use. The string `2/kernel-6.19.1-gentoo` means the kernel named /kernel-6.19.1-gentoo resides on the second partition. Beware, the path to the kernel is relative to the _boot_ partition, not to the root partition.

The second line indicates which recovery kernel to use. If it is the first install and there is no recovery kernel (yet), please comment this out. The third line indicates on which disk palo will reside.

To format the disk, palo must be run with certain arguments. This example uses _ext4_ for the first partition:

`root #` `palo --format-as=4 --init-partitioned=/dev/sda`

When configuration is done, simply run the palo command:

`root #` `palo`

The configuration must then be updated for post-first-install use:

FILE **`/etc/palo.conf`** **Simple PALO configuration example**

```
--commandline=2/kernel-6.19.1-gentoo root=/dev/sda4
--recoverykernel=/vmlinux.old
# Don't throw away the old partition, just update the existing one on `palo` runs.
--update-partitioned=/dev/sda
# --format-as has two meanings depending on whether --init-partitioned or --update-partitioned is used. Keep this line.
--format-as=4

```

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

Once rebooted in the fresh Gentoo environment, continue to the [Finalizing the Gentoo installation](/wiki/Handbook:HPPA/Installation/Finalizing "Handbook:HPPA/Installation/Finalizing") for important information on setting up the new installation, and finally some tips on how to start a productive Gentoo Linux journey.

[← Installing tools](/wiki/Handbook:HPPA/Installation/Tools "Previous part") [Home](/wiki/Handbook:HPPA "Handbook:HPPA") [Finalizing →](/wiki/Handbook:HPPA/Installation/Finalizing "Next part")