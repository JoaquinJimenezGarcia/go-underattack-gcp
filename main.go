package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

var project string
var action string
var ruleName string = "under-attack-rule"

func main() {
	flag.StringVar(&project, "project", "project-12345", "Project to be used.")
	flag.StringVar(&action, "action", "list", "Action to perform: list, activate, deactivate.")

	flag.Parse()

	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	if action == "list" {
		ListFirewallRules(*computeService, ctx)
	} else if action == "activate" {
		AddFirewallRule(*computeService, ctx)
	} else if action == "deactivate" {
		RemoveFirewallRule(*computeService, ctx)
	} else {
		log.Fatalln("Please choose one action to perform (list, activate, deactivate).")
	}
}

func ListFirewallRules(computeService compute.Service, ctx context.Context) {
	req := computeService.Firewalls.List(project)
	if err := req.Pages(ctx, func(page *compute.FirewallList) error {
		for _, firewall := range page.Items {
			fmt.Printf("%#v\n", firewall.Description)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func AddFirewallRule(computeService compute.Service, ctx context.Context) {
	var DeniedRules []*compute.FirewallDenied
	DeniedRules = append(DeniedRules, &compute.FirewallDenied{
		IPProtocol: "all",
	})

	rb := &compute.Firewall{
		Name:        ruleName,
		Description: ruleName,
		Denied:      DeniedRules,
		Priority:    1,
		Direction:   "INGRESS",
	}

	_, err := computeService.Firewalls.Insert(project, rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Under Attack mode enabled.")
}

func RemoveFirewallRule(computeService compute.Service, ctx context.Context) {
	_, err := computeService.Firewalls.Delete(project, ruleName).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Under Attack mode disabled.")
}
