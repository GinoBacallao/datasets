package datasets

import (
	"errors"

	"github.com/spf13/cobra"
)

var assemblyCmd = &cobra.Command{
	Deprecated: "please use \"datasets download genome\".",
	Use:        "assembly <accession ...>",
	Short:      "Download genome assembly data",
	Example:    "  datasets download assembly GCF_000001405.39 GCF_000001635.27",
	Long: `
Download data by assembly accession. Data is returned as a zip archive.
The default download package for a given assembly (or set of assemblies) includes all chromosomes and
unlocalized sequences and excludes any available annotation data.

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) documentation for information about getting started with the command-line tools.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(argIDArgs) == 0 {
			return errors.New("Input accessions not specified")
		}
		validAccList, warning, err := checkAssemblyAvailability(argIDArgs)
		if err != nil {
			return err
		}
		if warning != "" {
			cmd.Println(warning)
		}
		downloadOpts := getDownloadRequest(validAccList)
		err = downloadAssembly(downloadOpts, argDownloadFilename)
		return nil
	},
}

func init() {
	assemblyCmd.AddCommand(downloadAssmTaxonCmd)

	assemblyCmd.PersistentFlags().StringVar(&argInputFile, "inputfile", "", "file to read list of assembly accessions")
	assemblyCmd.PersistentFlags().StringSliceVarP(&argChromosomes, "chromosomes", "c", []string{"all"}, "limit to a specified, comma-delimited list of chromosomes")

	assemblyCmd.PersistentFlags().BoolVarP(&argExcludeSeq, "exclude-seq", "s", false, "Exclude genomic sequence")
	assemblyCmd.PersistentFlags().BoolVarP(&argExcludeRna, "exclude-rna", "r", false, "Exclude RNA sequence data")
	assemblyCmd.PersistentFlags().BoolVarP(&argExcludeProtein, "exclude-protein", "p", false, "Exclude protein sequence file")
	assemblyCmd.PersistentFlags().BoolVarP(&argExcludeGff3, "exclude-gff3", "g", false, "Exclude gff3 annotation file")
	assemblyCmd.PersistentFlags().BoolVarP(&argIncludeGbff, "include-gbff", "b", false, "Include gbff annotation file, if available")
	assemblyCmd.PersistentFlags().BoolVarP(&argIncludeGtf, "include-gtf", "e", false, "Include gtf annotation file, if available")
	assemblyCmd.PersistentFlags().BoolVarP(&argDehydrated, "dehydrated", "", false, "Download minimal package that includes data report and locations of data files. Use the rehydrate command to retrieve data files when needed.")
}
