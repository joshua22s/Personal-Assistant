package testalarmclock

import (
	"fmt"
	//	"os"
	//	"time"
	//	"github.com/stianeikeland/go-rpio"
	//	"gobot.io/x/gobot"
	//	"gobot.io/x/gobot/drivers/gpio"
	//	"gobot.io/x/gobot/platforms/raspi"
)

type TestAlarmClock struct {
	//	pin = rpio.Pin(10)

	//	r      = raspi.NewAdaptor()
	//	buzzer = gpio.NewBuzzerDriver(r, "7")
	//	robot  gobot.Robot
	//	done   = false
	//	pin    rpio.Pin
	//	r      raspi.Adaptor
	//	buzzer gpio.BuzzerDriver
	//	robot  gobot.Robot
	//	done   bool
}

func NewTestAlarmClock() *TestAlarmClock {
	a := TestAlarmClock{}
	//	a.pin = rpio.Pin(10)
	//	a.r = raspi.NewAdaptor()
	//	a.buzzer = gpio.NewBuzzerDriver(r, "7")
	return &a
}

func (this *TestAlarmClock) BuzzerWork() {
	//	type note struct {
	//		tone     float64
	//		duration float64
	//	}

	//	song := []note{
	//		{gpio.C4, gpio.Quarter},
	//		{gpio.C4, gpio.Quarter},
	//		{gpio.G4, gpio.Quarter},
	//		{gpio.G4, gpio.Quarter},
	//		{gpio.A4, gpio.Quarter},
	//		{gpio.A4, gpio.Quarter},
	//		{gpio.G4, gpio.Half},
	//		{gpio.F4, gpio.Quarter},
	//		{gpio.F4, gpio.Quarter},
	//		{gpio.E4, gpio.Quarter},
	//		{gpio.E4, gpio.Quarter},
	//		{gpio.D4, gpio.Quarter},
	//		{gpio.D4, gpio.Quarter},
	//		{gpio.C4, gpio.Half},
	//	}
	//	for _, val := range song {
	//		this.buzzer.Tone(val.tone, val.duration)
	//		time.Sleep(10 * time.Millisecond)
	//	}
}

func (this *TestAlarmClock) Activate() error {
	fmt.Println("run")
	return nil
	//	if err := this.rpio.Open(); err != nil {
	//		fmt.Println(err)
	//		os.Exit(1)
	//	}
	//	defer rpio.Close()
	//	this.pin.Output()
	//	this.pin.Toggle()
	//	time.Sleep(time.Second * 5)

	//	work := this.BuzzerWork

	//	this.robot = this.gobot.NewRobot("bot",
	//		[]this.gobot.Connection{r},
	//		[]this.gobot.Device{buzzer},
	//		work,
	//	)

	//	this.robot.Start()
	//	this.robot.Stop()
	//	this.pin.Toggle()
}
