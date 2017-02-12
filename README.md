# goSoundbench

goSoundbench is a collection of tools which helps conduction automated
tests for audio devices.

## How to install

In order to build the project, you need install the [Go language][1]. It's
recommended to install the latest version.

### MacOS

Its recommended to use the superb packet manager [homebrew][2] to install the
dependencies:

```bash

$ brew install portaudio
$ go get github.com/dh1tw/goSoundbench
$ cd $GOPATH/src/github.com/dh1tw/goSoundbench
$ make install

```

### Linux (Ubuntu)

```bash

$ sudo apt-get install libportaudio2 portaudio19-dev portaudio19-doc git
$ go get github.com/dh1tw/goSoundbench
$ cd $GOPATH/src/github.com/dh1tw/goSoundbench
$ make install

```

## Soundcards

Linux does not guarantee that USB Sound cards are always mapped as the same
device. However with a combination of udev and alsa rules we can make sure that
a particular USB Sound card gets a persistant handle.

Use the following command to see the path of the USB Device (you have to 
plug or unplug the device):

```bash

udevadm monitor --kernel --subsystem-match=sound

```

Once we know the path of the device, we create the following file:

```bash

sudo vim /etc/udev/rules.d/70-alsa-permanent.rules

```

and we add something like:

```bash

SUBSYSTEM!="sound", GOTO="my_usb_audio_end"
ACTION!="add", GOTO="my_usb_audio_end"

DEVPATH=="/devices/platform/soc/3f980000.usb/usb1/1-1/1-1.2/1-1.2:1.0/sound/card?", ATTR{id}="CH1"
DEVPATH=="/devices/platform/soc/3f980000.usb/usb1/1-1/1-1.3/1-1.3:1.0/sound/card?", ATTR{id}="CH2"
DEVPATH=="/devices/platform/soc/3f980000.usb/usb1/1-1/1-1.4/1-1.4:1.0/sound/card?", ATTR{id}="CH3"
DEVPATH=="/devices/platform/soc/3f980000.usb/usb1/1-1/1-1.5/1-1.5:1.0/sound/card?", ATTR{id}="CH4"


LABEL="my_usb_audio_end"

```

This above example maps all 4 USB ports of a Raspberry Pi to sound cards.

Now we can create an alsa configuration in `$HOME/.asoundrc`:

```bash

pcm.CH1 {
	type hw
	card "CH1"
    	device 0
}

ctl.CH1 {
	type hw
	card "CH1"
    	device 0
}

pcm.CH2 {
        type hw
        card "CH2"
        device 0
}

ctl.CH2 {
        type hw
        card "CH2"
        device 0
}
pcm.CH3 {
        type hw
        card "CH3"
        device 0
}

ctl.CH3 {
        type hw
        card "CH3"
        device 0
}
pcm.CH4 {
        type hw
        card "CH4"
        device 0
}

ctl.CH4 {
        type hw
        card "CH4"
        device 0
}

```

If you don't need the internal Sound card (of a Raspberry Pi), disable it:

```bash

sudo vim /etc/modprobe.d/sound.conf

```

and add:

```bash

blacklist snd-bcm2835
snd-usb-audio

```

You might have to the now remove or comment out invalid sound devices
from `/usr/share/alsa/alsa.conf`. In my case I had to remove:


```bash

#pcm.rear cards.pcm.rear
#pcm.center_lfe cards.pcm.center_lfe
#pcm.side cards.pcm.side
#pcm.surround21 cards.pcm.surround21
#pcm.surround40 cards.pcm.surround40
#pcm.surround41 cards.pcm.surround41
#pcm.surround50 cards.pcm.surround50
#pcm.surround51 cards.pcm.surround51
#pcm.surround71 cards.pcm.surround71
#pcm.iec958 cards.pcm.iec958
#pcm.spdif iec958
#pcm.hdmi cards.pcm.hdmi
#pcm.modem cards.pcm.modem
#pcm.phoneline cards.pcm.phoneline

```

[1]:https://golang.org
[2]:https://brew.sh