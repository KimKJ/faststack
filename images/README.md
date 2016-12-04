These are the images that Termbox provides by default.

To add an image, create a Dockerfile at `<name>/<tag>/Dockerfle` and add an
entry to `images.json`.

All images should have the following software installed:

* Basic unix tools (coreutils/busybox)
* tmux, curl
* Version control: git, subversion
* C/C++ compiler, make
* Nodejs, npm
* Ruby, bundler
* golang
* Rustc, cargo
* Python, pip