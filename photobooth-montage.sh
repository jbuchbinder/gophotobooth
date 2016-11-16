#!/bin/bash
# photobooth-montage.sh
# @jbuchbinder
#
# Create a photobooth-style montage of the four photos created by gophotobooth.
# Specify list of directories (without ending slashes) to create photobooth
# composites.
#
# Requires:
# - imagemagick's 'montage' utility
# - dcraw

SIZE=1000
EXT="CR2"

for dir in $*; do
	d=${dir//\/}
	if [ ! -f ${d}-photobooth.jpg ]; then
	REMOVELIST=""
	FILES=""
	for f in $d/*.${EXT}; do
		BN=$( basename "$f" )
		dcraw -e $f
		mv ${f//.$EXT}.thumb.jpg /tmp/${BN//.$EXT}.thumb.jpg
		REMOVELIST="$REMOVELIST /tmp/${BN//.$EXT}.thumb.jpg"
		FILES="$FILES /tmp/${BN//.$EXT}.thumb.jpg"
	done
	montage -verbose -label '' \
		-font Ubuntu -pointsize 32 \
		-background '#000000' -fill 'gray' -define jpeg:size=${SIZE}x${SIZE} \
		-geometry ${SIZE}x${SIZE}+2+2 -auto-orient $FILES ${d}-photobooth.jpg
	rm -f $REMOVELIST
	fi
done

