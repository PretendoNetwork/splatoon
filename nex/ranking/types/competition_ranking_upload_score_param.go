// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// CompetitionRankingUploadScoreParam is a type within the Ranking protocol
type CompetitionRankingUploadScoreParam struct {
	types.Structure
	Unknown1      types.UInt32
	SplatfestId   types.UInt32
	Unknown2      types.UInt32
	Score         types.UInt32
	TeamId        types.UInt8
	TeamScore     types.UInt32
	IsFirstUpload types.Bool
	AppData       types.QBuffer
}

// WriteTo writes the CompetitionRankingUploadScoreParam to the given writable
func (crusp CompetitionRankingUploadScoreParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	crusp.Unknown1.WriteTo(contentWritable)
	crusp.SplatfestId.WriteTo(contentWritable)
	crusp.Unknown2.WriteTo(contentWritable)
	crusp.Score.WriteTo(contentWritable)
	crusp.TeamId.WriteTo(contentWritable)
	crusp.TeamScore.WriteTo(contentWritable)
	crusp.IsFirstUpload.WriteTo(contentWritable)
	crusp.AppData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	crusp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the CompetitionRankingUploadScoreParam from the given readable
func (crusp CompetitionRankingUploadScoreParam) ExtractFrom(readable types.Readable) error {
	if err := crusp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingUploadScoreParam header. %s", err.Error())
	}

	if err := crusp.Unknown1.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingUploadScoreParam.Unknown1. %s", err.Error())
	}
	if err := crusp.SplatfestId.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingUploadScoreParam.SplatfestId. %s", err.Error())
	}
	if err := crusp.Unknown2.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingUploadScoreParam.Unknown2. %s", err.Error())
	}
	if err := crusp.Score.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingUploadScoreParam.Score. %s", err.Error())
	}
	if err := crusp.TeamId.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingUploadScoreParam.TeamId. %s", err.Error())
	}
	if err := crusp.TeamScore.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingUploadScoreParam.TeamWin. %s", err.Error())
	}
	if err := crusp.IsFirstUpload.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingUploadScoreParam.IsFirstUpload. %s", err.Error())
	}
	if err := crusp.AppData.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingUploadScoreParam.AppData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CompetitionRankingUploadScoreParam
func (crusp CompetitionRankingUploadScoreParam) Copy() types.RVType {
	copied := NewCompetitionRankingUploadScoreParam()

	copied.StructureVersion = crusp.StructureVersion
	copied.Unknown1 = crusp.Unknown1.Copy().(types.UInt32)
	copied.SplatfestId = crusp.SplatfestId.Copy().(types.UInt32)
	copied.Unknown2 = crusp.Unknown2.Copy().(types.UInt32)
	copied.Score = crusp.Score.Copy().(types.UInt32)
	copied.TeamId = crusp.TeamId.Copy().(types.UInt8)
	copied.TeamScore = crusp.TeamScore.Copy().(types.UInt32)
	copied.IsFirstUpload = crusp.IsFirstUpload.Copy().(types.Bool)
	copied.AppData = crusp.AppData.Copy().(types.QBuffer)

	return copied
}

// Equals checks if the given CompetitionRankingUploadScoreParam contains the same data as the current CompetitionRankingUploadScoreParam
func (crusp CompetitionRankingUploadScoreParam) Equals(o types.RVType) bool {
	if _, ok := o.(CompetitionRankingUploadScoreParam); !ok {
		return false
	}

	other := o.(CompetitionRankingUploadScoreParam)

	if crusp.StructureVersion != other.StructureVersion {
		return false
	}

	if !crusp.Unknown1.Equals(other.Unknown1) {
		return false
	}
	if !crusp.SplatfestId.Equals(other.SplatfestId) {
		return false
	}
	if !crusp.Unknown2.Equals(other.Unknown2) {
		return false
	}
	if !crusp.Score.Equals(other.Score) {
		return false
	}
	if !crusp.TeamId.Equals(other.TeamId) {
		return false
	}
	if !crusp.TeamScore.Equals(other.TeamScore) {
		return false
	}
	if !crusp.IsFirstUpload.Equals(other.IsFirstUpload) {
		return false
	}

	return crusp.AppData.Equals(other.AppData)
}

// CopyRef copies the current value of the CompetitionRankingUploadScoreParam
// and returns a pointer to the new copy
func (crusp CompetitionRankingUploadScoreParam) CopyRef() types.RVTypePtr {
	copied := crusp.Copy().(CompetitionRankingScoreInfo)
	return &copied
}

// Deref takes a pointer to the CompetitionRankingUploadScoreParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (crusp *CompetitionRankingUploadScoreParam) Deref() types.RVType {
	return *crusp
}

// String returns the string representation of the CompetitionRankingUploadScoreParam
func (crusp CompetitionRankingUploadScoreParam) String() string {
	return crusp.FormatToString(0)
}

// FormatToString pretty-prints the CompetitionRankingUploadScoreParam using the provided indentation level
func (crusp CompetitionRankingUploadScoreParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CompetitionRankingUploadScoreParam{\n")
	fmt.Fprintf(&b, "%sUnknown1: %s,\n", indentationValues, crusp.Unknown1)
	fmt.Fprintf(&b, "%sSplatfestId: %s,\n", indentationValues, crusp.SplatfestId)
	fmt.Fprintf(&b, "%sUnknown2: %s,\n", indentationValues, crusp.Unknown2)
	fmt.Fprintf(&b, "%sScore: %s,\n", indentationValues, crusp.Score)
	fmt.Fprintf(&b, "%sTeamId: %s,\n", indentationValues, crusp.TeamId)
	fmt.Fprintf(&b, "%TeamScore: %s,\n", indentationValues, crusp.TeamScore)
	fmt.Fprintf(&b, "%sIsFirstUpload: %s,\n", indentationValues, crusp.IsFirstUpload)
	fmt.Fprintf(&b, "%sAppData: %s,\n", indentationValues, crusp.AppData)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// CompetitionRankingUploadScoreParam returns a new CompetitionRankingUploadScoreParam
func NewCompetitionRankingUploadScoreParam() CompetitionRankingUploadScoreParam {
	return CompetitionRankingUploadScoreParam{
		Unknown1:      types.NewUInt32(0),
		SplatfestId:   types.NewUInt32(0),
		Unknown2:      types.NewUInt32(0),
		Score:         types.NewUInt32(0),
		TeamId:        types.NewUInt8(0),
		TeamScore:     types.NewUInt32(0),
		IsFirstUpload: types.NewBool(false),
		AppData:       types.NewQBuffer([]byte{}),
	}

}
