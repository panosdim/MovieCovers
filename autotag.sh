#!/bin/bash
# ---------------------------------------------------------------------------
# autotag.sh -- Watch for file changes in remote/local file systems and Auto 
#               Tag movies files with covers. It also triggers a miniDlna DB
#               rebuild.
#
# Copyright 2020 Panagiotis Dimopoulos (panosdim@gmail.com)
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License at (http://www.gnu.org/licenses/) for
# more details.
#
# Version: 1.0
# ---------------------------------------------------------------------------

TMP=$(mktemp)
MOD_TIMES=".modtimes"

find "$1" -type d -not -name "$2" -exec stat -c "%Y" '{}' \; > $TMP

if ! cmp --silent $TMP $MOD_TIMES; then 
    echo " `date` : New file changes detected. "
    cat $TMP > $MOD_TIMES
    /opt/autotag/autotag -p "$1" -e "$2"
    minidlnad -R
    systemctl restart minidnla.service
fi