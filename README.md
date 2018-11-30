# lichess socket proxy

This proxies lichess websockets to your local system.

To use:

1. Install ca cert from "li.cert" as a trusted root on your system. This is a dummy ca root used to sign a cert for `socket.lichess.org` so your browser will fall for this.
2. Edit host file to point `socket.lichess.org` at `127.0.0.1`.
3. Run this app.
4. Play on lichess.org
5. Don't cheat

Intended purpose is to adapt a physical board to play online. Bot accounts are fairly limited (can't matchmake at all), and the websockets are very finicky to reproduce so as to not get cut off. This will proxy and allow you to read and write moves mid-stream.

Yes, it can be used to cheat. Don't.