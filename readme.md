# notify

## why

sometimes popup messages are for great good.

## what

a fullscreen terminal app, typically launched in a new single use terminal, to display a popup message.

an optional y/n prompt will change the exit code.

an optional delay avoids accidental input.

## demo

![](https://github.com/nathants/notify/raw/master/demo.gif)

## install

`python3 -m pip install git+https://github.com/nathants/notify`

## usage

```bash
>> notify -h

usage: notify [-h] [-d DELAY] [-p] msg

    notify the user of a message with a fullscreen popup. hit any key to exit.


positional arguments:
  msg                   message to display

options:
  -h, --help            show this help message and exit
  -d DELAY, --delay DELAY
                        delay seconds before accepting user input for prompt (default: 0.5)
  -p, --prompt          prompt the user y/n then exit 0/1 (default: False)

```
