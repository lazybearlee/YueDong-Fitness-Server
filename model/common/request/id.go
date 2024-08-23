package request

// IDReq ID request struct
type IDReq struct {
	ID int `json:"id" form:"id" binding:"required"` // ID，必填
}

// Uint Convert ID to uint
func (r *IDReq) Uint() uint {
	return uint(r.ID)
}

// IDsReq ID request struct
type IDsReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"` // ID slice，必填
}

// AuthorityIdReq Authority ID request struct
type AuthorityIdReq struct {
	AuthorityId uint `json:"authorityId" form:"authorityId" binding:"required"` // Authority ID，必填
}
