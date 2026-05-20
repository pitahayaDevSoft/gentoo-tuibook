# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Bootloader/de "Handbuch:Alpha/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Bootloader "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Bootloader/es "Manual de Gentoo: Alpha/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Bootloader/fr "Manuel:Alpha/Installation/Système d'amorçage (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Bootloader/it "Handbook:Alpha/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Bootloader/hu "Handbook:Alpha/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Bootloader/pl "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Bootloader/pt-br "Manual:Alpha/Instalação/Gerenciador de boot (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Bootloader/cs "Handbook:Alpha/Installation/Bootloader/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Bootloader/ru "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Bootloader/ta "கையேடு:Alpha/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Bootloader/zh-cn "手册：Alpha/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Bootloader/ja "ハンドブック:Alpha/インストール/ブートローダー (100% translated)")
- 한국어

[Alpha 핸드북](/wiki/Handbook:Alpha/ko "Handbook:Alpha/ko")[설치](/wiki/Handbook:Alpha/Full/Installation/ko "Handbook:Alpha/Full/Installation/ko")[설치 정보](/wiki/Handbook:Alpha/Installation/About/ko "Handbook:Alpha/Installation/About/ko")[매체 선택](/wiki/Handbook:Alpha/Installation/Media/ko "Handbook:Alpha/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:Alpha/Installation/Networking/ko "Handbook:Alpha/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:Alpha/Installation/Disks/ko "Handbook:Alpha/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:Alpha/Installation/Stage/ko "Handbook:Alpha/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:Alpha/Installation/Base/ko "Handbook:Alpha/Installation/Base/ko")[커널 설정](/wiki/Handbook:Alpha/Installation/Kernel/ko "Handbook:Alpha/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:Alpha/Installation/System/ko "Handbook:Alpha/Installation/System/ko")[도구 설치](/wiki/Handbook:Alpha/Installation/Tools/ko "Handbook:Alpha/Installation/Tools/ko")부트로더 설정[마무리](/wiki/Handbook:Alpha/Installation/Finalizing/ko "Handbook:Alpha/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:Alpha/Full/Working/ko "Handbook:Alpha/Full/Working/ko")[포티지 소개](/wiki/Handbook:Alpha/Working/Portage/ko "Handbook:Alpha/Working/Portage/ko")[USE 플래그](/wiki/Handbook:Alpha/Working/USE/ko "Handbook:Alpha/Working/USE/ko")[포티지 기능](/wiki/Handbook:Alpha/Working/Features/ko "Handbook:Alpha/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:Alpha/Working/Initscripts/ko "Handbook:Alpha/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:Alpha/Working/EnvVar/ko "Handbook:Alpha/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:Alpha/Full/Portage/ko "Handbook:Alpha/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:Alpha/Portage/Files/ko "Handbook:Alpha/Portage/Files/ko")[변수](/wiki/Handbook:Alpha/Portage/Variables/ko "Handbook:Alpha/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:Alpha/Portage/Branches/ko "Handbook:Alpha/Portage/Branches/ko")[추가 도구](/wiki/Handbook:Alpha/Portage/Tools/ko "Handbook:Alpha/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:Alpha/Portage/CustomTree/ko "Handbook:Alpha/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:Alpha/Portage/Advanced/ko "Handbook:Alpha/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:Alpha/Full/Networking/ko "Handbook:Alpha/Full/Networking/ko")[시작하기](/wiki/Handbook:Alpha/Networking/Introduction/ko "Handbook:Alpha/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:Alpha/Networking/Advanced/ko "Handbook:Alpha/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:Alpha/Networking/Modular/ko "Handbook:Alpha/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:Alpha/Networking/Wireless/ko "Handbook:Alpha/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:Alpha/Networking/Extending/ko "Handbook:Alpha/Networking/Extending/ko")[동적 관리](/wiki/Handbook:Alpha/Networking/Dynamic/ko "Handbook:Alpha/Networking/Dynamic/ko")

## 선택

이제 커널을 설정했고 컴파일했으며 필요한 시스템 설정 파일의 내용을 올바르게 채워넣었으니, 시스템을 시작할 때 커널을 실행할 프로그램을 설치할 차례입니다. 이 프로그램을 부트로더라고합니다.

리눅스/알파용의 부트로더는 여러가지가 있습니다. 지원 부트로더 중 전부는 아니고 하나를 선택하십시오. [aBoot](#.EA.B8.B0.EB.B3.B8:_aBoot_.EC.82.AC.EC.9A.A9) 와 [MILO](#.EB.8C.80.EC.95.88:_MILO_.EC.82.AC.EC.9A.A9) 를 문서에 남겨두겠습니다.

## 기본: aBoot 사용

**참고**

aboot는 ext2와 ext3 파티션에서만 부팅합니다.

먼저 시스템에 aboot를 설치하십시오

`root #` `emerge --ask sys-boot/aboot`

다음 과정은 부팅 디스크를 부팅할 수 있게 만드는 작업입니다. 이 과정을 통해 시스템을 부팅하는 과정에서 aboot를 시작합니다. 디스크 동작을 시작할 때 aboot 부트로더를 기록하여 부팅 디스크를 부팅할 수 있도록 하겠습니다.

`root #` `swriteboot -f3 /dev/sda /boot/bootlx
`

`root #` `abootconf /dev/sda 2
`

**참고**

이 장에서 활용하는 분할 배치 방식과 다른 방식을 활용한다면, 그에 맞게 명령을 바꾸어야 합니다. 관련 설명서 페이지(man 8 swriteboot와 man 8 abootconf)를 반드시 읽으십시오. 또한 루트 파일 시스템이 JFS로 동작한다면, 우선 커널 옵션으로 ro를 붙여서 읽기 전용으로 마운트했는지 확인하십시오.

비록 aboot를 설치했지만, 여전히 파일을 작성하는 과정이 필요합니다. aboot는 각 설정에 대해 한 줄만 필요하므로, 우리는 다음과 같이 처리할 수 있습니다:

`root #` `echo '0:2/boot/vmlinux.gz root=/dev/sda3' > /etc/aboot.conf`

커널을 빌드하는 동안 마찬가지로 initramfs도 빌드했다면, 이 파일도 참조하도록 설정을 바꾸어야 하며, initramfs에 실제 루트 장치가 어디에 있는지도 알려야합니다:

`root #` `echo '0:2/boot/vmlinux.gz initrd=/boot/initramfs-genkernel-alpha-6.19.1-gentoo root=/dev/sda3' > /etc/aboot.conf`

추가로, SRM 변수를 설정하여 젠투를 자동으로 부팅하게 할 수 있습니다. 이 변수를 리눅스에서 설정하려 할 수 있겠지만 SRM 콘솔 자체에서 하는게 더 쉬울 지도 모릅니다.

`root #` `cd /proc/srm_environment/named_variables
`

`root #` `echo -n 0 > boot_osflags
`

`root #` `echo -n '' > boot_file
`

`root #` `echo -n 'BOOT' > auto_action
`

`root #` `echo -n 'dkc100' > bootdef_dev`

물론 dkc100을 부팅할 장치에 따라 다르게 설정하십시오.

SRM 콘솔을 나중에 띄우려면(젠투 설치 복구, 일부 변수 편집 등), 자동으로 불러오는 과정을 멈출때 `Ctrl` + `C` 를 치십시오.

직렬 콘솔을 활용하여 설치할 때, aboot.conf에 직렬 콘솔 부팅 플래그 설정을 꼭 포함하십시오. 더 많은 내용은 /etc/aboot.conf.example을 참조하십시오.

Aboot을 설정했고 쓸 준비가 끝났습니다. [시스템 다시 부팅](#.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.8B.A4.EC.8B.9C_.EB.B6.80.ED.8C.85) 으로 계속 진행하십시오.

이제 [시스템 다시 부팅](#.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.8B.A4.EC.8B.9C_.EB.B6.80.ED.8C.85) 으로 진행하십시오.

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

새로 설치한 젠투 환경으로 다시 부팅하고 나면, [젠투 설치 마무리](/wiki/Handbook:Alpha/Installation/Finalizing/ko "Handbook:Alpha/Installation/Finalizing/ko") 로 끝내십시오.

[← 시스템 도구 설치](/wiki/Handbook:Alpha/Installation/Tools/ko "이전 내용") [처음](/wiki/Handbook:Alpha/ko "Handbook:Alpha/ko") [마무리 →](/wiki/Handbook:Alpha/Installation/Finalizing/ko "다음 내용")