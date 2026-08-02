// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// CompetitionRankingGetParam is a type within the Ranking protocol
type CompetitionRankingGetParam struct {
	types.Structure
	Unknown     types.UInt32
	ResultRange types.ResultRange
	FestivalIds types.List[types.UInt32]
}

// WriteTo writes the CompetitionRankingGetParam to the given writable
func (crgp CompetitionRankingGetParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	crgp.Unknown.WriteTo(contentWritable)
	crgp.ResultRange.WriteTo(contentWritable)
	crgp.FestivalIds.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	crgp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the CompetitionRankingGetParam from the given readable
func (crgp CompetitionRankingGetParam) ExtractFrom(readable types.Readable) error {
	if err := crgp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingGetParam header. %s", err.Error())
	}

	if err := crgp.Unknown.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingGetParam.Unknown. %s", err.Error())
	}

	if err := crgp.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingGetParam.ResultRange. %s", err.Error())
	}

	if err := crgp.FestivalIds.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingGetParam.FestivalIds. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CompetitionRankingGetParam
func (crgp CompetitionRankingGetParam) Copy() types.RVType {
	copied := NewCompetitionRankingGetParam()

	copied.StructureVersion = crgp.StructureVersion
	copied.Unknown = crgp.Unknown.Copy().(types.UInt32)
	copied.ResultRange = crgp.Copy().(types.ResultRange)
	copied.FestivalIds = crgp.FestivalIds.Copy().(types.List[types.UInt32])

	return copied
}

// Equals checks if the given CompetitionRankingGetParam contains the same data as the current CompetitionRankingGetParam
func (crgp CompetitionRankingGetParam) Equals(o types.RVType) bool {
	if _, ok := o.(CompetitionRankingGetParam); !ok {
		return false
	}

	other := o.(CompetitionRankingGetParam)

	if crgp.StructureVersion != other.StructureVersion {
		return false
	}

	if !crgp.Unknown.Equals(other.Unknown) {
		return false
	}

	if !crgp.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return crgp.FestivalIds.Equals(other.FestivalIds)
}

// CopyRef copies the current value of the CompetitionRankingGetParam
// and returns a pointer to the new copy
func (crgp CompetitionRankingGetParam) CopyRef() types.RVTypePtr {
	copied := crgp.Copy().(CompetitionRankingGetParam)
	return &copied
}

// Deref takes a pointer to the CompetitionRankingGetParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (crgp *CompetitionRankingGetParam) Deref() types.RVType {
	return *crgp
}

// String returns the string representation of the CompetitionRankingGetParam
func (crgp CompetitionRankingGetParam) String() string {
	return crgp.FormatToString(0)
}

// FormatToString pretty-prints the CompetitionRankingGetParam using the provided indentation level
func (crgp CompetitionRankingGetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CompetitionRankingGetParam{\n")
	fmt.Fprintf(&b, "%sUnknown: %s,\n", indentationValues, crgp.Unknown)
	fmt.Fprintf(&b, "%sResultRange: %s,\n", indentationValues, crgp.ResultRange)
	fmt.Fprintf(&b, "%sFestivalIds: %s,\n", indentationValues, crgp.FestivalIds)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewCompetitionRankingGetParam returns a new CompetitionRankingGetParam
func NewCompetitionRankingGetParam() CompetitionRankingGetParam {
	return CompetitionRankingGetParam{
		Structure:   types.Structure{StructureVersion: 1},
		Unknown:     types.NewUInt32(0),
		ResultRange: types.NewResultRange(),
		FestivalIds: types.NewList[types.UInt32](),
	}

}
