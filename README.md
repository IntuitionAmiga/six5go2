# six5go2

Six5go2 - 6502 Emulator and Disassembler in Golang (c) 2022 Zayn Otley

USAGE   - ./six5go2 target_filename hex_entry_point dis/mon (Disassembler/Machine Monitor) hex (Hex opcodes as comments with disassembly)

EXAMPLE - ./six5go2 c64kernal.bin 0200 dis hex

EXAMPLE - ./six5go2 apple1basic.bin 0100 mon


Choose Disassembler or Machine Monitor at command line with dis or mon parameter.

Specify hex as optional parameter with the disassembler to have opcodes as comments in the source output.