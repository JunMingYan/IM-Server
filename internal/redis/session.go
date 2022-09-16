package redis

const sessionKey = "sessionID:"

//func SaveSession(sessionID string, session sessions.Session) error {
//	value, err := json.Marshal(session)
//	if err != nil {
//		return err
//	}
//	err := db.HMSet(sessionKey+sessionID, value).Err()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func GetSession(sessionID string) (interface{}, error) {
//	session, err := db.MGet(sessionKey + sessionID).Result()
//	if err != nil {
//		return nil, err
//	}
//
//	return session, err
//}
//
//func RemoveSession(sessionID string) error {
//	err := db.Del(sessionKey + sessionID).Err()
//	if err != nil {
//		return err
//	}
//	return nil
//}
