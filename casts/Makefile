WIDTH = 80
HEIGHT = 24
CMD = asciinema rec --cols $(WIDTH) --rows $(HEIGHT) --stdin --overwrite
SVG = docker run --rm -i bric3/svg-term-cli svg-term --width $(WIDTH) --height $(HEIGHT)
GIF = agg --theme solarized-dark --cols $(WIDTH) --rows $(HEIGHT)
CASTS = usage

all: $(addsuffix .gif, $(CASTS))

clean:
	rm -f *.cast
	rm -f *.svg
	rm -f *.gif
	rm -f *.mp4

$(addsuffix .cast, $(CASTS)): %.cast : %.sh
	bash $(basename $@).sh | $(CMD) -t "Trelldo: $(basename $@)" $@

$(addsuffix .svg, $(CASTS)): %.svg : %.cast
	$(SVG) < $(basename $@).cast > $@

$(addsuffix .gif, $(CASTS)): %.gif : %.cast
	$(GIF) $(basename $@).cast $@

$(addsuffix .mp4, $(CASTS)): %.mp4 : %.gif
	ffmpeg -i $(basename $@).gif -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" $@
