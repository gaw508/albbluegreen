package aws

// This package implements a BlueGreenService using AWS ELBv2. It
// relies on having a single load balancer, with 2 target groups,
// one for blue, one for green. The default rule for the ALB
// listener points at the blue target group, and there is a mutable
// rule which points at the green target group.
//
// The rule can have one of two values "X" or "Y". "X" is a regex
// which never matches, meaning the ALB listener points to the
// default target group (blue). "Y" is a regex which always passes,
// meaning the ALB listener points to the green target group.
//
// To toggle between blue and green, this service changes the value
// of the rule between these two regex.

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/gaw508/albbluegreen"
)

const blueRegex = "$a" // Regex which always fails
const greenRegex = "*" // Regex which always passes

var _ albbluegreen.BlueGreenService = &BlueGreenService{}

type BlueGreenService struct {
	AlbClient *elbv2.ELBV2
	RuleArn   string
}

func (s *BlueGreenService) Status() (albbluegreen.BlueGreenStatus, error) {
	input := &elbv2.DescribeRulesInput{
		RuleArns: []*string{&s.RuleArn},
	}

	output, err := s.AlbClient.DescribeRules(input)
	if err != nil {
		return albbluegreen.Blue, err
	}

	if len(output.Rules) != 1 {
		return albbluegreen.Blue, fmt.Errorf("unexpected no. of rules: got '%d', expected '1'", len(output.Rules))
	}

	if output.Rules[0] == nil {
		return albbluegreen.Blue, fmt.Errorf("rule was nil")
	}

	if len(output.Rules[0].Conditions) != 1 {
		return albbluegreen.Blue, fmt.Errorf(
			"unexpected no. of conditions: got '%d', expected '1'",
			len(output.Rules[0].Conditions),
		)
	}

	if output.Rules[0].Conditions[0] == nil {
		return albbluegreen.Blue, fmt.Errorf("condition was nil")
	}

	if output.Rules[0].Conditions[0].Field == nil || *output.Rules[0].Conditions[0].Field != "path-pattern" {
		return albbluegreen.Blue, fmt.Errorf("unexpected condition type: got '%+v', expected 'path-pattern'", output.Rules[0].Conditions[0].Field)
	}

	if len(output.Rules[0].Conditions[0].Values) != 1 {
		return albbluegreen.Blue, fmt.Errorf(
			"unexpected no. of condition values: got '%d', expected '1'",
			len(output.Rules[0].Conditions[0].Values),
		)
	}

	if output.Rules[0].Conditions[0].Values[0] == nil {
		return albbluegreen.Blue, fmt.Errorf("condition value was nil")
	}

	switch *output.Rules[0].Conditions[0].Values[0] {
	case blueRegex:
		return albbluegreen.Blue, nil
	case greenRegex:
		return albbluegreen.Green, nil
	default:
		return albbluegreen.Blue, fmt.Errorf(
			"unexpected condition value: got '%s', expected '%s' or '%s'",
			*output.Rules[0].Conditions[0].Values[0],
			blueRegex,
			greenRegex,
		)
	}
}

func (s *BlueGreenService) SetStatus(status albbluegreen.BlueGreenStatus) error {
	var regex string
	switch status {
	case albbluegreen.Blue:
		regex = blueRegex
	case albbluegreen.Green:
		regex = greenRegex
	default:
		return fmt.Errorf("invalid status: '%s'. Expected 'blue' or 'green'", status)
	}

	input := &elbv2.ModifyRuleInput{
		RuleArn: &s.RuleArn,
		Conditions: []*elbv2.RuleCondition{{
			Field:  aws.String("path-pattern"),
			Values: []*string{&regex},
		}},
	}

	_, err := s.AlbClient.ModifyRule(input)
	if err != nil {
		return err
	}

	return nil
}

func (s *BlueGreenService) Toggle() (albbluegreen.BlueGreenStatus, error) {
	status, err := s.Status()
	if err != nil {
		return albbluegreen.Blue, err
	}

	newStatus := albbluegreen.InvertStatus(status)

	err = s.SetStatus(newStatus)
	if err != nil {
		return albbluegreen.Blue, err
	}

	return newStatus, nil
}
