package datasets

import (
	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

var summaryGenomeTaxonCmd = &cobra.Command{
	Use:   "taxon",
	Short: "print a summary of a genome dataset by taxon (NCBI Taxonomy ID, scientific or common name at any tax rank)",
	Long: `
Print a summary of a genome dataset by taxon (NCBI Taxonomy ID, scientific or common name at any tax rank). The summary is returned in JSON format.

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets summary genome taxon human
  datasets summary genome taxon "mus musculus"
  datasets summary genome taxon 10116`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		request := openapi.NewV1AssemblyMetadataRequest()
		request.SetTaxon(args[0])

		err = updateAssemblyMetadataRequestOption(request)
		if err != nil {
			return err
		}

		metadata_err := printAssemblyMetadataWithPost(request)
		if metadata_err != nil {
			return metadata_err
		}
		return nil
	},
}

func init() {
	summaryGenomeTaxonCmd.PersistentFlags().BoolVar(&argTaxExactMatch, "tax-exact-match", false, "exclude sub-species when a species-level taxon is specified")
}
