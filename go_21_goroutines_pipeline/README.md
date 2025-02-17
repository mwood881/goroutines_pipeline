## PIPELINES FOR IMAGES USING CURRENCY

# Overview
This project is an image processing pipeline that uses Go currency. We need to analyze and optimize data processing times with goroutines. This is based on Amrit Singh (CODEHEIM) code and tutorial. 

# Objectives
- Clone and run Github as is
- Use error checking for input and outputting images
- Change the images with own images
- Unit tests
- Benchmark test for times
- Allow programs executed either with or without goroutines

## Setups

### Prerequisites: 

1) Install Go: https://go.dev/doc/install


2) Clone Github repository
  ```sh
  git clone <repository-url>
  cd image-processing-pipeline
  ```

3) Run origonal program
  ```sh
   go run main.go
   ```

4) Add 4 new images

   
5) Run Unit Tests
```sh
 go test ./...
```

6) Run benchmark Test
```sh
 go test -bench .
```

7) Challenge: A method to maintain aspect ratio instead of fixed 500x500 resizing.
  

8) This project demonstrates the impact of Go concurrency on image processing times. While concurrency improves performance, it is essential to manage resources efficiently to avoid bottlenecks. Further optimizations can enhance processing speed and output quality.

### References:
GITHUB to Clone: https://github.com/code-heim/go_21_goroutines_pipeline 
CODEHEIM: https://www.codeheim.io/ 
Tutorial: https://www.youtube.com/watch?v=8Rn8yOQH62k
