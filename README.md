Taskmaster
============
Taskmaster helps you focus on the important work at hand, so you can free up brain time for more important problems.

# Why use the Most Important Tasks (MIT) system?
[More Info](https://personalmba.com/most-important-tasks/)

Modern work life is a daily struggle to keep up with a never ending list of TODOs. It requires us to organize and prioritize work on the fly. the MIT task list system is a way to organize these competing priorities and become realistic about which tasks you can reasonably accomplish.

In the MIT list, your tasks are broken up into two types: Most Important Tasks (urgent, do or die) and TODOs (important, but not urgent). MITs are limited to a small number (3) to force you to prioritize whenever adding a new one. The TODOs list can grow larger, but must be groomed on a regular basis.

Taskmaster will force task limits on you: 3 MITs and 10 TODOs. It is very opinionated in not providing a backlog (which will just grow indefinitely) and forcing you to groom your tasks whenever hitting task limits.

```
$ tkm mit Email Widget Co.
                                                                                         
Error: cannot add any more tasks to list

MITs:
0) Plan for Retro
1) Schedule lunch with FizzBuzz CFO
2) Ticket FizzBuzz project tickets

TODOs:
3) Fix a non-urgent bug
```

```
$ tkm mv 2
MITs:
0) Plan for Retro
1) Schedule lunch with FizzBuzz CFO

TODOs:
2) Fix a non-urgent bug
3) Ticket FizzBuzz project tickets
```

```
$ tkm mit Email Widget Co.
MITs:
0) Plan for Retro
1) Schedule lunch with FizzBuzz CFO
2) Email Widget Co.

TODOs:
3) Fix a non-urgent bug
4) Ticket FizzBuzz project tickets
```

# Installation
### Via Homebrew
`$ brew install SimantovYousoufov/tap/taskmaster`

### Via Releases
1. Download the relevant [release](https://github.com/SimantovYousoufov/taskmaster/releases)
1. Drop the binary in your $PATH

### From source
1. Clone the repository
1. `$ go install .`

# Usage
```
$ tkm --help
Taskmaster is a better way to manage your tasks

Usage:
  tkm [flags]
  tkm [command]

Available Commands:
  daemon      Run the taskmaster daemon
  help        Help about any command
  mit         Add a task to the MIT list
  rm          Remove a task from the task list by index. Use `tkm stat` for indexes.
  stat        View the task set. Aliases: ls, list, all
  todo        Add a task to the TODO list
  version     Print the version number of Taskmaster
```

### Adding a task
```
$ tkm mit Email Widget Co.
MITs:
0) Plan for Retro
1) Schedule lunch with FizzBuzz CFO
2) Email Widget Co.

TODOs:
3) Fix a non-urgent bug
4) Ticket FizzBuzz project tickets
```

```
$ tkm todo Email Widget Co.
MITs:
0) Plan for Retro
1) Schedule lunch with FizzBuzz CFO

TODOs:
2) Fix a non-urgent bug
3) Ticket FizzBuzz project tickets
4) Email Widget Co.
```

### Removing a task
```
$ tkm rm 3
MITs:
0) Plan for Retro
1) Schedule lunch with FizzBuzz CFO

TODOs:
2) Fix a non-urgent bug
3) Email Widget Co.
```

### Reviewing your tasks
```
$ tkm
MITs:
0) Plan for Retro
1) Schedule lunch with FizzBuzz CFO

TODOs:
2) Fix a non-urgent bug
3) Email Widget Co.
```

```
$ tkm stat|ls|list|all
MITs:
0) Plan for Retro
1) Schedule lunch with FizzBuzz CFO

TODOs:
2) Fix a non-urgent bug
3) Email Widget Co.
```

### De-prioritizing a task
```
$ tkm mv 1
MITs:
0) Plan for Retro

TODOs:
1) Fix a non-urgent bug
2) Email Widget Co.
3) Schedule lunch with FizzBuzz CFO
```

### Configuring Taskmaster
```
$ tkm config
MIT Task Limit (number): 3
✔ Todo Task Limit (number): 10
```

# Data Privacy
Taskmaster keeps all its data local on your device (`$HOME/.taskmaster.json` by default). Your task list is stored plaintext in JSON.
