# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Bootloader/de "Handbuch:MIPS/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Bootloader "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Bootloader/es "Manual de Gentoo: MIPS/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Bootloader/fr "Handbook:MIPS/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Bootloader/it "Handbook:MIPS/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Bootloader/pl "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Bootloader/pt-br "Manual:MIPS/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Bootloader/ru "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Bootloader/ta "கையேடு:MIPS/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Bootloader/zh-cn "手册：MIPS/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Bootloader/ja "ハンドブック:MIPS/インストール/ブートローダー (100% translated)")
- 한국어

[MIPS 핸드북](/wiki/Handbook:MIPS/ko "Handbook:MIPS/ko")[설치](/wiki/Handbook:MIPS/Full/Installation/ko "Handbook:MIPS/Full/Installation/ko")[설치 정보](/wiki/Handbook:MIPS/Installation/About/ko "Handbook:MIPS/Installation/About/ko")[매체 선택](/wiki/Handbook:MIPS/Installation/Media/ko "Handbook:MIPS/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:MIPS/Installation/Networking/ko "Handbook:MIPS/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:MIPS/Installation/Disks/ko "Handbook:MIPS/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:MIPS/Installation/Stage/ko "Handbook:MIPS/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:MIPS/Installation/Base/ko "Handbook:MIPS/Installation/Base/ko")[커널 설정](/wiki/Handbook:MIPS/Installation/Kernel/ko "Handbook:MIPS/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:MIPS/Installation/System/ko "Handbook:MIPS/Installation/System/ko")[도구 설치](/wiki/Handbook:MIPS/Installation/Tools/ko "Handbook:MIPS/Installation/Tools/ko")부트로더 설정[마무리](/wiki/Handbook:MIPS/Installation/Finalizing/ko "Handbook:MIPS/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:MIPS/Full/Working/ko "Handbook:MIPS/Full/Working/ko")[포티지 소개](/wiki/Handbook:MIPS/Working/Portage/ko "Handbook:MIPS/Working/Portage/ko")[USE 플래그](/wiki/Handbook:MIPS/Working/USE/ko "Handbook:MIPS/Working/USE/ko")[포티지 기능](/wiki/Handbook:MIPS/Working/Features/ko "Handbook:MIPS/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:MIPS/Working/Initscripts/ko "Handbook:MIPS/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:MIPS/Working/EnvVar/ko "Handbook:MIPS/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:MIPS/Full/Portage/ko "Handbook:MIPS/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:MIPS/Portage/Files/ko "Handbook:MIPS/Portage/Files/ko")[변수](/wiki/Handbook:MIPS/Portage/Variables/ko "Handbook:MIPS/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:MIPS/Portage/Branches/ko "Handbook:MIPS/Portage/Branches/ko")[추가 도구](/wiki/Handbook:MIPS/Portage/Tools/ko "Handbook:MIPS/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:MIPS/Portage/CustomTree/ko "Handbook:MIPS/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:MIPS/Portage/Advanced/ko "Handbook:MIPS/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:MIPS/Full/Networking/ko "Handbook:MIPS/Full/Networking/ko")[시작하기](/wiki/Handbook:MIPS/Networking/Introduction/ko "Handbook:MIPS/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:MIPS/Networking/Advanced/ko "Handbook:MIPS/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:MIPS/Networking/Modular/ko "Handbook:MIPS/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:MIPS/Networking/Wireless/ko "Handbook:MIPS/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:MIPS/Networking/Extending/ko "Handbook:MIPS/Networking/Extending/ko")[동적 관리](/wiki/Handbook:MIPS/Networking/Dynamic/ko "Handbook:MIPS/Networking/Dynamic/ko")

## Contents

- [1실리콘 그래픽스 머신용 arcload](#.EC.8B.A4.EB.A6.AC.EC.BD.98_.EA.B7.B8.EB.9E.98.ED.94.BD.EC.8A.A4_.EB.A8.B8.EC.8B.A0.EC.9A.A9_arcload)
- [2Cobalt MicroServers용 CoLo](#Cobalt_MicroServers.EC.9A.A9_CoLo)
  - [2.1CoLo 설치](#CoLo_.EC.84.A4.EC.B9.98)
  - [2.2CoLo 설정](#CoLo_.EC.84.A4.EC.A0.95)
- [3직렬 콘솔 설정](#.EC.A7.81.EB.A0.AC_.EC.BD.98.EC.86.94_.EC.84.A4.EC.A0.95)
- [4SGI PROM 고급 설정](#SGI_PROM_.EA.B3.A0.EA.B8.89_.EC.84.A4.EC.A0.95)
  - [4.1일반 PROM 설정](#.EC.9D.BC.EB.B0.98_PROM_.EC.84.A4.EC.A0.95)
  - [4.2직접 볼륨 헤더 부팅 설정](#.EC.A7.81.EC.A0.91_.EB.B3.BC.EB.A5.A8_.ED.97.A4.EB.8D.94_.EB.B6.80.ED.8C.85_.EC.84.A4.EC.A0.95)
  - [4.3arcload 설정](#arcload_.EC.84.A4.EC.A0.95)
- [5시스템 다시 부팅](#.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.8B.A4.EC.8B.9C_.EB.B6.80.ED.8C.85)

## 실리콘 그래픽스 머신용 arcload

arcload는 64비트 커널이 필요하여 arcboot(64비트 바이너리로 쉽게 컴파일할 수 없음)를 쓸 수 없는 머신용으로 작성했습니다. 볼륨 헤더에서 커널을 직접 불러올 때 일어나는 특이한 현상을 처리할 수 있습니다. 이제 설치 과정을 진행해보겠습니다:

`root #` `emerge arcload dvhtool`

과정이 끝나면 /usr/lib/arcload/에서 arcload 바이너리를 찾아보십시오. 이제 두 파일을 찾아보실 수 있습니다:

- sashARCS: Indy, Indigo2 (R4k), Challenge S, O2 시스템용 32비트 바이너리
- sash64: Octane/Octane2, Origin 200/2000, Indigo2 Impact 시스템용 64비트 바이너리

- sashARCS: The 32-bit binary for Indy, Indigo2 (R4k), Challenge S and O2 systems
- sash64: The 64-bit binary for Octane/Octane2, Origin 200/2000 and Indigo2 Impact systems

볼륨 헤더에 적당한 시스템 바이너리를 설지하려면 `dvhtool` 을 사용하십시오:

Indy/Indigo2/Challenge S/O2 사용자라면:

`root #` `dvhtool --unix-to-vh /usr/lib/arcload/sashARCS sashARCS`

Indigo2 Impact/Octane/Octane2/Origin 200/Origin 2000 사용자라면:

`root #` `dvhtool --unix-to-vh /usr/lib/arcload/sash64 sash64`

**참고**

부팅 CD의 볼륨 헤더로 sashARCS 또는 sash64를 설치하기 전에는 sashARCS 또는 sash64 이름을 사용할 수 없습니다. 하드디스크에서 일반 부팅 과정을 진행한다면 사용자가 원하는대로 이름을 지정할 수 있습니다.

이제 `dvhtool` 명령을 사용하여 바이너리가 볼륨 헤더에 있는지 확인해보겠습니다:

`root #` `dvhtool --print-volume-directory`

```
----- directory entries -----
Entry #0, name "sash64", start 4, bytes 55859

```

arc.cf 파일의 문법은 C와 비슷합니다. 이 파일을 설정하는 전체 세부 내용은 Linux/MIPS 위키의 arcload 페이지를 살펴보십시오. 간단하게 말해, OSLoadFilename 변수를 활용하여 부팅시 활성/비활성화 할 여러가지 옵션을 지정하십시오.

파일 **`arc.cf`** **arc.cf 예제**

```
# ARCLoad Configuration

# Some default settings...
append  "root=/dev/sda5";
append  "ro";
append  "console=ttyS0,9600";

# Our main definition. ip28 may be changed if you wish.
ip28 {
        # Definition for a "working" kernel
        # Select this by setting OSLoadFilename="ip28(working)"
        working {
                description     "SGI Indigo2 Impact R10000\n\r";
                image system    "/working";
        }

        # Definition for a "new" kernel
        # Select this by setting OSLoadFilename="ip28(new)"
        new {
                description     "SGI Indigo2 Impact R10000 - Testing Kernel\n\r";
                image system    "/new";
        }

        # For debugging a kernel
        # Select this by setting OSLoadFilename="ip28(working,debug)"
        # or OSLoadFilename="ip28(new,debug)"
        debug {
                description     "Debug console";
                append          "init=/bin/bash";
        }
}

```

arcload-0.5부터는 볼륨 헤더 또는 파티션에 arc.cf와 커널을 함께 넣어야 합니다. 새 기능을 활용하려면 /boot/ 파티션(또는 부팅 파티션을 따로 나누어놓지 않았다면 /)에 파일을 넣으십시오. arcload는 잘 알려진 grub 부트로더의 파일 시스템 드라이버 코드를 사용하기 때문에 grub과 동일하게 다양한 파일 시스템을 지원합니다.

`root #` `dvhtool --unix-to-vh arc.cf arc.cf
`

`root #` `dvhtool --unix-to-vh /usr/src/linux/vmlinux new`

## Cobalt MicroServers용 CoLo

### CoLo 설치

Cobalt server에서는 칩에 설치한 펌웨어의 기능이 상당히 빈약합니다. Cobalt BOOTROM은 SGI PROM에 비교했을 때, 원시적이며, 여러가지 중요한 제한 사항이 있습니다.

- 커널 크기는 675kB(평균)로 제한합니다. 리눅스 2.4의 현재 크기로는 이 커널 크기에 맞추기란 거의 불가능합니다. 리눅스 2.6과 3.x는 뭐 말할 것도 없습니다.
- 64비트 커널은 내장 펌웨어에서 지원하지 않습니다(현재 Cobalt 머신에서 상당히 시험수준이긴 합니다만).
- 최선의 기본은 쉘입니다

이 제한을 극복하려고 대체 펌웨어인 CoLo(Cobalt Loader)를 개발했습니다. CoLo는 Cobalt 서버의 플래시 칩에 굽거나 기존 펌웨어에서 불러올 수 있는 BOOTROM 이미지입니다.

**참고**

이 안내서는 CoLo 설정으로 진행하여 저장한 펌웨어로 불러옵니다. 이 방법이 사실 안전하며 CoLo를 설정하는 방법으로 추천합니다.

**경고**

원하는 경우 서버 플래시에 구워넣어 기존 펌웨어를 완전히 엎을 수 있습니다만, 여러분의 최선의 노력을 다해야합니다. 무엇인가 잘못되었다는 판단이 확실하다면, BOOTROM 칩을 빼내고 기존의 펌웨어를 다시 구워넣어야합니다. 끔찍한 이야기 같다면 머신의 플래시에 _구워넣지 마십시오_. 여러분이 충고를 무시하고 무슨 일을 진행하든 저희는 책임지지않습니다.

이제 CoLo 설치를 진행하겠습니다. 꾸러미 이머지로 시작하십시오.

`root #` `emerge --ask sys-boot/colo`

설치가 끝나면 /usr/lib/colo/ 디렉터리에서 다음 두개의 파일을 찾아보십시오:

- colo-chain.elf (불러올 저장 펌웨어용 "커널")
- colo-rom-image.bin (BOOT 롬에 구워넣을 ROM 이미지)

- colo-chain.elf \- the "kernel" for the stock firmware to load.
- colo-rom-image.bin \- a ROM image for flashing into the BOOTROM.

/boot/를 마운트하고 시스템에서 찾을 /boot/ 위치에서 colo-chain.elf의 압축 사본을 덤핑하여 과정을 시작하겠습니다.

`root #` `gzip -9vc /usr/lib/colo/colo-chain.elf > /boot/vmlinux.gz`

### CoLo 설정

이제 시스템을 처음 부팅하면 LCD 배경에 메뉴를 띄우는 CoLo를 불러옵니다. 첫번째 옵션(그리고 5초간 기본값으로 대기)은 하드디스크로 부팅합니다. 시스템에서는 첫번째 리눅스 파티션을 찾아 마운트를 시도하고 default.colo를 실행합니다. 문법은 CoLo 문서(/usr/share/doc/colo-X.YY/README.shell.gz 에서 고르면 되며 X.YY는 설치 버전임)에 다 있으며 매우 간단합니다.

**참고**

그냥 단순히 요령입니다: 커널을 설치할 때 커널 이미지를 두개 만드는 것이 좋겠습니다. 동작하는 걸로 알려진 kernel.gz.working과 컴파일한 kernel.gz.new입니다. 현재 "새롭"고, "동작"하는 커널에 심볼릭 링크를 걸거나 그냥 커널 이미지의 이름을 바꾸기만 하면 됩니다.

파일 **`default.colo`** **CoLo 설정 예제**

```
#:CoLo:#
mount hda1
load /kernel.gz.working
execute root=/dev/sda5 ro console=ttyS0,115200

```

**참고**

CoLo는 `#:CoLo:#` 줄로 시작하지 않는 스크립트를 불러오지 않습니다. 쉘 스크립트의 `#!/bin/sh` 부분과 동일하게 생각하십시오.

어떤 커널과 설정을 부팅에 사용하고 기본 제한 시간을 어떻게 할지 물어볼 수 있습니다. 다음 설정에서는 어떤 커널을 쓸지, 어떤 이미지를 선택해서 실행할 지 물어보는 동작을 정확히 수행합니다. vmlinux.gz.new와 vmlinux.gz.working이 실제 커널이미지겠지만 디스크에 심볼릭 링크로 링크 걸어놓은 커널 이미지일 수도 있습니다. 50이라는 매개변수 값은 50/10 초 후 첫번째 옵션("Working")으로 과정을 진행하도록 설정합니다.

파일 **`default.colo`** **메뉴 기반 설정**

```
#:CoLo:#
lcd "Mounting hda1"
mount hda1
select "Which Kernel?" 50 Working New

goto {menu-option}
var image-name vmlinux.gz.working
goto 3f
@var image-name vmlinux.gz.working
goto 2f
@var image-name vmlinux.gz.new

@lcd "Loading Linux" {image-name}
load /{image-name}
lcd "Booting..."
execute root=/dev/sda5 ro console=ttyS0,115200
boot

```

더 자세한 내용은 /usr/share/doc/colo-VERSION의 문서를 살펴보십시오.

## 직렬 콘솔 설정

좋습니다. 이제 리눅스 설치 과정에서 부팅은 잘 진행이 되고 있습니다만, 실제 터미널에 사용자가 로그인 할 수도 있습니다. Cobalt 머신이라면 특히 바람직한 상황이 아닙니다. 실제 터미널에는 아무것도 없기 때문입니다.

**참고**

지원 비디오 칩셋으로 뭔가 화려한걸 해보려 하는 분은 원할 경우 이 절을 건너뛰시는게 좋습니다.

먼저 편집기를 띄우고 /etc/inittab를 뜯어고치겠습니다. 이 파일을 쭉 살펴보면 다음과 같은 내용이 나타납니다:

파일 **`/etc/inittab`** **inittab 일부**

```
# SERIAL CONSOLE
#c0:12345:respawn:/sbin/agetty 9600 ttyS0 vt102

# TERMINALS
c1:12345:respawn:/sbin/agetty 38400 tty1 linux
c2:12345:respawn:/sbin/agetty 38400 tty2 linux
c3:12345:respawn:/sbin/agetty 38400 tty3 linux
c4:12345:respawn:/sbin/agetty 38400 tty4 linux
c5:12345:respawn:/sbin/agetty 38400 tty5 linux
c6:12345:respawn:/sbin/agetty 38400 tty6 linux

# What to do at the "Three Finger Salute".
ca:12345:ctrlaltdel:/sbin/shutdown -r now

```

먼저 c0 줄의 주석을 제거하십시오. 기본적으로 Cobalt 서버에서는 9600bps 보 레이트로 터미널을 사용합니다. BOOT ROM에서 결정한 보 레이트에 일치하려면 115200으로 맞춰야 합니다. 다음 내용은 조치 사항을 이행한 후의 해당 섹션의 모습을 나타냅니다. 헤드리스 머신(예: Cobalt 서버)에서는 로컬 터미널 라인(c1부터 c6)의 줄을 열 수 없을 경우(/dev/ttyX) 동작이 잘못될 수 있으므로, 해당 줄의 주석 처리를 권합니다.

파일 **`/etc/inittab`** **inittab 일부 예제**

```
# SERIAL CONSOLE
c0:12345:respawn:/sbin/agetty 115200 ttyS0 vt102

# TERMINALS -- These are useless on a headless qube
#c1:12345:respawn:/sbin/agetty 38400 tty1 linux
#c2:12345:respawn:/sbin/agetty 38400 tty2 linux
#c3:12345:respawn:/sbin/agetty 38400 tty3 linux
#c4:12345:respawn:/sbin/agetty 38400 tty4 linux
#c5:12345:respawn:/sbin/agetty 38400 tty5 linux
#c6:12345:respawn:/sbin/agetty 38400 tty6 linux

```

이제 마지막으로 ... 보안 터미널 처럼 신뢰할 수 있는 로컬 직렬 포트로 시스템을 불러와야 합니다. 건드려야 할 파일은 /etc/securetty 입니다. 믿을만한 시스템의 터미널 목록이 여기에 있습니다. 단순히 루트 로그인에 사용할 직렬라인만 허용하도록, 두 줄의 내용만 더 넣겠습니다.

`root #` `echo 'ttyS0' >> /etc/securetty`

마지막으로 리눅스에서 /dev/tts/0도 호출합니다. 따라서 이 부분도 추가하겠습니다:

`root #` `echo 'tts/0' >> /etc/securetty`

## SGI PROM 고급 설정

### 일반 PROM 설정

부트로더를 설치하고 다시 부팅하면(잠시 후에 진행합니다), 시스템 관리 메뉴로 진입 후 시스템을 넷부팅할 때 처음 했던것과 마찬가지로 명령 모니터(5)를 선택하십시오.

코드 **부팅 진행 후 메뉴**

```
(1) Start System
(2) Install System Software
(3) Run Diagnostics
(4) Recover System
(5) Enter Command Monitor

```

볼륨 헤더 위치를 입력하십시오:

`>>` `setenv SystemPartition scsi(0)disk(1)rdisk(0)partition(8)`

젠투를 자동으로 부팅하십시오:

`>>` `setenv AutoLoad Yes`

시간대를 설정하십시오:

`>>` `setenv TimeZone EST5EDT`

직렬 콘솔을 사용하십시오 \- 그래픽 어댑터 사용자는 "d1"(하나) 대신 "g"를 넣어야 합니다:

`>>` `setenv console d1`

직렬 전송률을 설정하십시오. 9600 값이 기본값이며, 원한다면 38400까지 사용할 수도 있겠지만, 설정 하지 않아도 그만입니다:

`>>` `setenv dbaud 9600`

이제 시스템을 부팅하는 방법에 따라 다음 설정을 계속 진행하겠습니다.

### 직접 볼륨 헤더 부팅 설정

**참고**

여기서는 ,완벽함을 다룹니다. arcload를 대신 설치하는 방법을 추천합니다.

**참고**

Indy, Indigo2 (R4k), Challenge S 에서만 동작합니다.

루트 장치를 /dev/sda3와 같은 젠투 루트 파티션으로 지정하십시오:

`>>` `setenv OSLoadPartition <root device>`

존재하는 커널을 보려면 "ls"를 입력하십시오.

`>>` `setenv OSLoader <kernel name>
`

`>>` `setenv OSLoadFilename <kernel name>`

전달할 커널 매개변수를 선언하십시오:

`>>` `setenv OSLoadOptions <kernel parameters>`

커널 매개변수를 넣지 않고 커널을 사용하려면 boot -f PROM 명령을 사용하십시오:

`root #` `boot -f new root=/dev/sda5 ro`

### arcload 설정

arcload는 OSLoadFilename 옵션을 사용하여 arc.cf에서 어떤 옵션을 설정할지 지정합니다. 설정 파일은 본질적으로 스크립트이며 최상위 레벨 블록에서는 각각의 시스템에서 사용할 부팅 이미지를 정의하고, 해당 내용을 뒤져보면 추가 설정이 있습니다. 따라서 OSLoadFilename=mysys(sefial) 설정으로 mysys 블록의 설정을 가져온 후, 직렬 연결상에서 그 밖의 옵션을 중복 적용합니다.

위의 예제 파일에서, 시스템 블록을 정의했고 동작하는 ip28, 새로운 옵션과 디버그 옵션을 확보했습니다. 다음과 같이 PROM 변수를 정의하겠습니다:

arcload를 부트로더로 선택하십시오:- sash64 또는 sashARCS:

`>>` `setenv OSLoader sash64`

arc.cf의 "ip28" 섹션에 나타난 "동작하는" 커널 이미지를 사용하십시오:

`>>` `setenv OSLoadFilename ip28(working)`

arcload-0.5부터 볼륨 헤더에 파일을 더이상 놓을 필요가 없습니다 -- 대신 파이션에 넣으면 됩니다. 설정 파일과 커널을 어디에서 찾아야 하는지 arcload에 알려주려면 OSLoadPartition PROM 변수를 반드시 설정해야 합니다. 어떤 SCSI 버스에 디스크를 붙였느냐에 따라 정확한 값이 달라집니다. SystemPartition PROM 변수를 보조 안내용도로 사용하십시오 -- 파티션 번호만 바꿔야합니다.

**참고**

리눅스에서와 같이 파티션 번호가 1부터 시작하는 것이 아니라 0부터 시작합니다.

볼륨 헤더에서 불러오려면 8번 파티션을 활용하십시오:

`>>` `setenv OSLoadPartition scsi(0)disk(1)rdisk(0)partition(8)`

그렇지 않으면 파티션과 파일 시스템 형식을 지정하십시오:

`>>` `setenv OSLoadPartition scsi(0)disk(1)rdisk(0)partition(0)[ext2]`

## 시스템 다시 부팅

chroot로 진입한 환경을 빠져나가고 모든 파티션의 마운트를 해제하십시오. 그 다음 대미를 장식할 마법의 명령을 입력하여, 실제로 시험해보십시오:
reboot.

`root #` `exit`

`cdimage ~#` `cd
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`cdimage ~#` `umount -R /mnt/gentoo
`

`cdimage ~#` `reboot`

물론 부팅 CD를 제거하는걸 잊지 않으면 새 젠투 시스템 대신 CD로 부팅합니다.

새로 설치한 젠투 환경으로 다시 부팅하고 나면, [젠투 설치 마무리](/wiki/Handbook:MIPS/Installation/Finalizing/ko "Handbook:MIPS/Installation/Finalizing/ko") 로 끝내십시오.

[← 시스템 도구 설치](/wiki/Handbook:MIPS/Installation/Tools/ko "이전 내용") [처음](/wiki/Handbook:MIPS/ko "Handbook:MIPS/ko") [마무리 →](/wiki/Handbook:MIPS/Installation/Finalizing/ko "다음 내용")