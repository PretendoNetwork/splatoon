package globals

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"strconv"
)

var AuthenticationServerAccount *nex.Account

var SecureServerAccount *nex.Account

func InitAccounts() {
	AuthenticationServerAccount = nex.NewAccount(types.NewPID(1), "Quazal Authentication", KerberosPassword)
	SecureServerAccount = nex.NewAccount(types.NewPID(2), "Quazal Rendez-Vous", KerberosPassword)
}

func AccountDetailsByPID(pid *types.PID) (*nex.Account, *nex.Error) {
	if pid.Equals(AuthenticationServerAccount.PID) {
		return AuthenticationServerAccount, nil
	}

	if pid.Equals(SecureServerAccount.PID) {
		return SecureServerAccount, nil
	}

	password, errorCode := PasswordFromPID(pid)
	if errorCode != 0 {
		return nil, nex.NewError(errorCode, "Failed to get password from PID")
	}

	account := nex.NewAccount(pid, strconv.Itoa(int(pid.LegacyValue())), password)

	return account, nil
}

func AccountDetailsByUsername(username string) (*nex.Account, *nex.Error) {
	if username == AuthenticationServerAccount.Username {
		return AuthenticationServerAccount, nil
	}

	if username == SecureServerAccount.Username {
		return SecureServerAccount, nil
	}

	pidInt, err := strconv.Atoi(username)
	if err != nil {
		Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.RendezVous.InvalidUsername, "Invalid username")
	}

	pid := types.NewPID(uint64(pidInt))

	password, errorCode := PasswordFromPID(pid)
	if errorCode != 0 {
		Logger.Errorf("Password err: %v", errorCode)
		return nil, nex.NewError(errorCode, "Failed to get password from PID")
	}

	account := nex.NewAccount(pid, username, password)

	return account, nil
}
