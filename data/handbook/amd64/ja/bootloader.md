# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Bootloader/de "Handbook:AMD64/Installation/Bootloader/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Bootloader/es "Handbook:AMD64/Installation/Bootloader/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Bootloader/fr "Handbook:AMD64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Bootloader/it "Handbook:AMD64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Bootloader/hu "Handbook:AMD64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Bootloader/pl "Handbook:AMD64/Installation/Bootloader/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Bootloader/pt-br "Handbook:AMD64/Installation/Bootloader/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Bootloader/cs "Handbook:AMD64/Installation/Bootloader/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Bootloader/ru "Handbook:AMD64/Installation/Bootloader/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Bootloader/ta "Handbook:AMD64/Installation/Bootloader/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Bootloader/zh-cn "Handbook:AMD64/Installation/Bootloader/zh-cn (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:AMD64/Installation/Bootloader/ko "Handbook:AMD64/Installation/Bootloader/ko (100% translated)")

[AMD64 ハンドブック](/wiki/Handbook:AMD64 "Handbook:AMD64")[インストール](/wiki/Handbook:AMD64/Full/Installation "Handbook:AMD64/Full/Installation")[インストールについて](/wiki/Handbook:AMD64/Installation/About/ja "Handbook:AMD64/Installation/About/ja")[メディアの選択](/wiki/Handbook:AMD64/Installation/Media/ja "Handbook:AMD64/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:AMD64/Installation/Networking/ja "Handbook:AMD64/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:AMD64/Installation/Disks/ja "Handbook:AMD64/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:AMD64/Installation/Stage/ja "Handbook:AMD64/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:AMD64/Installation/Base/ja "Handbook:AMD64/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:AMD64/Installation/Kernel/ja "Handbook:AMD64/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:AMD64/Installation/System/ja "Handbook:AMD64/Installation/System/ja")[ツールのインストール](/wiki/Handbook:AMD64/Installation/Tools/ja "Handbook:AMD64/Installation/Tools/ja")ブートローダの設定[締めくくり](/wiki/Handbook:AMD64/Installation/Finalizing/ja "Handbook:AMD64/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:AMD64/Full/Working "Handbook:AMD64/Full/Working")[Portage について](/wiki/Handbook:AMD64/Working/Portage/ja "Handbook:AMD64/Working/Portage/ja")[USE フラグ](/wiki/Handbook:AMD64/Working/USE/ja "Handbook:AMD64/Working/USE/ja")[Portage の機能](/wiki/Handbook:AMD64/Working/Features/ja "Handbook:AMD64/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:AMD64/Working/Initscripts/ja "Handbook:AMD64/Working/Initscripts/ja")[環境変数](/wiki/Handbook:AMD64/Working/EnvVar/ja "Handbook:AMD64/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:AMD64/Full/Portage "Handbook:AMD64/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:AMD64/Portage/Files/ja "Handbook:AMD64/Portage/Files/ja")[変数](/wiki/Handbook:AMD64/Portage/Variables/ja "Handbook:AMD64/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:AMD64/Portage/Branches/ja "Handbook:AMD64/Portage/Branches/ja")[追加ツール](/wiki/Handbook:AMD64/Portage/Tools/ja "Handbook:AMD64/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:AMD64/Portage/CustomTree/ja "Handbook:AMD64/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:AMD64/Portage/Advanced/ja "Handbook:AMD64/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:AMD64/Full/Networking "Handbook:AMD64/Full/Networking")[はじめに](/wiki/Handbook:AMD64/Networking/Introduction/ja "Handbook:AMD64/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:AMD64/Networking/Advanced/ja "Handbook:AMD64/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:AMD64/Networking/Modular/ja "Handbook:AMD64/Networking/Modular/ja")[無線](/wiki/Handbook:AMD64/Networking/Wireless/ja "Handbook:AMD64/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:AMD64/Networking/Extending/ja "Handbook:AMD64/Networking/Extending/ja")[動的な管理](/wiki/Handbook:AMD64/Networking/Dynamic/ja "Handbook:AMD64/Networking/Dynamic/ja")

## Contents

- [1ブートローダを選ぶ](#.E3.83.96.E3.83.BC.E3.83.88.E3.83.AD.E3.83.BC.E3.83.80.E3.82.92.E9.81.B8.E3.81.B6)
- [2デフォルト: GRUB](#.E3.83.87.E3.83.95.E3.82.A9.E3.83.AB.E3.83.88:_GRUB)
  - [2.1Emerge](#Emerge)
  - [2.2インストール](#.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB)
    - [2.2.1DOS/レガシー BIOS システム](#DOS.2F.E3.83.AC.E3.82.AC.E3.82.B7.E3.83.BC_BIOS_.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
    - [2.2.2UEFI システム](#UEFI_.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
      - [2.2.2.1追加可能: セキュアブート](#.E8.BF.BD.E5.8A.A0.E5.8F.AF.E8.83.BD:_.E3.82.BB.E3.82.AD.E3.83.A5.E3.82.A2.E3.83.96.E3.83.BC.E3.83.88)
      - [2.2.2.2GRUB をデバッグする](#GRUB_.E3.82.92.E3.83.87.E3.83.90.E3.83.83.E3.82.B0.E3.81.99.E3.82.8B)
  - [2.3設定](#.E8.A8.AD.E5.AE.9A)
- [3代替案 1: systemd-boot](#.E4.BB.A3.E6.9B.BF.E6.A1.88_1:_systemd-boot)
  - [3.1Emerge](#Emerge_2)
  - [3.2インストール](#.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB_2)
  - [3.3追加可能: セキュアブート](#.E8.BF.BD.E5.8A.A0.E5.8F.AF.E8.83.BD:_.E3.82.BB.E3.82.AD.E3.83.A5.E3.82.A2.E3.83.96.E3.83.BC.E3.83.88_2)
- [4代替案 2: EFI スタブ](#.E4.BB.A3.E6.9B.BF.E6.A1.88_2:_EFI_.E3.82.B9.E3.82.BF.E3.83.96)
  - [4.1Unified カーネルイメージ](#Unified_.E3.82.AB.E3.83.BC.E3.83.8D.E3.83.AB.E3.82.A4.E3.83.A1.E3.83.BC.E3.82.B8)
- [5他の選択肢](#.E4.BB.96.E3.81.AE.E9.81.B8.E6.8A.9E.E8.82.A2)
- [6システムのリブート](#.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.81.AE.E3.83.AA.E3.83.96.E3.83.BC.E3.83.88)

## ブートローダを選ぶ

これまでLinuxカーネルを設定すると共に、システムツールをインストールし、設定ファイルを修正してきました。そして今、最も重要なLinuxインストールの最後の一片をインストールします。それがブートローダーです。

ブートローダーは、ブート中にLinuxカーネルを起動することに責任を負っています。ブートローダーがないと、システムは電源ボタンが押されたときに、どう事を進めればいいのかわからなくなってしまいます。

**amd64** に対して私たちは、DOS/レガシー BIOS ベースのシステムについては [GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/ja#Default:_GRUB "Handbook:AMD64/Blocks/Bootloader/ja") を設定する方法を、UEFI システムについては [GRUB](#Default:_GRUB)、 [systemd-boot](#Alternative_1:_systemd-boot)、または [efibootmgr](#Alternative_2:_efibootmgr) を設定する方法を文書化しています。

このセクションでは、ブートローダーパッケージの "emerge" と、ブートローダーのシステムへの "インストール" という表現を使っています。ここでいう "emerge" とは [Portage](/wiki/Portage "Portage") を使ってソフトウェアパッケージがシステムで利用できるようにすることです。そして "インストール" はブートローダーが必要なファイルをコピーしたりディスク上の特定の領域に変更を加えることで、ブートローダーを有効化し次回起動時に使用可能な状態にすることを指します。

## デフォルト: GRUB

デフォルトでは、大半の Gentoo システムが [GRUB Legacy](/wiki/GRUB_Legacy "GRUB Legacy") の後継である [GRUB](/wiki/GRUB/ja "GRUB/ja") ([sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) パッケージで利用できます) を使用しています。GRUB は追加の設定なしに従来の BIOS ("pc") システムで使うことができ、それ以外のプラットフォームでもビルド前のわずかな設定で済みます。詳しくは、 [GRUB](/wiki/GRUB/ja "GRUB/ja") ページの [前提条件](/wiki/GRUB/ja#.E5.89.8D.E6.8F.90.E6.9D.A1.E4.BB.B6 "GRUB/ja") 節をご覧ください。

### Emerge

MBR パーティションテーブルのみをサポートする従来の BIOS システムを使う場合、GRUBをインストールするのに追加の設定は必要ありません:

`root #` `emerge --ask --verbose sys-boot/grub`

UEFI ユーザーの方へ: 上記のコマンドを実行すると、現在有効な GRUB\_PLATFORMS 値が表示されます。UEFI 対応のシステムでは `GRUB_PLATFORMS="efi-64"` が有効になっていることを確認してください (これがデフォルトです) 。もし有効になっていなければ、GRUB パッケージを EFI の機能付きでビルドするために、 GRUB を emerge する _前に_/etc/portage/make.conf に `GRUB_PLATFORMS="efi-64"` を追加しなければなりません。

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub`

なんらかの経緯で `GRUB_PLATFORMS="efi-64"` を有効にしていない状態で GRUB が emerge されてしまった場合は、この行を make.conf に追加して、 emerge に `--update --newuse` オプションを渡せば、 [world パッケージセット](/wiki/World_set_(Portage) "World set (Portage)") の依存関係を再計算することができます:

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub`

これで GRUB ソフトウェアが _システム_ にマージされましたが、二次 _ブートローダ_ としてのインストールはまだ完了していません。

### インストール

つぎに、 grub-install コマンドを使って、必要な GRUB ファイルを /boot/grub/ ディレクトリにインストールします。もし（システムがブートする）一番目のディスクにインストールするなら、 /dev/sda ですので、以下のどちらかのコマンドでインストールすることができます:

#### DOS/レガシー BIOS システム

DOS/レガシー BIOS システムでは:

`root #` `grub-install /dev/sda`

#### UEFI システム

**重要**

grub-install を実行する _前に_ EFI システムパーティションがマウントされているか必ず確認してください。 grub-install が GRUB EFI ファイル (grubx64.efi) を間違ったディレクトリにインストールしてしまい、しかも間違ったディレクトリが使われた形跡を _まったく残さない_ ということが起こりえます。

UEFI システムでは:

`root #` `grub-install --efi-directory=/efi`

```
Installing for x86_64-efi platform.
Installation finished. No error reported.

```

インストールに成功したら、出力は上のコマンドの出力と一致するはずです。出力が完全に一致していない場合は、 [GRUB をデバッグする](/wiki/Handbook:AMD64/Blocks/Bootloader/ja#Debugging_GRUB "Handbook:AMD64/Blocks/Bootloader/ja") へと進み、そうでない場合は [設定の手順](/wiki/Handbook:AMD64/Blocks/Bootloader/ja#GRUB_Configure "Handbook:AMD64/Blocks/Bootloader/ja") へと進んでください。

##### 追加可能: セキュアブート

セキュアブートを有効化して正常にブートするためには、署名証明書が [UEFI](/wiki/UEFI "UEFI") ファームウェアによって受け入れられるか、 [shim](/wiki/Shim "Shim") をプリローダとして使用しなくてはなりません。shim はサードパーティの Microsoft の証明書によって署名済みで、ほとんどの UEFI マザーボードによってデフォルトで受け入れられます。

カスタム鍵を受け入れるように UEFI ファームウェアを設定する方法はファームウェアの製造元によって異なり、このハンドブックの対象外です。代わりに以下に shim をセットアップするための方法を示します。ここではユーザは既に以前の節の手順に従って、署名鍵を生成して portage にそれを使用させるように構成したものと仮定します。もしそうでない場合は、まず [カーネルのインストール](/wiki/Handbook:AMD64/Installation/Kernel/ja "Handbook:AMD64/Installation/Kernel/ja") の節に戻ってください。

パッケージ [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) は、[secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグが有効化されていれば、ビルド済みで署名済みのスタンドアローンな EFI 実行可能形式をインストールします。必要なパッケージをインストールして、スタンドアローンな grub、Shim、および MokManager を EFI システムパーティション上の同じディレクトリにコピーしてください。例えば:

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

次に、署名鍵を shims MOKlist に登録してください。sbsign とカーネルビルドシステムは PEM 形式の鍵を期待する一方で、shims MOKlist は DER 形式の鍵を要求します。ハンドブックの以前の節で示した例では、PEM 署名鍵などを生成していたので、ここでこの鍵を DER 形式に変換する必要があります:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**メモ**

ここで使用されるパスは、生成された鍵に属する証明書を含む pem ファイルへのパスでなくてはなりません。この例では鍵と証明書は同じ pem ファイル内にあります。

そして、変換された証明書を shims MOKlist にインポートしてください。このコマンドでは、インポートリクエスト用のいくつかのパスワードを設定するよう求められます:

`root #` `mokutil --import /path/to/kernel_key.der`

**メモ**

現在起動しているカーネルが、インポートしようとしている証明書を既に信頼している場合、"Already in kernel trusted keyring." というメッセージが表示されます。この場合には、先ほどのコマンドに --ignore-keyring という引数を追加して再実行してください。

次は、UEFI ファームウェアに Shim を登録します。次のコマンドでは、 `boot-disk` と `boot-partition-id` はそれぞれ、EFI システムパーティションのディスクとパーティションの識別子で置き換えてください:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\Gentoo\shimx64.efi' --label 'GRUB via Shim' --unicode`

このビルド済みで署名済みのスタンドアローンな grub は、通常と異なる場所から grub.cfg を読み込むことに注意してください。設定ファイルはデフォルトの /boot/grub/grub.cfg ではなく、grub の EFI 実行ファイルと同じディレクトリ、例えば /efi/EFI/Gentoo/grub.cfg に置く必要があります。カーネルのインストールと grub の設定の更新に [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) を使用しているなら、GRUB\_CFG 環境変数を使ってデフォルトの grub 設定ファイルの場所を上書きできます。

例えば:

`root #` `grub-mkconfig -o /efi/EFI/Gentoo/grub.cfg`

または、 [installkernel](/wiki/Installkernel/ja "Installkernel/ja") を利用して:

ファイル **`/etc/env.d/99grub`**

```
GRUB_CFG=/efi/EFI/Gentoo/grub.cfg

```

`root #` `env-update`

**メモ**

インポート処理はシステムを再起動するまで完了しません。このハンドブックのすべてのステップを終えた後にシステムを再起動して shim を読み込むと、shim は mokutil によって登録されたインポートリクエストを探します。そして、MokManager アプリケーションが開始され、インポートリクエストを作成した時に設定したパスワードを尋ねられます。画面に表示される手順にしたがって証明書のインポートを完了したら、システムを再起動して UEFI メニューに入り、セキュアブートの設定を有効にしてください。

##### GRUB をデバッグする

GRUB をデバッグするとき、新しい live イメージ環境へと再起動することなくインストール済みのシステムをブート可能にできるかもしれない簡単な解決策がいくつか存在します。

出力の中のどこかに "EFI variables are not supported on this system" が表示されている場合は、おそらく live イメージが EFI モードでブートされておらず、現在レガシー BIOS ブートモードにある可能性が高いです。解決方法としては、以下で言及されている [removable オプションを使う GRUB の手順](/wiki/Handbook:AMD64/Blocks/Bootloader/ja#GRUB_Install_EFI_systems_removable "Handbook:AMD64/Blocks/Bootloader/ja") を試してください。これは /EFI/BOOT/BOOTX64.EFI にある実行可能 EFI ファイルを上書きします。EFI モードで再起動できれば、マザーボードファームウェアはこのデフォルトブートエントリを実行し、GRUB を実行できるかもしれません。

grub-install が "Could not prepare Boot variable: Read-only file system" というエラーを返し、live 環境が正しく UEFI モードでブートされている場合は、efivars という特別なマウントを読み書き可能な状態で再マウントしたうえで [先述の grub-install コマンド](/wiki/Handbook:AMD64/Blocks/Bootloader/ja#GRUB_Install_EFI_systems_command "Handbook:AMD64/Blocks/Bootloader/ja") を再実行できるはずです:

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

これは特殊な EFI ファイルシステムをデフォルトではマウントしない、特定の非公式 Gentoo 環境によって引き起こされます。上のコマンドが実行できない場合は、公式 live イメージ環境を EFI モードで使用して再起動してください。

一部のマザーボードメーカの貧弱な UEFI 実装では、EFI システムパーティション (ESP) 内の .EFI ファイルの置き場所として /EFI/BOOT ディレクトリ _しか_ サポートしていないようです。GRUB のインストーラは、インストールコマンドに `--removable` オプションを付けることで、自動でこの場所に .EFI ファイルを作成することができます。ESP がマウントされていることを確認してから、以下のコマンドを実行してください; (以前に定義した通り) /efi にマウントされているとして、以下を実行してください:

`root #` `grub-install --target=x86_64-efi --efi-directory=/efi --removable`

このコマンドは UEFI 仕様で定義されている'デフォルトの'ディレクトリを作成し、デフォルトの名前 BOOTX64.EFI を持つファイルを作成します。

### 設定

次に、/etc/default/grub ファイルと /etc/grub.d スクリプトで指定されたユーザ固有の設定をもとに、GRUB 設定ファイルを生成します。GRUB はどのカーネルを起動するか（/boot/ 内で利用可能な最上位のもの）、どれがルートファイルシステムかを自動で検出してくれるので、ほとんどの場合、ユーザによる設定の必要はありません。カーネルパラメータは /etc/default/grub の GRUB\_CMDLINE\_LINUX 変数で指定することができます。

最終的な GRUB の設定ファイルを生成するには、 grub-mkconfig コマンドを実行します:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.18.8-gentoo
Found initrd image: /boot/initramfs-genkernel-amd64-6.18.8-gentoo
done

```

このコマンドの出力を見て、ブートに必要な Linux イメージが見つかったという報告が少なくともひとつあることを確認してください。もし initramfs を使っているか genkernel でカーネルをビルドしている場合は、正しい initrd イメージが認識されていることも確認してください。もし確認できなかった場合、/boot/ にそれらのファイルが存在するかどうか ls コマンドで確かめてください。必要なファイルが存在していなければ、カーネルの設定とインストールをやり直さなければなりません。

**ヒント**

接続されたドライブからほかのOSを検出するために、os-prober ユーティリティを使うこともできます。Windows 7、8.1、10、あるいはほかの Linux ディストリビューションが検出できるようになります。このようなデュアルブート環境を作るには、[sys-boot/os-prober](https://packages.gentoo.org/packages/sys-boot/os-prober) パッケージをインストールしてから grub-mkconfig コマンドを再実行するとよいでしょう。検出がうまくいかない時は、Gentoo コミュニティに助けを求める前に [GRUB](/wiki/GRUB/ja "GRUB/ja") 記事をよく読み直してみてください。

## 代替案 1: systemd-boot

別の選択肢として [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") もあり、これは OpenRC と systemd のマシンの両方で動作します。これは薄いチェーンローダであり、セキュアブートと合わせてうまく動作します。

### Emerge

systemd-boot をインストールするには、[boot](https://packages.gentoo.org/useflags/boot) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグを有効化して、[sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) (systemd の場合) または [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) (OpenRC の場合) を再インストールしてください:

ファイル **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot

```

`root #` `emerge --ask sys-apps/systemd`

または

`root #` `emerge --ask sys-apps/systemd-utils`

### インストール

そうしたら、systemd-boot ローダを [EFI システムパーティション](/wiki/EFI_System_Partition/ja "EFI System Partition/ja") にインストールしてください:

`root #` `bootctl install`

**重要**

bootctl install を実行する _前に_、EFI システムパーティションがマウントされていることを確認してください。

このブートローダを使用するときは、リブートする前に、以下を使用してブート可能なエントリが存在することを確認してください:

`root #` `bootctl list`

**警告**

新しい systemd-boot 用のカーネルコマンドラインは /etc/kernel/cmdline または /usr/lib/kernel/cmdline から読み込まれます。どちらのファイルも存在しない場合には、現在起動しているカーネルのカーネルコマンドラインが再利用されます (/proc/cmdline)。そのため、新規インストールでは、live CD のカーネルコマンドラインが新しいカーネルの起動に意図せず使用されてしまうことがあります。登録されているエントリーのカーネルコマンドラインは以下のようにして確認できます:

`root #` `bootctl list`

望みどおりのカーネルコマンドラインが表示されない場合には、正しいカーネルコマンドラインを含む /etc/kernel/cmdline を作成してからカーネルを再インストールしてください。

新しいエントリが存在しない場合は、[sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) パッケージが [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") および [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグを有効化した状態でインストールされていることを確認して、カーネルインストールを再実行してください。

ディストリビューションカーネルについては:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

手動で設定およびコンパイルされたカーネルについては:

`root #` `make install`

**重要**

systemd-boot 向けにカーネルをインストールする場合、root= カーネルコマンドライン引数はデフォルトでは追加されません。initramfs を使用する systemd システム上で、ブート時にルートパーティションを自動的に見つけるためには、ユーザは代わりに [systemd-gpt-auto-generator](/wiki/Systemd/ja#Automatic_mounting_of_partitions_at_boot "Systemd/ja") に頼ることができます。そうしない場合ユーザは、/etc/kernel/cmdline 内で、使用すべき他のカーネルコマンドライン引数と並んで root= を設定することで、手動でルートパーティションの場所を指定するべきです。その後、上記の通りにカーネルを再インストールしてください。

### 追加可能: セキュアブート

[secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグが有効化されている場合、systemd-boot EFI 実行可能形式は portage によって自動的に署名されるでしょう。さらに、bootctl install は自動的に署名されたバージョンをインストールするでしょう。

セキュアブートを有効化して正常にブートするためには、 [UEFI](/wiki/UEFI "UEFI") ファームウェアによって使用された証明書が受け入れられるか、 [shim](/wiki/Shim "Shim") をプリローダとして使用しなくてはなりません。shim はサードパーティの Microsoft の証明書によって署名済みで、ほとんどの UEFI マザーボードによってデフォルトで受け入れられます。

カスタム鍵を受け入れるように UEFI ファームウェアを設定する方法はファームウェアの製造元によって異なり、このハンドブックの対象外です。代わりに以下に shim をセットアップするための方法を示します。ここではユーザは既に以前の節の手順に従って、署名鍵を生成して portage にそれを使用させるように構成したものと仮定します。もしそうでない場合は、まず [カーネルのインストール](/wiki/Handbook:AMD64/Installation/Kernel/ja "Handbook:AMD64/Installation/Kernel/ja") の節に戻ってください。

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

sbsign とカーネルビルドシステムは PEM 形式の鍵を期待する一方で、shims MOKlist は DER 形式の鍵を要求します。ハンドブックの以前の節で示した例では、PEM 署名鍵などを生成していたので、ここでこの鍵を DER 形式に変換する必要があります:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**メモ**

ここで使用されるパスは、生成された鍵に属する証明書を含む pem ファイルへのパスでなくてはなりません。この例では鍵と証明書は同じ pem ファイル内にあります。

そして、変換された証明書を Shims MOKlist にインポートしてください:

`root #` `mokutil --import /path/to/kernel_key.der`

**メモ**

現在起動しているカーネルが、インポートしようとしている証明書を既に信頼している場合、"Already in kernel trusted keyring." というメッセージが表示されます。この場合には、先ほどのコマンドに --ignore-keyring という引数を追加して再実行してください。

最後に、UEFI ファームウェアに Shim を登録します。次のコマンドでは、 `boot-disk` と `boot-partition-id` はそれぞれ、EFI システムパーティションのディスクとパーティションの識別子で置き換えてください:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**メモ**

インポート処理はシステムを再起動するまで完了しません。このハンドブックのすべてのステップを終えた後にシステムを再起動して shim を読み込むと、shim は mokutil によって登録されたインポートリクエストを探します。そして、MokManager アプリケーションが開始され、インポートリクエストを作成した時に設定したパスワードを尋ねられます。画面に表示される手順にしたがって証明書のインポートを完了したら、システムを再起動して UEFI メニューに入り、セキュアブートの設定を有効にしてください。

## 代替案 2: EFI スタブ

UEFI ベースのファームウェアを搭載したコンピュータシステムは、カーネルをブートするために二次ブートローダ (GRUB 等) を厳密には必要としません。二次ブートローダはブートプロセス中の UEFI ファームウェアの機能を _拡張_ するために存在しています。GRUB を使用する (以前の節を参照してください) ほうが、ブート時のカーネルパラメータを簡単に変更するためのより柔軟なアプローチを提供しているので、典型的にはより簡単で、より堅牢です。

システムのブートについて最小限かつ厳格なアプローチをとるシステム管理者は、二次ブートローダを使用せずに、 [EFI スタブ](/wiki/EFI_stub/ja "EFI stub/ja") として Linux カーネルをブートすることができます。

[sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) アプリケーションは、システムの一次ブートローダである UEFI ファームウェアと通信するためのツールです。通常このツールは、ファームウェアのブート可能エントリの一覧にブートエントリを追加または削除するために使われます。また、既にブート可能エントリとして追加されている Linux カーネルを追加のオプションとともに実行できるように、ファームウェア設定を更新することもできます。これらの通信は EFI 変数と呼ばれる特別なデータ構造を通して行われます (したがって、EFI 変数のカーネルサポートが必要です)。

続ける _前_ に [EFI スタブ](/wiki/EFI_stub/ja "EFI stub/ja") カーネルの記事を必ず確認してださい。カーネルを UEFI ファームウェアから直接ブートできるようにするには、特有のオプションを有効化しなければなりません。このサポートを組み込むために、カーネルの再コンパイルが必要になる場合があります。

追加の情報を得るために [efibootmgr](/wiki/Efibootmgr/ja "Efibootmgr/ja") の記事を見てみるのも良い考えです。

**メモ**

繰り返しますが、efibootmgr は UEFI システムのブートにおいて必須では _ありません_; 単に UEFI ファームウェアに EFI スタブカーネルのためのエントリを追加するためだけに必要なものです。EFI スタブのサポート付きで適切にビルドされた Linux カーネルは、それ自体 _直接_ ブートさせることができます。追加のカーネルコマンドラインオプションも Linux カーネルの中に組み込むことができます (CONFIG\_CMDLINE というカーネルの設定オプションがあります)。同様に、initramfs のサポートもカーネルの中に'組み込む'ことができます。

カーネルをファームウェアから直接ブートするためには、カーネルイメージが [EFI システムパーティション](/wiki/EFI_System_Partition/ja "EFI System Partition/ja") 上に存在していなくてはなりません。これは [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) に対して [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/ja](/wiki/USE_flag/ja "USE flag/ja") USE フラグを有効化することで実現できます。EFI スタブブートはすべての UEFI システム上で機能することが保証されていないことを考慮し、この USE フラグは stable ではマスクされています。この機能を使用するためには、installkernel に対して testing キーワードを受諾しなくてはなりません。

ファイル **`/etc/portage/package.accept_keywords/installkernel`**

```
sys-kernel/installkernel
sys-boot/uefi-mkconfig
app-emulation/virt-firmware

```

ファイル **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel efistub

```

そして [installkernel](/wiki/Installkernel/ja "Installkernel/ja") を再インストールして、/efi/EFI/Gentoo ディレクトリを作成し、カーネルを再インストールしてください:

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi/EFI/Gentoo`

ディストリビューションカーネルでは:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel{,-bin}`

手動管理のカーネルでは:

`root #` `make install`

または、 [installkernel](/wiki/Installkernel/ja "Installkernel/ja") を使用しない場合は、カーネルを bzImage.efi という名前で EFI システムパーティションに手動でコピーしても構いません:

`root #` `mkdir -p /efi/EFI/Gentoo
`

`root #` `cp /boot/vmlinuz-* /efi/EFI/Gentoo/bzImage.efi
`

efibootmgr パッケージをインストールしてください:

`root #` `emerge --ask sys-boot/efibootmgr`

新しくコンパイルされた EFI スタブカーネルに対応する "Gentoo" という名称のブートエントリを、UEFI ファームウェア内に作成してください:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Gentoo\bzImage.efi"`

**メモ**

UEFI の定義には、ディレクトリパスの区切りにはバックスラッシュ ( `\`) を用いなければなりません。

初期 RAM ファイルシステム (initramfs) を用いるときには、これも EFI システムパーティションにコピーして、適切なブートオプションを加えてください:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Gentoo\bzImage.efi" --unicode "initrd=\EFI\Gentoo\initramfs.img"`

**ヒント**

追加のカーネルコマンドラインオプションは、上に示した initrd=... オプションとともに指定することで、ファームウェアによってカーネルにパースされます。

これらの変更が完了したら、システムを再起動後から、"gentoo" という名称のブートエントリーが利用可能になります。

### Unified カーネルイメージ

[installkernel](/wiki/Installkernel/ja "Installkernel/ja") が unified カーネルイメージをビルドしてインストールするように構成されているなら、unified カーネルイメージはすでに EFI システムパーティション上の EFI/Linux ディレクトリにインストールされているはずです。もしそうなっていない場合は、ディレクトリが存在することを確認してから、ハンドブックで既に説明した通りにカーネルインストールを再度実行してください。

インストールされた unified カーネルイメージのためのブートエントリを直接追加するには:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## 他の選択肢

ハンドブックでは取り扱わない他の選択肢については、利用可能な [ブートローダ](/wiki/Bootloader/ja "Bootloader/ja") の一覧を参照してください。

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

リブートして新しい Gentoo 環境に入ることができたら、最終章の [インストールの締めくくり](/wiki/Handbook:AMD64/Installation/Finalizing/ja "Handbook:AMD64/Installation/Finalizing/ja") に進むのがよいでしょう。

[← ツールのインストール](/wiki/Handbook:AMD64/Installation/Tools/ja "Previous part") [Home](/wiki/Handbook:AMD64/ja "Handbook:AMD64/ja") [インストールの締めくくり →](/wiki/Handbook:AMD64/Installation/Finalizing/ja "Next part")