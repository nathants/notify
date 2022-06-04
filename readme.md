# notify

## why

sometimes popup messages are for great good.

## what

a cli that triggers a fullscreen graphical popup with text.

optionally includes a y/n prompt which will change the exit code.

a delay before the prompt is answerable avoids accidental input.

## example

![](https://github.com/nathants/notify/raw/master/example.gif)

## install

`go install github.com/nathants/notify@latest`

## usage

```bash
>> notify -h

notify the user of a message with a fullscreen popup. hit Q or ENTER to exit.

Usage: notify [--prompt] [--delay-seconds DELAY-SECONDS] [--center] [MESSAGE]

Positional arguments:
  MESSAGE                the message to display on screen

Options:
  --prompt, -p           prompt the user for a y/n response, and exit 0/1 accordingly
  --delay-seconds DELAY-SECONDS, -d DELAY-SECONDS
                         delay seconds before accepting user input for prompted y/n [default: 1]
  --center, -c           horizontally center each line
  --help, -h             display this help and exit
```
