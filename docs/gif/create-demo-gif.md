# Create demo gif

This works with

ImageMagick 7.1.0-45

bash 5.1.16

iTerm2 3.4.16

macOS 12.5

1. Install requirements

   * Copy `demo-magic.sh` file from https://github.com/paxtonhare/demo-magic to here
   * https://github.com/icholy/ttygif
   * https://github.com/icetee/pv
   * ImageMagick

2. Uninstall global icm on your machine and install current development build to Go binary path

   ```
   brew uninstall icm
   make install
   ```

3. iTerm adjustments

   1. Use 80x25 terminal size

   2. Change font size to 10pt

      `iTerm -> Preferences -> Profiles -> Text`

4. Start recording

   ```
   ttyrec demo
   ```

5. Execute scripted session

   ```
   ./demo.sh
   ```

6. Stop recording

   Ctrl + D

7. Execute ttygif and abort **at the end** of the execution while generating gif

   `ttygif demo` + Ctrl + C

8. Copy aborted command in bash script

   ```
   echo '#!/bin/sh' > /tmp/convert.sh
   chmod +x /tmp/convert.sh
   # Copy convert command
   pbpaste >> /tmp/convert.sh
   ```

   If convert command does not work, try adding debug flags:

   ```
   convert -debug "Cache,Blob" ...
   ```

9. Remove unwanted frames in arguments

10. Add `-colors 16 -depth 4` flags

11. Execute script

    ```
    /tmp/convert.sh
    ```

12. Copy gif to here or rename it if `tty.gif` is in this directory

    ```
    cp /tmp/tty.gif ./demo.gif
    ```

    or

    ```
    mv tty.gif demo.gif
    ```

