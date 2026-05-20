# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbuch:HPPA/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader (100% translated)")
- español
- [français](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Bootloader/it "Handbook:HPPA/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Bootloader/pl "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Bootloader/pt-br "Manual:HPPA/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Bootloader/ru "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Bootloader/ta "கையேடு:HPPA/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Bootloader/zh-cn "手册：HPPA/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Bootloader/ja "ハンドブック:HPPA/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko (100% translated)")

[HPPA Handbook](/wiki/Handbook:HPPA "Handbook:HPPA")[Installation](/wiki/Handbook:HPPA/Full/Installation "Handbook:HPPA/Full/Installation")[About the installation](/wiki/Handbook:HPPA/Installation/About/es "Handbook:HPPA/Installation/About/es")[Choosing the media](/wiki/Handbook:HPPA/Installation/Media/es "Handbook:HPPA/Installation/Media/es")[Configuring the network](/wiki/Handbook:HPPA/Installation/Networking/es "Handbook:HPPA/Installation/Networking/es")[Preparing the disks](/wiki/Handbook:HPPA/Installation/Disks/es "Handbook:HPPA/Installation/Disks/es")[The stage file](/wiki/Handbook:HPPA/Installation/Stage/es "Handbook:HPPA/Installation/Stage/es")[Installing base system](/wiki/Handbook:HPPA/Installation/Base/es "Handbook:HPPA/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:HPPA/Installation/Kernel/es "Handbook:HPPA/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:HPPA/Installation/System/es "Handbook:HPPA/Installation/System/es")[Installing tools](/wiki/Handbook:HPPA/Installation/Tools/es "Handbook:HPPA/Installation/Tools/es")Configuring the bootloader[Finalizing](/wiki/Handbook:HPPA/Installation/Finalizing/es "Handbook:HPPA/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:HPPA/Full/Working "Handbook:HPPA/Full/Working")[Portage introduction](/wiki/Handbook:HPPA/Working/Portage/es "Handbook:HPPA/Working/Portage/es")[USE flags](/wiki/Handbook:HPPA/Working/USE/es "Handbook:HPPA/Working/USE/es")[Portage features](/wiki/Handbook:HPPA/Working/Features/es "Handbook:HPPA/Working/Features/es")[Initscript system](/wiki/Handbook:HPPA/Working/Initscripts/es "Handbook:HPPA/Working/Initscripts/es")[Environment variables](/wiki/Handbook:HPPA/Working/EnvVar/es "Handbook:HPPA/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:HPPA/Full/Portage "Handbook:HPPA/Full/Portage")[Files and directories](/wiki/Handbook:HPPA/Portage/Files/es "Handbook:HPPA/Portage/Files/es")[Variables](/wiki/Handbook:HPPA/Portage/Variables/es "Handbook:HPPA/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:HPPA/Portage/Branches/es "Handbook:HPPA/Portage/Branches/es")[Additional tools](/wiki/Handbook:HPPA/Portage/Tools/es "Handbook:HPPA/Portage/Tools/es")[Custom package repository](/wiki/Handbook:HPPA/Portage/CustomTree/es "Handbook:HPPA/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:HPPA/Portage/Advanced/es "Handbook:HPPA/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:HPPA/Full/Networking "Handbook:HPPA/Full/Networking")[Getting started](/wiki/Handbook:HPPA/Networking/Introduction/es "Handbook:HPPA/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:HPPA/Networking/Advanced/es "Handbook:HPPA/Networking/Advanced/es")[Modular networking](/wiki/Handbook:HPPA/Networking/Modular/es "Handbook:HPPA/Networking/Modular/es")[Wireless](/wiki/Handbook:HPPA/Networking/Wireless/es "Handbook:HPPA/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:HPPA/Networking/Extending/es "Handbook:HPPA/Networking/Extending/es")[Dynamic management](/wiki/Handbook:HPPA/Networking/Dynamic/es "Handbook:HPPA/Networking/Dynamic/es")

[Handbook:HPPA/Blocks/Bootloader/es](/index.php?title=Handbook:HPPA/Blocks/Bootloader/es&action=edit&redlink=1 "Handbook:HPPA/Blocks/Bootloader/es (page does not exist)")

## Reiniciar el sistema

Salga del entorno chroot y desmonte todas las particiones que continúen montadas. Luego escriba la orden mágica da inicio a la auténtica prueba final: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

¡No olvide retirar la imagen live, de lo contrario puede ser arrancada de nuevo en lugar del sistema Gentoo recién instalado!

Una vez reiniciado en el nuevo entorno Gentoo, es aconsejable finalizar con [Finalizando la instalación de Gentoo](/wiki/Handbook:HPPA/Installation/Finalizing/es "Handbook:HPPA/Installation/Finalizing/es").

[← Instalar las herramientas](/wiki/Handbook:HPPA/Installation/Tools/es "Previous part") [Inicio](/wiki/Handbook:HPPA/es "Handbook:HPPA/es") [Finalizar →](/wiki/Handbook:HPPA/Installation/Finalizing/es "Next part")