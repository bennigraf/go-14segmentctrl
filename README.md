# go-14segmentctrl

This tool is used on my Raspberry PI Zero W based pedal, that controls a Zoom B3 pedal via MIDI in 
conjunction with https://github.com/bennigraf/zoomb3ctrl. This tool basically acts as a bridge to
display the patch names to a 8-digit 14-segment display ([bought here](https://shop.wtihk.com/index.php?route=product/product&product_id=70)).

It Listens on port 3024 (tcp) and prints received messages that 14 segment display (driven by a HT16K33 driver chip).
