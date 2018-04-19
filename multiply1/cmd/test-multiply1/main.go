package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/ReconfigureIO/sdaccel/xcl"
)

func main() {
	// Allocate a 'world' for interacting with the FPGA
	world := xcl.NewWorld()
	defer world.Release()

	// Import the compiled code that will be loaded onto the FPGA (referred to here as a kernel)
	// Right now these two identifiers are hard coded as an output from the build process
	krnl := world.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl.Release()

	// Allocate a space in the shared memory to store the data you're sending to the FPGA and space
	// for the results from the FPGA
	//inputBuff := world.Malloc(xcl.ReadOnly, <size here>)
	//defer inputBuff.Free()

	outputBuff := world.Malloc(xcl.WriteOnly, 4)
	defer outputBuff.Free()

	// Create/get data and pass arguments to the FPGA as required. These could be small pieces of data,
	// pointers to memory, data lengths so the FPGA knows what to expect. This all depends on your project.
	// Usually, you will send data via shared memory, so you will need to write it to the space you allocated
	// above before passing the pointer to the FPGA.
	// We have passed three arguments here, you can pass more as neccessary

	// First argument
	krnl.SetArg(0, 3)
	// Second argument
	krnl.SetMemoryArg(1, outputBuff)

	// Run the FPGA with the supplied arguments. This is the same for all projects.
	// The arguments ``(1, 1, 1)`` relate to x, y, z co-ordinates and correspond to our current
	// underlying technology.
	krnl.Run(1, 1, 1)

	// Display/use the results returned from the FPGA as required!
	var result uint32
	err := binary.Read(outputBuff.Reader(), binary.LittleEndian, &result)
	if err != nil {
		log.Fatalln("Failed to read back result")
	}

	fmt.Println("Result: ", result)
}
