[![Gitlab pipeline status](https://gitlab.com/ollaww/goameoflife/badges/main/pipeline.svg)](https://gitlab.com/ollaww/goameoflife/commits/main)

# Goameoflife

Simple Go implementation of John Conway [Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) to learn and play with the language.

### Build and run
Just open a terminal big enough (I didn't care too much about sizing) and run

```bash
docker build . -t awesome-game
docker run -it awesome-game
```