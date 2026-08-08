// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// CompetitionRankingScoreInfo is a type within the Ranking protocol
type CompetitionRankingScoreInfo struct {
	types.Structure
	FestID    types.UInt32
	ScoreData types.List[CompetitionRankingScoreData]
	Unknown   types.UInt32
	TeamWins  types.List[types.UInt32]
	TeamVotes types.List[types.UInt32]
}

// WriteTo writes the CompetitionRankingScoreInfo to the given writable
func (crsi CompetitionRankingScoreInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	crsi.FestID.WriteTo(contentWritable)
	crsi.ScoreData.WriteTo(contentWritable)
	crsi.Unknown.WriteTo(contentWritable)
	crsi.TeamWins.WriteTo(contentWritable)
	crsi.TeamVotes.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	crsi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the CompetitionRankingScoreInfo from the given readable
func (crsi CompetitionRankingScoreInfo) ExtractFrom(readable types.Readable) error {
	if err := crsi.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreInfo header. %s", err.Error())
	}

	if err := crsi.FestID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreInfo.FestID. %s", err.Error())
	}

	if err := crsi.ScoreData.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreInfo.ScoreData. %s", err.Error())
	}

	if err := crsi.Unknown.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreInfo.Unknown. %s", err.Error())
	}

	if err := crsi.TeamWins.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreInfo.TeamWins. %s", err.Error())
	}

	if err := crsi.TeamVotes.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract CompetitionRankingScoreInfo.TeamVotes. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CompetitionRankingScoreInfo
func (crsi CompetitionRankingScoreInfo) Copy() types.RVType {
	copied := NewCompetitionRankingScoreInfo()

	copied.StructureVersion = crsi.StructureVersion
	copied.FestID = crsi.FestID.Copy().(types.UInt32)
	copied.ScoreData = crsi.ScoreData.Copy().(types.List[CompetitionRankingScoreData])
	copied.Unknown = crsi.Unknown.Copy().(types.UInt32)
	copied.TeamWins = crsi.TeamWins.Copy().(types.List[types.UInt32])
	copied.TeamVotes = crsi.TeamVotes.Copy().(types.List[types.UInt32])

	return copied
}

// Equals checks if the given CompetitionRankingScoreInfo contains the same data as the current CompetitionRankingScoreInfo
func (crsi CompetitionRankingScoreInfo) Equals(o types.RVType) bool {
	if _, ok := o.(CompetitionRankingScoreInfo); !ok {
		return false
	}

	other := o.(CompetitionRankingScoreInfo)

	if crsi.StructureVersion != other.StructureVersion {
		return false
	}

	if !crsi.FestID.Equals(other.FestID) {
		return false
	}

	if !crsi.ScoreData.Equals(other.ScoreData) {
		return false
	}

	if !crsi.Unknown.Equals(other.Unknown) {
		return false
	}

	if !crsi.TeamWins.Equals(other.TeamWins) {
		return false
	}

	if !crsi.TeamVotes.Equals(other.TeamVotes) {
		return false
	}

	return crsi.TeamVotes.Equals(other.TeamVotes)
}

// CopyRef copies the current value of the CompetitionRankingScoreInfo
// and returns a pointer to the new copy
func (crsi CompetitionRankingScoreInfo) CopyRef() types.RVTypePtr {
	copied := crsi.Copy().(CompetitionRankingScoreInfo)
	return &copied
}

// Deref takes a pointer to the CompetitionRankingScoreInfo
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (crsi *CompetitionRankingScoreInfo) Deref() types.RVType {
	return *crsi
}

// String returns the string representation of the CompetitionRankingScoreInfo
func (crsi CompetitionRankingScoreInfo) String() string {
	return crsi.FormatToString(0)
}

// FormatToString pretty-prints the CompetitionRankingScoreInfo using the provided indentation level
func (crsi CompetitionRankingScoreInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CompetitionRankingScoreInfo{\n")
	fmt.Fprintf(&b, "%sFestID: %s,\n", indentationValues, crsi.FestID)
	fmt.Fprintf(&b, "%sScoreData: %s,\n", indentationValues, crsi.ScoreData)
	fmt.Fprintf(&b, "%sUnknown: %s,\n", indentationValues, crsi.Unknown)
	fmt.Fprintf(&b, "%sTeamWins: %s,\n", indentationValues, crsi.TeamWins)
	fmt.Fprintf(&b, "%sTeamVotes: %s,\n", indentationValues, crsi.TeamVotes)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewCompetitionRankingScoreInfo returns a new CompetitionRankingScoreInfo
func NewCompetitionRankingScoreInfo() CompetitionRankingScoreInfo {
	return CompetitionRankingScoreInfo{
		FestID:    types.NewUInt32(0),
		ScoreData: types.NewList[CompetitionRankingScoreData](),
		Unknown:   types.NewUInt32(0),
		TeamWins:  types.NewList[types.UInt32](),
		TeamVotes: types.NewList[types.UInt32](),
	}

}
