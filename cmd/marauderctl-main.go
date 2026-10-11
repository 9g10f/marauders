package cmd

var Version = "0.0.2"

func MarauderS() error {
	Root.AddCommand(Enable())
	Root.AddCommand(Disable())
	Root.AddCommand(Start())
	Root.AddCommand(Run())

	err := Root.Execute()
	if err != nil {
		return err
	}

	return nil
}