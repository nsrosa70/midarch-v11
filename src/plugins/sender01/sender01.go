package main

import (
	"framework/messages"
	"strconv"
	"time"
)

type Sender struct{}

func GetTypeElem() interface{}{
	return Sender{}
}

func GetBehaviourExp() string {
	//return libraries.BehaviourLibrary[calculatorinvoker.CalculatorInvoker{}]
	return "B = InvR.e1 -> I_PosInvP -> TerR.e1 -> B"
}


var idx1 = 0
var idx2 = 0


func (Sender) I_PreInvR1(msg *messages.SAMessage, r *bool) {
	time.Sleep(100 * time.Millisecond)
	*msg = messages.SAMessage{Payload:"[Plugin 01] Message 01 ["+strconv.Itoa(idx1)+"]"}
	idx1++

	*r = true
	return
}

func (Sender) I_PreInvR2(msg *messages.SAMessage, r *bool) {
	time.Sleep(100 * time.Millisecond)
	*msg = messages.SAMessage{Payload:"[Plugin 01] Message 02 ["+strconv.Itoa(idx2)+"]"}
	idx2++

	*r = true
	return
}