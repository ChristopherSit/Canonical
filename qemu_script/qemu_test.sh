#!/bin/bash

# download and compile linux kernel into a bootable kernel image
wget https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.10.50.tar.xz
tar -xf linux-5.10.50.tar.xz
cd linux-5.10.50
make defconfig
make -j$(nproc) bzImage
cd ..

# create an very basic c program to be our init
cat > init.c << EOF
#include <stdio.h>
#include <unistd.h>
#include <sys/reboot.h>

void main() {
    printf("Hello, world!\n");
    reboot(RB_POWER_OFF);
}
EOF
gcc -static -o init init.c

# create initial ram file system which basically only has init in it
mkdir -p rootfs/{dev,etc,proc,sys,usr}

# the script broke unless i included the directory structure
sudo mknod rootfs/dev/null c 1 3
sudo mknod rootfs/dev/tty c 5 0
sudo mknod rootfs/dev/console c 5 1
sudo mknod rootfs/dev/zero c 1 5

# move init to rootfs/init
cp init rootfs/init

# create archive of the rootfs directory which acts as the initial ram file system
cd rootfs
find . | cpio -H newc -o > ../initramfs.cpio
cd ..

# pray this works and added flags to have qemu run in console
qemu-system-x86_64 -kernel linux-5.10.50/arch/x86/boot/bzImage -initrd initramfs.cpio -append "console=ttyS0 quiet init=/init" -nographic
