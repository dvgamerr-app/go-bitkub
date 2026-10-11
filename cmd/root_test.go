package cmd

import "testing"

func TestRootRunsPublicCommandsWithoutCredentials(t *testing.T) {
	t.Setenv("BTK_APIKEY", "")
	t.Setenv("BTK_SECRET", "")
	apiKey = ""
	secretKey = ""

	rootCmd.PersistentPreRun(rootCmd, nil)
}
