## ALB Blue Green Deployment Tool

[![Build Status](https://travis-ci.org/gaw508/albbluegreen.svg?branch=master)](https://travis-ci.org/gaw508/albbluegreen)

A tool for managing blue green deployments using AWS ALB (ELBv2).

It relies on having a single load balancer, with 2 target groups,
one for blue, one for green. The default rule for the ALB
listener points at the blue target group, and there is a mutable
rule which points at the green target group, using a path-prefix
regex.

The rule can have one of two values `$a` or `*`. `$a` is a regex
which never matches, meaning the ALB listener points to the
default target group (blue). `*` is a regex which always passes,
meaning the ALB listener points to the green target group.

To toggle between blue and green, this tool changes the value of
the rule between these two regex.

## Usage

```
# build the tool
go get -u github.com/gaw508/albbluegreen/cmd/albbluegreen

# get current status
albbluegreen --rule-arn=[RULE_ARN] --aws-profile=[AWS_PROFILE] --aws-region=[AWS_REGION] get-status

# set status
albbluegreen --rule-arn=[RULE_ARN] --aws-profile=[AWS_PROFILE] --aws-region=[AWS_REGION] set-status --status=blue

# toggle status (blue -> green OR green -> blue)
albbluegreen --rule-arn=[RULE_ARN] --aws-profile=[AWS_PROFILE] --aws-region=[AWS_REGION] toggle
```
