package main

import (
	"ImageProcessing/filter"
	"ImageProcessing/histogram"
	"ImageProcessing/io"
	"ImageProcessing/noise"
	"ImageProcessing/transform"
	"ImageProcessing/utils"
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
	// Box Blur
	boxBlur_ := filter.BoxBlur(grayscale, 50)
	err = io.Save("./img/boxBlur_.png", boxBlur_, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of boxBlur.png
	boxBlur_Histogram := histogram.GetGrayHistogram(boxBlur_)
	err = io.Save("./img/histogramImage/boxBlur_Histogram.png", boxBlur_Histogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// Box Blur
	boxBlur__ := filter.BoxBlur(grayscale, 1000)
	err = io.Save("./img/boxBlur__.png", boxBlur__, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of boxBlur.png
	boxBlur__Histogram := histogram.GetGrayHistogram(boxBlur__)
	err = io.Save("./img/histogramImage/boxBlur__Histogram.png", boxBlur__Histogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Gaussian Blur
	gaussianBlur := filter.GaussianBlur(grayscale, 50)
	err = io.Save("./img/gaussianBlur.png", gaussianBlur, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of boxBlur.png
	gaussianBlurHistogram := histogram.GetGrayHistogram(gaussianBlur)
	err = io.Save("./img/histogramImage/gaussianBlurHistogram.png", gaussianBlurHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Fixed Direction Blur
	fixedDirectionBlur := filter.FixedDirectionBlur(grayscale, 30)
	err = io.Save("./img/fixedDirectionBlur.png", fixedDirectionBlur, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of fixedDirectionBlur.png
	fixedDirectionBlurHistogram := histogram.GetGrayHistogram(fixedDirectionBlur)
	err = io.Save("./img/histogramImage/fixedDirectionBlurHistogram.png", fixedDirectionBlurHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Edge Detection with Laplacian
	edgeDetection := filter.EdgeDetection(grayscale, 10)
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
	sharpen := filter.Sharpen(grayscale, 10)
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

	// Median
	medianMedian := filter.Median(filter.Median(filter.Median(filter.Median(grayscale, 20), 20), 20), 20)
	err = io.Save("./img/medianMedian.png", medianMedian, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of median.png
	medianMedianHistogram := histogram.GetGrayHistogram(medianMedian)
	err = io.Save("./img/histogramImage/medianMedianHistogram.png", medianMedianHistogram.Y.Dump(), io.PNGEncoder())
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
	uniformNoise := noise.GenerateNoiseImage(200, 200, noise.Uniform)
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
	binaryNoise := noise.GenerateNoiseImage(200, 200, noise.Binary)
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
	gaussianNoise := noise.GenerateNoiseImage(200, 200, noise.Gaussian)
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
	spikeNoise := noise.GenerateNoiseImage(200, 200, noise.Spike)
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

	// smoothed grayscaleWithSpikeNoise.png by box blur filter
	smoothedGrayscaleWithSpikeNoiseByBoxBlur := filter.BoxBlur(grayscaleWithSpikeNoise, 3)
	err = io.Save("./img/smoothedGrayscaleWithSpikeNoiseByBoxBlur.png", smoothedGrayscaleWithSpikeNoiseByBoxBlur, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of smoothedGrayscaleWithSpikeNoiseByBoxBlur.png
	smoothedGrayscaleWithSpikeNoiseByBoxBlurHistogram := histogram.GetGrayHistogram(smoothedGrayscaleWithSpikeNoiseByBoxBlur)
	err = io.Save("./img/histogramImage/smoothedGrayscaleWithSpikeNoiseByBoxBlurHistogram.png", smoothedGrayscaleWithSpikeNoiseByBoxBlurHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// smoothed grayscaleWithSpikeNoise.png by gaussian blur filter
	smoothedGrayscaleWithSpikeNoiseByGaussianBlur := filter.GaussianBlur(grayscaleWithSpikeNoise, 3)
	err = io.Save("./img/smoothedGrayscaleWithSpikeNoiseByGaussianBlur.png", smoothedGrayscaleWithSpikeNoiseByGaussianBlur, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of smoothedGrayscaleWithSpikeNoiseByGaussianBlur.png
	smoothedGrayscaleWithSpikeNoiseByGaussianBlurHistogram := histogram.GetGrayHistogram(smoothedGrayscaleWithSpikeNoiseByGaussianBlur)
	err = io.Save("./img/histogramImage/smoothedGrayscaleWithSpikeNoiseByBoxBlurHistogram.png", smoothedGrayscaleWithSpikeNoiseByGaussianBlurHistogram.Y.Dump(), io.PNGEncoder())
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

	// Bilateral filter
	bilateralFilter := filter.BilateralFilter(grayscale, 20, 20, 16)
	err = io.Save("./img/bilateralFilter.png", bilateralFilter, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of smoothedGrayscaleWithSpikeNoiseByMedian.png
	bilateralFilterHistogram := histogram.GetGrayHistogram(bilateralFilter)
	err = io.Save("./img/histogramImage/bilateralFilterHistogram.png", bilateralFilterHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	bilateralFilter2 := filter.BilateralFilter(filter.BilateralFilter(filter.BilateralFilter(bilateralFilter, 20, 20, 16), 20, 20, 16), 20, 20, 16)
	err = io.Save("./img/bilateralFilter2.png", bilateralFilter2, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	bilateralFilter3 := filter.BilateralFilter(bilateralFilter2, 5, 12, 16)
	err = io.Save("./img/bilateralFilter3.png", bilateralFilter3, io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// smoothed grayscaleWithSpikeNoise.png by bilateral filter
	smoothedGrayscaleWithSpikeNoiseByBilateralFilter := filter.BilateralFilter(grayscaleWithSpikeNoise, 20, 20, 16)
	err = io.Save("./img/smoothedGrayscaleWithSpikeNoiseByBilateralFilter.png", smoothedGrayscaleWithSpikeNoiseByBilateralFilter, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of smoothedGrayscaleWithSpikeNoiseByBoxBlur.png
	smoothedGrayscaleWithSpikeNoiseByBilateralFilterHistogram := histogram.GetGrayHistogram(smoothedGrayscaleWithSpikeNoiseByBilateralFilter)
	err = io.Save("./img/histogramImage/smoothedGrayscaleWithSpikeNoiseByBoxBlurHistogram.png", smoothedGrayscaleWithSpikeNoiseByBilateralFilterHistogram.Y.Dump(), io.PNGEncoder())
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
	// Gamma correct
	gammaCorrect_ := filter.Gamma(grayscale, 0.5)
	err = io.Save("./img/gammaCorrect_.png", gammaCorrect_, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of gammaCorrect.png
	gammaCorrect_Histogram := histogram.GetGrayHistogram(gammaCorrect_)
	err = io.Save("./img/histogramImage/gammaCorrect_Histogram.png", gammaCorrect_Histogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// Gamma correct
	gammaCorrect__ := filter.Gamma(grayscale, 2)
	err = io.Save("./img/gammaCorrect__.png", gammaCorrect__, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of gammaCorrect.png
	gammaCorrect__Histogram := histogram.GetGrayHistogram(gammaCorrect__)
	err = io.Save("./img/histogramImage/gammaCorrect__Histogram.png", gammaCorrect__Histogram.Y.Dump(), io.PNGEncoder())
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
	// Histogram equalization
	gammaCorrect_histogramEqualization := filter.HistogramEqualization(gammaCorrect_)
	err = io.Save("./img/gammaCorrect_histogramEqualization.png", gammaCorrect_histogramEqualization, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of histogramEqualization.png
	gammaCorrect_histogramEqualizationHistogram := histogram.GetGrayHistogram(gammaCorrect_histogramEqualization)
	err = io.Save("./img/histogramImage/gammaCorrect_histogramEqualizationHistogram.png", gammaCorrect_histogramEqualizationHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// Cumulative histogram of histogramEqualization.png
	gammaCorrect_cumulativeHistogramEqualizationHistogram := histogram.GetGrayHistogram(gammaCorrect_histogramEqualization)
	err = io.Save("./img/histogramImage/gammaCorrect_cumulativeHistogramEqualizationHistogram.png", gammaCorrect_cumulativeHistogramEqualizationHistogram.Y.Cumulate().Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// Histogram equalization
	gammaCorrect__histogramEqualization := filter.HistogramEqualization(gammaCorrect__)
	err = io.Save("./img/gammaCorrect__histogramEqualization.png", gammaCorrect__histogramEqualization, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of histogramEqualization.png
	gammaCorrect__histogramEqualizationHistogram := histogram.GetGrayHistogram(gammaCorrect__histogramEqualization)
	err = io.Save("./img/histogramImage/gammaCorrect__histogramEqualizationHistogram.png", gammaCorrect__histogramEqualizationHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// Cumulative histogram of histogramEqualization.png
	gammaCorrect__cumulativeHistogramEqualizationHistogram := histogram.GetGrayHistogram(gammaCorrect__histogramEqualization)
	err = io.Save("./img/histogramImage/gammaCorrect__cumulativeHistogramEqualizationHistogram.png", gammaCorrect__cumulativeHistogramEqualizationHistogram.Y.Cumulate().Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Halftoning with dithering method
	halftoningWithDitheringMethodBayer := filter.HalftoningWithDitheringMethod(grayscale, filter.Bayer)
	err = io.Save("./img/halftoningWithDitheringMethodBayer.png", halftoningWithDitheringMethodBayer, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of halftoningWithDitheringMethod.png
	halftoningWithDitheringMethodBayerHistogram := histogram.GetGrayHistogram(halftoningWithDitheringMethodBayer)
	err = io.Save("./img/histogramImage/halftoningWithDitheringMethodBayerHistogram.png", halftoningWithDitheringMethodBayerHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Halftoning with dithering method
	halftoningWithDitheringMethodHalfTone := filter.HalftoningWithDitheringMethod(grayscale, filter.HalfTone)
	err = io.Save("./img/halftoningWithDitheringMethodHalfTone.png", halftoningWithDitheringMethodHalfTone, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of halftoningWithDitheringMethod.png
	halftoningWithDitheringMethodHalfToneHistogram := histogram.GetGrayHistogram(halftoningWithDitheringMethodHalfTone)
	err = io.Save("./img/histogramImage/halftoningWithDitheringMethodHalfToneHistogram.png", halftoningWithDitheringMethodHalfToneHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Halftoning with dithering method
	halftoningWithDitheringMethodScrew := filter.HalftoningWithDitheringMethod(grayscale, filter.Screw)
	err = io.Save("./img/halftoningWithDitheringMethodScrew.png", halftoningWithDitheringMethodScrew, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of halftoningWithDitheringMethod.png
	halftoningWithDitheringMethodScrewHistogram := histogram.GetGrayHistogram(halftoningWithDitheringMethodScrew)
	err = io.Save("./img/histogramImage/halftoningWithDitheringMethodScrewHistogram.png", halftoningWithDitheringMethodScrewHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Halftoning with error diffusion method
	halftoningWithErrorDiffusionMethod := filter.HalftoningWithErrorDiffusionMethod(grayscale)
	err = io.Save("./img/halftoningWithErrorDiffusionMethod.png", halftoningWithErrorDiffusionMethod, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of halftoningWithDitheringMethod.png
	halftoningWithErrorDiffusionMethodHistogram := histogram.GetGrayHistogram(halftoningWithErrorDiffusionMethod)
	err = io.Save("./img/histogramImage/halftoningWithErrorDiffusionMethodHistogram.png", halftoningWithErrorDiffusionMethodHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// open keystoneEffectSample.png
	keystoneEffectSampleOriginal, err := io.Open("./img/keystoneEffectSample.png")
	if err != nil {
		panic(err)
	}

	// convert original.png to grayscale
	keystoneEffectSampleOriginalGrayscale := filter.Grayscale(keystoneEffectSampleOriginal)
	err = io.Save("./img/keystoneEffectSampleOriginalGrayscale.png", keystoneEffectSampleOriginalGrayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of grayscale.png
	keystoneEffectSampleOriginalGrayscaleHistogram := histogram.GetGrayHistogram(keystoneEffectSampleOriginalGrayscale)
	err = io.Save("./img/histogramImage/keystoneEffectSampleOriginalGrayscaleHistogram.png", keystoneEffectSampleOriginalGrayscaleHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// default
	ua, va := 836.0, 1152.0
	ub, vb := 1224.0, 1788.0
	uc, vc := 2596.0, 1188.0
	ud, vd := 2020.0, 768.0
	x, y := 800.0, 600.0
	a := [][]float64{
		{x, y, 0, 0, 0, 0, -uc * x, -uc * y, uc - ua},
		{0, y, 0, 0, 0, 0, 0, -ub * y, ub - ua},
		{0, 0, 1, 0, 0, 0, 0, 0, ua},
		{0, 0, 0, x, y, 0, -vc * x, -vc * y, vc - va},
		{0, 0, 0, 0, y, 0, 0, -vb * y, vb - va},
		{0, 0, 0, 0, 0, 1, 0, 0, va},
		{x, 0, 0, 0, 0, 0, -ud * x, 0, ud - ua},
		{0, 0, 0, x, 0, 0, -vd * x, 0, vd - va},
	}
	utils.GaussElimination(&a)
	// for i := range a {
	// 	fmt.Println(a[i])
	// }
	planeProjectionMatrix := [][]float64{
		{a[0][len(a[0])-1], a[1][len(a[0])-1], a[2][len(a[0])-1]},
		{a[3][len(a[0])-1], a[4][len(a[0])-1], a[5][len(a[0])-1]},
		{a[6][len(a[0])-1], a[7][len(a[0])-1], 1},
	}
	// for i := range planeProjectionMatrix {
	// 	fmt.Println(planeProjectionMatrix[i])
	// }
	// fmt.Println("A")
	// fmt.Println(planeProjectionMatrix[0][0]*0 + planeProjectionMatrix[0][1]*0 + planeProjectionMatrix[0][2]*1)
	// fmt.Println(planeProjectionMatrix[1][0]*0 + planeProjectionMatrix[1][1]*0 + planeProjectionMatrix[1][2]*1)
	// fmt.Println(planeProjectionMatrix[2][0]*0 + planeProjectionMatrix[2][1]*0 + planeProjectionMatrix[2][2]*1)
	// fmt.Println()
	// fmt.Println("B")
	// fmt.Println(planeProjectionMatrix[0][0]*0 + planeProjectionMatrix[0][1]*y + planeProjectionMatrix[0][2]*1)
	// fmt.Println(planeProjectionMatrix[1][0]*0 + planeProjectionMatrix[1][1]*y + planeProjectionMatrix[1][2]*1)
	// fmt.Println(planeProjectionMatrix[2][0]*0 + planeProjectionMatrix[2][1]*y + planeProjectionMatrix[2][2]*1)
	// fmt.Println()
	// fmt.Println("C")
	// fmt.Println(planeProjectionMatrix[0][0]*x + planeProjectionMatrix[0][1]*y + planeProjectionMatrix[0][2]*1)
	// fmt.Println(planeProjectionMatrix[1][0]*x + planeProjectionMatrix[1][1]*y + planeProjectionMatrix[1][2]*1)
	// fmt.Println(planeProjectionMatrix[2][0]*x + planeProjectionMatrix[2][1]*y + planeProjectionMatrix[2][2]*1)
	// fmt.Println()
	// fmt.Println("D")
	// fmt.Println(planeProjectionMatrix[0][0]*x + planeProjectionMatrix[0][1]*0 + planeProjectionMatrix[0][2]*1)
	// fmt.Println(planeProjectionMatrix[1][0]*x + planeProjectionMatrix[1][1]*0 + planeProjectionMatrix[1][2]*1)
	// fmt.Println(planeProjectionMatrix[2][0]*x + planeProjectionMatrix[2][1]*0 + planeProjectionMatrix[2][2]*1)

	// transform keystoneEffectSampleOriginalGrayscale.png
	keystoneEffect := transform.KeystoneEffect(keystoneEffectSampleOriginalGrayscale, &planeProjectionMatrix, int(x), int(y))
	err = io.Save("./img/keystoneEffect.png", keystoneEffect, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of keystoneEffect.png
	keystoneEffectHistogram := histogram.GetGrayHistogram(keystoneEffect)
	err = io.Save("./img/histogramImage/keystoneEffectHistogram.png", keystoneEffectHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// open C2_0_original.png and C2_2_original.png
	C2_0_original, err := io.Open("./img/C2_0_original.png")
	if err != nil {
		panic(err)
	}
	C2_1_original, err := io.Open("./img/C2_1_original.png")
	if err != nil {
		panic(err)
	}

	// convert C2_0_original.png to grayscale
	C2_0_grayscale := filter.Grayscale(C2_0_original)
	err = io.Save("./img/C2_0_grayscale.png", C2_0_grayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of C2_0_grayscale.png
	C2_0_grayscaleHistogram := histogram.GetGrayHistogram(C2_0_grayscale)
	err = io.Save("./img/histogramImage/C2_0_grayscaleHistogram.png", C2_0_grayscaleHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// convert C2_1_grayscale.png to grayscale
	C2_1_grayscale := filter.Grayscale(C2_1_original)
	err = io.Save("./img/C2_1_grayscale.png", C2_1_grayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of C2_1_grayscale.png
	C2_1_grayscaleHistogram := histogram.GetGrayHistogram(C2_1_grayscale)
	err = io.Save("./img/histogramImage/C2_1_grayscaleHistogram.png", C2_1_grayscaleHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// C2_0
	ua, va = 1030.0, 970.0
	ub, vb = 1046.0, 1710.0
	uc, vc = 2560.0, 1964.0
	ud, vd = 2604.0, 1002.0
	x, y = 600.0, 200.0
	a = [][]float64{
		{x, y, 0, 0, 0, 0, -uc * x, -uc * y, uc - ua},
		{0, y, 0, 0, 0, 0, 0, -ub * y, ub - ua},
		{0, 0, 1, 0, 0, 0, 0, 0, ua},
		{0, 0, 0, x, y, 0, -vc * x, -vc * y, vc - va},
		{0, 0, 0, 0, y, 0, 0, -vb * y, vb - va},
		{0, 0, 0, 0, 0, 1, 0, 0, va},
		{x, 0, 0, 0, 0, 0, -ud * x, 0, ud - ua},
		{0, 0, 0, x, 0, 0, -vd * x, 0, vd - va},
	}
	utils.GaussElimination(&a)
	planeProjectionMatrix = [][]float64{
		{a[0][len(a[0])-1], a[1][len(a[0])-1], a[2][len(a[0])-1]},
		{a[3][len(a[0])-1], a[4][len(a[0])-1], a[5][len(a[0])-1]},
		{a[6][len(a[0])-1], a[7][len(a[0])-1], 1},
	}
	// transform keystoneEffectC2_0_grayscale.png
	keystoneEffect_C2_0_grayscale := transform.KeystoneEffect(C2_0_grayscale, &planeProjectionMatrix, int(x), int(y))
	err = io.Save("./img/keystoneEffect_C2_0_grayscale.png", keystoneEffect_C2_0_grayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of keystoneEffect.png
	keystoneEffect_C2_0_grayscaleHistogram := histogram.GetGrayHistogram(keystoneEffect_C2_0_grayscale)
	err = io.Save("./img/histogramImage/keystoneEffect_C2_0_grayscaleHistogram.png", keystoneEffect_C2_0_grayscaleHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// C2_1
	ua, va = 350.0, 1084.0
	ub, vb = 894.0, 2095.0
	uc, vc = 2723.0, 1020.0
	ud, vd = 2140.0, 318.0
	x, y = 600.0, 200.0
	a = [][]float64{
		{x, y, 0, 0, 0, 0, -uc * x, -uc * y, uc - ua},
		{0, y, 0, 0, 0, 0, 0, -ub * y, ub - ua},
		{0, 0, 1, 0, 0, 0, 0, 0, ua},
		{0, 0, 0, x, y, 0, -vc * x, -vc * y, vc - va},
		{0, 0, 0, 0, y, 0, 0, -vb * y, vb - va},
		{0, 0, 0, 0, 0, 1, 0, 0, va},
		{x, 0, 0, 0, 0, 0, -ud * x, 0, ud - ua},
		{0, 0, 0, x, 0, 0, -vd * x, 0, vd - va},
	}
	utils.GaussElimination(&a)
	planeProjectionMatrix = [][]float64{
		{a[0][len(a[0])-1], a[1][len(a[0])-1], a[2][len(a[0])-1]},
		{a[3][len(a[0])-1], a[4][len(a[0])-1], a[5][len(a[0])-1]},
		{a[6][len(a[0])-1], a[7][len(a[0])-1], 1},
	}
	// transform keystoneEffectC2_1_grayscale.png
	keystoneEffect_C2_1_grayscale := transform.KeystoneEffect(C2_1_grayscale, &planeProjectionMatrix, int(x), int(y))
	err = io.Save("./img/keystoneEffect_C2_1_grayscale.png", keystoneEffect_C2_1_grayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of keystoneEffect.png
	keystoneEffect_C2_1_grayscaleHistogram := histogram.GetGrayHistogram(keystoneEffect_C2_1_grayscale)
	err = io.Save("./img/histogramImage/keystoneEffect_C2_1_grayscaleHistogram.png", keystoneEffect_C2_1_grayscaleHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// get C2_result.png
	C2_result := transform.GetUnion(keystoneEffect_C2_0_grayscale, keystoneEffect_C2_1_grayscale)
	err = io.Save("./img/C2_result.png", C2_result, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of keystoneEffect.png
	C2_resultHistogram := histogram.GetGrayHistogram(C2_result)
	err = io.Save("./img/histogramImage/C2_resultHistogram.png", C2_resultHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// open C3_1.png
	C3_0Original, err := io.Open("./img/C3_0.png")
	if err != nil {
		panic(err)
	}

	// convert C3_1.png to grayscale
	C3_0Grayscale := filter.Grayscale(C3_0Original)
	err = io.Save("./img/C3_0Grayscale.png", C3_0Grayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of C3_1.png
	C3_0GrayscaleHistogram := histogram.GetGrayHistogram(C3_0Grayscale)
	err = io.Save("./img/histogramImage/C3_0GrayscaleHistogram.png", C3_0GrayscaleHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// open C3_2.png
	C3_1Original, err := io.Open("./img/C3_1.png")
	if err != nil {
		panic(err)
	}

	// convert C3_2.png to grayscale
	C3_1Grayscale := filter.Grayscale(C3_1Original)
	err = io.Save("./img/C3_1Grayscale.png", C3_1Grayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of C3_2.png
	C3_1GrayscaleHistogram := histogram.GetGrayHistogram(C3_1Grayscale)
	err = io.Save("./img/histogramImage/C3_1GrayscaleHistogram.png", C3_1GrayscaleHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// open C3_0.png
	C3_2Original, err := io.Open("./img/C3_2.png")
	if err != nil {
		panic(err)
	}

	// convert C3_0.png to grayscale
	C3_2Grayscale := filter.Grayscale(C3_2Original)
	err = io.Save("./img/C3_2Grayscale.png", C3_2Grayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of C3_2.png
	C3_2GrayscaleHistogram := histogram.GetGrayHistogram(C3_2Grayscale)
	err = io.Save("./img/histogramImage/C3_2GrayscaleHistogram.png", C3_2GrayscaleHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// C3_1をC3_0の座標系に変換
	// C3_1
	ua01, va01 := 564.0, 1742.0
	ub01, vb01 := 574.0, 1872.0
	uc01, vc01 := 928.0, 1843.0
	ud01, vd01 := 854.0, 1724.0
	// C3_0
	xa01, ya01 := 2312.0, 1751.0
	xb01, yb01 := 2324.0, 1877.0
	xc01, yc01 := 2680.0, 1901.0
	xd01, yd01 := 2594.0, 1766.0

	a01 := [][]float64{
		{ua01, va01, 1, 0, 0, 0, -xa01*ua01, -xa01*va01, xa01},
		{0, 0, 0, ua01, va01, 1, -ya01*ua01, -ya01*va01, ya01},
		{ub01, vb01, 1, 0, 0, 0, -xb01*ub01, -xb01*vb01, xb01},
		{0, 0, 0, ub01, vb01, 1, -yb01*ub01, -yb01*vb01, yb01},
		{uc01, vc01, 1, 0, 0, 0, -xc01*uc01, -xc01*vc01, xc01},
		{0, 0, 0, uc01, vc01, 1, -yc01*uc01, -yc01*vc01, yc01},
		{ud01, vd01, 1, 0, 0, 0, -xd01*ud01, -xd01*vd01, xd01},
		{0, 0, 0, ud01, vd01, 1, -yd01*ud01, -yd01*vd01, yd01},
	}

	utils.GaussElimination(&a01)
	planeProjectionMatrix01 := [][]float64{
		{a01[0][len(a01[0])-1], a01[1][len(a01[0])-1], a01[2][len(a01[0])-1]},
		{a01[3][len(a01[0])-1], a01[4][len(a01[0])-1], a01[5][len(a01[0])-1]},
		{a01[6][len(a01[0])-1], a01[7][len(a01[0])-1], 1},
	}

	// C3_2
	ua02, va02 := 2016.0, 1824.0
	ub02, vb02 := 2016.0, 1996.0
	uc02, vc02 := 2550.0, 2083.0
	ud02, vd02 := 2220.0, 1867.0
	// C3_0
	xa02, ya02 := 595.0, 1880.0
	xb02, yb02 := 592.0, 2063.0
	xc02, yc02 := 1135.0, 2069.0
	xd02, yd02 := 815.0, 1899.0
	a02 := [][]float64{
		{ua02, va02, 1, 0, 0, 0, -xa02*ua02, -xa02*va02, xa02},
		{0, 0, 0, ua02, va02, 1, -ya02*ua02, -ya02*va02, ya02},
		{ub02, vb02, 1, 0, 0, 0, -xb02*ub02, -xb02*vb02, xb02},
		{0, 0, 0, ub02, vb02, 1, -yb02*ub02, -yb02*vb02, yb02},
		{uc02, vc02, 1, 0, 0, 0, -xc02*uc02, -xc02*vc02, xc02},
		{0, 0, 0, uc02, vc02, 1, -yc02*uc02, -yc02*vc02, yc02},
		{ud02, vd02, 1, 0, 0, 0, -xd02*ud02, -xd02*vd02, xd02},
		{0, 0, 0, ud02, vd02, 1, -yd02*ud02, -yd02*vd02, yd02},
	}

	utils.GaussElimination(&a02)
	planeProjectionMatrix02 := [][]float64{
		{a02[0][len(a02[0])-1], a02[1][len(a02[0])-1], a02[2][len(a02[0])-1]},
		{a02[3][len(a02[0])-1], a02[4][len(a02[0])-1], a02[5][len(a02[0])-1]},
		{a02[6][len(a02[0])-1], a02[7][len(a02[0])-1], 1},
	}
	C3_Stitched_012 := transform.Stitching(C3_0Grayscale, C3_1Grayscale, C3_2Grayscale,  2000, &planeProjectionMatrix01, &planeProjectionMatrix02)
	err = io.Save("./img/C3_Stitched012.png", C3_Stitched_012, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	C3_Stitched_012Histogram := histogram.GetGrayHistogram(C3_Stitched_012)
	err = io.Save("./img/histogramImage/C3_Stitched_012Histogram.png", C3_Stitched_012Histogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
}
