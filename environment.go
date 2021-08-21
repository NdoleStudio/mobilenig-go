package mobilenig

// Environment is the URL of the mobilenig environment
type Environment string

const (
	// TestEnvironment is the test Environment
	TestEnvironment = Environment("TEST")

	// LiveEnvironment is the production Environment
	LiveEnvironment = Environment("LIVE")
)

func (e Environment) String() string {
	return string(e)
}
