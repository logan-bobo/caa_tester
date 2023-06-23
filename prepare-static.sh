#!/usr/bin/env bash

#We need to move the stylesheet out of node modules so it can be served by the static directory on our server
cp node_modules/dracula-ui/styles/dracula-ui.css static/