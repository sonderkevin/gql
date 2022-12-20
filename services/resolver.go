package services

// import (
// 	"context"

// 	"github.com/nrfta/go-paging"
// 	"github.com/volatiletech/sqlboiler/v4/queries/qm"
// )

// //go:generate go run github.com/99designs/gqlgen generate

// type Resolver struct{}

// type queryResolver struct {
// 	*Resolver
// }

// func (r *Resolver) Query() QueryResolver {
// 	return &queryResolver{r}
// }

// func (r *Resolver) PageInfo() PageInfoResolver {
// 	return paging.NewPageInfoResolver()
// }

// func (r *queryResolver) Users(ctx context.Context, page *paging.PageArgs) (*UserFriendsConnection, error) {
// 	var mods []qm.QueryMod

// 	totalCount := 3
// 	// if err != nil {
// 	// 	return &UserFriendsConnection{
// 	// 		PageInfo: paging.NewEmptyPageInfo(),
// 	// 	}, err
// 	// }

// 	paginator := paging.NewOffsetPaginator(page, int64(totalCount))
// 	mods = append(mods, paginator.QueryMods()...)

// 	records := users
// 	// if err != nil {
// 	// 	return &PostConnection{
// 	// 		PageInfo: paging.NewEmptyPageInfo(),
// 	// 	}, err
// 	// }

// 	result := &UserFriendsConnection{
// 		PageInfo: &paginator.PageInfo,
// 	}

// 	for _, row := range records {

// 		result.Edges = append(result.Edges, &UserFriendsEdge{
// 			// Cursor: paging.EncodeOffsetCursor(paginator.Offset + i + 1),
// 			Cursor: "xxx",
// 			Node:   row,
// 		})
// 	}
// 	return result, nil
// }

// // func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
// // 	var result []*User

// // 	for _, userData := range users {
// // 		user := &User{
// // 			ID:   userData.ID,
// // 			Name: userData.Name,
// // 			FriendsConnection: &UserFriendsConnection{
// // 				Edges: []*UserFriendsEdge{},
// // 			},
// // 		}

// // 		// result = append(result)

// // 		if userData.UserFriends != nil {
// // 			// var userFriends []*User

// // 			for _, friendId := range userData.UserFriends {
// // 				userIndex, _ := strconv.ParseInt(friendId, 10, 0)
// // 				friendData := users[userIndex-1]

// // 				userFriend := &User{
// // 					friendData.ID,
// // 					friendData.Name,
// // 					&UserFriendsConnection{},
// // 				}

// // 				// userFriends = append(userFriends, userFriend)
// // 				user.FriendsConnection.Edges = append(user.FriendsConnection.Edges, &UserFriendsEdge{
// // 					Cursor: "0",
// // 					Node:   userFriend,
// // 				})
// // 			}
// // 		}

// // 	}

// // 	return result, nil
// // }
