# Go notes

This is a simple Go project to help me get notified when I need to attend a
meeting.

I use a pretty simple workflow to take notes. I have a directory in my home
directory called `notes`, which is synced between my computers via a private
repository hosted on GitHub.

Inside this `notes` directory, I keep a subdirectory called `journal` where I
keep my journals for each day. Those journal files include thoughts, tasks,
meetings and notes in general. I use `marksman` LSP to link my notes and quick
jump between them.

But also, I'd like to get notified when I need to attend a meeting, so I created
this project alongside with a [systemd service and timer]() that triggers this
program every minute to check if I need to attend a meeting.

I also spend almost all my time during coding inside a TMUX server, running
inside Alacritty occupying the whole screen and hiding interruptions, this is
why I also need to get notified at TMUX.

## Setup

Clone this directory anywhere and then run `go build -o notes && mv notes
~/.bin/notes`.
