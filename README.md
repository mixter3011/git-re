# Git-Re
A continuation of the Git-Add CLI tool. Stages, Adds and commits changes to an existing repository by taking a commit message

### Features
- Stages the Changes made to the directory into the Git pipeline.
- Adds the Staged changes to the repository path.
- Commits the changes to the repository.

### Usage 
- First up clone this repo in your root directory
``` 
    git clone https://github.com/yourusername/git-automator.git

```
- Build the Go executable:
```
cd git-re
go build -o git-re main.go
```
- Then add the this Cloned directory path to your environment variables. How to add path to environment variables can be found on this link : http://surl.li/sxrdz

- To use Git Re, simply run the following command in your terminal:
``` 
    git-re

```
- Follow the prompts to enter a commit message when prompted.

### Prerequisites
- Go programming language installed
- GitHub CLI (gh) installed and authenticated
