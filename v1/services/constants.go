package services

const (
	membershipModule  = "/membership"
	membershipVersion = membershipModule + "/v1"
	membersEndpoint   = membershipVersion + "/members"
	memberEndpoint    = membersEndpoint + "/%s"
)
