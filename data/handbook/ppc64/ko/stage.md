# Stage

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Stage/de "Handbuch:PPC64/Installation/Stage (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Stage "Handbook:PPC64/Installation/Stage (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Stage/es "Manual de Gentoo: PPC64/Instalación/Stage (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Stage/fr "Handbook:PPC64/Installation/Stage/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Stage/it "Handbook:PPC64/Installation/Stage/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Stage/hu "Handbook:PPC64/Installation/Stage/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Stage/pl "Handbook:PPC64/Installation/Stage (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Stage/pt-br "Handbook:PPC64/Installation/Stage/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Stage/ru "Handbook:PPC64/Installation/Stage (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Stage/ta "கையேடு:PPC64/நிறுவல்/நிலை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Stage/zh-cn "手册：PPC64/安装/安装stage3 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Stage/ja "ハンドブック:PPC64/インストール/Stage (100% translated)")
- 한국어

[PPC64 핸드북](/wiki/Handbook:PPC64/ko "Handbook:PPC64/ko")[설치](/wiki/Handbook:PPC64/Full/Installation/ko "Handbook:PPC64/Full/Installation/ko")[설치 정보](/wiki/Handbook:PPC64/Installation/About/ko "Handbook:PPC64/Installation/About/ko")[매체 선택](/wiki/Handbook:PPC64/Installation/Media/ko "Handbook:PPC64/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:PPC64/Installation/Networking/ko "Handbook:PPC64/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:PPC64/Installation/Disks/ko "Handbook:PPC64/Installation/Disks/ko")스테이지 3 설치[베이스 시스템 설치](/wiki/Handbook:PPC64/Installation/Base/ko "Handbook:PPC64/Installation/Base/ko")[커널 설정](/wiki/Handbook:PPC64/Installation/Kernel/ko "Handbook:PPC64/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:PPC64/Installation/System/ko "Handbook:PPC64/Installation/System/ko")[도구 설치](/wiki/Handbook:PPC64/Installation/Tools/ko "Handbook:PPC64/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko")[마무리](/wiki/Handbook:PPC64/Installation/Finalizing/ko "Handbook:PPC64/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:PPC64/Full/Working/ko "Handbook:PPC64/Full/Working/ko")[포티지 소개](/wiki/Handbook:PPC64/Working/Portage/ko "Handbook:PPC64/Working/Portage/ko")[USE 플래그](/wiki/Handbook:PPC64/Working/USE/ko "Handbook:PPC64/Working/USE/ko")[포티지 기능](/wiki/Handbook:PPC64/Working/Features/ko "Handbook:PPC64/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:PPC64/Working/Initscripts/ko "Handbook:PPC64/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:PPC64/Working/EnvVar/ko "Handbook:PPC64/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:PPC64/Full/Portage/ko "Handbook:PPC64/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:PPC64/Portage/Files/ko "Handbook:PPC64/Portage/Files/ko")[변수](/wiki/Handbook:PPC64/Portage/Variables/ko "Handbook:PPC64/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:PPC64/Portage/Branches/ko "Handbook:PPC64/Portage/Branches/ko")[추가 도구](/wiki/Handbook:PPC64/Portage/Tools/ko "Handbook:PPC64/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:PPC64/Portage/CustomTree/ko "Handbook:PPC64/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:PPC64/Portage/Advanced/ko "Handbook:PPC64/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:PPC64/Full/Networking/ko "Handbook:PPC64/Full/Networking/ko")[시작하기](/wiki/Handbook:PPC64/Networking/Introduction/ko "Handbook:PPC64/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:PPC64/Networking/Advanced/ko "Handbook:PPC64/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:PPC64/Networking/Modular/ko "Handbook:PPC64/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:PPC64/Networking/Wireless/ko "Handbook:PPC64/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:PPC64/Networking/Extending/ko "Handbook:PPC64/Networking/Extending/ko")[동적 관리](/wiki/Handbook:PPC64/Networking/Dynamic/ko "Handbook:PPC64/Networking/Dynamic/ko")

## Contents

- [1스테이지 타르볼 선택](#.EC.8A.A4.ED.85.8C.EC.9D.B4.EC.A7.80_.ED.83.80.EB.A5.B4.EB.B3.BC_.EC.84.A0.ED.83.9D)
- [2OpenRC](#OpenRC)
- [3systemd](#systemd)
  - [3.1Multilib (32비트 및 64비트)](#Multilib_.2832.EB.B9.84.ED.8A.B8_.EB.B0.8F_64.EB.B9.84.ED.8A.B8.29)
  - [3.2no-multilib(순수 64비트)](#no-multilib.28.EC.88.9C.EC.88.98_64.EB.B9.84.ED.8A.B8.29)
- [4스테이지 타르볼 다운로드](#.EC.8A.A4.ED.85.8C.EC.9D.B4.EC.A7.80_.ED.83.80.EB.A5.B4.EB.B3.BC_.EB.8B.A4.EC.9A.B4.EB.A1.9C.EB.93.9C)
- [5날짜 및 시간 설정](#.EB.82.A0.EC.A7.9C_.EB.B0.8F_.EC.8B.9C.EA.B0.84_.EC.84.A4.EC.A0.95)
  - [5.1자동](#.EC.9E.90.EB.8F.99)
  - [5.2수동](#.EC.88.98.EB.8F.99)
  - [5.3그래픽 브라우저](#.EA.B7.B8.EB.9E.98.ED.94.BD_.EB.B8.8C.EB.9D.BC.EC.9A.B0.EC.A0.80)
  - [5.4명령행 브라우저](#.EB.AA.85.EB.A0.B9.ED.96.89_.EB.B8.8C.EB.9D.BC.EC.9A.B0.EC.A0.80)
  - [5.5검증 및 유효화](#.EA.B2.80.EC.A6.9D_.EB.B0.8F_.EC.9C.A0.ED.9A.A8.ED.99.94)
- [6스테이지 타르볼 설치하기](#.EC.8A.A4.ED.85.8C.EC.9D.B4.EC.A7.80_.ED.83.80.EB.A5.B4.EB.B3.BC_.EC.84.A4.EC.B9.98.ED.95.98.EA.B8.B0)
- [7컴파일 옵션 설정](#.EC.BB.B4.ED.8C.8C.EC.9D.BC_.EC.98.B5.EC.85.98_.EC.84.A4.EC.A0.95)
  - [7.1도입부](#.EB.8F.84.EC.9E.85.EB.B6.80)
  - [7.2CFLAGS와 CXXFLAGS](#CFLAGS.EC.99.80_CXXFLAGS)
  - [7.3RUSTFLAGS](#RUSTFLAGS)
  - [7.4MAKEOPTS](#MAKEOPTS)
  - [7.5준비, 시, 작!](#.EC.A4.80.EB.B9.84.2C_.EC.8B.9C.2C_.EC.9E.91.21)
- [8References](#References)

### 스테이지 타르볼 선택

**요령**

On supported architectures, it is recommended for users targeting a desktop (graphical) operating system environment to use a stage file with the term `desktop` within the name. These files include packages such as [llvm-core/llvm](https://packages.gentoo.org/packages/llvm-core/llvm) and [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) and USE flag tuning which will greatly improve install time.

The [stage file](/wiki/Stage_file "Stage file") acts as the seed of a Gentoo install. Stage files are generated by the [Release Engineering Team](/wiki/Project:RelEng "Project:RelEng") with [Catalyst](/wiki/Catalyst "Catalyst"). Stage files are based on specific [profiles](/wiki/Profile_(Portage) "Profile (Portage)"), and contain an almost-complete system.

When choosing a stage file, it's important to pick one with profile targets corresponding to the desired system type.

**중요**

While it's possible to make major profile changes after an installation has been established, switching requires substantial effort and consideration, and is outside the scope of this installation manual. Switching init systems is difficult, but switching from `no-multilib` to `multilib` requires extensive Gentoo and low-level toolchain knowledge.

대부분 사용자는 '고급' 타르볼 선택 항목을 활용하면 안됩니다. 이 항목은 특정 프로그램 또는 하드웨어 설정용으로 만들었습니다.

### OpenRC

[OpenRC](/wiki/OpenRC "OpenRC") is a dependency-based init system (responsible for starting up system services once the kernel has booted) that maintains compatibility with the system provided init program, normally located in /sbin/init. It is Gentoo's native and original init system, but is also deployed by a few other Linux distributions and BSD systems.

### systemd

systemd is a modern SysV-style init and rc replacement for Linux systems. It is used as the primary init system by a majority of Linux distributions. systemd is fully supported in Gentoo and works for its intended purpose. If something seems lacking in the Handbook for a systemd install path, review the [systemd article](/wiki/Systemd/ko "Systemd/ko") _before_ asking for support.

#### Multilib (32비트 및 64비트)

**참고**

Not every architecture has a multilib option. Many only run with native code. Multilib is most commonly applied to **amd64**.

시스텝 기본 타르볼을 선택하면, 특히 [시스템 프로파일 선택](/wiki/Handbook:PPC64/Installation/Base/ko#.EC.A0.81.EC.A0.88.ED.95.9C_.ED.94.84.EB.A1.9C.ED.8C.8C.EC.9D.BC_.EC.84.A0.ED.83.9D "Handbook:PPC64/Installation/Base/ko") 을 진행할 때, 설치 과정의 나머지 시간을 절약할 수 있습니다. 스테이지 타르볼 선택은 앞으로의 시스템 설정에 직접적인 영향을 주며 이후에 발생할 두통을 예방할 수 있습니다. multilib 타르볼은 가능하면 64비트 라이브러리를 사용하며 호환성이 필요하면 32비트 버전을 대신 사용합니다. 주요 설치 과정에 있어 최상의 옵션이며, 앞으로 다양한 개별 설정을 유연하게 처리할 수 있기 때문입니다. 프로파일을 쉽게 바꿀 수 있는 시스템을 원한다면 프로세서 아키텍처에 해당하는 multilib 타르볼 옵션을 다운로드해야합니다.

**요령**

Using `multilib` targets makes it easier to switch profiles later, compared to `no-multilib`

#### no-multilib(순수 64비트)

**경고**

no-multilib 시스템에서 multilib 시스템으로 이전할 경우 젠투 동작과 하부 단계 툴체인에 대한 해박한 지식이 필요합니다(게다가 이 문제는 [툴체인 개발자](/wiki/Project:Toolchain "Project:Toolchain") 들 마저도 이를 갈게 만듭니다). 비굴해서 그런게 아니며 이 안내서에서 다룰 내용의 범위를 벗어납니다.

no-multilib 타르볼 시스템 베이스로 선택하면 완벽한 64비트 처리 시스템 환경을 갖춥니다. 이 환경은 multilib 프로파일로 전환하는 희한하지만 가능한 기능을 갖추고 있습니다. **정말 필요한 경우가 아닌 이상 처음 젠투를 사용하는 사람이라면 no-multilib 타르볼을 젠투에서 사용하지 않는게 좋습니다.**

### 스테이지 타르볼 다운로드

Before downloading the _stage file_, the current directory should be set to the location of the mount used for the install:

`root #` `cd /mnt/gentoo`

### 날짜 및 시간 설정

Stage archives are generally obtained using HTTPS which requires relatively accurate system time. Clock skew can prevent downloads from working, and can cause unpredictable errors if the system time is adjusted by any considerable amount after installation.

date 명령으로 현재 날짜와 시각을 확인하십시오:

`root #` `date`

```
Mon Oct  3 13:16:22 PDT 2016

```

날짜 시각이 잘못 나타났다면 아래 방식으로 고치십시오.

#### 자동

Using [Network\_Time\_Protocol](/wiki/Network_Time_Protocol "Network Time Protocol") to correct clock skew is typically easier and more reliable than manually setting the system clock.

chronyd, part of [net-misc/chrony](https://packages.gentoo.org/packages/net-misc/chrony) can be used to update the system clock to UTC with:

`root #` `ntpd -q -g`

**중요**

Systems without a functioning Real-Time Clock (RTC) must sync the system clock at every system start, and on regular intervals thereafter. This is also beneficial for systems with a RTC, as the battery could fail, and clock skew can accumulate.

**경고**

Standard NTP traffic is not authenticated, it is important to verify time data obtained from the network.

#### 수동

When NTP access is unavailable, date can be used to manually set the system clock.

모든 리눅스 시스템에 UTC 시계 설정을 추천합니다. 시간대는 설치 과정 후반에 설정합니다. 시간대를 설정하면 시계를 지역 시간으로 나타냅니다.

date 명령으로 시스템 시계를 직접 설정할 수도 있습니다. `MMDDhhmmYYYY` 형식으로 설정합니다(월,일,시,분,연도).

예를 들어 2016년 10월 03일 13시 16분을 설정하려면:

`root #` `date 100313162016`

#### 그래픽 브라우저

완벽한 그래픽 여건을 갖춘 웹 브라우저에서 주 웹사이트 [다운로드](https://www.gentoo.org/downloads/#other-arches) 페이지에서 스테이지 파일 URL을 그대로 복사해서 쓰는데 아무런 문제가 없습니다. 간단히 적절한 탭을 선택한 다음, 스테이지 파일 링크 위에 마우스 커서를 가져간 후, 마우스 오른쪽 단추를 누르고 링크 주소 복사 (Firefox) 또는 링크 위치 복사 (Chromium)를 선택하여 클립보드에 링크를 복사합니다. 그 후 명령행에서 wget 유틸리티 매개변수 자리에 붙여넣어 스테이지 타르볼을 내려받으십시오:

`root #` `wget <PASTED_STAGE_URL>`

#### 명령행 브라우저

좀 더 예전부터 사용해온 독자 또는 젠투 사용자라면, 그래픽 방식이 아닌 메뉴기반 브라우저 links를 대신 선호할 지도 모르겠습니다. 스테이지 파일을 다운로드하려면 다음 명령으로 젠투 미러를 찾아보십시오:

`root #` `links https://www.gentoo.org/downloads/mirrors/`

links에서 HTTP 프록시를 사용하려면, `-http-proxy` 옵션으로 URL을 전달하십시오:

`root #` `links -http-proxy proxy.server.com:8080 https://www.gentoo.org/downloads/mirrors/`

links 다음에는 lynx 브라우저도 있습니다. links와 유사하게 비-그래픽 브라우저지만, 메뉴 기반 브라우저도 아닙니다.

`root #` `lynx https://www.gentoo.org/downloads/mirrors/`

프록시를 지정해야 한다면, http\_proxy 또는 ftp\_proxy 변수 값을 export로 처리하십시오:

`root #` `export http_proxy="http://proxy.server.com:port"
`

`root #` `export ftp_proxy="http://proxy.server.com:port"`

미러 목록에서 가까운 미러를 선택하십시오. HTTP 미러를 사용하는 걸로 충분합니다만, 다른 프로토콜로도 쓸 수 있습니다. releases/ppc64/autobuilds/ 디렉터리로 이동하십시오. 존재하는 모든 스테이지 파일이 나타납니다(아마도 각각의 하위 아키텍처에 있는 하위 디렉터리에 있을지도 모릅니다). 그 중 하나를 선택하고 `d` 를 눌러 다운로드하십시오.

스테이지 파일 다운로드가 끝나면 스테이지 타르볼 내용의 무결성을 검증하고 유효화할 수 있습니다. 이 부분은 [다음 장](#.EA.B2.80.EC.A6.9D_.EB.B0.8F_.EC.9C.A0.ED.9A.A8.ED.99.94) 에서 진행합니다.

스테이지 파일 검증 및 유효화에 관심 없는 분들은 `q` 키를 눌러 명령행 브라우저를 닫고 바로 [스테이지 타르볼 압축 해제](#.EC.8A.A4.ED.85.8C.EC.9D.B4.EC.A7.80_.ED.83.80.EB.A5.B4.EB.B3.BC_.EC.95.95.EC.B6.95_.ED.95.B4.EC.A0.9C) 부분으로 넘어갈 수 있습니다.

#### 검증 및 유효화

최소 설치 CD 처럼, 추가로 검증하고 유효화할 다운로드 파일이 있습니다. 이 단계는 건너뛸 수 있지만 그냥 다운로드한 파일의 무결성에 신경 쓰는 사용자를 위해 제공합니다.

`root #` `wget https://distfiles.gentoo.org/releases/`

- 스테이지 타르볼 파일 목록이 있는 .CONTENTS 파일.
- 각각의 알고리즘으로 만든 스테이지 파일의 체크섬이 있는 .DIGESTS 파일.
- .DIGESTS 파일과 마찬가지로 각각의 알고리즘으로 만든 스테이지 파일의 체크섬이 있지만, 젠투 프로젝트에서 제공했음을 확인할때 쓰는 암호화 서명도 들어있는 .DIGESTS.asc.

ISO 파일과 마찬가지로, gpg를 활용하여 .DIGESTS.asc 파일의 암호화 서명을 검증하여 누군가가 체크섬에 손을 댔는지 여부를 확인할 수 있습니다:

For official Gentoo live images, the [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release) package provides PGP signing keys for automated releases. The keys must first be imported into the user's session in order to be used for verification:

`root #` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

For all non-official live images which offer gpg and wget in the live environment, refer to the earlier [Verifying the downloaded files](/wiki/Handbook:PPC64/Installation/Media/ko#Verifying_the_downloaded_files "Handbook:PPC64/Installation/Media/ko") section.

Verify the signature of the tarball and, optionally, associated checksum files:

`root #` `gpg --verify stage3-ppc64-<release>.tar.bz2,.DIGESTS.asc`

If verification succeeds, "Good signature from" will be in the output of the previous command(s).

The fingerprints of the OpenPGP keys used for signing release media can be found on the [release media signatures page](https://www.gentoo.org/downloads/signatures/).

**참고**

Most stages are now explicitly [suffixed](https://www.gentoo.org/news/2021/07/20/more-downloads.html) with the init system type (openrc or systemd), although some architectures may still be missing these for now.

openssl을 사용하여 .DIGESTS 또는 .DIGESTS.asc 파일에서 제공하는 체크섬 출력을 비교하십시오.

SHA512 체크섬을 검증한다면:

`root #` `openssl dgst -r -sha512 stage3-ppc64-<release>.tar.bz2`

`dgst` instructs the openssl command to use the Message Digest sub-command, `-r` prints the digest output in coreutils format, and `-sha512` selects the SHA512 digest.

월풀 체크섬을 검증하려면:

`root #` `openssl dgst -r -whirlpool stage3-ppc64-<release>.tar.bz2`

.DIGESTS(.asc) 파일에 등록한 값을 이 명령의 출력과 비교하십시오. 값이 일치해야 하며, 그렇지 않으면 다운로드한 파일(또는 digests 파일)이 깨진 상태입니다.

sha512sum 명령을 사용하는 다른 방법도 있습니다:

`root #` `sha512sum stage3-ppc64-<release>.tar.bz2`

The `--check` option instructs sha256sum to read a list of expected files and associated hashes, and then print an associated "OK" for each file that calculates correctly or a "FAILED" for files that do not.

## 스테이지 타르볼 설치하기

이제 다운로드한 스테이지를 시스템에 압축해제하십시오. tar를 사용하여 진행하겠습니다:

`root #` `tar xpvf stage3-*.tar.bz2 --xattrs-include='*.*' --numeric-owner`

동일한 옵션( `xpf` 와 `--xattrs-include='*.*'`)을 사용했는지 확인하십시오. `x` 는 추출, `p` 는 권한 플래그 유지, `f` 는 표준 출력이 아닌 파일로 추출함을 나타냅니다. `--xattrs-include='*.*'` 는 압축 파일에 저장한 이름 영역 전체의 확장 속성을 포함합니다. 마지막으로 `--numeric-owner` 옵션은 얼리어답터가 공식 젠투 설치 미디어를 활용하지 않는다 할 지라도, 젠투 릴리즈 엔지니어링 팀이 의도한 대로 타르볼에서 풀려나온 파일의 사용자 ID와 그룹 ID가 동일하게 나왔는지 확인합니다.

- `x` e **x** tract, instructs tar to extract the contents of the archive.
- `p` **p** reserve permissions.
- `v` **v** erbose output.
- `f` **f** ile, provides tar with the name of the input archive.
- `--xattrs-include='*.*'` Preserves extended attributes in all namespaces stored in the archive.
- `--numeric-owner` Ensure that the user and group IDs of files being extracted from the tarball remain the same as Gentoo's release engineering team intended (even if adventurous users are not using official Gentoo live environments for the installation process).
- `-C /mnt/gentoo` Extract files to the root partition regardless of the current directory.

이제 스테이지 파일의 압축을 해제했으니, [컴파일 옵션 설정](#.EC.BB.B4.ED.8C.8C.EC.9D.BC_.EC.98.B5.EC.85.98_.EC.84.A4.EC.A0.95) 으로 진행하십시오.

## 컴파일 옵션 설정

### 도입부

젠투를 최적화 하려는 목적으로 젠투에서 공식적으로 지원하는 꾸러미 관리자 포티지 동작에 영향을 줄 여러가지 변수를 설정할 수 있습니다. 이들 변수는 (export로) 환경 변수처럼 설정할 수 있습니다만 언제든 값이 유지되는 것은 아닙니다. 설정값을 유지하려, 포티지의 설정 파일 /etc/portage/make.conf 파일을 포티지에서 읽습니다.

**참고**

Technically variables can be exported via the [shell's](/wiki/Shell "Shell") profile or rc files, however that is not best practice for basic system administration.

Portage reads in the [make.conf](/wiki/Make.conf "Make.conf") file when it runs, which will change runtime behavior depending on the values saved in the file. make.conf can be considered the primary configuration file for Portage, so treat its content carefully.

**참고**

쓸 수 있도록 주석 처리하여 준비한 모든 변수 목록은 /mnt/gentoo/usr/share/portage/config/make.conf.example에 있습니다. 젠투 설치를 제대로 하기위해 설정이 필요한 변수만을 아래에 언급했습니다.

For a successful Gentoo installation only the variables that are mentioned below need to be set.}}

편집기를 실행(이 안내서에서는 nano를 사용합니다)하여 이 다음에 언급할 최적화 변수값을 바꾸어보겠습니다.

`root #` `nano -w /mnt/gentoo/etc/portage/make.conf`

make.conf.example 파일에서 파일을 어떤 식으로 구성해야 하는지 분명히 나타납니다: "#"(으)로 시작하는 줄은 주석이며, 다른 줄은 VARIABLE="content" 문법으로 작성한 변수 설정 부분입니다. 다양한 이들 변수에 대해서는 다음에 이야기하겠습니다.

### CFLAGS와 CXXFLAGS

CFLAGS와 CXXFLAGS 변수는 gcc C/C++ 컴파일러의 최적화 플래그를 각각 지정합니다. 보통 여기에 지정하지만, 최적의 성능을 위해서는 각각의 프로그램에 플래그를 최적화해야합니다. 각각의 프로그램이 다르기 때문입니다. 그러나 그리 관리하기 쉬운게 아니므로 이 플래그 정의를 make.conf 파일에 다룹니다.

make.conf 에서는 보통 시스템에 가장 많이 영향을 줄 최적화 플래그를 지정해야합니다. 이 변수에 시험적인 설정은 넣지 마십시오. 최적화를 과도하게 하면 프로그램 동작이 잘못되는 수가 있습니다(깨지거나, 잘못되거나, 기능이 망가지거나).

가능한 모든 최적화 옵션을 설명하지는 않겠습니다. 이들을 전부 이해하려면 [GNU 온라인 문서](https://gcc.gnu.org/onlinedocs/) en 또는 gcc 정보 페이지(info gcc \- 리눅스 시스템에서만 동작)를 참고하십시오. make.conf.example 파일 자체에 상당한 양의 예제와 정보를 담고 있습니다. 이것 또한 잊지 말고 살펴보십시오.

첫번째 설정은 대상 아키텍처 이름을 지정하는 `-march=` 또는 `-mtune=` 플래그입니다. 사용할 수 있는 옵션은 make.conf.example 파일에 (주석으로) 들어있습니다. 보통 사용하는 값은 컴파일러가 대상 아키텍처를 (사용자가 젠투를 설치하려는) 현재 시스템으로 설정하도록 하는 _native_ 값입니다.

두번째는 gcc 최적화 수준 플래그를 지정하는 `-O` 플래그(숫자 영이 아닌 대문자 O임) 입니다. 가능한 클래스는 s(크기 최적화), 0(영. 최적화 안함), 1, 2, 또는 속도 최적화 를 위한 3 플래그(모든 클래스는 이전 클래스와 비슷하지만, 몇가지 특징을 추가함)입니다. 기본적으로 `-O2` 를 추천합니다. 시스템 전반적인 영역에 있어 `-O3` 이 문제를 일으키는것으로 알려져 있어 `-O2` 에 집착하기를 추천합니다.

다른 최적화 플래그는 `-pipe`(다중 스테이지 컴파일간 통신에 임시 파일을 쓰는 대신 파이프를 활용)입니다. 생성 코드에 영향을 주지는 않지만 더 많은 메모리를 사용합니다. 메모리가 부족해지면, gcc를 강제로 끝냅니다. 이 경우 이 플래그를 사용하지 마십시오.

`-fomit-frame-pointer`(필요하지 않은 함수에 대한 프레임 포인터를 레지스터에서 계속 가지고 있지 않도록 하는 옵션)를 사용하면 프로그램을 디버깅하는동안 심각한 문제가 생길지도 모릅니다.

CFLAGS와 CXXFLAGS 변수를 지정하면, 각각의 최적화 플래그를 하나의 문자열로 합칩니다. 스테이지 3 아카이브에 들어있는 기본값은 풀려나온 값 자체로도 충분합니다. 다음 플래그는 예제일뿐입니다:

코드 **CFLAGS와 CXXFLAGS 변수 예제**

```
CFLAGS="-O2 -pipe"
# Use the same settings for both variables
CXXFLAGS="${CFLAGS}"

```

**요령**

[GCC 최적화 안내서](/wiki/GCC_optimization/ko "GCC optimization/ko") 에 다양한 컴파일 옵션이 시스템에 어떻게 영향을 주는지 많은 설명을 넣었지만, 시스템 최적화를 시작하려는 초보자에게는 [Safe CFLAGS](/wiki/Safe_CFLAGS "Safe CFLAGS")(en)가 더 도움이 될 수도 있습니다.

### RUSTFLAGS

Many programs are now written in Rust which has its own way of optimising. By default Rust optimizes to level 3 on all release builds unless a project changes it so this should be left as is. A full optimization list (known as codegen) that can be passed to the rust compiler is available at [https://doc.rust-lang.org/rustc/codegen-options/index.html](https://doc.rust-lang.org/rustc/codegen-options/index.html).

The most useful optimization would be to tell Rust to compile for your CPU using the following example:

파일 **`/etc/portage/make.conf`** **RUSTFLAGS Example**

```
RUSTFLAGS="${RUSTFLAGS} -C target-cpu=native"

```

To get a list of supported CPUs in Rust run:

`root #` `rustc -C target-cpu=help`

This will show what the default and also which CPU will be selected with native.

**참고**

The above command only works on desktop stage3 tarballs or after emerging [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) or [dev-lang/rust](https://packages.gentoo.org/packages/dev-lang/rust).

### MAKEOPTS

MAKEOPTS 변수는 꾸러미를 설치하는 동안 컴파일을 동시에 몇개를 진행하는지 지정합니다. 최적의 값은 시스템에 붙은 CPU(또는 CPU 코어)의 갯수에 1을 더한 값이지만 이 안내서가 언제나 완벽하진 않습니다.

Further, as of Portage 3.0.53[\[1\]](#cite_note-1), if left undefined, Portage's default behavior is to set the MAKEOPTS load-average value to the same number of threads returned by nproc.

A good choice is the smaller of: the number of threads the CPU has, or the total amount of system RAM divided by 2 GiB.

**경고**

Using a large number of jobs can significantly impact memory consumption. A good recommendation is to have at least 2 GiB of RAM for every job specified (so, e.g. `-j6` requires _at least_ 12 GiB). To avoid running out of memory, lower the number of jobs to fit the available memory.

**요령**

When using parallel emerges ( `--jobs`), the effective number of jobs run can grow exponentially (up to make jobs multiplied by emerge jobs). This can be worked around by running a localhost-only distcc configuration that will limit the number of compiler instances per host.

코드 **make.conf의 MAKEOPTS 선언 예제**

```
MAKEOPTS="-j2"

```

Search for MAKEOPTS in man 5 make.conf for more details.

### 준비, 시, 작!

개인 취향에 맞춰 /mnt/gentoo/etc/portage/make.conf를 업데이트한 후 저장하십시오(나노 사용자는 `Ctrl` + `X` 를 치십시오).

## References

1. [↑](#cite_ref-1)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e](https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e)

[← 디스크 준비](/wiki/Handbook:PPC64/Installation/Disks/ko "이전 내용") [처음](/wiki/Handbook:PPC64/ko "Handbook:PPC64/ko") [젠투 베이스 시스템 설치 →](/wiki/Handbook:PPC64/Installation/Base/ko "다음 내용")