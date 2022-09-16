package sonyflake

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonFlakeUser  *sonyflake.Sonyflake
	sonFlakeMsg   *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func Init(machineID uint16) (err error) {
	sonyMachineID = machineID
	t, err := time.Parse("2006-01-02", "2022-06-20")
	settingUser := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonFlakeUser = sonyflake.NewSonyflake(settingUser)
	//
	t, err = time.Parse("2006-01-02", "2022-08-08")
	settingMsg := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonFlakeMsg = sonyflake.NewSonyflake(settingMsg)
	return
}

func GetUserID() (id uint64, err error) {
	if sonFlakeUser == nil {
		err = fmt.Errorf("sony flake not inited")
		return 0, err
	}
	id, err = sonFlakeUser.NextID()
	return
}

func GetMsgID() (id uint64, err error) {
	if sonFlakeMsg == nil {
		err = fmt.Errorf("sony flake not inited")
		return 0, err
	}
	id, err = sonFlakeMsg.NextID()
	return
}

func GetGroupID() (id uint64, err error) {
	if sonFlakeMsg == nil {
		err = fmt.Errorf("sony flake not inited")
		return 0, err
	}
	id, err = sonFlakeMsg.NextID()
	return
}
