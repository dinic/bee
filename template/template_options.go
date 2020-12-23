package template

var tplOption string = `echo "auto option ..."
help=no

ProjectHolder_PREFIX=
ProjectHolder_BIN=
ProjectHolder_CONF=
VERSION=$$$$cat VERSION$$$$
opt=

for option
do
    opt="$opt $$$$echo $option | sed -e \"s/\(--[^=]*=\)\(.* .*\)/\1'\2'/\"$$$$"
    case "$option" in
        --*=*) value=$$$$echo "$option" | sed -e 's/[-_a-zA-Z0-9]*=//'$$$$ ;;
        *) value="" ;;
    esac

    case "$option" in
        --help)                 help=yes;;
        --prefix=*)             ProjectHolder_PREFIX="$value";;
        --bin-dir=*)            ProjectHolder_BIN="$value";;
		--conf-dir=*)           ProjectHolder_CONF="$value";;
		--build-dir=*)			ProjectHolder_BUILD_DIR="$value";;
        *)
            echo "$0: error: invalid option \"$option\""
            exit 1
        ;;
    esac

done

if [ $help = yes ]; then
cat << END

    --help                  print this message
    --prefix=PATH           set installation prefix
    --bin-dir=PATH          set bin dir
    --conf-dir=PATH         set conf dir
	--build-dir=PATH		set build dir
END
	exit 1
fi
`
