package main

import (
	"ImageProcessing/filter"
	"ImageProcessing/histogram"
	"ImageProcessing/io"
	"ImageProcessing/noise"
)

func main() {
	// open original.png
	original, err := io.Open("./img/original.png")
	if err != nil {
		panic(err)
	}

	// convert original.png to grayscale
	grayscale := filter.Grayscale(original)
	err = io.Save("./img/grayscale.png", grayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of grayscale.png
	grayscalHistogram := histogram.GetGrayHistogram(grayscale)
	err = io.Save("./img/histogramImage/grayscalHistogram.png", grayscalHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Box Blur
	boxBlur := filter.BoxBlur(grayscale, 10)
	err = io.Save("./img/boxBlur.png", boxBlur, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of boxBlur.png
	boxBlurHistogram := histogram.GetGrayHistogram(boxBlur)
	err = io.Save("./img/histogramImage/boxBlurHistogram.png", boxBlurHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Gaussian Blur
	gaussianBlur := filter.GaussianBlur(grayscale, 10)
	err = io.Save("./img/gaussianBlur.png", boxBlur, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of boxBlur.png
	gaussianBlurHistogram := histogram.GetGrayHistogram(gaussianBlur)
	err = io.Save("./img/histogramImage/gaussianBlurHistogram.png", gaussianBlurHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Edge Detection with Laplacian
	edgeDetection := filter.EdgeDetection(grayscale, 5)
	err = io.Save("./img/edgeDetection.png", edgeDetection, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of edgeDetection.png
	edgeDetectionHistogram := histogram.GetGrayHistogram(edgeDetection)
	err = io.Save("./img/histogramImage/edgeDetectionHistogram.png", edgeDetectionHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Fixed Edge Detection
	fixedEdgeDetection := filter.FixedEdgeDetection(grayscale)
	err = io.Save("./img/fixEdgeDetection.png", fixedEdgeDetection, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of fixedEdgeDetection.png
	fixedEdgeDetectionHistogram := histogram.GetGrayHistogram(fixedEdgeDetection)
	err = io.Save("./img/histogramImage/fixedEdgeDetectionHistogram.png", fixedEdgeDetectionHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Sharpen
	sharpen := filter.Sharpen(grayscale, 2)
	err = io.Save("./img/sharpen.png", sharpen, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of sharpen.png
	sharpenHistogram := histogram.GetGrayHistogram(sharpen)
	err = io.Save("./img/histogramImage/sharpenHistogram.png", sharpenHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// FixedSharpen
	fixedSharpen := filter.FixedSharpen(grayscale)
	err = io.Save("./img/fixedSharpen.png", fixedSharpen, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of fixedSharpen.png
	fixedSharpenHistogram := histogram.GetGrayHistogram(fixedSharpen)
	err = io.Save("./img/histogramImage/fixedSharpenHistogram.png", fixedSharpenHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Median
	median := filter.Median(grayscale, 5)
	err = io.Save("./img/median.png", median, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of median.png
	medianHistogram := histogram.GetGrayHistogram(median)
	err = io.Save("./img/histogramImage/medianHistogram.png", medianHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Dilate
	dilate := filter.Dilate(grayscale, 5)
	err = io.Save("./img/dilate.png", dilate, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of dilate.png
	dilateHistogram := histogram.GetGrayHistogram(dilate)
	err = io.Save("./img/histogramImage/dilateHistogram.png", dilateHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Erode
	erode := filter.Erode(grayscale, 5)
	err = io.Save("./img/erode.png", erode, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of erode.png
	erodeHistogram := histogram.GetGrayHistogram(erode)
	err = io.Save("./img/histogramImage/erodeHistogram.png", erodeHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Uniform noise
	uniformNoise := noise.GenerateNoiseImage(original.Bounds().Dx(), original.Bounds().Dy(), noise.Uniform)
	err = io.Save("./img/uniformNoise.png", uniformNoise, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of uniformNoise.png
	uniformNoiseHistogram := histogram.GetGrayHistogram(uniformNoise)
	err = io.Save("./img/histogramImage/uniformNoiseHistogram.png", uniformNoiseHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Binary noise
	binaryNoise := noise.GenerateNoiseImage(original.Bounds().Dx(), original.Bounds().Dy(), noise.Binary)
	err = io.Save("./img/binaryNoise.png", binaryNoise, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of binaryNoise.png
	binaryNoiseHistogram := histogram.GetGrayHistogram(binaryNoise)
	err = io.Save("./img/histogramImage/binaryNoiseHistogram.png", binaryNoiseHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Gaussian noise
	gaussianNoise := noise.GenerateNoiseImage(original.Bounds().Dx(), original.Bounds().Dy(), noise.Gaussian)
	err = io.Save("./img/gaussianNoise.png", gaussianNoise, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of gaussianNoise.png
	gaussianNoiseHistogram := histogram.GetGrayHistogram(gaussianNoise)
	err = io.Save("./img/histogramImage/gaussianNoiseHistogram.png", gaussianNoiseHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Spike noise
	spikeNoise := noise.GenerateNoiseImage(original.Bounds().Dx(), original.Bounds().Dy(), noise.Spike)
	err = io.Save("./img/spikeNoise.png", spikeNoise, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of spikeNoise.png
	spikeNoiseHistogram := histogram.GetGrayHistogram(spikeNoise)
	err = io.Save("./img/histogramImage/spikeNoiseHistogram.png", spikeNoiseHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// grayscale.png with spike noise
	grayscaleWithSpikeNoise := noise.GenerateSpikeNoiseOn(grayscale)
	err = io.Save("./img/grayscaleWithSpikeNoise.png", grayscaleWithSpikeNoise, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of grayscaleWithSpikeNoise.png
	grayscaleWithSpikeNoiseHistogram := histogram.GetGrayHistogram(grayscaleWithSpikeNoise)
	err = io.Save("./img/histogramImage/grayscaleWithSpikeNoiseHistogram.png", grayscaleWithSpikeNoiseHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// smoothed grayscaleWithSpikeNoise.png by median filter
	smoothedGrayscaleWithSpikeNoiseByMedian := filter.Median(grayscaleWithSpikeNoise, 3)
	err = io.Save("./img/smoothedGrayscaleWithSpikeNoiseByMedian.png", smoothedGrayscaleWithSpikeNoiseByMedian, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of smoothedGrayscaleWithSpikeNoiseByMedian.png
	smoothedGrayscaleWithSpikeNoiseByMedianHistogram := histogram.GetGrayHistogram(smoothedGrayscaleWithSpikeNoiseByMedian)
	err = io.Save("./img/histogramImage/smoothedGrayscaleWithSpikeNoiseByMedianHistogram.png", smoothedGrayscaleWithSpikeNoiseByMedianHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Gamma correct
	gammaCorrect := filter.Gamma(grayscale, 0.5)
	err = io.Save("./img/gammaCorrect.png", gammaCorrect, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of gammaCorrect.png
	gammaCorrectHistogram := histogram.GetGrayHistogram(gammaCorrect)
	err = io.Save("./img/histogramImage/gammaCorrectHistogram.png", gammaCorrectHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Histogram equalization
	histogramEqualization := filter.HistogramEqualization(grayscale)
	err = io.Save("./img/histogramEqualization.png", histogramEqualization, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of histogramEqualization.png
	histogramEqualizationHistogram := histogram.GetGrayHistogram(histogramEqualization)
	err = io.Save("./img/histogramImage/histogramEqualizationHistogram.png", histogramEqualizationHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// Cumulative histogram of histogramEqualization.png
	cumulativeHistogramEqualizationHistogram := histogram.GetGrayHistogram(histogramEqualization)
	err = io.Save("./img/histogramImage/cumulativeHistogramEqualizationHistogram.png", cumulativeHistogramEqualizationHistogram.Y.Cumulate().Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
}