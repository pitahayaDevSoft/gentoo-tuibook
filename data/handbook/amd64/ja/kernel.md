# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Kernel/de "Handbook:AMD64/Installation/Kernel/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Kernel "Handbook:AMD64/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Kernel/es "Handbook:AMD64/Installation/Kernel/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Kernel/fr "Handbook:AMD64/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Kernel/it "Handbook:AMD64/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Kernel/hu "Handbook:AMD64/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Kernel/pl "Handbook:AMD64/Installation/Kernel/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Kernel/pt-br "Handbook:AMD64/Installation/Kernel/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Kernel/cs "Handbook:AMD64/Installation/Kernel/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Kernel/ru "Handbook:AMD64/Installation/Kernel/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Kernel/ta "Handbook:AMD64/Installation/Kernel/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Kernel/zh-cn "Handbook:AMD64/Installation/Kernel/zh-cn (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:AMD64/Installation/Kernel/ko "Handbook:AMD64/Installation/Kernel/ko (100% translated)")

[AMD64 ハンドブック](/wiki/Handbook:AMD64 "Handbook:AMD64")[インストール](/wiki/Handbook:AMD64/Full/Installation "Handbook:AMD64/Full/Installation")[インストールについて](/wiki/Handbook:AMD64/Installation/About/ja "Handbook:AMD64/Installation/About/ja")[メディアの選択](/wiki/Handbook:AMD64/Installation/Media/ja "Handbook:AMD64/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:AMD64/Installation/Networking/ja "Handbook:AMD64/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:AMD64/Installation/Disks/ja "Handbook:AMD64/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:AMD64/Installation/Stage/ja "Handbook:AMD64/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:AMD64/Installation/Base/ja "Handbook:AMD64/Installation/Base/ja")カーネルの設定[システムの設定](/wiki/Handbook:AMD64/Installation/System/ja "Handbook:AMD64/Installation/System/ja")[ツールのインストール](/wiki/Handbook:AMD64/Installation/Tools/ja "Handbook:AMD64/Installation/Tools/ja")[ブートローダの設定](/wiki/Handbook:AMD64/Installation/Bootloader/ja "Handbook:AMD64/Installation/Bootloader/ja")[締めくくり](/wiki/Handbook:AMD64/Installation/Finalizing/ja "Handbook:AMD64/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:AMD64/Full/Working "Handbook:AMD64/Full/Working")[Portage について](/wiki/Handbook:AMD64/Working/Portage/ja "Handbook:AMD64/Working/Portage/ja")[USE フラグ](/wiki/Handbook:AMD64/Working/USE/ja "Handbook:AMD64/Working/USE/ja")[Portage の機能](/wiki/Handbook:AMD64/Working/Features/ja "Handbook:AMD64/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:AMD64/Working/Initscripts/ja "Handbook:AMD64/Working/Initscripts/ja")[環境変数](/wiki/Handbook:AMD64/Working/EnvVar/ja "Handbook:AMD64/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:AMD64/Full/Portage "Handbook:AMD64/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:AMD64/Portage/Files/ja "Handbook:AMD64/Portage/Files/ja")[変数](/wiki/Handbook:AMD64/Portage/Variables/ja "Handbook:AMD64/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:AMD64/Portage/Branches/ja "Handbook:AMD64/Portage/Branches/ja")[追加ツール](/wiki/Handbook:AMD64/Portage/Tools/ja "Handbook:AMD64/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:AMD64/Portage/CustomTree/ja "Handbook:AMD64/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:AMD64/Portage/Advanced/ja "Handbook:AMD64/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:AMD64/Full/Networking "Handbook:AMD64/Full/Networking")[はじめに](/wiki/Handbook:AMD64/Networking/Introduction/ja "Handbook:AMD64/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:AMD64/Networking/Advanced/ja "Handbook:AMD64/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:AMD64/Networking/Modular/ja "Handbook:AMD64/Networking/Modular/ja")[無線](/wiki/Handbook:AMD64/Networking/Wireless/ja "Handbook:AMD64/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:AMD64/Networking/Extending/ja "Handbook:AMD64/Networking/Extending/ja")[動的な管理](/wiki/Handbook:AMD64/Networking/Dynamic/ja "Handbook:AMD64/Networking/Dynamic/ja")

## Contents

- [1ファームウェアとマイクロコードのインストール](#.E3.83.95.E3.82.A1.E3.83.BC.E3.83.A0.E3.82.A6.E3.82.A7.E3.82.A2.E3.81.A8.E3.83.9E.E3.82.A4.E3.82.AF.E3.83.AD.E3.82.B3.E3.83.BC.E3.83.89.E3.81.AE.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB)
  - [1.1ファームウェア](#.E3.83.95.E3.82.A1.E3.83.BC.E3.83.A0.E3.82.A6.E3.82.A7.E3.82.A2)
    - [1.1.1提案: Linux ファームウェア](#.E6.8F.90.E6.A1.88:_Linux_.E3.83.95.E3.82.A1.E3.83.BC.E3.83.A0.E3.82.A6.E3.82.A7.E3.82.A2)
      - [1.1.1.1ファームウェアのロード](#.E3.83.95.E3.82.A1.E3.83.BC.E3.83.A0.E3.82.A6.E3.82.A7.E3.82.A2.E3.81.AE.E3.83.AD.E3.83.BC.E3.83.89)
    - [1.1.2SOF Firmware](#SOF_Firmware)
  - [1.2提案: マイクロコード](#.E6.8F.90.E6.A1.88:_.E3.83.9E.E3.82.A4.E3.82.AF.E3.83.AD.E3.82.B3.E3.83.BC.E3.83.89)
- [2sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [2.1ブートローダ](#.E3.83.96.E3.83.BC.E3.83.88.E3.83.AD.E3.83.BC.E3.83.80)
    - [2.1.1GRUB](#GRUB)
    - [2.1.2systemd-boot](#systemd-boot)
    - [2.1.3EFI スタブ](#EFI_.E3.82.B9.E3.82.BF.E3.83.96)
    - [2.1.4伝統的なレイアウト、他のブートローダ (例: (e)lilo、syslinux 等)](#.E4.BC.9D.E7.B5.B1.E7.9A.84.E3.81.AA.E3.83.AC.E3.82.A4.E3.82.A2.E3.82.A6.E3.83.88.E3.80.81.E4.BB.96.E3.81.AE.E3.83.96.E3.83.BC.E3.83.88.E3.83.AD.E3.83.BC.E3.83.80_.28.E4.BE.8B:_.28e.29lilo.E3.80.81syslinux_.E7.AD.89.29)
  - [2.2Initramfs](#Initramfs)
    - [2.2.1Chroot detection](#Chroot_detection)
  - [2.3省略可能: Unified カーネルイメージ](#.E7.9C.81.E7.95.A5.E5.8F.AF.E8.83.BD:_Unified_.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.A4.E3.83.A1.E3.83.BC.E3.82.B8)
    - [2.3.1ジェネリック Unified カーネルイメージ (systemd のみ)](#.E3.82.B8.E3.82.A7.E3.83.8D.E3.83.AA.E3.83.83.E3.82.AF_Unified_.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.A4.E3.83.A1.E3.83.BC.E3.82.B8_.28systemd_.E3.81.AE.E3.81.BF.29)
    - [2.3.2セキュアブート](#.E3.82.BB.E3.82.AD.E3.83.A5.E3.82.A2.E3.83.96.E3.83.BC.E3.83.88)
- [3カーネルのコンフィギュレーションとコンパイル](#.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.81.AE.E3.82.B3.E3.83.B3.E3.83.95.E3.82.A3.E3.82.AE.E3.83.A5.E3.83.AC.E3.83.BC.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.A8.E3.82.B3.E3.83.B3.E3.83.91.E3.82.A4.E3.83.AB)
  - [3.1ディストリビューションカーネル](#.E3.83.87.E3.82.A3.E3.82.B9.E3.83.88.E3.83.AA.E3.83.93.E3.83.A5.E3.83.BC.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB)
    - [3.1.1省略可能: 署名付きカーネルモジュール](#.E7.9C.81.E7.95.A5.E5.8F.AF.E8.83.BD:_.E7.BD.B2.E5.90.8D.E4.BB.98.E3.81.8D.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.83.A2.E3.82.B8.E3.83.A5.E3.83.BC.E3.83.AB)
    - [3.1.2省略可能: カーネルイメージに署名する (セキュアブート)](#.E7.9C.81.E7.95.A5.E5.8F.AF.E8.83.BD:_.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.A4.E3.83.A1.E3.83.BC.E3.82.B8.E3.81.AB.E7.BD.B2.E5.90.8D.E3.81.99.E3.82.8B_.28.E3.82.BB.E3.82.AD.E3.83.A5.E3.82.A2.E3.83.96.E3.83.BC.E3.83.88.29)
    - [3.1.3ディストリビューションカーネルをインストールする](#.E3.83.87.E3.82.A3.E3.82.B9.E3.83.88.E3.83.AA.E3.83.93.E3.83.A5.E3.83.BC.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.92.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB.E3.81.99.E3.82.8B)
    - [3.1.4アップグレードと後処理](#.E3.82.A2.E3.83.83.E3.83.97.E3.82.B0.E3.83.AC.E3.83.BC.E3.83.89.E3.81.A8.E5.BE.8C.E5.87.A6.E7.90.86)
    - [3.1.5インストール/アップグレード後タスク](#.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB.2F.E3.82.A2.E3.83.83.E3.83.97.E3.82.B0.E3.83.AC.E3.83.BC.E3.83.89.E5.BE.8C.E3.82.BF.E3.82.B9.E3.82.AF)
      - [3.1.5.1手動で initramfs または Unified カーネルイメージを再ビルドする](#.E6.89.8B.E5.8B.95.E3.81.A7_initramfs_.E3.81.BE.E3.81.9F.E3.81.AF_Unified_.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.A4.E3.83.A1.E3.83.BC.E3.82.B8.E3.82.92.E5.86.8D.E3.83.93.E3.83.AB.E3.83.89.E3.81.99.E3.82.8B)
  - [3.2別の方法: マニュアル設定](#.E5.88.A5.E3.81.AE.E6.96.B9.E6.B3.95:_.E3.83.9E.E3.83.8B.E3.83.A5.E3.82.A2.E3.83.AB.E8.A8.AD.E5.AE.9A)
  - [3.3カーネルソースのインストール](#.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.BD.E3.83.BC.E3.82.B9.E3.81.AE.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB)
    - [3.3.1Modprobed-db プロセス](#Modprobed-db_.E3.83.97.E3.83.AD.E3.82.BB.E3.82.B9)
    - [3.3.2手動プロセス](#.E6.89.8B.E5.8B.95.E3.83.97.E3.83.AD.E3.82.BB.E3.82.B9)
      - [3.3.2.1必要なオプションを有効化する](#.E5.BF.85.E8.A6.81.E3.81.AA.E3.82.AA.E3.83.97.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E6.9C.89.E5.8A.B9.E5.8C.96.E3.81.99.E3.82.8B)
    - [3.3.3省略可能: 署名済みカーネルモジュール](#.E7.9C.81.E7.95.A5.E5.8F.AF.E8.83.BD:_.E7.BD.B2.E5.90.8D.E6.B8.88.E3.81.BF.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.83.A2.E3.82.B8.E3.83.A5.E3.83.BC.E3.83.AB)
    - [3.3.4省略可能: カーネルイメージに署名する (セキュアブート)](#.E7.9C.81.E7.95.A5.E5.8F.AF.E8.83.BD:_.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.A4.E3.83.A1.E3.83.BC.E3.82.B8.E3.81.AB.E7.BD.B2.E5.90.8D.E3.81.99.E3.82.8B_.28.E3.82.BB.E3.82.AD.E3.83.A5.E3.82.A2.E3.83.96.E3.83.BC.E3.83.88.29_2)
    - [3.3.5コンパイルおよびインストール](#.E3.82.B3.E3.83.B3.E3.83.91.E3.82.A4.E3.83.AB.E3.81.8A.E3.82.88.E3.81.B3.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB)

## ファームウェアとマイクロコードのインストール

### ファームウェア

#### 提案: Linux ファームウェア

多くのシステムでは、特定のハードウェアが正しく機能するための非 FOSS なファームウェアが必要です。[sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) パッケージは、すべてではありませんが、多くのデバイスのためのファームウェアを含んでいます。

**ヒント**

ほとんどの無線カードと GPU は、機能するためにファームウェアを必要とします。

`root #` `emerge --ask sys-kernel/linux-firmware`

**メモ**

一部のファームウェアパッケージのインストールには、関連するファームウェアライセンスを受諾する必要があることがよくあります。必要であれば、ライセンスの受諾についてはハンドブックの [ライセンスの取り扱いの節](/wiki/Handbook:AMD64/Working/Portage/ja#Licenses "Handbook:AMD64/Working/Portage/ja") を確認してください。

##### ファームウェアのロード

ファームウェアファイルは、典型的には、それに関連付けられたカーネルモジュールがロードされたときにロードされます。これはつまり、カーネルモジュールが _M_ ではなく _Y_ に設定されている場合、ファームウェアは **CONFIG\_EXTRA\_FIRMWARE** を使用してカーネルに組み込まれていなくてはならないということです。ほとんどの場合、ファームウェアを必要とするモジュールの組み込みは、ロードを複雑にしたり、失敗の原因になったりすることがあります。

#### SOF Firmware

Sound Open Firmware (SOF) is a new open source audio driver meant to replace the proprietary Smart Sound Technology (SST) audio driver from Intel. 10th gen+ and Apollo Lake (Atom E3900, Celeron N3350, and Pentium N4200) Intel CPUs require this firmware for certain features and certain AMD APUs also have support for this firmware. SOF's supported platforms matrix can be found [here](https://thesofproject.github.io/latest/platforms/index.html) for more information.

`root #` `emerge --ask sys-firmware/sof-firmware`

### 提案: マイクロコード

個別のグラフィックスハードウェアやネットワークインターフェースに加えて、CPU もまたファームウェアアップデートを必要とすることがあります。こうしたファームウェアは典型的には _マイクロコード_ と呼ばれます。新しいリビジョンのマイクロコードは、動作の不安定さ、セキュリティ上の懸念、その他の CPU ハードウェアのさまざまなバグに対するパッチとして、必要になることがあります。

AMD CPU に対するマイクロコードアップデートは、先述の [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) パッケージとともに配布されます。Intel CPU に対するマイクロコードは [sys-firmware/intel-microcode](https://packages.gentoo.org/packages/sys-firmware/intel-microcode) パッケージ内で見つかりますので、これを個別にインストールする必要があります。マイクロコードアップデートを適用する方法についてのさらなる情報は、 [マイクロコードの記事](/wiki/Microcode "Microcode") を確認してください。

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

#### systemd-boot

ブートローダとして [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") (旧 gummiboot) を使用している場合は、systemd の kernel-install を使用しなくてはなりません。そのため、[sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) に対して [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") および [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグが有効化されていることを確認して、systemd-boot のための関連するパッケージをインストールしてください。

ファイル **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot kernel-install
sys-kernel/installkernel systemd systemd-boot

```

新しいカーネルのために使用するカーネルコマンドラインは /etc/kernel/cmdline および /etc/cmdline 内で指定すべきです。両ファイルは共に更新する必要があるため、ここでシンボリックリンクにしておくといいでしょう。

ファイル **`/etc/kernel/cmdline`**

```
quiet splash

```

`root #` `mkdir -p /etc/cmdline.d`

`root #` `ln -s /etc/kernel/cmdline /etc/cmdline.d/00-installkernel.conf`

`root #` `emerge --ask sys-kernel/installkernel`

**メモ**

systemd-boot requires kernels to be installed to /efi.

#### EFI スタブ

UEFI ベースのコンピュータシステムでは、厳密には、カーネルをブートするために二次ブートローダは必要ありません。二次ブートローダは、ブートプロセス中の UEFI ファームウェアの機能性を _拡張_ するために存在しています。とはいえ、典型的には二次ブートローダを使用するほうがより簡単で、ブート時にカーネルパラメータを手っ取り早く変更するための柔軟なアプローチを提供しているため、より堅牢です。 UEFI 実装は、ベンダやモデルによって大きく異なり、また与えられたファームウェアが UEFI 仕様に従っている保証はないことにも注意してください。したがって、EFI スタブブートはすべての UEFI ベースのシステムで機能する保証はありません。

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi/EFI/Gentoo`

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

### 省略可能: Unified カーネルイメージ

[Unified カーネルイメージ](/wiki/Unified_Kernel_Image "Unified Kernel Image") (UKI) は、まず何よりもカーネル、そして initramfs、それとカーネルコマンドラインを単一の実行可能形式に合成したものです。カーネルコマンドラインは unified カーネルイメージ内に埋め込まれるので、unified カーネルイメージを生成する前に指定するべきです (下を参照してください)。セキュアブートを有効化してブートする場合、ブートローダまたはファームウェアによってブート時に提供されるあらゆるカーネルコマンドライン引数は無視されることに注意してください。

unified カーネルイメージはスタブローダを必要とします。現時点で唯一利用可能なのは systemd-stub です。これを有効化するには:

systemd システムでは:

ファイル **`/etc/portage/package.use/uki`**

```
sys-apps/systemd boot

```

`root #` `emerge --ask sys-apps/systemd`

OpenRC システムでは:

ファイル **`/etc/portage/package.use/uki`**

```
sys-apps/systemd-utils boot kernel-install

```

`root #` `emerge --ask sys-apps/systemd-utils`

[Installkernel](/wiki/Installkernel/ja "Installkernel/ja") は [dracut](/wiki/Unified_kernel_image#dracut "Unified kernel image") または [ukify](/wiki/Unified_kernel_image#ukify "Unified kernel image") のいずれかを使用して、自動的に unified カーネルイメージを生成することができます。これは、それぞれに対応するフラグと、[uki](https://packages.gentoo.org/useflags/uki) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグを有効化することで行えます。

dracut を使用する場合は:

ファイル **`/etc/portage/package.use/uki`**

```
sys-kernel/installkernel dracut uki

```

ファイル **`/etc/dracut.conf.d/uki.conf`**

```
uefi="yes"
kernel_cmdline="some-kernel-command-line-arguments"

```

`root #` `emerge --ask sys-kernel/installkernel`

ukify を使用する場合は:

ファイル **`/etc/portage/package.use/uki`**

```
sys-apps/systemd boot ukify                         # systemd システムの場合
sys-apps/systemd-utils kernel-install boot ukify    # OpenRC システムの場合
sys-kernel/installkernel dracut ukify uki

```

ファイル **`/etc/kernel/cmdline`**

```
some-kernel-command-line-arguments

```

`root #` `emerge --ask sys-kernel/installkernel`

dracut は initramfs と unified カーネルイメージの両方を生成することができる一方で、ukify は後者しか生成できないので、initramfs は dracut を使用して個別に生成しなくてはならないことに注意してください。

**重要**

上の設定例 (Dracut および ukify の両方) では、unified カーネルイメージがルートパーティションを見つけることができるようにするために、少なくとも適切な root= カーネルコマンドラインパラメータを指定することが重要です。これは Discoverable Partitions Specification (DPS) に従う systemd ベースのシステムでは必要ありません。この場合は、埋め込まれた initramfs がルートパーティションを動的に見つけることができるでしょう。

#### ジェネリック Unified カーネルイメージ (systemd のみ)

ビルド済みの [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) は任意で、ほとんどの systemd ベースのシステムをブートすることができるジェネリックな initramfs を含む、ビルド済みのジェネリック unified カーネルイメージをインストールすることができます。これは、[generic-uki](https://packages.gentoo.org/useflags/generic-uki) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグを有効化して、 [installkernel](/wiki/Installkernel/ja "Installkernel/ja") をカスタムの initramfs または unified カーネルイメージを生成しないように設定することで、インストールすることができます:

ファイル **`/etc/portage/package.use/uki`**

```
sys-kernel/gentoo-kernel-bin generic-uki
sys-kernel/installkernel -dracut -ukify -ugrd uki

```

#### セキュアブート

**警告**

この節に従い、かつ自身のカーネルを手動でコンパイルする場合、 [カーネルに署名する](/wiki/Handbook:AMD64/Installation/Kernel/ja#Optional:_Signing_the_kernel_image_.28Secure_Boot.29 "Handbook:AMD64/Installation/Kernel/ja") で触れられている手順に従うようにしてください。

[sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) によって任意に配布されているジェネリック Unified カーネルイメージは事前署名済みです。ローカルで生成された unified カーネルイメージに署名する方法は dracut または ukify のどちらを使用するかで異なります。鍵と証明書の場所は /etc/portage/make.conf 内で指定された SECUREBOOT\_SIGN\_KEY および SECUREBOOT\_SIGN\_CERT と同一であるべきということに注意してください。

dracut を使用する場合は:

ファイル **`/etc/dracut.conf.d/uki.conf`**

```
uefi="yes"
kernel_cmdline="some-kernel-command-line-arguments"
uefi_secureboot_key="/path/to/kernel_key.pem"
uefi_secureboot_cert="/path/to/kernel_key.pem"

```

ukify を使用する場合は:

ファイル **`/etc/kernel/uki.conf`**

```
[UKI]
SecureBootPrivateKey=/path/to/kernel_key.pem
SecureBootCertificate=/path/to/kernel_key.pem

```

## カーネルのコンフィギュレーションとコンパイル

**ヒント**

最初のブートでは dist-kernel を使用するのが懸命な選択かもしれません。システムの問題とカーネルコンフィグの問題を除外するための、とても簡潔な手法を提供しているからです。機能すると分かっているカーネルをフォールバックとして常に持っておくことは、デバッグの効率を高め、更新時にシステムがブートできなくなるかもという不安を緩和させるでしょう。

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**重要**

以下のサブ節のうち一つだけが必要です。どれを使用すべきか分からない場合は、とりあえずは最初に記載されているもので進めてください。もし必要なら、後でいつでも切り換えることができます。

簡単なものから込み入ったものへ、順に並べると:

[完全自動アプローチ: ディストリビューションカーネル](/wiki/Handbook:AMD64/Installation/Kernel/ja#Distribution_kernels "Handbook:AMD64/Installation/Kernel/ja")[ディストリビューションカーネル](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") は、Linux カーネル、関連するモジュール、および (必須ではありませんがデフォルトでは有効化されている) initramfs ファイルを、設定、自動でビルド、インストールするために利用されます。将来のカーネル更新はパッケージマネージャを介して扱われるため、他のシステムパッケージとまったく同様に完全に自動で行われます。カスタマイズが必要な場合は [カスタムのカーネルコンフィグファイルを提供する](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel") ことも可能です。これが最も簡単なプロセスで、すぐ動作するものが手に入りシステム管理者による関与を最小にできるため、新規の Gentoo ユーザには完璧です。[手動アプローチ](/wiki/Handbook:AMD64/Installation/Kernel/ja#Alternative:_Manual_configuration "Handbook:AMD64/Installation/Kernel/ja")新しいカーネルのソースがシステムパッケージマネージャを通じてインストールされます。カーネルは eselect kernel と無数の make コマンドを利用して、手動で設定、ビルド、インストールされます。将来のカーネル更新はカーネルファイルの設定、ビルド、インストールの手動プロセスを繰り返して行います。これが最も込み入ったプロセスですが、カーネル更新プロセスに関して最大限の制御を行えます。

すべてのディストリビューションが構築されるその中心にあるのが Linux カーネルです。カーネルレイヤーはユーザのプログラムとハードウェアの間に存在します。ハンドブックではカーネルソースについていくつかの可能な選択肢を提供しますが、より詳しい説明付きで、より完全なカーネルソースのリストは、 [カーネルパッケージのページ](/wiki/Kernel/Packages/ja "Kernel/Packages/ja") で見ることができます。

**ヒント**

/boot または [EFI システムパーティション](/wiki/EFI_System_Partition/ja "EFI System Partition/ja") へのカーネルイメージのコピー、 [initramfs](/wiki/Initramfs "Initramfs") や [Unified カーネルイメージ](/wiki/Unified_Kernel_Image "Unified Kernel Image") の生成、ブートローダの設定の更新など、カーネルインストールのタスクは、 [installkernel](/wiki/Installkernel/ja "Installkernel/ja") で自動化することができます。先に進める前に、[sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) をインストールして設定するとよいかもしれません。さらなる情報については [下のカーネルインストールの節](/wiki/Handbook:AMD64/Installation/Kernel/ja#Kernel_installation "Handbook:AMD64/Installation/Kernel/ja") を参照してください。

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

##### 省略可能: カーネルイメージに署名する (セキュアブート)

ビルド済みディストリビューションカーネル ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) 内のカーネルイメージは既に [セキュアブート](/wiki/Secure_Boot "Secure Boot") 用に署名されています。ソースからビルドされたカーネルイメージに署名するには、 [/etc/portage/make.conf](/wiki//etc/portage/make.conf/ja "/etc/portage/make.conf/ja") で [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグを有効化し、署名のためにどの鍵を使用するかを任意で指定してください。セキュアブートとともに使用するためにカーネルイメージに署名するには、カーネルモジュールも署名する必要があることに注意してください。カーネルイメージとカーネルモジュールの両方に署名するために、同一の鍵を使用しても構いません:

ファイル **`/etc/portage/make.conf`** **カスタム署名鍵を有効化する**

```
USE="modules-sign secureboot"

# 任意で、カスタム署名鍵を使用するために。
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # MODULES_SIGN_KEY が証明書を含まない場合のみ必須です。
MODULES_SIGN_HASH="sha512" # デフォルトは sha512 です。

# 任意で、セキュアブートを有効化してブートするために。署名鍵は同一でも別でも構いません。
SECUREBOOT_SIGN_KEY="/path/to/kernel_key.pem"
SECUREBOOT_SIGN_CERT="/path/to/kernel_key.pem"

```

**メモ**

SECUREBOOT\_SIGN\_KEY と SECUREBOOT\_SIGN\_CERT は異なるファイルにすることもできます。この例では、OpenSSL によって生成される pem ファイルには鍵とそれに伴う証明書の両方を含んでいるため、両変数は同じ値に設定されています。

**メモ**

この例では、カーネルイメージに署名するために、モジュールに署名するために生成されたのと同じ鍵が使用されています。カーネルイメージに署名するために、2 個目の別の鍵を生成して使用することも可能です。前節と同じ OpenSSL コマンドをもう一度使用することができます。

新しい鍵を生成する手順については上の節を参照してください。カーネルイメージを署名するために別の鍵を使用する場合は、同じ手順を繰り返すことができます。

セキュアブートを有効化して正しくブートさせるためには、使用するブートローダにも署名し、証明書が [UEFI](/wiki/UEFI "UEFI") ファームウェアまたは [Shim](/wiki/Shim "Shim") によって許可されなくてはなりません。これはハンドブック内で後で説明します。

#### ディストリビューションカーネルをインストールする

Gentoo パッチが当てられたカーネルをソースからビルドするには、次をタイプしてください:

`root #` `emerge --ask sys-kernel/gentoo-kernel`

システムの管理者として、カーネルのソースをローカルでコンパイルするのを避けたい場合は、代わりにコンパイル済みのカーネルイメージを使用することができます:

`root #` `emerge --ask sys-kernel/gentoo-kernel-bin`

**重要**

[sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) および [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) のような _[ディストリビューションカーネル](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ は、デフォルトでは、 [initramfs](/wiki/Handbook:AMD64/Installation/Kernel/ja#Initramfs "Handbook:AMD64/Installation/Kernel/ja") とともにインストールされることを想定しています。カーネルをインストールするために emerge を実行する前に、 [installkernel の節](/wiki/Handbook:AMD64/Installation/Kernel/ja#Initramfs "Handbook:AMD64/Installation/Kernel/ja") の記載に従って、initramfs ジェネレータ ( [Dracut](/wiki/Dracut "Dracut") など) を利用するように [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) が設定されていることを確認するべきです。

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

ディストリビューションカーネルを正常にインストールできたら、次の節に進むときです: 次は [システムの設定](/wiki/Handbook:AMD64/Installation/System/ja "Handbook:AMD64/Installation/System/ja") です。

### 別の方法: マニュアル設定

### カーネルソースのインストール

amd64 ベースのシステムにカーネルを手動でインストールしてコンパイルする場合には、Gentoo は[sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources) パッケージを推奨しています。

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

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:AMD64/Installation/Kernel#Distribution_kernels.2Fja "Handbook:AMD64/Installation/Kernel").

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

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:AMD64/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fja "Handbook:AMD64/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:AMD64/Installation/Kernel#Compiling_and_installing.2Fja "Handbook:AMD64/Installation/Kernel")

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

#### 省略可能: カーネルイメージに署名する (セキュアブート)

( [セキュアブート](/wiki/Secure_Boot "Secure Boot") が有効化されたシステムで使用するために) カーネルイメージに署名する場合には、以下のカーネルコンフィグオプションを設定することが推奨されます:

カーネル **セキュアブートのためのロックダウン**

```
General setup  --->
  Kexec and crash features  --->
    [*] Enable kexec system call
    [*] Enable kexec file based system call
    [*]   Verify kernel signature during kexec_file_load() syscall
    [*]     Require a valid signature in kexec_file_load() syscall
    [*]     Enable ""image"" signature verification support

[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Require modules to be validly signed
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

Security options  --->
[*] Integrity subsystem
  [*] Basic module for enforcing kernel lockdown
  [*]   Enable lockdown LSM early in init
        Kernel default lockdown mode (Integrity)  --->

  [*]   Digital signature verification using multiple keyrings
  [*]     Enable asymmetric keys support
  -*-       Require all keys on the integrity keyrings be signed
  [*]       Provide keyring for platform/firmware trusted keys
  [*]       Provide a keyring to which Machine Owner Keys may be added
  [ ]         Enforce Machine Keyring CA Restrictions

```

ただし ""image"" はアーキテクチャ固有のイメージ名に対するプレースホルダです。これらのオプションは上から順に、kexec の呼び出し時にカーネルイメージが署名されていることを強制し (kexec はカーネルをその場で置き換えることを許可します)、カーネルモジュールが署名されていることを強制し、integrity ロックダウンモードを有効化し (実行時にカーネルを変更することを防止します)、各種キーチェーンを有効化します。

圧縮されたカーネルの展開をネイティブにサポートしていないアーキテクチャ (例: **arm64** および **riscv**) では、カーネルは自身の展開プログラム (zboot) とともにビルドしなくてはなりません:

カーネル **zboot CONFIG\_EFI\_ZBOOT**

```
Device Drivers --->
  Firmware Drivers --->
    EFI (Extensible Firmware Interface) Support --->
      [*] Enable the generic EFI decompressor

```

カーネルのコンパイル後は、次の節で説明する通り、カーネルイメージに署名しなくてはなりません。まず [app-crypt/sbsigntools](https://packages.gentoo.org/packages/app-crypt/sbsigntools) をインストールして、次にカーネルイメージに署名してください:

`root #` `emerge --ask app-crypt/sbsigntools`

`root #` `sbsign /usr/src/linux-x.y.z/path/to/kernel-image --cert /path/to/kernel_key.pem --key /path/to/kernel_key.pem --output /usr/src/linux-x.y.z/path/to/kernel-image`

**メモ**

この例では、カーネルイメージに署名するために、モジュールに署名するために生成されたのと同じ鍵が使用されています。カーネルイメージに署名するために、2 個目の別の鍵を生成して使用することも可能です。前節と同じ OpenSSL コマンドをもう一度使用することができます。

それではインストールを続行してください。

他のパッケージによってインストールされる EFI 実行可能形式に自動的に署名するには、グローバルに [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグを有効化してください:

ファイル **`/etc/portage/make.conf`** **セキュアブートを有効化する**

```
USE="modules-sign secureboot"

# 任意で、カスタム署名鍵を使用するために。
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # MODULES_SIGN_KEY が証明書を含まない場合のみ必須です。
MODULES_SIGN_HASH="sha512" # デフォルトは sha512 です

# 任意で、セキュアブートを有効化してブートするために。署名鍵は同一でも別でも構いません。
SECUREBOOT_SIGN_KEY="/path/to/kernel_key.pem"
SECUREBOOT_SIGN_CERT="/path/to/kernel_key.pem"

```

**メモ**

SECUREBOOT\_SIGN\_KEY と SECUREBOOT\_SIGN\_CERT は異なるファイルを指していても構いません。この例では、OpenSSL によって生成される pem ファイルには鍵とそれに伴う証明書の両方を含んでいるため、両変数は同じ値に設定されています。

**メモ**

systemd の `ukify` を使用して [Unified カーネルイメージ](/wiki/Unified_Kernel_Image "Unified Kernel Image") を生成する場合は、カーネルイメージは unified カーネルイメージに含められる前に自動的に署名されるので、手動で署名する必要はありません。

#### コンパイルおよびインストール

さあ、カーネルのコンフィグレーションが完了し、コンパイルとインストールをする時がきました。コンフィグレーションを終了させ、コンパイル作業を開始しましょう:

`root #` `make -j$(nproc) && make modules_install`

**メモ**

make -jX とすることで、並列ビルド数を減らすことができます (X には、並行処理を許可するビルドプロセスの数を整数で指定します)。/etc/portage/make.conf の説明中の、MAKEOPTS 変数についてと同様です。

カーネルのコンパイルが完了したら、カーネルのイメージを /boot/ にコピーします。これは make install コマンドで行います。

`root #` `make install`

このコマンドはカーネルイメージを /boot にコピーします。もし [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) がインストールされていれば、このコマンドは代わりに /sbin/installkernel を呼び出し、カーネルインストールを委譲します。 [Installkernel](/wiki/Installkernel/ja "Installkernel/ja") は単にカーネルを /boot にコピーするのではなく、各カーネルのファイル名にそのバージョン番号を追加してインストールします。さらに、installkernel はカーネルインストールに関連する様々な作業、例えば、 [initramfs](/wiki/Initramfs "Initramfs") の生成、 [Unified カーネルイメージ](/wiki/Unified_Kernel_Image "Unified Kernel Image") のビルド、 [ブートローダ](/wiki/Bootloader/ja "Bootloader/ja") の構成の更新などを、自動的に実現するためのフレームワークも提供しています。

では、 [システムの設定](/wiki/Handbook:AMD64/Installation/System/ja "Handbook:AMD64/Installation/System/ja") に進み、インストールを続けましょう。

[← Gentooベースシステムのインストール](/wiki/Handbook:AMD64/Installation/Base/ja "Previous part") [Home](/wiki/Handbook:AMD64/ja "Handbook:AMD64/ja") [システムの設定 →](/wiki/Handbook:AMD64/Installation/System/ja "Next part")