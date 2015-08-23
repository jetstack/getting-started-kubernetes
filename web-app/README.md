# Simple CMS web app

A very simple CMS web app built with Go that serves up a directory of files, that is used 
throughout the workshop [labs](labs/).

Version 1.0.0 lists files in a directory *inside* the container.
Version 1.1.0 instead (and more properly) lists files in a Kubernetes volume. It also has a small extra UI feature to differentiate.

Credit: [@kelseyhightower](https://twitter.com/kelseyhightower)'s excellent Kubernetes [inspector](https://github.com/kelseyhightower/inspector) 
pod that provided the barebones for the app.
