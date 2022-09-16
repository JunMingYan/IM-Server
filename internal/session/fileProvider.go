package session

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

type FileProvider struct {
	savePath string
	muxMap   map[string]*sync.RWMutex
}

func NewFileProvider(savePath string) *FileProvider {
	return &FileProvider{
		savePath: savePath,
		muxMap:   make(map[string]*sync.RWMutex),
	}
}

func (fp FileProvider) fileName(sessionID string) string {
	return fp.savePath + "/" + sessionID
}

func (fp FileProvider) toString(value interface{}) (string, error) {
	var str string
	vType := reflect.TypeOf(value)
	switch vType.Name() {
	case "int":
		i, _ := value.(int)
		str = strconv.Itoa(i)
	case "string":
		str, _ = value.(string)
	case "int64":
		i, _ := value.(int64)
		str = strconv.FormatInt(i, 10)
	default:
		return "", errors.New("Unsupported type: " + vType.Name())
	}
	return str, nil
}

func (fp FileProvider) write(sessionID string, data map[string]string, newFile bool) error {
	_, exist := fp.muxMap[sessionID]
	if !exist {
		fp.muxMap[sessionID] = new(sync.RWMutex)
	}
	fp.muxMap[sessionID].Lock()
	defer func() {
		fp.muxMap[sessionID].Unlock()
	}()
	fName := fp.fileName(sessionID)
	_, err := os.Stat(fName)
	var f *os.File
	if newFile {
		if err == nil {
			os.Remove(fName)
		}
		f, err = os.Create(fName)
		if err != nil {
			return errors.New("Creating session file failed: " + err.Error())
		}
	} else {
		if err != nil {
			return errors.New("Session file does not exists: " + err.Error())
		}
		f, err = os.OpenFile(fName, os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			return errors.New("Opening session file failed: " + err.Error())
		}
	}
	defer func() {
		os.Chtimes(fName, time.Now(), time.Now())
		f.Close()
	}()
	for key, value := range data {
		_, err = fmt.Fprintln(f, key+":"+value)
		if err != nil {
			return errors.New("Setting session key value failed: " + err.Error())
		}
	}
	return nil
}

func (fp FileProvider) read(sessionID string) (map[string]string, error) {
	fName := fp.fileName(sessionID)
	_, err := os.Stat(fName)
	if err != nil {
		return nil, errors.New("Session file does not exists: " + fName)
	}
	_, exist := fp.muxMap[sessionID]
	if !exist {
		fp.muxMap[sessionID] = new(sync.RWMutex)
	}
	fp.muxMap[sessionID].Lock()
	defer func() {
		fp.muxMap[sessionID].Unlock()
	}()
	f, err := os.Open(fName)
	if err != nil {
		return nil, errors.New("Opening session file failed: " + err.Error())
	}
	defer func() {
		os.Chtimes(fName, time.Now(), time.Now())
		f.Close()
	}()
	data := make(map[string]string)
	scaner := bufio.NewScanner(f)
	for scaner.Scan() {
		kv := strings.Split(scaner.Text(), ":")
		if len(kv) != 2 {
			continue
		}
		data[kv[0]] = kv[1]
	}
	if len(data) == 0 {
		return nil, errors.New("no data in session file")
	}
	return data, nil
}

func (fp FileProvider) create(sessionID string, data map[string]interface{}) error {
	strData := make(map[string]string)
	for key, value := range data {
		strValue, err := fp.toString(value)
		if err != nil {
			return err
		}
		strData[key] = strValue
	}
	return fp.write(sessionID, strData, true)
}

func (fp FileProvider) get(sessionID, key string) (string, error) {
	data, err := fp.read(sessionID)
	if err != nil {
		return "", err
	}
	value, ok := data[key]
	if !ok {
		return "", errors.New("Session key does not exists: " + key)
	}
	return value, err
}

func (fp FileProvider) getAll(sessionID string) (map[string]string, error) {
	return fp.read(sessionID)
}

func (fp FileProvider) set(sessionID, key string, value interface{}) error {
	data, err := fp.read(sessionID)
	if err != nil {
		return err
	}
	str, err := fp.toString(value)
	if err != nil {
		return err
	}
	data[key] = str
	return fp.write(sessionID, data, false)
}

func (fp FileProvider) destroy(sessionID string) error {
	fName := fp.fileName(sessionID)
	_, err := os.Stat(fName)
	if err != nil {
		return errors.New("Session file does not exists: " + fName)
	}
	_, exist := fp.muxMap[sessionID]
	if !exist {
		fp.muxMap[sessionID] = new(sync.RWMutex)
	}
	fp.muxMap[sessionID].Lock()
	err = os.Remove(fName)
	fp.muxMap[sessionID].Unlock()
	if err != nil {
		return errors.New("Removing session file failed: " + err.Error())
	}
	delete(fp.muxMap, sessionID)
	return nil
}

func (fp FileProvider) gc(expire int64) error {
	now := time.Now().Unix()
	for sessionID, mux := range fp.muxMap {
		fName := fp.fileName(sessionID)
		if len(fName) == 0 {
			continue
		}
		mux.Lock()
		info, err := os.Stat(fName)
		if err != nil {
			mux.Unlock()
			continue
		}
		modTime := info.ModTime().Unix()
		if modTime+expire*60 < now {
			err = os.Remove(fName)
			mux.Unlock()
			if err != nil {
				delete(fp.muxMap, sessionID)
			}
		} else {
			mux.Unlock()
		}
	}
	return nil
}
