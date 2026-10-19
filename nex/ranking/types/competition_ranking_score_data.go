// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// CompetitionRankingScoreData is a type within the Ranking protocol
type CompetitionRankingScoreData struct {
	types.Structure
	Unknown1 types.UInt32
	PID      types.PID
	Score    types.UInt32
	Modified types.DateTime
	Unknown2 types.UInt8
	AppData  types.QBuffer
}

// WriteTo writes the CompetitionRankingScoreData to the given writable
func (crsd CompetitionRankingScoreData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	crsd.Unknown1.WriteTo(contentWritable)
	crsd.PID.WriteTo(contentWritable)
	crsd.Score.WriteTo(contentWritable)
	crsd.Modified.WriteTo(contentWritable)
	crsd.Unknown2.WriteTo(contentWritable)
	crsd.AppData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	crsd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the CompetitionRankingScoreData from the given readable
func (crsd CompetitionRankingScoreData) ExtractFrom(readable types.Readable) error {
	if err := crsd.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreData header. %s", err.Error())
	}

	if err := crsd.Unknown1.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreData.Unknown1. %s", err.Error())
	}

	if err := crsd.PID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreData.PID. %s", err.Error())
	}

	if err := crsd.Score.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreData.Score. %s", err.Error())
	}

	if err := crsd.Modified.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreData.Modified. %s", err.Error())
	}

	if err := crsd.Unknown2.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreData.Unknown2. %s", err.Error())
	}

	if err := crsd.AppData.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreData.AppData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CompetitionRankingScoreData
func (csrd CompetitionRankingScoreData) Copy() types.RVType {
	copied := NewCompetitionRankingScoreData()

	copied.StructureVersion = csrd.StructureVersion
	copied.Unknown1 = csrd.Unknown1.Copy().(types.UInt32)
	copied.PID = csrd.PID.Copy().(types.PID)
	copied.Score = csrd.Score.Copy().(types.UInt32)
	copied.Modified = csrd.Modified.Copy().(types.DateTime)
	copied.Unknown2 = csrd.Unknown2.Copy().(types.UInt8)
	copied.AppData = csrd.AppData.Copy().(types.QBuffer)

	return copied
}

// Equals checks if the given CompetitionRankingScoreData contains the same data as the current CompetitionRankingScoreData
func (csrd CompetitionRankingScoreData) Equals(o types.RVType) bool {
	if _, ok := o.(CompetitionRankingScoreData); !ok {
		return false
	}

	other := o.(CompetitionRankingScoreData)

	if csrd.StructureVersion != other.StructureVersion {
		return false
	}

	if !csrd.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !csrd.PID.Equals(other.PID) {
		return false
	}

	if !csrd.Score.Equals(other.Score) {
		return false
	}

	if !csrd.Modified.Equals(other.Modified) {
		return false
	}

	if !csrd.Unknown2.Equals(other.Unknown2) {
		return false
	}

	return csrd.AppData.Equals(other.AppData)
}

// CopyRef copies the current value of the CompetitionRankingScoreData
// and returns a pointer to the new copy
func (csrd CompetitionRankingScoreData) CopyRef() types.RVTypePtr {
	copied := csrd.Copy().(CompetitionRankingScoreData)
	return &copied
}

// Deref takes a pointer to the CompetitionRankingScoreData
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (csrd *CompetitionRankingScoreData) Deref() types.RVType {
	return *csrd
}

// String returns the string representation of the CompetitionRankingScoreData
func (csrd CompetitionRankingScoreData) String() string {
	return csrd.FormatToString(0)
}

// FormatToString pretty-prints the CompetitionRankingScoreData using the provided indentation level
func (csrd CompetitionRankingScoreData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CompetitionRankingScoreData{\n")
	fmt.Fprintf(&b, "%sUnknown1: %s,\n", indentationValues, csrd.Unknown1)
	fmt.Fprintf(&b, "%sPID: %s,\n", indentationValues, csrd.PID)
	fmt.Fprintf(&b, "%sScore: %s,\n", indentationValues, csrd.PID)
	fmt.Fprintf(&b, "%sModified: %s,\n", indentationValues, csrd.Modified)
	fmt.Fprintf(&b, "%ssUnknown2: %s,\n", indentationValues, csrd.Unknown2)
	fmt.Fprintf(&b, "%sAppData: %s,\n", indentationValues, csrd.AppData)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewCompetitionRankingScoreData returns a new CompetitionRankingScoreData
func NewCompetitionRankingScoreData() CompetitionRankingScoreData {
	return CompetitionRankingScoreData{
		Unknown1: types.NewUInt32(0),
		PID:      types.NewPID(0),
		Score:    types.NewUInt32(0),
		Modified: types.NewDateTime(0),
		Unknown2: types.NewUInt8(0),
		AppData:  types.NewQBuffer([]byte{}),
	}

}
