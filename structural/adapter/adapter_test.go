package adapter

import "testing"

func TestAdapter(t *testing.T) {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	client.insertSquareUsbInComputer(windowsMachineAdapter)

	/*
	   Insert square port into mac machine
	   Insert circle port into windows machine
	*/
}
