# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbuch:PPC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Bootloader "Handbook:PPC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Bootloader/es "Manual de Gentoo: PPC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Bootloader/fr "Handbook:PPC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Bootloader/it "Handbook:PPC/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Bootloader/hu "Handbook:PPC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Bootloader/pl "Handbook:PPC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Bootloader/pt-br "Handbook:PPC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Bootloader/ru "Handbook:PPC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Bootloader/ta "கையேடு:PPC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Bootloader/zh-cn "手册：PPC/安装/配置系统引导程序Bootloader (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:PPC/Installation/Bootloader/ko "Handbook:PPC/Installation/Bootloader/ko (100% translated)")

[PPC ハンドブック](/wiki/Handbook:PPC "Handbook:PPC")[インストール](/wiki/Handbook:PPC/Full/Installation "Handbook:PPC/Full/Installation")[インストールについて](/wiki/Handbook:PPC/Installation/About/ja "Handbook:PPC/Installation/About/ja")[メディアの選択](/wiki/Handbook:PPC/Installation/Media/ja "Handbook:PPC/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:PPC/Installation/Networking/ja "Handbook:PPC/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:PPC/Installation/Disks/ja "Handbook:PPC/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:PPC/Installation/Stage/ja "Handbook:PPC/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:PPC/Installation/Base/ja "Handbook:PPC/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:PPC/Installation/Kernel/ja "Handbook:PPC/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:PPC/Installation/System/ja "Handbook:PPC/Installation/System/ja")[ツールのインストール](/wiki/Handbook:PPC/Installation/Tools/ja "Handbook:PPC/Installation/Tools/ja")ブートローダの設定[締めくくり](/wiki/Handbook:PPC/Installation/Finalizing/ja "Handbook:PPC/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:PPC/Full/Working "Handbook:PPC/Full/Working")[Portage について](/wiki/Handbook:PPC/Working/Portage/ja "Handbook:PPC/Working/Portage/ja")[USE フラグ](/wiki/Handbook:PPC/Working/USE/ja "Handbook:PPC/Working/USE/ja")[Portage の機能](/wiki/Handbook:PPC/Working/Features/ja "Handbook:PPC/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:PPC/Working/Initscripts/ja "Handbook:PPC/Working/Initscripts/ja")[環境変数](/wiki/Handbook:PPC/Working/EnvVar/ja "Handbook:PPC/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:PPC/Full/Portage "Handbook:PPC/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:PPC/Portage/Files/ja "Handbook:PPC/Portage/Files/ja")[変数](/wiki/Handbook:PPC/Portage/Variables/ja "Handbook:PPC/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:PPC/Portage/Branches/ja "Handbook:PPC/Portage/Branches/ja")[追加ツール](/wiki/Handbook:PPC/Portage/Tools/ja "Handbook:PPC/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:PPC/Portage/CustomTree/ja "Handbook:PPC/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:PPC/Portage/Advanced/ja "Handbook:PPC/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:PPC/Full/Networking "Handbook:PPC/Full/Networking")[はじめに](/wiki/Handbook:PPC/Networking/Introduction/ja "Handbook:PPC/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:PPC/Networking/Advanced/ja "Handbook:PPC/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:PPC/Networking/Modular/ja "Handbook:PPC/Networking/Modular/ja")[無線](/wiki/Handbook:PPC/Networking/Wireless/ja "Handbook:PPC/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:PPC/Networking/Extending/ja "Handbook:PPC/Networking/Extending/ja")[動的な管理](/wiki/Handbook:PPC/Networking/Dynamic/ja "Handbook:PPC/Networking/Dynamic/ja")

## Contents

- [1ブートローダの選択肢](#.E3.83.96.E3.83.BC.E3.83.88.E3.83.AD.E3.83.BC.E3.83.80.E3.81.AE.E9.81.B8.E6.8A.9E.E8.82.A2)
- [2NewWorld Macs](#NewWorld_Macs)
  - [2.1GRUB](#GRUB)
  - [2.2インストール](#.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB)
  - [2.3Setup bootstrap partition](#Setup_bootstrap_partition)
  - [2.4Setup GRUB](#Setup_GRUB)
- [3OldWorld Macs](#OldWorld_Macs)
  - [3.1BootX](#BootX)
- [4Pegasos](#Pegasos)
  - [4.1BootCreator](#BootCreator)
- [5システムのリブート](#.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.81.AE.E3.83.AA.E3.83.96.E3.83.BC.E3.83.88)

## ブートローダの選択肢

Now that the kernel is configured and compiled and the necessary system configuration files are filled in correctly, it is time to install a program that will fire up the kernel when the system is started. Such a program is called a boot loader.

The boot loader to use depends upon the type of PPC machine.

For a NewWorld Apple or IBM machine, grub needs to be selected. OldWorld Apple machines havs one option: BootX. The Pegasos does not require a boot loader, but it is necessary to emerge bootcreator to create SmartFirmware boot menus.

## NewWorld Macs

### GRUB

### インストール

`root #` `emerge --ask sys-boot/grub`

### Setup bootstrap partition

First, prepare the bootstrap partition that was created created during the [preparing the disk](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks") step. Following the example, this partition should be /dev/sda2. Optionally, confirm this by using parted:

Replace /dev/sda with the correct device if required.

`root #` `parted /dev/sda print`

```
Model: ATA Patriot Burst El (scsi)
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Disk /dev/sda: 120GB
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Sector size (logical/physical): 512B/512B
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Partition Table: mac
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Disk Flags:
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Number  Start   End     Size    File system  Name       Flags
 1      512B    32.8kB  32.3kB               Apple
 2      32.8kB  852kB   819kB   hfs          bootstrap  boot
 3      852kB   538MB   537MB   ext4         Boot
 4      538MB   54.2GB  53.7GB  ext4         Gentoo

```

In this output, partition 2 has the bootstrap information so /dev/sda2 is the correct path to use.

Format this partition as [HFS](/wiki/HFS "HFS") using the hformat command which is part of the [sys-fs/hfsutils](https://packages.gentoo.org/packages/sys-fs/hfsutils) package:

`root #` `dd if=/dev/zero of=/dev/sda2 bs=512`

`root #` `hformat -l bootstrap /dev/sda2`

Create a directory to mount the bootstrap partition and then mount it:

`root #` `mkdir /boot/NWBB
`

`root #` `mount --types hfs /dev/sda2 /boot/NWBB`

### Setup GRUB

`root #` `grub-install --macppc-directory=/boot/NWBB /dev/sda2`

If it installs without errors, unmount the bootstrap:

`root #` `umount /boot/NWBB`

Next, bless the partition so it will boot:

`root #` `hmount /dev/sda2
`

`root #` `hattrib -t tbxi -c UNIX :System:Library:CoreServices:BootX
`

`root #` `hattrib -b :System:Library:CoreServices
`

`root #` `humount`

Finally, build the grub.cfg file:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

## OldWorld Macs

### BootX

**重要**

BootX can only be used on OldWorld Apple systems with MacOS classic 7 to 9. For machines under 32MB of RAM use MacOS classic 7.

BootX can be downloaded at [https://github.com/immolo/BootX/archive/refs/tags/1.2.2.tar.gz](https://github.com/immolo/BootX/archive/refs/tags/1.2.2.tar.gz)

Since BootX boots Linux from within MacOS, the kernel will need to be copied from the Linux Partition to the MacOS partition. First, mount the MacOS partition from outside of the chroot. Use mac-fdisk -l to find the MacOS partition number, sda6 is used as an example here. Once the partition is mounted, we'll copy the kernel to the system folder so BootX can find it.

`root #` `exit`

`cdimage ~#` `mkdir /mnt/mac
`

`cdimage ~#` `mount /dev/sda6 /mnt/mac -t hfs
`

`cdimage ~#` `cp /mnt/gentoo/usr/src/linux/vmlinux "/mnt/mac/System Folder/Linux Kernels/kernel-6.18.8-gentoo"`

Now that the kernel is copied over, we'll need to reboot to set up BootX.

`cdimage ~#` `cd /
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/pts,/shm,}
`

`cdimage ~#` `umount -l /mnt/gentoo{/proc,/sys,}
`

`cdimage ~#` `umount -l /mnt/mac
`

`cdimage ~#` `reboot`

Of course, don't forget to remove the bootable CD, otherwise the CD will be booted again instead of MacOS.

Once the machine has booted into MacOS, open the BootX control panel. When not using genkernel, select Options and uncheck Use specified RAM disk. If genkernel is used, ensure that the genkernel initrd is selected instead of the Installation CD initrd. If not using genkernel, there is now an option to specify the machine's Linux root disk and partition. Fill these in with the appropriate values. Depending upon the kernel configuration, additional boot arguments may need to be applied.

BootX can be configured to start Linux upon boot. If this is done, then the machine will boot into MacOS first and, during startup, BootX will load and start Linux. See the BootX home page for more information.

**重要**

Make sure to include support for the HFS and HFS+ filesystems in the kernel, otherwise upgrades or changes to the kernel on the MacOS partition will not be possible.

## Pegasos

**メモ**

Pegasos also has Grub support but this is currently undocumented in Gentoo. Please add this to the main wiki and notify on this page's discussion once ready to migrated here.

### BootCreator

**重要**

BootCreator will build a nice SmartFirmware bootmenu written in Forth for the Pegasos.

First make sure to have bootcreator installed on the system:

`root #` `emerge --ask sys-boot/bootcreator`

Now copy the file /etc/bootmenu.example into /etc/bootmenu/ and edit it to suit personal needs:

`root #` `cp /etc/bootmenu.example /etc/bootmenu
`

`root #` `nano -w /etc/bootmenu`

Below is a complete /etc/bootmenu config file. vmlinux and initrd should be replaced by the kernel and initrd image names.

ファイル **`/etc/bootmenu`** **Example bootcreator configuration file**

```
#
# Example description file for bootcreator 1.1
#

[VERSION]
1

[TITLE]
Boot Menu

[SETTINGS]
AbortOnKey = false
Timeout    = 9
Default    = 1

[SECTION]
Local HD -> Morphos      (Normal)
ide:0 boot2.img ramdebug edebugflags="logkprintf"

[SECTION]
Local HD -> Linux (Normal)
ide:0 kernel-6.18.8-gentoo video=radeonfb:1024x768@70 root=/dev/sda3

[SECTION]
Local HD -> Genkernel (Normal)
ide:0 kernel-genkernel-ppc-6.18.8-gentoo root=/dev/ram0
root=/dev/sda3 initrd=initramfs-genkernel-ppc-6.18.8-gentoo

```

Finally the bootmenu must be transferred into Forth and copied to the boot partition, so that the SmartFirmware can read it. Therefore it is necessar to call bootcreator:

`root #` `bootcreator /etc/bootmenu /boot/menu`

**メモ**

Be sure to have a look into the SmartFirmware's settings when rebooting, that menu is the file that will be loaded by default.

## システムのリブート

chroot環境を出て、全てのパーティションをアンマウントします。次に、最終かつ真のテストを実行するためのマジカルコマンドrebootを入力しましょう。

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

live イメージを取り出すのを忘れないでください。そうしないと新しくインストールされた Gentoo ではなく、live イメージが再度ブート対象になってしまうかもしれません！

リブートして新しい Gentoo 環境に入ることができたら、最終章の [インストールの締めくくり](/wiki/Handbook:PPC/Installation/Finalizing/ja "Handbook:PPC/Installation/Finalizing/ja") に進むのがよいでしょう。

[← ツールのインストール](/wiki/Handbook:PPC/Installation/Tools/ja "Previous part") [Home](/wiki/Handbook:PPC/ja "Handbook:PPC/ja") [インストールの締めくくり →](/wiki/Handbook:PPC/Installation/Finalizing/ja "Next part")