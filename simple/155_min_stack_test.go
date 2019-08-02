package simple

import (
	"fmt"
	"testing"
	"time"
)

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

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(fmt.Sprintf("TestMinStack getMin result is %v", minStack.GetMin()))
	time.Sleep(1 * time.Second)
	minStack.Pop()
	fmt.Println(fmt.Sprintf("TestMinStack top result is %v", minStack.Top()))
	time.Sleep(1 * time.Second)
	fmt.Println(fmt.Sprintf("TestMinStack getMin result is %v", minStack.GetMin()))
	time.Sleep(1 * time.Second)



	minStack = Constructor()
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)

	fmt.Println(fmt.Sprintf("TestMinStack top result is %v", minStack.Top()))
	time.Sleep(1 * time.Second)

	minStack.Pop()

	fmt.Println(fmt.Sprintf("TestMinStack getMin result is %v", minStack.GetMin()))
	time.Sleep(1 * time.Second)

	minStack.Pop()

	fmt.Println(fmt.Sprintf("TestMinStack getMin result is %v", minStack.GetMin()))
	time.Sleep(1 * time.Second)

	minStack.Pop()
	minStack.Push(2147483647)

	fmt.Println(fmt.Sprintf("TestMinStack top result is %v", minStack.Top()))
	time.Sleep(1 * time.Second)

	fmt.Println(fmt.Sprintf("TestMinStack getMin result is %v", minStack.GetMin()))
	time.Sleep(1 * time.Second)

	minStack.Push(-2147483648)

	fmt.Println(fmt.Sprintf("TestMinStack top result is %v", minStack.Top()))
	time.Sleep(1 * time.Second)

	fmt.Println(fmt.Sprintf("TestMinStack getMin result is %v", minStack.GetMin()))
	time.Sleep(1 * time.Second)

	minStack.Pop()

	fmt.Println(fmt.Sprintf("TestMinStack getMin result is %v", minStack.GetMin()))
	time.Sleep(1 * time.Second)

}
