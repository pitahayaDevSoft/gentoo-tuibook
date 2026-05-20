# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Bootloader/de "Handbuch:PPC64/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Bootloader "Handbook:PPC64/Installation/Bootloader (100% translated)")
- español
- [français](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Bootloader/it "Handbook:PPC64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Handbook:PPC64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Bootloader/pt-br "Handbook:PPC64/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Bootloader/ru "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Bootloader/ta "கையேடு:PPC64/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Bootloader/zh-cn "手册：PPC64/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Bootloader/ja "ハンドブック:PPC64/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko (100% translated)")

[PPC64 Handbook](/wiki/Handbook:PPC64 "Handbook:PPC64")[Installation](/wiki/Handbook:PPC64/Full/Installation "Handbook:PPC64/Full/Installation")[About the installation](/wiki/Handbook:PPC64/Installation/About/es "Handbook:PPC64/Installation/About/es")[Choosing the media](/wiki/Handbook:PPC64/Installation/Media/es "Handbook:PPC64/Installation/Media/es")[Configuring the network](/wiki/Handbook:PPC64/Installation/Networking/es "Handbook:PPC64/Installation/Networking/es")[Preparing the disks](/wiki/Handbook:PPC64/Installation/Disks/es "Handbook:PPC64/Installation/Disks/es")[The stage file](/wiki/Handbook:PPC64/Installation/Stage/es "Handbook:PPC64/Installation/Stage/es")[Installing base system](/wiki/Handbook:PPC64/Installation/Base/es "Handbook:PPC64/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:PPC64/Installation/Kernel/es "Handbook:PPC64/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:PPC64/Installation/System/es "Handbook:PPC64/Installation/System/es")[Installing tools](/wiki/Handbook:PPC64/Installation/Tools/es "Handbook:PPC64/Installation/Tools/es")Configuring the bootloader[Finalizing](/wiki/Handbook:PPC64/Installation/Finalizing/es "Handbook:PPC64/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:PPC64/Full/Working "Handbook:PPC64/Full/Working")[Portage introduction](/wiki/Handbook:PPC64/Working/Portage/es "Handbook:PPC64/Working/Portage/es")[USE flags](/wiki/Handbook:PPC64/Working/USE/es "Handbook:PPC64/Working/USE/es")[Portage features](/wiki/Handbook:PPC64/Working/Features/es "Handbook:PPC64/Working/Features/es")[Initscript system](/wiki/Handbook:PPC64/Working/Initscripts/es "Handbook:PPC64/Working/Initscripts/es")[Environment variables](/wiki/Handbook:PPC64/Working/EnvVar/es "Handbook:PPC64/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:PPC64/Full/Portage "Handbook:PPC64/Full/Portage")[Files and directories](/wiki/Handbook:PPC64/Portage/Files/es "Handbook:PPC64/Portage/Files/es")[Variables](/wiki/Handbook:PPC64/Portage/Variables/es "Handbook:PPC64/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:PPC64/Portage/Branches/es "Handbook:PPC64/Portage/Branches/es")[Additional tools](/wiki/Handbook:PPC64/Portage/Tools/es "Handbook:PPC64/Portage/Tools/es")[Custom package repository](/wiki/Handbook:PPC64/Portage/CustomTree/es "Handbook:PPC64/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:PPC64/Portage/Advanced/es "Handbook:PPC64/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:PPC64/Full/Networking "Handbook:PPC64/Full/Networking")[Getting started](/wiki/Handbook:PPC64/Networking/Introduction/es "Handbook:PPC64/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:PPC64/Networking/Advanced/es "Handbook:PPC64/Networking/Advanced/es")[Modular networking](/wiki/Handbook:PPC64/Networking/Modular/es "Handbook:PPC64/Networking/Modular/es")[Wireless](/wiki/Handbook:PPC64/Networking/Wireless/es "Handbook:PPC64/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:PPC64/Networking/Extending/es "Handbook:PPC64/Networking/Extending/es")[Dynamic management](/wiki/Handbook:PPC64/Networking/Dynamic/es "Handbook:PPC64/Networking/Dynamic/es")

[Handbook:PPC64/Blocks/Bootloader/es](/index.php?title=Handbook:PPC64/Blocks/Bootloader/es&action=edit&redlink=1 "Handbook:PPC64/Blocks/Bootloader/es (page does not exist)")

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

Una vez reiniciado en el nuevo entorno Gentoo, es aconsejable finalizar con [Finalizando la instalación de Gentoo](/wiki/Handbook:PPC64/Installation/Finalizing/es "Handbook:PPC64/Installation/Finalizing/es").

[← Instalar las herramientas](/wiki/Handbook:PPC64/Installation/Tools/es "Previous part") [Inicio](/wiki/Handbook:PPC64/es "Handbook:PPC64/es") [Finalizar →](/wiki/Handbook:PPC64/Installation/Finalizing/es "Next part")