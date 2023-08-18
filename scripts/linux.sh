#!/bin/sh

show_help(){

	cat << EOF
linux.sh

Choose one of the available commands:
	build
	setup
	install-dependencies
	help | --help | -h
	
EOF


}


if [ $# -eq 0 ]; then
	show_help
	exit
fi

_install_dependencies() {
	echo "Install dependencies in Linux"
}

_setup(){
	echo "Setup in Linux"

}

_build(){

	echo "Building in Linux"
	go build
	./ryuu
}



main() {

	case "$1" in 
		(build)
			shift
			_build "$@"
			;;
		(install-dependencies)
			shift
			_install_dependencies "$@"
			;;
		(setup)
			shift
			_setup "$@"
			;;
		(help | --help | -h)
			show_help 
			exit 0 
			;;
		(*)
			printf >&2 "Error: invalid command\n"
			show_help
			exit 1
			;;

	esac
	

}

main "$@"
