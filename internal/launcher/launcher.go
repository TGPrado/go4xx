package launcher

import (
	"github.com/spf13/cobra"
)

type Flags struct {
	Headers  []string
	ExC      bool
	ExH      bool
	ExP      bool
	ExE      bool
	V        bool
	FileName string
	Output   string
}

func ParseFlags() *Flags {
	var flags Flags

	var rootCmd = &cobra.Command{
		Use:   "go4xx",
		Short: "A golang app to bypass 4xx status code.",
		Long:  "A golang app to bypass 4xx status code.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.Flags().StringSliceVarP(&flags.Headers, "header", "H", []string{}, "Headers send in request.")
	rootCmd.Flags().BoolVar(&flags.ExC, "exC", false, "Remove credentials from payload.")
	rootCmd.Flags().BoolVar(&flags.ExP, "exP", false, "Remove paths from payload.")
	rootCmd.Flags().BoolVar(&flags.ExE, "exE", false, "Remove extensions from payload.")
	rootCmd.Flags().BoolVar(&flags.ExH, "exH", false, "Remove headers from payload.")
	rootCmd.Flags().BoolVarP(&flags.V, "verbose", "v", false, "Verbose mode.")
	rootCmd.Flags().StringVarP(&flags.FileName, "file", "f", "", "File to read.")
	rootCmd.Flags().StringVarP(&flags.Output, "output", "o", "", "Output.")

	if err := rootCmd.Execute(); err != nil {
	}

	return &flags
}
