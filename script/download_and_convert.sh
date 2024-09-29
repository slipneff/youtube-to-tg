#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <YouTube URL> <output filename>"
    exit 1
fi

youtube_url="$1"
output_filename="$2"

youtube-dl -f bestaudio --extract-audio --audio-format mp3 --output "$output_filename" "$youtube_url"

if [ $? -eq 0 ]; then
    echo "Video downloaded and converted to MP3 successfully."
else
    echo "Failed to download or convert the video."
    exit 1
fi