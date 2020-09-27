#!/bin/bash

pkgname=iwsp

# 编译多平台包
GOOS=linux go build -ldflags "-w -s -X main.version=$1" -o release/${pkgname}-linux .
GOOS=darwin go build -ldflags "-w -s -X main.version=$1" -o release/${pkgname}-macos .
GOOS=windows go build -ldflags "-w -s -X main.version=$1" -o release/${pkgname}-windows.exe .

# 得到md5
md5=$(md5sum release/${pkgname}-linux)

# 使用md5和版本号拼接pkgbuild
cat <<EOF > PKGBUILD
# Maintainer: amtoaer <amtoaer@outlook.com>
pkgname=iwsp-bin
pkgver=$1
pkgrel=1
pkgdesc="Northeastern University Gymnasium Reservation."
arch=('x86_64')
url="https://github.com/amtoaer/iwsp"
license=('MIT')

source=(
        "iwsp-linux::https://github.com/amtoaer/iwsp/releases/download/\${pkgver}/iwsp-linux"
        "LICENSE::https://raw.githubusercontent.com/amtoaer/iwsp/master/LICENSE"
)

md5sums=(
        '${md5:0:32}'
        '31fad0aacc583d621612630ce8f5a26c'
)

package(){
        install -D -m 755 \$srcdir/iwsp-linux \$pkgdir/usr/bin/iwsp
        install -D -m 644 \$srcdir/LICENSE \$pkgdir/usr/share/licenses/\$pkgname/LICENSE
}
EOF
