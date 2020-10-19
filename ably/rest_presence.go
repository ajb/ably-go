package ably

type RESTPresence struct {
	client  *REST
	channel *RESTChannel
}

// Get gives the channel's presence messages according to the given parameters.
// The returned result can be inspected for the presence messages via
// the PresenceMessages() method.
func (p *RESTPresence) Get(params *PaginateParams) (*PaginatedResult, error) {
	path := p.channel.baseURL + "/presence"
	return newPaginatedResult(nil, paginatedRequest{typ: presMsgType, path: path, params: params, query: query(p.client.get), logger: p.logger(), respCheck: checkValidHTTPResponse})
}

// History gives the channel's presence messages history according to the given
// parameters. The returned result can be inspected for the presence messages
// via the PresenceMessages() method.
func (p *RESTPresence) History(params *PaginateParams) (*PaginatedResult, error) {
	path := p.channel.baseURL + "/presence/history"
	return newPaginatedResult(nil, paginatedRequest{typ: presMsgType, path: path, params: params, query: query(p.client.get), logger: p.logger(), respCheck: checkValidHTTPResponse})
}

func (p *RESTPresence) logger() *LoggerOptions {
	return p.client.logger()
}
