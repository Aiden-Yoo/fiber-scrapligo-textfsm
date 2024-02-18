package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
)

func main2() {
	app := fiber.New()

	// POST 요청을 처리하는 라우트
	app.Post("/execute", func(c *fiber.Ctx) error {
		// 장비 정보 설정
		target := Target{
			Hostname: "192.168.1.93",
			Platform: "cisco_nxos",
			Username: "admin",
			Password: "Qwe123!@#",
		}

		response, err := executeCommand(&target)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(fiber.Map{"response": response})
	})

	log.Fatal(app.Listen(":3000"))
}

// Target 구조체는 네트워크 장치에 연결하기 위한 정보를 정의합니다.
type Target struct {
	Hostname string
	Platform string
	Username string
	Password string
}

// executeCommand는 주어진 타겟 장치에서 ScrapliGo를 사용하여 명령을 실행합니다.
func executeCommand(target *Target) (string, error) {
	p, err := platform.NewPlatform(
		target.Platform,
		target.Hostname,
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername(target.Username),
		options.WithAuthPassword(target.Password),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create platform: %w", err)
	}

	d, err := p.GetNetworkDriver()
	if err != nil {
		return "", fmt.Errorf("failed to get network driver: %w", err)
	}

	err = d.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open driver: %w", err)
	}
	defer d.Close()

	response, err := d.SendCommand("show version")
	if err != nil {
		return "", fmt.Errorf("failed to send command: %w", err)
	}

	return response.Result, nil
}
