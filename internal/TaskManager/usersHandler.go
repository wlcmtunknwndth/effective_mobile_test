package TaskManager

import (
	"context"
	"encoding/json"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"github.com/wlcmtunknwndth/effective_mobile_test/lib/httpResp"
	"github.com/wlcmtunknwndth/effective_mobile_test/lib/sl"
	"io"
	"net/http"
)

const (
	statusUserCreated         = "User created"          // 201
	statusBadRequest          = "Bad request"           // 400
	statusInternalServerError = "Internal server error" // 500
)

func (s *Service) CreateUser(w http.ResponseWriter, r *http.Request) {
	const op = scope + "GetUserByPassport"

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			s.log.Error(errCloseBody, sl.Op(op), sl.Err(err))
			return
		}
		return
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.log.Error(errReadBody, sl.Op(op), sl.Err(err))
		httpResp.WriteResponse(w, http.StatusBadRequest, statusBadRequest)
		return
	}

	var userApi models.CreateUserAPI
	if err = json.Unmarshal(body, &userApi); err != nil {
		s.log.Error(errCantUnmarshal, sl.Op(op), sl.Err(err))

		return
	}

	usr, usrInfo, err := models.CreateUserToUsersDB(&userApi)
	if err != nil {
		s.log.Error("couldn't parse request", sl.Op(op), sl.Err(err))
		httpResp.WriteResponse(w, http.StatusBadRequest, statusBadRequest)
		return
	}

	var ctx context.Context

	id, err := s.users.CreateUser(ctx, usr)
	if err != nil {
		s.log.Error("couldn't create user", sl.Op(op), sl.Err(err))
		httpResp.WriteResponse(w, http.StatusInternalServerError, statusInternalServerError)
		return
	}

	usrInfo.UserID = id
	if err = s.users.AddUserInfo(ctx, usrInfo); err != nil {
		s.log.Error("couldn't save user info", sl.Op(op), sl.Err(err))
		httpResp.WriteResponse(w, http.StatusInternalServerError, statusInternalServerError)
		return
	}

	httpResp.WriteResponse(w, http.StatusCreated, statusUserCreated)
	return
}
