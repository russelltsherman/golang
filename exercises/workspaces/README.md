# workspaces

[Go 1.18 Workspaces](https://www.youtube.com/watch?v=NrN2OpVt804)

GOPATH
offered a logical way or organizaing packages for a project
prescribed a narrow project organization

workspaces allow dependent modules to be organized as a logical unit that is resolved locally on disk

`go work` and go.work

a workspace is comprised of local modules and directories

once you declare a workspace
run your go commands from within that workspace

benefits
simpler workflow
code with multiple dependent modules
work with unpublished modules
unique workspace config per project/profile
