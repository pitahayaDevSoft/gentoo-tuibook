# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbuch:SPARC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Bootloader "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Bootloader/es "Manual de Gentoo: SPARC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Bootloader/fr "Handbook:SPARC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Bootloader/pl "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Bootloader/pt-br "Handbook:SPARC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Bootloader/ru "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Bootloader/ta "கையேடு:SPARC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Bootloader/zh-cn "手册：SPARC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Bootloader/ja "ハンドブック:SPARC/インストール/ブートローダー (100% translated)")
- 한국어

[SPARC 핸드북](/wiki/Handbook:SPARC/ko "Handbook:SPARC/ko")[설치](/wiki/Handbook:SPARC/Full/Installation/ko "Handbook:SPARC/Full/Installation/ko")[설치 정보](/wiki/Handbook:SPARC/Installation/About/ko "Handbook:SPARC/Installation/About/ko")[매체 선택](/wiki/Handbook:SPARC/Installation/Media/ko "Handbook:SPARC/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:SPARC/Installation/Networking/ko "Handbook:SPARC/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:SPARC/Installation/Disks/ko "Handbook:SPARC/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:SPARC/Installation/Stage/ko "Handbook:SPARC/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:SPARC/Installation/Base/ko "Handbook:SPARC/Installation/Base/ko")[커널 설정](/wiki/Handbook:SPARC/Installation/Kernel/ko "Handbook:SPARC/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:SPARC/Installation/System/ko "Handbook:SPARC/Installation/System/ko")[도구 설치](/wiki/Handbook:SPARC/Installation/Tools/ko "Handbook:SPARC/Installation/Tools/ko")부트로더 설정[마무리](/wiki/Handbook:SPARC/Installation/Finalizing/ko "Handbook:SPARC/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:SPARC/Full/Working/ko "Handbook:SPARC/Full/Working/ko")[포티지 소개](/wiki/Handbook:SPARC/Working/Portage/ko "Handbook:SPARC/Working/Portage/ko")[USE 플래그](/wiki/Handbook:SPARC/Working/USE/ko "Handbook:SPARC/Working/USE/ko")[포티지 기능](/wiki/Handbook:SPARC/Working/Features/ko "Handbook:SPARC/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:SPARC/Working/Initscripts/ko "Handbook:SPARC/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:SPARC/Working/EnvVar/ko "Handbook:SPARC/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:SPARC/Full/Portage/ko "Handbook:SPARC/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:SPARC/Portage/Files/ko "Handbook:SPARC/Portage/Files/ko")[변수](/wiki/Handbook:SPARC/Portage/Variables/ko "Handbook:SPARC/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:SPARC/Portage/Branches/ko "Handbook:SPARC/Portage/Branches/ko")[추가 도구](/wiki/Handbook:SPARC/Portage/Tools/ko "Handbook:SPARC/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:SPARC/Portage/CustomTree/ko "Handbook:SPARC/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:SPARC/Portage/Advanced/ko "Handbook:SPARC/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:SPARC/Full/Networking/ko "Handbook:SPARC/Full/Networking/ko")[시작하기](/wiki/Handbook:SPARC/Networking/Introduction/ko "Handbook:SPARC/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:SPARC/Networking/Advanced/ko "Handbook:SPARC/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:SPARC/Networking/Modular/ko "Handbook:SPARC/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:SPARC/Networking/Wireless/ko "Handbook:SPARC/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:SPARC/Networking/Extending/ko "Handbook:SPARC/Networking/Extending/ko")[동적 관리](/wiki/Handbook:SPARC/Networking/Dynamic/ko "Handbook:SPARC/Networking/Dynamic/ko")

## Contents

- [1GRUB](#GRUB)
  - [1.1Emerge](#Emerge)
  - [1.2Install](#Install)
    - [1.2.1GPT](#GPT)
    - [1.2.2Sun partition table](#Sun_partition_table)
  - [1.3Configure](#Configure)
- [2SPARC 부트로더 SILO](#SPARC_.EB.B6.80.ED.8A.B8.EB.A1.9C.EB.8D.94_SILO)
- [3시스템 다시 부팅](#.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.8B.A4.EC.8B.9C_.EB.B6.80.ED.8C.85)

## GRUB

When [selecting a 64-bit profile](/wiki/Handbook:SPARC/Installation/Base/ko#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/ko") during installation, then [GRUB](/wiki/GRUB/ko "GRUB/ko") is the only supported bootloader.

### Emerge

GRUB should be correctly configured for the platform automatically based on the profile. To make it explicit, however, specify it using:

`root #` `echo 'GRUB_PLATFORMS="ieee1275"' >> /etc/portage/make.conf`

`root #` `emerge --ask --verbose sys-boot/grub`

The GRUB software has now been merged to the system, but not yet installed as a bootloader.

### Install

#### GPT

If the [disk is partitioned](/wiki/Handbook:SPARC/Installation/Disks/ko#Using_fdisk_to_partition_the_disk "Handbook:SPARC/Installation/Disks/ko") using GPT (the preferred method), then install GRUB to the BIOS boot partition. Presuming the first disk (the one where the system boots from) is /dev/sda, the following commands will do:

`root #` `grub-install --target=sparc64-ieee1275 --recheck /dev/sda`

**요령**

{{{1}}}

#### Sun partition table

If the disk is partitioned using a Sun partition table instead, GRUB must be installed using blocklists. In this mode, instead of providing the physical disk as an argument, provide the path to the partition on which /boot/grub is mounted.

`root #` `grub-install --target=sparc64-ieee1275 --recheck --force --skip-fs-probe /dev/sda1`

### Configure

Next, generate the GRUB2 configuration based on the user configuration specified in the /etc/default/grub file and /etc/grub.d scripts. In most cases, no configuration is needed by users as GRUB2 will automatically detect which kernel to boot (the highest one available in /boot/) and what the root file system is. It is also possible to append kernel parameters in /etc/default/grub using the GRUB\_CMDLINE\_LINUX variable.

To generate the final GRUB2 configuration, run the grub-mkconfig command:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.19.3-gentoo
Found initrd image: /boot/initramfs-genkernel-sparc-6.19.3-gentoo
done

```

The output of the command must mention that at least one Linux image is found, as those are needed to boot the system. If an initramfs is used or genkernel was used to build the kernel, the correct initrd image should be detected as well. If this is not the case, go to /boot/ and check the contents using the ls command. If the files are indeed missing, go back to the kernel configuration and installation instructions.

## SPARC 부트로더 SILO

이제 Sparc Improved boot LOader [SILO](http://www.sparc-boot.org/) 를 설치하고 설정할 차례입니다.

`root #` `emerge --ask sys-boot/silo`

그 다음 /etc/silo.conf를 만드십시오:

`root #` `nano -w /etc/silo.conf`

아래에 silo.conf 파일 내용을 나타냈습니다. 이 책에서 언급한 공간 분할 방식과 kernel-6.19.3-gentoo 커널 이미지, initramfs-genkernel-sparc64-6.19.3-gentoo initramfs를 활용합니다.

파일 **`/etc/silo.conf`** **예제 설정 파일**

```
partition = 1         # Boot partition (= root partition)
root = /dev/sda1      # Root partition
timeout = 150         # Wait 15 seconds before booting the default section

image = /boot/kernel-6.19.3-gentoo
  label = linux
  append = "initrd=/boot/initramfs-genkernel-sparc64-6.19.3-gentoo real_root=/dev/sda1"

```

포티지로 가져온 silo.conf 예제 파일을 사용할 때, 모든 줄이 필요한건 아니므로 주석처리했는지 확인하십시오.

SILO(부트로더)를 설치하는 물리 디스크와 /etc/silo.conf를 저장하는 물리 디스크가 다르다면, /etc/silo.conf 파일을 부트로더 디스크에 먼저 복사하십시오. 디스크에서 /boot/ 파티션을 따로 나누었다면, /boot/에 파일을 복사하시고 /sbin/silo를 실행하십시오:

`root #` `cp /etc/silo.conf /boot
`

`root #` `/sbin/silo -C /boot/silo.conf`

```
/boot/silo.conf appears to be valid

```

아니면 그냥 /sbin/silo를 실행하십시오:

`root #` `/sbin/silo`

```
/etc/silo.conf appears to be valid

```

**참고**

업데이트 할 때, 또는 [sys-boot/silo](https://packages.gentoo.org/packages/sys-boot/silo) 꾸러미를 설치할 때 매번 silo(필요한 경우 매개변수 붙임)를 실행하십시오.

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

새로 설치한 젠투 환경으로 다시 부팅하고 나면, [젠투 설치 마무리](/wiki/Handbook:SPARC/Installation/Finalizing/ko "Handbook:SPARC/Installation/Finalizing/ko") 로 끝내십시오.

[← 시스템 도구 설치](/wiki/Handbook:SPARC/Installation/Tools/ko "이전 내용") [처음](/wiki/Handbook:SPARC/ko "Handbook:SPARC/ko") [마무리 →](/wiki/Handbook:SPARC/Installation/Finalizing/ko "다음 내용")