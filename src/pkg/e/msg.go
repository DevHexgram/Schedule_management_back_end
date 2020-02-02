package e

var MsgFlags = map[int]string {
	30200:"Token Expired",
	30210:"Token Format Changed",
	40000:"Wrong Format Of JSON",
	40010:"Wrong Format Of Header",
	40020:"Wrong Format of Token",
	40030:"Duplicate username",
	40040:"Poor Authority",
	40400:"Unable To Parse Parameters`",
	40410:"Username or Password Wrong",
	40420:"Invitation Code Wrong",
	40430:"Not Found",
	40440:"User Status Not Exist",
	50000:"Can't Insert Into Database",
	50010:"Can't Generate Token",
	50020:"Middleware Wrong",
}

func GetMsg(code int) string {
	msg,ok := MsgFlags[code]
	if ok {
		return msg
	}
	return "failed"
}
