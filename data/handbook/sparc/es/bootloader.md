# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbuch:SPARC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Bootloader "Handbook:SPARC/Installation/Bootloader (100% translated)")
- español
- [français](/wiki/Handbook:SPARC/Installation/Bootloader/fr "Handbook:SPARC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Bootloader/pl "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Bootloader/pt-br "Handbook:SPARC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Bootloader/ru "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Bootloader/ta "கையேடு:SPARC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Bootloader/zh-cn "手册：SPARC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Bootloader/ja "ハンドブック:SPARC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Bootloader/ko "Handbook:SPARC/Installation/Bootloader/ko (100% translated)")

[SPARC Handbook](/wiki/Handbook:SPARC "Handbook:SPARC")[Installation](/wiki/Handbook:SPARC/Full/Installation "Handbook:SPARC/Full/Installation")[About the installation](/wiki/Handbook:SPARC/Installation/About/es "Handbook:SPARC/Installation/About/es")[Choosing the media](/wiki/Handbook:SPARC/Installation/Media/es "Handbook:SPARC/Installation/Media/es")[Configuring the network](/wiki/Handbook:SPARC/Installation/Networking/es "Handbook:SPARC/Installation/Networking/es")[Preparing the disks](/wiki/Handbook:SPARC/Installation/Disks/es "Handbook:SPARC/Installation/Disks/es")[The stage file](/wiki/Handbook:SPARC/Installation/Stage/es "Handbook:SPARC/Installation/Stage/es")[Installing base system](/wiki/Handbook:SPARC/Installation/Base/es "Handbook:SPARC/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:SPARC/Installation/Kernel/es "Handbook:SPARC/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:SPARC/Installation/System/es "Handbook:SPARC/Installation/System/es")[Installing tools](/wiki/Handbook:SPARC/Installation/Tools/es "Handbook:SPARC/Installation/Tools/es")Configuring the bootloader[Finalizing](/wiki/Handbook:SPARC/Installation/Finalizing/es "Handbook:SPARC/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:SPARC/Full/Working "Handbook:SPARC/Full/Working")[Portage introduction](/wiki/Handbook:SPARC/Working/Portage/es "Handbook:SPARC/Working/Portage/es")[USE flags](/wiki/Handbook:SPARC/Working/USE/es "Handbook:SPARC/Working/USE/es")[Portage features](/wiki/Handbook:SPARC/Working/Features/es "Handbook:SPARC/Working/Features/es")[Initscript system](/wiki/Handbook:SPARC/Working/Initscripts/es "Handbook:SPARC/Working/Initscripts/es")[Environment variables](/wiki/Handbook:SPARC/Working/EnvVar/es "Handbook:SPARC/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:SPARC/Full/Portage "Handbook:SPARC/Full/Portage")[Files and directories](/wiki/Handbook:SPARC/Portage/Files/es "Handbook:SPARC/Portage/Files/es")[Variables](/wiki/Handbook:SPARC/Portage/Variables/es "Handbook:SPARC/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:SPARC/Portage/Branches/es "Handbook:SPARC/Portage/Branches/es")[Additional tools](/wiki/Handbook:SPARC/Portage/Tools/es "Handbook:SPARC/Portage/Tools/es")[Custom package repository](/wiki/Handbook:SPARC/Portage/CustomTree/es "Handbook:SPARC/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:SPARC/Portage/Advanced/es "Handbook:SPARC/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:SPARC/Full/Networking "Handbook:SPARC/Full/Networking")[Getting started](/wiki/Handbook:SPARC/Networking/Introduction/es "Handbook:SPARC/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:SPARC/Networking/Advanced/es "Handbook:SPARC/Networking/Advanced/es")[Modular networking](/wiki/Handbook:SPARC/Networking/Modular/es "Handbook:SPARC/Networking/Modular/es")[Wireless](/wiki/Handbook:SPARC/Networking/Wireless/es "Handbook:SPARC/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:SPARC/Networking/Extending/es "Handbook:SPARC/Networking/Extending/es")[Dynamic management](/wiki/Handbook:SPARC/Networking/Dynamic/es "Handbook:SPARC/Networking/Dynamic/es")

[Handbook:SPARC/Blocks/Bootloader/es](/index.php?title=Handbook:SPARC/Blocks/Bootloader/es&action=edit&redlink=1 "Handbook:SPARC/Blocks/Bootloader/es (page does not exist)")

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

Una vez reiniciado en el nuevo entorno Gentoo, es aconsejable finalizar con [Finalizando la instalación de Gentoo](/wiki/Handbook:SPARC/Installation/Finalizing/es "Handbook:SPARC/Installation/Finalizing/es").

[← Instalar las herramientas](/wiki/Handbook:SPARC/Installation/Tools/es "Previous part") [Inicio](/wiki/Handbook:SPARC/es "Handbook:SPARC/es") [Finalizar →](/wiki/Handbook:SPARC/Installation/Finalizing/es "Next part")