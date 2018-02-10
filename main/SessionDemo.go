package main


/**
	session管理器底层数据结构
 */
type Provider interface {
	sessionInit(sid string)(Session,error)

}

/**
	session的接口设计
 */
type Session interface {


}

func main() {



}
