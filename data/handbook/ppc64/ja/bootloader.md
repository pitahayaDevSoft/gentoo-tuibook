# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Bootloader/de "Handbuch:PPC64/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Bootloader "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Bootloader/es "Manual de Gentoo: PPC64/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Bootloader/it "Handbook:PPC64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Handbook:PPC64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Bootloader/pt-br "Handbook:PPC64/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Bootloader/ru "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Bootloader/ta "கையேடு:PPC64/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Bootloader/zh-cn "手册：PPC64/安装/配置系统引导程序Bootloader (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko (100% translated)")

[PPC64 ハンドブック](/wiki/Handbook:PPC64 "Handbook:PPC64")[インストール](/wiki/Handbook:PPC64/Full/Installation "Handbook:PPC64/Full/Installation")[インストールについて](/wiki/Handbook:PPC64/Installation/About/ja "Handbook:PPC64/Installation/About/ja")[メディアの選択](/wiki/Handbook:PPC64/Installation/Media/ja "Handbook:PPC64/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:PPC64/Installation/Networking/ja "Handbook:PPC64/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:PPC64/Installation/Disks/ja "Handbook:PPC64/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:PPC64/Installation/Stage/ja "Handbook:PPC64/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:PPC64/Installation/Base/ja "Handbook:PPC64/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:PPC64/Installation/Kernel/ja "Handbook:PPC64/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:PPC64/Installation/System/ja "Handbook:PPC64/Installation/System/ja")[ツールのインストール](/wiki/Handbook:PPC64/Installation/Tools/ja "Handbook:PPC64/Installation/Tools/ja")ブートローダの設定[締めくくり](/wiki/Handbook:PPC64/Installation/Finalizing/ja "Handbook:PPC64/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:PPC64/Full/Working "Handbook:PPC64/Full/Working")[Portage について](/wiki/Handbook:PPC64/Working/Portage/ja "Handbook:PPC64/Working/Portage/ja")[USE フラグ](/wiki/Handbook:PPC64/Working/USE/ja "Handbook:PPC64/Working/USE/ja")[Portage の機能](/wiki/Handbook:PPC64/Working/Features/ja "Handbook:PPC64/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:PPC64/Working/Initscripts/ja "Handbook:PPC64/Working/Initscripts/ja")[環境変数](/wiki/Handbook:PPC64/Working/EnvVar/ja "Handbook:PPC64/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:PPC64/Full/Portage "Handbook:PPC64/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:PPC64/Portage/Files/ja "Handbook:PPC64/Portage/Files/ja")[変数](/wiki/Handbook:PPC64/Portage/Variables/ja "Handbook:PPC64/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:PPC64/Portage/Branches/ja "Handbook:PPC64/Portage/Branches/ja")[追加ツール](/wiki/Handbook:PPC64/Portage/Tools/ja "Handbook:PPC64/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:PPC64/Portage/CustomTree/ja "Handbook:PPC64/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:PPC64/Portage/Advanced/ja "Handbook:PPC64/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:PPC64/Full/Networking "Handbook:PPC64/Full/Networking")[はじめに](/wiki/Handbook:PPC64/Networking/Introduction/ja "Handbook:PPC64/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:PPC64/Networking/Advanced/ja "Handbook:PPC64/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:PPC64/Networking/Modular/ja "Handbook:PPC64/Networking/Modular/ja")[無線](/wiki/Handbook:PPC64/Networking/Wireless/ja "Handbook:PPC64/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:PPC64/Networking/Extending/ja "Handbook:PPC64/Networking/Extending/ja")[動的な管理](/wiki/Handbook:PPC64/Networking/Dynamic/ja "Handbook:PPC64/Networking/Dynamic/ja")

With the kernel configured and compiled and the necessary system configuration files filled in correctly, it is time to install a program that will fire up the kernel when the system boots. Such a program is called a boot loader.

**メモ**

Currently using Petitboot on Talos systems is undocumented in Gentoo. Please add the steps to [TalosII#Bootloader](/wiki/TalosII#Bootloader.2Fja "TalosII") and notify on this Discussion page when ready to merge into the Handbook.

## Contents

- [1Using GRUB](#Using_GRUB)
  - [1.1インストール](#.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB)
  - [1.2Mac hardware (G5)](#Mac_hardware_.28G5.29)
    - [1.2.1Setup bootstrap partition](#Setup_bootstrap_partition)
    - [1.2.2Setup GRUB](#Setup_GRUB)
  - [1.3IBM hardware](#IBM_hardware)
    - [1.3.1Setup GRUB](#Setup_GRUB_2)
    - [1.3.2Grub config](#Grub_config)
- [2システムのリブート](#.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.81.AE.E3.83.AA.E3.83.96.E3.83.BC.E3.83.88)

## Using GRUB

GRUB が PPC64 ベースの Linux マシン向けのブートローダです。

### インストール

`root #` `emerge --ask sys-boot/grub`

### Mac hardware (G5)

#### Setup bootstrap partition

First, prepare the bootstrap partition that was created created during the [preparing the disk](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks") step. Following the example, this partition should be /dev/sda2. Optionally, confirm this by using parted:

Replace /dev/sda with the correct device if required.

`root #` `parted /dev/sda print`

```
Model: ATA Patriot Burst El (scsi)

Disk /dev/sda: 120GB

Sector size (logical/physical): 512B/512B

Partition Table: mac

Disk Flags:

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

#### Setup GRUB

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

### IBM hardware

Setting up Grub on IBM hardware is as simple as:

#### Setup GRUB

`root #` `grub-install /dev/sda1`

**メモ**

/dev/sda1 is the PReP boot partition made in the partitioning stage

#### Grub config

Finally. build the grub.cfg file:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

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

リブートして新しい Gentoo 環境に入ることができたら、最終章の [インストールの締めくくり](/wiki/Handbook:PPC64/Installation/Finalizing/ja "Handbook:PPC64/Installation/Finalizing/ja") に進むのがよいでしょう。

[← ツールのインストール](/wiki/Handbook:PPC64/Installation/Tools/ja "Previous part") [Home](/wiki/Handbook:PPC64/ja "Handbook:PPC64/ja") [インストールの締めくくり →](/wiki/Handbook:PPC64/Installation/Finalizing/ja "Next part")