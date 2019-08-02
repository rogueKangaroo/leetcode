package simple


/**
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

**/
type MinStack struct {
	Val int
	LastStack *MinStack
	NextStack *MinStack
	NextMin *MinStack
	LastMin *MinStack
	MinHeader *MinStack
	StackHeader *MinStack
	isSet bool
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	stack := &MinStack{
		Val:x,
		isSet:true,
	}
	if !this.isSet{
		this.isSet = true
		this.Val = x
		this.StackHeader = this
	} else {
		temp := this
		for {
			if temp.NextStack != nil {
				temp = temp.NextStack
			} else {
				temp.NextStack = stack
				stack.LastStack = temp
				break
			}
		}
		this.StackHeader = stack
	}
	minTemp := this.MinHeader
	if minTemp == nil {
		this.MinHeader = stack
		return
	}
	if minTemp.Val > x {
		stack.NextMin = minTemp
		minTemp.LastMin = stack
		this.MinHeader = stack
		return
	}
	for {
		if minTemp.Val < x && minTemp.NextMin != nil{
			minTemp = minTemp.NextMin
		} else {
			stack.NextMin = minTemp.NextMin
			minTemp.NextMin = stack
			stack.LastMin = minTemp
			if stack.NextMin != nil {
				stack.NextMin.LastMin = stack
			}
			break
		}
	}
}


func (this *MinStack) Pop()  {
	if this.StackHeader != this {
		if this.StackHeader.LastStack != nil {
			this.StackHeader.LastStack.NextStack = nil
		}
		if this.StackHeader.LastMin != nil {
			this.StackHeader.LastMin.NextMin = this.StackHeader.NextMin
		} else {
			this.MinHeader = this.StackHeader.NextMin
		}
		if this.StackHeader.NextMin != nil {
			this.StackHeader.NextMin.LastMin = this.StackHeader.LastMin
		}
		this.StackHeader = this.StackHeader.LastStack
	}else {
		this.StackHeader = nil
		this.MinHeader = nil
	}
}


func (this *MinStack) Top() int {
	return this.StackHeader.Val
}


func (this *MinStack) GetMin() int {
	return this.MinHeader.Val
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */