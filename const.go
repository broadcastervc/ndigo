// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 14 Jul 2022 12:03:24 BST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package ndigo

/*
#cgo linux LDFLAGS: -L/usr/local/lib -lndi
#cgo darwin LDFLAGS: -Wl,-rpath,/Library/NDI\ SDK\ for\ Apple/lib/macOS -L/Library/NDI\ SDK\ for\ Apple/lib/macOS -lndi
#cgo windows LDFLAGS: -L'C:\Program/ Files\NDI\NDI 5 Runtime\v5' -lProcessing.NDI.Lib.x64
#include <stdlib.h>
#include "include/Processing.NDI.Lib.h"
#include "cgo_helpers.h"
*/
import "C"

// FrameType as declared in include/Processing.NDI.structs.h:51
type FrameType int32

// FrameType enumeration from include/Processing.NDI.structs.h:51
const (
	FrameTypeNone         FrameType = iota
	FrameTypeVideo        FrameType = 1
	FrameTypeAudio        FrameType = 2
	FrameTypeMetadata     FrameType = 3
	FrameTypeError        FrameType = 4
	FrameTypeStatusChange FrameType = 100
	FrameTypeMax          FrameType = 2147483647
)

// FourCCVideoType as declared in include/Processing.NDI.structs.h:122
type FourCCVideoType int32

// FourCCVideoType enumeration from include/Processing.NDI.structs.h:122
const (
	FourCCVideoTypeUYVY FourCCVideoType = 1498831189
	FourCCTypeUYVY      FourCCVideoType = 1498831189
	FourCCVideoTypeUYVA FourCCVideoType = 1096178005
	FourCCTypeUYVA      FourCCVideoType = 1096178005
	FourCCVideoTypeP216 FourCCVideoType = 909193808
	FourCCTypeP216      FourCCVideoType = 909193808
	FourCCVideoTypePA16 FourCCVideoType = 909197648
	FourCCTypePA16      FourCCVideoType = 909197648
	FourCCVideoTypeYV12 FourCCVideoType = 842094169
	FourCCTypeYV12      FourCCVideoType = 842094169
	FourCCVideoTypeI420 FourCCVideoType = 808596553
	FourCCTypeI420      FourCCVideoType = 808596553
	FourCCVideoTypeNV12 FourCCVideoType = 842094158
	FourCCTypeNV12      FourCCVideoType = 842094158
	FourCCVideoTypeBGRA FourCCVideoType = 1095911234
	FourCCTypeBGRA      FourCCVideoType = 1095911234
	FourCCVideoTypeBGRX FourCCVideoType = 1481787202
	FourCCTypeBGRX      FourCCVideoType = 1481787202
	FourCCVideoTypeRGBA FourCCVideoType = 1094862674
	FourCCTypeRGBA      FourCCVideoType = 1094862674
	FourCCVideoTypeRGBX FourCCVideoType = 1480738642
	FourCCTypeRGBX      FourCCVideoType = 1480738642
	FourCCVideoTypeMax  FourCCVideoType = 2147483647
)

// FourCCAudioType as declared in include/Processing.NDI.structs.h:136
type FourCCAudioType int32

// FourCCAudioType enumeration from include/Processing.NDI.structs.h:136
const (
	FourCCAudioTypeFLTP FourCCAudioType = 1884572742
	FourCCTypeFLTP      FourCCAudioType = 1884572742
	FourCCAudioTypeMax  FourCCAudioType = 2147483647
)

// FrameFormatType as declared in include/Processing.NDI.structs.h:152
type FrameFormatType int32

// FrameFormatType enumeration from include/Processing.NDI.structs.h:152
const (
	FrameFormatTypeProgressive FrameFormatType = 1
	FrameFormatTypeInterleaved FrameFormatType = 0
	FrameFormatTypeField0      FrameFormatType = 2
	FrameFormatTypeField1      FrameFormatType = 3
	FrameFormatTypeMax         FrameFormatType = 2147483647
)

// RecvBandwidth as declared in include/Processing.NDI.Recv.h:43
type RecvBandwidth int32

// RecvBandwidth enumeration from include/Processing.NDI.Recv.h:43
const (
	RecvBandwidthMetadataOnly RecvBandwidth = -10
	RecvBandwidthAudioOnly    RecvBandwidth = 10
	RecvBandwidthLowest       RecvBandwidth = 0
	RecvBandwidthHighest      RecvBandwidth = 100
	RecvBandwidthMax          RecvBandwidth = 2147483647
)

// RecvColorFormat as declared in include/Processing.NDI.Recv.h:101
type RecvColorFormat int32

// RecvColorFormat enumeration from include/Processing.NDI.Recv.h:101
const (
	RecvColorFormatBGRXBGRA  RecvColorFormat = iota
	RecvColorFormatUYVYBGRA  RecvColorFormat = 1
	RecvColorFormatRGBXRGBA  RecvColorFormat = 2
	RecvColorFormatUYVYRGBA  RecvColorFormat = 3
	RecvColorFormatFastest   RecvColorFormat = 100
	RecvColorFormatBest      RecvColorFormat = 101
	RecvColorFormatEBGRXBGRA RecvColorFormat = 0
	RecvColorFormatEUYVYBGRA RecvColorFormat = 1
	RecvColorFormatERGBXRGBA RecvColorFormat = 2
	RecvColorFormatEUYVYRGBA RecvColorFormat = 3
	RecvColorFormatMax       RecvColorFormat = 2147483647
)
