## My notes on Michael's [corresponding day](https://www.90daysofdevops.com/2022/day15/)
Michael explains a lot of basic commands on that day. But I have to confess I 
didn't know the "locate" command. Had to install it first, but nice command,
because the the syntax is easier than "find" and it's also faster, because it
just scans a prebuilt Linux database instead of the whole filesystem. Than of
course one must be aware of updating the DB manually if files are not found.

I have to say, that I don't like Michael's writing much, it's kind of sloppy
and hard to follow. I know, my writing isn't much better, but I hope Michael's
blogging gets better, otherwise I might change to another resource for learning.
To be fair, it started as his own learning path, but now it's promoted as a
learning resource for everybody.

## Commands new to me
### locate
- easier syntax than "find"
- faster than find

### awk
- scripting language for searching patterns in text and performing actions on the
results

### cut
- cut parts of lines from files and print the result to stout
- cut by delimiter, byte position, character
- no, this doesn't remove anything from the file this is applied to (other than
Michael states)

### xargs
- converts input from stdin into arguments to a command
- grep, for example, can take input either as arguments or from stdin, but cp
or echo can only take
input as arguments, which is why xargs is necessary
- example: 'find -name "test.txt" | xargs grep "testing"'

### getent
- get entries from all textfiles or databases that are configured in /etc/nsswitch.conf
- for examples to find out which service runs on which port: "getent services <portnumber>"
or which port does a service use: "getent services <servicename>"

I'll some resources, sorry some of them are in german, for anybody that ever
actually follows this :wink:)

## Resources
- [awk command](https://www.tutorialspoint.com/unix_commands/awk.htm)
- [awk with good examples, german resources](https://www.linux-community.de/ausgaben/linuxuser/2005/10/awk-werkzeug-und-skriptsprache/)
- [cut command](https://linuxize.com/post/linux-cut-command/)
- [xargs command](https://en.wikipedia.org/wiki/Xargs)
- [xargs command, german resource](https://www.ionos.de/digitalguide/server/konfiguration/linux-xargs-befehl/)
- [getent command, german resource](https://www.my-it-brain.de/wordpress/getent-informationen-ueber-user-gruppen-services-und-mehr-anzeigen/)
