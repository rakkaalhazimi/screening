package main

import (
	"fmt"
	"github.com/codingsince1985/checksum"
)

func main() {
	file := "./compare/target/folder1/subfolder1/hello.txt"

	md5, _ := checksum.MD5sum(file)
	fmt.Println(md5)

	sha256, _ := checksum.SHA256sum(file)
	fmt.Println(sha256)

	sha1, _ := checksum.SHA1sum(file)
	fmt.Println(sha1)

	crc32, _ := checksum.CRC32(file)
	fmt.Println(crc32)

	blake2s256, _ := checksum.Blake2s256(file)
	fmt.Println(blake2s256)
}
