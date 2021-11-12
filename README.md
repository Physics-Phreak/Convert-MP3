# Convert-MP3

Convert all the songs in a folder into mp3 and store them in a sub folder

## What is This?

This is just a small program I wrote to make it easier for me to convert my .flac files to mp3. I wrote it in GoLang because I had just started learning it, and thought it would be an interesting beginner project.<br>

The code is very simple, and has a ton of room for improvement. Feel free to mess with the code to suit your needs.

## How To Run
**All instructions are for Linux specifically, you can adapt it to suit your os**<br>
<br>

*Please Note: This program is basically using FFmpeg to convert the songs, so you will need it installed and added to path*<br>

- Download the GoLang files and build it, or download the convert-mp3.sh file
- Move the convert-mp3.sh files to ~/bin to access it from any folder
- Navigate to the folder with the music files
- Run `convert-mp3` to run the program. This will create a sub-directory called "MP3", inside of which the mp3 files for all your songs will be saved