package parser

import (
	"strings"
)

type Parser struct {
	PlanPrefix     string
	ErrorPrefix    string
	ResourcePrefix string
}

type PlanResources []string

type PlanResult struct {
	Status  int
	Body    string
	Plan    PlanResources
	Summary string
	Error   string
}

func NewParser() *Parser {
	return &Parser{
		PlanPrefix:     "Plan:",
		ErrorPrefix:    "Error:",
		ResourcePrefix: "#",
	}
}

func (p *Parser) Do(body string) PlanResult {
	var (
		planSummary string
		planRes     PlanResources
		errRes      string
		status      int
	)

	begin := 0
	lines := strings.Split(body, "\n")
	for i, line := range lines {
		line = strings.TrimLeft(line, " ")
		if strings.HasPrefix(line, p.PlanPrefix) {
			planSummary = line
			status = 0
		} else if strings.HasPrefix(line, p.ErrorPrefix) {
			errRes = strings.Join(lines[i:], "\n")
			status = 1
			break
		} else if strings.HasPrefix(line, p.ResourcePrefix) {
			planRes = append(planRes, strings.Join(lines[begin:i], "\n"))
			begin = i
		}
	}

	return PlanResult{
		Status:  status,
		Body:    body,
		Plan:    planRes,
		Summary: planSummary,
		Error:   errRes,
	}
}
