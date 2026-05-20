# Disks

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Disks/de "Handbuch:PPC64/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Disks/es "Manual de Gentoo: PPC64/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Disks/fr "Handbook:PPC64/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Disks/it "Handbook:PPC64/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Disks/hu "Handbook:PPC64/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Disks/pl "Handbook:PPC64/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Disks/pt-br "Handbook:PPC64/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Disks/ru "Handbook:PPC64/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Disks/ta "கையேடு:PPC64/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Disks/zh-cn "手册：PPC64/安装/准备磁盘 (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:PPC64/Installation/Disks/ko "Handbook:PPC64/Installation/Disks/ko (100% translated)")

[PPC64 ハンドブック](/wiki/Handbook:PPC64 "Handbook:PPC64")[インストール](/wiki/Handbook:PPC64/Full/Installation "Handbook:PPC64/Full/Installation")[インストールについて](/wiki/Handbook:PPC64/Installation/About/ja "Handbook:PPC64/Installation/About/ja")[メディアの選択](/wiki/Handbook:PPC64/Installation/Media/ja "Handbook:PPC64/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:PPC64/Installation/Networking/ja "Handbook:PPC64/Installation/Networking/ja")ディスクの準備[stage ファイル](/wiki/Handbook:PPC64/Installation/Stage/ja "Handbook:PPC64/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:PPC64/Installation/Base/ja "Handbook:PPC64/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:PPC64/Installation/Kernel/ja "Handbook:PPC64/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:PPC64/Installation/System/ja "Handbook:PPC64/Installation/System/ja")[ツールのインストール](/wiki/Handbook:PPC64/Installation/Tools/ja "Handbook:PPC64/Installation/Tools/ja")[ブートローダの設定](/wiki/Handbook:PPC64/Installation/Bootloader/ja "Handbook:PPC64/Installation/Bootloader/ja")[締めくくり](/wiki/Handbook:PPC64/Installation/Finalizing/ja "Handbook:PPC64/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:PPC64/Full/Working "Handbook:PPC64/Full/Working")[Portage について](/wiki/Handbook:PPC64/Working/Portage/ja "Handbook:PPC64/Working/Portage/ja")[USE フラグ](/wiki/Handbook:PPC64/Working/USE/ja "Handbook:PPC64/Working/USE/ja")[Portage の機能](/wiki/Handbook:PPC64/Working/Features/ja "Handbook:PPC64/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:PPC64/Working/Initscripts/ja "Handbook:PPC64/Working/Initscripts/ja")[環境変数](/wiki/Handbook:PPC64/Working/EnvVar/ja "Handbook:PPC64/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:PPC64/Full/Portage "Handbook:PPC64/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:PPC64/Portage/Files/ja "Handbook:PPC64/Portage/Files/ja")[変数](/wiki/Handbook:PPC64/Portage/Variables/ja "Handbook:PPC64/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:PPC64/Portage/Branches/ja "Handbook:PPC64/Portage/Branches/ja")[追加ツール](/wiki/Handbook:PPC64/Portage/Tools/ja "Handbook:PPC64/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:PPC64/Portage/CustomTree/ja "Handbook:PPC64/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:PPC64/Portage/Advanced/ja "Handbook:PPC64/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:PPC64/Full/Networking "Handbook:PPC64/Full/Networking")[はじめに](/wiki/Handbook:PPC64/Networking/Introduction/ja "Handbook:PPC64/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:PPC64/Networking/Advanced/ja "Handbook:PPC64/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:PPC64/Networking/Modular/ja "Handbook:PPC64/Networking/Modular/ja")[無線](/wiki/Handbook:PPC64/Networking/Wireless/ja "Handbook:PPC64/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:PPC64/Networking/Extending/ja "Handbook:PPC64/Networking/Extending/ja")[動的な管理](/wiki/Handbook:PPC64/Networking/Dynamic/ja "Handbook:PPC64/Networking/Dynamic/ja")

## Contents

- [1ブロックデバイスの概要](#.E3.83.96.E3.83.AD.E3.83.83.E3.82.AF.E3.83.87.E3.83.90.E3.82.A4.E3.82.B9.E3.81.AE.E6.A6.82.E8.A6.81)
  - [1.1ブロックデバイス](#.E3.83.96.E3.83.AD.E3.83.83.E3.82.AF.E3.83.87.E3.83.90.E3.82.A4.E3.82.B9)
  - [1.2Partitions and slices](#Partitions_and_slices)
- [2パーティション構成の設計](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E6.A7.8B.E6.88.90.E3.81.AE.E8.A8.AD.E8.A8.88)
  - [2.1パーティション数とサイズ](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E6.95.B0.E3.81.A8.E3.82.B5.E3.82.A4.E3.82.BA)
  - [2.2スワップ領域について](#.E3.82.B9.E3.83.AF.E3.83.83.E3.83.97.E9.A0.98.E5.9F.9F.E3.81.AB.E3.81.A4.E3.81.84.E3.81.A6)
- [3Default: Using mac-fdisk](#Default:_Using_mac-fdisk)
- [4代替案: fdisk を用いる](#.E4.BB.A3.E6.9B.BF.E6.A1.88:_fdisk_.E3.82.92.E7.94.A8.E3.81.84.E3.82.8B)
  - [4.1Viewing current partition layout](#Viewing_current_partition_layout)
  - [4.2Removing all partitions](#Removing_all_partitions)
  - [4.3Creating the PPC PReP boot partition](#Creating_the_PPC_PReP_boot_partition)
  - [4.4スワップパーティションを作成する](#.E3.82.B9.E3.83.AF.E3.83.83.E3.83.97.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [4.5ルートパーティションを作成する](#.E3.83.AB.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [4.6パーティションのレイアウトを保存する](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E3.83.AC.E3.82.A4.E3.82.A2.E3.82.A6.E3.83.88.E3.82.92.E4.BF.9D.E5.AD.98.E3.81.99.E3.82.8B)
- [5ファイルシステムを作成する](#.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [5.1はじめに](#.E3.81.AF.E3.81.98.E3.82.81.E3.81.AB)
  - [5.2ファイルシステム](#.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
  - [5.3パーティションにファイルシステムを適用する](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AB.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.82.92.E9.81.A9.E7.94.A8.E3.81.99.E3.82.8B)
    - [5.3.1レガシー BIOS ブートパーティションのファイルシステム](#.E3.83.AC.E3.82.AC.E3.82.B7.E3.83.BC_BIOS_.E3.83.96.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
    - [5.3.2小さな ext4 パーティション](#.E5.B0.8F.E3.81.95.E3.81.AA_ext4_.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3)
  - [5.4スワップパーティションを有効にする](#.E3.82.B9.E3.83.AF.E3.83.83.E3.83.97.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E6.9C.89.E5.8A.B9.E3.81.AB.E3.81.99.E3.82.8B)
- [6ルートパーティションのマウント](#.E3.83.AB.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E3.83.9E.E3.82.A6.E3.83.B3.E3.83.88)

## ブロックデバイスの概要

### ブロックデバイス

Gentoo Linuxの、そしてLinux一般の、ブロックデバイス、パーティション、Linuxファイルシステムを含めた、ディスクやファイルシステム中心の考え方について詳しく見てみましょう。ディスクの入出力とファイルシステムについて理解することで、インストールのためのパーティションとファイルシステムを構築できるようになります。

まずはブロックデバイスについて見ていきます。SCSIドライブやシリアルATAドライブは両方とも/dev/sdaや/dev/sdb、/dev/sdcなどのようなデバイスハンドルとしてラベル付されます。更にモダンなマシンでは、PCI ExpressベースのNVMeソリッドステートディスクは、/dev/nvme0n1、/dev/nvme0n2などのようなデバイスハンドルを持ちます。

下の表は、各種のブロックデバイスがシステム上のどこにあるかを判断するのに役立つでしょう:

デバイスの種類デフォルトのデバイスハンドル編集者メモと、考慮すべき点
IDE、SATA、SAS、SCSI、または USB フラッシュメモリ/dev/sda2007 年頃から現在までに製造されたハードウェアで見られます。このデバイスハンドルはおそらく Linux 上でもっともよく使用されているものでしょう。この種のデバイスは [SATA バス](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA")、 [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI")、 [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB") バスを介してブロックストレージとして接続されます。例えば、最初の SATA デバイス上の最初のパーティションは /dev/sda1 という名前になります。
NVM Express (NVMe)/dev/nvme0n1ソリッドステートテクノロジとして最新の [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") ドライブは PCI Express バスに接続され、一般市場でもっとも高速な転送速度を持っています。2014 年頃以降のシステムは NVMe ハードウェアのサポートを備えているかもしれません。最初の NVMe デバイスの最初のパーティションは /dev/nvme0n1p1 という名前になります。
MMC、eMMC、および SD カード/dev/mmcblk0[embedded MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard") デバイス、SD カード、そして他の種類のメモリーカードはデータ用のストレージとして有用です。つまり、多くのシステムはこれらの種類のデバイスからのブートを許可していないかもしれません。これらのデバイスに Linux をインストールして常用するのは **おすすめできません**。それらの典型的な設計意図である、ファイルの交換用に使うものと考えてください。この種のストレージは短期的なファイルバックアップまたはスナップショットとして使用すると便利かもしれません。
VirtIO/dev/vda[QEMU](/wiki/QEMU "QEMU") 仮想エミュレータ内でのみ見られる [仮想](/wiki/Virtualization "Virtualization") デバイス。 `virtio_blk` ドライバは仮想マシンに対してホストのドライブ空間へのアクセスを提供します。とはいえ、仮想マシン内で Gentoo を試すための優れた方法です。

上のブロックデバイスは、ディスクへの抽象的なインターフェースを表しています。ユーザープログラムはこれらのブロックデバイスを用いて、デバイスが SATA、SCSI、もしくは他のものであるかどうかを心配することなしにディスクと通信することができます。プログラムは容易にディスク上の記憶領域を、ランダムアクセスできる 4096 バイト (4K) ごとの連続領域としてアドレッシングできます。

### Partitions and slices

Although it is theoretically possible to use a full disk to house a Linux system, this is almost never done in practice. Instead, full disk block devices are split up in smaller, more manageable block devices. On most systems, these are called partitions. Other architectures use a similar technique, called _slices_.

## パーティション構成の設計

### パーティション数とサイズ

ディスクのパーティションレイアウトの設計は、システムに対する要求と、デバイスに適用されるファイルシステムに大きく依存します。多数のユーザがいる場合、セキュリティを向上し、バックアップの作成とその他のメンテナンスを容易にするために、/home を分離されたパーティションに配置することが推奨されます。もし メールサーバとして動作する場合は、/var を分離されたパーティションとし、すべてのメールを /var ディレクトリに保存すべきでしょう。ゲームサーバでは、ほとんどのゲームサーバソフトウェアは /opt にインストールされるので、/opt を分離されたパーティションとすることができます。これらが推奨される理由は最初の /home ディレクトリと同様で、セキュリティ、バックアップ、そしてメンテナンスです。

Gentoo では多くの場合、/usr と /var は相対的に大きい容量を確保すべきです。/usr にはシステム上で利用可能なアプリケーションの大部分と、Linux カーネルソース (/usr/src 配下) が配置されます。デフォルトでは、/var には Gentoo ebuild リポジトリが (/var/db/repos/gentoo 配下に) 配置され、ファイルシステム依存ではあるものの通常 650 MiB ほどのディスク容量を消費します。この推定容量には /var/cache/distfiles と /var/cache/binpkgs ディレクトリは _含まれていません_。これらはそれぞれ、ソースファイルとバイナリパッケージ (使用している場合) を格納するディレクトリで、システムに追加すればするほど大きくなっていきます。

適切なパーティションの数とサイズは、システムを取り巻く環境と、トレードオフを考慮することで大きく変わります。パーティションやボリュームを分離することには下記の利点があります:

- それぞれのパーティションまたはボリュームに対して、最も性能が高いファイルシステムを選択できます
- ゾンビプロセスがパーティションまたはボリュームに継続的に書き込みをした場合でも、システム全体の空き領域を使い切ることはありません
- 必要ならば、複数のチェックを並行して実行することで、ファイルシステムチェックの時間を短縮できます (複数のパーティションよりも複数のディスクの方が効果を実感できます)
- リードのみ、 `nosuid`(setuidビット無効)、 `noexec`(実行ビット無効)等のマウントオプションによって、セキュリティが向上します

しかし、複数パーティションにはデメリットもあります:

- もし適切に設定されていないと、あるパーティションが空き領域をたくさん持ち、別のパーティションにはまったく空き領域がなくなるといったことが起こり得ます。
- /usr/ を独立したパーティションにすると、他のブートスクリプトが動作する前にパーティションをマウントするために、initramfs を使ってブートする必要があるかもしれません。initramfs の生成と保守はこのハンドブックのスコープの範囲外ですので、 **慣れていない方が /usr を独立したパーティションとすることは推奨しません。**
- SCSI や SATA では仕様上の制約により、GPT ラベルを使用しない限りは 15 個までしかパーティションを作れません。

**メモ**

サービスおよび init システムとして systemd を使うつもりのインストールでは、/usr ディレクトリはルートファイルシステムの一部とするか、または initramfs によりマウントされるようにして、ブート時に利用できるようにしなくてはなりません。

### スワップ領域について

スワップ領域サイズの推奨値
RAM サイズサスペンド対応時ハイバネーション対応時
2 GB 以下4GB4GB
2 から 8 GBRAM 容量2 \* RAM
8 から 64 GB最小 8 GB、最大 16 GB1.5 \* RAM
64 GB 以上最小 8 GBハイバネーションは推奨されません! 非常に大きな容量のメモリを持つシステムでは、ハイバネーションは推奨 _されません_。可能ではありますが、正しくハイバネーションするためには、メモリの内容全体をディスクに書き込まなくてはなりません。数十 GB (またはそれ以上!) のデータをディスクに書き出すのには、回転式ディスクを使用する場合は特に、非常に長い時間がかかることがあります。この状況ではサスペンドするのが最善です。

スワップ領域のサイズについて完璧な値というものはありません。スワップ領域の目的は、メインメモリ（RAM）が逼迫した際、カーネルにディスク領域を提供するためにあります。スワップ領域があれば、カーネルは最近最も使われていないメモリページをディスクに書き出し（スワップもしくはページアウト）、現在のタスクのために RAM 上に置かれたメモリを開放します。もちろん、もしディスクにスワップされたページが急に必要になった場合は、これらのページはメモリに戻す（ページイン）必要があります。これには、RAM から読み込むより相当長い時間がかかります（メインメモリと比較してディスクはとても遅いためです）。

システムがメモリを大量に消費するアプリケーションを実行しないとき、またシステムが多くの RAM を持っているときは、それほど大きいスワップ領域は必要ではありません。しかし、ハイバネーションの際に、スワップ領域は _メモリの内容すべて_ を保存するために使われる（サーバシステムよりも、デスクトップやラップトップシステムでよくあることです）ことに留意してください。システムにハイバネーションのサポートが必要な場合は、メモリの全体量以上のサイズのスワップ領域が必要です。

RAM 容量が 4GB より少ない場合の一般的なルールとして、スワップ領域のサイズは内部メモリ (RAM) の 2 倍であることが推奨されます。複数のハードディスクを備えるシステムでは、並列して読み込み/書き込み操作が行えるように、それぞれのディスクに 1 つずつスワップパーティションを作成するのが賢い方法です。スワップ空間内のデータにアクセスしなくてはならないときに、ディスクがより高速にスワップできるほど、システムもより高速に動作するでしょう。回転式ディスクとソリッドステートディスクを比較すると、ソリッドステートハードウェア上にスワップを置いたほうが高いパフォーマンスが発揮できます。

スワップ _パーティション_ の代わりに、スワップ _ファイル_ を使用することができることも特筆に値します。これは主にディスク容量が非常に限られたシステムで役に立つものです。

## Default: Using mac-fdisk

**重要**

These instructions are for the Apple G5 system.

Partition
Description
/dev/sda1Apple partition map, automatically created when the disk is formatted with a "mac" partition table.
/dev/sda2New World boot block
/dev/sda3Swap partition
/dev/sda4Root partition

Start mac-fdisk:

`root #` `mac-fdisk /dev/sda`

First delete the partitions that have been cleared previously to make room for Linux partitions. Use the `d` key in mac-fdisk to delete those partition(s). It will ask for the partition number to delete.

Second, create an Apple\_Bootstrap partition by pressing the `b` key. It will ask for a block from which to start. Enter the number of the first free partition, followed by entering a `p`. For instance this is _2p_.

**メモ**

This partition is not a "boot" partition. It is not used by Linux at all; there is no need to place any filesystem on it and it should never be mounted. PPC users do _not_ need an extra partition for /boot.

Now create a swap partition by pressing the `c` key. Again mac-fdisk will ask what block to start from. As we used 2 before to create the Apple\_Bootstrap partition, enter _3p_. When asked for the size, enter **512M** (or whatever size needed). When asked for a name, enter _swap_ (mandatory).

To create the root partition, enter `c`, followed by _4p_ to select from what block the root partition should start. When asked for the size, enter _4p_ again. mac-fdisk will interpret this as "Use all available space". When asked for the name, enter _root_ (mandatory).

To finish up, write the partition to the disk using `w` and `q` to quit mac-fdisk.

**メモ**

To make sure everything is okay, run mac-fdisk once more to verify all the partitions are present. If a partition is absent, or it is missing some of the changes that were made, then reinitialize the partitions by pressing `i` in mac-fdisk. Note that this will recreate the partition map and thus remove all the partitions.

## 代替案: fdisk を用いる

**重要**

The following instructions are for IBM pSeries, iSeries, and OpenPower systems.

**メモ**

When planning to use a RAID disk array for the Gentoo installation on POWER5-based hardware, first run iprconfig to format the disks to Advanced Function format and create the disk array. Emerge [sys-fs/iprutils](https://packages.gentoo.org/packages/sys-fs/iprutils) after the installation is complete.

If the system has an ipr-based SCSI adapter, start the ipr utilities now.

`root #` `/etc/init.d/iprinit start`

The following parts explain how to create the example partition layout described previously, namely:

Partition
Description
/dev/sda1PPC PReP Boot partition
/dev/sda2Swap partition
/dev/sda3Root partition

Change or modify the partition layout according to personal preference.

### Viewing current partition layout

[fdisk](/wiki/Fdisk "Fdisk") is a popular and powerful tool to split a disk into partitions. Fire up fdisk on the current disk (in our example, we use /dev/sda):

`root #` `fdisk /dev/sda`

```
Command (m for help)

```

If there is still an AIX partition layout on the system, then the following error message will be displayed:

`root #` `fdisk /dev/sda`

```
  There is a valid AIX label on this disk.
  Unfortunately Linux cannot handle these
  disks at the moment.  Nevertheless some
  advice:
  1. fdisk will destroy its contents on write.
  2. Be sure that this disk is NOT a still vital
     part of a volume group. (Otherwise you may
     erase the other disks as well, if unmirrored.)
  3. Before deleting this physical volume be sure
     to remove the disk logically from your AIX
     machine.  (Otherwise you become an AIXpert).

```

Don't worry, new empty DOS partition table can be created by pressing `o`.

**警告**

This will destroy any installed AIX version!

Type `p` to display the disk current partition configuration:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          12       53266+  83  Linux
/dev/sda2              13         233      981571+  82  Linux swap
/dev/sda3             234         674     1958701+  83  Linux
/dev/sda4             675        6761    27035410+   5  Extended
/dev/sda5             675        2874     9771268+  83  Linux
/dev/sda6            2875        2919      199836   83  Linux
/dev/sda7            2920        3008      395262   83  Linux
/dev/sda8            3009        6761    16668918   83  Linux

```

This particular disk is configured to house six Linux filesystems (each with a corresponding partition listed as "Linux") as well as a swap partition (listed as "Linux swap").

### Removing all partitions

First remove all existing partitions from the disk. Type `d` to delete a partition. For instance, to delete an existing /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

The partition has been scheduled for deletion. It will no longer show up when typing `p`, but it will not be erased until the changes have been saved. If a mistake was made and the session needs to be aborted, then type `q` immediately and hit `Enter` and none of the partitions will be deleted or modified.

Now, assuming that indeed all partitions need to be wiped out, repeatedly type `p` to print out a partition listing and then type `d` and the number of the partition to delete it. Eventually, the partition table will show no more partitions:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

Device Boot    Start       End    Blocks   Id  System

```

Now that the in-memory partition table is empty, let's create the partitions. We will use a default partitioning scheme as discussed previously. Of course, don't follow these instructions to the letter but adjust to personal preference.

### Creating the PPC PReP boot partition

First create a small PReP boot partition. Type `n` to create a new partition, then `p` to select a primary partition, followed by `1` to select the first primary partition. When prompted for the first cylinder, hit `Enter`. When prompted for the last cylinder, type +7M to create a partition 7 MB in size. After this, type `t` to set the partition type, `1` to select the partition just created and then type in _41_ to set the partition type to "PPC PReP Boot". Finally, mark the PReP partition as bootable.

**メモ**

The PReP partition has to be smaller than 8 MB!

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System

```

`Command (m for help):` `n`

```
Command action
      e   extended
      p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-6761, default 1):
Using default value 1
Last cylinder or +size or +sizeM or +sizeK (1-6761, default
6761): +8M

```

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 41
Changed system type of partition 1 to 41 (PPC PReP Boot)

```

`Command (m for help):` `a`

```
Partition number (1-4): 1
Command (m for help):

```

Now, when looking at the partition table again (through `p`), the following partition information should be shown:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1  *            1           3       13293   41  PPC PReP Boot

```

### スワップパーティションを作成する

Now create the swap partition. To do this, type `n` to create a new partition, then `p` to tell fdisk to create a primary partition. Then type `2` to create the second primary partition, /dev/sda2 in our case. When prompted for the first cylinder, hit `Enter`. When prompted for the last cylinder, type +512M to create a partition 512MB in size. After this, type `t` to set the partition type, `2` to select the partition just created and then type in _82_ to set the partition type to "Linux Swap". After completing these steps, typing `p` should display a partition table that looks similar to this:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1           3       13293   41  PPC PReP Boot
/dev/sda2               4         117      506331   82  Linux swap

```

### ルートパーティションを作成する

Finally, create the root partition. To do this, type `n` to create a new partition, then `p` to tell fdisk to create a primary partition. Then type `3` to create the third primary partition, /dev/sda3 in our case. When prompted for the first cylinder, hit `Enter`. When prompted for the last cylinder, hit enter to create a partition that takes up the rest of the remaining space on the disk. After completing these steps, typing `p` should display a partition table that looks similar to this:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1           3       13293   41  PPC PReP Boot
/dev/sda2               4         117      506331   82  Linux swap
/dev/sda3             118        6761    29509326   83  Linux

```

### パーティションのレイアウトを保存する

To save the partition layout and exit fdisk, type `w`.

`Command (m for help):` `w`

## ファイルシステムを作成する

**警告**

SSD または NVMe ドライブを使用する場合は、ファームウェアアップグレードがあるかどうか確認するのが賢明です。特に一部の Intel SSDs (600p および 6000p) は、XFS の I/O 使用量パターンによって誘発される [データ破損の可能性](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) のためのファームウェアアップグレードが必要です。この問題はファームウェアレベルのもので、XFS ファイルシステムの欠陥ではありません。smartctl ユーティリティはデバイスのモデルとファームウェアバージョンを確認するのに役立ちます。

### はじめに

パーティションが作成できたら、その上にファイルシステムを作成します。次の節ではLinuxがサポートする各種ファイルシステムを紹介します。どのファイルシステムを使うかをすでに決めているなら、 [パーティションにファイルシステムを適用する](/wiki/Handbook:PPC64/Installation/Disks/ja#Applying_a_filesystem_to_a_partition "Handbook:PPC64/Installation/Disks/ja") へ進みましょう。そうでなければ、次の節を読んで利用可能なファイルシステムについて知るのがよいでしょう。

### ファイルシステム

Linux は多くのファイルシステムをサポートしていますが、それらの多くは特定の目的をもって配備するのが賢明なものです。特定のファイルシステムのみが ppc64 アーキテクチャ上で安定して動作するとされています - 重要なパーティションに実験的なファイルシステムを選択するときは、事前にファイルシステムのサポート状況を十分に知っておくことを推奨します。 **XFS はすべてのプラットフォームで、すべての目的で推奨されるファイルシステムです。** 以下は、網羅的ではないリストです:

[XFS](/wiki/XFS/ja "XFS/ja")メタデータジャーナリングのあるファイルシステムで、堅牢な機能セットを持ち、スケーラビリティに最適化されています。新しい機能を取り入れながら継続的にアップグレードされ続けています。唯一の欠点は、現在対応中ではあるものの、 XFS パーティションはまだ縮小できないという点です。XFS の特筆すべき点として reflink とコピーオンライト (CoW) に対応しており、これは Gentoo システム上ではユーザが実施するコンパイル量の多さから特に有用です。XFS は全目的、全プラットフォームで利用できる、おすすめの現代的なファイルシステムです。パーティションは少なくとも 300MB ある必要があります。[ext4](/wiki/Ext4/ja "Ext4/ja")ext4 は reflink などの現代的な機能を欠いてはいるものの、信頼性があり、全目的、全プラットフォームで使用できるファイルシステムです。[VFAT](/wiki/VFAT "VFAT")別名 FAT32。Linux でサポートされていますが、標準的な UNIX パーミッションの設定をサポートしていません。ほとんど、他の OS (Microsoft Windows または Apple macOS) との相互運用性/交換のために使われていますが、いくつかのシステムブートローダーファームウェア (たとえば UEFI) でも必要になります。UEFI システムを使用している場合は、システムをブートするためには VFAT でフォーマットされた [EFI システムパーティション](/wiki/EFI_System_Partition/ja "EFI System Partition/ja") が必要になるでしょう。[btrfs](/wiki/Btrfs/ja "Btrfs/ja")次世代のファイルシステムです。スナップショット、チェックサムによる自己修復、透過的圧縮、サブボリューム、RAID の統合などの先進的な機能を提供します。RAID 5/6 とクオータグループは、btrfs のすべてのバージョンで安全ではありません。[F2FS](/wiki/F2FS/ja "F2FS/ja")Flash-Friendly File System はもともと、Samsung によって NAND フラッシュメモリで利用するために作られました。Gentoo を microSD カードや USB スティックや他のフラッシュベースの記憶装置にインストールする際にはすばらしい選択でしょう。[NTFS](/wiki/NTFS/ja "NTFS/ja")この "New Technology" ファイルシステムは、Windows NT 3.1 以降の Microsoft Windows のフラッグシップファイルシステムです。VFAT と同様、BSD や Linux が正しく動作するために必要な UNIX パーミッション設定や拡張属性を保持しないため、ほとんどの場合ルートファイルシステムとして使うべきではありません。Microsoft Windows とデータ交換の相互運用のために _のみ_ 使うべきです ( _のみ_ の強調に注意してください)。[ZFS](/wiki/ZFS "ZFS")**重要:** ZFS プールは minimal インストール CD で作成することができます。さらなる情報については [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs") を参照してください。Matthew Ahrens と Jeff Bonwick によって作成された次世代のファイルシステム。いくつかのキーとなる理想に基づいて設計されました: ストレージの管理はシンプルであるべき、冗長性はファイルシステムによって処理されるべき、ファイルシステムは修復のためにオフラインにされるべきではない、コードをリリースする前に最悪のシナリオを自動的にシミュレーションすることは重要である、そしてデータの整合性は最優先である。

ファイルシステムについてのより広範な情報は、コミュニティによって維持されている [ファイルシステムの記事](/wiki/Filesystem/ja "Filesystem/ja") で見つけることができます。

### パーティションにファイルシステムを適用する

**メモ**

再起動する前に、選択したファイルシステムに関連するユーザ空間ユーティリティを emerge しておいてください。インストールプロセスの終わりのあたりでまた、そうするように再通知します。

パーティションまたはボリュームの上にファイルシステムを作成するには、ファイルシステムごとに異なるユーザースペースのユーティリティが利用可能です。下表でファイルシステムの名前をクリックすると、それぞれに追加の情報が得られます：

ファイルシステム
作成コマンド
live 環境内にある？
パッケージ
[XFS](/wiki/XFS/ja "XFS/ja")mkfs.xfs はい
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/ja "Ext4/ja")mkfs.ext4 はい
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
mkfs.vfat はい
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs/ja "Btrfs/ja")mkfs.btrfs はい
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS/ja "F2FS/ja")mkfs.f2fs はい
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS/ja "NTFS/ja")mkfs.ntfs はい
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... はい
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

**重要**

ハンドブックではインストールプロセスの一部として新しいパーティションを使用することを推奨しますが、重要なこととして、mkfs コマンドを実行すると、パーティションに含まれるすべてのデータは消去されることに注意してください。必要な場合は、新しいファイルシステムを作成する前に、中に存在する任意のデータが適切にバックアップされていることを確認してください。

例えば、パーティション構造例の通りに、ルートパーティション (/dev/sda3) を xfs として設定するには、次のコマンドが使えます:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf /dev/sda3`

#### レガシー BIOS ブートパーティションのファイルシステム

MBR/DOS ディスクラベルを持ち、レガシー BIOS を利用してブートするシステムは、ブートローダがサポートする任意のファイルシステムフォーマットを使用することができます。

例えば、XFS でフォーマットするには:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf /dev/sda1`

#### 小さな ext4 パーティション

ext4 ファイルシステムを (8 GiB 未満の) 小さいパーティションに使用する場合は、十分な inode 数を確保できるように適切なオプションを指定してファイルシステムを作成するべきです。これは `-T small` オプションを使用して指定することができます:

`root #` `mkfs.ext4 -T small /dev/<device>`

こうすることで「inode あたりのバイト数」が 16kB から 4kB に減るので、ファイルシステムに 4 倍の inode 数を確保できます。

### スワップパーティションを有効にする

mkswapはスワップパーティションを初期化するために使われるコマンドです：

`root #` `mkswap /dev/sda2`

**メモ**

以前に開始したインストール手順を完了しなかった場合は、ハンドブックのこの時点からインストールを再開することができます。このリンクを permalink として使用してください: [インストールの再開はここから](/wiki/Handbook:PPC64/Installation/Disks/ja#Resumed_installations_start_here "Handbook:PPC64/Installation/Disks/ja")。

スワップパーティションを有効化するには、swaponを使います：

`root #` `swapon /dev/sda2`

この「有効化」の手順は、このスワップパーティションが live 環境内に新しく作成されたという理由から必要になっているものです。システムのリブート後は、スワップパーティションが fstab またはその他のマウント機構で適切に定義されている限り、スワップ領域は自動的に有効化されるでしょう。

## ルートパーティションのマウント

一部の live 環境は、Gentoo のルートパーティションのために提案されているマウントポイント (/mnt/gentoo) や、パーティショニングの節で作成された追加のパーティションのためのマウントポイントを持たないかもしれません:

`root #` `mkdir --parents /mnt/gentoo`

以前の手順で作成した追加の (カスタムの) パーティションがあれば、mkdir コマンドを使用して、追加で必要なマウントポイントの作成を続けてください。

マウントポイントが作成できたら、mount コマンドで、パーティションにアクセスできるようにしましょう。

ルートパーティションをマウントしてください:

`root #` `mount /dev/sda3 /mnt/gentoo`

必要に応じて、mount コマンドを使用して、追加の (カスタムの) パーティションのマウントを続けてください。

**メモ**

もし/tmp/を別のパーティションに置く必要があるなら、マウントしたあと権限の変更を忘れずに行ってください:

`root #` `chmod 1777 /mnt/gentoo/tmp`

/var/tmpについても同様です。

このあと解説の中で、proc ファイルシステム (仮想的なカーネルとのインターフェース) が、他のカーネル擬似ファイルシステムと同様にマウントされますが、まず最初は、 [Gentoo stage ファイル](/wiki/Handbook:PPC64/Installation/Stage/ja "Handbook:PPC64/Installation/Stage/ja") を展開する必要があります。

[← ネットワーク設定](/wiki/Handbook:PPC64/Installation/Networking/ja "Previous part") [Home](/wiki/Handbook:PPC64/ja "Handbook:PPC64/ja") [Gentoo インストールファイルをインストール →](/wiki/Handbook:PPC64/Installation/Stage/ja "Next part")