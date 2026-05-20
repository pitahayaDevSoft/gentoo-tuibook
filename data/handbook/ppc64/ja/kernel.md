# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Kernel/de "Handbuch:PPC64/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Kernel "Handbook:PPC64/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Kernel/es "Manual de Gentoo: PPC64/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Kernel/fr "Handbook:PPC64/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Kernel/it "Handbook:PPC64/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Kernel/hu "Handbook:PPC64/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Kernel/pl "Handbook:PPC64/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Kernel/pt-br "Handbook:PPC64/Installation/Kernel/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Kernel/ru "Handbook:PPC64/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Kernel/ta "கையேடு:PPC64/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Kernel/zh-cn "手册：PPC64/安装/配置Linux内核 (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:PPC64/Installation/Kernel/ko "Handbook:PPC64/Installation/Kernel/ko (100% translated)")

[PPC64 ハンドブック](/wiki/Handbook:PPC64 "Handbook:PPC64")[インストール](/wiki/Handbook:PPC64/Full/Installation "Handbook:PPC64/Full/Installation")[インストールについて](/wiki/Handbook:PPC64/Installation/About/ja "Handbook:PPC64/Installation/About/ja")[メディアの選択](/wiki/Handbook:PPC64/Installation/Media/ja "Handbook:PPC64/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:PPC64/Installation/Networking/ja "Handbook:PPC64/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:PPC64/Installation/Disks/ja "Handbook:PPC64/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:PPC64/Installation/Stage/ja "Handbook:PPC64/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:PPC64/Installation/Base/ja "Handbook:PPC64/Installation/Base/ja")カーネルの設定[システムの設定](/wiki/Handbook:PPC64/Installation/System/ja "Handbook:PPC64/Installation/System/ja")[ツールのインストール](/wiki/Handbook:PPC64/Installation/Tools/ja "Handbook:PPC64/Installation/Tools/ja")[ブートローダの設定](/wiki/Handbook:PPC64/Installation/Bootloader/ja "Handbook:PPC64/Installation/Bootloader/ja")[締めくくり](/wiki/Handbook:PPC64/Installation/Finalizing/ja "Handbook:PPC64/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:PPC64/Full/Working "Handbook:PPC64/Full/Working")[Portage について](/wiki/Handbook:PPC64/Working/Portage/ja "Handbook:PPC64/Working/Portage/ja")[USE フラグ](/wiki/Handbook:PPC64/Working/USE/ja "Handbook:PPC64/Working/USE/ja")[Portage の機能](/wiki/Handbook:PPC64/Working/Features/ja "Handbook:PPC64/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:PPC64/Working/Initscripts/ja "Handbook:PPC64/Working/Initscripts/ja")[環境変数](/wiki/Handbook:PPC64/Working/EnvVar/ja "Handbook:PPC64/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:PPC64/Full/Portage "Handbook:PPC64/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:PPC64/Portage/Files/ja "Handbook:PPC64/Portage/Files/ja")[変数](/wiki/Handbook:PPC64/Portage/Variables/ja "Handbook:PPC64/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:PPC64/Portage/Branches/ja "Handbook:PPC64/Portage/Branches/ja")[追加ツール](/wiki/Handbook:PPC64/Portage/Tools/ja "Handbook:PPC64/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:PPC64/Portage/CustomTree/ja "Handbook:PPC64/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:PPC64/Portage/Advanced/ja "Handbook:PPC64/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:PPC64/Full/Networking "Handbook:PPC64/Full/Networking")[はじめに](/wiki/Handbook:PPC64/Networking/Introduction/ja "Handbook:PPC64/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:PPC64/Networking/Advanced/ja "Handbook:PPC64/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:PPC64/Networking/Modular/ja "Handbook:PPC64/Networking/Modular/ja")[無線](/wiki/Handbook:PPC64/Networking/Wireless/ja "Handbook:PPC64/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:PPC64/Networking/Extending/ja "Handbook:PPC64/Networking/Extending/ja")[動的な管理](/wiki/Handbook:PPC64/Networking/Dynamic/ja "Handbook:PPC64/Networking/Dynamic/ja")

## Contents

- [1ファームウェアとマイクロコードのインストール](#.E3.83.95.E3.82.A1.E3.83.BC.E3.83.A0.E3.82.A6.E3.82.A7.E3.82.A2.E3.81.A8.E3.83.9E.E3.82.A4.E3.82.AF.E3.83.AD.E3.82.B3.E3.83.BC.E3.83.89.E3.81.AE.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB)
  - [1.1ファームウェア](#.E3.83.95.E3.82.A1.E3.83.BC.E3.83.A0.E3.82.A6.E3.82.A7.E3.82.A2)
    - [1.1.1提案: Linux ファームウェア](#.E6.8F.90.E6.A1.88:_Linux_.E3.83.95.E3.82.A1.E3.83.BC.E3.83.A0.E3.82.A6.E3.82.A7.E3.82.A2)
      - [1.1.1.1ファームウェアのロード](#.E3.83.95.E3.82.A1.E3.83.BC.E3.83.A0.E3.82.A6.E3.82.A7.E3.82.A2.E3.81.AE.E3.83.AD.E3.83.BC.E3.83.89)
- [2sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [2.1ブートローダ](#.E3.83.96.E3.83.BC.E3.83.88.E3.83.AD.E3.83.BC.E3.83.80)
    - [2.1.1GRUB](#GRUB)
    - [2.1.2伝統的なレイアウト、他のブートローダ (例: (e)lilo、syslinux 等)](#.E4.BC.9D.E7.B5.B1.E7.9A.84.E3.81.AA.E3.83.AC.E3.82.A4.E3.82.A2.E3.82.A6.E3.83.88.E3.80.81.E4.BB.96.E3.81.AE.E3.83.96.E3.83.BC.E3.83.88.E3.83.AD.E3.83.BC.E3.83.80_.28.E4.BE.8B:_.28e.29lilo.E3.80.81syslinux_.E7.AD.89.29)
  - [2.2Initramfs](#Initramfs)
    - [2.2.1Chroot detection](#Chroot_detection)
- [3カーネルのコンフィギュレーションとコンパイル](#.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.81.AE.E3.82.B3.E3.83.B3.E3.83.95.E3.82.A3.E3.82.AE.E3.83.A5.E3.83.AC.E3.83.BC.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.A8.E3.82.B3.E3.83.B3.E3.83.91.E3.82.A4.E3.83.AB)
  - [3.1ディストリビューションカーネル](#.E3.83.87.E3.82.A3.E3.82.B9.E3.83.88.E3.83.AA.E3.83.93.E3.83.A5.E3.83.BC.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB)
    - [3.1.1省略可能: 署名付きカーネルモジュール](#.E7.9C.81.E7.95.A5.E5.8F.AF.E8.83.BD:_.E7.BD.B2.E5.90.8D.E4.BB.98.E3.81.8D.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.83.A2.E3.82.B8.E3.83.A5.E3.83.BC.E3.83.AB)
    - [3.1.2ディストリビューションカーネルをインストールする](#.E3.83.87.E3.82.A3.E3.82.B9.E3.83.88.E3.83.AA.E3.83.93.E3.83.A5.E3.83.BC.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.92.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB.E3.81.99.E3.82.8B)
    - [3.1.3アップグレードと後処理](#.E3.82.A2.E3.83.83.E3.83.97.E3.82.B0.E3.83.AC.E3.83.BC.E3.83.89.E3.81.A8.E5.BE.8C.E5.87.A6.E7.90.86)
    - [3.1.4インストール/アップグレード後タスク](#.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB.2F.E3.82.A2.E3.83.83.E3.83.97.E3.82.B0.E3.83.AC.E3.83.BC.E3.83.89.E5.BE.8C.E3.82.BF.E3.82.B9.E3.82.AF)
      - [3.1.4.1手動で initramfs または Unified カーネルイメージを再ビルドする](#.E6.89.8B.E5.8B.95.E3.81.A7_initramfs_.E3.81.BE.E3.81.9F.E3.81.AF_Unified_.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.A4.E3.83.A1.E3.83.BC.E3.82.B8.E3.82.92.E5.86.8D.E3.83.93.E3.83.AB.E3.83.89.E3.81.99.E3.82.8B)
  - [3.2別の方法: マニュアル設定](#.E5.88.A5.E3.81.AE.E6.96.B9.E6.B3.95:_.E3.83.9E.E3.83.8B.E3.83.A5.E3.82.A2.E3.83.AB.E8.A8.AD.E5.AE.9A)
  - [3.3カーネルソースのインストール](#.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.BD.E3.83.BC.E3.82.B9.E3.81.AE.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB)
    - [3.3.1Modprobed-db プロセス](#Modprobed-db_.E3.83.97.E3.83.AD.E3.82.BB.E3.82.B9)
    - [3.3.2手動プロセス](#.E6.89.8B.E5.8B.95.E3.83.97.E3.83.AD.E3.82.BB.E3.82.B9)
      - [3.3.2.1必要なオプションを有効化する](#.E5.BF.85.E8.A6.81.E3.81.AA.E3.82.AA.E3.83.97.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E6.9C.89.E5.8A.B9.E5.8C.96.E3.81.99.E3.82.8B)
    - [3.3.3省略可能: 署名済みカーネルモジュール](#.E7.9C.81.E7.95.A5.E5.8F.AF.E8.83.BD:_.E7.BD.B2.E5.90.8D.E6.B8.88.E3.81.BF.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.83.A2.E3.82.B8.E3.83.A5.E3.83.BC.E3.83.AB)
    - [3.3.4コンパイルおよびインストール](#.E3.82.B3.E3.83.B3.E3.83.91.E3.82.A4.E3.83.AB.E3.81.8A.E3.82.88.E3.81.B3.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB)

## ファームウェアとマイクロコードのインストール

### ファームウェア

#### 提案: Linux ファームウェア

多くのシステムでは、特定のハードウェアが正しく機能するための非 FOSS なファームウェアが必要です。[sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) パッケージは、すべてではありませんが、多くのデバイスのためのファームウェアを含んでいます。

**ヒント**

ほとんどの無線カードと GPU は、機能するためにファームウェアを必要とします。

`root #` `emerge --ask sys-kernel/linux-firmware`

**メモ**

一部のファームウェアパッケージのインストールには、関連するファームウェアライセンスを受諾する必要があることがよくあります。必要であれば、ライセンスの受諾についてはハンドブックの [ライセンスの取り扱いの節](/wiki/Handbook:PPC64/Working/Portage/ja#Licenses "Handbook:PPC64/Working/Portage/ja") を確認してください。

##### ファームウェアのロード

ファームウェアファイルは、典型的には、それに関連付けられたカーネルモジュールがロードされたときにロードされます。これはつまり、カーネルモジュールが _M_ ではなく _Y_ に設定されている場合、ファームウェアは **CONFIG\_EXTRA\_FIRMWARE** を使用してカーネルに組み込まれていなくてはならないということです。ほとんどの場合、ファームウェアを必要とするモジュールの組み込みは、ロードを複雑にしたり、失敗の原因になったりすることがあります。

## sys-kernel/installkernel

[Installkernel](/wiki/Installkernel/ja "Installkernel/ja") は、カーネルのインストール、 [initramfs](/wiki/Initramfs "Initramfs") の生成、 [unified カーネルイメージ](/wiki/Unified_kernel_image "Unified kernel image") の生成、そして何よりもブートローダの設定を自動化するために、使用することができます。[sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) はこれを達成するための 2 通りの道を実装しています: Debian に由来する伝統的な installkernel と、 [systemd](/wiki/Systemd/ja "Systemd/ja") の kernel-install です。どちらを選ぶべきかは、何よりもシステムのブートローダによって変わります。デフォルトでは、systemd プロファイルでは systemd の kernel-install が使用される一方、それ以外のプロファイルでは伝統的な installkernel がデフォルトです。

### ブートローダ

今こそ、システムでどのブートローダを使いたいかについて考えるときです。よく分からない場合は、下の「伝統的なレイアウト」のサブ節に従ってください。

**重要**

以下のサブ節のうち一つだけが必要です。どれを使用すべきか分からない場合は、とりあえずは最初に記載されているもので進めてください。もし必要なら、後でいつでも切り換えることができます。

#### GRUB

GRUB を使用する場合は、systemd の kernel-install または伝統的な Debian installkernel のいずれかを使用することができます。[systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグはこれらの実装の切り換えを行います。カーネルをインストールするときに自動的に grub-mkconfig を実行するためには、[grub](https://packages.gentoo.org/useflags/grub) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") [USE フラグ](/wiki/USE_flag/ja "USE flag/ja") を有効化してください。

**メモ**

GRUB requires kernels to be installed to /boot.

ファイル **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

**メモ**

systemd-boot requires kernels to be installed to /efi.

**メモ**

UEFI の構成設定のために [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) を使用する場合は、カーネルをインストールしようとする前に kernel-bootcfg-boot-successful サービスが有効化されていることを確認してください。この EFI スタブブートの実装は systemd システムでのデフォルトです。

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**メモ**

EFIstub requires kernels to be installed to /efi.

#### 伝統的なレイアウト、他のブートローダ (例: (e)lilo、syslinux 等)

[grub](https://packages.gentoo.org/useflags/grub) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja")、[systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja")、[efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") または [uki](https://packages.gentoo.org/useflags/uki) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグが有効化されて **いない** 場合は、伝統的な /boot レイアウト (例えば (e)LILO、syslinux 等向けの) がデフォルトで使用されます。追加の操作は必要ありません。

### Initramfs

システムがブートするために、初期 RAM ファイルシステム ( **init** ial **ram**-based **f** ile **s** ystem)、または [initramfs](/wiki/Initramfs "Initramfs") が必要になる場合があります。必要になるケースは様々なものがありますが、よくあるのは以下の場合です:

- ストレージ/ファイルシステムのドライバーがモジュールになっているカーネル。
- /usr/ または /var/ が分離されたパーティション上にあるレイアウト。
- 暗号化されたルートファイルシステム。

**ヒント**

[ディストリビューションカーネル](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") は、多くのストレージとファイルシステムのドライバがモジュールとしてビルドされるので、initramfs とともに使用するように設計されています。

ルートファイルシステムのマウントに加えて、initramfs は以下のような他のタスクも行うことができます:

- 正常でないシステムシャットダウンの発生時にファイルシステムの一貫性を確認し修復するためのツールである、 **f** ile **s** ystem **c** onsistency chec **k** fsck を実行する。
- ブート後期での失敗時に回復環境を提供する。

[Installkernel](/wiki/Installkernel/ja "Installkernel/ja") は、[dracut](https://packages.gentoo.org/useflags/dracut) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") または [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグが有効化されている場合、カーネルのインストール時に initramfs を自動的に生成することができます:

ファイル **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Chroot detection

Bootloaders such as [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") and [EFI stub](/wiki/EFI_stub/ja "EFI stub/ja") use the kernel arguments of the running system as set in /proc/cmdline by default. Because of the wide range of ways Gentoo can be installed users will randomly get tripped up by this.

To help solve any issues this may cause, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) will check if it is running in a chroot and fail if the kernel command line is not explicitly configured. Otherwise the bootloader would use the install media's boot arguments which would lead to boot failure.

One way to satisfy [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is by using Dracut's config file to set the root partition UUID as shown below, or alternatively for more information on what this check helps with and different ways to configure it, see [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check.2Fja "Installkernel").

`root #` `mkdir /etc/dracut.conf.d`

`root #` `blkid`

```
/dev/sda3: UUID="2122cd72-94d7-4dcc-821e-3705926deecc"
```

上の例では、ルートパーティションは /dev/sda3 で、UUID は 2122cd72-94d7-4dcc-821e-3705926deecc です。
その場合 dracut コンフィグファイルは次のようになるでしょう:

ファイル **`/etc/dracut.conf.d/00-installkernel.conf`**

```
kernel_cmdline=" root=UUID=2122cd72-94d7-4dcc-821e-3705926deecc "

```

`root #` `emerge --ask sys-kernel/installkernel`

## カーネルのコンフィギュレーションとコンパイル

**ヒント**

最初のブートでは dist-kernel を使用するのが懸命な選択かもしれません。システムの問題とカーネルコンフィグの問題を除外するための、とても簡潔な手法を提供しているからです。機能すると分かっているカーネルをフォールバックとして常に持っておくことは、デバッグの効率を高め、更新時にシステムがブートできなくなるかもという不安を緩和させるでしょう。

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**重要**

以下のサブ節のうち一つだけが必要です。どれを使用すべきか分からない場合は、とりあえずは最初に記載されているもので進めてください。もし必要なら、後でいつでも切り換えることができます。

簡単なものから込み入ったものへ、順に並べると:

[完全自動アプローチ: ディストリビューションカーネル](/wiki/Handbook:PPC64/Installation/Kernel/ja#Distribution_kernels "Handbook:PPC64/Installation/Kernel/ja")[ディストリビューションカーネル](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") は、Linux カーネル、関連するモジュール、および (必須ではありませんがデフォルトでは有効化されている) initramfs ファイルを、設定、自動でビルド、インストールするために利用されます。将来のカーネル更新はパッケージマネージャを介して扱われるため、他のシステムパッケージとまったく同様に完全に自動で行われます。カスタマイズが必要な場合は [カスタムのカーネルコンフィグファイルを提供する](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel") ことも可能です。これが最も簡単なプロセスで、すぐ動作するものが手に入りシステム管理者による関与を最小にできるため、新規の Gentoo ユーザには完璧です。[手動アプローチ](/wiki/Handbook:PPC64/Installation/Kernel/ja#Alternative:_Manual_configuration "Handbook:PPC64/Installation/Kernel/ja")新しいカーネルのソースがシステムパッケージマネージャを通じてインストールされます。カーネルは eselect kernel と無数の make コマンドを利用して、手動で設定、ビルド、インストールされます。将来のカーネル更新はカーネルファイルの設定、ビルド、インストールの手動プロセスを繰り返して行います。これが最も込み入ったプロセスですが、カーネル更新プロセスに関して最大限の制御を行えます。

すべてのディストリビューションが構築されるその中心にあるのが Linux カーネルです。カーネルレイヤーはユーザのプログラムとハードウェアの間に存在します。ハンドブックではカーネルソースについていくつかの可能な選択肢を提供しますが、より詳しい説明付きで、より完全なカーネルソースのリストは、 [カーネルパッケージのページ](/wiki/Kernel/Packages/ja "Kernel/Packages/ja") で見ることができます。

**ヒント**

/boot または [EFI システムパーティション](/wiki/EFI_System_Partition/ja "EFI System Partition/ja") へのカーネルイメージのコピー、 [initramfs](/wiki/Initramfs "Initramfs") や [Unified カーネルイメージ](/wiki/Unified_Kernel_Image "Unified Kernel Image") の生成、ブートローダの設定の更新など、カーネルインストールのタスクは、 [installkernel](/wiki/Installkernel/ja "Installkernel/ja") で自動化することができます。先に進める前に、[sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) をインストールして設定するとよいかもしれません。さらなる情報については [下のカーネルインストールの節](/wiki/Handbook:PPC64/Installation/Kernel/ja#Kernel_installation "Handbook:PPC64/Installation/Kernel/ja") を参照してください。

### ディストリビューションカーネル

_[ディストリビューションカーネル](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ は、カーネルを展開、構成設定、コンパイル、インストールする完全なプロセスをカバーする ebuild です。この手法を利用する最大の利点は、@world アップグレードの一部として、パッケージマネージャによってカーネルが新しいバージョンに更新されることです。これは emerge コマンドを実行する以外にユーザの関与を必要としません。ディストリビューションカーネルはデフォルトでは、大部分のハードウェアをサポートするように構成されますが、カスタマイズするための 2 つの機構が提供されています: savedconfig とコンフィグスニペットです。 [設定についてのさらなる詳細](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel") はプロジェクトページを参照してください。

##### 省略可能: 署名付きカーネルモジュール

ビルド済みディストリビューションカーネル ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) 内のカーネルモジュールは既に署名されています。ソースからビルドしたカーネルのモジュールに署名するためには、 [/etc/portage/make.conf](/wiki//etc/portage/make.conf/ja "/etc/portage/make.conf/ja") で [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグを有効化し、署名のためにどの鍵を使用するかを任意で指定してください:

ファイル **`/etc/portage/make.conf`** **モジュールの署名を有効化する**

```
USE="modules-sign"

# 任意で、カスタム署名鍵を使用するために。
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # MODULES_SIGN_KEY が証明書を含まない場合のみ必須です。
MODULES_SIGN_HASH="sha512" # デフォルトは sha512 です。

```

MODULES\_SIGN\_KEY が指定されない場合は、カーネルビルドシステムが鍵を生成し、鍵は /usr/src/linux-x.y.z/certs に保存されるでしょう。各カーネルリリースに対して同じ鍵が使用されることを確実にするためには、手動で鍵を生成することを推奨します。鍵は以下で生成することができます:

`root #` `openssl req -new -noenc -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

**メモ**

MODULES\_SIGN\_KEY と MODULES\_SIGN\_CERT は異なるファイルにすることもできます。この例では、OpenSSL によって生成される pem ファイルには鍵とそれに伴う証明書の両方を含んでいるため、両変数は同じ値に設定されています。

OpenSSL が鍵の生成についてのいくつか質問をしてくるでしょうが、できる限り詳細にこれらの質問に答えることを推奨します。

安全な場所に鍵を保存してください。最低限の措置として、鍵は root ユーザによる読み取りのみ可能にすべきです。このことを以下のようにして確認してください:

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

これが上と違ったものを出力する場合は、以下で権限を修正してください:

`root #` `chown root:root kernel_key.pem`

`root #` `chmod 400 kernel_key.pem`

#### ディストリビューションカーネルをインストールする

Gentoo パッチが当てられたカーネルをソースからビルドするには、次をタイプしてください:

`root #` `emerge --ask sys-kernel/gentoo-kernel`

システムの管理者として、カーネルのソースをローカルでコンパイルするのを避けたい場合は、代わりにコンパイル済みのカーネルイメージを使用することができます:

`root #` `emerge --ask sys-kernel/gentoo-kernel-bin`

**重要**

[sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) および [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) のような _[ディストリビューションカーネル](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ は、デフォルトでは、 [initramfs](/wiki/Handbook:PPC64/Installation/Kernel/ja#Initramfs "Handbook:PPC64/Installation/Kernel/ja") とともにインストールされることを想定しています。カーネルをインストールするために emerge を実行する前に、 [installkernel の節](/wiki/Handbook:PPC64/Installation/Kernel/ja#Initramfs "Handbook:PPC64/Installation/Kernel/ja") の記載に従って、initramfs ジェネレータ ( [Dracut](/wiki/Dracut "Dracut") など) を利用するように [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) が設定されていることを確認するべきです。

#### アップグレードと後処理

一度カーネルがインストールされたら、パッケージマネージャが自動的にカーネルを新しいバージョンに更新するでしょう。古いバージョンは、パッケージマネージャに古いパッケージを片付けるように指示するまで残ります。ディスク容量を再利用するためには、定期的に `--depclean` オプション付きで emerge を実行することで、古いパッケージを片付けることができます:

`root #` `emerge --depclean`

あるいは、古いカーネルのバージョンだけを片付けるには:

`root #` `emerge --prune sys-kernel/gentoo-kernel sys-kernel/gentoo-kernel-bin`

**ヒント**

設計として、emerge はカーネルビルドディレクトリのみを削除します。カーネルモジュールや、インストールされたカーネルイメージは削除しません。古いカーネルを完全に削除するためには、[app-admin/eclean-kernel](https://packages.gentoo.org/packages/app-admin/eclean-kernel) ツールを使用することができます。

#### インストール/アップグレード後タスク

ディストリビューションカーネルのアップグレードは、他のパッケージ (例: [sys-fs/zfs-kmod](https://packages.gentoo.org/packages/sys-fs/zfs-kmod) または [x11-drivers/nvidia-drivers](https://packages.gentoo.org/packages/x11-drivers/nvidia-drivers)) によってインストールされた外部カーネルモジュールに対して、自動的な再ビルドを発動させることができます。この自動化された挙動は [dist-kernel](https://packages.gentoo.org/useflags/dist-kernel) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグによって有効化されます。このフラグはまた、必要に応じて [initramfs](/wiki/Initramfs "Initramfs") の再生成も発動させるでしょう。

ディストリビューションカーネルを使用する場合は、このフラグを /etc/portage/make.conf から有効化しておくことが強く推奨されます:

ファイル **`/etc/portage/make.conf`** **USE=dist-kernel を有効化する**

```
USE="dist-kernel"

```

##### 手動で initramfs または Unified カーネルイメージを再ビルドする

必要であれば、カーネルアップグレード後に、以下を実行して手動で再ビルドを実行してください:

`root #` `emerge --ask @module-rebuild`

カーネルモジュール (ZFS 等) のうちいずれかが初期ブートで必要であれば、続けて次を利用して initramfs を再ビルドしてください:

`root #` `emerge --config sys-kernel/gentoo-kernel
`

`root #` `emerge --config sys-kernel/gentoo-kernel-bin
`

ディストリビューションカーネルを正常にインストールできたら、次の節に進むときです: 次は [システムの設定](/wiki/Handbook:PPC64/Installation/System/ja "Handbook:PPC64/Installation/System/ja") です。

### 別の方法: マニュアル設定

### カーネルソースのインストール

ppc64 ベースのシステムにカーネルを手動でインストールしてコンパイルする場合には、Gentoo は[sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources) パッケージを推奨しています。

適切なカーネルソースを選択して、emerge でインストールします:

`root #` `emerge --ask sys-kernel/gentoo-sources`

このコマンドはカーネルソースを /usr/src/ の下に、カーネルバージョン毎のパスを分けてインストールします。選択されたカーネルソースパッケージに対して [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグが有効化されていなければ、シンボリックリンクは自動で作成されません。

現在実行しているカーネルに対応するソースを指すように、/usr/src/linux シンボリックリンクを維持することは慣例となっています。しかし、このシンボリックリンクはデフォルトでは作成されないでしょう。シンボリックリンクを作成する簡単な方法は、eselect の kernel モジュールを利用することです。

シンボリックリンクの目的と、それを管理する方法についてのさらなる情報は、 [Kernel/Upgrade](/wiki/Kernel/Upgrade/ja "Kernel/Upgrade/ja") を参照してください。

まず、インストールされているカーネルを一覧表示します:

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.18.8-gentoo

```

linux シンボリックリンクを作成するには、次を使用してください:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    12 Oct 13 11:04 /usr/src/linux -> linux-6.18.8-gentoo

```

カーネルのマニュアル設定は一般的に、システム管理者がしなければならない最も難しい手続きのひとつと見なされています。これは真実ではありません。カーネルを数回設定してみれば、それが難しいと言われていたことなど忘れてしまうでしょう！ Gentoo ユーザが手動でカーネルシステムを管理するための方法は 2 通りあり、その両方を以下に示します:

**重要**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### Modprobed-db プロセス

カーネルを管理するための非常に簡単な方法のひとつは、まず [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) をインストールし、システムが必要とするものについての情報を収集するために [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) を使用するという方法です。modprobed-db は crontab を介してシステムを監視し、システムが使用されている間のすべてのデバイスのすべてのモジュールを追加し、ユーザが必要とするものすべてがサポートされていることを確実にするツールです。例えば、インストール後に Xbox コントローラが追加されたら modprobed-db は、カーネルが次回再ビルドされるときにビルドされる対象としてそのモジュールを追加するでしょう。この話題についてさらに詳しくは、 [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") の記事で読むことができます。

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:PPC64/Installation/Kernel#Distribution_kernels.2Fja "Handbook:PPC64/Installation/Kernel").

Next, install [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db):

`root #` `emerge --ask sys-kernel/modprobed-db`

Please watch out for further steps related to modprobed-db in the Handbook.

More on this topic can be found in the [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") article.
}}

#### 手動プロセス

この方法では、外部ツールの利用を望む限り最小限にすることによって、カーネルのビルド方法を完全に管理することができます。人によってはこの方法のことを、わざわざ難しくやっていると思うかもしれません。

しかし、この選択肢を取る場合でも一つだけ真実があります。カーネルを手動で設定する時、ハードウェア情報を知ることはとても役に立ちます。ほとんどの情報は、lspci コマンドを含む [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils) をインストールすることで得られます。

`root #` `emerge --ask sys-apps/pciutils`

**メモ**

chroot環境では、lspciが出力していると思われる( _pcilib: cannot open /sys/bus/pci/devices_ のような)pcilibの警告は、無視しても構いません。

システム情報を得るための別の方法は、lsmodを使ってインストールCDが使っているカーネルモジュールを把握することです。その情報は何を有効にすべきかとてもよいヒントを与えてくれるでしょう。

では、カーネルソースがあるディレクトリに移動しましょう。

`root #` `cd /usr/src/linux
`

**ヒント**

カーネルに対して利用可能な make の引数の一覧を確認するには、 `make help` を実行してください。

カーネルは現在 installcd で使用されているモジュールを自動検出することができ、これはユーザーが独自の設定をするにあたってよい出発点になるでしょう。呼び出すには以下のようにします:

`root #` `make localmodconfig`

それでは、nconfig を使って設定しましょう:

`root #` `make nconfig`

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:PPC64/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fja "Handbook:PPC64/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:PPC64/Installation/Kernel#Compiling_and_installing.2Fja "Handbook:PPC64/Installation/Kernel")

###### 必要なオプションを有効化する

#### 省略可能: 署名済みカーネルモジュール

自動的にカーネルモジュールに署名するためには、CONFIG\_MODULE\_SIG\_ALL を有効化してください:

カーネル **カーネルモジュールに署名する CONFIG\_MODULE\_SIG\_ALL**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

お望みであれば任意でハッシュアルゴリズムを変更してください。

すべてのモジュールが有効な署名で署名されていることを強制するためには、CONFIG\_MODULE\_SIG\_FORCE も有効化してください:

カーネル **署名済みカーネルモジュールを強制する CONFIG\_MODULE\_SIG\_FORCE**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Require modules to be validly signed
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

カスタム鍵を使用するためには、CONFIG\_MODULE\_SIG\_KEY にこの鍵の場所を指定してください。指定されていない場合は、カーネルビルドシステムが鍵を生成するでしょう。それに任せずに手動で鍵を生成することが推奨されます。これは、以下で行えます:

`root #` `openssl req -new -nodes -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

OpenSSL は鍵を生成するユーザについて、いくつかの質問をしてくるでしょう。これらの質問にできる限り詳細に答えることが推奨されます。

鍵を安全な場所に保存してください。少なくとも、鍵は root ユーザによってのみ読み込み可能であるべきです。以下でこれを検証してください:

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

もしこれが上と何か違うものを出力する場合は、パーミッションを以下で訂正してください:

`root #` `chown root:root kernel_key.pem
`

`root #` `chmod 400 kernel_key.pem
`

カーネル **署名鍵を指定する CONFIG\_MODULE\_SIG\_KEY**

```
-*- Cryptographic API  --->
  Certificates for signature checking  --->
    (/path/to/kernel_key.pem) File name or PKCS#11 URI of module signing key

```

`linux-mod-r1.eclass` を利用する他のパッケージによってインストールされた外部カーネルモジュールにも署名するには、グローバルに [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグを有効化してください:

ファイル **`/etc/portage/make.conf`** **モジュールの署名を有効化する**

```
USE="modules-sign"

# 任意で、カスタム署名鍵を使用する場合。
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # 鍵 MODULES_SIGN_KEY が証明書を含まない場合のみ必須です
MODULES_SIGN_HASH="sha512" # デフォルトは sha512 です

```

**メモ**

MODULES\_SIGN\_KEY と MODULES\_SIGN\_CERT は異なるファイルを指していても構いません。この例では、OpenSSL によって生成される pem ファイルには鍵とそれに伴う証明書の両方を含んでいるため、両変数は同じ値に設定されています。

#### コンパイルおよびインストール

カーネルのコンフィグレーションが完了し、コンパイルとインストールをする時がきました。コンフィグレーションを終了させ、コンパイル作業を開始しましょう:

`root #` `make && make modules_install`

**メモ**

make -j N とすることで、ビルドを並行処理させることができます (N には、並行処理を許可するビルドプロセスの _数_ を整数で指定します)。/etc/portage/make.conf の説明中の、 MAKEOPTS 変数についてと同様です。

When the kernel has finished compiling, copy the kernel image to /boot/. This is handled by the make install command:

`root #` `make install`

This command will copy the kernel image to /boot. If [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is installed it will call /sbin/installkernel instead and delegate the kernel installation. Instead of simply copying the kernel to /boot, [Installkernel](/wiki/Installkernel/ja "Installkernel/ja") installs each kernel with its version number in the file name. Additionally, installkernel provides a framework for automatically accomplishing various tasks relating to kernel installation, such as: generating an [initramfs](/wiki/Initramfs "Initramfs"), building an [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), and updating the [bootloader](/wiki/Bootloader/ja "Bootloader/ja") configuration.

では、 [システムの設定](/wiki/Handbook:PPC64/Installation/System/ja "Handbook:PPC64/Installation/System/ja") に進み、インストールを続けましょう。

[← Gentooベースシステムのインストール](/wiki/Handbook:PPC64/Installation/Base/ja "Previous part") [Home](/wiki/Handbook:PPC64/ja "Handbook:PPC64/ja") [システムの設定 →](/wiki/Handbook:PPC64/Installation/System/ja "Next part")