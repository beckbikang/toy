package limit

import(
	"time"
	"sync.Mutex"
)


var(

	FullError = error.New("error:limit is full")

)


type Limiter interface{
	GetToken()(bool,error) //判断是否可以访问
	GetToken(n int)(int,error) //判断是否可以访问多个
	Reset()(bool,error)//重置
}


/**
1. 漏桶算法
漏桶(Leaky Bucket)算法思路很简单，水(请求)先进入到漏桶里，漏桶以一定的速度出水(接口有响应速率)，
当水流入速度过大会直接溢出(访问频率超过接口响应速率)，然后就拒绝请求，可以看出漏桶算法能强行限制数据的传输速率


漏通结构的构成
	1 桶有固定的容量
	2 固定的流速
	3 记录时间

**/



//方案一，加锁的版本
type leakyBucket struct{
	capacity int //容量
	rate float32 //速率
	currentNum int  //当前的数量
	currentTime int64 //毫秒数
	addNum int //每次添加的数量，默认是1个
	timeout time.Duration
	l sync.Mutex
}

//新建漏通对象
func NewLeakyBucket(rate float32, capacity int,  addNum int){
	return &leakyBucket{
		capacity:capacity,
		rate:rate,
		currentNum, 0,
		addNum:addNum,
		currentTime:0,
		l:sync.Mutex{},
	};
}

//获取访问权限
func (l*leakyBucket) GetToken()(bool,error){
	i,err := GetToken(1)
	if err != nil {
		return false, err
	}else{
		return true,nil
	}
}

/**


*/

//获取token
func (l*leakyBucket) GetToken (n int)(int,error){
	l.Lock()
	defer l.Unlock()
	if n > l.capacity {
		return 0, error.New("n is bigger than capacity")
	}
	now_time := time.Now().UnixNano()/1000;

	//首次进入
	if l.currentNum == 0{
		l.currentNum = now_time
	}

	//根据流速判断水位
	leakWater := l.rate *(now_time - l.currentTime)
	if leakWater < 0{
		leakWater = 0
	}

	//这个时间段内，本次已经流出去的水
	l.currentNum := l.currentNum - leakWater

	//流出大于进入则为0
	if l.currentNum < 0 {
		l.currentNum = 0
	}
	l.currentTime = now_time

	if l.currentNum + n <= l.capacity {
		l.currentNum += n
		return n,nil
	}else{
		return 0,error.New("")
	}

	return 0,nil
}

//清空数据
func (l*leakyBucket) Reset()(bool,error){
	l.currentNum = 0
	l.currentTime = time.Now().UnixNano()/1000;

	return true,nils
}


//方案二，利用chan和select做超时处理





/**
令牌桶算法
**/
type tokenBucket struct{

}






