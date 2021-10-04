package client

import (
// You have to import it yourself because it bugs
)

type PIF XenAPIObject

func (self *PIF) GetRecord() (record map[string]interface{}, err error) {
	record = make(map[string]interface{})
	result := APIResult{}
	err = self.Client.APICall(&result, "PIF.get_record", self.Ref)
	if err != nil {
		return record, err`
	}
	for k, v := range result.Value.(xmlrpc.Struct) {
		record[k] = v
	}
	return record, nil
}
