# Disks

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Disks/de "Handbook:AMD64/Installation/Disks/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Disks "Handbook:AMD64/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:AMD64/Installation/Disks/tr "Handbook:AMD64/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Disks/es "Handbook:AMD64/Installation/Disks/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Disks/fr "Handbook:AMD64/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Disks/it "Handbook:AMD64/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Disks/hu "Handbook:AMD64/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Disks/pl "Handbook:AMD64/Installation/Disks/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Disks/pt-br "Handbook:AMD64/Installation/Disks/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Disks/cs "Handbook:AMD64/Installation/Disks/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Disks/ru "Handbook:AMD64/Installation/Disks/ru (100% translated)")
- [العربية](/wiki/Handbook:AMD64/Installation/Disks/ar "Handbook:AMD64/Installation/Disks/ar (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Disks/ta "Handbook:AMD64/Installation/Disks/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Disks/zh-cn "Handbook:AMD64/Installation/Disks/zh-cn (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:AMD64/Installation/Disks/ko "Handbook:AMD64/Installation/Disks/ko (100% translated)")

[AMD64 ハンドブック](/wiki/Handbook:AMD64 "Handbook:AMD64")[インストール](/wiki/Handbook:AMD64/Full/Installation "Handbook:AMD64/Full/Installation")[インストールについて](/wiki/Handbook:AMD64/Installation/About/ja "Handbook:AMD64/Installation/About/ja")[メディアの選択](/wiki/Handbook:AMD64/Installation/Media/ja "Handbook:AMD64/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:AMD64/Installation/Networking/ja "Handbook:AMD64/Installation/Networking/ja")ディスクの準備[stage ファイル](/wiki/Handbook:AMD64/Installation/Stage/ja "Handbook:AMD64/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:AMD64/Installation/Base/ja "Handbook:AMD64/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:AMD64/Installation/Kernel/ja "Handbook:AMD64/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:AMD64/Installation/System/ja "Handbook:AMD64/Installation/System/ja")[ツールのインストール](/wiki/Handbook:AMD64/Installation/Tools/ja "Handbook:AMD64/Installation/Tools/ja")[ブートローダの設定](/wiki/Handbook:AMD64/Installation/Bootloader/ja "Handbook:AMD64/Installation/Bootloader/ja")[締めくくり](/wiki/Handbook:AMD64/Installation/Finalizing/ja "Handbook:AMD64/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:AMD64/Full/Working "Handbook:AMD64/Full/Working")[Portage について](/wiki/Handbook:AMD64/Working/Portage/ja "Handbook:AMD64/Working/Portage/ja")[USE フラグ](/wiki/Handbook:AMD64/Working/USE/ja "Handbook:AMD64/Working/USE/ja")[Portage の機能](/wiki/Handbook:AMD64/Working/Features/ja "Handbook:AMD64/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:AMD64/Working/Initscripts/ja "Handbook:AMD64/Working/Initscripts/ja")[環境変数](/wiki/Handbook:AMD64/Working/EnvVar/ja "Handbook:AMD64/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:AMD64/Full/Portage "Handbook:AMD64/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:AMD64/Portage/Files/ja "Handbook:AMD64/Portage/Files/ja")[変数](/wiki/Handbook:AMD64/Portage/Variables/ja "Handbook:AMD64/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:AMD64/Portage/Branches/ja "Handbook:AMD64/Portage/Branches/ja")[追加ツール](/wiki/Handbook:AMD64/Portage/Tools/ja "Handbook:AMD64/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:AMD64/Portage/CustomTree/ja "Handbook:AMD64/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:AMD64/Portage/Advanced/ja "Handbook:AMD64/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:AMD64/Full/Networking "Handbook:AMD64/Full/Networking")[はじめに](/wiki/Handbook:AMD64/Networking/Introduction/ja "Handbook:AMD64/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:AMD64/Networking/Advanced/ja "Handbook:AMD64/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:AMD64/Networking/Modular/ja "Handbook:AMD64/Networking/Modular/ja")[無線](/wiki/Handbook:AMD64/Networking/Wireless/ja "Handbook:AMD64/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:AMD64/Networking/Extending/ja "Handbook:AMD64/Networking/Extending/ja")[動的な管理](/wiki/Handbook:AMD64/Networking/Dynamic/ja "Handbook:AMD64/Networking/Dynamic/ja")

## Contents

- [1ブロックデバイスの概要](#.E3.83.96.E3.83.AD.E3.83.83.E3.82.AF.E3.83.87.E3.83.90.E3.82.A4.E3.82.B9.E3.81.AE.E6.A6.82.E8.A6.81)
  - [1.1ブロックデバイス](#.E3.83.96.E3.83.AD.E3.83.83.E3.82.AF.E3.83.87.E3.83.90.E3.82.A4.E3.82.B9)
  - [1.2パーティションテーブル](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.83.86.E3.83.BC.E3.83.96.E3.83.AB)
    - [1.2.1GUID パーティションテーブル (GPT)](#GUID_.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.83.86.E3.83.BC.E3.83.96.E3.83.AB_.28GPT.29)
    - [1.2.2マスターブートレコード (MBR) あるいは DOS ブートセクタ](#.E3.83.9E.E3.82.B9.E3.82.BF.E3.83.BC.E3.83.96.E3.83.BC.E3.83.88.E3.83.AC.E3.82.B3.E3.83.BC.E3.83.89_.28MBR.29_.E3.81.82.E3.82.8B.E3.81.84.E3.81.AF_DOS_.E3.83.96.E3.83.BC.E3.83.88.E3.82.BB.E3.82.AF.E3.82.BF)
  - [1.3高度なストレージ](#.E9.AB.98.E5.BA.A6.E3.81.AA.E3.82.B9.E3.83.88.E3.83.AC.E3.83.BC.E3.82.B8)
  - [1.4デフォルトのパーティション構成](#.E3.83.87.E3.83.95.E3.82.A9.E3.83.AB.E3.83.88.E3.81.AE.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E6.A7.8B.E6.88.90)
- [2パーティション構成の設計](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E6.A7.8B.E6.88.90.E3.81.AE.E8.A8.AD.E8.A8.88)
  - [2.1パーティション数とサイズ](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E6.95.B0.E3.81.A8.E3.82.B5.E3.82.A4.E3.82.BA)
  - [2.2スワップ領域について](#.E3.82.B9.E3.83.AF.E3.83.83.E3.83.97.E9.A0.98.E5.9F.9F.E3.81.AB.E3.81.A4.E3.81.84.E3.81.A6)
  - [2.3EFI システムパーティション (ESP) とは](#EFI_.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3_.28ESP.29_.E3.81.A8.E3.81.AF)
  - [2.4BIOS ブートパーティションとは](#BIOS_.E3.83.96.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.A8.E3.81.AF)
- [3UEFI 向けに GPT でディスクをパーティショニングする](#UEFI_.E5.90.91.E3.81.91.E3.81.AB_GPT_.E3.81.A7.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.82.92.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.8B.E3.83.B3.E3.82.B0.E3.81.99.E3.82.8B)
  - [3.1現在のパーティションレイアウトを表示する](#.E7.8F.BE.E5.9C.A8.E3.81.AE.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.83.AC.E3.82.A4.E3.82.A2.E3.82.A6.E3.83.88.E3.82.92.E8.A1.A8.E7.A4.BA.E3.81.99.E3.82.8B)
  - [3.2新しいディスクラベルを作成する / すべてのパーティションを削除する](#.E6.96.B0.E3.81.97.E3.81.84.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.83.A9.E3.83.99.E3.83.AB.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B_.2F_.E3.81.99.E3.81.B9.E3.81.A6.E3.81.AE.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E5.89.8A.E9.99.A4.E3.81.99.E3.82.8B)
  - [3.3EFI システムパーティション (ESP) を作成する](#EFI_.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3_.28ESP.29_.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [3.4スワップパーティションを作成する](#.E3.82.B9.E3.83.AF.E3.83.83.E3.83.97.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [3.5ルートパーティションを作成する](#.E3.83.AB.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [3.6パーティションのレイアウトを保存する](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E3.83.AC.E3.82.A4.E3.82.A2.E3.82.A6.E3.83.88.E3.82.92.E4.BF.9D.E5.AD.98.E3.81.99.E3.82.8B)
- [4BIOS / legacy ブート向けに MBR でディスクをパーティショニングする](#BIOS_.2F_legacy_.E3.83.96.E3.83.BC.E3.83.88.E5.90.91.E3.81.91.E3.81.AB_MBR_.E3.81.A7.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.82.92.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.8B.E3.83.B3.E3.82.B0.E3.81.99.E3.82.8B)
  - [4.1現在のパーティションレイアウトを表示する](#.E7.8F.BE.E5.9C.A8.E3.81.AE.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.83.AC.E3.82.A4.E3.82.A2.E3.82.A6.E3.83.88.E3.82.92.E8.A1.A8.E7.A4.BA.E3.81.99.E3.82.8B_2)
  - [4.2新しいディスクラベルを作成する / すべてのパーティションを削除する](#.E6.96.B0.E3.81.97.E3.81.84.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.83.A9.E3.83.99.E3.83.AB.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B_.2F_.E3.81.99.E3.81.B9.E3.81.A6.E3.81.AE.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E5.89.8A.E9.99.A4.E3.81.99.E3.82.8B_2)
  - [4.3ブートパーティションを作成する](#.E3.83.96.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [4.4スワップパーティションを作成する](#.E3.82.B9.E3.83.AF.E3.83.83.E3.83.97.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B_2)
  - [4.5ルートパーティションを作成する](#.E3.83.AB.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B_2)
  - [4.6パーティションのレイアウトを保存する](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E3.83.AC.E3.82.A4.E3.82.A2.E3.82.A6.E3.83.88.E3.82.92.E4.BF.9D.E5.AD.98.E3.81.99.E3.82.8B_2)
- [5ファイルシステムを作成する](#.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [5.1はじめに](#.E3.81.AF.E3.81.98.E3.82.81.E3.81.AB)
  - [5.2ファイルシステム](#.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
  - [5.3パーティションにファイルシステムを適用する](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AB.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.82.92.E9.81.A9.E7.94.A8.E3.81.99.E3.82.8B)
    - [5.3.1EFI システムパーティションのファイルシステム](#EFI_.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
    - [5.3.2レガシー BIOS ブートパーティションのファイルシステム](#.E3.83.AC.E3.82.AC.E3.82.B7.E3.83.BC_BIOS_.E3.83.96.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
    - [5.3.3小さな ext4 パーティション](#.E5.B0.8F.E3.81.95.E3.81.AA_ext4_.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3)
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

### パーティションテーブル

Linux システムを入れるために、（btrfs RAID を作成した場合のように）パーティショニングされていない生のディスクを使うことも理論上は可能ですが、実際にそのようなことを行うことはほとんどありません。代わりに、ディスク全体のブロックデバイスをより小さく扱いやすいブロックデバイスに分けて使います。 **amd64** システムでは、この分けられたブロックデバイスのことをパーティションと呼びます。現在主流なパーティショニング技術は、MBR（DOS ディスクラベルとも呼ばれる）と GPT の 2 つがあります。これらは 2 種類のブートプロセスに関連しています: レガシー BIOS ブートと UEFIです。

#### GUID パーティションテーブル (GPT)

_GUID パーティションテーブル (GPT)_ 構成 (GPT ディスクラベルとも呼ばれます) は、パーティションの識別子として 64 ビットの値を使います。パーティション情報を格納する領域は MBR パーティションテーブル (DOS ディスクラベル) の 512 バイトよりもずっと大きいため、パーティション数の制限はないようなものです。さらに、最大パーティションサイズもより大きいです (およそ 8 ZiB、そう、ゼビバイトです)。

オペレーティングシステムとファームウェアの間のソフトウェアインターフェースが（BIOS ではなく）UEFI ならば、DOS ディスクラベルでは互換性の問題が発生するので、GPT はほぼ必須となります。

GPTはまたチェックサムと冗長性も備えています。具体的にはヘッダやパーティションテーブルのエラーを検出するCRC32チェックサムや、ディスクの末尾にバックアップのGPTを持っています。もしディスク先頭にあるプライマリGPTに損害があっても、バックアップのGPTを使って回復できます。

**重要**

GPT にはいくつか注意点があります:

- BIOS ベースのコンピュータで GPT を使うことは可能ではありますが、Microsoft Windows は BIOS モードでは GPT パーティションからのブートを拒否するため、Microsoft Windows オペレーティングシステムとのデュアルブートを行うことはできないでしょう。
- 一部、バグのある（古い）マザーボードのファームウェアは、BIOS/CSM/legacy モードで起動するように設定されていると、GPT ラベルのディスクから起動する際に問題が発生する場合があります。

#### マスターブートレコード (MBR) あるいは DOS ブートセクタ

_[マスターブートレコード](https://en.wikipedia.org/wiki/Master_boot_record "wikipedia:Master boot record")_ ブートセクタ (DOS ブートセクタ、または DOS ディスクラベルとも呼ばれ、最近では GPT/UEFI 構成と対比してレガシー BIOS ブートとも呼ばれます) は、1983 年に PC DOS 2.x とともに最初に導入されました。MBR はパーティションの識別子として、32 ビットで、開始セクタとパーティションのセクタ数を使い、3 種類のパーティションタイプ (プライマリ、拡張、論理) を持っています。プライマリパーティションは、ディスク先頭のとても小さい領域 (ふつうは 512 バイト) にある MBR の中に、その情報が格納されます。この小ささのために、たった 4 つのプライマリパーティションしか使うことができません (例えば /dev/sda1 から /dev/sda4 まで)。

より多くのパーティションを使うために、プライマリパーティションのうちのひとつを _拡張_ パーティションとしてマークすることができます。拡張パーティションは追加の論理パーティションを複数格納することができます（パーティションの中にパーティションが存在することになります）。

**重要**

まだほとんどのマザーボードメーカーがサポートしてはいるものの、MBR ブートセクタと、それに関連するパーティションの制限は既に過去のものと考えられます。2010 年以前のハードウェアを扱っているのでない限り、 [GUID パーティションテーブル](https://en.wikipedia.org/wiki/GUID_Partition_Table "wikipedia:GUID Partition Table") でディスクをパーティショニングする方が良いでしょう。このセットアップを使って作業を続ける必要がある読者は、以下のことを認識しておいてください:

- 2010 年以降のほとんどのマザーボードは、MBR ブートセクタを利用するのを過去の(サポートはされているが理想的でない)ブートモードとみなします。
- 32 ビットの識別子を使用しているため、MBR のパーティションテーブルは 2 TiB を超えるサイズのストレージ空間のアドレスを指定することができません。
- 拡張パーティションを作成しない限り、MBR は最大で4つまでのパーティションしかサポートしません。
- このセットアップではバックアップのブートセクタは一切提供されないので、何かがパーティションテーブルを上書きしてしまうとすべてのパーティション情報が失われます。

とはいえ、MBR とレガシー BIOS ブートは AWS などの仮想化されたクラウド環境ではいまだに使われていることがあります。

ハンドブックの著者たちは、可能であればいつでも、Gentoo をインストールするためには [GPT](#GPT) を使うことを提案します。

### 高度なストレージ

公式 Gentoo ブートメディアは、高度なファイルシステムとツールを使用する様々な構成をサポートしています。これらを使用することで、より柔軟な変更、スナップショット、より高度なキャッシュなどの機能を使用することができます。

ハンドブックではその使用法について取り扱いませんが、そのようなシステムを稼働させるために役立つガイドのリストを以下に列挙しておきます:

- [Btrfs](/wiki/Btrfs/Native_System_Root_Guide "Btrfs/Native System Root Guide")
- [ZFS](/wiki/ZFS/rootfs "ZFS/rootfs")
- [LVM (Logical Volume Manager)](/wiki/LVM/ja "LVM/ja")

ディスクの暗号化も同様に取り扱います:

- [Rootfs の暗号化](/wiki/Rootfs_encryption/ja "Rootfs encryption/ja")

### デフォルトのパーティション構成

これよりこのハンドブックでは、ふたつの場合を考察し説明します:

1. GUID パーティションテーブル (GPT) のディスクとともに、UEFI ファームウェアを使用する。
2. MBR パーティションテーブルのディスクとともに、MBR DOS/レガシー BIOS ファームウェアを使用する。

特定のマザーボードファームウェアではブートタイプを混ぜて組み合わせることも可能ではありますが、このハンドブックの対象範囲からは外れます。上述の通り、現代的なハードウェアに対するインストールでは、GPT ディスクラベルを持つディスクで UEFI ブートを使用することが強く推奨されます。

シンプルな例として以下のパーティション構造を使います。

**重要**

以下の表の一行目は、GPT ディスクラベル _**または**_ MBR DOS/レガシー BIOS ディスクラベルの _**どちらか一方**_ を選択したときの、それぞれに対応する排他的な情報を記載しています。疑わしい場合は、2010 年以降に製造された **amd64** マシンは一般的に UEFI ファームウェアと GPT ブートセクタをサポートしているので、GPT で進めてください。

パーティション
ファイルシステム
サイズ
説明
/dev/sda1fat32 EFI システムパーティションのために必要なファイルシステム。常に GPT ディスクラベルと関連付けられます。1 GiB
EFI システムパーティションの詳細。_UEFI 実装をサポートするシステムファームウェアで適用可能です。これは典型的には 2010 年頃から現在までに製造されたシステムが該当します。_xfs MBR パーティションテーブルのブートパーティションにおすすめのファイルシステム。DOS/レガシー BIOS ディスクラベルに制限された古いファームウェアと組み合わせて使用されます。MBR DOS/レガシー BIOS ブートパーティションの詳細。_レガシー BIOS マシンのファームウェアで適用可能です。この種のシステムは典型的には 2010 年<u>より前</u>に製造されたシステムが該当し、一般論として徐々に製造が中止されてきています。_/dev/sda2linux-swap
RAM サイズ \* 2
スワップパーティションの詳細。
/dev/sda3xfs
ディスクの残り 選択された _**プロファイル**_、追加の _**パーティション**_ (任意)、およびシステムの _**目的**_ によって rootfs のサイズの見積もりの複雑さは高まるので、ハンドブックの作者は rootfs パーティションに対する「万能」の提案を提供することはできません。<br></br>そのディスクを使用する唯一の OS が Gentoo なら、ディスクの残りを選択するのが最も安全かつ提案される選択肢です。ルートパーティションの詳細。

もしこの情報だけで十分なほど熟練した読者は、実際のパーティション作成に進んで構いません。

fdisk と parted はいずれも、公式 Gentoo live イメージ環境に含まれるパーティショニングのためのユーティリティです。fdisk はよく知られていて、安定した、MBR と GPT の両ディスクを扱うことができるツールです。parted は GPT パーティションをサポートした、最初期の Linux ブロックデバイスの管理ツールの一つです。読者が望むのであれば fdisk への代替として使用することができますが、ハンドブックでは、多くの Linux 環境で広く利用できる fdisk のための手順のみを提供します。

パーティションの生成方法に進む前に、以降の数セッションでパーティション構造がどのように生成されるのかについて、その詳細を述べ、いくつかの共通した落とし穴について触れておきます。

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

### EFI システムパーティション (ESP) とは

オペレーティングシステムを起動するのに (BIOS ではなく) UEFI を使うシステムに Gentoo をインストールするときは、EFI システムパーティションを作成することが必須です。この手順については後述の説明でも述べます。 **BIOS/レガシーモードで起動する場合には、EFI システムパーティションは不要です。**

ESP は FAT 系列のファイルシステム (Linux システムでは _vfat_ と表示することもあります) である必要があります。 [UEFI specification](http://www.uefi.org/sites/default/files/resources/UEFI%202_5.pdf) では、UEFI ファームウェアは FAT12、16、32 を認識すると書かれている一方で、ESP には FAT32 を推奨しています。パーティションを作成したら、ESP をフォーマットしてください:

`root #` `mkfs.fat -F 32 /dev/sda1`

**重要**

ESP が FAT 系列のファイルシステムでフォーマットされていないと、UEFI ファームウェアはブートローダー (か Linux カーネル) を見つけられず、おそらくシステムをブートすることができません！

### BIOS ブートパーティションとは

_BIOS ブートパーティション_ はとても小さい (1 - 2 MB) パーティションで、GRUB2 などのブートローダが、与えられた領域に収まらないようなデータを置くためのパーティションです。ディスクが GPT ディスクラベルを使ってフォーマットされているが、システムのファームウェアがレガシー BIOS/MBR DOS ブートモードの GRUB2 を利用してブートしようとしている場合にのみ、必要になります。 **EFI/UEFI モードで起動する場合には _不要_ で、MBR/レガシー DOS ディスクラベルを使用する場合にも _不要_ です。** _BIOS ブートパーティション_ はこのガイドでは使用しません。

## UEFI 向けに GPT でディスクをパーティショニングする

以降の部分では、単一の GPT ディスクデバイスのためのパーティションレイアウト例を、UEFI 仕様書および Discoverable Partitions Specification (DPS) に準拠するように作成する方法を説明します。DPS は Linux Userspace API (UAPI) Group Specification の一部として提供されている仕様であり、準拠することが推奨されますが、必須ではありません。この仕様を、[sys-apps/util-linux](https://packages.gentoo.org/packages/sys-apps/util-linux) パッケージの一部である fdisk ユーティリティを使って実装します。

表は単純な Gentoo インストールのために推奨されるデフォルトを提供します。個人の好みやシステムの設計目的に応じて追加のパーティションを追加してもかまいません。

デバイスパス (sysfs)
マウントポイント
ファイルシステム
DPS UUID (Type-UUID)
説明
/dev/sda1/efivfat
c12a7328-f81f-11d2-ba4b-00a0c93ec93b
EFI システムパーティション (ESP) の詳細。
/dev/sda2なし。スワップはデバイスファイルのようにファイルシステムにマウントされることはありません。swap
0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
スワップパーティションの詳細。
/dev/sda3/xfs
4f68bce3-e8cd-4db1-96e7-fbcaf984b709
ルートパーティションの詳細。

### 現在のパーティションレイアウトを表示する

fdiskは、ディスクをパーティション分割するためのポピュラーでパワフルなツールです。ディスク（我々の例では/dev/sda）に対してfdiskを起動しましょう。

`root #` `fdisk /dev/sda`

`p` キーを使えば、現在のディスクのパーティション構成を表示できます。

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

Device        Start        End    Sectors   Size Type
/dev/sda1      2048    2099199    2097152     1G EFI System
/dev/sda2   2099200   10487807    8388608     4G Linux swap
/dev/sda3  10487808 1953523711 1943035904 926.5G Linux root (x86-64)

```

このディスクは 2 つの Linux ファイルシステム ("Linux" と書かれているパーティションに対応します) と 1 つの swap パーティション ("Linux swap" と書かれているパーティション) で構成されているようです。

### 新しいディスクラベルを作成する / すべてのパーティションを削除する

`g` キーを押下するとすぐに、既存のパーティションがすべて削除され、新しい GPT ディスクラベルが作成されるでしょう:

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 3E56EE74-0571-462B-A992-9872E3855D75).

```

または、既存の GPT ディスクラベル (上の `p` の出力を確認してください) を保つために、代わりに既存のパーティションをひとつずつ削除することを検討してください。パーティションを削除するには `d` を押下します。例えば /dev/sda1 を削除するにはこのようにします:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

これで指定したパーティションの削除が予約されました。パーティションの一覧 ( `p`) にはもう現れませんが、変更を保存するまで実際の消去は行われないので、間違えて操作してしまった場合は中止することができます。すぐに `q` を押して `Enter` を押せば、パーティションは削除されません。

`p` でパーティションの一覧を表示して `d` とパーティション番号を入力する、という作業を繰り返すと、パーティションテーブルは空っぽになります。

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

```

さて、メモリ内のパーティションテーブルが空になり、パーティションを作る準備ができました。

### EFI システムパーティション (ESP) を作成する

**メモ**

より小さい ESP にすることもできますが、推奨はされません。他の OS と共有するかもしれない場合は特に。

まずは、/efi としてマウントされることになる小さな EFI システムパーティションを作成します。新規パーティションを作るので `n` を入力し、 `1` で最初の基本パーティションを選択しましょう。開始セクタについて聞かれたら、2048 (ブートローダーのために必要になるかもしれません) になっていることを確認して `Enter` を押しましょう。終了セクタの指定では、1 ギビバイトのパーティションを作るので +1G と入力します:

`Command (m for help):` `n`

```
Partition number (1-128, default 1): 1
First sector (2048-1953525134, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525134, default 1953523711): +1G

Created a new partition 1 of type 'Linux filesystem' and of size 1 GiB.
Partition #1 contains a vfat signature.

Do you want to remove the signature? [Y]es/[N]o: Y
The signature will be removed by a write command.

```

パーティションを EFI システムパーティションとしてマークしてください:

`Command (m for help):` `t`

```
Selected partition 1
Partition type or alias (type L to list all): 1
Changed type of partition 'Linux filesystem' to 'EFI System'.

```

### スワップパーティションを作成する

次に、スワップパーティションを作成したいので、新規パーティション作成の `n` を押下し、 `2` で 2 番目のパーティション、/dev/sda2 を選択しましょう。開始セクタの指定ではそのまま `Enter` を押します。終了セクタの指定では、4 GiB のパーティションを作るので +4G (もしくはお好みの swap 領域のサイズ) と入力します。

`Command (m for help):` `n`

```
Partition number (2-128, default 2):
First sector (2099200-1953525134, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525134, default 1953523711): +4G

Created a new partition 2 of type 'Linux filesystem' and of size 4 GiB.

```

この後、パーティションタイプを設定するために `t` を押下し、今作成したパーティション `2` を選択、そしてパーティションタイプ "Linux Swap" を意味する _19_ を入力します。

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type or alias (type L to list all): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### ルートパーティションを作成する

最後に、ルートパーティションを作成します。 `n` で新規パーティション作成、3 番目のパーティションである /dev/sda3 を作成するために `3` を押下、最初のセクタはそのまま `Enter` を押します。最後のセクタを聞かれたら、ディスクの空き領域全てをこのパーティションに割り当てたいのでそのまま `Enter` を押しましょう。

`Command (m for help):` `n`

```
Partition number (3-128, default 3): 3
First sector (10487808-1953525134, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525134, default 1953523711):

Created a new partition 3 of type 'Linux filesystem' and of size 926.5 GiB..

```

**メモ**

ルートパーティションのタイプを "Linux root (x86-64)" に設定する必要はなく、"Linux filesystem" タイプに設定されていてもシステムは通常通り機能するでしょう。このファイルシステムタイプは、それをサポートするブートローダ (つまり systemd-boot) を使用し、fstab ファイルを使用したくない場合のみ必要です。

ルートパーティションを作成したら、パーティションタイプを設定するために `t` を押下し、今作成したパーティション `3` を選択、そしてパーティションタイプ "Linux Root (x86-64)" を意味する _23_ を入力します。

`Command(m for help):` `t`

```
Partition number (1-3, default 3): 3
Partition type or alias (type L to list all): 23

Changed type of partition 'Linux filesystem' to 'Linux root (x86-64)'

```

これが終わったら、 `p` で次のようなパーティションテーブルが表示されるはずです:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

Device        Start        End    Sectors   Size Type
/dev/sda1      2048    2099199    2097152     1G EFI System
/dev/sda2   2099200   10487807    8388608     4G Linux swap
/dev/sda3  10487808 1953523711 1943035904 926.5G Linux root (x86-64)

Filesystem/RAID signature on partition 1 will be wiped.

```

### パーティションのレイアウトを保存する

このパーティションレイアウトを保存して fdisk ユーティリティを終了するために、 `w` を押下してください。

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

これでパーティションが利用可能になったので、次のインストール手順はここにファイルシステムを置いていくことです。

## BIOS / legacy ブート向けに MBR でディスクをパーティショニングする

以下の表は、単純な MBR / レガシー BIOS ブートでのインストールに推奨されるパーティションレイアウトを提供します。個人の好みやシステムの設計目的に応じて追加のパーティションを追加してもかまいません。

デバイスパス (sysfs)
マウントポイント
ファイルシステム
DPS UUID (PARTUUID)
説明
/dev/sda1/bootxfs
なし
MBR DOS / レガシー BIOS ブートパーティションの詳細。
/dev/sda2なし。スワップはデバイスファイルのようにファイルシステムにマウントされることはありません。swap
0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
スワップパーティションの詳細。
/dev/sda3/xfs
4f68bce3-e8cd-4db1-96e7-fbcaf984b709
ルートパーティションの詳細。

パーティションレイアウトはお好みで変更してください。

### 現在のパーティションレイアウトを表示する

ディスク（我々の例では/dev/sda）に対してfdiskを起動しましょう。

`root #` `fdisk /dev/sda`

`p` キーを使えば、現在のディスクのパーティション構成を表示できます。

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

このディスクは現時点まで、GPT テーブルを使用して、2 つの Linux ファイルシステム ("Linux" と書かれているパーティションに対応します) と 1 つの swap パーティション ("Linux swap" と書かれているパーティション) で構成されているようです。

### 新しいディスクラベルを作成する / すべてのパーティションを削除する

`o` キーを押下するとすぐに、既存のパーティションがすべて削除され、新しい MBR ディスクラベル (DOS ディスクラベルとも呼ばれます) が作成されるでしょう:

`Command (m for help):` `o`

```
Created a new DOS disklabel with disk identifier 0xf163b576.
The device contains 'gpt' signature and it will be removed by a write command. See fdisk(8) man page and --wipe option for more details.

```

または、既存の DOS ディスクラベル (上の `p` の出力を確認してください) を保つために、代わりに既存のパーティションをひとつずつ削除することを検討してください。パーティションを削除するには `d` を押下します。例えば /dev/sda1 を削除するにはこのようにします:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

これで指定したパーティションの削除が予約されました。パーティションの一覧 ( `p`) にはもう現れませんが、変更を保存するまで実際の消去は行われないので、間違えて操作してしまった場合は中止することができます。すぐに `q` を入力して `Enter` を押せば、パーティションは削除されません。

`p` でパーティションの一覧を表示して `d` とパーティション番号を押下する、という作業を繰り返すと、パーティションテーブルは空っぽになります。

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

```

これで、新しいパーティションを作るためのディスクの準備ができました。

### ブートパーティションを作成する

まずは、/boot としてマウントされることになる小さなパーティションを作成します。新規パーティションを作るので `n` を押下し、 `p` で _基本_ パーティションを選択、 `1` で最初の基本パーティションを選択しましょう。開始セクタについて聞かれたら、2048 (ブートローダーのために必要になるかもしれません) になっていることを確認して `Enter` を押しましょう。終了セクタの指定では、1 GB のパーティションを作るので +1G と入力します:

`Command (m for help):` `n`

```
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-1953525167, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525167, default 1953525167): +1G

Created a new partition 1 of type 'Linux' and of size 1 GiB.

```

`a` キーを押して、 `Enter` を押すことで、パーティションをブート可能としてマークしてください:

`Command (m for help):` `a`

```
Selected partition 1
The bootable flag on partition 1 is enabled now.

```

メモ: ディスク上で複数のパーティションが利用可能な場合は、ブート可能としてマークしたいパーティションを選択する必要があるでしょう。

### スワップパーティションを作成する

次に、スワップパーティションを作成したいので、新規パーティション作成の `n` を押下し、 `p` で基本パーティションを選択し、 `2` で 2 番目の基本パーティション、/dev/sda2 を選択しましょう。開始セクタの指定ではそのまま `Enter` を押します。終了セクタの指定では、4GB のパーティションを作るので +4G (もしくはお好みの swap 領域のサイズ) と入力します。

`Command (m for help):` `n`

```
Partition type
   p   primary (1 primary, 0 extended, 3 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (2-4, default 2): 2
First sector (2099200-1953525167, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525167, default 1953525167): +4G

Created a new partition 2 of type 'Linux' and of size 4 GiB.

```

ここまでできたら、パーティションタイプを設定するために `t` を押下し、今作成したパーティション `2` を選択、そしてパーティションタイプ "Linux Swap" を意味する _82_ を入力します。

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Hex code (type L to list all codes): 82

Changed type of partition 'Linux' to 'Linux swap / Solaris'.

```

### ルートパーティションを作成する

最後に、ルートパーティションを作成します。 `n` で新規パーティション作成、3 番目の基本パーティション、/dev/sda3 を作成するために `p` と `3` を押下、最初のセクタはそのまま `Enter` を押します。最後のセクタを聞かれたら、ディスクの空き領域全てをこのパーティションに割り当てたいのでそのまま `Enter` を押しましょう:

`Command (m for help):` `n`

```
Partition type
   p   primary (2 primary, 0 extended, 2 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (3,4, default 3): 3
First sector (10487808-1953525167, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525167, default 1953525167):

Created a new partition 3 of type 'Linux' and of size 926.5 GiB.

```

これが終わったら、 `p` で次のようなパーティションテーブルが表示されるはずです:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

### パーティションのレイアウトを保存する

このパーティションレイアウトを書き込んで fdisk を終了するために、 `w` を押下してください:

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

それでは、パーティション上にファイルシステムを作成していきましょう。

## ファイルシステムを作成する

**警告**

SSD または NVMe ドライブを使用する場合は、ファームウェアアップグレードがあるかどうか確認するのが賢明です。特に一部の Intel SSDs (600p および 6000p) は、XFS の I/O 使用量パターンによって誘発される [データ破損の可能性](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) のためのファームウェアアップグレードが必要です。この問題はファームウェアレベルのもので、XFS ファイルシステムの欠陥ではありません。smartctl ユーティリティはデバイスのモデルとファームウェアバージョンを確認するのに役立ちます。

### はじめに

パーティションが作成できたら、その上にファイルシステムを作成します。次の節ではLinuxがサポートする各種ファイルシステムを紹介します。どのファイルシステムを使うかをすでに決めているなら、 [パーティションにファイルシステムを適用する](/wiki/Handbook:AMD64/Installation/Disks/ja#Applying_a_filesystem_to_a_partition "Handbook:AMD64/Installation/Disks/ja") へ進みましょう。そうでなければ、次の節を読んで利用可能なファイルシステムについて知るのがよいでしょう。

### ファイルシステム

Linux は多くのファイルシステムをサポートしていますが、それらの多くは特定の目的をもって配備するのが賢明なものです。特定のファイルシステムのみが amd64 アーキテクチャ上で安定して動作するとされています - 重要なパーティションに実験的なファイルシステムを選択するときは、事前にファイルシステムのサポート状況を十分に知っておくことを推奨します。 **XFS はすべてのプラットフォームで、すべての目的で推奨されるファイルシステムです。** 以下は、網羅的ではないリストです:

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

#### EFI システムパーティションのファイルシステム

EFI システムパーティション (/dev/sda1) は FAT32 としてフォーマットする必要があります:

`root #` `mkfs.vfat -F 32 /dev/sda1`

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

以前に開始したインストール手順を完了しなかった場合は、ハンドブックのこの時点からインストールを再開することができます。このリンクを permalink として使用してください: [インストールの再開はここから](/wiki/Handbook:AMD64/Installation/Disks/ja#Resumed_installations_start_here "Handbook:AMD64/Installation/Disks/ja")。

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

EFI インストールの場合、ESP はルートパーティションの場所の下にマウントすべきです:

`root #` `mkdir --parents /mnt/gentoo/efi`

必要に応じて、mount コマンドを使用して、追加の (カスタムの) パーティションのマウントを続けてください。

**メモ**

もし/tmp/を別のパーティションに置く必要があるなら、マウントしたあと権限の変更を忘れずに行ってください:

`root #` `chmod 1777 /mnt/gentoo/tmp`

/var/tmpについても同様です。

このあと解説の中で、proc ファイルシステム (仮想的なカーネルとのインターフェース) が、他のカーネル擬似ファイルシステムと同様にマウントされますが、まず最初は、 [Gentoo stage ファイル](/wiki/Handbook:AMD64/Installation/Stage/ja "Handbook:AMD64/Installation/Stage/ja") を展開する必要があります。

[← ネットワーク設定](/wiki/Handbook:AMD64/Installation/Networking/ja "Previous part") [Home](/wiki/Handbook:AMD64/ja "Handbook:AMD64/ja") [Gentoo インストールファイルをインストール →](/wiki/Handbook:AMD64/Installation/Stage/ja "Next part")