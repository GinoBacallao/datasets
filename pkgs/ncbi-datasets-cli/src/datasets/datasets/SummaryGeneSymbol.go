package datasets

import (
	openapi "datasets_cli/v1/openapi"
	"errors"

	"github.com/spf13/cobra"
)

func cmdRunSummaryGeneSymbol(cmd *cobra.Command, args []string) (err error) {
	if len(argIDArgs) == 0 {
		return errors.New("No gene symbols provided")
	}

	req := openapi.NewV1GeneDatasetRequest()
	symbols_taxon := openapi.NewV1GeneDatasetRequestSymbolsForTaxon()
	symbols_taxon.SetSymbols(argIDArgs)
	symbols_taxon.SetTaxon(argTaxon)
	req.SetSymbolsForTaxon(*symbols_taxon)

	if argJsonLinesFormat {
		err = streamGeneMatch(req, &JsonLinesStreamProcessor{})
	} else {
		err = streamGeneMatch(req, &JsonStreamProcessor{wrapperName: "genes"})
	}
	return
}

var summaryGeneSymbolCmd = &cobra.Command{
	Use:   "symbol <gene symbol ...> [flags]",
	Short: "print a summary of a gene dataset by gene symbol",
	Long:  `Print a summary of a gene dataset by gene symbol and taxon (species name or species-level NCBI Taxonomy ID). If no taxon is specified, data will be returned for human. The summary is returned in JSON format.`,
	Example: `  datasets summary gene symbol tp53
  datasets summary gene symbol brca1 --taxon mouse`,
	RunE: cmdRunSummaryGeneSymbol,
}

func init() {
	flags := summaryGeneSymbolCmd.Flags()
	registerHiddenStringPair(flags, &argInputFile, "inputfile", "i", "", "read a list of gene symbols from a file to use as input")
	registerHiddenStringPair(flags, &argTaxon, "taxon", "t", "human", "specify a species name (common or scientific) or species-level NCBI Taxonomy ID")
}
