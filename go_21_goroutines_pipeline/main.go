package main

import (
	"fmt"
	"image"
	"strings"
	"time"

	imageprocessing "goroutines_pipeline/image_processing"
)

type Job struct {
	InputPath string      // Path to the input image
	Image     image.Image // The loaded image to be processed
	OutPath   string      // Path where the output image will be saved
	Err       error       // Error field for error tracking during processing
}

// loadImage loads an image file and checks for errors.
func loadImage(paths []string) <-chan Job {
	out := make(chan Job) // Create a channel to send jobs (image paths and image data)
	go func() {
		// Iterate over each image path
		for _, p := range paths {
			// Create a job for each image path
			job := Job{InputPath: p, OutPath: strings.Replace(p, "images/", "images/output/", 1)}

			// Read the image and store it in the job
			job.Image = imageprocessing.ReadImage(p)

			// If the image is nil, there was an error loading it
			if job.Image == nil {
				job.Err = fmt.Errorf("failed to load image %s", p) // Set the error field
				fmt.Printf("Error loading image %s\n", p)          // Print error message
			}

			// Send the job to the channel
			out <- job
		}
		close(out) // Close the channel when done sending all jobs
	}()
	return out // Return the channel of jobs
}

// resize resizes the image and checks for errors.
func resize(input <-chan Job) <-chan Job {
	out := make(chan Job) // Create a new channel for resized images
	go func() {
		// Loop through all incoming jobs from the input channel
		for job := range input {
			// If there was an error with the job, skip the resizing
			if job.Err != nil {
				out <- job
				continue
			}

			// Resize the image and update the job
			job.Image = imageprocessing.Resize(job.Image)

			// Send the updated job to the output channel
			out <- job
		}
		close(out) // Close the channel when done
	}()
	return out // Return the channel of resized jobs
}

// convertToGrayscale converts the image to grayscale and checks for errors.
func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job) // Create a new channel for grayscale images
	go func() {
		// Loop through all incoming jobs from the input channel
		for job := range input {
			// If there was an error with the job, skip the grayscale conversion
			if job.Err != nil {
				out <- job
				continue
			}

			// Convert the image to grayscale and update the job
			job.Image = imageprocessing.Grayscale(job.Image)

			// Send the updated job to the output channel
			out <- job
		}
		close(out) // Close the channel when done
	}()
	return out // Return the channel of grayscale jobs
}

// saveImage saves the processed image and handles errors.
func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool) // Create a new channel for success/failure results
	go func() {
		// Loop through all incoming jobs from the input channel
		for job := range input {
			// If there was an error with the job, don't try to save the image
			if job.Err != nil {
				out <- false // Send a failure result
				continue
			}

			// Save the processed image to the output path
			imageprocessing.WriteImage(job.OutPath, job.Image)

			// Send a success result
			out <- true
		}
		close(out) // Close the channel when done
	}()
	return out // Return the channel of success/failure results
}

// processPipeline processes images through the pipeline with or without goroutines.
func processPipeline(imagePaths []string, useGoroutines bool) {
	// Record the start time to measure the pipeline execution duration
	start := time.Now()

	// Initialize the pipeline steps: load image, resize, convert to grayscale, and save
	channel1 := loadImage(imagePaths)        // Load the images
	channel2 := resize(channel1)             // Resize the images
	channel3 := convertToGrayscale(channel2) // Convert to grayscale
	writeResults := saveImage(channel3)      // Save the processed images

	// Loop through the results of the saveImage step and log success or failure
	for success := range writeResults {
		if success {
			// If the result is true, print success
			fmt.Println("Success!")
		} else {
			// If the result is false, print failure
			fmt.Println("Failed!")
		}
	}

	// Calculate and print the total duration of the pipeline execution
	duration := time.Since(start)
	fmt.Printf("Pipeline completed in: %v\n", duration)
}

func main() {
	// Define a list of image paths for testing
	imagePaths := []string{
		"images/image1.jpeg", // Replace with your own image
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	// Run pipeline with goroutines
	fmt.Println("Running pipeline with goroutines:")
	processPipeline(imagePaths, true)

	// Run pipeline without goroutines
	fmt.Println("Running pipeline without goroutines:")
	processPipeline(imagePaths, false) // Process without goroutines
}
