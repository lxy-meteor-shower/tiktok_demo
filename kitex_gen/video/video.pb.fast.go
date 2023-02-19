// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package video

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 11:
		offset, err = x.fastReadField11(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.FollowCount = &tmp
	return offset, err
}

func (x *User) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.FollowerCount = &tmp
	return offset, err
}

func (x *User) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.IsFollow, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *User) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Avatar = &tmp
	return offset, err
}

func (x *User) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.BackgroundImage = &tmp
	return offset, err
}

func (x *User) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Signature = &tmp
	return offset, err
}

func (x *User) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.TotalFavorited = &tmp
	return offset, err
}

func (x *User) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.WorkCount = &tmp
	return offset, err
}

func (x *User) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.FavoriteCount = &tmp
	return offset, err
}

func (x *Video) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Video[number], err)
}

func (x *Video) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Author = &v
	return offset, nil
}

func (x *Video) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PlayUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.FavoriteCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.CommentCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.IsFavorite, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *Video) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseResp[number], err)
}

func (x *BaseResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMessage, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ServiceTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FeedRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FeedRequest[number], err)
}

func (x *FeedRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LatestTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FeedResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FeedResponse[number], err)
}

func (x *FeedResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.NextTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FeedResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Videos = append(x.Videos, &v)
	return offset, nil
}

func (x *FeedResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *VideoActionRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VideoActionRequest[number], err)
}

func (x *VideoActionRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Video = &v
	return offset, nil
}

func (x *VideoActionResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VideoActionResponse[number], err)
}

func (x *VideoActionResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *ListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListRequest[number], err)
}

func (x *ListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserID, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListResponse[number], err)
}

func (x *ListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Videos = append(x.Videos, &v)
	return offset, nil
}

func (x *ListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *MGetVideoRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MGetVideoRequest[number], err)
}

func (x *MGetVideoRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.VideoIDs = append(x.VideoIDs, v)
			return offset, err
		})
	return offset, err
}

func (x *MGetVideoResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MGetVideoResponse[number], err)
}

func (x *MGetVideoResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Videos = append(x.Videos, &v)
	return offset, nil
}

func (x *MGetVideoResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *User) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.Name)
	return offset
}

func (x *User) fastWriteField3(buf []byte) (offset int) {
	if x.FollowCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, *x.FollowCount)
	return offset
}

func (x *User) fastWriteField4(buf []byte) (offset int) {
	if x.FollowerCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, *x.FollowerCount)
	return offset
}

func (x *User) fastWriteField5(buf []byte) (offset int) {
	if !x.IsFollow {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 5, x.IsFollow)
	return offset
}

func (x *User) fastWriteField6(buf []byte) (offset int) {
	if x.Avatar == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, *x.Avatar)
	return offset
}

func (x *User) fastWriteField7(buf []byte) (offset int) {
	if x.BackgroundImage == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, *x.BackgroundImage)
	return offset
}

func (x *User) fastWriteField8(buf []byte) (offset int) {
	if x.Signature == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, *x.Signature)
	return offset
}

func (x *User) fastWriteField9(buf []byte) (offset int) {
	if x.TotalFavorited == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 9, *x.TotalFavorited)
	return offset
}

func (x *User) fastWriteField10(buf []byte) (offset int) {
	if x.WorkCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 10, *x.WorkCount)
	return offset
}

func (x *User) fastWriteField11(buf []byte) (offset int) {
	if x.FavoriteCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 11, *x.FavoriteCount)
	return offset
}

func (x *Video) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *Video) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *Video) fastWriteField2(buf []byte) (offset int) {
	if x.Author == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.Author)
	return offset
}

func (x *Video) fastWriteField3(buf []byte) (offset int) {
	if x.PlayUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.PlayUrl)
	return offset
}

func (x *Video) fastWriteField4(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.CoverUrl)
	return offset
}

func (x *Video) fastWriteField5(buf []byte) (offset int) {
	if x.FavoriteCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.FavoriteCount)
	return offset
}

func (x *Video) fastWriteField6(buf []byte) (offset int) {
	if x.CommentCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.CommentCount)
	return offset
}

func (x *Video) fastWriteField7(buf []byte) (offset int) {
	if !x.IsFavorite {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 7, x.IsFavorite)
	return offset
}

func (x *Video) fastWriteField8(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.Title)
	return offset
}

func (x *BaseResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *BaseResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *BaseResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMessage == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMessage)
	return offset
}

func (x *BaseResp) fastWriteField3(buf []byte) (offset int) {
	if x.ServiceTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.ServiceTime)
	return offset
}

func (x *FeedRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *FeedRequest) fastWriteField1(buf []byte) (offset int) {
	if x.LatestTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.LatestTime)
	return offset
}

func (x *FeedResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *FeedResponse) fastWriteField1(buf []byte) (offset int) {
	if x.NextTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.NextTime)
	return offset
}

func (x *FeedResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Videos == nil {
		return offset
	}
	for i := range x.Videos {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.Videos[i])
	}
	return offset
}

func (x *FeedResponse) fastWriteField3(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.BaseResp)
	return offset
}

func (x *VideoActionRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *VideoActionRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Video == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.Video)
	return offset
}

func (x *VideoActionResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *VideoActionResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *ListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserID == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserID)
	return offset
}

func (x *ListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Videos == nil {
		return offset
	}
	for i := range x.Videos {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.Videos[i])
	}
	return offset
}

func (x *ListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.BaseResp)
	return offset
}

func (x *MGetVideoRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *MGetVideoRequest) fastWriteField1(buf []byte) (offset int) {
	if len(x.VideoIDs) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.VideoIDs),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.VideoIDs[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *MGetVideoResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *MGetVideoResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Videos == nil {
		return offset
	}
	for i := range x.Videos {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.Videos[i])
	}
	return offset
}

func (x *MGetVideoResponse) fastWriteField2(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.BaseResp)
	return offset
}

func (x *User) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *User) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.Name)
	return n
}

func (x *User) sizeField3() (n int) {
	if x.FollowCount == nil {
		return n
	}
	n += fastpb.SizeInt64(3, *x.FollowCount)
	return n
}

func (x *User) sizeField4() (n int) {
	if x.FollowerCount == nil {
		return n
	}
	n += fastpb.SizeInt64(4, *x.FollowerCount)
	return n
}

func (x *User) sizeField5() (n int) {
	if !x.IsFollow {
		return n
	}
	n += fastpb.SizeBool(5, x.IsFollow)
	return n
}

func (x *User) sizeField6() (n int) {
	if x.Avatar == nil {
		return n
	}
	n += fastpb.SizeString(6, *x.Avatar)
	return n
}

func (x *User) sizeField7() (n int) {
	if x.BackgroundImage == nil {
		return n
	}
	n += fastpb.SizeString(7, *x.BackgroundImage)
	return n
}

func (x *User) sizeField8() (n int) {
	if x.Signature == nil {
		return n
	}
	n += fastpb.SizeString(8, *x.Signature)
	return n
}

func (x *User) sizeField9() (n int) {
	if x.TotalFavorited == nil {
		return n
	}
	n += fastpb.SizeInt64(9, *x.TotalFavorited)
	return n
}

func (x *User) sizeField10() (n int) {
	if x.WorkCount == nil {
		return n
	}
	n += fastpb.SizeInt64(10, *x.WorkCount)
	return n
}

func (x *User) sizeField11() (n int) {
	if x.FavoriteCount == nil {
		return n
	}
	n += fastpb.SizeInt64(11, *x.FavoriteCount)
	return n
}

func (x *Video) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	return n
}

func (x *Video) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *Video) sizeField2() (n int) {
	if x.Author == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.Author)
	return n
}

func (x *Video) sizeField3() (n int) {
	if x.PlayUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.PlayUrl)
	return n
}

func (x *Video) sizeField4() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.CoverUrl)
	return n
}

func (x *Video) sizeField5() (n int) {
	if x.FavoriteCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.FavoriteCount)
	return n
}

func (x *Video) sizeField6() (n int) {
	if x.CommentCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(6, x.CommentCount)
	return n
}

func (x *Video) sizeField7() (n int) {
	if !x.IsFavorite {
		return n
	}
	n += fastpb.SizeBool(7, x.IsFavorite)
	return n
}

func (x *Video) sizeField8() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(8, x.Title)
	return n
}

func (x *BaseResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *BaseResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.StatusCode)
	return n
}

func (x *BaseResp) sizeField2() (n int) {
	if x.StatusMessage == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMessage)
	return n
}

func (x *BaseResp) sizeField3() (n int) {
	if x.ServiceTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.ServiceTime)
	return n
}

func (x *FeedRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *FeedRequest) sizeField1() (n int) {
	if x.LatestTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.LatestTime)
	return n
}

func (x *FeedResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *FeedResponse) sizeField1() (n int) {
	if x.NextTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.NextTime)
	return n
}

func (x *FeedResponse) sizeField2() (n int) {
	if x.Videos == nil {
		return n
	}
	for i := range x.Videos {
		n += fastpb.SizeMessage(2, x.Videos[i])
	}
	return n
}

func (x *FeedResponse) sizeField3() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.BaseResp)
	return n
}

func (x *VideoActionRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *VideoActionRequest) sizeField1() (n int) {
	if x.Video == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.Video)
	return n
}

func (x *VideoActionResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *VideoActionResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *ListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ListRequest) sizeField1() (n int) {
	if x.UserID == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserID)
	return n
}

func (x *ListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ListResponse) sizeField1() (n int) {
	if x.Videos == nil {
		return n
	}
	for i := range x.Videos {
		n += fastpb.SizeMessage(1, x.Videos[i])
	}
	return n
}

func (x *ListResponse) sizeField2() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.BaseResp)
	return n
}

func (x *MGetVideoRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *MGetVideoRequest) sizeField1() (n int) {
	if len(x.VideoIDs) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.VideoIDs),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.VideoIDs[numIdxOrVal])
			return n
		})
	return n
}

func (x *MGetVideoResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *MGetVideoResponse) sizeField1() (n int) {
	if x.Videos == nil {
		return n
	}
	for i := range x.Videos {
		n += fastpb.SizeMessage(1, x.Videos[i])
	}
	return n
}

func (x *MGetVideoResponse) sizeField2() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.BaseResp)
	return n
}

var fieldIDToName_User = map[int32]string{
	1:  "Id",
	2:  "Name",
	3:  "FollowCount",
	4:  "FollowerCount",
	5:  "IsFollow",
	6:  "Avatar",
	7:  "BackgroundImage",
	8:  "Signature",
	9:  "TotalFavorited",
	10: "WorkCount",
	11: "FavoriteCount",
}

var fieldIDToName_Video = map[int32]string{
	1: "Id",
	2: "Author",
	3: "PlayUrl",
	4: "CoverUrl",
	5: "FavoriteCount",
	6: "CommentCount",
	7: "IsFavorite",
	8: "Title",
}

var fieldIDToName_BaseResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMessage",
	3: "ServiceTime",
}

var fieldIDToName_FeedRequest = map[int32]string{
	1: "LatestTime",
}

var fieldIDToName_FeedResponse = map[int32]string{
	1: "NextTime",
	2: "Videos",
	3: "BaseResp",
}

var fieldIDToName_VideoActionRequest = map[int32]string{
	1: "Video",
}

var fieldIDToName_VideoActionResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_ListRequest = map[int32]string{
	1: "UserID",
}

var fieldIDToName_ListResponse = map[int32]string{
	1: "Videos",
	2: "BaseResp",
}

var fieldIDToName_MGetVideoRequest = map[int32]string{
	1: "VideoIDs",
}

var fieldIDToName_MGetVideoResponse = map[int32]string{
	1: "Videos",
	2: "BaseResp",
}