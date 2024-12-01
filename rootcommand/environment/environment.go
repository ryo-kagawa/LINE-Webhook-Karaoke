package environment

type Environment struct {
	Database EnvironmentDatabase
	Line     EnvironmentLine
}

func GetEnvironment() Environment {
	return Environment{
		Database: GetEnvironmentDatabase(),
		Line:     GetEnvironmentLine(),
	}
}

func (e Environment) Validate() error {
	if err := e.Database.Validate(); err != nil {
		return err
	}
	if err := e.Line.Validate(); err != nil {
		return err
	}
	return nil
}
