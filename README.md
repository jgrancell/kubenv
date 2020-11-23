# Kubenv

A simple Kubernetes configuration manager.

[![kubenv](https://img.shields.io/travis/com/jgrancell/kubenv?style=for-the-badge&logo=travis)](https://travis-ci.com/github/jgrancell/kubenv)
[![coverage](https://img.shields.io/codecov/c/github/jgrancell/kubenv?color=65187a&style=for-the-badge&token=p8NQJsRPDX)](https://codecov.io/gh/jgrancell/kubenv/)

[![releases](https://img.shields.io/github/v/release/jgrancell/kubenv?style=for-the-badge)](https://github.com/jgrancell/kubenv/releases)
[![GitHub license](https://img.shields.io/github/license/jgrancell/kubenv?color=333333&style=for-the-badge)](https://github.com/Jguer/yay/blob/master/LICENSE)

## Objectives

Kubenv was designed to solve a single purpose -- management of a large number of
Kubernetes kubeconfig files while:
* Not having to constantly pass `KUBECONFIG=` environment values
* Avoiding an unwieldly bash_aliases file.
* Constantly remember which alias/kubeconfig you're currently running under.

## Features

* Importing kubeconfig files from anywhere in your filesystem.
* Listing all currently-registered kubeconfig files.
* Easily swapping between registered kubeconfig files.
* Setting custom locations for your registered kubeconfig files.

Planed features:
* Omni-OS compatibility (coming soon™)

## Installation

All releases contain binary packages for the following operating systems:

| Operating System |  Arch |   Support    |
| :--------------- |:-----:| :----------: |
| Linux            | amd64 |   [![supported](https://img.shields.io/badge/status-supported-blue?style=flat-square)](https://github.com/jgrancell/kubenv/releases)  |
| Linux            |  arm  |   [![supported](https://img.shields.io/badge/status-supported-blue?style=flat-square)](https://github.com/jgrancell/kubenv/releases)  |
| OSX              | amd64 |   [![supported](https://img.shields.io/badge/status-supported-blue?style=flat-square)](https://github.com/jgrancell/kubenv/releases)  |
| OSX              | arm |   [![coming soon](https://img.shields.io/badge/status-planned-orange?style=flat-square)](https://github.com/jgrancell/kubenv/releases)  |

Packages for some Linux operating systems as well as OSX homebrew are planned.
