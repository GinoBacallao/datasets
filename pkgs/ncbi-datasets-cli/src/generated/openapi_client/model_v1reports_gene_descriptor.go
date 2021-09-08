/*
 * NCBI Datasets API
 *
 * ### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1reportsGeneDescriptor struct for V1reportsGeneDescriptor
type V1reportsGeneDescriptor struct {
	GeneId *string `json:"gene_id,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Description *string `json:"description,omitempty"`
	TaxId *string `json:"tax_id,omitempty"`
	Taxname *string `json:"taxname,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	Type *V1reportsGeneDescriptorGeneType `json:"type,omitempty"`
	RnaType *V1reportsGeneDescriptorRnaType `json:"rna_type,omitempty"`
	Orientation *V1reportsOrientation `json:"orientation,omitempty"`
	GenomicRanges *[]V1reportsSeqRangeSet `json:"genomic_ranges,omitempty"`
	ReferenceStandards *[]V1reportsGenomicRegion `json:"reference_standards,omitempty"`
	GenomicRegions *[]V1reportsGenomicRegion `json:"genomic_regions,omitempty"`
	Transcripts *[]V1reportsTranscript `json:"transcripts,omitempty"`
	Proteins *[]V1reportsProtein `json:"proteins,omitempty"`
	Chromosome *string `json:"chromosome,omitempty"`
	Chromosomes *[]string `json:"chromosomes,omitempty"`
	NomenclatureAuthority *V1reportsNomenclatureAuthority `json:"nomenclature_authority,omitempty"`
	SwissProtAccessions *[]string `json:"swiss_prot_accessions,omitempty"`
	EnsemblGeneIds *[]string `json:"ensembl_gene_ids,omitempty"`
	OmimIds *[]string `json:"omim_ids,omitempty"`
	Synonyms *[]string `json:"synonyms,omitempty"`
	ReplacedGeneId *string `json:"replaced_gene_id,omitempty"`
	Annotations *[]V1reportsAnnotation `json:"annotations,omitempty"`
}

// NewV1reportsGeneDescriptor instantiates a new V1reportsGeneDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsGeneDescriptor() *V1reportsGeneDescriptor {
	this := V1reportsGeneDescriptor{}
	var type_ V1reportsGeneDescriptorGeneType = V1REPORTSGENEDESCRIPTORGENETYPE_UNKNOWN
	this.Type = &type_
	var rnaType V1reportsGeneDescriptorRnaType = V1REPORTSGENEDESCRIPTORRNATYPE_RNA_UNKNOWN
	this.RnaType = &rnaType
	var orientation V1reportsOrientation = V1REPORTSORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// NewV1reportsGeneDescriptorWithDefaults instantiates a new V1reportsGeneDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsGeneDescriptorWithDefaults() *V1reportsGeneDescriptor {
	this := V1reportsGeneDescriptor{}
	var type_ V1reportsGeneDescriptorGeneType = V1REPORTSGENEDESCRIPTORGENETYPE_UNKNOWN
	this.Type = &type_
	var rnaType V1reportsGeneDescriptorRnaType = V1REPORTSGENEDESCRIPTORRNATYPE_RNA_UNKNOWN
	this.RnaType = &rnaType
	var orientation V1reportsOrientation = V1REPORTSORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// GetGeneId returns the GeneId field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetGeneId() string {
	if o == nil || o.GeneId == nil {
		var ret string
		return ret
	}
	return *o.GeneId
}

// GetGeneIdOk returns a tuple with the GeneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetGeneIdOk() (*string, bool) {
	if o == nil || o.GeneId == nil {
		return nil, false
	}
	return o.GeneId, true
}

// HasGeneId returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasGeneId() bool {
	if o != nil && o.GeneId != nil {
		return true
	}

	return false
}

// SetGeneId gets a reference to the given string and assigns it to the GeneId field.
func (o *V1reportsGeneDescriptor) SetGeneId(v string) {
	o.GeneId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *V1reportsGeneDescriptor) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V1reportsGeneDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetTaxId() string {
	if o == nil || o.TaxId == nil {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetTaxIdOk() (*string, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *V1reportsGeneDescriptor) SetTaxId(v string) {
	o.TaxId = &v
}

// GetTaxname returns the Taxname field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetTaxname() string {
	if o == nil || o.Taxname == nil {
		var ret string
		return ret
	}
	return *o.Taxname
}

// GetTaxnameOk returns a tuple with the Taxname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetTaxnameOk() (*string, bool) {
	if o == nil || o.Taxname == nil {
		return nil, false
	}
	return o.Taxname, true
}

// HasTaxname returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasTaxname() bool {
	if o != nil && o.Taxname != nil {
		return true
	}

	return false
}

// SetTaxname gets a reference to the given string and assigns it to the Taxname field.
func (o *V1reportsGeneDescriptor) SetTaxname(v string) {
	o.Taxname = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *V1reportsGeneDescriptor) SetCommonName(v string) {
	o.CommonName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetType() V1reportsGeneDescriptorGeneType {
	if o == nil || o.Type == nil {
		var ret V1reportsGeneDescriptorGeneType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetTypeOk() (*V1reportsGeneDescriptorGeneType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given V1reportsGeneDescriptorGeneType and assigns it to the Type field.
func (o *V1reportsGeneDescriptor) SetType(v V1reportsGeneDescriptorGeneType) {
	o.Type = &v
}

// GetRnaType returns the RnaType field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetRnaType() V1reportsGeneDescriptorRnaType {
	if o == nil || o.RnaType == nil {
		var ret V1reportsGeneDescriptorRnaType
		return ret
	}
	return *o.RnaType
}

// GetRnaTypeOk returns a tuple with the RnaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetRnaTypeOk() (*V1reportsGeneDescriptorRnaType, bool) {
	if o == nil || o.RnaType == nil {
		return nil, false
	}
	return o.RnaType, true
}

// HasRnaType returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasRnaType() bool {
	if o != nil && o.RnaType != nil {
		return true
	}

	return false
}

// SetRnaType gets a reference to the given V1reportsGeneDescriptorRnaType and assigns it to the RnaType field.
func (o *V1reportsGeneDescriptor) SetRnaType(v V1reportsGeneDescriptorRnaType) {
	o.RnaType = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetOrientation() V1reportsOrientation {
	if o == nil || o.Orientation == nil {
		var ret V1reportsOrientation
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetOrientationOk() (*V1reportsOrientation, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given V1reportsOrientation and assigns it to the Orientation field.
func (o *V1reportsGeneDescriptor) SetOrientation(v V1reportsOrientation) {
	o.Orientation = &v
}

// GetGenomicRanges returns the GenomicRanges field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetGenomicRanges() []V1reportsSeqRangeSet {
	if o == nil || o.GenomicRanges == nil {
		var ret []V1reportsSeqRangeSet
		return ret
	}
	return *o.GenomicRanges
}

// GetGenomicRangesOk returns a tuple with the GenomicRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetGenomicRangesOk() (*[]V1reportsSeqRangeSet, bool) {
	if o == nil || o.GenomicRanges == nil {
		return nil, false
	}
	return o.GenomicRanges, true
}

// HasGenomicRanges returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasGenomicRanges() bool {
	if o != nil && o.GenomicRanges != nil {
		return true
	}

	return false
}

// SetGenomicRanges gets a reference to the given []V1reportsSeqRangeSet and assigns it to the GenomicRanges field.
func (o *V1reportsGeneDescriptor) SetGenomicRanges(v []V1reportsSeqRangeSet) {
	o.GenomicRanges = &v
}

// GetReferenceStandards returns the ReferenceStandards field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetReferenceStandards() []V1reportsGenomicRegion {
	if o == nil || o.ReferenceStandards == nil {
		var ret []V1reportsGenomicRegion
		return ret
	}
	return *o.ReferenceStandards
}

// GetReferenceStandardsOk returns a tuple with the ReferenceStandards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetReferenceStandardsOk() (*[]V1reportsGenomicRegion, bool) {
	if o == nil || o.ReferenceStandards == nil {
		return nil, false
	}
	return o.ReferenceStandards, true
}

// HasReferenceStandards returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasReferenceStandards() bool {
	if o != nil && o.ReferenceStandards != nil {
		return true
	}

	return false
}

// SetReferenceStandards gets a reference to the given []V1reportsGenomicRegion and assigns it to the ReferenceStandards field.
func (o *V1reportsGeneDescriptor) SetReferenceStandards(v []V1reportsGenomicRegion) {
	o.ReferenceStandards = &v
}

// GetGenomicRegions returns the GenomicRegions field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetGenomicRegions() []V1reportsGenomicRegion {
	if o == nil || o.GenomicRegions == nil {
		var ret []V1reportsGenomicRegion
		return ret
	}
	return *o.GenomicRegions
}

// GetGenomicRegionsOk returns a tuple with the GenomicRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetGenomicRegionsOk() (*[]V1reportsGenomicRegion, bool) {
	if o == nil || o.GenomicRegions == nil {
		return nil, false
	}
	return o.GenomicRegions, true
}

// HasGenomicRegions returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasGenomicRegions() bool {
	if o != nil && o.GenomicRegions != nil {
		return true
	}

	return false
}

// SetGenomicRegions gets a reference to the given []V1reportsGenomicRegion and assigns it to the GenomicRegions field.
func (o *V1reportsGeneDescriptor) SetGenomicRegions(v []V1reportsGenomicRegion) {
	o.GenomicRegions = &v
}

// GetTranscripts returns the Transcripts field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetTranscripts() []V1reportsTranscript {
	if o == nil || o.Transcripts == nil {
		var ret []V1reportsTranscript
		return ret
	}
	return *o.Transcripts
}

// GetTranscriptsOk returns a tuple with the Transcripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetTranscriptsOk() (*[]V1reportsTranscript, bool) {
	if o == nil || o.Transcripts == nil {
		return nil, false
	}
	return o.Transcripts, true
}

// HasTranscripts returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasTranscripts() bool {
	if o != nil && o.Transcripts != nil {
		return true
	}

	return false
}

// SetTranscripts gets a reference to the given []V1reportsTranscript and assigns it to the Transcripts field.
func (o *V1reportsGeneDescriptor) SetTranscripts(v []V1reportsTranscript) {
	o.Transcripts = &v
}

// GetProteins returns the Proteins field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetProteins() []V1reportsProtein {
	if o == nil || o.Proteins == nil {
		var ret []V1reportsProtein
		return ret
	}
	return *o.Proteins
}

// GetProteinsOk returns a tuple with the Proteins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetProteinsOk() (*[]V1reportsProtein, bool) {
	if o == nil || o.Proteins == nil {
		return nil, false
	}
	return o.Proteins, true
}

// HasProteins returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasProteins() bool {
	if o != nil && o.Proteins != nil {
		return true
	}

	return false
}

// SetProteins gets a reference to the given []V1reportsProtein and assigns it to the Proteins field.
func (o *V1reportsGeneDescriptor) SetProteins(v []V1reportsProtein) {
	o.Proteins = &v
}

// GetChromosome returns the Chromosome field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetChromosome() string {
	if o == nil || o.Chromosome == nil {
		var ret string
		return ret
	}
	return *o.Chromosome
}

// GetChromosomeOk returns a tuple with the Chromosome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetChromosomeOk() (*string, bool) {
	if o == nil || o.Chromosome == nil {
		return nil, false
	}
	return o.Chromosome, true
}

// HasChromosome returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasChromosome() bool {
	if o != nil && o.Chromosome != nil {
		return true
	}

	return false
}

// SetChromosome gets a reference to the given string and assigns it to the Chromosome field.
func (o *V1reportsGeneDescriptor) SetChromosome(v string) {
	o.Chromosome = &v
}

// GetChromosomes returns the Chromosomes field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetChromosomes() []string {
	if o == nil || o.Chromosomes == nil {
		var ret []string
		return ret
	}
	return *o.Chromosomes
}

// GetChromosomesOk returns a tuple with the Chromosomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetChromosomesOk() (*[]string, bool) {
	if o == nil || o.Chromosomes == nil {
		return nil, false
	}
	return o.Chromosomes, true
}

// HasChromosomes returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasChromosomes() bool {
	if o != nil && o.Chromosomes != nil {
		return true
	}

	return false
}

// SetChromosomes gets a reference to the given []string and assigns it to the Chromosomes field.
func (o *V1reportsGeneDescriptor) SetChromosomes(v []string) {
	o.Chromosomes = &v
}

// GetNomenclatureAuthority returns the NomenclatureAuthority field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetNomenclatureAuthority() V1reportsNomenclatureAuthority {
	if o == nil || o.NomenclatureAuthority == nil {
		var ret V1reportsNomenclatureAuthority
		return ret
	}
	return *o.NomenclatureAuthority
}

// GetNomenclatureAuthorityOk returns a tuple with the NomenclatureAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetNomenclatureAuthorityOk() (*V1reportsNomenclatureAuthority, bool) {
	if o == nil || o.NomenclatureAuthority == nil {
		return nil, false
	}
	return o.NomenclatureAuthority, true
}

// HasNomenclatureAuthority returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasNomenclatureAuthority() bool {
	if o != nil && o.NomenclatureAuthority != nil {
		return true
	}

	return false
}

// SetNomenclatureAuthority gets a reference to the given V1reportsNomenclatureAuthority and assigns it to the NomenclatureAuthority field.
func (o *V1reportsGeneDescriptor) SetNomenclatureAuthority(v V1reportsNomenclatureAuthority) {
	o.NomenclatureAuthority = &v
}

// GetSwissProtAccessions returns the SwissProtAccessions field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetSwissProtAccessions() []string {
	if o == nil || o.SwissProtAccessions == nil {
		var ret []string
		return ret
	}
	return *o.SwissProtAccessions
}

// GetSwissProtAccessionsOk returns a tuple with the SwissProtAccessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetSwissProtAccessionsOk() (*[]string, bool) {
	if o == nil || o.SwissProtAccessions == nil {
		return nil, false
	}
	return o.SwissProtAccessions, true
}

// HasSwissProtAccessions returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasSwissProtAccessions() bool {
	if o != nil && o.SwissProtAccessions != nil {
		return true
	}

	return false
}

// SetSwissProtAccessions gets a reference to the given []string and assigns it to the SwissProtAccessions field.
func (o *V1reportsGeneDescriptor) SetSwissProtAccessions(v []string) {
	o.SwissProtAccessions = &v
}

// GetEnsemblGeneIds returns the EnsemblGeneIds field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetEnsemblGeneIds() []string {
	if o == nil || o.EnsemblGeneIds == nil {
		var ret []string
		return ret
	}
	return *o.EnsemblGeneIds
}

// GetEnsemblGeneIdsOk returns a tuple with the EnsemblGeneIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetEnsemblGeneIdsOk() (*[]string, bool) {
	if o == nil || o.EnsemblGeneIds == nil {
		return nil, false
	}
	return o.EnsemblGeneIds, true
}

// HasEnsemblGeneIds returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasEnsemblGeneIds() bool {
	if o != nil && o.EnsemblGeneIds != nil {
		return true
	}

	return false
}

// SetEnsemblGeneIds gets a reference to the given []string and assigns it to the EnsemblGeneIds field.
func (o *V1reportsGeneDescriptor) SetEnsemblGeneIds(v []string) {
	o.EnsemblGeneIds = &v
}

// GetOmimIds returns the OmimIds field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetOmimIds() []string {
	if o == nil || o.OmimIds == nil {
		var ret []string
		return ret
	}
	return *o.OmimIds
}

// GetOmimIdsOk returns a tuple with the OmimIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetOmimIdsOk() (*[]string, bool) {
	if o == nil || o.OmimIds == nil {
		return nil, false
	}
	return o.OmimIds, true
}

// HasOmimIds returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasOmimIds() bool {
	if o != nil && o.OmimIds != nil {
		return true
	}

	return false
}

// SetOmimIds gets a reference to the given []string and assigns it to the OmimIds field.
func (o *V1reportsGeneDescriptor) SetOmimIds(v []string) {
	o.OmimIds = &v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return *o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetSynonymsOk() (*[]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *V1reportsGeneDescriptor) SetSynonyms(v []string) {
	o.Synonyms = &v
}

// GetReplacedGeneId returns the ReplacedGeneId field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetReplacedGeneId() string {
	if o == nil || o.ReplacedGeneId == nil {
		var ret string
		return ret
	}
	return *o.ReplacedGeneId
}

// GetReplacedGeneIdOk returns a tuple with the ReplacedGeneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetReplacedGeneIdOk() (*string, bool) {
	if o == nil || o.ReplacedGeneId == nil {
		return nil, false
	}
	return o.ReplacedGeneId, true
}

// HasReplacedGeneId returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasReplacedGeneId() bool {
	if o != nil && o.ReplacedGeneId != nil {
		return true
	}

	return false
}

// SetReplacedGeneId gets a reference to the given string and assigns it to the ReplacedGeneId field.
func (o *V1reportsGeneDescriptor) SetReplacedGeneId(v string) {
	o.ReplacedGeneId = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *V1reportsGeneDescriptor) GetAnnotations() []V1reportsAnnotation {
	if o == nil || o.Annotations == nil {
		var ret []V1reportsAnnotation
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsGeneDescriptor) GetAnnotationsOk() (*[]V1reportsAnnotation, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *V1reportsGeneDescriptor) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []V1reportsAnnotation and assigns it to the Annotations field.
func (o *V1reportsGeneDescriptor) SetAnnotations(v []V1reportsAnnotation) {
	o.Annotations = &v
}

func (o V1reportsGeneDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeneId != nil  {
		toSerialize["gene_id"] = o.GeneId
	}
	if o.Symbol != nil  {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Description != nil  {
		toSerialize["description"] = o.Description
	}
	if o.TaxId != nil  {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.Taxname != nil  {
		toSerialize["taxname"] = o.Taxname
	}
	if o.CommonName != nil  {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Type != nil  {
		toSerialize["type"] = o.Type
	}
	if o.RnaType != nil  {
		toSerialize["rna_type"] = o.RnaType
	}
	if o.Orientation != nil  {
		toSerialize["orientation"] = o.Orientation
	}
	if o.GenomicRanges != nil && len(o.GetGenomicRanges()) > 0  {
		toSerialize["genomic_ranges"] = o.GenomicRanges
	}
	if o.ReferenceStandards != nil && len(o.GetReferenceStandards()) > 0  {
		toSerialize["reference_standards"] = o.ReferenceStandards
	}
	if o.GenomicRegions != nil && len(o.GetGenomicRegions()) > 0  {
		toSerialize["genomic_regions"] = o.GenomicRegions
	}
	if o.Transcripts != nil && len(o.GetTranscripts()) > 0  {
		toSerialize["transcripts"] = o.Transcripts
	}
	if o.Proteins != nil && len(o.GetProteins()) > 0  {
		toSerialize["proteins"] = o.Proteins
	}
	if o.Chromosome != nil  {
		toSerialize["chromosome"] = o.Chromosome
	}
	if o.Chromosomes != nil && len(o.GetChromosomes()) > 0  {
		toSerialize["chromosomes"] = o.Chromosomes
	}
	if o.NomenclatureAuthority != nil  {
		toSerialize["nomenclature_authority"] = o.NomenclatureAuthority
	}
	if o.SwissProtAccessions != nil && len(o.GetSwissProtAccessions()) > 0  {
		toSerialize["swiss_prot_accessions"] = o.SwissProtAccessions
	}
	if o.EnsemblGeneIds != nil && len(o.GetEnsemblGeneIds()) > 0  {
		toSerialize["ensembl_gene_ids"] = o.EnsemblGeneIds
	}
	if o.OmimIds != nil && len(o.GetOmimIds()) > 0  {
		toSerialize["omim_ids"] = o.OmimIds
	}
	if o.Synonyms != nil && len(o.GetSynonyms()) > 0  {
		toSerialize["synonyms"] = o.Synonyms
	}
	if o.ReplacedGeneId != nil  {
		toSerialize["replaced_gene_id"] = o.ReplacedGeneId
	}
	if o.Annotations != nil && len(o.GetAnnotations()) > 0  {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsGeneDescriptor struct {
	value *V1reportsGeneDescriptor
	isSet bool
}

func (v NullableV1reportsGeneDescriptor) Get() *V1reportsGeneDescriptor {
	return v.value
}

func (v *NullableV1reportsGeneDescriptor) Set(val *V1reportsGeneDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsGeneDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsGeneDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsGeneDescriptor(val *V1reportsGeneDescriptor) *NullableV1reportsGeneDescriptor {
	return &NullableV1reportsGeneDescriptor{value: val, isSet: true}
}

func (v NullableV1reportsGeneDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsGeneDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

