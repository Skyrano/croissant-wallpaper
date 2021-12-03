Small project to change automatically the desktop wallpaper for Linux, Windows and Mac.

Usage:

 - Copy the binary files (*bin* directory) to the usage directory (like a USB key for example)

 - Copy the *config.yml* file (*resources* directory) alongside the binaries

 - Update the configuration file with the path to the image or folder you want to use

If *random* is set to false, the *wallpaper* path will be used to set the wallpaper image.
If *random* is set to true, an image from the *imagesDir* directory will be choosen randomly and set as wallpaper.
Take note that all the paths given in the *config.yml* file are relative to the binary executed (as the config file itself which must be in the same directory).