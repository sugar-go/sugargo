#!/bin/bash

go install github.com/lingdor/gmodeltool@v0.0.9
curl -O -s https://raw.githubusercontent.com/sugar-go/sugargo/main/demo.tgz
tar zxvf demo.tgz
rm demo.tgz
