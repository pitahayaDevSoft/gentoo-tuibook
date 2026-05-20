# About

## Contents

- [1はじめに](#.E3.81.AF.E3.81.98.E3.82.81.E3.81.AB)
  - [1.1ようこそ](#.E3.82.88.E3.81.86.E3.81.93.E3.81.9D)
  - [1.2インストール作業の順序](#.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB.E4.BD.9C.E6.A5.AD.E3.81.AE.E9.A0.86.E5.BA.8F)
  - [1.3インストールの選択肢について](#.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB.E3.81.AE.E9.81.B8.E6.8A.9E.E8.82.A2.E3.81.AB.E3.81.A4.E3.81.84.E3.81.A6)
  - [1.4トラブルがあったときは](#.E3.83.88.E3.83.A9.E3.83.96.E3.83.AB.E3.81.8C.E3.81.82.E3.81.A3.E3.81.9F.E3.81.A8.E3.81.8D.E3.81.AF)

## はじめに

### ようこそ

まずはじめに、Gentooへようこそ。あなたは、選択と性能の世界に入ろうとしています。Gentooでは、全てが選択です。そのことは、インストールするときにもたびたび明らかになります。コンパイルを何回やらせたいか、Gentooをどのようにインストールしたいか、システムのロギングツールに何を使用したいかなどなど、ユーザに選択権があります。

Gentooは、高速で現代的なメタディストリビューションで、整理された柔軟なデザインになっています。Gentooは、フリーソフトウェアで造られており、内部をユーザに隠匿することがありません。Gentooで用いるパッケージ管理システムであるPortageはPythonで書かれていますので、ソースコードを容易に閲覧したり改変したりすることができます。Gentooで採用しているパッケージ化の手法は、（コンパイル済パッケージにも対応していますが、）ソースコードを用いるものです。Gentooの設定は、一般的なテキストファイルで行います。つまり言い換えれば、Gentooは、全てをユーザーに開放しています。

Gentooに実行させることの選択肢を全員が理解していることは、きわめて重要です。望まざることをユーザーに強いないよう、我々は努めています。もしかりにこれに反するようなことがあったならば、 [バグリポート](https://bugs.gentoo.org) を送ってください。

### インストール作業の順序

The Gentoo Installation can be seen as a 10-step procedure, corresponding to the next set of chapters. Every step results in a certain state:

01. After step 1, the user is in a working environment ready to install Gentoo
02. After step 2, the Internet connection is ready to install Gentoo
03. After step 3, the hard disks are initialized to host the Gentoo installation
04. After step 4, the installation environment is prepared and the user is ready to chroot into the new environment
05. After step 5, core packages, which are the same on all Gentoo installations, are installed
06. After step 6, the Linux kernel is installed
07. After step 7, the user will have configured most of the Gentoo system configuration files
08. After step 8, the necessary system tools are installed
09. After step 9, the proper boot loader has been installed and configured
10. After step 10, the freshly installed Gentoo Linux environment is ready to be explored

このハンドブックでは、一定の選択肢を提示したときには必ず、賛否両論の併記に努めます。デフォルトの選択肢で進める記載をした際にも（見出しに「デフォルト:」と記載）、他に取りうる選択肢も記載します（見出しに「代替案:」と記載）。決して、「デフォルトは Gentoo のお勧めだ」と考えないでください。デフォルトはあくまでも、多くのユーザーが利用すると思われる選択肢にすぎません。

ときには、追加可能な手順が続くことがあります。そのような手順は「追加可能:」と記載します。つまりこの手順は、Gentoo のインストール自体には必須ではありません。とはいえ、以前にした決断によっては必須になる追加手順もあります。その際には、その追加手順の説明の直前に、この旨を明記するとともに、原因となった決断をした時期も記載します。

### インストールの選択肢について

Gentooは、多くの異なる方法でインストールすることができます。 ARM64のハードウェアの場合、一般的にこれはEMMC、SD、またはUSDメディアにDD'ingのためのシステムイメージをダウンロードすることを含みます。

This document covers the installation specifically for ARM64 hardware. We try to cover the specifics for known hardware that we've been able to test and validate. ARM64 hardware not listed will probably work however you'll need to understand what is needed for your hardware and adjust your installation accordingly.

**メモ**

Gentoo 以外のCDを用いる場合などの、ほかのインストール方法については、 [『Alternative Installation Guide』](/wiki/Installation_alternatives/ja "Installation alternatives/ja") を読んでください。

また、我々が提供している [『Gentoo Installation Tips & Tricks』](/wiki/Gentoo_installation_tips_and_tricks/ja "Gentoo installation tips and tricks/ja") という文書も役にたつかもしれません。

### トラブルがあったときは

インストール中に (またはインストール文書に) 何か問題を見つけたら、 [バグトラッキングシステム](https://bugs.gentoo.org) で既知のバグとして報告されていないかどうか、確認してみてください。 もし無いようであれば、私たちが対応できるように、その問題をバグ報告してください。 その (あなたが報告した) バグを担当する開発者たちを恐れないでください。取って喰われるようなことは (滅多に) ありませんから。

あなたが今読んでいる文書は、特定のアーキテクチャ向けということになっていますが、 他のアーキテクチャの情報も、その中に紛れ込んでしまっているかもしれない、ということを一応、先に言っておきます。これはGentooハンドブックの多くの部分が、全てのアーキテクチャに共通のソースコードを使用していることに因ります (重複作業を減らして開発リソースを節約するため)。混乱しないように、このような部分は最小限に抑えていきたいと思っています。

その問題が、ユーザーの問題 (文書をよく読んだにもかかわらず起きたあなたのミス) なのか、ソフトウェアの問題 (インストール/文書をよくテストしたにもかかわらず起きた私たちのミス) なのか、はっきりしないときには、IRCチャンネルの[#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo))に気軽に参加してみてください。そんなときじゃなくても全然かまわないんですけどね。

そういえば、Gentooについて何か分からないことがあったら、 [Gentoo Wiki](/wiki/Main_Page "Main Page") から利用できる [よくある質問](/wiki/FAQ/ja "FAQ/ja") を見てみてください。 [Gentoo Forums](https://forums.gentoo.org) 上にある [FAQs](https://forums.gentoo.org/viewforum.php?f=40) もあります。

[←](/index.php?title=Handbook:ARM64/ja&action=edit&redlink=1 "Previous part") [Home](/index.php?title=Handbook:ARM64/ja&action=edit&redlink=1 "Handbook:ARM64/ja (page does not exist)") [→](/index.php?title=Handbook:ARM64/Installation/Media/ja&action=edit&redlink=1 "Next part")