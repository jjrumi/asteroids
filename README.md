# Asteroids

This is a hand-crafted Asteroids game built on top of the OpenGL-based library Pixel: https://github.com/faiface/pixel

## Build & Run

Use the following make target to build & run the project:
```
$ make run
```

### Build requirements
This project requires C code from the https://github.com/go-gl/glfw library to compile.
We use https://github.com/goware/modvendor to download it locally along with the vendor libraries. 

# ToDo
- Draw thrust
- Ship can fire
- Explode asteroids when hit by fire