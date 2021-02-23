package robot_service_rest

import (
	"errors"
	"github.com/IDzetI/Cable-robot/internal/robot"
	robot_service "github.com/IDzetI/Cable-robot/internal/robot/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type service struct {
	uc *robot.UseCase
	c  chan string
}

func New(uc *robot.UseCase) robot_service.Service {
	s := service{
		uc: uc,
		c:  make(chan string),
	}
	return &s
}

func (s *service) logger() {
	for msg := range s.c {
		log.Println(msg)
	}
}

func (s *service) Start() (err error) {
	go s.logger()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	g := r.Group("api")
	{
		g.POST("/exec", s.exec)
	}
	return r.Run(":8000")
}

func (s *service) exec(c *gin.Context) {
	var body struct {
		Command string    `json:"command"`
		Point   []float64 `json:"point"`
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	}
	switch body.Command {
	case "init":
		err = s.do(func() error {
			return s.uc.MoveInJoinSpace(body.Point, s.c)
		})
	case "control":
		err = s.do(func() error {
			return s.uc.ExternalControl(body.Point)
		})
	default:
		err = errors.New("invalid command")
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		})
	}
}

func (s *service) do(f func() error) (err error) {
	go func() {
		err = f()
	}()
	time.Sleep(time.Millisecond * 100)
	return
}
