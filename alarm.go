package main

type Alarm struct {
	io     *PacketIO
	sender *SenderType
}

/****************************************************************
*
*  FUNCTION : This function call used by main（） .
*               初始化Alarm对象（包含初始化 IO & sender 对象）
*  DESCRIPTION :
*         syntax  -- NewAlarm() (*Alarm, error)；
*	input :
*
*  DESIGNER :
*
*****************************************************************/
func NewAlarm() (*Alarm, error) {
	var err error
	alarm := &Alarm{}
	alarm.io = RoutineInit()
	// error?
	alarm.sender, err = ReadSenderInfo()
	if err != nil {
		return nil, err
	}

	return alarm, nil
}
