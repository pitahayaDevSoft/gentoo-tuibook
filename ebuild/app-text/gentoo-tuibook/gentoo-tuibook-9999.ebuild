# Copyright 2026 Gentoo Authors
# Distributed under the terms of the MIT License

EAPI=8

inherit git-r3 go-module

DESCRIPTION="TUI app for reading the Gentoo Linux installation handbook in the terminal"
HOMEPAGE="https://github.com/julesklord/gentoo-tuibook"
EGIT_REPO_URI="https://github.com/julesklord/gentoo-tuibook.git"

LICENSE="MIT"
SLOT="0"
KEYWORDS=""

BDEPEND="dev-lang/go"

src_unpack() {
	git-r3_src_unpack
	# Pre-fetch Go modules during unpack phase before network restriction sandbox is applied
	go-module_live_vendor
}

src_compile() {
	# Build natively using our industrial Makefile with Gentoo's Go environment
	emake GOFLAGS="-mod=vendor" build
}

src_install() {
	# Install into Portage's image directory under the standard /usr prefix
	emake DESTDIR="${D}" PREFIX="/usr" install
	einstalldocs
}
