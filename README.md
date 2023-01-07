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

My journals have this structure:

```md
## Meetings
- [ ] 10:00 Daily
- [ ] 14:30 1:1 with some person

## Tasks
- [ ] I need to get this done

## Notes
blah blah
```

I also spend almost all my time during coding inside a TMUX server, running
inside Alacritty occupying the whole screen and hiding interruptions, this is
why I also need to get notified at TMUX.

## Setup

Clone this directory anywhere and then run `go build -o notes && mv notes
~/.bin/notes`.

My [dotfiles]() are my home directory. To enable the systemd service and timer,
I must run:

```
$ systemctl --user enable schedule-notes.service
$ systemctl --user enable schedule-notes.timer
```

Then run the following command to start the timer service:

```
$ systemctl --user start schedule-notes.timer
```

The service must be running now

```
$ systemctl --user status schedule-notes.timer

● schedule-notes.timer - Schedule a message every 1 minute
     Loaded: loaded (/home/gustavo/.config/systemd/user/schedule-notes.timer; enabled; vendor preset: enabl>
     Active: active (running) since Sat 2023-01-07 12:14:21 -03; 2min 26s ago
    Trigger: n/a
   Triggers: ● schedule-notes.service

Jan 07 12:14:21 tinhoso-pop systemd[2701]: Started Schedule a message every 1 minute.
```
