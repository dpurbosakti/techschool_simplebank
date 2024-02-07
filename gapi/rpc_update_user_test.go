package gapi

// import (
// 	"context"
// 	"database/sql"
// 	mockdb "simple-bank/db/mock"
// 	db "simple-bank/db/sqlc"
// 	"simple-bank/pb"
// 	"simple-bank/util"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// func TestUpdateUserAPI(t *testing.T) {
// 	user, _ := randomUser(t)

// 	newName := util.RandomOwner()
// 	newEmail := util.RandomEmail()

// 	testCase := []struct {
// 		name          string
// 		req           *pb.UpdateUserRequest
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(t *testing.T, res *pb.UpdateUserResponse, err error)
// 	}{
// 		{
// 			name: "OK",
// 			req: &pb.UpdateUserRequest{
// 				FullName: &newName,
// 				Email:    &newEmail,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				arg := db.UpdateUserParams{
// 					Username: user.Username,
// 					FullName: sql.NullString{
// 						String: newName,
// 						Valid:  true,
// 					},
// 					Email: sql.NullString{
// 						String: newEmail,
// 						Valid:  true,
// 					},
// 				}

// 				updatedUser := db.User{
// 					Username:          user.Username,
// 					HashedPassword:    user.HashedPassword,
// 					FullName:          newName,
// 					Email:             newEmail,
// 					PasswordChangedAt: user.PasswordChangedAt,
// 					CreatedAt:         user.CreatedAt,
// 					IsEmailVerified:   user.IsEmailVerified,
// 				}
// 				store.EXPECT().
// 					UpdateUser(gomock.Any(), gomock.Eq(arg)).
// 					Times(1).
// 					Return(updatedUser, nil)
// 			},
// 			checkResponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
// 				// check response
// 				require.NoError(t, err)
// 				require.NotNil(t, res)
// 				updatedUser := res.GetUser()
// 				require.Equal(t, user.Username, updatedUser.Username)
// 				require.Equal(t, newName, updatedUser.FullName)
// 				require.Equal(t, newEmail, updatedUser.Email)
// 			},
// 		},
// 		// {
// 		// 	name: "OK",
// 		// 	req: &pb.UpdateUserRequest{
// 		// 		Username: user.Username,
// 		// 		FullName: &newName,
// 		// 		Email:    &newEmail,
// 		// 	},
// 		// 	buildStubs: func(store *mockdb.MockStore) {
// 		// 		store.EXPECT().
// 		// 			CreateUserTx(gomock.Any(), gomock.Any()).
// 		// 			Times(1).
// 		// 			Return(db.CreateUserTxResult{}, nil)
// 		// 	},
// 		// 	checkResponse: func(t *testing.T, res *pb.CreateUserResponse, err error) {
// 		// 		// check response
// 		// 		require.Error(t, err)
// 		// 		st, ok := status.FromError(err)
// 		// 		require.True(t, ok)
// 		// 		require.Equal(t, codes.Internal, st.Code())
// 		// 	},
// 		// },
// 	}

// 	for i := range testCase {
// 		tc := testCase[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			storeCtrl := gomock.NewController(t)
// 			defer storeCtrl.Finish()

// 			store := mockdb.NewMockStore(storeCtrl)

// 			tc.buildStubs(store)

// 			// start test server and send request
// 			server := newTestServer(t, store, nil)
// 			res, err := server.UpdateUser(context.Background(), tc.req)
// 			tc.checkResponse(t, res, err)
// 		})
// 	}
// }
