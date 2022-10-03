# six5go2

Six5go2 - 6502 Emulator and Disassembler in Golang (c) 2022 Zayn Otley

USAGE   - ./six5go2 <target_filename> <hex_entry_point> <dis>/<mon> (Disassembler/Machine Monitor) <hex> (Hex opcodes as comments with disassembly)

EXAMPLE - ./six5go2 AllSuiteA.bin 4000 mon

EXAMPLE - ./six5go2 AllSuiteA.bin 4000 dis

EXAMPLE - ./six5go2 AllSuiteA.bin 4000 dis hex

Choose Disassembler or Machine Monitor at command line with dis or mon parameter.

Specify hex as optional parameter with the disassembler to have opcodes as comments in the source output.


To build the project, run the following command:

    git clone https://github.com/intuitionamiga/six5go2.git
    cd six5go2
    go build .

To run the disassembler on the AllSuiteA 6502 opcode test suite, run the following command:
    ./six5go2 AllSuiteA.bin 4000 dis

To run the disassembler with hex opcodes as comments, run the following command:
    ./six5go2 AllSuiteA.bin 4000 dis hex

To run the machine monitor, run the following command:
    ./six5go2 AllSuiteA.bin 4000 mon
