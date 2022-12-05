# Create demo gif

1. Install https://github.com/charmbracelet/vhs.
2. Uninstall icm

   ```bash
   brew uninstall icm
   ```

3. Install local build

   ```bash
   (cd ../.. && make install)
   ```

4. Run vhs command

   ```bash
   vhs < demo.tape
   ```