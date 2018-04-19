package main

import (
	// Import the entire framework for interracting with SDAccel from Go (including bundled verilog)
	_ "github.com/ReconfigureIO/sdaccel"

	// Use the new AXI protocol package for interracting with memory
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"
)

func Top(
	// Specify inputs and outputs to and from the FPGA. Tell the FPGA where to find data in shared memory, what data type
	// to expect or pass single integers directly to the FPGA by sending them to the control register

	in uint32,
	out uintptr,

	// Set up channels for interacting with the shared memory
	//memReadAddr chan<- axiprotocol.Addr,
	//memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	// Do whatever needs doing with the data from the host

	result := in * 2

	// Write the result to the location in shared memory as requested by the host
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, true, out, result)
}
