//go:generate go install github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=templates
package main

import (
	"fmt"
	"log"

	"git.mills.io/prologic/bitcask"
	"git.sr.ht/~u472892/dhcpdadm/templates"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

var (
	APPNAME    = "dhcpdadm"
	MAJOR      = 0
	MINOR      = 1
	REVISION   = 221112
	CONFIG     AppConfig
	DATABASE   *bitcask.Bitcask
	cssHandler = fasthttp.FSHandler("./static/css", 1)
)

func version() string {
	return fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, REVISION)
}

func config() {
	CONFIG = loadAppConfigDefaults(CONFIG)

	var err error
	CONFIG, err = loadAppConfigFromFile(CONFIG, fmt.Sprintf("%s.conf", APPNAME))
	if err != nil {
		out(DebugLevelCritical, "main", fmt.Sprintf("loading configuration file '%s' failed: %s", fmt.Sprintf("%s.conf", APPNAME), err))
	}
}

func index(c *fiber.Ctx) error {
	p := &templates.Index{}
	c.Set("Content-Type", "text/html; charset=utf-8")
	templates.WritePageTemplate(c, p)
	return nil
}

func css(c *fiber.Ctx) error {
	s := &templates.BaseStyle{}
	c.Set("Content-Type", "text/css")
	templates.WriteStyleTemplate(c, s)
	return nil
}

func server() {
	app := fiber.New()

	app.Get("/api/v1/ping/", apiPing)
	app.Get("/api/v1/agent/", apiAgent)
	app.Get("/admin/agents/add/new/", adminAgentsAddNew)
	app.Get("/admin/agents/add/", adminAgentsAdd)
	app.Get("/admin/agents/", adminAgentsIndex)
	app.Get("/admin/", adminIndex)
	app.Get("/css/style.css", css)
	app.Get("/", index)

	out(DebugLevelInfo, "server", fmt.Sprintf("listening on %s:%d", CONFIG.Host, CONFIG.Port))
	log.Fatalln(app.Listen(fmt.Sprintf("%s:%d", CONFIG.Host, CONFIG.Port)))
}

func main() {
	config()

	out(DebugLevelInfo, "main", fmt.Sprintf("starting version %s", version()))
	out(DebugLevelInfo, "main", fmt.Sprintf("current debug level is %s", CONFIG.DebugLevel))
	out(DebugLevelInfo, "main", fmt.Sprintf("local agent is '%s'", CONFIG.Agent))

	var dbErr error
	DATABASE, dbErr = bitcask.Open(CONFIG.DatabasePath)
	if dbErr != nil {
		out(DebugLevelInfo, "main", fmt.Sprintf("unable to open database file '%s': %s", CONFIG.DatabasePath, dbErr))
	}
	defer DATABASE.Close()

	server()
}
