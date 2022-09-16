package mysql

//func SaveMsg(message *models.Message) (bool, error) {
//	sqlStr := `INSERT INTO t_message(F_MsgID,F_SendID,F_RecipientID,F_Content,F_SendTime) VALUES(?,?,?,?,?)`
//	_, err := DB.Exec(sqlStr, message.MsgID, message.SendID, message.RecipientID, message.Content, message.SendTime)
//	if err != nil {
//		logrus.Errorf("保存消息出错!,err:->%s", err)
//		return false, err
//	}
//	return true, nil
//}
//
//func SaveOfflineMsg(message *models.Message) error {
//	sqlStr := `INSERT INTO t_offline_message(F_MsgID,F_SendID,F_RecipientID,F_Content,F_SendTime) VALUES(?,?,?,?,?)`
//	_, err := DB.Exec(sqlStr, message.MsgID, message.SendID, message.RecipientID, message.Content, message.SendTime)
//	if err != nil {
//		logrus.Errorf("保存离线消息出错!,err:->%s", err)
//		return err
//	}
//	return nil
//}
//
//func OfflineMsgNumber(message *models.OfflineMsg) (number int, err error) {
//	sqlStr := `SELECT COUNT(F_ID) FROM t_offline_message WHERE F_SendID = ? AND F_RecipientID = ?`
//	err = DB.Select(&number, sqlStr, message.SendID, message.RecipientID)
//	if err != nil {
//		logrus.Errorf("查询离线消息数量出错!,err:->%s", err)
//		return 0, err
//	}
//	return number, nil
//}
//
//func OfflineMsg(message *models.OfflineMsg) (msgList []models.OfflineMessage, err error) {
//	sqlStr := `SELECT F_ID,F_MsgID,F_SendID,F_RecipientID,F_Content,F_SendTime FROM t_offline_message WHERE F_SendID = ? AND F_RecipientID = ?`
//	err = DB.Select(&msgList, sqlStr, message.SendID, message.RecipientID)
//	if err != nil {
//		logrus.Errorf("查询离线消息出错!,err:->%s", err)
//		return nil, err
//	}
//	return msgList, nil
//}
//
//func DeleteOfflineMsg(message *models.OfflineMsg) error {
//	sqlStr := `DELETE FROM t_offline_message WHERE F_SendID = ? AND F_RecipientID = ?`
//	_, err := DB.Exec(sqlStr, message.SendID, message.RecipientID)
//	if err != nil {
//		logrus.Errorf("删除离线消息出错,err:%s\n", err)
//		return err
//	}
//	return nil
//}
