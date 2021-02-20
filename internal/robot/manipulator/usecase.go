package manipulator

type Manipulator struct {
	commands chan command
	joints []joint
}

type Joint struct {
	delta []float64
	currentAngle float64
}

type command struct {
	join int
	degree float64
}

func New(joints []Joint) Manipulator {
	return Manipulator{
		commands: make(chan command),
		joints: joints,
	}
}

func (m Manipulator)sendTrajectory(trajectory []command)  {
	for _, cmd := range trajectory{
		m.commands <- cmd
	}
}

func (m Manipulator)Move(point []float64)(err error)  {
	trajectory, err := m.getTrajectory(point)
	if err != nil{
		return
	}
	m.sendTrajectory(trajectory)
	return
}