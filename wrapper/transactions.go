package wrapper

/********** GLOBAL VARIABLES **********/
const _clientTransactionsURL = _url + "/trans"

/********** METHODS  **********/

// GetClientTransactions returns transactions made by all clients
func GetClientTransactions(cred ClientCredentials) Users {
	return handleRequestMulti(cred, "GET", _clientTransactionsURL, "trans", nil)
}

// GetUserTransactions returns transactions made by client users
// *CHECK* need OAuth key to make request
func GetUserTransactions(cred ClientCredentials, userID string) Users {
	url := _usersURL + "/" + userID + "/trans"

	return handleRequestMulti(cred, "GET", url, "trans", nil)
}
