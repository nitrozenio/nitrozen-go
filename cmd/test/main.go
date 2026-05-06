package main

import (
	"context"
	"fmt"
	"os"

	nitrozen "github.com/nitrozenio/nitrozen-go"
)

func main() {
	token := os.Getenv("NITROZEN_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "Set NITROZEN_TOKEN env var first:\n  NITROZEN_TOKEN=<your-token> go run ./cmd/test/")
		os.Exit(1)
	}

	cfg := nitrozen.NewConfiguration()
	client := nitrozen.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), nitrozen.ContextAccessToken, token)

	pass := 0
	fail := 0

	check := func(name string, err error) bool {
		if err != nil {
			fmt.Printf("  FAIL  %s: %v\n", name, err)
			fail++
			return false
		}
		fmt.Printf("  OK    %s\n", name)
		pass++
		return true
	}

	fmt.Println("\n── User ──────────────────────────────────────────")
	user, _, err := client.UsersAPI.UserGet(ctx).Execute()
	if check("UserGet", err) {
		plan := user.Data.GetPlan()
		fmt.Printf("        %s (%s)  plan=%s\n", user.Data.GetName(), user.Data.GetEmail(), plan.GetName())
	}

	usage, _, err := client.UsersAPI.UserUsageGet(ctx).Execute()
	if check("UserUsageGet", err) {
		limits := usage.Data.GetLimits()
		fmt.Printf("        projects=%d/%d  entries=%d/%d\n",
			usage.Data.GetProjectsCount(), limits.GetMaxProjects(),
			usage.Data.GetEntriesThisMonth(), limits.GetMaxEntriesPerMonth(),
		)
	}

	fmt.Println("\n── Tokens ────────────────────────────────────────")
	tokens, _, err := client.AuthenticationAPI.TokensGet(ctx).Execute()
	if check("TokensGet", err) {
		fmt.Printf("        %d token(s) found\n", len(tokens.Data))
	}

	fmt.Println("\n── Projects ──────────────────────────────────────")
	projects, _, err := client.ProjectsAPI.ProjectsGet(ctx).Execute()
	if !check("ProjectsGet", err) {
		summary(pass, fail)
		return
	}
	fmt.Printf("        %d project(s)\n", len(projects.Data))
	for _, p := range projects.Data {
		fmt.Printf("        [%d] %s  slug=%s\n", p.GetId(), p.GetName(), p.GetSlug())
	}

	if len(projects.Data) == 0 {
		fmt.Println("        (no projects to test entries with)")
		summary(pass, fail)
		return
	}

	first := projects.Data[0]
	projectId := first.GetId()
	fmt.Printf("\n── Entries for \"%s\" ───────────────────────────\n", first.GetName())

	entries, _, err := client.EntriesAPI.ProjectsProjectEntriesGet(ctx, projectId).PerPage(5).Execute()
	if check("EntriesGet", err) {
		fmt.Printf("        %d entry(ies) (page 1, up to 5)\n", len(entries.Data))
		for _, e := range entries.Data {
			status := "draft"
			if e.GetIsPublished() {
				status = "published"
			}
			fmt.Printf("        [%d] %s  (%s, %s)\n", e.GetId(), e.GetTitle(), e.GetCategory(), status)
		}
	}

	fmt.Println("\n── Public entries (no auth) ──────────────────────")
	pub, _, err := client.PublicAPI.PublicProjectsSlugEntriesGet(ctx, first.GetSlug()).Execute()
	if check("PublicEntriesGet", err) {
		fmt.Printf("        %d published entry(ies) via public API\n", len(pub.Data))
	}

	fmt.Println("\n── Get single project ────────────────────────────")
	proj, _, err := client.ProjectsAPI.ProjectsProjectGet(ctx, projectId).Execute()
	if check("ProjectGet", err) {
		fmt.Printf("        id=%d name=%s\n", proj.Data.GetId(), proj.Data.GetName())
	}

	summary(pass, fail)
}

func summary(pass, fail int) {
	fmt.Printf("\n%d passed, %d failed\n\n", pass, fail)
	if fail > 0 {
		os.Exit(1)
	}
}
