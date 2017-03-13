# ![TinderOnline](http://image.prntscr.com/image/9d2b8e726a944215be4338849ea9c0cd.png)

# TinderOnline ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/TinderOnline) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

TinderOnline is a Go program made on the top of Tinder API.

The TinderOnline's goal is to be a perfect tool providing a stupidly easy-to-use and fast program to find out when was the last time your friends used the Tinder App. It uses your Facebook Token to autenticate and find your friends, then show to you how much time passed since their last use.

## Project Status

TinderOnline is on beta. Pull Requests [are welcome](https://github.com/DiSiqueira/TinderOnline#social-coding)

![](https://i.imgur.com/W7JSwzA.jpg)

## Features

- Open source - You can check out our code
- Secure
- Only one parameter
- 100% satisfaction guaranteed
- It's perfect to find out if your girlfriend or boyfriend are using Tinder
- Instantly see who are online now
- STUPIDLY [EASY TO USE](https://github.com/DiSiqueira/Gorganizer#usage)
- Very fast start up and response time
- Uses native libs

## Installation

### Option 1: Go Get

```bash
$ go get github.com/disiqueira/tinderonline
$ tinderonline -h
```

Make sure that `$GOPATH\bin` is in your PATH, if it is not add it by doing `export PATH=$PATH:$GOPATH/bin`

### Option 2: From source

```bash
$ git clone https://github.com/disiqueira/tinderonline.git tinderonline
$ cd tinderonline/
$ go get -v -d
$ go build *.go
```

## Usage

### Show Help

```bash
# Show help
$ tinderonline -h
```

### Show your friends

```bash
# Show your friends that use Tinder, and who much time since their last usage
$ tinderonline -token=YOURFACEBOOKTOKENHERE
```

## How to get your Facebook Token

1. Open [This link] (https://www.facebook.com/v2.6/dialog/oauth?redirect_uri=fb464891386855067%3A%2F%2Fauthorize%2F&scope=user_birthday,user_photos,user_education_history,email,user_relationship_details,user_friends,user_work_history,user_likes&response_type=token%2Csigned_request&client_id=464891386855067)

2. You will see a dialog that says you have already authorized Tinder. At this point, open your browser’s developer tools. (Cmd + Option + I on Mac) or (F12, Ctrl + Shift + I on Windows) or right click the page anywhere and select 'Inspect'. Switch to the 'Network' tab in your developer tools window. Your dev tools window should look like the image below.
![](https://tinderface.herokuapp.com/fb-auth-window.png)

3. Press the 'Okay' button on the Tinder dialog, and you will see some activity in the Network tab. In the Network tab, locate the new 'confirm?dpr=2' entry, and right click it. Select the 'Copy Response' option from the context menu like in the image below.
![](https://tinderface.herokuapp.com/dev-tools-window.png)

4. After copying the response, paste it into some notepad and remove everything before the `&access_token=` and everything after the `&expires_in=`, you will get only the Facebook Token, it should looks like it:

EAATm0PX4ZCpsAAG1dVq03YCE7BZB58ejN1MvkZAV6NFlJg4BKjhYZAz1G38b9CQfZCbzZCfEXkgkCkGkP1SEARXr3F1uQKjDSZA9RVd0LFt9ZBZDFh3mWo370IayhPIsVi7xJC4SvLEu99mxIHyG2VG07c5EyUI9MGLGC2sGTccHvdw3h6FTZAu6KCv

5. This is the value you should pass when calling the TinderOnline

## Program Help

![](http://image.prntscr.com/image/3a12514f944b48078863961920a7da18.png)

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/TinderOnline/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ git clone https://github.com/disiqueira/tinderonline.git tinderonline
$ cd tinderonline/
$ go get -v -d
$ go run tinderonline.go
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it] (https://github.com/DiSiqueira/TinderOnline/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2017 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
