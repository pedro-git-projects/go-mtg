for f in BG.svg BR.svg CP.svg GP.svg GU.svg RG.svg R.svg T.svg \
         UP.svg U.svg WP.svg WU.svg BP.svg B.svg C.svg G.svg GW.svg \
         RP.svg RW.svg UB.svg UR.svg WB.svg w.svg; do
  magick "$f" \
    -background none \
    -alpha set \
    "${f%.svg}.png"
done
