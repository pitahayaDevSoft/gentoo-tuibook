# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbuch:SPARC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Bootloader "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Bootloader/es "Manual de Gentoo: SPARC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Bootloader/fr "Handbook:SPARC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Bootloader/pl "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Bootloader/pt-br "Handbook:SPARC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Bootloader/ru "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Bootloader/ta "கையேடு:SPARC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Bootloader/zh-cn "手册：SPARC/安装/配置系统引导程序Bootloader (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:SPARC/Installation/Bootloader/ko "Handbook:SPARC/Installation/Bootloader/ko (100% translated)")

[SPARC ハンドブック](/wiki/Handbook:SPARC "Handbook:SPARC")[インストール](/wiki/Handbook:SPARC/Full/Installation "Handbook:SPARC/Full/Installation")[インストールについて](/wiki/Handbook:SPARC/Installation/About/ja "Handbook:SPARC/Installation/About/ja")[メディアの選択](/wiki/Handbook:SPARC/Installation/Media/ja "Handbook:SPARC/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:SPARC/Installation/Networking/ja "Handbook:SPARC/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:SPARC/Installation/Disks/ja "Handbook:SPARC/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:SPARC/Installation/Stage/ja "Handbook:SPARC/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:SPARC/Installation/Base/ja "Handbook:SPARC/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:SPARC/Installation/Kernel/ja "Handbook:SPARC/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:SPARC/Installation/System/ja "Handbook:SPARC/Installation/System/ja")[ツールのインストール](/wiki/Handbook:SPARC/Installation/Tools/ja "Handbook:SPARC/Installation/Tools/ja")ブートローダの設定[締めくくり](/wiki/Handbook:SPARC/Installation/Finalizing/ja "Handbook:SPARC/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:SPARC/Full/Working "Handbook:SPARC/Full/Working")[Portage について](/wiki/Handbook:SPARC/Working/Portage/ja "Handbook:SPARC/Working/Portage/ja")[USE フラグ](/wiki/Handbook:SPARC/Working/USE/ja "Handbook:SPARC/Working/USE/ja")[Portage の機能](/wiki/Handbook:SPARC/Working/Features/ja "Handbook:SPARC/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:SPARC/Working/Initscripts/ja "Handbook:SPARC/Working/Initscripts/ja")[環境変数](/wiki/Handbook:SPARC/Working/EnvVar/ja "Handbook:SPARC/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:SPARC/Full/Portage "Handbook:SPARC/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:SPARC/Portage/Files/ja "Handbook:SPARC/Portage/Files/ja")[変数](/wiki/Handbook:SPARC/Portage/Variables/ja "Handbook:SPARC/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:SPARC/Portage/Branches/ja "Handbook:SPARC/Portage/Branches/ja")[追加ツール](/wiki/Handbook:SPARC/Portage/Tools/ja "Handbook:SPARC/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:SPARC/Portage/CustomTree/ja "Handbook:SPARC/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:SPARC/Portage/Advanced/ja "Handbook:SPARC/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:SPARC/Full/Networking "Handbook:SPARC/Full/Networking")[はじめに](/wiki/Handbook:SPARC/Networking/Introduction/ja "Handbook:SPARC/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:SPARC/Networking/Advanced/ja "Handbook:SPARC/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:SPARC/Networking/Modular/ja "Handbook:SPARC/Networking/Modular/ja")[無線](/wiki/Handbook:SPARC/Networking/Wireless/ja "Handbook:SPARC/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:SPARC/Networking/Extending/ja "Handbook:SPARC/Networking/Extending/ja")[動的な管理](/wiki/Handbook:SPARC/Networking/Dynamic/ja "Handbook:SPARC/Networking/Dynamic/ja")

## Contents

- [1GRUB](#GRUB)
  - [1.1Emerge](#Emerge)
  - [1.2インストール](#.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB)
    - [1.2.1GPT](#GPT)
    - [1.2.2Sun partition table](#Sun_partition_table)
  - [1.3Configure](#Configure)
- [2SILO, the SPARC bootloader](#SILO.2C_the_SPARC_bootloader)
- [3システムのリブート](#.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.81.AE.E3.83.AA.E3.83.96.E3.83.BC.E3.83.88)

## GRUB

When [selecting a 64-bit profile](/wiki/Handbook:SPARC/Installation/Base/ja#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/ja") during installation, then [GRUB](/wiki/GRUB/ja "GRUB/ja") is the only supported bootloader.

### Emerge

GRUB should be correctly configured for the platform automatically based on the profile. To make it explicit, however, specify it using:

`root #` `echo 'GRUB_PLATFORMS="ieee1275"' >> /etc/portage/make.conf`

`root #` `emerge --ask --verbose sys-boot/grub`

The GRUB software has now been merged to the system, but not yet installed as a bootloader.

### インストール

#### GPT

If the [disk is partitioned](/wiki/Handbook:SPARC/Installation/Disks/ja#Using_fdisk_to_partition_the_disk "Handbook:SPARC/Installation/Disks/ja") using GPT (the preferred method), then install GRUB to the BIOS boot partition. Presuming the first disk (the one where the system boots from) is /dev/sda, the following commands will do:

`root #` `grub-install --target=sparc64-ieee1275 --recheck /dev/sda`

**ヒント**

{{{1}}}

#### Sun partition table

If the disk is partitioned using a Sun partition table instead, GRUB must be installed using blocklists. In this mode, instead of providing the physical disk as an argument, provide the path to the partition on which /boot/grub is mounted.

`root #` `grub-install --target=sparc64-ieee1275 --recheck --force --skip-fs-probe /dev/sda1`

### Configure

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

## SILO, the SPARC bootloader

When [selecting a 32-bit profile](/wiki/Handbook:SPARC/Installation/Base/ja#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/ja") during installation, then [SILO](https://git.kernel.org/?p=linux/kernel/git/davem/silo.git;a=summary) (Sparc Improved boot LOader) is the only supported bootloader.

`root #` `emerge --ask sys-boot/silo`

Next create /etc/silo.conf:

`root #` `nano -w /etc/silo.conf`

Below an example silo.conf file is shown. It uses the partitioning scheme we use throughout this book, kernel-6.19.3-gentoo as kernel image and initramfs-genkernel-sparc64-6.19.3-gentoo as initramfs.

ファイル **`/etc/silo.conf`** **Example configuration file**

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

**メモ**

Run silo (with parameters if necessary) again each time after updating or installing the [sys-boot/silo](https://packages.gentoo.org/packages/sys-boot/silo) package.

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

リブートして新しい Gentoo 環境に入ることができたら、最終章の [インストールの締めくくり](/wiki/Handbook:SPARC/Installation/Finalizing/ja "Handbook:SPARC/Installation/Finalizing/ja") に進むのがよいでしょう。

[← ツールのインストール](/wiki/Handbook:SPARC/Installation/Tools/ja "Previous part") [Home](/wiki/Handbook:SPARC/ja "Handbook:SPARC/ja") [インストールの締めくくり →](/wiki/Handbook:SPARC/Installation/Finalizing/ja "Next part")