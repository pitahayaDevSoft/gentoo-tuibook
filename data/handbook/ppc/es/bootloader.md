# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbuch:PPC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Bootloader "Handbook:PPC/Installation/Bootloader (100% translated)")
- español
- [français](/wiki/Handbook:PPC/Installation/Bootloader/fr "Handbook:PPC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Bootloader/it "Handbook:PPC/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Bootloader/hu "Handbook:PPC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Bootloader/pl "Handbook:PPC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Bootloader/pt-br "Handbook:PPC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Bootloader/ru "Handbook:PPC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Bootloader/ta "கையேடு:PPC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Bootloader/zh-cn "手册：PPC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Bootloader/ja "ハンドブック:PPC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Bootloader/ko "Handbook:PPC/Installation/Bootloader/ko (100% translated)")

[PPC Handbook](/wiki/Handbook:PPC "Handbook:PPC")[Installation](/wiki/Handbook:PPC/Full/Installation "Handbook:PPC/Full/Installation")[About the installation](/wiki/Handbook:PPC/Installation/About/es "Handbook:PPC/Installation/About/es")[Choosing the media](/wiki/Handbook:PPC/Installation/Media/es "Handbook:PPC/Installation/Media/es")[Configuring the network](/wiki/Handbook:PPC/Installation/Networking/es "Handbook:PPC/Installation/Networking/es")[Preparing the disks](/wiki/Handbook:PPC/Installation/Disks/es "Handbook:PPC/Installation/Disks/es")[The stage file](/wiki/Handbook:PPC/Installation/Stage/es "Handbook:PPC/Installation/Stage/es")[Installing base system](/wiki/Handbook:PPC/Installation/Base/es "Handbook:PPC/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:PPC/Installation/Kernel/es "Handbook:PPC/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:PPC/Installation/System/es "Handbook:PPC/Installation/System/es")[Installing tools](/wiki/Handbook:PPC/Installation/Tools/es "Handbook:PPC/Installation/Tools/es")Configuring the bootloader[Finalizing](/wiki/Handbook:PPC/Installation/Finalizing/es "Handbook:PPC/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:PPC/Full/Working "Handbook:PPC/Full/Working")[Portage introduction](/wiki/Handbook:PPC/Working/Portage/es "Handbook:PPC/Working/Portage/es")[USE flags](/wiki/Handbook:PPC/Working/USE/es "Handbook:PPC/Working/USE/es")[Portage features](/wiki/Handbook:PPC/Working/Features/es "Handbook:PPC/Working/Features/es")[Initscript system](/wiki/Handbook:PPC/Working/Initscripts/es "Handbook:PPC/Working/Initscripts/es")[Environment variables](/wiki/Handbook:PPC/Working/EnvVar/es "Handbook:PPC/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:PPC/Full/Portage "Handbook:PPC/Full/Portage")[Files and directories](/wiki/Handbook:PPC/Portage/Files/es "Handbook:PPC/Portage/Files/es")[Variables](/wiki/Handbook:PPC/Portage/Variables/es "Handbook:PPC/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:PPC/Portage/Branches/es "Handbook:PPC/Portage/Branches/es")[Additional tools](/wiki/Handbook:PPC/Portage/Tools/es "Handbook:PPC/Portage/Tools/es")[Custom package repository](/wiki/Handbook:PPC/Portage/CustomTree/es "Handbook:PPC/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:PPC/Portage/Advanced/es "Handbook:PPC/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:PPC/Full/Networking "Handbook:PPC/Full/Networking")[Getting started](/wiki/Handbook:PPC/Networking/Introduction/es "Handbook:PPC/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:PPC/Networking/Advanced/es "Handbook:PPC/Networking/Advanced/es")[Modular networking](/wiki/Handbook:PPC/Networking/Modular/es "Handbook:PPC/Networking/Modular/es")[Wireless](/wiki/Handbook:PPC/Networking/Wireless/es "Handbook:PPC/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:PPC/Networking/Extending/es "Handbook:PPC/Networking/Extending/es")[Dynamic management](/wiki/Handbook:PPC/Networking/Dynamic/es "Handbook:PPC/Networking/Dynamic/es")

[Handbook:PPC/Blocks/Bootloader/es](/index.php?title=Handbook:PPC/Blocks/Bootloader/es&action=edit&redlink=1 "Handbook:PPC/Blocks/Bootloader/es (page does not exist)")

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

Una vez reiniciado en el nuevo entorno Gentoo, es aconsejable finalizar con [Finalizando la instalación de Gentoo](/wiki/Handbook:PPC/Installation/Finalizing/es "Handbook:PPC/Installation/Finalizing/es").

[← Instalar las herramientas](/wiki/Handbook:PPC/Installation/Tools/es "Previous part") [Inicio](/wiki/Handbook:PPC/es "Handbook:PPC/es") [Finalizar →](/wiki/Handbook:PPC/Installation/Finalizing/es "Next part")