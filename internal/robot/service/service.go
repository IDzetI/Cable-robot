package robot_service

type Service interface {
	Start() (err error)
	Stop() (err error)
}
