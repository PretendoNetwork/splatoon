package nex

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	commonmatchmaking "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	commonmatchmakingext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	commonmatchmakeextension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	commonnattraversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	commonranking "github.com/PretendoNetwork/nex-protocols-common-go/v2/ranking"
	commonsecure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	matchmaking "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	matchmakingext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	matchmakeextension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	nattraversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/splatoon"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	"github.com/PretendoNetwork/splatoon/globals"
)

func CreateReportDBRecord(_ types.PID, _ types.UInt32, _ types.QBuffer) error {
	return nil
}

func stubGetPlayingSession(err error, packet nex.PacketInterface, callID uint32, _ types.List[types.PID]) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	lstPlayingSession := types.NewList[*match_making_types.PlayingSession]()

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	lstPlayingSession.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = matchmakeextension.ProtocolID
	rmcResponse.MethodID = matchmakeextension.MethodGetSimplePlayingSession
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

func cleanupMatchmakeSessionSearchCriteriasHandler(searchCriterias types.List[match_making_types.MatchmakeSessionSearchCriteria]) {
	for _, searchCriteria := range searchCriterias {
		searchCriteria.Attribs[4] = types.NewString("")
	}
}

func onAfterAutoMatchmakeWithParamPostpone(_ nex.PacketInterface, _ match_making_types.AutoMatchmakeParam) {
	globals.MatchmakingManager.Mutex.Lock()

	_, err := globals.MatchmakingManager.Database.Exec(`UPDATE matchmaking.matchmake_sessions SET open_participation=true WHERE game_mode=12`)
	if err != nil {
		globals.Logger.Error(err.Error())
	}

	globals.MatchmakingManager.Mutex.Unlock()
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
	commonMatchMakingProtocol := commonmatchmaking.NewCommonProtocol(matchMakingProtocol)
	commonMatchMakingProtocol.SetManager(globals.MatchmakingManager)

	matchMakingExtProtocol := matchmakingext.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol := commonmatchmakingext.NewCommonProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol.SetManager(globals.MatchmakingManager)

	matchmakeExtensionProtocol := matchmakeextension.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := commonmatchmakeextension.NewCommonProtocol(matchmakeExtensionProtocol)
	matchmakeExtensionProtocol.SetHandlerGetPlayingSession(stubGetPlayingSession)
	commonMatchmakeExtensionProtocol.CleanupMatchmakeSessionSearchCriterias = cleanupMatchmakeSessionSearchCriteriasHandler
	commonMatchmakeExtensionProtocol.OnAfterAutoMatchmakeWithParamPostpone = onAfterAutoMatchmakeWithParamPostpone
	commonMatchmakeExtensionProtocol.SetManager(globals.MatchmakingManager)

	rankingProtocol := ranking.NewProtocol(globals.SecureEndpoint)
	globals.SecureEndpoint.RegisterServiceProtocol(rankingProtocol)
	commonranking.NewCommonProtocol(rankingProtocol)
}
