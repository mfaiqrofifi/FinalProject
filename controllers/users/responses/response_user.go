package responses

import (
	"social_media/business/users/statusdesc"
	"social_media/business/users/userRegister"
	"social_media/business/users/userStatus"
	"time"
)

type Response struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Address   string `json:"alamat"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type deteilResponse struct {
// 	response response
// 	Status   []_help.DBStatus
// }
type comment struct {
	ID        int    `json:"id"`
	Coment    string `json:"comment"`
	UserID    int    `json:"userid"`
	Like      int    `json:"like"`
	Dislike   int    `json:"dislike"`
	StatusId  int    `json:"statusid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type responseDeteil struct{
// 	ID        int    `json:"id"`
// 	Nama      string `json:"nama"`
// 	Address   string `json:"alamat"`
// 	Status
// }
type loginResponse struct {
	ResponseLogin Response `json:"user"`
	Token         string   `json:"token"`
}

type responseStatus struct {
	ID        int    `json:"id"`
	Status    string `json:"status"`
	IdUser    int    `json:"iduser"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type responseDeteil struct {
// 	ID      int                `json:"id"`
// 	Nama    string             `json:"nama"`
// 	Address string             `json:"alamat"`
// 	Status  []users.UserStatus `json:"status"`
// }
func FromDomainComment(domain statusdesc.DetailStatus) comment {
	return comment{
		ID:        domain.ID,
		Coment:    domain.Coment,
		UserID:    domain.ID,
		Like:      domain.Like,
		StatusId:  domain.StatusId,
		Dislike:   domain.Dislike,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomain(domain userRegister.DomainRegister) Response {
	return Response{
		ID:        domain.ID,
		Nama:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func FromLoginResponse(domain userRegister.DomainRegister) loginResponse {
	user := Response{
		ID:        domain.ID,
		Nama:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
	return loginResponse{
		ResponseLogin: user,
		Token:         domain.Token,
	}

}

// func FromLoginDeteil(domain userRegister.DomainRegister) deteilResponse {
// 	user := response{
// 		ID:        domain.ID,
// 		Nama:      domain.Name,
// 		Email:     domain.Email,
// 		Address:   domain.Address,
// 		CreatedAt: domain.CreatedAt,
// 		UpdatedAt: domain.UpdatedAt,
// 	}
// 	return deteilResponse{
// 		response: user,
// 		Status:   domain.Status,
// 	}
// }

func FromDomainStatus(status userStatus.UserStatus) responseStatus {
	return responseStatus{
		ID:     status.ID,
		Status: status.Status,
		IdUser: status.IdUser,
	}
}

// func FromListDeteil(data []userRegister.DomainRegister) (result []deteilResponse) {
// 	result = []deteilResponse{}
// 	for _, user := range data {
// 		result = append(result, FromLoginDeteil(user))
// 	}
// 	return
// }

// func FromDomainDetiel(deteil users.UserDeteil)responseDeteil{
// 	return responseDeteil{
// 		ID: deteil.DomainRegister.ID,
// 		Nama: deteil.DomainRegister.Name,
// 		Address: deteil.DomainRegister.Address,
// 		Status: deteil.Status,
// 	}
// }
