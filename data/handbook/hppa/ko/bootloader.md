# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbuch:HPPA/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Bootloader/es "Manual de Gentoo: HPPA/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Bootloader/it "Handbook:HPPA/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Bootloader/pl "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Bootloader/pt-br "Manual:HPPA/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Bootloader/ru "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Bootloader/ta "கையேடு:HPPA/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Bootloader/zh-cn "手册：HPPA/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Bootloader/ja "ハンドブック:HPPA/インストール/ブートローダー (100% translated)")
- 한국어

[HPPA 핸드북](/wiki/Handbook:HPPA/ko "Handbook:HPPA/ko")[설치](/wiki/Handbook:HPPA/Full/Installation/ko "Handbook:HPPA/Full/Installation/ko")[설치 정보](/wiki/Handbook:HPPA/Installation/About/ko "Handbook:HPPA/Installation/About/ko")[매체 선택](/wiki/Handbook:HPPA/Installation/Media/ko "Handbook:HPPA/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:HPPA/Installation/Networking/ko "Handbook:HPPA/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:HPPA/Installation/Disks/ko "Handbook:HPPA/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:HPPA/Installation/Stage/ko "Handbook:HPPA/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:HPPA/Installation/Base/ko "Handbook:HPPA/Installation/Base/ko")[커널 설정](/wiki/Handbook:HPPA/Installation/Kernel/ko "Handbook:HPPA/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:HPPA/Installation/System/ko "Handbook:HPPA/Installation/System/ko")[도구 설치](/wiki/Handbook:HPPA/Installation/Tools/ko "Handbook:HPPA/Installation/Tools/ko")부트로더 설정[마무리](/wiki/Handbook:HPPA/Installation/Finalizing/ko "Handbook:HPPA/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:HPPA/Full/Working/ko "Handbook:HPPA/Full/Working/ko")[포티지 소개](/wiki/Handbook:HPPA/Working/Portage/ko "Handbook:HPPA/Working/Portage/ko")[USE 플래그](/wiki/Handbook:HPPA/Working/USE/ko "Handbook:HPPA/Working/USE/ko")[포티지 기능](/wiki/Handbook:HPPA/Working/Features/ko "Handbook:HPPA/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:HPPA/Working/Initscripts/ko "Handbook:HPPA/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:HPPA/Working/EnvVar/ko "Handbook:HPPA/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:HPPA/Full/Portage/ko "Handbook:HPPA/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:HPPA/Portage/Files/ko "Handbook:HPPA/Portage/Files/ko")[변수](/wiki/Handbook:HPPA/Portage/Variables/ko "Handbook:HPPA/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:HPPA/Portage/Branches/ko "Handbook:HPPA/Portage/Branches/ko")[추가 도구](/wiki/Handbook:HPPA/Portage/Tools/ko "Handbook:HPPA/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:HPPA/Portage/CustomTree/ko "Handbook:HPPA/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:HPPA/Portage/Advanced/ko "Handbook:HPPA/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:HPPA/Full/Networking/ko "Handbook:HPPA/Full/Networking/ko")[시작하기](/wiki/Handbook:HPPA/Networking/Introduction/ko "Handbook:HPPA/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:HPPA/Networking/Advanced/ko "Handbook:HPPA/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:HPPA/Networking/Modular/ko "Handbook:HPPA/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:HPPA/Networking/Wireless/ko "Handbook:HPPA/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:HPPA/Networking/Extending/ko "Handbook:HPPA/Networking/Extending/ko")[동적 관리](/wiki/Handbook:HPPA/Networking/Dynamic/ko "Handbook:HPPA/Networking/Dynamic/ko")

## PALO 설치

PA-RISC 플랫폼의 부트로더를 palo라고 합니다. 먼저 프로그램을 시스템에 이머지하십시오:

`root #` `emerge --ask sys-boot/palo`

설정은 /etc/palo.conf에서 찾을 수 있습니다. 아래는 설정 예제입니다:

**경고**

This configuration **must** be changed after running palo for the first time! See below for after first setup.

파일 **`/etc/palo.conf`** **간단한 PALO 설정 예제**

```
'"`UNIQ--pre-00000001-QINU`"'

```

첫번째 줄은 커널 위치와 사용할 매개 변수를 palo 명령에 알립니다. `2/kernel-6.19.1-gentoo` 문자열은 /kernel-6.19.1-gentoo 이름의 커널이 두번째 분할 영역에 있음을 의미합니다. 커널의 경로는 루트 분할 영역이 아닌 _부팅_ 분할 영역과 관련 있음을 주의 깊게 참고하십시오.

두번째 줄은 사용할 복구 커널을 나타냅니다. 첫 설치 단계이며, 복구 커널이(아직) 없다면, 이 부분을 주석처리하십시오. 세번째 줄은 palo가 있는 디스크를 나타냅니다.

To format the disk, palo must be run with certain arguments. This example uses _ext4_ for the first partition:

`root #` `palo --format-as=4 --init-partitioned=/dev/sda`

설정이 끝났으면, 간단하게 palo 명령을 실행하십시오.

`root #` `palo`

The configuration must then be updated for post-first-install use:

파일 **`/etc/palo.conf`** **Simple PALO configuration example**

```
--commandline=2/kernel-6.19.1-gentoo root=/dev/sda4
--recoverykernel=/vmlinux.old
# Don't throw away the old partition, just update the existing one on `palo` runs.
--update-partitioned=/dev/sda
# --format-as has two meanings depending on whether --init-partitioned or --update-partitioned is used. Keep this line.
--format-as=4

```

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

새로 설치한 젠투 환경으로 다시 부팅하고 나면, [젠투 설치 마무리](/wiki/Handbook:HPPA/Installation/Finalizing/ko "Handbook:HPPA/Installation/Finalizing/ko") 로 끝내십시오.

[← 시스템 도구 설치](/wiki/Handbook:HPPA/Installation/Tools/ko "이전 내용") [처음](/wiki/Handbook:HPPA/ko "Handbook:HPPA/ko") [마무리 →](/wiki/Handbook:HPPA/Installation/Finalizing/ko "다음 내용")