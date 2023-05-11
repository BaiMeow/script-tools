package cmd

type Commander interface {
	// Run command and get its output
	Run(cmd string) ([]byte, error)
	// Send command but not get its output
	Send(cmd string) error
	// Close the commander
	Close() error
}
