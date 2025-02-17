package imageprocessing_test

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"testing"
)

// Test ReadImage function
func TestReadImage(t *testing.T) {
	// Define the image path to test reading the image
	imagePath := "test_image.jpeg"

	// Call the ReadImage function from the imageprocessing package
	img := imageprocessing.ReadImage(imagePath)

	// Check if the image is nil, which indicates an error occurred during loading
	if img == nil {
		t.Errorf("Error reading image %s: got nil", imagePath)
	}

	// Ensure that we don't get a nil image
	if img == nil {
		t.Errorf("Expected non-nil image, but got nil")
	}
}

// Test Resize function
func TestResize(t *testing.T) {
	// Create a new image (500x500) to test resizing
	originalImage := image.NewRGBA(image.Rect(0, 0, 500, 500))

	// Call the Resize function from the imageprocessing package
	resizedImage := imageprocessing.Resize(originalImage)

	// Check if the resized image has the expected dimensions (500x500)
	if resizedImage.Bounds().Dx() != 500 || resizedImage.Bounds().Dy() != 500 {
		// If the dimensions are incorrect, we report the error
		t.Errorf("Resize failed, expected 100x100 image but got %dx%d", resizedImage.Bounds().Dx(), resizedImage.Bounds().Dy())
	}
}

// Test Grayscale function
func TestGrayscale(t *testing.T) {
	// Create a new image (100x100) to test grayscale conversion
	originalImage := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Call the Grayscale function from the imageprocessing package
	grayImage := imageprocessing.Grayscale(originalImage)

	// Check if the grayscale function returned a non-nil image
	if grayImage == nil {
		t.Errorf("Grayscale function returned nil image")
	}
}

// processPipeline simulates image processing with or without goroutines
func processPipeline(imagePaths []string, useGoroutines bool) {
	// Loop through all image paths
	for _, path := range imagePaths {
		// Load the image using ReadImage function
		img := imageprocessing.ReadImage(path)

		// Check if we are using goroutines to process images
		if useGoroutines {
			// If using goroutines, process resizing and grayscale in a separate goroutine
			go func() {
				img = imageprocessing.Resize(img)    // Resize the image
				img = imageprocessing.Grayscale(img) // Convert the image to grayscale
			}()
		} else {
			// If not using goroutines, process sequentially
			img = imageprocessing.Resize(img)
			img = imageprocessing.Grayscale(img)
		}
	}
}

// BenchmarkPipelineWithGoroutines benchmarks the pipeline with goroutines enabled
func BenchmarkPipelineWithGoroutines(b *testing.B) {
	// Define a list of image paths to be processed
	imagePaths := []string{
		"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	// Benchmark the process by running it b.N times
	for i := 0; i < b.N; i++ {
		// Call the processPipeline function with goroutines enabled
		processPipeline(imagePaths, true)
	}
}

// BenchmarkPipelineWithoutGoroutines benchmarks the pipeline without goroutines
func BenchmarkPipelineWithoutGoroutines(b *testing.B) {
	// Define a list of image paths to be processed
	imagePaths := []string{
		"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	// Benchmark the process by running it b.N times
	for i := 0; i < b.N; i++ {
		// Call the processPipeline function without goroutines
		processPipeline(imagePaths, false)
	}
}
