package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"git.sr.ht/~u472892/dhcpdadm/templates"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func adminIndex(c *fiber.Ctx) error {
	p := &templates.AdminIndex{}
	c.Set("Content-Type", "text/html; charset=utf-8")
	templates.WritePageTemplate(c, p)
	return nil
}

func adminAgentsIndex(c *fiber.Ctx) error {
	p := &templates.AdminAgentsIndex{}
	c.Set("Content-Type", "text/html; charset=utf-8")
	templates.WritePageTemplate(c, p)
	return nil
}

func adminAgentsAdd(c *fiber.Ctx) error {
	p := &templates.AdminAgentsAdd{}
	d, err := validateAgentsAddData(c)
	if err != nil {
		p.ErrorMessage = fmt.Sprintf("%s", err)
	}
	p.Agent = d

	c.Set("Content-Type", "text/html; charset=utf-8")
	templates.WritePageTemplate(c, p)
	return nil
}

func adminAgentsAddNew(c *fiber.Ctx) error {
	p := &templates.AdminAgentsAddNew{}
	d, err := validateAgentsAddData(c)
	if err != nil {
		p.ErrorMessage = fmt.Sprintf("%s", err)
		c.Set("Content-Type", "text/html; charset=utf-8")
		templates.WritePageTemplate(c, p)
		return err
	}

	data := fetchAgentsInfo(d)
	d.Name = data.Agent
	d.Version = data.Version

	if d.Name == "" {
		p.ErrorMessage = "Unable to fetch information from agent. Check your input and try again."
	}

	p.Agent = d

	c.Set("Content-Type", "text/html; charset=utf-8")
	templates.WritePageTemplate(c, p)
	return nil
}

func validateAgentsAddData(c *fiber.Ctx) (templates.AdminAgentsAddAgent, error) {
	d := templates.AdminAgentsAddAgent{}

	d.Name = string(c.FormValue("agent_name"))
	d.Hostname = string(c.FormValue("agent_hostname"))
	d.Port = string(c.FormValue("agent_port"))
	d.Confirmed = string(c.FormValue("creation_confirm"))

	if d.Hostname == "" {
		return d, errors.New("Hostname cannot be empty. Check your input and try again.")
	}

	if d.Port == "" {
		d.Port = "9091"
	}

	return d, nil
}

func fetchAgentsInfo(agent templates.AdminAgentsAddAgent) AppConfig {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(fmt.Sprintf("http://%s:%s/api/v1/ping", agent.Hostname, agent.Port))
	resp := fasthttp.AcquireResponse()

	client := &fasthttp.Client{}
	client.Do(req, resp)

	b := resp.Body()
	remote := AppConfig{}
	json.Unmarshal(b, &remote)

	return remote
}
