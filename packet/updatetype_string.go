// Code generated by "stringer -type=UpdateType"; DO NOT EDIT

package packet

import "fmt"

const _UpdateType_name = "UPDATETYPE_VALUESUPDATETYPE_MOVEMENTUPDATETYPE_CREATE_OBJECTUPDATETYPE_CREATE_OBJECT2UPDATETYPE_OUT_OF_RANGE_OBJECTSUPDATETYPE_NEAR_OBJECTS"

var _UpdateType_index = [...]uint8{0, 17, 36, 60, 85, 116, 139}

func (i UpdateType) String() string {
	if i >= UpdateType(len(_UpdateType_index)-1) {
		return fmt.Sprintf("UpdateType(%d)", i)
	}
	return _UpdateType_name[_UpdateType_index[i]:_UpdateType_index[i+1]]
}
