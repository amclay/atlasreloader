# atlasreloader
atlas minigame player


How it works:
  Takes screenshot to memory, then looks for two adjacent red/white pixels within the minigame area. Clicks when it finds it.
  
Tried this with AHK, but it was way too slow. Additionally calling PixelGetColor many times in a row was extremely slow (6 secoonds). This is ~20-30ms per loop.
