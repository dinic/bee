package template

var tplConfig string = `echo "auto config ..."

VERSION=$$$$cat VERSION$$$$
ProjectHolder_AUTO_CONFIG="${ProjectHolder_ROOT}/auto/auto_config.go"

cat << END > $ProjectHolder_AUTO_CONFIG
package auto

const (
	VERION = "$VERSION"
	Prefix  = "$ProjectHolder_PREFIX/"
	ConfDir = "$ProjectHolder_CONF/"
)
END

go fmt $ProjectHolder_AUTO_CONFIG > /dev/null

ProjectHolder_BUILDINFO="${ProjectHolder_ROOT}/auto/buildinfo"
cat << END > $ProjectHolder_BUILDINFO
ProjectHolder_ROOT=$ProjectHolder_ROOT
ProjectHolder_PREFIX=$ProjectHolder_PREFIX
BINPATH=$ProjectHolder_BIN
CONFPATH=$ProjectHolder_CONF
BUILDPATH=$ProjectHolder_ROOT
VERSION=$VERSION
END
`
