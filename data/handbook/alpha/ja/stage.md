# Stage

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Stage/de "Handbuch:Alpha/Installation/Stage (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Stage "Handbook:Alpha/Installation/Stage (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Stage/tr "Handbook:Alpha/Installation/Stage/tr (0% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Stage/es "Manual de Gentoo/Instalación/Stage (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Stage/fr "Handbook:Alpha/Installation/Stage/fr (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Stage/it "Handbook:Alpha/Installation/Stage/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Stage/hu "Handbook:Alpha/Installation/Stage/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Stage/pl "Handbook:Alpha/Installation/Stage (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Stage/pt-br "Manual:Alpha/Instalação/Stage (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Stage/cs "Handbook:Alpha/Installation/Stage/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Stage/ru "Handbook:Alpha/Installation/Stage (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Stage/ta "கையேடு:Alpha/நிறுவல்/நிலை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Stage/zh-cn "手册：Alpha/安装/安装stage3 (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:Alpha/Installation/Stage/ko "Handbook:Alpha/Installation/Stage/ko (100% translated)")

[Alpha ハンドブック](/wiki/Handbook:Alpha "Handbook:Alpha")[インストール](/wiki/Handbook:Alpha/Full/Installation "Handbook:Alpha/Full/Installation")[インストールについて](/wiki/Handbook:Alpha/Installation/About/ja "Handbook:Alpha/Installation/About/ja")[メディアの選択](/wiki/Handbook:Alpha/Installation/Media/ja "Handbook:Alpha/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:Alpha/Installation/Networking/ja "Handbook:Alpha/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:Alpha/Installation/Disks/ja "Handbook:Alpha/Installation/Disks/ja")stage ファイル[ベースシステムのインストール](/wiki/Handbook:Alpha/Installation/Base/ja "Handbook:Alpha/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:Alpha/Installation/Kernel/ja "Handbook:Alpha/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:Alpha/Installation/System/ja "Handbook:Alpha/Installation/System/ja")[ツールのインストール](/wiki/Handbook:Alpha/Installation/Tools/ja "Handbook:Alpha/Installation/Tools/ja")[ブートローダの設定](/wiki/Handbook:Alpha/Installation/Bootloader/ja "Handbook:Alpha/Installation/Bootloader/ja")[締めくくり](/wiki/Handbook:Alpha/Installation/Finalizing/ja "Handbook:Alpha/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:Alpha/Full/Working "Handbook:Alpha/Full/Working")[Portage について](/wiki/Handbook:Alpha/Working/Portage/ja "Handbook:Alpha/Working/Portage/ja")[USE フラグ](/wiki/Handbook:Alpha/Working/USE/ja "Handbook:Alpha/Working/USE/ja")[Portage の機能](/wiki/Handbook:Alpha/Working/Features/ja "Handbook:Alpha/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:Alpha/Working/Initscripts/ja "Handbook:Alpha/Working/Initscripts/ja")[環境変数](/wiki/Handbook:Alpha/Working/EnvVar/ja "Handbook:Alpha/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:Alpha/Full/Portage "Handbook:Alpha/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:Alpha/Portage/Files/ja "Handbook:Alpha/Portage/Files/ja")[変数](/wiki/Handbook:Alpha/Portage/Variables/ja "Handbook:Alpha/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:Alpha/Portage/Branches/ja "Handbook:Alpha/Portage/Branches/ja")[追加ツール](/wiki/Handbook:Alpha/Portage/Tools/ja "Handbook:Alpha/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:Alpha/Portage/CustomTree/ja "Handbook:Alpha/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:Alpha/Portage/Advanced/ja "Handbook:Alpha/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:Alpha/Full/Networking "Handbook:Alpha/Full/Networking")[はじめに](/wiki/Handbook:Alpha/Networking/Introduction/ja "Handbook:Alpha/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:Alpha/Networking/Advanced/ja "Handbook:Alpha/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:Alpha/Networking/Modular/ja "Handbook:Alpha/Networking/Modular/ja")[無線](/wiki/Handbook:Alpha/Networking/Wireless/ja "Handbook:Alpha/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:Alpha/Networking/Extending/ja "Handbook:Alpha/Networking/Extending/ja")[動的な管理](/wiki/Handbook:Alpha/Networking/Dynamic/ja "Handbook:Alpha/Networking/Dynamic/ja")

## Contents

- [1stage ファイルを選択する](#stage_.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.92.E9.81.B8.E6.8A.9E.E3.81.99.E3.82.8B)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
  - [1.3multilib (32 ビットと 64 ビット)](#multilib_.2832_.E3.83.93.E3.83.83.E3.83.88.E3.81.A8_64_.E3.83.93.E3.83.83.E3.83.88.29)
  - [1.4非 multilib (64 ビットのみ)](#.E9.9D.9E_multilib_.2864_.E3.83.93.E3.83.83.E3.83.88.E3.81.AE.E3.81.BF.29)
- [2stage ファイルをダウンロードする](#stage_.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.92.E3.83.80.E3.82.A6.E3.83.B3.E3.83.AD.E3.83.BC.E3.83.89.E3.81.99.E3.82.8B)
  - [2.1日時を設定する](#.E6.97.A5.E6.99.82.E3.82.92.E8.A8.AD.E5.AE.9A.E3.81.99.E3.82.8B)
    - [2.1.1自動](#.E8.87.AA.E5.8B.95)
    - [2.1.2手動](#.E6.89.8B.E5.8B.95)
  - [2.2グラフィカルブラウザ](#.E3.82.B0.E3.83.A9.E3.83.95.E3.82.A3.E3.82.AB.E3.83.AB.E3.83.96.E3.83.A9.E3.82.A6.E3.82.B6)
  - [2.3コマンドラインブラウザ](#.E3.82.B3.E3.83.9E.E3.83.B3.E3.83.89.E3.83.A9.E3.82.A4.E3.83.B3.E3.83.96.E3.83.A9.E3.82.A6.E3.82.B6)
  - [2.4検証して確認する](#.E6.A4.9C.E8.A8.BC.E3.81.97.E3.81.A6.E7.A2.BA.E8.AA.8D.E3.81.99.E3.82.8B)
- [3stage ファイルをインストールする](#stage_.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.92.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB.E3.81.99.E3.82.8B)
- [4コンパイルオプションを設定する](#.E3.82.B3.E3.83.B3.E3.83.91.E3.82.A4.E3.83.AB.E3.82.AA.E3.83.97.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E8.A8.AD.E5.AE.9A.E3.81.99.E3.82.8B)
  - [4.1はじめに](#.E3.81.AF.E3.81.98.E3.82.81.E3.81.AB)
  - [4.2CFLAGS と CXXFLAGS](#CFLAGS_.E3.81.A8_CXXFLAGS)
  - [4.3RUSTFLAGS](#RUSTFLAGS)
  - [4.4MAKEOPTS](#MAKEOPTS)
  - [4.5よーい、ドン！](#.E3.82.88.E3.83.BC.E3.81.84.E3.80.81.E3.83.89.E3.83.B3.EF.BC.81)
- [5参照](#.E5.8F.82.E7.85.A7)

## stage ファイルを選択する

**ヒント**

デスクトップ (グラフィカル) OS 環境を目的にするユーザは、それがサポートされているアーキテクチャでは、名前に `desktop` の項を持つ stage ファイルを使用することが推奨されます。これらのファイルは [llvm-core/llvm](https://packages.gentoo.org/packages/llvm-core/llvm) や [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) などのパッケージと、USE フラグのチューニングを含んでおり、インストール時間を大きく改善します。

[stage ファイル](/wiki/Stage_file/ja "Stage file/ja") は Gentoo インストールの種として機能します。stage ファイルは [リリースエンジニアリングチーム](/wiki/Project:RelEng "Project:RelEng") によって [Catalyst](/wiki/Catalyst "Catalyst") を使用して生成されます。stage ファイルは、特定の [プロファイル](/wiki/Profile_(Portage) "Profile (Portage)") を基礎としており、ほぼ完全なシステムを含んでいます。

stage ファイルを選択するときには、所望のシステムの種別に応じたプロファイルターゲットを持つもの選ぶことが重要です。

**重要**

インストールが確立された後にメジャープロファイルを変更することは、技術的には可能ですが、変更にはかなりの労力と熟慮が必要で、このインストールマニュアルの範囲外です。init システムの切り換えは難しいですが、 `no-multilib` から `multilib` への切り換えは Gentoo と低レベルツールチェーンについての深い知識を必要とします。

**ヒント**

大部分のユーザは、'advanced' な tarball を選択する必要は無いはずです。これらは、典型的でないあるいは高度な、ソフトウェアもしくはハードウェアのためのものです。

### OpenRC

[OpenRC](/wiki/OpenRC/ja "OpenRC/ja") は依存関係に基づく init システム (カーネルが起動した後にシステムサービスを開始するためのシステム) で、通常は /sbin/init にある、システムが提供する init プログラムとの互換性を保っています。Gentoo に由来するオリジナルの init システムですが、他の一部の Linux ディストリビューションや BSD システムでも採用されています。

### systemd

systemd は SysV スタイルの init と rc の、Linux システム向けの現代的な代替です。Linux ディストリビューションの大多数では、第一の init システムとして使用されています。systemd は Gentoo で完全にサポートされており、その意図した目的に合うように動作します。もし systemd インストールパスについて、何かがハンドブックから欠けているようであれば、助けを求める _前に_ [systemd の記事](/wiki/Systemd/ja "Systemd/ja") を確認してください。

### multilib (32 ビットと 64 ビット)

**メモ**

すべてのアーキテクチャが multilib オプションを持っているわけではありません。多くは、ネイティブコードでしか動作しません。multilib が最もよく適用されているのは **amd64** です。

multilib プロファイルは、可能な場合は 64 ビットのライブラリを使用し、互換性のために厳密に必要な場合のみ 32 ビットのライブラリにフォールバックします。これは将来のカスタマイズのために大きな柔軟性をもたらすので、ほとんどのシステムにとってすばらしい選択肢となります。

**ヒント**

`multilib` ターゲットを使用すると、 `no-multilib` と比較して、後でプロファイルを切り換えるのが簡単になります。

### 非 multilib (64 ビットのみ)

**警告**

非 mutilib を必要とする明確な理由がなく、単に Gentoo を使いたいというケースでは、非 multilib を選択すべきでは _ありません_。 `no-multilib` から `mutilib` システムへの変更は、Gentoo の深い知識と低レベルのツールチェーンが必要です (これはおそらく私たちの [Toolchain developers](/wiki/Project:Toolchain "Project:Toolchain") を身震いさせるでしょう)。これは気の弱い人への警告ではなく、このガイドの範囲外になるということです。

システムのベースとして非 multilib の tarball を選択することで、32 ビットソフトウェアの無い、完全な 64 ビット環境を構築できます。これは事実上、 `multilib` プロファイルへの変更を (技術的には可能ではありますが) 面倒なものにします。

## stage ファイルをダウンロードする

_stage ファイル_ をダウンロードする前に、インストールのために使用されるマウントポイントにカレントディレクトリを設定すべきです:

`root #` `cd /mnt/gentoo`

### 日時を設定する

stage アーカイブは一般的には HTTPS を使用して取得され、これには比較的正確なシステム時刻の設定が必要です。時刻がずれているとダウンロードができないことがあるほか、インストール後に大きくシステム時刻を修正すると、予測不可能なエラーを発生させることがあります。

date を使って、現在の日付と時刻を検証することができます:

`root #` `date`

```
Mon Oct  3 13:16:22 UTC 2021

```

表示された日時が 2、3 分以上ずれている場合は、以下に示す方法のうちいずれかに従って更新してください。

#### 自動

時計のずれを修正するためには、手動でシステム時計を設定するよりも、典型的には [Network\_Time\_Protocol](/wiki/Network_Time_Protocol "Network Time Protocol") を使用するのがより簡単でより信頼できるでしょう。

[net-misc/chrony](https://packages.gentoo.org/packages/net-misc/chrony) の一部である chronyd は、システム時計を UTC と更新するために、次のようにして使用することができます:

`root #` `chronyd -q`

**重要**

動作するリアルタイムクロック (RTC) を持たないシステムは、システム開始のたびに、そしてその後は定期的に、システム時計を同期しなくてはなりません。RTC を持つシステムでも、バッテリが故障して時計のずれが蓄積することがあるため、これは有用です。

**警告**

標準的な NTP のトラフィックは認証されていません。ネットワークから取得した時刻データを検証することは重要です。

#### 手動

NTP へのアクセスが利用できない場合は、date を使用してシステム時計を手動で設定することができます。

**メモ**

すべての Linux システムでは UTC で時刻を設定することが推奨されます。あとでシステムのタイムゾーンを定義して、時刻を表示するときのオフセットを変更します。

時刻を設定するためには、以下の引数フォーマットが使用されます: `MMDDhhmmYYYY` 構文 ( **M** onth (月)、 **D** ay (日)、 **h** our (時)、 **m** inute (分)、 **Y** ear (年))。

例えば、2021年の10月3日 13時16分に設定するには、以下を実行してください:

`root #` `date 100313162021`

### グラフィカルブラウザ

完全なグラフィカルウェブブラウザがある環境を使っているひとには、stage ファイルの URL をメインウェブサイトの [ダウンロードセクション](https://www.gentoo.org/downloads/#other-arches) からコピーするのに何の問題も無いでしょう。単純に適切なタブを選択して、stage ファイルへのリンクを右クリックして、Copy Link してクリップボードにリンクをコピーして、コマンドライン上で wget ユーティリティにリンクをペーストして、stage ファイルをダウンロードします:

`root #` `wget <PASTED_STAGE_FILE_URL>`

### コマンドラインブラウザ

伝統的な読者や'古参'の Gentoo ユーザで、コマンドラインのみで作業をする人たちは、グラフィカル環境を必要としないメニュー形式のブラウザである links ([www-client/links](https://packages.gentoo.org/packages/www-client/links)) を使うほうを好むかもしれません。stage をダウンロードするために、Gentoo ミラーリストに飛んでください。

`root #` `links https://www.gentoo.org/downloads/mirrors/`

linksでHTTPプロキシを使うには、 `-http-proxy` オプションにプロキシのURLを渡してください。

`root #` `links -http-proxy proxy.server.com:8080 https://www.gentoo.org/downloads/mirrors/`

links に似た lynx ([www-client/lynx](https://packages.gentoo.org/packages/www-client/lynx)) というブラウザもあります。links と同じくグラフィカル環境を必要としませんが、メニューはありません。

`root #` `lynx https://www.gentoo.org/downloads/mirrors/`

プロキシを定義する必要があるならば、http\_proxyやftp\_proxy変数をexportしてください。

`root #` `export http_proxy="http://proxy.server.com:port"
`

`root #` `export ftp_proxy="http://proxy.server.com:port"`

ミラーリストから、近くのミラーを選んでください。通常はHTTPミラーで十分ですが、他のプロトコルも使えます。releases/alpha/autobuilds/ ディレクトリに移動してください。入手可能なすべてのstageファイルが列挙されています。ファイルは、サブアーキテクチャにちなんだ名前のサブティレクトリの中にあることもあります。ファイルを選び、 `d` を押してダウンロードしてください。

stage ファイルのダウンロードが完了したら、整合性を検証して stage ファイルのコンテンツが正当か確認することができます。興味のあるひとは [次節](/wiki/Handbook:Alpha/Installation/Stage/ja#Verifying_and_validating "Handbook:Alpha/Installation/Stage/ja") へ進んでください。

stage ファイルの検証と確認に興味が無いひとは、 `q` を押すことでコマンドラインブラウザを終了して、 [stage ファイルを展開する](/wiki/Handbook:Alpha/Installation/Stage/ja#Unpacking_the_stage_file "Handbook:Alpha/Installation/Stage/ja") 節へすぐに進むことができます。

### 検証して確認する

Minimal インストール CD のときと同じく、stage ファイルを検証して確認するためのファイルもダウンロードすることができます。これらの手順は飛ばしてもかまいませんが、ダウンロードしたファイルの整合性を気にするユーザのためにこれらのファイルが提供されています。これらの追加のファイルはミラーディレクトリのルートの下から取得できます。ハードウェアアーキテクチャとシステムプロファイルごとの適切な場所にブラウザで移動して、関連する .CONTENTS.gz、.DIGESTS、そして .sha256 ファイルをダウンロードしてください。

`root #` `wget https://distfiles.gentoo.org/releases/`

- .CONTENTS.gz \- stage ファイル内のファイル一覧を含む圧縮ファイル。
- .DIGESTS \- stage ファイルの、複数の暗号学的ハッシュアルゴリズムを用いたチェックサムを含みます。
- .sha256 \- stage ファイルの、PGP で署名された SHA256 チェックサムを含みます。このファイルはすべての stage ファイルで取得できるとは限りません。

ISO ファイルと同様に、tarball が改竄されていないことを確認するために、gpg を使って .tar.xz ファイルの電子署名を検証することもできます。

公式 Gentoo live イメージに関しては、自動化されたリリースのために [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release) パッケージが PGP 署名鍵を提供しています。鍵を検証に使用するためには、まずユーザのセッションに鍵をインポートする必要があります:

`root #` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

非公式 live イメージを使用していて、live 環境内で gpg と wget が提供されている場合は、先の [ダウンロードしたファイルを検証する](/wiki/Handbook:Alpha/Installation/Media/ja#Verifying_the_downloaded_files "Handbook:Alpha/Installation/Media/ja") の節を参照してください。

tarball の署名と、追加で関連するチェックサムファイルを検証してください:

`root #` `gpg --verify stage3-alpha-<release>-<init>.tar.xz.asc stage3-alpha-<release>-<init>.tar.xz
`

`root #` `gpg --output stage3-alpha-<release>-<init>.tar.xz.DIGESTS.verified --verify stage3-alpha-<release>-<init>.tar.xz.DIGESTS
`

`root #` `gpg --output stage3-alpha-<release>-<init>.tar.xz.sha256.verified --verify stage3-alpha-<release>-<init>.tar.xz.sha256
`

検証に成功した場合は、上のコマンドの出力に "Good signature from" が含まれるでしょう。

リリースメディアに署名するのに使用された OpenPGP 鍵のフィンガープリントは、 [リリースメディアの署名ページ](https://www.gentoo.org/downloads/signatures/) で確認できます。

**メモ**

今では、ほとんどの stage は init システムの種類に応じて明示的に (openrc または systemd と) [接尾辞](https://www.gentoo.org/news/2021/07/20/more-downloads.html) が付けられていますが、アーキテクチャによっては、これらがまだ無いものもあります。

openssl、sha256sum、または sha512sum などの暗号ツールおよびユーティリティを使用し、その出力を提供された .DIGESTS ファイルに含まれるチェックサムと比較することができます。

例えば、openssl で SHA512 チェックサムを検証するには:

`root #` `openssl dgst -r -sha512 stage3-alpha-<release>-<init>.tar.xz`

`dgst` は openssl コマンドにメッセージダイジェストのサブコマンドを使うように指示し、 `-r` はダイジェスト出力を coreutils フォーマットで印字し、 `-sha512` は SHA512 ダイジェストを選択します。

openssl で BLAKE2B512 チェックサムを検証するには:

`root #` `openssl dgst -r -blake2b512 stage3-alpha-<release>-<init>.tar.xz`

チェックサムコマンドの出力を、.DIGESTS.verified ファイルに含まれているハッシュとファイル名の値の組み合わせと比較してください。これらの値の組み合わせはチェックサムコマンドの出力と一致している必要があります。一致していない場合、ダウンロードしたファイルが破損しているため、削除して再ダウンロードするべきです。

sha256sum ユーティリティを使用して、関連する .sha256 ファイルに含まれる SHA256 ハッシュを検証するには:

`root #` `sha256sum --check stage3-alpha-<release>-<init>.tar.xz.sha256.verified`

`--check` オプションは sha256sum に期待されるファイルと対応するハッシュのリストを読み込ませ、各ファイルごとに、正しく計算できれば "OK" を、そうでない場合は "FAILED" を印字するように指示します。

## stage ファイルをインストールする

_stage ファイル_ のダウンロードと検証が完了したら、tar を使用してこれを展開することができます:

`root #` `tar xpvf stage3-*.tar.xz --xattrs-include='*.*' --numeric-owner -C /mnt/gentoo`

展開する前に、オプションを確認してください:

- `x` 展開 (e **x** tract)。アーカイブから内容を展開するように tar に指示します。
- `p` パーミッションを保持します ( **p** reserve permissions)。
- `v` 詳細な出力 ( **v** erbose output)。
- `f` ファイル ( **f** ile)。入力アーカイブの名前を tar に提供します。
- `--xattrs-include='*.*'` アーカイブに保存されている拡張属性をすべての名前空間について保持します。
- `--numeric-owner` (たとえ冒険的なユーザが公式 Gentoo live 環境を使わずにインストール作業をしている場合であっても) tarball から展開されるファイルのユーザ ID とグループ ID が Gentoo リリースエンジニアリングチームの意図通りになるようにします。
- `-C /mnt/gentoo` カレントディレクトリがどこかにかかわらず、ルートパーティションにファイルを展開します。

これでステージファイルは展開されました。この続きは [コンパイルオプションを設定する](/wiki/Handbook:Alpha/Installation/Stage/ja#Configuring_compile_options "Handbook:Alpha/Installation/Stage/ja") で。

## コンパイルオプションを設定する

### はじめに

システムを最適化するために、Portage（Gentooの公式なパッケージマネージャ）の挙動に影響する変数を設定できます。これらの変数はすべて環境変数として（exportを使って）設定できますが、export による設定は永続的なものではありません。

**メモ**

[シェル](/wiki/Shell/ja "Shell/ja") のプロファイルまたは rc ファイルを利用して変数を export することは技術的には可能ですが、これは基本的なシステム管理の方法としてはベストプラクティスではありません。

Portage は実行時に [make.conf](/wiki/Make.conf "Make.conf") を読み、ファイルに保存された値に基づいて実行時の振る舞いを変えます。make.conf は Portage の第一の設定ファイルと見ることができますので、その内容には注意して取り扱ってください。

**メモ**

/mnt/gentoo/usr/share/portage/config/make.conf.exampleに、すべての利用可能な変数のリストが、コメント付きで記載されています。make.conf についてのさらなるドキュメンテーションは、man -LC 5 make.conf を実行することで確認できます。(訳注: 日本語版は 10 年以上更新されていません。 `-LC` を付けて最新の英語版を参照してください。)

Gentoo のインストールを成功させるために最低限設定する必要がある変数は、以降で示すものだけです。

これから詳しく見ていく最適化変数を設定するために、エディタ（このガイドではnanoを使います）を起動してください。

`root #` `nano /mnt/gentoo/etc/portage/make.conf`

make.conf.example ファイルを読めば、記述形式は分かるでしょう。コメント行は `#` で始まり、他の行は「 `変数="内容"`」の形式で変数を定義します。これらの変数のうちのいくつかについて、次の節で見ていきます。

### CFLAGS と CXXFLAGS

CFLAGSとCXXFLAGS変数はそれぞれ、GCC CコンパイラとC++コンパイラのための最適化フラグを定義します。この2つの変数は通常ここで定義されますが、真に最高のパフォーマンスを発揮するためには、このフラグはプログラム毎に別々に設定する必要があるでしょう。すべてのプログラムは異なるからです。しかし、それでは管理が大変なので、make.confファイルでこれらのフラグを定義します。

make.confでは、一般にシステムの応答が速くなるように最適化フラグを設定するべきです。この変数に実験的な設定を書かないでください。過剰な最適化はプログラムの挙動をおかしくすることがあり、クラッシュや誤動作の元となります。

ハンドブックではすべての最適化オプションを説明することはしません。すべてを理解するためには、 [GNU オンラインマニュアル](https://gcc.gnu.org/onlinedocs/) や GCC info ページ (info gcc) を読んでください。make.conf.example ファイルにはたくさんの設定例と情報が含まれているので、これを読むこともお忘れなく。

最初の設定は `-march=` または `-mtune=` フラグです。これはターゲットアーキテクチャの名前を指定します。可能な選択肢はmake.conf.exampleファイル内にコメントとして書かれています。 _native_ を指定すると、コンパイラは（Gentooをインストールしようとしている）現在のシステムのアーキテクチャをターゲットとして選択してくれるので、よく使われます。

ふたつめの設定は `-O` フラグ（ゼロではなく大文字のオー）です。これはgcc最適化クラスフラグを指定します。可能なクラスは、s（サイズ最適化）、0（ゼロ、最適化無し）、1、2、3（速度最適化）です。速度最適化については、各クラスは1段階前のクラスが持つものと同じフラグに加えて、追加のフラグを持ちます。 `-O2` は推奨されるデフォルト設定です。 `-O3` をシステム全体で使うと問題を起こすことが知られているので、 `-O2` にとどめることをおすすめします。

他によく使われる最適化フラグには `-pipe` があります。これは、コンパイルステージ間での連絡方法として、一時ファイルではなくパイプを使うよう指定します。生成されるコードには影響しませんが、より多くのメモリを使うようになります。メモリの少ないシステムでは、gccが強制終了するかもしれません。そのような場合には、このフラグは使わないでください。

`-fomit-frame-pointer` を使うと、必要の無い場合にはフレームポインタをレジスタに保持しなくなります。これはアプリケーションのデバッグ時に深刻な影響を与えるかもしれません。

CFLAGS と CXXFLAGS 変数を定義するときには、最適化フラグは 1 つの文字列として結合してください。stage ファイルアーカイブに含まれるデフォルト値で十分でしょう。以下に例を示します:

コード **CFLAGSとCXXFLAGS変数の設定例**

```
# すべての言語において設定するコンパイラフラグ
COMMON_FLAGS="-mieee -pipe -O2 -mcpu=ev6"
# 同じ設定を両方の変数に使用
CFLAGS="${COMMON_FLAGS}"
CXXFLAGS="${COMMON_FLAGS}"

```

**ヒント**

各種コンパイルオプションがどのようにシステムに影響するかについては [GCC の最適化](/wiki/GCC_optimization/ja "GCC optimization/ja") の記事に詳しい情報がありますが、初心者がシステムの最適化を始めるには [Safe CFLAGS](/wiki/Safe_CFLAGS/ja "Safe CFLAGS/ja") の記事のほうがもっと実践的な場所かもしれません。

### RUSTFLAGS

現在では多くのプログラムが Rust で書かれていますが、Rust は最適化のための独自の方法を持っています。デフォルトでは、Rust はすべてのリリースビルドに対してレベル 3 の最適化を行います。ただし、プロジェクト側で変更されている場合もあり、その場合これはそのままにしておくべきです。Rust コンパイラに渡すことができる最適化の完全な一覧 (codegen として知られています) は、 [https://doc.rust-lang.org/rustc/codegen-options/index.html](https://doc.rust-lang.org/rustc/codegen-options/index.html) で利用可能です。

最も有用な最適化は、以下の例を使用して、使用中の CPU 向けにコンパイルするように Rust に指示することでしょう:

ファイル **`/etc/portage/make.conf`** **RUSTFLAGS の例**

```
RUSTFLAGS="${RUSTFLAGS} -C target-cpu=native"

```

Rust でサポートされている CPU の一覧を取得するには、以下を実行してください:

`root #` `rustc -C target-cpu=help`

このコマンドにより、デフォルトと、native でどの CPU が選択されるかが表示されるでしょう。

**メモ**

上のコマンドは、desktop stage3 tarball に対してか、[dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) または [dev-lang/rust](https://packages.gentoo.org/packages/dev-lang/rust) を emerge した後でのみ機能します。

### MAKEOPTS

MAKEOPTS 変数は、パッケージのインストール時にどれだけ並行してコンパイルを走らせるかを定義します。Portage バージョン 3.0.31[\[1\]](#cite_note-1) 時点において、未定義のままの場合、Portage のデフォルトの挙動では MAKEOPTS の jobs の値は nproc が返すスレッド数と同じ数に設定されます。

さらに Portage 3.0.53 の時点では[\[2\]](#cite_note-2)、未定義のままの場合、Portage のデフォルトの挙動では MAKEOPTS の load-average の値は nproc が返すスレッド数と同じ数に設定されます。

CPU のスレッド数か、システム全体の RAM 容量を 2 GiB で割った数のうち、小さい方を選択するのがよい選択とされています。

**警告**

ジョブ数を大きくすると、メモリ使用量にきわめて大きな影響を及ぼします。目安は、指定したジョブ数の各ジョブに対し、最低 2 GiB の RAM が割り当てられるようにすることです (つまり、例えば `-j6` は _最低でも_ 12 GiB を要求します)。メモリが枯渇しないようにするには、利用可能なメモリ容量に合うようにジョブ数を減らしてください。

**ヒント**

並列 emerge を使用する ( `--jobs`) と、実効的なジョブ数が指数関数的に (make ジョブ数 × emerge ジョブ数まで) 増大することがあります。これに対しては、localhost-only distcc 構成によって、ホスト当たりのコンパイラインスタンス数を制限することで対処することができます。

ファイル **`/etc/portage/make.conf`** **MAKEOPTS 設定例**

```
# 未定義のままの場合、Portage のデフォルトの挙動は以下の通りです:
# - MAKEOPTS の jobs 値を `nproc` が返すスレッド数と同じ数に設定する
# - MAKEOPTS の load-average 値を `nproc` が返すスレッド数よりわずかに大きい数に設定する (減衰された値であるため)
# '4' をシステムに応じて適切に (min(RAM/2GB, スレッド数) 置き換えるか、未設定のままにしておいてください。
MAKEOPTS="-j4 -l5"

```

さらなる詳細については man 5 make.conf 内で MAKEOPTS を検索してください。

### よーい、ドン！

好みの設定に合わせて /mnt/gentoo/etc/portage/make.conf を変更し、保存してください。nano では `Ctrl` + `o` で変更を保存して、 `Ctrl` + `x` で終了できます。

## 参照

1. [↑](#cite_ref-1)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2](https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2)
2. [↑](#cite_ref-2)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e](https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e)

[← ディスクの準備](/wiki/Handbook:Alpha/Installation/Disks/ja "Previous part") [Home](/wiki/Handbook:Alpha/ja "Handbook:Alpha/ja") [Gentooベースシステムのインストール →](/wiki/Handbook:Alpha/Installation/Base/ja "Next part")