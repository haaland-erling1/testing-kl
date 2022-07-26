package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2"
)

func WaitForFileChange(pathAddr string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case _, ok := <-watcher.Events:
				if !ok {
					return
				}
				// log.Printf("%s %s\n", event.Name, event.Op)
				return
			case _, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}

	}()

	err = watcher.Add(pathAddr)
	if err != nil {
		fmt.Println("Add failed:", err)
		// log.Fatal("Add failed:", err)
	}
	<-done
}

func remove(s []net.Listener, i int) []net.Listener {
	if len(s)-1 < i {
		return s
	}
	// fmt.Println(len(s))
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

type Service struct {
	Name   string `json:"name"`
	Port   int    `json:"proxyPort"`
	Target int    `json:"servicePort"`
}

func startProxy() {
	confFile := os.Getenv("CONFIG_FILE")

	var data struct {
		Services []Service `json:"services"`
	}

	listeners := []net.Listener{}

	for {

		configData, err := ioutil.ReadFile(confFile)
		if err != nil {
			fmt.Println("Error reading config file:", confFile, err)
		}

		err = json.Unmarshal(configData, &data)
		if err != nil {
			fmt.Println("Error unmarshalling config file:", err)
		}

		for len(listeners) > 0 {
			i := len(listeners) - 1
			listener := listeners[i]
			fmt.Println("closing: ", listener.Addr().String())
			listener.Close()
			listeners = remove(listeners, i)
		}

		// for i, listener := range listeners {
		// }

		// for _, conn := range connections {
		// 	if conn != nil {
		// 		fmt.Println("closing: ", conn.RemoteAddr().String())
		// 		conn.Close()

		// 	}
		// }

		// listeners = []net.Listener{}
		// connections = []net.Conn{}

		for _, service := range data.Services {
			go func(service Service) {

				fmt.Println(fmt.Sprint(service.Name, ":", service.Target, "->", service.Port))

				listener, err := net.Listen("tcp", fmt.Sprintf(":%d", service.Port))
				if err != nil {
					fmt.Println("Error listening:", err)
					return
				}

				defer listener.Close()

				listeners = append(listeners, listener)

				for {
					conn, err := listener.Accept()
					if err != nil {
						listener.Close()
						// fmt.Println("Error accepting connection: ", err)
						break
					}
					upconn, err := net.Dial("tcp", fmt.Sprint(service.Name, ":", service.Target))
					if err != nil {
						fmt.Println("Error dialing target: ", err)
						conn.Close()
						continue
					}
					go func() {
						io.Copy(conn, upconn)
						conn.Close()
						upconn.Close()
					}()
					go func() {
						io.Copy(upconn, conn)
						upconn.Close()
						conn.Close()
					}()
				}
			}(service)
		}

		WaitForFileChange(confFile)
		// time.Sleep(time.Second)

	}
}

func startApi() {

	confFile := os.Getenv("CONFIG_FILE")

	app := fiber.New()
	app.Post("/update", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/post", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Body()))

		os.WriteFile(confFile, c.Body(), fs.ModeAppend)

		// err := c.BodyParser(&data)
		// if err != nil {
		// 	return err
		// }

		// if err != nil {
		// 	return err
		// }

		c.Send([]byte("done"))
		return nil

	})

	app.Listen(":2999")
}

func main() {
	go startApi()
	go startProxy()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}
