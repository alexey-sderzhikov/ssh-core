package main

func mapConnectionDBtoHTTP(c connectionModelDB) connectionHTTP {
	return connectionHTTP{
		ID:       c.id,
		Address:  c.address,
		Username: c.username,
	}
}

func mapConnectionsListDBtoHTTP(cs []connectionModelDB) []connectionHTTP {
	result := make([]connectionHTTP, len(cs))

	for i, c := range cs {
		result[i] = mapConnectionDBtoHTTP(c)
	}

	return result
}
