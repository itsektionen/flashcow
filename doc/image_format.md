# The Receipt Image Phormat (.rip) version 0
It's little endian. Deal with IT.

x86 (and modern ARM and RISC-V) on top! Fuck big-endian MIPS and ARM and especially IBM Z.

Header:
- Magic number: 'RIPKISTA'
- SHA-256 Checksum (includes data past this point)
- Version: u16 = 0x0000
- Reserved: 14 bytes
- Width: u32
- Height: u32
- Pixel data row-major from the top as 8bpp grayscale, no padding
