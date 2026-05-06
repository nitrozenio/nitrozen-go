# Nitrozen Go SDK

Official Go client for the [Nitrozen.io](https://nitrozen.io) changelog API. Manage projects, entries, and API tokens programmatically.

## Installation

```sh
go get github.com/nitrozenio/nitrozen-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    nitrozen "github.com/nitrozenio/nitrozen-go"
)

func main() {
    cfg := nitrozen.NewConfiguration()
    client := nitrozen.NewAPIClient(cfg)

    // Authenticate using context
    ctx := context.WithValue(context.Background(), nitrozen.ContextAccessToken, "YOUR_API_TOKEN")

    // List projects
    projects, _, err := client.ProjectsAPI.ProjectsGet(ctx).Execute()
    if err != nil {
        log.Fatal(err)
    }
    for _, p := range projects.Data {
        fmt.Printf("[%d] %s  slug=%s\n", p.GetId(), p.GetName(), p.GetSlug())
    }

    // Create a changelog entry
    entry, _, err := client.EntriesAPI.
        ProjectsProjectEntriesPost(ctx, projects.Data[0].GetId()).
        EntryInput(nitrozen.EntryInput{
            Title:    nitrozen.PtrString("Dark mode shipped"),
            Content:  nitrozen.PtrString("Users can now toggle dark mode from their profile settings."),
            Category: nitrozen.PtrString("new"),
        }).
        Execute()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created entry: %d\n", entry.Data.GetId())
}
```

## Authentication

All authenticated endpoints require a Bearer token. Create one from your [API Tokens](https://nitrozen.io/api-tokens) page and pass it via context:

```go
ctx := context.WithValue(context.Background(), nitrozen.ContextAccessToken, "YOUR_API_TOKEN")
```

## Rate Limiting

Rate limits are determined by your plan (`api_rate_limit` requests per minute). The API returns `429 Too Many Requests` when exceeded. Unauthenticated requests are limited to 60 req/min per IP.

## Pagination

List endpoints accept `page` and `per_page` query parameters. Responses include a `meta` object with pagination details:

```go
resp, _, err := client.EntriesAPI.
    ProjectsProjectEntriesGet(ctx, projectId).
    Page(2).
    PerPage(20).
    Execute()

fmt.Println(resp.Meta.GetTotal())
```

## API Endpoints

All URIs are relative to `https://nitrozen.io/api/v1`.

| API | Method | HTTP | Description |
|-----|--------|------|-------------|
| *AuthenticationAPI* | [**TokensGet**](docs/AuthenticationAPI.md#tokensget) | GET /tokens | List API tokens |
| *AuthenticationAPI* | [**TokensPost**](docs/AuthenticationAPI.md#tokenspost) | POST /tokens | Create an API token |
| *AuthenticationAPI* | [**TokensTokenDelete**](docs/AuthenticationAPI.md#tokenstokendelete) | DELETE /tokens/{token} | Revoke an API token |
| *EntriesAPI* | [**ProjectsProjectEntriesGet**](docs/EntriesAPI.md#projectsprojectentriesget) | GET /projects/{project}/entries | List entries |
| *EntriesAPI* | [**ProjectsProjectEntriesPost**](docs/EntriesAPI.md#projectsprojectentriespost) | POST /projects/{project}/entries | Create an entry |
| *EntriesAPI* | [**ProjectsProjectEntriesEntryGet**](docs/EntriesAPI.md#projectsprojectentriesentryget) | GET /projects/{project}/entries/{entry} | Get an entry |
| *EntriesAPI* | [**ProjectsProjectEntriesEntryPut**](docs/EntriesAPI.md#projectsprojectentriesentryput) | PUT /projects/{project}/entries/{entry} | Update an entry |
| *EntriesAPI* | [**ProjectsProjectEntriesEntryDelete**](docs/EntriesAPI.md#projectsprojectentriesentrydelete) | DELETE /projects/{project}/entries/{entry} | Delete an entry |
| *ProjectsAPI* | [**ProjectsGet**](docs/ProjectsAPI.md#projectsget) | GET /projects | List projects |
| *ProjectsAPI* | [**ProjectsPost**](docs/ProjectsAPI.md#projectspost) | POST /projects | Create a project |
| *ProjectsAPI* | [**ProjectsProjectGet**](docs/ProjectsAPI.md#projectsprojectget) | GET /projects/{project} | Get a project |
| *ProjectsAPI* | [**ProjectsProjectPut**](docs/ProjectsAPI.md#projectsprojectput) | PUT /projects/{project} | Update a project |
| *ProjectsAPI* | [**ProjectsProjectDelete**](docs/ProjectsAPI.md#projectsprojectdelete) | DELETE /projects/{project} | Delete a project |
| *PublicAPI* | [**PublicProjectsSlugEntriesGet**](docs/PublicAPI.md#publicprojectsslugentriesget) | GET /public/projects/{slug}/entries | List published entries (no auth) |
| *UsersAPI* | [**UserGet**](docs/UsersAPI.md#userget) | GET /user | Get current user |
| *UsersAPI* | [**UserUsageGet**](docs/UsersAPI.md#userusageget) | GET /user/usage | Get usage statistics |

## Models

- [Entry](docs/Entry.md)
- [EntryInput](docs/EntryInput.md)
- [PaginationMeta](docs/PaginationMeta.md)
- [Project](docs/Project.md)
- [ProjectInput](docs/ProjectInput.md)
- [TokenInfo](docs/TokenInfo.md)
- [Usage](docs/Usage.md)
- [UsageLimits](docs/UsageLimits.md)
- [User](docs/User.md)
- [UserPlan](docs/UserPlan.md)
- [ValidationError](docs/ValidationError.md)

## Utility Functions

All model fields are pointers. Use these helpers to construct values inline:

```go
nitrozen.PtrString("value")
nitrozen.PtrInt(42)
nitrozen.PtrInt32(42)
nitrozen.PtrBool(true)
nitrozen.PtrFloat64(3.14)
```

## License

MIT — see [LICENSE](LICENSE).
