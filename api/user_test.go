package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/devphasex/cedar-bank-api/db/mock"
	db "github.com/devphasex/cedar-bank-api/db/sqlc"
	"github.com/devphasex/cedar-bank-api/util"
	"github.com/devphasex/cedar-bank-api/util/hash"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
)

func randomUser(t *testing.T) (db.User, string) {
	ag := hash.DefaultArgonHash()
	p := util.RandomString(8)
	passwordHash, err := ag.GenerateHash([]byte(p), nil)

	require.NoError(t, err)
	passwordHashStr, passwordSaltStr := hash.ArgonStringEncode(passwordHash)
	return db.User{
		ID:             util.RandomInt(2000, 2500),
		Username:       util.RandomOwner(),
		Email:          util.RandomEmail(),
		Fullname:       util.RandomOwner(),
		HashedPassword: passwordHashStr,
		PasswordSalt:   passwordSaltStr,
	}, p
}

func TestCreateUserApi(t *testing.T) {
	user, p := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: map[string]any{
				"username": user.Username,
				"fullname": user.Fullname,
				"password": p,
				"email":    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},

		{
			name: "InternalServerError",
			body: map[string]any{
				"username": user.Username,
				"fullname": user.Fullname,
				"password": p,
				"email":    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).Return(db.User{}, pgx.ErrTxClosed)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},

		{
			name: "DuplicateUsername",
			body: map[string]any{
				"username": user.Username,
				"fullname": user.Fullname,
				"password": p,
				"email":    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).Return(db.User{}, &pgconn.PgError{
					Code:           "23505",
					ConstraintName: "users_username_key",
				})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusConflict, recorder.Code)
			},
		},

		{
			name: "DuplicateEmail",
			body: map[string]any{
				"username": user.Username,
				"fullname": user.Fullname,
				"password": p,
				"email":    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).Return(db.User{}, &pgconn.PgError{
					Code:           "23505",
					ConstraintName: "users_email_key",
				})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusConflict, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)

			tc.buildStubs(store)
			server := NewServer(store, nil)
			recorder := httptest.NewRecorder()
			url := "/auth/sign-up"

			b, _ := json.Marshal(tc.body)
			body := bytes.NewBuffer(b)
			request, err := http.NewRequest(http.MethodPost, url, body)

			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}

}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	var resp SuccessResponse
	err := json.NewDecoder(body).Decode(&resp)

	require.NoError(t, err)
	require.True(t, resp.Status)

	var gotUser CreateUserResponse
	dataJSON, err := json.Marshal(resp.Data)

	require.NoError(t, err)
	err = json.Unmarshal(dataJSON, &gotUser)
	require.NoError(t, err)
	require.Equal(t, CreateUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Fullname: user.Fullname,
	}, gotUser)
}