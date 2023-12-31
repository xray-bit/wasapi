package com

import "golang.org/x/sys/windows"

//
// Device properties
// These PKEYs correspond to the old setupapi SPDRP_XXX properties
//
var (
	_PKEY_Device_DeviceDesc   = PROPERTYKEY{windows.GUID{Data1: 0xa45c254e, Data2: 0xdf1c, Data3: 0x4efd, Data4: [8]byte{0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0}}, 2}
	_PKEY_Device_FriendlyName = PROPERTYKEY{windows.GUID{Data1: 0xa45c254e, Data2: 0xdf1c, Data3: 0x4efd, Data4: [8]byte{0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0}}, 14}
	_PKEY_Device_InstanceId   = PROPERTYKEY{windows.GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 256}
	_PKEY_Device_ContainerId  = PROPERTYKEY{windows.GUID{Data1: 0x8c7ed206, Data2: 0x3f8a, Data3: 0x4827, Data4: [8]byte{0xb3, 0xab, 0xae, 0x9e, 0x1f, 0xae, 0xfc, 0x6c}}, 2}
)

func PKEY_Device_DeviceDesc() PROPERTYKEY {
	return _PKEY_Device_DeviceDesc
}

func PKEY_Device_FriendlyName() PROPERTYKEY {
	return _PKEY_Device_FriendlyName
}

func PKEY_Device_InstanceId() PROPERTYKEY {
	return _PKEY_Device_InstanceId
}

func PKEY_Device_ContainerId() PROPERTYKEY {
	return _PKEY_Device_ContainerId
}
