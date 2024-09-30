package nex

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
//	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	commonnattraversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	commonsecure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	nattraversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	"github.com/PretendoNetwork/splatoon/globals"
	"strconv"
	"strings"
	"fmt"

	commonmatchmaking "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	commonmatchmakingext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	commonmatchmakeextension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	commonranking "github.com/PretendoNetwork/nex-protocols-common-go/v2/ranking"
	matchmaking "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	matchmakingext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	matchmakingtypes "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	matchmakeextension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/splatoon"
)

func CreateReportDBRecord(_ *types.PID, _ *types.PrimitiveU32, _ *types.QBuffer) error {
	return nil
}

func stubGetPlayingSession(err error, packet nex.PacketInterface, callID uint32, _ *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	lstPlayingSession := types.NewList[*matchmakingtypes.PlayingSession]()

	// * There are no sessions, I tell you!
	//for _, playingSession := range playingSessions {
	//	lstPlayingSession.Append(playingSession)
	//}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	lstPlayingSession.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = matchmakeextension.ProtocolID
	rmcResponse.MethodID = matchmakeextension.MethodGetSimplePlayingSession
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

// from nex-protocols-common-go/matchmaking_utils.go
func compareSearchCriteria[T ~uint16 | ~uint32](original T, search string) bool {
	if search == "" { // * Accept any value
		return true
	}

	before, after, found := strings.Cut(search, ",")
	if found {
		min, err := strconv.ParseUint(before, 10, 64)
		if err != nil {
			return false
		}

		max, err := strconv.ParseUint(after, 10, 64)
		if err != nil {
			return false
		}

		return min <= uint64(original) && max >= uint64(original)
	} else {
		searchNum, err := strconv.ParseUint(before, 10, 64)
		if err != nil {
			return false
		}

		return searchNum == uint64(original)
	}
}

func cleanupMatchmakeSessionSearchCriteriasHandler(searchCriterias *types.List[*match_making_types.MatchmakeSessionSearchCriteria]) {
	for _, searchCriteria := range searchCriterias.Slice() {
	  searchCriteria.Attribs.SetIndex(1, "")
	  searchCriteria.Attribs.SetIndex(4, "")
	}
  }


func gameSpecificMatchmakeSessionSearchCriteriaChecksHandler(searchCriteria *matchmakingtypes.MatchmakeSessionSearchCriteria, matchmakeSession *matchmakingtypes.MatchmakeSession) bool {
	original := matchmakeSession.Attributes.Slice()
	search := searchCriteria.Attribs.Slice()
	if len(original) != len(search) {
		return false
	}

	for index, originalAttribute := range original {
		// ignore dummy criterias for matchmaking
		// everyone ends up in different rooms if you don't skip these
		if index == 1 || index == 4 {
			continue
		}
		searchAttribute := search[index]

		if !compareSearchCriteria(originalAttribute.Value, searchAttribute.Value) {
			return false
		}
	}

	return true
}

func onAfterAutoMatchmakeWithParamPostpone(_ nex.PacketInterface, _ *matchmakingtypes.AutoMatchmakeParam) {
	// * This is ugly but I can't work out a better way to do this
	// * Set Splatfest rooms to open participation
//	for _, session := range common_globals.Sessions {
//		if session.GameMatchmakeSession != nil && session.GameMatchmakeSession.GameMode.Value == 12 {
//			session.GameMatchmakeSession.OpenParticipation.Value = true
//		}
//	}
}

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	commonSecureProtocol := commonsecure.NewCommonProtocol(secureProtocol)

	commonSecureProtocol.CreateReportDBRecord = CreateReportDBRecord

	natTraversalProtocol := nattraversal.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	commonnattraversal.NewCommonProtocol(natTraversalProtocol)

	matchMakingProtocol := matchmaking.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	//commonmatchmaking.NewCommonProtocol(matchMakingProtocol)
	commonMatchMakingProtocol := commonmatchmaking.NewCommonProtocol(matchMakingProtocol)
	commonMatchMakingProtocol.SetManager(globals.MatchmakingManager)

	matchMakingExtProtocol := matchmakingext.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
//	commonmatchmakingext.NewCommonProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol := commonmatchmakingext.NewCommonProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol.SetManager(globals.MatchmakingManager)

	matchmakeExtensionProtocol := matchmakeextension.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := commonmatchmakeextension.NewCommonProtocol(matchmakeExtensionProtocol)
	matchmakeExtensionProtocol.SetHandlerGetPlayingSession(stubGetPlayingSession)
	commonMatchmakeExtensionProtocol.CleanupMatchmakeSessionSearchCriterias = cleanupMatchmakeSessionSearchCriteriasHandler
	fmt.Printf("Test")
	commonMatchmakeExtensionProtocol.OnAfterAutoMatchmakeWithParamPostpone = onAfterAutoMatchmakeWithParamPostpone
	commonMatchmakeExtensionProtocol.SetManager(globals.MatchmakingManager)

	rankingProtocol := ranking.NewProtocol(globals.SecureEndpoint)
	globals.SecureEndpoint.RegisterServiceProtocol(rankingProtocol)
	commonranking.NewCommonProtocol(rankingProtocol)
}
