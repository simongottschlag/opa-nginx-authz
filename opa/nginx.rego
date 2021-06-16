package nginx

default authz = false

authz {
	is_valid_token
}

is_valid_token {
	some i
	authorization_header = input.headers.Authorization[i]
	token = split(authorization_header, " ")[1]
	token == "test"
}
