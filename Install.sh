#!/bin/bash
curl -LO https://github.com/yourname/hinglish/releases/latest/download/hinglish
chmod +x hinglish
sudo mv hinglish /usr/local/bin
echo "Hinglish installed! Run: hinglish <file.hl>"
