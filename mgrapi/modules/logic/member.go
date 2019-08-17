package logic

// //IMember 用户登录
// type IMemberLogic interface {
// 	QueryUserInfo(u string, ident string) (info db.QueryRow, err error) //在用
// 	Query(uid int64) (db.QueryRow, error)
// 	CacheQuery(u string, ident string) (ls *model.MemberState, err error)
// 	QueryAuth(sysID, userID int64) (err error)
// 	QueryRoleByNameAndIdent(name, password, ident string) (s *model.MemberState, err error)
// 	SaveLoginStateToCache(s *model.MemberState) error
// }

// //MemberLogic 用户登录管理
// type MemberLogic struct {
// 	c     component.IContainer
// 	cache member.ICacheMember
// 	db    member.IDBMember
// 	http  *http.Client
// }

// //NewMemberLogic 创建登录对象
// func NewMemberLogic(c component.IContainer) *MemberLogic {
// 	return &MemberLogic{
// 		c:     c,
// 		cache: member.NewCacheMember(c),
// 		db:    member.NewDBMember(c),
// 		http:  &http.Client{},
// 	}
// }

// //Query 查询用户信息
// func (m *MemberLogic) Query(uid int64) (db.QueryRow, error) {
// 	return m.db.QueryByID(uid)
// }

// //QueryRoleByNameAndIdent xx
// func (m *MemberLogic) QueryRoleByNameAndIdent(name, password, ident string) (s *model.MemberState, err error) {
// 	return m.db.Query(name, password, ident)
// }

// // QueryAuth xx
// func (m *MemberLogic) QueryAuth(sysID, userID int64) (err error) {
// 	err = m.cache.QueryAuth(sysID, userID)
// 	if err != nil {
// 		data, err := m.db.QueryAuth(sysID, userID)
// 		if err != nil || data == nil {
// 			return err
// 		}
// 		if err = m.cache.SaveAuth(sysID, userID, data); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// // QueryUserInfo 返回用户信息
// func (m *MemberLogic) QueryUserInfo(u string, ident string) (ls db.QueryRow, err error) {

// 	if ls, err = m.db.QueryByUserName(u, ident); err != nil {
// 		return nil, err
// 	}
// 	//检查用户是否已锁定
// 	if ls.GetInt("status") == enum.UserLock {
// 		return nil, context.NewError(context.ERR_LOCKED, "用户被锁定暂时无法登录")
// 	}
// 	//检查用户是否已禁用
// 	if ls.GetInt("status") == enum.UserDisable {
// 		return nil, context.NewError(context.ERR_FORBIDDEN, "用户被禁用请联系管理员")
// 	}
// 	return ls, err
// }
