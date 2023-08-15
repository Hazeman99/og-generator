# OG Portfolio Image Generator CLI Tool

## Description

Create an OG image for an article on your website with this CLI tool.

**Note: The code is pretty messy and needs to be refactored.**

## Pre-requisites

- Go 1.18
- Background Image - 1200 X 630 pixels
- Icon Image - 48 X 48 pixels

## Usage

Make sure that you have the **Background Image** (file named background.png) and **Icon Image** (file names icon.png) in the same directory as the executable.

Run the following command:

```bash
go run . --title="<TITLE HERE>" --description="<DESCRIPTION HERE>" --minutes=7 --author="<AUTHOR NAME HERE>" --website="<WEBSITE URL HERE>"
```

Or you can build the executable and run it:

```bash
go build .
./og-image-generator --title="<TITLE HERE>" --description="<DESCRIPTION HERE>" --minutes=7 --author="<AUTHOR NAME HERE>" --website="<WEBSITE URL HERE>"
```

## Credits

Heavily relies on the below referenced article:
https://pace.dev/blog/2020/03/02/dynamically-generate-social-images-in-golang-by-mat-ryer.html